// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeat},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities int

const (
	// SeatCapabilitiesNone: no input capabilities
	SeatCapabilitiesNone SeatCapabilities = 0
	// SeatCapabilitiesPointer: the seat has a pointer (e.g. mouse)
	SeatCapabilitiesPointer SeatCapabilities = 1
	// SeatCapabilitiesTouch: the seat has touchscreen(s) attached
	SeatCapabilitiesTouch SeatCapabilities = 2
	// SeatCapabilitiesTabletStylus: the seat has drawing tablet(s) attached
	SeatCapabilitiesTabletStylus SeatCapabilities = 4
	// SeatCapabilitiesKeyboard: the seat has keyboard(s) attached
	SeatCapabilitiesKeyboard SeatCapabilities = 8
	// SeatCapabilitiesTabletPad: the seat has drawing tablet pad(s) attached
	SeatCapabilitiesTabletPad SeatCapabilities = 16
	// SeatCapabilitiesAllPointing: the union of all pointing capabilities
	SeatCapabilitiesAllPointing SeatCapabilities = 7
	// SeatCapabilitiesAll: the union of all capabilities
	SeatCapabilitiesAll SeatCapabilities = 15
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Seat: the `GdkSeat` object represents a collection of input devices that
// belong to a user.
type Seat interface {
	gextras.Objector

	// Capabilities returns the capabilities this `GdkSeat` currently has.
	Capabilities() SeatCapabilities
	// Display returns the `GdkDisplay` this seat belongs to.
	Display() Display
	// Keyboard returns the device that routes keyboard events.
	Keyboard() Device
	// Pointer returns the device that routes pointer events.
	Pointer() Device
}

// seat implements the Seat class.
type seat struct {
	gextras.Objector
}

var _ Seat = (*seat)(nil)

// WrapSeat wraps a GObject to the right type. It is
// primarily used internally.
func WrapSeat(obj *externglib.Object) Seat {
	return seat{
		Objector: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeat(obj), nil
}

// Capabilities returns the capabilities this `GdkSeat` currently has.
func (s seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = SeatCapabilities(_cret)

	return _seatCapabilities
}

// Display returns the `GdkDisplay` this seat belongs to.
func (s seat) Display() Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

// Keyboard returns the device that routes keyboard events.
func (s seat) Keyboard() Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

// Pointer returns the device that routes pointer events.
func (s seat) Pointer() Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}
