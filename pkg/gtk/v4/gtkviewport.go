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

// GTypeViewport returns the GType for the type Viewport.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeViewport() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Viewport").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalViewport)
	return gtype
}

// Viewport: GtkViewport implements scrollability for widgets that lack their
// own scrolling capabilities.
//
// Use GtkViewport to scroll child widgets such as GtkGrid, GtkBox, and so on.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
//
//
// Accessibility
//
// GtkViewport uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Viewport struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Scrollable
}

var (
	_ Widgetter         = (*Viewport)(nil)
	_ coreglib.Objector = (*Viewport)(nil)
)

func wrapViewport(obj *coreglib.Object) *Viewport {
	return &Viewport{
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
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	return wrapViewport(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewport creates a new GtkViewport.
//
// The new viewport uses the given adjustments, or default adjustments if none
// are given.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal adjustment.
//    - vadjustment (optional): vertical adjustment.
//
// The function returns the following values:
//
//    - viewport: new GtkViewport.
//
func NewViewport(hadjustment, vadjustment *Adjustment) *Viewport {
	var _args [2]girepository.Argument

	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Viewport").InvokeMethod("new_Viewport", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _viewport *Viewport // out

	_viewport = wrapViewport(coreglib.Take(unsafe.Pointer(_cret)))

	return _viewport
}

// Child gets the child widget of viewport.
//
// The function returns the following values:
//
//    - widget (optional): child widget of viewport.
//
func (viewport *Viewport) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_gret := girepository.MustFind("Gtk", "Viewport").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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

// ScrollToFocus gets whether the viewport is scrolling to keep the focused
// child in view.
//
// The function returns the following values:
//
//    - ok: TRUE if the viewport keeps the focus child scrolled to view.
//
func (viewport *Viewport) ScrollToFocus() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_gret := girepository.MustFind("Gtk", "Viewport").InvokeMethod("get_scroll_to_focus", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of viewport.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (viewport *Viewport) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "Viewport").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(viewport)
	runtime.KeepAlive(child)
}

// SetScrollToFocus sets whether the viewport should automatically scroll to
// keep the focused child in view.
//
// The function takes the following parameters:
//
//    - scrollToFocus: whether to keep the focus widget scrolled to view.
//
func (viewport *Viewport) SetScrollToFocus(scrollToFocus bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if scrollToFocus {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Viewport").InvokeMethod("set_scroll_to_focus", _args[:], nil)

	runtime.KeepAlive(viewport)
	runtime.KeepAlive(scrollToFocus)
}
