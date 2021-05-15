package girgen

import (
	"fmt"
	"log"
	"path"
	"strings"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
	"golang.org/x/sync/singleflight"
)

// ResolvedType is a resolved type from a given gir.Type.
type ResolvedType struct {
	// either or
	Extern  *ExternType // optional
	Builtin *string     // optional

	Import  string // Full import path.
	Package string // Package name, also import alias.

	Parent string // GIR Type, optional
	CType  string
	Ptr    uint8
}

// ExternType is an externally resolved type.
type ExternType struct {
	Result *gir.TypeFindResult
}

var (
	typeCache  sync.Map
	typeFlight singleflight.Group
)

func countPtrs(typ gir.Type, result *gir.TypeFindResult) uint8 {
	ptr := uint8(strings.Count(typ.CType, "*"))

	// Edge case: interfaces must not be pointers. We should still sometimes
	// allow for pointers to interfaces, if needed, but this likely won't work.
	if result != nil && result.Interface != nil && ptr > 0 {
		ptr--
	}
	// Edge case: a string is a gchar*, so we don't need a pointer.
	if typ.Name == "utf8" && ptr > 0 {
		ptr--
	}

	return ptr
}

// builtinType is a convenient function to make a new resolvedType.
func builtinType(imp, typ string, girType gir.Type) *ResolvedType {
	if imp != "" {
		// Create the actual type if there's an import path.
		typ = path.Base(imp) + "." + typ
	}

	return &ResolvedType{
		Builtin: &typ,
		Import:  imp,
		Package: path.Base(imp),
		CType:   girType.CType,
		Ptr:     countPtrs(girType, nil),
	}
}

