// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_root_get_type()), F: marshalRoot},
	})
}

// Root: `GtkRoot` is the interface implemented by all widgets that can act as a
// toplevel widget.
//
// The root widget takes care of providing the connection to the windowing
// system and manages layout, drawing and event delivery for its widget
// hierarchy.
//
// The obvious example of a `GtkRoot` is `GtkWindow`.
//
// To get the display to which a `GtkRoot` belongs, use
// [method@Gtk.Root.get_display].
//
// `GtkRoot` also maintains the location of keyboard focus inside its widget
// hierarchy, with [method@Gtk.Root.set_focus] and [method@Gtk.Root.get_focus].
type Root interface {
	NativeWidget

	// Display returns the display that this `GtkRoot` is on.
	Display() gdk.Display
	// Focus retrieves the current focused widget within the root.
	//
	// Note that this is the widget that would have the focus if the root is
	// active; if the root is not focused then `gtk_widget_has_focus (widget)`
	// will be false for the widget.
	Focus() Widget
	// SetFocus: if @focus is not the current focus widget, and is focusable,
	// sets it as the focus widget for the root.
	//
	// If @focus is nil, unsets the focus widget for the root.
	//
	// To set the focus to a particular widget in the root, it is usually more
	// convenient to use [method@Gtk.Widget.grab_focus] instead of this
	// function.
	SetFocus(focus Widget)
}

// root implements the Root interface.
type root struct {
	Native
	Widget
}

var _ Root = (*root)(nil)

// WrapRoot wraps a GObject to a type that implements interface
// Root. It is primarily used internally.
func WrapRoot(obj *externglib.Object) Root {
	return Root{
		Native: WrapNative(obj),
		Widget: WrapWidget(obj),
	}
}

func marshalRoot(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRoot(obj), nil
}

// Display returns the display that this `GtkRoot` is on.
func (s root) Display() gdk.Display {
	var arg0 *C.GtkRoot

	arg0 = (*C.GtkRoot)(unsafe.Pointer(s.Native()))

	var cret *C.GdkDisplay
	var goret gdk.Display

	cret = C.gtk_root_get_display(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Display)

	return goret
}

// Focus retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus if the root is
// active; if the root is not focused then `gtk_widget_has_focus (widget)`
// will be false for the widget.
func (s root) Focus() Widget {
	var arg0 *C.GtkRoot

	arg0 = (*C.GtkRoot)(unsafe.Pointer(s.Native()))

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_root_get_focus(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// SetFocus: if @focus is not the current focus widget, and is focusable,
// sets it as the focus widget for the root.
//
// If @focus is nil, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually more
// convenient to use [method@Gtk.Widget.grab_focus] instead of this
// function.
func (s root) SetFocus(focus Widget) {
	var arg0 *C.GtkRoot
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkRoot)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(focus.Native()))

	C.gtk_root_set_focus(arg0, arg1)
}
