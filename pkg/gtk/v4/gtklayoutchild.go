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
	GTypeLayoutChild = coreglib.Type(C.gtk_layout_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLayoutChild, F: marshalLayoutChild},
	})
}

// LayoutChildOverrider contains methods that are overridable.
type LayoutChildOverrider interface {
}

// LayoutChild: GtkLayoutChild is the base class for objects that are meant to
// hold layout properties.
//
// If a GtkLayoutManager has per-child properties, like their packing type, or
// the horizontal and vertical span, or the icon name, then the layout manager
// should use a GtkLayoutChild implementation to store those properties.
//
// A GtkLayoutChild instance is only ever valid while a widget is part of a
// layout.
type LayoutChild struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*LayoutChild)(nil)
)

// LayoutChilder describes types inherited from class LayoutChild.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type LayoutChilder interface {
	coreglib.Objector
	baseLayoutChild() *LayoutChild
}

var _ LayoutChilder = (*LayoutChild)(nil)

func initClassLayoutChild(gclass unsafe.Pointer, goval any) {
}

func wrapLayoutChild(obj *coreglib.Object) *LayoutChild {
	return &LayoutChild{
		Object: obj,
	}
}

func marshalLayoutChild(p uintptr) (interface{}, error) {
	return wrapLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (layoutChild *LayoutChild) baseLayoutChild() *LayoutChild {
	return layoutChild
}

// BaseLayoutChild returns the underlying base object.
func BaseLayoutChild(obj LayoutChilder) *LayoutChild {
	return obj.baseLayoutChild()
}

// ChildWidget retrieves the GtkWidget associated to the given layout_child.
//
// The function returns the following values:
//
//    - widget: Widget.
//
func (layoutChild *LayoutChild) ChildWidget() Widgetter {
	var _arg0 *C.GtkLayoutChild // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(coreglib.InternObject(layoutChild).Native()))

	_cret = C.gtk_layout_child_get_child_widget(_arg0)
	runtime.KeepAlive(layoutChild)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// LayoutManager retrieves the GtkLayoutManager instance that created the given
// layout_child.
//
// The function returns the following values:
//
//    - layoutManager: GtkLayoutManager.
//
func (layoutChild *LayoutChild) LayoutManager() LayoutManagerer {
	var _arg0 *C.GtkLayoutChild   // out
	var _cret *C.GtkLayoutManager // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(coreglib.InternObject(layoutChild).Native()))

	_cret = C.gtk_layout_child_get_layout_manager(_arg0)
	runtime.KeepAlive(layoutChild)

	var _layoutManager LayoutManagerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.LayoutManagerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(LayoutManagerer)
			return ok
		})
		rv, ok := casted.(LayoutManagerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.LayoutManagerer")
		}
		_layoutManager = rv
	}

	return _layoutManager
}
