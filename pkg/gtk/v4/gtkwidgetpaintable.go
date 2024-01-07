// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeWidgetPaintable = coreglib.Type(girepository.MustFind("Gtk", "WidgetPaintable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWidgetPaintable, F: marshalWidgetPaintable},
	})
}

// WidgetPaintableOverrides contains methods that are overridable.
type WidgetPaintableOverrides struct {
}

func defaultWidgetPaintableOverrides(v *WidgetPaintable) WidgetPaintableOverrides {
	return WidgetPaintableOverrides{}
}

// WidgetPaintable: GtkWidgetPaintable is a GdkPaintable that displays the
// contents of a widget.
//
// GtkWidgetPaintable will also take care of the widget not being in a state
// where it can be drawn (like when it isn't shown) and just draw nothing or
// where it does not have a size (like when it is hidden) and report no size in
// that case.
//
// Of course, GtkWidgetPaintable allows you to monitor widgets for size changes
// by emitting the gdk.Paintable::invalidate-size signal whenever the size of
// the widget changes as well as for visual changes by emitting the
// gdk.Paintable::invalidate-contents signal whenever the widget changes.
//
// You can use a GtkWidgetPaintable everywhere a GdkPaintable is allowed,
// including using it on a GtkPicture (or one of its parents) that it was set on
// itself via gtk_picture_set_paintable(). The paintable will take care of
// recursion when this happens. If you do this however, ensure that the
// gtk.Picture:can-shrink property is set to TRUE or you might end up with an
// infinitely growing widget.
type WidgetPaintable struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gdk.Paintable
}

var (
	_ coreglib.Objector = (*WidgetPaintable)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WidgetPaintable, *WidgetPaintableClass, WidgetPaintableOverrides](
		GTypeWidgetPaintable,
		initWidgetPaintableClass,
		wrapWidgetPaintable,
		defaultWidgetPaintableOverrides,
	)
}

func initWidgetPaintableClass(gclass unsafe.Pointer, overrides WidgetPaintableOverrides, classInitFunc func(*WidgetPaintableClass)) {
	if classInitFunc != nil {
		class := (*WidgetPaintableClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWidgetPaintable(obj *coreglib.Object) *WidgetPaintable {
	return &WidgetPaintable{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalWidgetPaintable(p uintptr) (interface{}, error) {
	return wrapWidgetPaintable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WidgetPaintableClass: instance of this type is always passed by reference.
type WidgetPaintableClass struct {
	*widgetPaintableClass
}

// widgetPaintableClass is the struct that's finalized.
type widgetPaintableClass struct {
	native unsafe.Pointer
}

var GIRInfoWidgetPaintableClass = girepository.MustFind("Gtk", "WidgetPaintableClass")
