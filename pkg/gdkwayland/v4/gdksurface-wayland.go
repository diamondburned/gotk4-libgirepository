// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
// void _gotk4_gdkwayland4_WaylandToplevelExported(GdkToplevel*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_popup_get_type()), F: marshalWaylandPopupper},
		{T: externglib.Type(C.gdk_wayland_surface_get_type()), F: marshalWaylandSurfacer},
		{T: externglib.Type(C.gdk_wayland_toplevel_get_type()), F: marshalWaylandTopleveller},
	})
}

// WaylandPopup: wayland implementation of GdkPopup.
type WaylandPopup struct {
	WaylandSurface

	*externglib.Object
	gdk.Popup
	gdk.Surface
}

var (
	_ externglib.Objector = (*WaylandPopup)(nil)
	_ gdk.Surfacer        = (*WaylandPopup)(nil)
)

func wrapWaylandPopup(obj *externglib.Object) *WaylandPopup {
	return &WaylandPopup{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Popup: gdk.Popup{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandPopupper(p uintptr) (interface{}, error) {
	return wrapWaylandPopup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandSurface: wayland implementation of GdkSurface.
//
// Beyond the gdk.Surface API, the Wayland implementation offers access to the
// Wayland wl_surface object with gdkwayland.WaylandSurface.GetWlSurface().
type WaylandSurface struct {
	gdk.Surface
}

var (
	_ gdk.Surfacer = (*WaylandSurface)(nil)
)

func wrapWaylandSurface(obj *externglib.Object) *WaylandSurface {
	return &WaylandSurface{
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandSurfacer(p uintptr) (interface{}, error) {
	return wrapWaylandSurface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandToplevel: wayland implementation of GdkToplevel.
//
// Beyond the gdk.Toplevel API, the Wayland implementation has API to set up
// cross-process parent-child relationships between surfaces with
// gdkwayland.WaylandToplevel.ExportHandle() and
// gdkwayland.WaylandToplevel.SetTransientForExported().
type WaylandToplevel struct {
	WaylandSurface

	*externglib.Object
	gdk.Surface
	gdk.Toplevel
}

var (
	_ externglib.Objector = (*WaylandToplevel)(nil)
	_ gdk.Surfacer        = (*WaylandToplevel)(nil)
)

func wrapWaylandToplevel(obj *externglib.Object) *WaylandToplevel {
	return &WaylandToplevel{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Surface: gdk.Surface{
			Object: obj,
		},
		Toplevel: gdk.Toplevel{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
	}
}

func marshalWaylandTopleveller(p uintptr) (interface{}, error) {
	return wrapWaylandToplevel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ExportHandle: asynchronously obtains a handle for a surface that can be
// passed to other processes.
//
// When the handle has been obtained, callback will be called.
//
// It is an error to call this function on a surface that is already exported.
//
// When the handle is no longer needed,
// gdkwayland.WaylandToplevel.UnexportHandle() should be called to clean up
// resources.
//
// The main purpose for obtaining a handle is to mark a surface from another
// surface as transient for this one, see
// gdkwayland.WaylandToplevel.SetTransientForExported().
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//    - callback to call with the handle.
//
func (toplevel *WaylandToplevel) ExportHandle(callback WaylandToplevelExported) bool {
	var _arg0 *C.GdkToplevel               // out
	var _arg1 C.GdkWaylandToplevelExported // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret C.gboolean // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gdkwayland4_WaylandToplevelExported)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_wayland_toplevel_export_handle(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(callback)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetApplicationID sets the application id on a GdkToplevel.
//
// The function takes the following parameters:
//
//    - applicationId: application id for the toplevel.
//
func (toplevel *WaylandToplevel) SetApplicationID(applicationId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(applicationId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_toplevel_set_application_id(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(applicationId)
}

// SetTransientForExported marks toplevel as transient for the surface to which
// the given parent_handle_str refers.
//
// Typically, the handle will originate from a
// gdkwayland.WaylandToplevel.ExportHandle() call in another process.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//    - parentHandleStr: exported handle for a surface.
//
func (toplevel *WaylandToplevel) SetTransientForExported(parentHandleStr string) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parentHandleStr)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_toplevel_set_transient_for_exported(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(parentHandleStr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnexportHandle destroys the handle that was obtained with
// gdk_wayland_toplevel_export_handle().
//
// It is an error to call this function on a surface that does not have a
// handle.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (toplevel *WaylandToplevel) UnexportHandle() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	C.gdk_wayland_toplevel_unexport_handle(_arg0)
	runtime.KeepAlive(toplevel)
}
