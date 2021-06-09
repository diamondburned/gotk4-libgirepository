// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
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

// IconFactory: an icon factory manages a collection of IconSet; a IconSet
// manages a set of variants of a particular icon (i.e. a IconSet contains
// variants for different sizes and widget states). Icons in an icon factory are
// named by a stock ID, which is a simple string identifying the icon. Each
// Style has a list of IconFactory derived from the current theme; those icon
// factories are consulted first when searching for an icon. If the theme
// doesn’t set a particular icon, GTK+ looks for the icon in a list of default
// icon factories, maintained by gtk_icon_factory_add_default() and
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
	Buildable

	// Add adds the given @icon_set to the icon factory, under the name
	// @stock_id. @stock_id should be namespaced for your application, e.g.
	// “myapp-whatever-icon”. Normally applications create a IconFactory, then
	// add it to the list of default factories with
	// gtk_icon_factory_add_default(). Then they pass the @stock_id to widgets
	// such as Image to display the icon. Themes can provide an icon with the
	// same name (such as "myapp-whatever-icon") to override your application’s
	// default icons. If an icon already existed in @factory for @stock_id, it
	// is unreferenced and replaced with the new @icon_set.
	Add(stockId string, iconSet *IconSet)
	// AddDefault adds an icon factory to the list of icon factories searched by
	// gtk_style_lookup_icon_set(). This means that, for example,
	// gtk_image_new_from_stock() will be able to find icons in @factory. There
	// will normally be an icon factory added for each library or application
	// that comes with icons. The default icon factories can be overridden by
	// themes.
	AddDefault()
	// Lookup looks up @stock_id in the icon factory, returning an icon set if
	// found, otherwise nil. For display to the user, you should use
	// gtk_style_lookup_icon_set() on the Style for the widget that will display
	// the icon, instead of using this function directly, so that themes are
	// taken into account.
	Lookup(stockId string) *IconSet
	// RemoveDefault removes an icon factory from the list of default icon
	// factories. Not normally used; you might use it for a library that can be
	// unloaded or shut down.
	RemoveDefault()
}

// iconFactory implements the IconFactory interface.
type iconFactory struct {
	gextras.Objector
	Buildable
}

var _ IconFactory = (*iconFactory)(nil)

// WrapIconFactory wraps a GObject to the right type. It is
// primarily used internally.
func WrapIconFactory(obj *externglib.Object) IconFactory {
	return IconFactory{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalIconFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIconFactory(obj), nil
}

// NewIconFactory constructs a class IconFactory.
func NewIconFactory() IconFactory {
	var _cret C.GtkIconFactory

	cret = C.gtk_icon_factory_new()

	var _iconFactory IconFactory

	_iconFactory = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(IconFactory)

	return _iconFactory
}

// Add adds the given @icon_set to the icon factory, under the name
// @stock_id. @stock_id should be namespaced for your application, e.g.
// “myapp-whatever-icon”. Normally applications create a IconFactory, then
// add it to the list of default factories with
// gtk_icon_factory_add_default(). Then they pass the @stock_id to widgets
// such as Image to display the icon. Themes can provide an icon with the
// same name (such as "myapp-whatever-icon") to override your application’s
// default icons. If an icon already existed in @factory for @stock_id, it
// is unreferenced and replaced with the new @icon_set.
func (f iconFactory) Add(stockId string, iconSet *IconSet) {
	var _arg0 *C.GtkIconFactory
	var _arg1 *C.gchar
	var _arg2 *C.GtkIconSet

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkIconSet)(unsafe.Pointer(iconSet.Native()))

	C.gtk_icon_factory_add(_arg0, _arg1, _arg2)
}

// AddDefault adds an icon factory to the list of icon factories searched by
// gtk_style_lookup_icon_set(). This means that, for example,
// gtk_image_new_from_stock() will be able to find icons in @factory. There
// will normally be an icon factory added for each library or application
// that comes with icons. The default icon factories can be overridden by
// themes.
func (f iconFactory) AddDefault() {
	var _arg0 *C.GtkIconFactory

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))

	C.gtk_icon_factory_add_default(_arg0)
}

// Lookup looks up @stock_id in the icon factory, returning an icon set if
// found, otherwise nil. For display to the user, you should use
// gtk_style_lookup_icon_set() on the Style for the widget that will display
// the icon, instead of using this function directly, so that themes are
// taken into account.
func (f iconFactory) Lookup(stockId string) *IconSet {
	var _arg0 *C.GtkIconFactory
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GtkIconSet

	cret = C.gtk_icon_factory_lookup(_arg0, _arg1)

	var _iconSet *IconSet

	_iconSet = WrapIconSet(unsafe.Pointer(_cret))

	return _iconSet
}

// RemoveDefault removes an icon factory from the list of default icon
// factories. Not normally used; you might use it for a library that can be
// unloaded or shut down.
func (f iconFactory) RemoveDefault() {
	var _arg0 *C.GtkIconFactory

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(f.Native()))

	C.gtk_icon_factory_remove_default(_arg0)
}
