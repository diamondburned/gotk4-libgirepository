package girgen

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// Go to C type conversions.

// GoCConversion describes the information needed to generate Go code to convert
// Go types to C types.
type GoCConversion struct {
	Value  string
	Target string
	Type   gir.AnyType
	Owner  gir.TransferOwnership

	// ArgAt is used for array and closure generation.
	ArgAt ArgAtFunc
}

// func (ng *NamespaceGenerator) GoCConverter(value, target string, any gir.AnyType) string {

// }

// func (ng *NamespaceGenerator) gocArrayConverter(value, target string, array gir.Array) string {

// }

// func (ng *NamespaceGenerator) gocTypeConverter(value, target string, typ gir.Type) string {

// }

func (ng *NamespaceGenerator) _gocTypeConverter(conv GoCConversion, typ gir.Type, create bool) string {
	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef(directCallOrCreate(conv.Value, conv.Target, "C.CString", create))
			p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Value)
			return p.String()
		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return directCallOrCreate(conv.Value, conv.Target, "gextras.Cbool", create)
		default:
			return directCallOrCreate(conv.Value, conv.Target, "C."+typ.CType, create)
		}
	}

	switch typ.Name {
	case "gpointer":
		ng.addImport("github.com/diamondburned/gotk4/internal/box")

		b := pen.NewBlock()
		b.Linef("id := box.Assign(box.Boxed, %s)", conv.Value)
		b.Linef("%s = C.gpointer(id)", conv.Target)
		return b.String()

	case "GLib.DestroyNotify", "DestroyNotify":
		// This should never be called, because the caller should never see a
		// DestroyNotify, so there's no use to convert from Go to C.
		ng.logln(logError, "unexpected DestroyNotify conversion from Go to C")
		return ""
	}

	// TODO
	return ""
}