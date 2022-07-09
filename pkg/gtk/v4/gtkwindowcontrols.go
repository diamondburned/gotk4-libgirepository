// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeWindowControls returns the GType for the type WindowControls.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeWindowControls() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "WindowControls").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalWindowControls)
	return gtype
}

// WindowControlsOverrider contains methods that are overridable.
type WindowControlsOverrider interface {
}

// WindowControls: GtkWindowControls shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// GtkWindowControls only displays start or end side of the controls (see
// gtk.WindowControls:side), so it's intended to be always used in pair with
// another GtkWindowControls for the opposite side, for example:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">start</property>
//        </object>
//      </child>
//
//      ...
//
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">end</property>
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    windowcontrols
//    ├── [image.icon]
//    ├── [button.minimize]
//    ├── [button.maximize]
//    ╰── [button.close]
//
//
// A GtkWindowControls' CSS node is called windowcontrols. It contains subnodes
// corresponding to each title button. Which of the title buttons exist and
// where they are placed exactly depends on the desktop environment and
// gtk.WindowControls:decoration-layout value.
//
// When gtk.WindowControls:empty is TRUE, it gets the .empty style class.
//
//
// Accessibility
//
// GtkWindowControls uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*WindowControls)(nil)
)

func classInitWindowControlser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWindowControls(obj *coreglib.Object) *WindowControls {
	return &WindowControls{
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

func marshalWindowControls(p uintptr) (interface{}, error) {
	return wrapWindowControls(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DecorationLayout gets the decoration layout of this GtkWindowControls.
//
// The function returns the following values:
//
//    - utf8 (optional): decoration layout or NULL if it is unset.
//
func (self *WindowControls) DecorationLayout() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "WindowControls").InvokeMethod("get_decoration_layout", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Empty gets whether the widget has any window buttons.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget has window buttons, otherwise FALSE.
//
func (self *WindowControls) Empty() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "WindowControls").InvokeMethod("get_empty", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetDecorationLayout sets the decoration layout for the title buttons.
//
// This overrides the gtk.Settings:gtk-decoration-layout setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close and icon (the window
// icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon on the left, and
// minimize, maximize and close buttons on the right.
//
// If gtk.WindowControls:side value is GTK_PACK_START, self will display the
// part before the colon, otherwise after that.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (self *WindowControls) SetDecorationLayout(layout string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if layout != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "WindowControls").InvokeMethod("set_decoration_layout", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(layout)
}
