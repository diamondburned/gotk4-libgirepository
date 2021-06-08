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
	// ImContext gets the input method context of the key @controller.
	ImContext() IMContext
	// SetImContext sets the input method context of the key @controller.
	SetImContext(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey interface.
type eventControllerKey struct {
	EventController
}

var _ EventControllerKey = (*eventControllerKey)(nil)

// WrapEventControllerKey wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return EventControllerKey{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

// NewEventControllerKey constructs a class EventControllerKey.
func NewEventControllerKey() EventControllerKey {
	cret := new(C.GtkEventControllerKey)
	var goret EventControllerKey

	cret = C.gtk_event_controller_key_new()

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(EventControllerKey)

	return goret
}

// Forward forwards the current event of this @controller to a @widget.
//
// This function can only be used in handlers for the
// [signal@Gtk.EventControllerKey::key-pressed],
// [signal@Gtk.EventControllerKey::key-released] or
// [signal@Gtk.EventControllerKey::modifiers] signals.
func (c eventControllerKey) Forward(widget Widget) bool {
	var arg0 *C.GtkEventControllerKey
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_event_controller_key_forward(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Group gets the key group of the current event of this @controller.
//
// See [method@Gdk.KeyEvent.get_layout].
func (c eventControllerKey) Group() uint {
	var arg0 *C.GtkEventControllerKey

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_event_controller_key_get_group(arg0)

	goret = uint(cret)

	return goret
}

// ImContext gets the input method context of the key @controller.
func (c eventControllerKey) ImContext() IMContext {
	var arg0 *C.GtkEventControllerKey

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var cret *C.GtkIMContext
	var goret IMContext

	cret = C.gtk_event_controller_key_get_im_context(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(IMContext)

	return goret
}

// SetImContext sets the input method context of the key @controller.
func (c eventControllerKey) SetImContext(imContext IMContext) {
	var arg0 *C.GtkEventControllerKey
	var arg1 *C.GtkIMContext

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(arg0, arg1)
}
