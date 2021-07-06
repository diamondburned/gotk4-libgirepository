// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_size_group_get_type()), F: marshalSizeGroup},
	})
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
// element with `class="GtkSizeGroup"` somewhere in the UI definition. The
// widgets that belong to the size group are specified by a <widgets> element
// that may contain multiple <widget> elements, one for each member of the size
// group. The ”name” attribute gives the id of the widget.
//
// An example of a UI definition fragment with GtkSizeGroup:
//
//    <object class="GtkSizeGroup">
//      <property name="mode">GTK_SIZE_GROUP_HORIZONTAL</property>
//      <widgets>
//        <widget name="radio1"/>
//        <widget name="radio2"/>
//      </widgets>
//    </object>
type SizeGroup interface {
	gextras.Objector

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	//
	// This method is inherited from Buildable
	AddChild(builder Builder, child gextras.Objector, typ string)
	// ConstructChild constructs a child of @buildable with the name @name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// This method is inherited from Buildable
	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the @buildable.
	//
	// This method is inherited from Buildable
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// This method is inherited from Buildable
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// This method is inherited from Buildable
	CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool)
	// GetInternalChild: get the internal child called @childname of the
	// @buildable object.
	//
	// This method is inherited from Buildable
	GetInternalChild(builder Builder, childname string) gextras.Objector
	// GetName gets the name of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetName() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// This method is inherited from Buildable
	ParserFinished(builder Builder)
	// SetBuildableProperty sets the property name @name to @value on the
	// @buildable object.
	//
	// This method is inherited from Buildable
	SetBuildableProperty(builder Builder, name string, value externglib.Value)
	// SetName sets the name of the @buildable object.
	//
	// This method is inherited from Buildable
	SetName(name string)

	// AddWidget adds a widget to a SizeGroup. In the future, the requisition of
	// the widget will be determined as the maximum of its requisition and the
	// requisition of the other widgets in the size group. Whether this applies
	// horizontally, vertically, or in both directions depends on the mode of
	// the size group. See gtk_size_group_set_mode().
	//
	// When the widget is destroyed or no longer referenced elsewhere, it will
	// be removed from the size group.
	AddWidget(widget Widget)
	// IgnoreHidden returns if invisible widgets are ignored when calculating
	// the size.
	//
	// Deprecated: since version 3.22.
	IgnoreHidden() bool
	// Mode gets the current mode of the size group. See
	// gtk_size_group_set_mode().
	Mode() SizeGroupMode
	// RemoveWidget removes a widget from a SizeGroup.
	RemoveWidget(widget Widget)
	// SetIgnoreHidden sets whether unmapped widgets should be ignored when
	// calculating the size.
	//
	// Deprecated: since version 3.22.
	SetIgnoreHidden(ignoreHidden bool)
	// SetMode sets the SizeGroupMode of the size group. The mode of the size
	// group determines whether the widgets in the size group should all have
	// the same horizontal requisition (GTK_SIZE_GROUP_HORIZONTAL) all have the
	// same vertical requisition (GTK_SIZE_GROUP_VERTICAL), or should all have
	// the same requisition in both directions (GTK_SIZE_GROUP_BOTH).
	SetMode(mode SizeGroupMode)
}

// sizeGroup implements the SizeGroup interface.
type sizeGroup struct {
	*externglib.Object
}

var _ SizeGroup = (*sizeGroup)(nil)

// WrapSizeGroup wraps a GObject to a type that implements
// interface SizeGroup. It is primarily used internally.
func WrapSizeGroup(obj *externglib.Object) SizeGroup {
	return sizeGroup{obj}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSizeGroup(obj), nil
}

// NewSizeGroup: create a new SizeGroup.
func NewSizeGroup(mode SizeGroupMode) SizeGroup {
	var _arg1 C.GtkSizeGroupMode // out
	var _cret *C.GtkSizeGroup    // in

	_arg1 = C.GtkSizeGroupMode(mode)

	_cret = C.gtk_size_group_new(_arg1)

	var _sizeGroup SizeGroup // out

	_sizeGroup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(SizeGroup)

	return _sizeGroup
}

func (s sizeGroup) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(s))
}

func (b sizeGroup) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b sizeGroup) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b sizeGroup) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b sizeGroup) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b sizeGroup) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b sizeGroup) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b sizeGroup) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b sizeGroup) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b sizeGroup) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b sizeGroup) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (s sizeGroup) AddWidget(widget Widget) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_size_group_add_widget(_arg0, _arg1)
}

func (s sizeGroup) IgnoreHidden() bool {
	var _arg0 *C.GtkSizeGroup // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_size_group_get_ignore_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s sizeGroup) Mode() SizeGroupMode {
	var _arg0 *C.GtkSizeGroup    // out
	var _cret C.GtkSizeGroupMode // in

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_size_group_get_mode(_arg0)

	var _sizeGroupMode SizeGroupMode // out

	_sizeGroupMode = SizeGroupMode(_cret)

	return _sizeGroupMode
}

func (s sizeGroup) RemoveWidget(widget Widget) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_size_group_remove_widget(_arg0, _arg1)
}

func (s sizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	var _arg0 *C.GtkSizeGroup // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	if ignoreHidden {
		_arg1 = C.TRUE
	}

	C.gtk_size_group_set_ignore_hidden(_arg0, _arg1)
}

func (s sizeGroup) SetMode(mode SizeGroupMode) {
	var _arg0 *C.GtkSizeGroup    // out
	var _arg1 C.GtkSizeGroupMode // out

	_arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkSizeGroupMode(mode)

	C.gtk_size_group_set_mode(_arg0, _arg1)
}
