// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GObject* _gotk4_gtk3_BuildableIface_construct_child(void*, void*, void*);
// extern GObject* _gotk4_gtk3_BuildableIface_get_internal_child(void*, void*, void*);
// extern gboolean _gotk4_gtk3_BuildableIface_custom_tag_start(void*, void*, void*, void*, void*, void*);
// extern gchar* _gotk4_gtk3_BuildableIface_get_name(void*);
// extern void _gotk4_gtk3_BuildableIface_add_child(void*, void*, void*, void*);
// extern void _gotk4_gtk3_BuildableIface_custom_finished(void*, void*, void*, void*, gpointer);
// extern void _gotk4_gtk3_BuildableIface_custom_tag_end(void*, void*, void*, void*, void*);
// extern void _gotk4_gtk3_BuildableIface_parser_finished(void*, void*);
// extern void _gotk4_gtk3_BuildableIface_set_buildable_property(void*, void*, void*, void*);
// extern void _gotk4_gtk3_BuildableIface_set_name(void*, void*);
import "C"

// glib.Type values for gtkbuildable.go.
var GTypeBuildable = coreglib.Type(C.gtk_buildable_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBuildable, F: marshalBuildable},
	})
}

// BuildableOverrider contains methods that are overridable.
type BuildableOverrider interface {
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child to add.
	//    - typ (optional): kind of child or NULL.
	//
	AddChild(builder *Builder, child *coreglib.Object, typ string)
	// ConstructChild constructs a child of buildable with the name name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - name of child to construct.
	//
	// The function returns the following values:
	//
	//    - object: constructed child.
	//
	ConstructChild(builder *Builder, name string) *coreglib.Object
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of the tag.
	//    - data (optional): user data created in custom_tag_start.
	//
	CustomFinished(builder *Builder, child *coreglib.Object, tagname string, data unsafe.Pointer)
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//    - data (optional): user data that will be passed in to parser
	//      functions.
	//
	CustomTagEnd(builder *Builder, child *coreglib.Object, tagname string, data *unsafe.Pointer)
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//
	// The function returns the following values:
	//
	//    - parser to fill in.
	//    - data (optional): return location for user data that will be passed in
	//      to parser functions.
	//    - ok: TRUE if a object has a custom implementation, FALSE if it
	//      doesn't.
	//
	CustomTagStart(builder *Builder, child *coreglib.Object, tagname string) (*glib.MarkupParser, unsafe.Pointer, bool)
	// InternalChild: get the internal child called childname of the buildable
	// object.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - childname: name of child.
	//
	// The function returns the following values:
	//
	//    - object: internal child of the buildable object.
	//
	InternalChild(builder *Builder, childname string) *coreglib.Object
	// Name gets the name of the buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the buildable.
	//
	// The function returns the following values:
	//
	//    - utf8: name set with gtk_buildable_set_name().
	//
	Name() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//
	ParserFinished(builder *Builder)
	// SetBuildableProperty sets the property name name to value on the
	// buildable object.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - name of property.
	//    - value of property.
	//
	SetBuildableProperty(builder *Builder, name string, value *coreglib.Value)
	// SetName sets the name of the buildable object.
	//
	// The function takes the following parameters:
	//
	//    - name to set.
	//
	SetName(name string)
}

// Buildable allows objects to extend and customize their deserialization from
// [GtkBuilder UI descriptions][BUILDER-UI]. The interface includes methods for
// setting names and properties of objects, parsing custom tags and constructing
// child objects.
//
// The GtkBuildable interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK+. The main user of this interface
// is Builder. There should be very little need for applications to call any of
// these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// Builder format or run any extra routines at deserialization time.
//
// Buildable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Buildable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Buildable)(nil)
)

