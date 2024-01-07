// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_Seat_ConnectToolRemoved(gpointer, void*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectToolAdded(gpointer, void*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectDeviceRemoved(gpointer, void*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectDeviceAdded(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeSeat = coreglib.Type(girepository.MustFind("Gdk", "Seat").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeat, F: marshalSeat},
	})
}

// Seat object represents a collection of input devices that belong to a user.
type Seat struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Seat)(nil)
)

// Seater describes types inherited from class Seat.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Seater interface {
	coreglib.Objector
	baseSeat() *Seat
}

var _ Seater = (*Seat)(nil)

func wrapSeat(obj *coreglib.Object) *Seat {
	return &Seat{
		Object: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	return wrapSeat(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Seat) baseSeat() *Seat {
	return v
}

// BaseSeat returns the underlying base object.
func BaseSeat(obj Seater) *Seat {
	return obj.baseSeat()
}

// ConnectDeviceAdded signal is emitted when a new input device is related to
// this seat.
func (v *Seat) ConnectDeviceAdded(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "device-added", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectDeviceAdded), f)
}

// ConnectDeviceRemoved signal is emitted when an input device is removed (e.g.
// unplugged).
func (v *Seat) ConnectDeviceRemoved(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "device-removed", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectDeviceRemoved), f)
}

// ConnectToolAdded signal is emitted whenever a new tool is made known to the
// seat. The tool may later be assigned to a device (i.e. on proximity with a
// tablet). The device will emit the Device::tool-changed signal accordingly.
//
// A same tool may be used by several devices.
func (v *Seat) ConnectToolAdded(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "tool-added", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectToolAdded), f)
}

// ConnectToolRemoved: this signal is emitted whenever a tool is no longer known
// to this seat.
func (v *Seat) ConnectToolRemoved(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "tool-removed", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectToolRemoved), f)
}
