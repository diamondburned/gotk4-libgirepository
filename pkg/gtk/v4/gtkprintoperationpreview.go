// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_PrintOperationPreview_ConnectReady(gpointer, void*, guintptr);
// extern void _gotk4_gtk4_PrintOperationPreview_ConnectGotPageSize(gpointer, void*, void*, guintptr);
import "C"

// GType values.
var (
	GTypePrintOperationPreview = coreglib.Type(girepository.MustFind("Gtk", "PrintOperationPreview").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePrintOperationPreview, F: marshalPrintOperationPreview},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable.
type PrintOperationPreviewOverrider interface {
}

// PrintOperationPreview: GtkPrintOperationPreview is the interface that is used
// to implement print preview.
//
// A GtkPrintOperationPreview object is passed to the
// gtk.PrintOperation::preview signal by gtk.PrintOperation.
//
// PrintOperationPreview wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PrintOperationPreview struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PrintOperationPreview)(nil)
)

// PrintOperationPreviewer describes PrintOperationPreview's interface methods.
type PrintOperationPreviewer interface {
	coreglib.Objector

	basePrintOperationPreview() *PrintOperationPreview
}

var _ PrintOperationPreviewer = (*PrintOperationPreview)(nil)

func ifaceInitPrintOperationPreviewer(gifacePtr, data C.gpointer) {
}

func wrapPrintOperationPreview(obj *coreglib.Object) *PrintOperationPreview {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	return wrapPrintOperationPreview(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *PrintOperationPreview) basePrintOperationPreview() *PrintOperationPreview {
	return v
}

// BasePrintOperationPreview returns the underlying base object.
func BasePrintOperationPreview(obj PrintOperationPreviewer) *PrintOperationPreview {
	return obj.basePrintOperationPreview()
}

// ConnectGotPageSize is emitted once for each page that gets rendered to the
// preview.
//
// A handler for this signal should update the context according to page_setup
// and set up a suitable cairo context, using
// gtk.PrintContext.SetCairoContext().
func (v *PrintOperationPreview) ConnectGotPageSize(f func(context *PrintContext, pageSetup *PageSetup)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "got-page-size", false, unsafe.Pointer(C._gotk4_gtk4_PrintOperationPreview_ConnectGotPageSize), f)
}

// ConnectReady signal gets emitted once per preview operation, before the first
// page is rendered.
//
// A handler for this signal can be used for setup tasks.
func (v *PrintOperationPreview) ConnectReady(f func(context *PrintContext)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "ready", false, unsafe.Pointer(C._gotk4_gtk4_PrintOperationPreview_ConnectReady), f)
}

// PrintOperationPreviewIface: instance of this type is always passed by
// reference.
type PrintOperationPreviewIface struct {
	*printOperationPreviewIface
}

// printOperationPreviewIface is the struct that's finalized.
type printOperationPreviewIface struct {
	native unsafe.Pointer
}

var GIRInfoPrintOperationPreviewIface = girepository.MustFind("Gtk", "PrintOperationPreviewIface")
