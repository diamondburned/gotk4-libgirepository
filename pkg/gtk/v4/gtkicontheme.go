// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_theme_error_get_type()), F: marshalIconThemeError},
		{T: externglib.Type(C.gtk_icon_lookup_flags_get_type()), F: marshalIconLookupFlags},
		{T: externglib.Type(C.gtk_icon_paintable_get_type()), F: marshalIconPaintabler},
		{T: externglib.Type(C.gtk_icon_theme_get_type()), F: marshalIconThemer},
	})
}

// IconThemeError: error codes for `GtkIconTheme` operations.
type IconThemeError int

const (
	// NotFound: icon specified does not exist in the theme
	IconThemeErrorNotFound IconThemeError = iota
	// Failed: unspecified error occurred.
	IconThemeErrorFailed
)

func marshalIconThemeError(p uintptr) (interface{}, error) {
	return IconThemeError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IconLookupFlags: used to specify options for gtk_icon_theme_lookup_icon().
type IconLookupFlags int

const (
	// IconLookupFlagsForceRegular: try to always load regular icons, even when
	// symbolic icon names are given
	IconLookupFlagsForceRegular IconLookupFlags = 0b1
	// IconLookupFlagsForceSymbolic: try to always load symbolic icons, even
	// when regular icon names are given
	IconLookupFlagsForceSymbolic IconLookupFlags = 0b10
	// IconLookupFlagsPreload starts loading the texture in the background so it
	// is ready when later needed.
	IconLookupFlagsPreload IconLookupFlags = 0b100
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	return IconLookupFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IconPaintabler describes IconPaintable's methods.
type IconPaintabler interface {
	// File gets the `GFile` that was used to load the icon.
	File() *gio.File
	// IconName: get the icon name being used for this icon.
	IconName() string
	// IsSymbolic checks if the icon is symbolic or not.
	IsSymbolic() bool
}

// IconPaintable contains information found when looking up an icon in
// `GtkIconTheme`.
//
// `GtkIconPaintable` implements `GdkPaintable`.
type IconPaintable struct {
	*externglib.Object

	gdk.Paintable
}

var (
	_ IconPaintabler  = (*IconPaintable)(nil)
	_ gextras.Nativer = (*IconPaintable)(nil)
)

func wrapIconPaintable(obj *externglib.Object) IconPaintabler {
	return &IconPaintable{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalIconPaintabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconPaintable(obj), nil
}

// NewIconPaintableForFile creates a `GtkIconPaintable` for a file with a given
// size and scale.
//
// The icon can then be rendered by using it as a `GdkPaintable`.
func NewIconPaintableForFile(file gio.Filer, size int, scale int) *IconPaintable {
	var _arg1 *C.GFile            // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _cret *C.GtkIconPaintable // in

	_arg1 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))
	_arg2 = C.int(size)
	_arg3 = C.int(scale)

	_cret = C.gtk_icon_paintable_new_for_file(_arg1, _arg2, _arg3)

	var _iconPaintable *IconPaintable // out

	_iconPaintable = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*IconPaintable)

	return _iconPaintable
}

// File gets the `GFile` that was used to load the icon.
//
// Returns nil if the icon was not loaded from a file.
func (self *IconPaintable) File() *gio.File {
	var _arg0 *C.GtkIconPaintable // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_get_file(_arg0)

	var _file *gio.File // out

	_file = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.File)

	return _file
}

// IconName: get the icon name being used for this icon.
//
// When an icon looked up in the icon theme was not available, the icon theme
// may use fallback icons - either those specified to
// gtk_icon_theme_lookup_icon() or the always-available "image-missing". The
// icon chosen is returned by this function.
//
// If the icon was created without an icon theme, this function returns nil.
func (self *IconPaintable) IconName() string {
	var _arg0 *C.GtkIconPaintable // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_get_icon_name(_arg0)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(_cret))

	return _filename
}

// IsSymbolic checks if the icon is symbolic or not.
//
// This currently uses only the file name and not the file contents for
// determining this. This behaviour may change in the future.
//
// Note that to render a symbolic `GtkIconPaintable` properly (with recoloring),
// you have to set its icon name on a `GtkImage`.
func (self *IconPaintable) IsSymbolic() bool {
	var _arg0 *C.GtkIconPaintable // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_is_symbolic(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconThemer describes IconTheme's methods.
type IconThemer interface {
	// AddResourcePath adds a resource path that will be looked at when looking
	// for icons, similar to search paths.
	AddResourcePath(path string)
	// AddSearchPath appends a directory to the search path.
	AddSearchPath(path string)
	// Display returns the display that the `GtkIconTheme` object was created
	// for.
	Display() *gdk.Display
	// IconNames lists the names of icons in the current icon theme.
	IconNames() []string
	// IconSizes returns an array of integers describing the sizes at which the
	// icon is available without scaling.
	IconSizes(iconName string) []int
	// ResourcePath gets the current resource path.
	ResourcePath() []string
	// SearchPath gets the current search path.
	SearchPath() []string
	// ThemeName gets the current icon theme name.
	ThemeName() string
	// HasGIcon checks whether an icon theme includes an icon for a particular
	// `GIcon`.
	HasGIcon(gicon gio.Iconner) bool
	// HasIcon checks whether an icon theme includes an icon for a particular
	// name.
	HasIcon(iconName string) bool
	// SetSearchPath sets the search path for the icon theme object.
	SetSearchPath(path []string)
	// SetThemeName sets the name of the icon theme that the `GtkIconTheme`
	// object uses overriding system configuration.
	SetThemeName(themeName string)
}

// IconTheme: `GtkIconTheme` provides a facility for loading themed icons.
//
// The main reason for using a name rather than simply providing a filename is
// to allow different icons to be used depending on what “icon theme” is
// selected by the user. The operation of icon themes on Linux and Unix follows
// the Icon Theme Specification
// (http://www.freedesktop.org/Standards/icon-theme-spec) There is a fallback
// icon theme, named `hicolor`, where applications should install their icons,
// but additional icon themes can be installed as operating system vendors and
// users choose.
//
// In many cases, named themes are used indirectly, via [class@Gtk.Image] rather
// than directly, but looking up icons directly is also simple. The
// `GtkIconTheme` object acts as a database of all the icons in the current
// theme. You can create new `GtkIconTheme` objects, but it’s much more
// efficient to use the standard icon theme of the `GtkWidget` so that the icon
// information is shared with other people looking up icons.
//
// “`c GtkIconTheme *icon_theme; GtkIconPaintable *icon; GdkPaintable
// *paintable;
//
// icon_theme = gtk_icon_theme_get_for_display (gtk_widget_get_display
// (my_widget)); icon = gtk_icon_theme_lookup_icon (icon_theme, "my-icon-name",
// // icon name 48, // icon size 1, // scale 0, // flags); paintable =
// GDK_PAINTABLE (icon); // Use the paintable g_object_unref (icon); “`
type IconTheme struct {
	*externglib.Object
}

var (
	_ IconThemer      = (*IconTheme)(nil)
	_ gextras.Nativer = (*IconTheme)(nil)
)

func wrapIconTheme(obj *externglib.Object) IconThemer {
	return &IconTheme{
		Object: obj,
	}
}

func marshalIconThemer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconTheme(obj), nil
}

