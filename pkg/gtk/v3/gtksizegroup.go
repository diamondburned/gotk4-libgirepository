// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtksizegroup.go.
var GTypeSizeGroup = externglib.Type(C.gtk_size_group_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSizeGroup, F: marshalSizeGroup},
	})
}

// SizeGroupOverrider contains methods that are overridable.
type SizeGroupOverrider interface {
	externglib.Objector
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
	*externglib.Object

	Buildable
}

var (
	_ externglib.Objector = (*SizeGroup)(nil)
)

func classInitSizeGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSizeGroup(obj *externglib.Object) *SizeGroup {
	return &SizeGroup{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	return wrapSizeGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSizeGroup: create a new SizeGroup.
//
// The function takes the following parameters:
//
//    - mode for the new size group.
//
// The function returns the following values:
//
//    - sizeGroup: newly created SizeGroup.
//
func NewSizeGroup(mode SizeGroupMode) *SizeGroup {
	var _arg1 C.GtkSizeGroupMode // out
	var _cret *C.GtkSizeGroup    // in

	_arg1 = C.GtkSizeGroupMode(mode)

	_cret = C.gtk_size_group_new(_arg1)
	runtime.KeepAlive(mode)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sizeGroup
}

// AddWidget adds a widget to a SizeGroup. In the future, the requisition of the
// widget will be determined as the maximum of its requisition and the
// requisition of the other widgets in the size group. Whether this applies
// horizontally, vertically, or in both directions depends on the mode of the
// size group. See gtk_size_group_set_mode().
//
// When the widget is destroyed or no longer referenced elsewhere, it will be
// removed from the size group.
//
// The function takes the following parameters:
//
//    - widget to add.
//
func (sizeGroup *SizeGroup) AddWidget(widget Widgetter) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_size_group_add_widget(_arg0, _arg1)
	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(widget)
}

// IgnoreHidden returns if invisible widgets are ignored when calculating the
// size.
//
// Deprecated: Measuring the size of hidden widgets has not worked reliably for
// a long time. In most cases, they will report a size of 0 nowadays, and thus,
// their size will not affect the other size group members. In effect, size
// groups will always operate as if this property was TRUE. Use a Stack instead
// to hide widgets while still having their size taken into account.
//
// The function returns the following values:
//
//    - ok: TRUE if invisible widgets are ignored.
//
func (sizeGroup *SizeGroup) IgnoreHidden() bool {
	var _arg0 *C.GtkSizeGroup // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))

	_cret = C.gtk_size_group_get_ignore_hidden(_arg0)
	runtime.KeepAlive(sizeGroup)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Mode gets the current mode of the size group. See gtk_size_group_set_mode().
//
// The function returns the following values:
//
//    - sizeGroupMode: current mode of the size group.
//
func (sizeGroup *SizeGroup) Mode() SizeGroupMode {
	var _arg0 *C.GtkSizeGroup    // out
	var _cret C.GtkSizeGroupMode // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))

	_cret = C.gtk_size_group_get_mode(_arg0)
	runtime.KeepAlive(sizeGroup)

	var _sizeGroupMode SizeGroupMode // out

	_sizeGroupMode = SizeGroupMode(_cret)

	return _sizeGroupMode
}

// Widgets returns the list of widgets associated with size_group.
//
// The function returns the following values:
//
//    - sList of widgets. The list is owned by GTK+ and should not be modified.
//
func (sizeGroup *SizeGroup) Widgets() []Widgetter {
	var _arg0 *C.GtkSizeGroup // out
	var _cret *C.GSList       // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))

	_cret = C.gtk_size_group_get_widgets(_arg0)
	runtime.KeepAlive(sizeGroup)

	var _sList []Widgetter // out

	_sList = make([]Widgetter, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkWidget)(v)
		var dst Widgetter // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Widgetter is nil")
			}

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			dst = rv
		}
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveWidget removes a widget from a SizeGroup.
//
// The function takes the following parameters:
//
//    - widget to remove.
//
func (sizeGroup *SizeGroup) RemoveWidget(widget Widgetter) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_size_group_remove_widget(_arg0, _arg1)
	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(widget)
}

// SetIgnoreHidden sets whether unmapped widgets should be ignored when
// calculating the size.
//
// Deprecated: Measuring the size of hidden widgets has not worked reliably for
// a long time. In most cases, they will report a size of 0 nowadays, and thus,
// their size will not affect the other size group members. In effect, size
// groups will always operate as if this property was TRUE. Use a Stack instead
// to hide widgets while still having their size taken into account.
//
// The function takes the following parameters:
//
//    - ignoreHidden: whether unmapped widgets should be ignored when calculating
//      the size.
//
func (sizeGroup *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))
	if ignoreHidden {
		_arg1 = C.TRUE
	}

	C.gtk_size_group_set_ignore_hidden(_arg0, _arg1)
	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(ignoreHidden)
}

// SetMode sets the SizeGroupMode of the size group. The mode of the size group
// determines whether the widgets in the size group should all have the same
// horizontal requisition (GTK_SIZE_GROUP_HORIZONTAL) all have the same vertical
// requisition (GTK_SIZE_GROUP_VERTICAL), or should all have the same
// requisition in both directions (GTK_SIZE_GROUP_BOTH).
//
// The function takes the following parameters:
//
//    - mode to set for the size group.
//
func (sizeGroup *SizeGroup) SetMode(mode SizeGroupMode) {
	var _arg0 *C.GtkSizeGroup    // out
	var _arg1 C.GtkSizeGroupMode // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(externglib.InternObject(sizeGroup).Native()))
	_arg1 = C.GtkSizeGroupMode(mode)

	C.gtk_size_group_set_mode(_arg0, _arg1)
	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(mode)
}
