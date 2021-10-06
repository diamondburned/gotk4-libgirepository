// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontexter},
	})
}

// IMMulticontext: GtkIMMulticontext is input method supporting multiple,
// switchable input methods.
//
// Text widgets such as GtkText or GtkTextView use a GtkIMMultiContext to
// implement their im-module property for switching between different input
// methods.
type IMMulticontext struct {
	IMContext
}

func wrapIMMulticontext(obj *externglib.Object) *IMMulticontext {
	return &IMMulticontext{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMMulticontexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIMMulticontext(obj), nil
}

// NewIMMulticontext creates a new GtkIMMulticontext.
func NewIMMulticontext() *IMMulticontext {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_multicontext_new()

	var _imMulticontext *IMMulticontext // out

	_imMulticontext = wrapIMMulticontext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imMulticontext
}

// ContextID gets the id of the currently active delegate of the context.
func (context *IMMulticontext) ContextID() string {
	var _arg0 *C.GtkIMMulticontext // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_multicontext_get_context_id(_arg0)
	runtime.KeepAlive(context)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetContextID sets the context id for context.
//
// This causes the currently active delegate of context to be replaced by the
// delegate corresponding to the new context id.
//
// The function takes the following parameters:
//
//    - contextId: id to use.
//
func (context *IMMulticontext) SetContextID(contextId string) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contextId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_multicontext_set_context_id(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(contextId)
}