// Buildabler describes Buildable's interface methods.
type Buildabler interface {
	coreglib.Objector

	// AddChild adds a child to buildable.
	AddChild(builder *Builder, child *coreglib.Object, typ string)
	// ConstructChild constructs a child of buildable with the name name.
	ConstructChild(builder *Builder, name string) *coreglib.Object
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the buildable.
	CustomFinished(builder *Builder, child *coreglib.Object, tagname string, data unsafe.Pointer)
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	CustomTagEnd(builder *Builder, child *coreglib.Object, tagname string, data *unsafe.Pointer)
	// CustomTagStart: this is called for each unknown element under <child>.
	CustomTagStart(builder *Builder, child *coreglib.Object, tagname string) (*glib.MarkupParser, unsafe.Pointer, bool)
	// InternalChild: get the internal child called childname of the buildable
	// object.
	InternalChild(builder *Builder, childname string) *coreglib.Object
	// Name gets the name of the buildable object.
	Name() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI].
	ParserFinished(builder *Builder)
	// SetBuildableProperty sets the property name name to value on the
	// buildable object.
	SetBuildableProperty(builder *Builder, name string, value *coreglib.Value)
	// SetName sets the name of the buildable object.
	SetName(name string)
}

var _ Buildabler = (*Buildable)(nil)

func ifaceInitBuildabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "BuildableIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("add_child"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_add_child)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("construct_child"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_construct_child)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_finished"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_custom_finished)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_tag_end"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_custom_tag_end)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_tag_start"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_custom_tag_start)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_internal_child"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_get_internal_child)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_name"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_get_name)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("parser_finished"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_parser_finished)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_buildable_property"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_set_buildable_property)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_name"))) = unsafe.Pointer(C._gotk4_gtk3_BuildableIface_set_name)
}

//export _gotk4_gtk3_BuildableIface_add_child
func _gotk4_gtk3_BuildableIface_add_child(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _typ string             // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_child = coreglib.Take(unsafe.Pointer(arg2))
	if arg3 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	}

	iface.AddChild(_builder, _child, _typ)
}

//export _gotk4_gtk3_BuildableIface_construct_child
func _gotk4_gtk3_BuildableIface_construct_child(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret *C.GObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _name string      // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.ConstructChild(_builder, _name)

	cret = (*C.void)(unsafe.Pointer(object.Native()))
	C.g_object_ref(C.gpointer(object.Native()))

	return cret
}

//export _gotk4_gtk3_BuildableIface_custom_finished
func _gotk4_gtk3_BuildableIface_custom_finished(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out
	var _data unsafe.Pointer    // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (unsafe.Pointer)(unsafe.Pointer(arg4))

	iface.CustomFinished(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk3_BuildableIface_custom_tag_end
func _gotk4_gtk3_BuildableIface_custom_tag_end(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out
	var _data *unsafe.Pointer   // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	if arg4 != nil {
		_data = (*unsafe.Pointer)(unsafe.Pointer(arg4))
	}

	iface.CustomTagEnd(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk3_BuildableIface_custom_tag_start
func _gotk4_gtk3_BuildableIface_custom_tag_start(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 *C.void, arg5 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	parser, data, ok := iface.CustomTagStart(_builder, _child, _tagname)

	*arg4 = (*C.void)(gextras.StructNative(unsafe.Pointer(parser)))
	*arg5 = (*C.void)(unsafe.Pointer(data))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_BuildableIface_get_internal_child
func _gotk4_gtk3_BuildableIface_get_internal_child(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret *C.GObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _childname string // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_childname = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.InternalChild(_builder, _childname)

	cret = (*C.void)(unsafe.Pointer(object.Native()))

	return cret
}

//export _gotk4_gtk3_BuildableIface_get_name
func _gotk4_gtk3_BuildableIface_get_name(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	utf8 := iface.Name()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk3_BuildableIface_parser_finished
func _gotk4_gtk3_BuildableIface_parser_finished(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))

	iface.ParserFinished(_builder)
}

//export _gotk4_gtk3_BuildableIface_set_buildable_property
func _gotk4_gtk3_BuildableIface_set_buildable_property(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder      // out
	var _name string           // out
	var _value *coreglib.Value // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_value = coreglib.ValueFromNative(unsafe.Pointer(arg3))

	iface.SetBuildableProperty(_builder, _name, _value)
}

//export _gotk4_gtk3_BuildableIface_set_name
func _gotk4_gtk3_BuildableIface_set_name(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetName(_name)
}

func wrapBuildable(obj *coreglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	return wrapBuildable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddChild adds a child to buildable. type is an optional string describing how
// the child should be added.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - child to add.
//    - typ (optional): kind of child or NULL.
//
func (buildable *Buildable) AddChild(builder *Builder, child *coreglib.Object, typ string) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(child.Native()))
	if typ != "" {
		*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(typ)))
		defer C.free(unsafe.Pointer(_args[3]))
	}

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(typ)
}

