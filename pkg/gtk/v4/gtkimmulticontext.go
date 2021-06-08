// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontext},
	})
}

// IMMulticontext: `GtkIMMulticontext` is input method supporting multiple,
// switchable input methods.
//
// Text widgets such as `GtkText` or `GtkTextView` use a `GtkIMMultiContext` to
// implement their `im-module` property for switching between different input
// methods.
type IMMulticontext interface {
	IMContext

	// ContextID gets the id of the currently active delegate of the @context.
	ContextID() string
	// SetContextID sets the context id for @context.
	//
	// This causes the currently active delegate of @context to be replaced by
	// the delegate corresponding to the new context id.
	SetContextID(contextID string)
}

// imMulticontext implements the IMMulticontext interface.
type imMulticontext struct {
	IMContext
}

var _ IMMulticontext = (*imMulticontext)(nil)

// WrapIMMulticontext wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMMulticontext(obj *externglib.Object) IMMulticontext {
	return IMMulticontext{
		IMContext: WrapIMContext(obj),
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMMulticontext(obj), nil
}

// NewIMMulticontext constructs a class IMMulticontext.
func NewIMMulticontext() IMMulticontext {
	cret := new(C.GtkIMMulticontext)
	var goret IMMulticontext

	cret = C.gtk_im_multicontext_new()

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(IMMulticontext)

	return goret
}

// ContextID gets the id of the currently active delegate of the @context.
func (c imMulticontext) ContextID() string {
	var arg0 *C.GtkIMMulticontext

	arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))

	var cret *C.char
	var goret string

	cret = C.gtk_im_multicontext_get_context_id(arg0)

	goret = C.GoString(cret)

	return goret
}

// SetContextID sets the context id for @context.
//
// This causes the currently active delegate of @context to be replaced by
// the delegate corresponding to the new context id.
func (c imMulticontext) SetContextID(contextID string) {
	var arg0 *C.GtkIMMulticontext
	var arg1 *C.char

	arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(contextID))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_im_multicontext_set_context_id(arg0, arg1)
}
