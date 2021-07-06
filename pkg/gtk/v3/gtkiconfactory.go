// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_icon_factory_get_type()), F: marshalIconFactory},
	})
}

// IconFactory: icon factory manages a collection of IconSet; a IconSet manages
// a set of variants of a particular icon (i.e. a IconSet contains variants for
// different sizes and widget states). Icons in an icon factory are named by a
// stock ID, which is a simple string identifying the icon. Each Style has a
// list of IconFactory derived from the current theme; those icon factories are
// consulted first when searching for an icon. If the theme doesn’t set a
// particular icon, GTK+ looks for the icon in a list of default icon factories,
// maintained by gtk_icon_factory_add_default() and
// gtk_icon_factory_remove_default(). Applications with icons should add a
// default icon factory with their icons, which will allow themes to override
// the icons for the application.
//
// To display an icon, always use gtk_style_lookup_icon_set() on the widget that
// will display the icon, or the convenience function gtk_widget_render_icon().
// These functions take the theme into account when looking up the icon to use
// for a given stock ID.
//
//
// GtkIconFactory as GtkBuildable
//
// GtkIconFactory supports a custom <sources> element, which can contain
// multiple <source> elements. The following attributes are allowed:
//
// - stock-id
//
//    The stock id of the source, a string. This attribute is
//    mandatory
//
// - filename
//
//    The filename of the source, a string.  This attribute is
//    optional
//
// - icon-name
//
//    The icon name for the source, a string.  This attribute is
//    optional.
//
// - size
//
//    Size of the icon, a IconSize enum value.  This attribute is
//    optional.
//
// - direction
//
//    Direction of the source, a TextDirection enum value.  This
//    attribute is optional.
//
// - state
//
//    State of the source, a StateType enum value.  This
//    attribute is optional.
//
// A IconFactory UI definition fragment. ##
//
//    <object class="GtkIconFactory" id="iconfactory1">
//      <sources>
//        <source stock-id="apple-red" filename="apple-red.png"/>
//      </sources>
//    </object>
//    <object class="GtkWindow" id="window1">
//      <child>
//        <object class="GtkButton" id="apple_button">
//          <property name="label">apple-red</property>
//          <property name="use-stock">True</property>
//        </object>
//      </child>
//    </object>
type IconFactory interface {
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

	// Add adds the given @icon_set to the icon factory, under the name
	// @stock_id. @stock_id should be namespaced for your application, e.g.
	// “myapp-whatever-icon”. Normally applications create a IconFactory, then
	// add it to the list of default factories with
	// gtk_icon_factory_add_default(). Then they pass the @stock_id to widgets
	// such as Image to display the icon. Themes can provide an icon with the
	// same name (such as "myapp-whatever-icon") to override your application’s
	// default icons. If an icon already existed in @factory for @stock_id, it
	// is unreferenced and replaced with the new @icon_set.
	//
	// Deprecated: since version 3.10.
	Add(stockId string, iconSet *IconSet)
	// AddDefault adds an icon factory to the list of icon factories searched by
	// gtk_style_lookup_icon_set(). This means that, for example,
	// gtk_image_new_from_stock() will be able to find icons in @factory. There
	// will normally be an icon factory added for each library or application
	// that comes with icons. The default icon factories can be overridden by
	// themes.
	//
	// Deprecated: since version 3.10.
	AddDefault()
	// Lookup looks up @stock_id in the icon factory, returning an icon set if
	// found, otherwise nil. For display to the user, you should use
	// gtk_style_lookup_icon_set() on the Style for the widget that will display
	// the icon, instead of using this function directly, so that themes are
	// taken into account.
	//
	// Deprecated: since version 3.10.
	Lookup(stockId string) *IconSet
	// RemoveDefault removes an icon factory from the list of default icon
	// factories. Not normally used; you might use it for a library that can be
	// unloaded or shut down.
	//
	// Deprecated: since version 3.10.
	RemoveDefault()
}

// iconFactory implements the IconFactory interface.
type iconFactory struct {
	*externglib.Object
}

var _ IconFactory = (*iconFactory)(nil)

// WrapIconFactory wraps a GObject to a type that implements
// interface IconFactory. It is primarily used internally.
func WrapIconFactory(obj *externglib.Object) IconFactory {
	return iconFactory{obj}
}

func marshalIconFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIconFactory(obj), nil
}

// NewIconFactory creates a new IconFactory. An icon factory manages a
// collection of IconSets; a IconSet manages a set of variants of a particular
// icon (i.e. a IconSet contains variants for different sizes and widget
// states). Icons in an icon factory are named by a stock ID, which is a simple
// string identifying the icon. Each Style has a list of IconFactorys derived
// from the current theme; those icon factories are consulted first when
// searching for an icon. If the theme doesn’t set a particular icon, GTK+ looks
// for the icon in a list of default icon factories, maintained by
// gtk_icon_factory_add_default() and gtk_icon_factory_remove_default().
// Applications with icons should add a default icon factory with their icons,
// which will allow themes to override the icons for the application.
//
// Deprecated: since version 3.10.
func NewIconFactory() IconFactory {
	var _cret *C.GtkIconFactory // in

	_cret = C.gtk_icon_factory_new()

	var _iconFactory IconFactory // out

	_iconFactory = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(IconFactory)

	return _iconFactory
}

func (i iconFactory) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(i))
}

func (b iconFactory) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b iconFactory) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b iconFactory) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b iconFactory) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b iconFactory) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b iconFactory) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b iconFactory) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b iconFactory) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b iconFactory) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b iconFactory) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (f iconFactory) Add(stockId string, iconSet *IconSet) {
	var _arg0 *C.GtkIconFactory // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.GtkIconSet     // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	C.gtk_icon_factory_add(_arg0, _arg1, _arg2)
}

func (f iconFactory) AddDefault() {
	var _arg0 *C.GtkIconFactory // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))

	C.gtk_icon_factory_add_default(_arg0)
}

func (f iconFactory) Lookup(stockId string) *IconSet {
	var _arg0 *C.GtkIconFactory // out
	var _arg1 *C.gchar          // out
	var _cret *C.GtkIconSet     // in

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_factory_lookup(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(unsafe.Pointer(_cret))
	C.gtk_icon_set_ref(_cret)
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(unsafe.Pointer(v)))
	})

	return _iconSet
}

func (f iconFactory) RemoveDefault() {
	var _arg0 *C.GtkIconFactory // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))

	C.gtk_icon_factory_remove_default(_arg0)
}
