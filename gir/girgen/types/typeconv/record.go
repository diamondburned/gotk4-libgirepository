package typeconv

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// RecordHasFree returns the free/unref method if it has one.
func RecordHasFree(record *gir.Record) *gir.Method {
	for _, name := range []string{"unref", "free", "destroy"} {
		if m := types.FindMethodName(record.Methods, name); m != nil {
			return m
		}
	}
	return nil
}

// RecordHasUnref returns the unref method if it has one.
func RecordHasUnref(record *gir.Record) *gir.Method {
	// TODO: runtime link mode support
	return types.FindMethodName(record.Methods, "unref")
}

// RecordHasRef returns the ref method if it has one.
func RecordHasRef(record *gir.Record) *gir.Method {
	// TODO: runtime link mode support
	return types.FindMethodName(record.Methods, "ref")
}

type RecordPrintedMethod struct {
	Header file.Header
	Body   string
}

// RecordPrintFreeMethod generates a call with 1 argument for either free or
// unref. If method is nil, then a C.free call is generated. Value is assumed to
// be an unsafe.Pointer.
//
// The caller must import girepository.Argument manually.
func RecordPrintFree(gen types.FileGenerator, parent *types.Resolved, value string) file.Value[string] {
	rec := parent.Extern.Type.(*gir.Record)

	free := RecordHasFree(rec)
	if free == nil {
		return file.Value[string]{
			Value: fmt.Sprintf("C.free(%s)", value),
		}
	}

	return recordPrintMethod(gen, parent, free, value)
}

func recordPrintMethod(gen types.FileGenerator, parent *types.Resolved, method *gir.Method, value string) file.Value[string] {
	rec := parent.Extern.Type.(*gir.Record)

	// TODO: refactor thoughts: should typeconv and girgen/callable be combined?
	// typeconv has to generate function calling code in some cases for freeing,
	// and we have various different ways of doing that scattered throughout the
	// program. It would be far better to have 1 place of doing them all.

	var v file.Value[string]

	switch gen.LinkMode() {
	case types.RuntimeLinkMode:
		girInfo := fmt.Sprintf("GIRInfo%s", strcases.PascalToGo(rec.Name))
		if !parent.NeedsNamespace(gen.Namespace()) {
			parent.ImportPubl(gen, &v.Header)
			girInfo = fmt.Sprintf("%s.%s", parent.Extern.Namespace.Name, girInfo)
		}

		p := pen.NewBlock()
		p.Linef("var args [1]girepository.Argument")
		p.Linef("*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(%s)", value)
		p.Linef("%s.InvokeRecordMethod(%q, args[:], nil)", girInfo, method.Name)
		v.Value = p.String()

	case types.DynamicLinkMode:
		v.Value = fmt.Sprintf(
			"C.%s((%s)(%s))",
			method.CIdentifier,
			types.AnyTypeCGo(method.Parameters.InstanceParameter.AnyType),
			value)

	default:
		panic("unreachable")
	}

	return v
}
