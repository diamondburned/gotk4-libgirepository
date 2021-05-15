// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk"
	"github.com/gotk3/gotk3/glib"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Objects/Classes
	})
}

type WaylandDevice struct {
	gdk.Device
}

func wrapWaylandDevice(obj *glib.Object) *WaylandDevice {
	return &WaylandDevice{Device{*externglib.Object{obj}}}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandDisplay struct {
	gdk.Display
}

func wrapWaylandDisplay(obj *glib.Object) *WaylandDisplay {
	return &WaylandDisplay{Display{*externglib.Object{obj}}}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandMonitor struct {
	gdk.Monitor
}

func wrapWaylandMonitor(obj *glib.Object) *WaylandMonitor {
	return &WaylandMonitor{Monitor{*externglib.Object{obj}}}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandPopup struct {
	WaylandSurface
}

func wrapWaylandPopup(obj *glib.Object) *WaylandPopup {
	return &WaylandPopup{WaylandSurface{Surface{*externglib.Object{obj}}}}
}

func marshalWaylandPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandSeat struct {
	gdk.Seat
}

func wrapWaylandSeat(obj *glib.Object) *WaylandSeat {
	return &WaylandSeat{Seat{*externglib.Object{obj}}}
}

func marshalWaylandSeat(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandSurface struct {
	gdk.Surface
}

func wrapWaylandSurface(obj *glib.Object) *WaylandSurface {
	return &WaylandSurface{Surface{*externglib.Object{obj}}}
}

func marshalWaylandSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type WaylandToplevel struct {
	WaylandSurface
}

func wrapWaylandToplevel(obj *glib.Object) *WaylandToplevel {
	return &WaylandToplevel{WaylandSurface{Surface{*externglib.Object{obj}}}}
}

func marshalWaylandToplevel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}
