// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_PrintOperationPreviewIface_is_selected(void*, gint);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_end_preview(void*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_got_page_size(void*, void*, void*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_ready(void*, void*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_render_page(void*, gint);
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize(gpointer, void*, void*, guintptr);
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectReady(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtkprintoperationpreview.go.
var GTypePrintOperationPreview = coreglib.Type(C.gtk_print_operation_preview_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePrintOperationPreview, F: marshalPrintOperationPreview},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()
	// The function takes the following parameters:
	//
	//    - context
	//    - pageSetup
	//
	GotPageSize(context *PrintContext, pageSetup *PageSetup)
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	//
	// The function takes the following parameters:
	//
	//    - pageNr: page number.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the page has been selected for printing.
	//
	IsSelected(pageNr int32) bool
	// The function takes the following parameters:
	//
	Ready(context *PrintContext)
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	//
	// A custom iprint preview should use this function in its ::expose handler
	// to render the currently selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	//
	// The function takes the following parameters:
	//
	//    - pageNr: page to render.
	//
	RenderPage(pageNr int32)
}

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

	// EndPreview ends a preview.
	EndPreview()
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int32) bool
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	RenderPage(pageNr int32)

	// Got-page-size signal is emitted once for each page that gets rendered to
	// the preview.
	ConnectGotPageSize(func(context *PrintContext, pageSetup *PageSetup)) coreglib.SignalHandle
	// Ready signal gets emitted once per preview operation, before the first
	// page is rendered.
	ConnectReady(func(context *PrintContext)) coreglib.SignalHandle
}

var _ PrintOperationPreviewer = (*PrintOperationPreview)(nil)

func ifaceInitPrintOperationPreviewer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkPrintOperationPreviewIface)(unsafe.Pointer(gifacePtr))
	iface.end_preview = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_end_preview)
	iface.got_page_size = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_got_page_size)
	iface.is_selected = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_is_selected)
	iface.ready = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_ready)
	iface.render_page = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_render_page)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_end_preview
func _gotk4_gtk3_PrintOperationPreviewIface_end_preview(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	iface.EndPreview()
}

//export _gotk4_gtk3_PrintOperationPreviewIface_got_page_size
func _gotk4_gtk3_PrintOperationPreviewIface_got_page_size(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _context *PrintContext // out
	var _pageSetup *PageSetup  // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageSetup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg2)))

	iface.GotPageSize(_context, _pageSetup)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_is_selected
func _gotk4_gtk3_PrintOperationPreviewIface_is_selected(arg0 *C.void, arg1 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _pageNr int32 // out

	_pageNr = int32(arg1)

	ok := iface.IsSelected(_pageNr)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_PrintOperationPreviewIface_ready
func _gotk4_gtk3_PrintOperationPreviewIface_ready(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _context *PrintContext // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))

	iface.Ready(_context)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_render_page
func _gotk4_gtk3_PrintOperationPreviewIface_render_page(arg0 *C.void, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _pageNr int32 // out

	_pageNr = int32(arg1)

	iface.RenderPage(_pageNr)
}

func wrapPrintOperationPreview(obj *coreglib.Object) *PrintOperationPreview {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	return wrapPrintOperationPreview(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize
func _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(context *PrintContext, pageSetup *PageSetup)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext, pageSetup *PageSetup))
	}

	var _context *PrintContext // out
	var _pageSetup *PageSetup  // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageSetup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg2)))

	f(_context, _pageSetup)
}

// ConnectGotPageSize signal is emitted once for each page that gets rendered to
// the preview.
//
// A handler for this signal should update the context according to page_setup
// and set up a suitable cairo context, using
// gtk_print_context_set_cairo_context().
func (preview *PrintOperationPreview) ConnectGotPageSize(f func(context *PrintContext, pageSetup *PageSetup)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(preview, "got-page-size", false, unsafe.Pointer(C._gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize), f)
}

//export _gotk4_gtk3_PrintOperationPreview_ConnectReady
func _gotk4_gtk3_PrintOperationPreview_ConnectReady(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(context *PrintContext)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext))
	}

	var _context *PrintContext // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))

	f(_context)
}

// ConnectReady signal gets emitted once per preview operation, before the first
// page is rendered.
//
// A handler for this signal can be used for setup tasks.
func (preview *PrintOperationPreview) ConnectReady(f func(context *PrintContext)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(preview, "ready", false, unsafe.Pointer(C._gotk4_gtk3_PrintOperationPreview_ConnectReady), f)
}

// EndPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (preview *PrintOperationPreview) EndPreview() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(preview).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	runtime.KeepAlive(preview)
}

// IsSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
//
// The function takes the following parameters:
//
//    - pageNr: page number.
//
// The function returns the following values:
//
//    - ok: TRUE if the page has been selected for printing.
//
func (preview *PrintOperationPreview) IsSelected(pageNr int32) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenderPage renders a page to the preview, using the print context that was
// passed to the PrintOperation::preview handler together with preview.
//
// A custom iprint preview should use this function in its ::expose handler to
// render the currently selected page.
//
// Note that this function requires a suitable cairo context to be associated
// with the print context.
//
// The function takes the following parameters:
//
//    - pageNr: page to render.
//
func (preview *PrintOperationPreview) RenderPage(pageNr int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)
}
