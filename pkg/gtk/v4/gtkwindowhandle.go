// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeWindowHandle = coreglib.Type(C.gtk_window_handle_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowHandle, F: marshalWindowHandle},
	})
}

// WindowHandleOverrider contains methods that are overridable.
type WindowHandleOverrider interface {
}

// WindowHandle: GtkWindowHandle is a titlebar area widget.
//
// When added into a window, it can be dragged to move the window, and handles
// right click, double click and middle click as expected of a titlebar.
//
//
// CSS nodes
//
// GtkWindowHandle has a single CSS node with the name windowhandle.
//
//
// Accessibility
//
// GtkWindowHandle uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowHandle struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*WindowHandle)(nil)
)

func initClassWindowHandle(gclass unsafe.Pointer, goval any) {
}

func wrapWindowHandle(obj *coreglib.Object) *WindowHandle {
	return &WindowHandle{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalWindowHandle(p uintptr) (interface{}, error) {
	return wrapWindowHandle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowHandle creates a new GtkWindowHandle.
//
// The function returns the following values:
//
//    - windowHandle: new GtkWindowHandle.
//
func NewWindowHandle() *WindowHandle {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_window_handle_new()

	var _windowHandle *WindowHandle // out

	_windowHandle = wrapWindowHandle(coreglib.Take(unsafe.Pointer(_cret)))

	return _windowHandle
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self.
//
func (self *WindowHandle) Child() Widgetter {
	var _arg0 *C.GtkWindowHandle // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_window_handle_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _widget
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *WindowHandle) SetChild(child Widgetter) {
	var _arg0 *C.GtkWindowHandle // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.gtk_window_handle_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
