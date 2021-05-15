package girgen

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path"
	"strconv"
	"text/template"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
	"github.com/fatih/color"
)

func newGoTemplate(block string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"PascalToGo":               PascalToGo,
		"SnakeToGo":                SnakeToGo,
		"FirstChar":                FirstChar,
		"GoDoc":                    GoDoc,
		"CommentReflowLinesIndent": CommentReflowLinesIndent,
	})
	t = template.Must(t.Parse(block))
	return t
}

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	// KnownTypes contains a list of type checks that return true if the given
	// type matches a known type.
	KnownTypes []func(string) bool
	// NoMarshalPkgs contains a list of namespace package names that don't have
	// marshalers generated. GLib is by default in this list, because it
	// shouldn't have any marshalers.
	NoMarshalPkgs []string

	Repos      gir.Repositories
	RootModule string

	logger *log.Logger
	color  bool
}

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, root string) *Generator {
	return &Generator{
		Repos:      repos,
		RootModule: root,
	}
}

// WithLogger sets the generator's logger.
func (g *Generator) WithLogger(logger *log.Logger, color bool) {
	g.logger = logger
	g.color = color
}

type logLevel uint8

const (
	logInfo logLevel = iota
	logSummary
	logWarn
	logError
)

func (lvl logLevel) prefix() string {
	switch lvl {
	case logInfo:
		return "info:"
	case logSummary:
		return "summary:"
	case logWarn:
		return "warning:"
	case logError:
		return "error:"
	default:
		return ""
	}
}

func (lvl logLevel) colorf(f string, v ...interface{}) string {
	switch lvl {
	case logInfo:
		return color.HiBlueString(f, v...)
	case logSummary:
		return color.HiGreenString(f, v...)
	case logWarn:
		return color.HiYellowString(f, v...)
	case logError:
		return color.HiRedString(f, v...)
	default:
		return fmt.Sprintf(f, v...)
	}
}

func (g *Generator) logln(level logLevel, v ...interface{}) {
	if g.logger == nil {
		return
	}

	prefix := level.prefix()
	if g.color {
		prefix = level.colorf(prefix)
	}

	v = append([]interface{}{prefix}, v...)
	g.logger.Println(v...)
}

func (g *Generator) ImportPath(pkgPath string) string {
	return gir.ImportPath(g.RootModule, pkgPath)
}

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace string) *NamespaceGenerator {
	res := g.Repos.FindNamespace(namespace)
	if res == nil {
		return nil
	}

	buf := bytes.Buffer{}
	buf.Grow(4 * 1024 * 1024) // 4MB

	return &NamespaceGenerator{
		pen:  pen.New(&buf),
		body: &buf,

		imports: map[string]string{},
		pkgPath: g.ImportPath(gir.GoNamespace(res.Namespace)),
		gen:     g,
		current: res,
	}
}

// NamespaceGenerator is a generator for a specific namespace.
type NamespaceGenerator struct {
	pen  *pen.Pen
	body *bytes.Buffer

	pkgPath string            // package name
	imports map[string]string // optional alias value

	gen     *Generator
	current *gir.NamespaceFindResult
}

// Generate generates the current namespace into the given writer.
func (ng *NamespaceGenerator) Generate(w io.Writer) error {
	ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
	ng.addImport("github.com/diamondburned/gotk4/internal/callback")

	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateInit()
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateFuncs()
	ng.generateRecords()
	ng.generateClasses()

	if err := ng.pen.Flush(); err != nil {
		return err
	}

	pen := pen.New(w)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.Line()

	pen.Words("package", ng.PackageName())
	pen.Line()

	if len(ng.imports) > 0 {
		pen.Words("import (")
		for imp, alias := range ng.imports {
			// Only use the import alias if it's provided and does not match the
			// base name of the import path for idiomaticity.
			if alias != "" && alias != path.Base(imp) {
				pen.Words(alias, "", strconv.Quote(imp))
			} else {
				pen.Words(strconv.Quote(imp))
			}
		}
		pen.Block(")")
		pen.Line()
	}

	pkgs := []string{"// #cgo pkg-config:"}
	for _, pkg := range ng.current.Repository.Packages {
		pkgs = append(pkgs, pkg.Name)
	}

	pen.Words(pkgs...)
	pen.Words("// #cgo CFLAGS: -Wno-deprecated-declarations")
	for _, cIncl := range ng.current.Repository.CIncludes {
		pen.Words("// #include", fmt.Sprintf("<%s>", cIncl.Name))
	}
	// pen.Words("// extern void callbackDelete(gpointer);")
	pen.Words(`import "C"`)
	pen.Line()

	// TODO: detect when this is needed.
	// pen.Words("//export callbackDelete")
	// pen.Block(`func callbackDelete(ptr C.gpointer) { callback.Delete(uintptr(ptr)) }`)

	pen.Write(ng.body.Bytes())

	return pen.Flush()
}

// PackageName returns the current namespace's package name.
func (ng *NamespaceGenerator) PackageName() string {
	return gir.GoNamespace(ng.current.Namespace)
}

// Namespace returns the generator's namespace.
func (ng *NamespaceGenerator) Namespace() *gir.Namespace {
	return ng.current.Namespace
}

// Repository returns the generator's repository.
func (ng *NamespaceGenerator) Repository() *gir.PkgRepository {
	return ng.current.Repository
}

func (ng *NamespaceGenerator) logln(level logLevel, v ...interface{}) {
	prefix := []interface{}{"package", ng.current.Namespace.Name + ":"}
	prefix = append(prefix, v...)

	ng.gen.logln(level, prefix...)
}

func (ng *NamespaceGenerator) warnUnknownType(typ string) {
	ng.logln(logWarn, "unknown gir type", strconv.Quote(typ))
}

func (ng *NamespaceGenerator) addImport(pkgPath string) {
	ng.addImportAlias(pkgPath, "")
}

func (ng *NamespaceGenerator) addImportAlias(pkgPath, alias string) {
	_, ok := ng.imports[pkgPath]
	if ok {
		return
	}

	ng.imports[pkgPath] = alias
}