// externGLibType returns an external GLib type from gotk3.
func externGLibType(glibType string, typ gir.Type) *ResolvedType {
	ptrs := strings.Count(glibType, "*")
	glibType = strings.Repeat("*", ptrs) + "externglib." + strings.TrimPrefix(glibType, "*")

	return &ResolvedType{
		Builtin: &glibType,
		Import:  "github.com/gotk3/gotk3/glib",
		Package: "externglib",
		CType:   typ.CType,
		Ptr:     uint8(ptrs),
	}
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(gen *Generator, typ gir.Type, result *gir.TypeFindResult) *ResolvedType {
	var parent string
	if result.Class != nil {
		parent = result.Class.Parent
	}

	pkg := gir.GoNamespace(result.Namespace)

	return &ResolvedType{
		Extern: &ExternType{
			Result: result,
		},
		Import:  gen.ImportPath(pkg),
		Package: pkg,
		Parent:  parent,
		CType:   typ.CType,
		Ptr:     countPtrs(typ, result),
	}
}

// NeedsNamespace returns true if the returned Go type needs a namespace to be
// referenced properly.
func (typ *ResolvedType) NeedsNamespace(current *gir.NamespaceFindResult) bool {
	if typ.Extern == nil {
		return false
	}

	return !typ.Extern.Result.Eq(current)
}

// GoType formats the Go type.
func (typ *ResolvedType) GoType(needsNamespace bool) string {
	if typ.Builtin != nil {
		return *typ.Builtin
	}

	name, _ := typ.Extern.Result.Info()
	name = PascalToGo(name)

	ptr := strings.Repeat("*", int(typ.Ptr))

	if !needsNamespace {
		return ptr + name
	}
	return ptr + typ.Package + "." + name
}

// CGoType returns the CGo type.
func (typ *ResolvedType) CGoType() string {
	ptr := strings.Count(typ.CType, "*")
	val := strings.TrimSuffix(typ.CType, "*")

	return strings.Repeat("*", ptr) + "C." + val
}

// movePtr moves the same number of pointers from the given orig string into
// another string.
func movePtr(orig, into string) string {
	ptr := strings.Count(orig, "*")
	return strings.Repeat("*", ptr) + into
}

// GoType returns the generated Go type of the given resolved type.
func (ng *NamespaceGenerator) GoType(resolved *ResolvedType) string {
	return resolved.GoType(resolved.NeedsNamespace(ng.current))
}

// arrayType generates the Go type signature for the given array.
func (ng *NamespaceGenerator) resolveArrayType(array gir.Array) (string, bool) {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := ng.ResolveAnyType(array.AnyType)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// ResolveAnyType generates the Go type signature for the AnyType union. An
// empty string returned is an invalid type.
func (ng *NamespaceGenerator) ResolveAnyType(any gir.AnyType) (string, bool) {
	switch {
	case any.Array != nil:
		return ng.resolveArrayType(*any.Array)
	case any.Type != nil:
		return ng.ResolveToGoType(*any.Type)
	}

	// Probably varargs, ignore because Cgo.
	return "", false
}

// ResolveToGoType is a convenient function that wraps around ResolveType and
// returns the Go type.
func (ng *NamespaceGenerator) ResolveToGoType(typ gir.Type) (string, bool) {
	resolved := ng.ResolveType(typ)
	if resolved == nil {
		return "", false
	}

	return ng.GoType(resolved), true
}

// ResolveTypeName resolves the given GIR type name. The resolved type will
// always have no pointer.
func (ng *NamespaceGenerator) ResolveTypeName(girType string) *ResolvedType {
	var cType string

	// FindType is cached, so we can afford to do this.
	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, girType)
	if result != nil {
		// Use the CType result ONLY. The returned Name from Info does NOT have
		// the namespace prepended.
		_, cType = result.Info()
	}

	return ng.ResolveType(gir.Type{
		Name:  girType,
		CType: cType,
	})
}

// ResolveType resolves the given type from the GIR type field. It returns nil
// if the type is not known. It does not recursively traverse the type.
func (ng *NamespaceGenerator) ResolveType(typ gir.Type) *ResolvedType {
	if typ.CType == "" {
		// Some aliases may not have a CType, for some reason. Create our own.
		// Mutating typ here is fine, since it's a copy.
		typ.CType = "gir_" + ng.current.Namespace.Name + "." + typ.Name
	}

	v, ok := typeCache.Load(typ.CType)
	if ok {
		return v.(*ResolvedType)
	}

	// Cache miss. Use singleflight to ensure we're not looking up multiple
	// versions of the same type to prevent cache stampede.
	v, _, _ = typeFlight.Do(typ.CType, func() (interface{}, error) {
		resolved := ng.resolveTypeUncached(typ)
		if resolved != nil {
			// Save into the cache within the singleflight callback.
			typeCache.Store(typ.CType, resolved)

			// Add the import in the same singleflight callback, but only if the
			// namespace is not the current one.
			if resolved.Import != "" && resolved.Import != ng.pkgPath {
				ng.addImportAlias(resolved.Import, resolved.Package)
			}
		}

		return resolved, nil
	})

	// may be a non-nil interface to a nil pointer.
	return v.(*ResolvedType)
}

// girPrimitiveGo maps the given GIR primitive type to a Go primitive type.
var girPrimitiveGo = map[string]string{
	"none":        "",
	"gboolean":    "bool",
	"gfloat":      "float32",
	"gdouble":     "float64",
	"long double": "float64",
	"gint":        "int",
	"gssize":      "int",
	"gint8":       "int8",
	"gint16":      "int16",
	"gshort":      "int16",
	"gint32":      "int32",
	"glong":       "int32",
	"int32":       "int32",
	"gint64":      "int64",
	"guint":       "uint",
	"gsize":       "uint",
	"guchar":      "byte",
	"gchar":       "byte",
	"guint8":      "uint8",
	"guint16":     "uint16",
	"gushort":     "uint16",
	"guint32":     "uint32",
	"gulong":      "uint32",
	"gunichar":    "uint32",
	"guint64":     "uint64",
	"utf8":        "string",
	"filename":    "string",
}

func (ng *NamespaceGenerator) resolveTypeUncached(typ gir.Type) *ResolvedType {
	if typ.Name == "" {
		ng.logln(logWarn, "empty gir type", typ)
		return nil
	}

	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		return builtinType("", prim, typ)
	}

	// Resolve the unknown namespace that is GLib and primitive types.
	switch typ.Name {
	// TODO: ignore field
	// TODO: aaaaaaaaaaaaaaaaaaaaaaa
	case "gpointer":
		return builtinType("unsafe", "Pointer", typ)
	case "GLib.DestroyNotify", "DestroyNotify": // This should be handled externally.
		return builtinType("unsafe", "Pointer", typ)
	case "GType":
		return externGLibType("Type", typ)
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return externGLibType("*Value", typ)
	case "GObject.Object":
		return externGLibType("*Object", typ)
	case "GObject.Closure":
		return externGLibType("*Closure", typ)
	case "GObject.Callback":
		// Callback is a special func(Any) Any type, so we treat it as
		// interface{} similarly to object.Connect(). We can use glib's Closure
		// APIs to parse this interface{}.
		return builtinType("", "interface{}", typ)

	case "GObject.InitiallyUnowned":
		return externGLibType("InitiallyUnowned", typ)

	case "va_list":
		// CGo cannot handle variadic argument lists.
		return nil

	// We don't know what these types translates to.
	case "GObject.TypeModule":
		return nil
	case "GObject.ParamSpec": // this is deprecated
		return nil
	case "GObject.Parameter": // also deprecated I think
		return nil
	// TODO: Find a way to map EnumValue type.
	// TODO: Add _full function support.
	case "GObject.EnumValue":
		return nil
	}

	// Types that aren't in the switch tree that match any of these patterns are
	// types that must be in the switch tree, so them not being in there is a
	// bug.
	for _, check := range ng.gen.KnownTypes {
		if check(typ.Name) {
			log.Fatalf("missing gir type %s in the type tree\n", typ.Name)
		}
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		ng.warnUnknownType(typ.Name)
		return nil
	}

	return typeFromResult(ng.gen, typ, result)
}