// ConstructChild constructs a child of buildable with the name name.
//
// Builder calls this function if a “constructor” has been specified in the UI
// definition.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - name of child to construct.
//
// The function returns the following values:
//
//    - object: constructed child.
//
func (buildable *Buildable) ConstructChild(builder *Builder, name string) *coreglib.Object {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)

	var _object *coreglib.Object // out

	_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
// called once for each custom tag handled by the buildable.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of the tag.
//    - data (optional): user data created in custom_tag_start.
//
func (buildable *Buildable) CustomFinished(builder *Builder, child *coreglib.Object, tagname string, data unsafe.Pointer) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(child.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_args[3]))
	*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (C.gpointer)(unsafe.Pointer(data))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)
	runtime.KeepAlive(data)
}

// CustomTagEnd: this is called at the end of each custom element handled by the
// buildable.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of tag.
//    - data (optional): user data that will be passed in to parser functions.
//
func (buildable *Buildable) CustomTagEnd(builder *Builder, child *coreglib.Object, tagname string, data *unsafe.Pointer) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(child.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_args[3]))
	if data != nil {
		*(**C.void)(unsafe.Pointer(&_args[4])) = (*C.void)(unsafe.Pointer(data))
	}

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)
	runtime.KeepAlive(data)
}

// CustomTagStart: this is called for each unknown element under <child>.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of tag.
//
// The function returns the following values:
//
//    - parser to fill in.
//    - data (optional): return location for user data that will be passed in to
//      parser functions.
//    - ok: TRUE if a object has a custom implementation, FALSE if it doesn't.
//
func (buildable *Buildable) CustomTagStart(builder *Builder, child *coreglib.Object, tagname string) (*glib.MarkupParser, unsafe.Pointer, bool) {
	var _args [4]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(child.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_args[3]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)

	var _parser *glib.MarkupParser // out
	var _data unsafe.Pointer       // out
	var _ok bool                   // out

	_parser = (*glib.MarkupParser)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_data = (unsafe.Pointer)(unsafe.Pointer(_outs[1]))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _parser, _data, _ok
}

// InternalChild: get the internal child called childname of the buildable
// object.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - childname: name of child.
//
// The function returns the following values:
//
//    - object: internal child of the buildable object.
//
func (buildable *Buildable) InternalChild(builder *Builder, childname string) *coreglib.Object {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(childname)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(childname)

	var _object *coreglib.Object // out

	_object = coreglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Name gets the name of the buildable object.
//
// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
// used to construct the buildable.
//
// The function returns the following values:
//
//    - utf8: name set with gtk_buildable_set_name().
//
func (buildable *Buildable) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buildable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ParserFinished: called when the builder finishes the parsing of a [GtkBuilder
// UI definition][BUILDER-UI]. Note that this will be called once for each time
// gtk_builder_add_from_file() or gtk_builder_add_from_string() is called on a
// builder.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//
func (buildable *Buildable) ParserFinished(builder *Builder) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
}

// SetBuildableProperty sets the property name name to value on the buildable
// object.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - name of property.
//    - value of property.
//
func (buildable *Buildable) SetBuildableProperty(builder *Builder, name string, value *coreglib.Value) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(value.Native()))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// SetName sets the name of the buildable object.
//
// The function takes the following parameters:
//
//    - name to set.
//
func (buildable *Buildable) SetName(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(buildable)
	runtime.KeepAlive(name)
}
