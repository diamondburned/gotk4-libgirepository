// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_controller_motion_get_type()), F: marshalDropControllerMotioner},
	})
}

// DropControllerMotion: GtkDropControllerMotion is an event controller tracking
// the pointer during Drag-and-Drop operations.
//
// It is modeled after gtk.EventControllerMotion so if you have used that, this
// should feel really familiar.
//
// This controller is not able to accept drops, use gtk.DropTarget for that
// purpose.
type DropControllerMotion struct {
	EventController
}

func wrapDropControllerMotion(obj *externglib.Object) *DropControllerMotion {
	return &DropControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropControllerMotioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDropControllerMotion(obj), nil
}

// NewDropControllerMotion creates a new event controller that will handle
// pointer motion events during drag and drop.
func NewDropControllerMotion() *DropControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_drop_controller_motion_new()

	var _dropControllerMotion *DropControllerMotion // out

	_dropControllerMotion = wrapDropControllerMotion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dropControllerMotion
}

// ContainsPointer returns if a Drag-and-Drop operation is within the widget
// self or one of its children.
func (self *DropControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_contains_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Drop returns the GdkDrop of a current Drag-and-Drop operation over the widget
// of self.
func (self *DropControllerMotion) Drop() gdk.Dropper {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret *C.GdkDrop                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_get_drop(_arg0)
	runtime.KeepAlive(self)

	var _drop gdk.Dropper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gdk.Dropper)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Dropper")
			}
			_drop = rv
		}
	}

	return _drop
}

// IsPointer returns if a Drag-and-Drop operation is within the widget self, not
// one of its children.
func (self *DropControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_is_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectEnter signals that the pointer has entered the widget.
func (self *DropControllerMotion) ConnectEnter(f func(x, y float64)) externglib.SignalHandle {
	return self.Connect("enter", f)
}

// ConnectLeave signals that the pointer has left the widget.
func (self *DropControllerMotion) ConnectLeave(f func()) externglib.SignalHandle {
	return self.Connect("leave", f)
}

// ConnectMotion: emitted when the pointer moves inside the widget.
func (self *DropControllerMotion) ConnectMotion(f func(x, y float64)) externglib.SignalHandle {
	return self.Connect("motion", f)
}
