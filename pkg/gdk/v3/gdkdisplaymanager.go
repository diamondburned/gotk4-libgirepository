// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_manager_get_type()), F: marshalDisplayManager},
	})
}

// DisplayManager: the purpose of the DisplayManager singleton object is to
// offer notification when displays appear or disappear or the default display
// changes.
//
// You can use gdk_display_manager_get() to obtain the DisplayManager singleton,
// but that should be rarely necessary. Typically, initializing GTK+ opens a
// display that you can work with without ever accessing the DisplayManager.
//
// The GDK library can be built with support for multiple backends. The
// DisplayManager object determines which backend is used at runtime.
//
// When writing backend-specific code that is supposed to work with multiple GDK
// backends, you have to consider both compile time and runtime. At compile
// time, use the K_WINDOWING_X11, K_WINDOWING_WIN32 macros, etc. to find out
// which backends are present in the GDK library you are building your
// application against. At runtime, use type-check macros like
// GDK_IS_X11_DISPLAY() to find out which backend is in use:
//
// Backend-specific code
//
//    #ifdef GDK_WINDOWING_X11
//      if (GDK_IS_X11_DISPLAY (display))
//        {
//          // make X11-specific calls here
//        }
//      else
//    #endif
//    #ifdef GDK_WINDOWING_QUARTZ
//      if (GDK_IS_QUARTZ_DISPLAY (display))
//        {
//          // make Quartz-specific calls here
//        }
//      else
//    #endif
//      g_error ("Unsupported GDK backend");
type DisplayManager interface {
	gextras.Objector

	// DefaultDisplay gets the default Display.
	DefaultDisplay() Display
	// OpenDisplay opens a display.
	OpenDisplay(name string) Display
	// SetDefaultDisplay sets @display as the default display.
	SetDefaultDisplay(display Display)
}

// displayManager implements the DisplayManager class.
type displayManager struct {
	gextras.Objector
}

var _ DisplayManager = (*displayManager)(nil)

// WrapDisplayManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapDisplayManager(obj *externglib.Object) DisplayManager {
	return displayManager{
		Objector: obj,
	}
}

func marshalDisplayManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDisplayManager(obj), nil
}

// DefaultDisplay gets the default Display.
func (m displayManager) DefaultDisplay() Display {
	var _arg0 *C.GdkDisplayManager // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(m.Native()))

	_cret = C.gdk_display_manager_get_default_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

// OpenDisplay opens a display.
func (m displayManager) OpenDisplay(name string) Display {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.gchar             // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_manager_open_display(_arg0, _arg1)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

// SetDefaultDisplay sets @display as the default display.
func (m displayManager) SetDefaultDisplay(display Display) {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_manager_set_default_display(_arg0, _arg1)
}
