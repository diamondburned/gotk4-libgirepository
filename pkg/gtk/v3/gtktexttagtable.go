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
//
// void gotk4_TextTagTableForeach(GtkTextTag*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTable},
	})
}

type TextTagTableForeach func(tag TextTag)

//export gotk4_TextTagTableForeach
func gotk4_TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var tag TextTag // out

	tag = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TextTag)

	fn := v.(TextTagTableForeach)
	fn(tag)
}

// TextTagTable: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The GtkTextTagTable implementation of the GtkBuildable interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags:
//
//    <object class="GtkTextTagTable">
//     <child type="tag">
//       <object class="GtkTextTag"/>
//     </child>
//    </object>
type TextTagTable interface {
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

	// Add a tag to the table. The tag is assigned the highest priority in the
	// table.
	//
	// @tag must not be in a tag table already, and may not have the same name
	// as an already-added tag.
	Add(tag TextTag) bool
	// Foreach calls @func on each tag in @table, with user data @data. Note
	// that the table may not be modified while iterating over it (you can’t
	// add/remove tags).
	Foreach(fn TextTagTableForeach)
	// Size returns the size of the table (number of tags)
	Size() int
	// Lookup: look up a named tag.
	Lookup(name string) TextTag
	// Remove a tag from the table. If a TextBuffer has @table as its tag table,
	// the tag is removed from the buffer. The table’s reference to the tag is
	// removed, so the tag will end up destroyed if you don’t have a reference
	// to it.
	Remove(tag TextTag)
}

// textTagTable implements the TextTagTable interface.
type textTagTable struct {
	*externglib.Object
}

var _ TextTagTable = (*textTagTable)(nil)

// WrapTextTagTable wraps a GObject to a type that implements
// interface TextTagTable. It is primarily used internally.
func WrapTextTagTable(obj *externglib.Object) TextTagTable {
	return textTagTable{obj}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextTagTable(obj), nil
}

// NewTextTagTable creates a new TextTagTable. The table contains no tags by
// default.
func NewTextTagTable() TextTagTable {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable TextTagTable // out

	_textTagTable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TextTagTable)

	return _textTagTable
}

func (t textTagTable) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(t))
}

func (b textTagTable) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b textTagTable) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b textTagTable) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b textTagTable) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b textTagTable) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b textTagTable) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b textTagTable) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b textTagTable) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b textTagTable) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b textTagTable) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (t textTagTable) Add(tag TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t textTagTable) Foreach(fn TextTagTableForeach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*[0]byte)(C.gotk4_TextTagTableForeach)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
}

func (t textTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t textTagTable) Lookup(name string) TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.gchar           // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)

	var _textTag TextTag // out

	_textTag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TextTag)

	return _textTag
}

func (t textTagTable) Remove(tag TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
}