// NewIconTheme creates a new icon theme object.
//
// Icon theme objects are used to lookup up an icon by name in a particular icon
// theme. Usually, you’ll want to use [func@Gtk.IconTheme.get_for_display]
// rather than creating a new icon theme object for scratch.
func NewIconTheme() *IconTheme {
	var _cret *C.GtkIconTheme // in

	_cret = C.gtk_icon_theme_new()

	var _iconTheme *IconTheme // out

	_iconTheme = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*IconTheme)

	return _iconTheme
}

// AddResourcePath adds a resource path that will be looked at when looking for
// icons, similar to search paths.
//
// See [method@Gtk.IconTheme.set_resource_path].
//
// This function should be used to make application-specific icons available as
// part of the icon theme.
func (self *IconTheme) AddResourcePath(path string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_add_resource_path(_arg0, _arg1)
}

// AddSearchPath appends a directory to the search path.
//
// See [method@Gtk.IconTheme.set_search_path].
func (self *IconTheme) AddSearchPath(path string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_add_search_path(_arg0, _arg1)
}

// Display returns the display that the `GtkIconTheme` object was created for.
func (self *IconTheme) Display() *gdk.Display {
	var _arg0 *C.GtkIconTheme // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_display(_arg0)

	var _display *gdk.Display // out

	_display = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Display)

	return _display
}

// IconNames lists the names of icons in the current icon theme.
func (self *IconTheme) IconNames() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_icon_names(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(src[i]))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// IconSizes returns an array of integers describing the sizes at which the icon
// is available without scaling.
//
// A size of -1 means that the icon is available in a scalable format. The array
// is zero-terminated.
func (self *IconTheme) IconSizes(iconName string) []int {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out
	var _cret *C.int

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_theme_get_icon_sizes(_arg0, _arg1)

	var _gints []int

	{
		var i int
		var z C.int
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_gints = make([]int, i)
		for i := range src {
			_gints[i] = int(src[i])
		}
	}

	return _gints
}

// ResourcePath gets the current resource path.
//
// See [method@Gtk.IconTheme.set_resource_path].
func (self *IconTheme) ResourcePath() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_resource_path(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(src[i]))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// SearchPath gets the current search path.
//
// See [method@Gtk.IconTheme.set_search_path].
func (self *IconTheme) SearchPath() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_search_path(_arg0)

	var _filenames []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(src[i]))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// ThemeName gets the current icon theme name.
//
// Returns (transfer full): the current icon theme name,
func (self *IconTheme) ThemeName() string {
	var _arg0 *C.GtkIconTheme // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_theme_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HasGIcon checks whether an icon theme includes an icon for a particular
// `GIcon`.
func (self *IconTheme) HasGIcon(gicon gio.Iconner) bool {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.GIcon        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer((gicon).(gextras.Nativer).Native()))

	_cret = C.gtk_icon_theme_has_gicon(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasIcon checks whether an icon theme includes an icon for a particular name.
func (self *IconTheme) HasIcon(iconName string) bool {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_theme_has_icon(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSearchPath sets the search path for the icon theme object.
//
// When looking for an icon theme, GTK will search for a subdirectory of one or
// more of the directories in @path with the same name as the icon theme
// containing an index.theme file. (Themes from multiple of the path elements
// are combined to allow themes to be extended by adding icons in the user’s
// home directory.)
//
// In addition if an icon found isn’t found either in the current icon theme or
// the default icon theme, and an image file with the right name is found
// directly in one of the elements of @path, then that image will be used for
// the icon name. (This is legacy feature, and new icons should be put into the
// fallback icon theme, which is called hicolor, rather than directly on the
// icon path.)
func (self *IconTheme) SetSearchPath(path []string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 **C.char

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(path)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(path))
		for i := range path {
			out[i] = (*C.char)(C.CString(path[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_icon_theme_set_search_path(_arg0, _arg1)
}

// SetThemeName sets the name of the icon theme that the `GtkIconTheme` object
// uses overriding system configuration.
//
// This function cannot be called on the icon theme objects returned from
// [type_func@Gtk.IconTheme.get_for_display].
func (self *IconTheme) SetThemeName(themeName string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(themeName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_set_theme_name(_arg0, _arg1)
}
