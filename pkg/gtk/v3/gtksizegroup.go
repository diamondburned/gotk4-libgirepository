// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSizeGroup = coreglib.Type(girepository.MustFind("Gtk", "SizeGroup").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSizeGroup, F: marshalSizeGroup},
	})
}

// SizeGroupOverrides contains methods that are overridable.
type SizeGroupOverrides struct {
}

func defaultSizeGroupOverrides(v *SizeGroup) SizeGroupOverrides {
	return SizeGroupOverrides{}
}

// SizeGroup provides a mechanism for grouping a number of widgets together so
// they all request the same amount of space. This is typically useful when you
// want a column of widgets to have the same size, but you can’t use a Grid
// widget.
//
// In detail, the size requested for each widget in a SizeGroup is the maximum
// of the sizes that would have been requested for each widget in the size group
// if they were not in the size group. The mode of the size group (see
// gtk_size_group_set_mode()) determines whether this applies to the horizontal
// size, the vertical size, or both sizes.
//
// Note that size groups only affect the amount of space requested, not the size
// that the widgets finally receive. If you want the widgets in a SizeGroup to
// actually be the same size, you need to pack them in such a way that they get
// the size they request and not more. For example, if you are packing your
// widgets into a table, you would not include the GTK_FILL flag.
//
// SizeGroup objects are referenced by each widget in the size group, so once
// you have added all widgets to a SizeGroup, you can drop the initial reference
// to the size group with g_object_unref(). If the widgets in the size group are
// subsequently destroyed, then they will be removed from the size group and
// drop their references on the size group; when all widgets have been removed,
// the size group will be freed.
//
// Widgets can be part of multiple size groups; GTK+ will compute the horizontal
// size of a widget from the horizontal requisition of all widgets that can be
// reached from the widget by a chain of size groups of type
// GTK_SIZE_GROUP_HORIZONTAL or GTK_SIZE_GROUP_BOTH, and the vertical size from
// the vertical requisition of all widgets that can be reached from the widget
// by a chain of size groups of type GTK_SIZE_GROUP_VERTICAL or
// GTK_SIZE_GROUP_BOTH.
//
// Note that only non-contextual sizes of every widget are ever consulted by
// size groups (since size groups have no knowledge of what size a widget will
// be allocated in one dimension, it cannot derive how much height a widget will
// receive for a given width). When grouping widgets that trade height for width
// in mode GTK_SIZE_GROUP_VERTICAL or GTK_SIZE_GROUP_BOTH: the height for the
// minimum width will be the requested height for all widgets in the group. The
// same is of course true when horizontally grouping width for height widgets.
//
// Widgets that trade height-for-width should set a reasonably large minimum
// width by way of Label:width-chars for instance. Widgets with static sizes as
// well as widgets that grow (such as ellipsizing text) need no such
// considerations.
//
//
// GtkSizeGroup as GtkBuildable
//
// Size groups can be specified in a UI definition by placing an <object>
// element with class="GtkSizeGroup" somewhere in the UI definition. The widgets
// that belong to the size group are specified by a <widgets> element that may
// contain multiple <widget> elements, one for each member of the size group.
// The ”name” attribute gives the id of the widget.
//
// An example of a UI definition fragment with GtkSizeGroup:
//
//    <object class="GtkSizeGroup">
//      <property name="mode">GTK_SIZE_GROUP_HORIZONTAL</property>
//      <widgets>
//        <widget name="radio1"/>
//        <widget name="radio2"/>
//      </widgets>
//    </object>.
type SizeGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
}

var (
	_ coreglib.Objector = (*SizeGroup)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SizeGroup, *SizeGroupClass, SizeGroupOverrides](
		GTypeSizeGroup,
		initSizeGroupClass,
		wrapSizeGroup,
		defaultSizeGroupOverrides,
	)
}

func initSizeGroupClass(gclass unsafe.Pointer, overrides SizeGroupOverrides, classInitFunc func(*SizeGroupClass)) {
	if classInitFunc != nil {
		class := (*SizeGroupClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSizeGroup(obj *coreglib.Object) *SizeGroup {
	return &SizeGroup{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	return wrapSizeGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SizeGroupClass: instance of this type is always passed by reference.
type SizeGroupClass struct {
	*sizeGroupClass
}

// sizeGroupClass is the struct that's finalized.
type sizeGroupClass struct {
	native unsafe.Pointer
}

var GIRInfoSizeGroupClass = girepository.MustFind("Gtk", "SizeGroupClass")
