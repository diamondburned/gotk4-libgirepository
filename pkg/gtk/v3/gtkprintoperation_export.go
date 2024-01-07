// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_PageSetupDoneFunc
func _gotk4_gtk3_PageSetupDoneFunc(arg1 *C.void, arg2 C.gpointer) {
	var fn PageSetupDoneFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PageSetupDoneFunc)
	}

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg1)))

	fn(_pageSetup)
}

//export _gotk4_gtk3_PrintOperation_ConnectBeginPrint
func _gotk4_gtk3_PrintOperation_ConnectBeginPrint(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
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

//export _gotk4_gtk3_PrintOperation_ConnectCreateCustomWidget
func _gotk4_gtk3_PrintOperation_ConnectCreateCustomWidget(arg0 C.gpointer, arg1 C.guintptr) (cret C.GObject) {
	var f func() (object *coreglib.Object)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (object *coreglib.Object))
	}

	object := f()

	var _ *coreglib.Object

	cret = *(*C.GObject)(unsafe.Pointer(object.Native()))

	return cret
}

//export _gotk4_gtk3_PrintOperation_ConnectCustomWidgetApply
func _gotk4_gtk3_PrintOperation_ConnectCustomWidgetApply(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

//export _gotk4_gtk3_PrintOperation_ConnectDrawPage
func _gotk4_gtk3_PrintOperation_ConnectDrawPage(arg0 C.gpointer, arg1 *C.void, arg2 C.gint, arg3 C.guintptr) {
	var f func(context *PrintContext, pageNr int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext, pageNr int))
	}

	var _context *PrintContext // out
	var _pageNr int            // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageNr = int(arg2)

	f(_context, _pageNr)
}

//export _gotk4_gtk3_PrintOperation_ConnectEndPrint
func _gotk4_gtk3_PrintOperation_ConnectEndPrint(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
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

//export _gotk4_gtk3_PrintOperation_ConnectPaginate
func _gotk4_gtk3_PrintOperation_ConnectPaginate(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) (cret C.gboolean) {
	var f func(context *PrintContext) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext) (ok bool))
	}

	var _context *PrintContext // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))

	ok := f(_context)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_PrintOperation_ConnectPreview
func _gotk4_gtk3_PrintOperation_ConnectPreview(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.guintptr) (cret C.gboolean) {
	var f func(preview PrintOperationPreviewer, context *PrintContext, parent *Window) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(preview PrintOperationPreviewer, context *PrintContext, parent *Window) (ok bool))
	}

	var _preview PrintOperationPreviewer // out
	var _context *PrintContext           // out
	var _parent *Window                  // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.PrintOperationPreviewer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(PrintOperationPreviewer)
			return ok
		})
		rv, ok := casted.(PrintOperationPreviewer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.PrintOperationPreviewer")
		}
		_preview = rv
	}
	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg2)))
	if arg3 != nil {
		_parent = wrapWindow(coreglib.Take(unsafe.Pointer(arg3)))
	}

	ok := f(_preview, _context, _parent)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_PrintOperation_ConnectRequestPageSetup
func _gotk4_gtk3_PrintOperation_ConnectRequestPageSetup(arg0 C.gpointer, arg1 *C.void, arg2 C.gint, arg3 *C.void, arg4 C.guintptr) {
	var f func(context *PrintContext, pageNr int, setup *PageSetup)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext, pageNr int, setup *PageSetup))
	}

	var _context *PrintContext // out
	var _pageNr int            // out
	var _setup *PageSetup      // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageNr = int(arg2)
	_setup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg3)))

	f(_context, _pageNr, _setup)
}

//export _gotk4_gtk3_PrintOperation_ConnectStatusChanged
func _gotk4_gtk3_PrintOperation_ConnectStatusChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_PrintOperation_ConnectUpdateCustomWidget
func _gotk4_gtk3_PrintOperation_ConnectUpdateCustomWidget(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.guintptr) {
	var f func(widget Widgetter, setup *PageSetup, settings *PrintSettings)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter, setup *PageSetup, settings *PrintSettings))
	}

	var _widget Widgetter        // out
	var _setup *PageSetup        // out
	var _settings *PrintSettings // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}
	_setup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg2)))
	_settings = wrapPrintSettings(coreglib.Take(unsafe.Pointer(arg3)))

	f(_widget, _setup, _settings)
}