// TODO: GoTypeConverter converts Go types to C with GIR type.

// AnyTypeConverter returns Go code that is the conversion from the given C
// value type to its respective Go value type. An empty string returned is
// invalid.
func (ng *NamespaceGenerator) AnyTypeConverter(value, target string, any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return ng.arrayConverter(value, target, *any.Array)
	case any.Type != nil:
		return ng.typeConverter(value, target, *any.Type)
	}

	// Ignore VarArgs.
	return ""
}

func directCallOrCreate(value, target, typ string, create bool) string {
	var op = " = "
	if create {
		op = " := "
	}

	return target + op + typ + "(" + value + ")"
}

func (ng *NamespaceGenerator) arrayConverter(value, target string, array gir.Array) string {
	if array.Type != nil {
		ng.gen.logln(logWarn, "skipping nested array", array)
		return ""
	}

	// Detect if the underlying is a compatible Go primitive type. If it is,
	// then we can directly cast a fixed-size array.
	if array.Type != nil {
		if primitiveGo, ok := girPrimitiveGo[array.Type.Name]; ok {
			return fmt.Sprintf("%s = ([%d]%s)(%s)", target, array.FixedSize, primitiveGo, value)
		}
	}

	innerType := ng.ResolveType(*array.Type)
	if innerType == nil {
		return ""
	}
	innerCGoType := innerType.CGoType()

	// Generate a type converter from "src" to "dst" variables.
	innerConv := ng.typeConverter("src", "dst", *array.Type)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		b.Linef("var a [%d]%s", array.FixedSize, innerType)
		// TODO: nested array support
		b.Linef("cArray := ([%d]%s)(%s)", array.FixedSize, array.Type.CType, value)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := cArray[i]")
		b.Linef("  dst := &a[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		return "" // TODO

	case array.Name == "GLib.Array": // treat as Go array
		b.Linef("var length uintptr")
		b.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&length))", value)
		b.Linef("a := make([]%s, length)", innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		b.Linef("  dst := &a[i]")
		// TODO: nested array support
		b.Linef("  " + innerConv)
		b.Linef("}")

	default: // null-terminated
		// Scan for the length.
		b.Linef("var length uint")
		b.Linef("for p := unsafe.Pointer(%s); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {", value)
		b.Linef("  length++")
		b.Linef("}")

		b.EmptyLine()

		// Preallocate the slice.
		b.Linef("a := make([]%s, length)", innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + 1))", innerCGoType, value)
		b.Linef("  dst := &a[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")
	}

	return b.String()
}

func (ng *NamespaceGenerator) typeConverter(value, target string, typ gir.Type) string {
	return ng._typeConverter(value, target, typ, false)
}

func (ng *NamespaceGenerator) _typeConverter(value, target string, typ gir.Type, create bool) string {
	// Resolve primitive types.
	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			return directCallOrCreate(value, target, "C.GoString", create)
		default:
			return directCallOrCreate(value, target, prim, create)
		}
	}

	// Resolve special-case GLib types.
	switch typ.Name {
	case "gpointer":
		return directCallOrCreate(value, target, "unsafe.Pointer", create)
	case "GLib.DestroyNotify", "DestroyNotify":
		return ""
	case "GType":
		return ""
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return ""
	case "GObject.Object":
		return directCallOrCreate(value, target, "glib.Take", create)
	case "GObject.Closure":
		return ""
	case "GObject.InitiallyUnowned":
		return ""
	case "GObject.Callback":
		// TODO: When is this ever needed? How do I even do this?
		return ""
	case "va_list":
		// CGo cannot handle variadic argument lists.
		return ""
	case "GObject.EnumValue", "GObject.TypeModule", "GObject.ParamSpec", "GObject.Parameter":
		// Refer to ResolveType.
		return ""
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		// Probably already warned.
		return ""
	}

	resolved := typeFromResult(ng.gen, typ, result)
	goType := resolved.GoType(false)

	// Resolve alias.
	if result.Alias != nil {
		var b pen.Block
		b.Line(ng._typeConverter(value, "tmp", result.Alias.Type, true))
		b.EmptyLine()
		b.Line(directCallOrCreate("tmp", target, goType, false))
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return directCallOrCreate(value, target, goType, false)
	}

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.
	// TODO: interfaces should be wrapped by an unexported type.

	// Assume the wrap function. This so far works for classes and records.
	// TODO: handle wrap functions from another package.
	return directCallOrCreate(value, target, "wrap"+goType, false)
}
