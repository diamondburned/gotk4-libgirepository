// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_operation_preview_get_type()), F: marshalPrintOperationPreview},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable. This
// interface is a subset of the interface PrintOperationPreview.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()

	GotPageSize(context PrintContext, pageSetup PageSetup)
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool

	Ready(context PrintContext)
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with @preview.
	//
	// A custom iprint preview should use this function in its ::expose handler
	// to render the currently selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	RenderPage(pageNr int)
}

type PrintOperationPreview interface {
	gextras.Objector
	PrintOperationPreviewOverrider
}

// printOperationPreview implements the PrintOperationPreview interface.
type printOperationPreview struct {
	gextras.Objector
}

var _ PrintOperationPreview = (*printOperationPreview)(nil)

// WrapPrintOperationPreview wraps a GObject to a type that implements interface
// PrintOperationPreview. It is primarily used internally.
func WrapPrintOperationPreview(obj *externglib.Object) PrintOperationPreview {
	return PrintOperationPreview{
		Objector: obj,
	}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintOperationPreview(obj), nil
}

// EndPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (p printOperationPreview) EndPreview() {
	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))

	C.gtk_print_operation_preview_end_preview(_arg0)
}

// IsSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
func (p printOperationPreview) IsSelected(pageNr int) bool {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))
	_arg1 = C.gint(pageNr)

	var _cret C.gboolean // in

	_cret = C.gtk_print_operation_preview_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenderPage renders a page to the preview, using the print context that
// was passed to the PrintOperation::preview handler together with @preview.
//
// A custom iprint preview should use this function in its ::expose handler
// to render the currently selected page.
//
// Note that this function requires a suitable cairo context to be
// associated with the print context.
func (p printOperationPreview) RenderPage(pageNr int) {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))
	_arg1 = C.gint(pageNr)

	C.gtk_print_operation_preview_render_page(_arg0, _arg1)
}
