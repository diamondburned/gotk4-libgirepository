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
// #include <glib-object.h>
import "C"

// GTypeIMMulticontext returns the GType for the type IMMulticontext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeIMMulticontext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "IMMulticontext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalIMMulticontext)
	return gtype
}

// IMMulticontextOverrider contains methods that are overridable.
type IMMulticontextOverrider interface {
}

// IMMulticontext: GtkIMMulticontext is input method supporting multiple,
// switchable input methods.
//
// Text widgets such as GtkText or GtkTextView use a GtkIMMultiContext to
// implement their im-module property for switching between different input
// methods.
type IMMulticontext struct {
	_ [0]func() // equal guard
	IMContext
}

var (
	_ IMContexter = (*IMMulticontext)(nil)
)

func classInitIMMulticontexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIMMulticontext(obj *coreglib.Object) *IMMulticontext {
	return &IMMulticontext{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	return wrapIMMulticontext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMMulticontext creates a new GtkIMMulticontext.
//
// The function returns the following values:
//
//    - imMulticontext: new GtkIMMulticontext.
//
func NewIMMulticontext() *IMMulticontext {
	_gret := girepository.MustFind("Gtk", "IMMulticontext").InvokeMethod("new_IMMulticontext", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _imMulticontext *IMMulticontext // out

	_imMulticontext = wrapIMMulticontext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imMulticontext
}

// ContextID gets the id of the currently active delegate of the context.
//
// The function returns the following values:
//
//    - utf8: id of the currently active delegate.
//
func (context *IMMulticontext) ContextID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gtk", "IMMulticontext").InvokeMethod("get_context_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(contextId)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "IMMulticontext").InvokeMethod("set_context_id", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(contextId)
}
