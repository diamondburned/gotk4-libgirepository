// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey: `GtkEventControllerKey` is an event controller that
// provides access to key events.
type EventControllerKey interface {
	EventController

	// Forward forwards the current event of this @controller to a @widget.
	//
	// This function can only be used in handlers for the
	// [signal@Gtk.EventControllerKey::key-pressed],
	// [signal@Gtk.EventControllerKey::key-released] or
	// [signal@Gtk.EventControllerKey::modifiers] signals.
	Forward(widget Widget) bool
	// Group gets the key group of the current event of this @controller.
	//
	// See [method@Gdk.KeyEvent.get_layout].
	Group() uint
	// SetImContext sets the input method context of the key @controller.
	SetImContext(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey class.
type eventControllerKey struct {
	EventController
}

var _ EventControllerKey = (*eventControllerKey)(nil)

// WrapEventControllerKey wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return eventControllerKey{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

// Forward forwards the current event of this @controller to a @widget.
//
// This function can only be used in handlers for the
// [signal@Gtk.EventControllerKey::key-pressed],
// [signal@Gtk.EventControllerKey::key-released] or
// [signal@Gtk.EventControllerKey::modifiers] signals.
func (c eventControllerKey) Forward(widget Widget) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Group gets the key group of the current event of this @controller.
//
// See [method@Gdk.KeyEvent.get_layout].
func (c eventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var _cret C.guint // in

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// SetImContext sets the input method context of the key @controller.
func (c eventControllerKey) SetImContext(imContext IMContext) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
