// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_factory_get_type()), F: marshalIconFactorier},
	})
}

// IconSizeFromName looks up the icon size associated with name.
//
// Deprecated: Use IconTheme instead.
func IconSizeFromName(name string) int {
	var _arg1 *C.gchar      // out
	var _cret C.GtkIconSize // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_size_from_name(_arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IconSizeGetName gets the canonical name of the given icon size. The returned
// string is statically allocated and should not be freed.
//
// Deprecated: Use IconTheme instead.
func IconSizeGetName(size int) string {
	var _arg1 C.GtkIconSize // out
	var _cret *C.gchar      // in

	_arg1 = C.GtkIconSize(size)

	_cret = C.gtk_icon_size_get_name(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IconSizeLookup obtains the pixel size of a semantic icon size size:
// K_ICON_SIZE_MENU, K_ICON_SIZE_BUTTON, etc. This function isn’t normally
// needed, gtk_icon_theme_load_icon() is the usual way to get an icon for
// rendering, then just look at the size of the rendered pixbuf. The rendered
// pixbuf may not even correspond to the width/height returned by
// gtk_icon_size_lookup(), because themes are free to render the pixbuf however
// they like, including changing the usual size.
func IconSizeLookup(size int) (width int, height int, ok bool) {
	var _arg1 C.GtkIconSize // out
	var _arg2 C.gint        // in
	var _arg3 C.gint        // in
	var _cret C.gboolean    // in

	_arg1 = C.GtkIconSize(size)

	_cret = C.gtk_icon_size_lookup(_arg1, &_arg2, &_arg3)

	var _width int  // out
	var _height int // out
	var _ok bool    // out

	_width = int(_arg2)
	_height = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _width, _height, _ok
}

// IconSizeLookupForSettings obtains the pixel size of a semantic icon size,
// possibly modified by user preferences for a particular Settings. Normally
// size would be K_ICON_SIZE_MENU, K_ICON_SIZE_BUTTON, etc. This function isn’t
// normally needed, gtk_widget_render_icon_pixbuf() is the usual way to get an
// icon for rendering, then just look at the size of the rendered pixbuf. The
// rendered pixbuf may not even correspond to the width/height returned by
// gtk_icon_size_lookup(), because themes are free to render the pixbuf however
// they like, including changing the usual size.
//
// Deprecated: Use gtk_icon_size_lookup() instead.
func IconSizeLookupForSettings(settings *Settings, size int) (width int, height int, ok bool) {
	var _arg1 *C.GtkSettings // out
	var _arg2 C.GtkIconSize  // out
	var _arg3 C.gint         // in
	var _arg4 C.gint         // in
	var _cret C.gboolean     // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_icon_size_lookup_for_settings(_arg1, _arg2, &_arg3, &_arg4)

	var _width int  // out
	var _height int // out
	var _ok bool    // out

	_width = int(_arg3)
	_height = int(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _width, _height, _ok
}

// IconSizeRegister registers a new icon size, along the same lines as
// K_ICON_SIZE_MENU, etc. Returns the integer value for the size.
//
// Deprecated: Use IconTheme instead.
func IconSizeRegister(name string, width int, height int) int {
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out
	var _cret C.GtkIconSize // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(width)
	_arg3 = C.gint(height)

	_cret = C.gtk_icon_size_register(_arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IconSizeRegisterAlias registers alias as another name for target. So calling
// gtk_icon_size_from_name() with alias as argument will return target.
//
// Deprecated: Use IconTheme instead.
func IconSizeRegisterAlias(alias string, target int) {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(alias)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(target)

	C.gtk_icon_size_register_alias(_arg1, _arg2)
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
type IconFactory struct {
	*externglib.Object

	Buildable
}

func wrapIconFactory(obj *externglib.Object) *IconFactory {
	return &IconFactory{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalIconFactorier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconFactory(obj), nil
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
// Deprecated: Use IconTheme instead.
func NewIconFactory() *IconFactory {
	var _cret *C.GtkIconFactory // in

	_cret = C.gtk_icon_factory_new()

	var _iconFactory *IconFactory // out

	_iconFactory = wrapIconFactory(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconFactory
}

// Add adds the given icon_set to the icon factory, under the name stock_id.
// stock_id should be namespaced for your application, e.g.
// “myapp-whatever-icon”. Normally applications create a IconFactory, then add
// it to the list of default factories with gtk_icon_factory_add_default(). Then
// they pass the stock_id to widgets such as Image to display the icon. Themes
// can provide an icon with the same name (such as "myapp-whatever-icon") to
// override your application’s default icons. If an icon already existed in
// factory for stock_id, it is unreferenced and replaced with the new icon_set.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) Add(stockId string, iconSet *IconSet) {
	var _arg0 *C.GtkIconFactory // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.GtkIconSet     // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(factory.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(iconSet)))

	C.gtk_icon_factory_add(_arg0, _arg1, _arg2)
}

// AddDefault adds an icon factory to the list of icon factories searched by
// gtk_style_lookup_icon_set(). This means that, for example,
// gtk_image_new_from_stock() will be able to find icons in factory. There will
// normally be an icon factory added for each library or application that comes
// with icons. The default icon factories can be overridden by themes.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) AddDefault() {
	var _arg0 *C.GtkIconFactory // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_icon_factory_add_default(_arg0)
}

// Lookup looks up stock_id in the icon factory, returning an icon set if found,
// otherwise NULL. For display to the user, you should use
// gtk_style_lookup_icon_set() on the Style for the widget that will display the
// icon, instead of using this function directly, so that themes are taken into
// account.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) Lookup(stockId string) *IconSet {
	var _arg0 *C.GtkIconFactory // out
	var _arg1 *C.gchar          // out
	var _cret *C.GtkIconSet     // in

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(factory.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_factory_lookup(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _iconSet
}

// RemoveDefault removes an icon factory from the list of default icon
// factories. Not normally used; you might use it for a library that can be
// unloaded or shut down.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) RemoveDefault() {
	var _arg0 *C.GtkIconFactory // out

	_arg0 = (*C.GtkIconFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_icon_factory_remove_default(_arg0)
}

// IconFactoryLookupDefault looks for an icon in the list of default icon
// factories. For display to the user, you should use
// gtk_style_lookup_icon_set() on the Style for the widget that will display the
// icon, instead of using this function directly, so that themes are taken into
// account.
//
// Deprecated: Use IconTheme instead.
func IconFactoryLookupDefault(stockId string) *IconSet {
	var _arg1 *C.gchar      // out
	var _cret *C.GtkIconSet // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_factory_lookup_default(_arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _iconSet
}
