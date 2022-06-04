// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_IconTheme_ConnectChanged(gpointer, guintptr);
import "C"

// glib.Type values for gtkicontheme.go.
var (
	GTypeIconThemeError  = coreglib.Type(C.gtk_icon_theme_error_get_type())
	GTypeIconLookupFlags = coreglib.Type(C.gtk_icon_lookup_flags_get_type())
	GTypeIconPaintable   = coreglib.Type(C.gtk_icon_paintable_get_type())
	GTypeIconTheme       = coreglib.Type(C.gtk_icon_theme_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeIconThemeError, F: marshalIconThemeError},
		{T: GTypeIconLookupFlags, F: marshalIconLookupFlags},
		{T: GTypeIconPaintable, F: marshalIconPaintable},
		{T: GTypeIconTheme, F: marshalIconTheme},
	})
}

// IconThemeError: error codes for GtkIconTheme operations.
type IconThemeError C.gint

const (
	// IconThemeNotFound: icon specified does not exist in the theme.
	IconThemeNotFound IconThemeError = iota
	// IconThemeFailed: unspecified error occurred.
	IconThemeFailed
)

func marshalIconThemeError(p uintptr) (interface{}, error) {
	return IconThemeError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for IconThemeError.
func (i IconThemeError) String() string {
	switch i {
	case IconThemeNotFound:
		return "NotFound"
	case IconThemeFailed:
		return "Failed"
	default:
		return fmt.Sprintf("IconThemeError(%d)", i)
	}
}

// IconLookupFlags: used to specify options for gtk_icon_theme_lookup_icon().
type IconLookupFlags C.guint

const (
	// IconLookupForceRegular: try to always load regular icons, even when
	// symbolic icon names are given.
	IconLookupForceRegular IconLookupFlags = 0b1
	// IconLookupForceSymbolic: try to always load symbolic icons, even when
	// regular icon names are given.
	IconLookupForceSymbolic IconLookupFlags = 0b10
	// IconLookupPreload starts loading the texture in the background so it is
	// ready when later needed.
	IconLookupPreload IconLookupFlags = 0b100
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	return IconLookupFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for IconLookupFlags.
func (i IconLookupFlags) String() string {
	if i == 0 {
		return "IconLookupFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(64)

	for i != 0 {
		next := i & (i - 1)
		bit := i - next

		switch bit {
		case IconLookupForceRegular:
			builder.WriteString("ForceRegular|")
		case IconLookupForceSymbolic:
			builder.WriteString("ForceSymbolic|")
		case IconLookupPreload:
			builder.WriteString("Preload|")
		default:
			builder.WriteString(fmt.Sprintf("IconLookupFlags(0b%b)|", bit))
		}

		i = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if i contains other.
func (i IconLookupFlags) Has(other IconLookupFlags) bool {
	return (i & other) == other
}

// IconPaintable contains information found when looking up an icon in
// GtkIconTheme.
//
// GtkIconPaintable implements GdkPaintable.
type IconPaintable struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gdk.Paintable
}

var (
	_ coreglib.Objector = (*IconPaintable)(nil)
)

func wrapIconPaintable(obj *coreglib.Object) *IconPaintable {
	return &IconPaintable{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalIconPaintable(p uintptr) (interface{}, error) {
	return wrapIconPaintable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIconPaintableForFile creates a GtkIconPaintable for a file with a given
// size and scale.
//
// The icon can then be rendered by using it as a GdkPaintable.
//
// The function takes the following parameters:
//
//    - file: GFile.
//    - size: desired icon size.
//    - scale: desired scale.
//
// The function returns the following values:
//
//    - iconPaintable: GtkIconPaintable containing for the icon. Unref with
//      g_object_unref().
//
func NewIconPaintableForFile(file gio.Filer, size, scale int32) *IconPaintable {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(size)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(scale)

	_gret := girepository.MustFind("Gtk", "IconPaintable").InvokeMethod("new_IconPaintable_for_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(file)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scale)

	var _iconPaintable *IconPaintable // out

	_iconPaintable = wrapIconPaintable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconPaintable
}

// File gets the GFile that was used to load the icon.
//
// Returns NULL if the icon was not loaded from a file.
//
// The function returns the following values:
//
//    - file (optional) for the icon, or NULL. Free with g_object_unref().
//
func (self *IconPaintable) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconPaintable").InvokeMethod("get_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _file *gio.File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// IconName: get the icon name being used for this icon.
//
// When an icon looked up in the icon theme was not available, the icon theme
// may use fallback icons - either those specified to
// gtk_icon_theme_lookup_icon() or the always-available "image-missing". The
// icon chosen is returned by this function.
//
// If the icon was created without an icon theme, this function returns NULL.
//
// The function returns the following values:
//
//    - filename (optional): themed icon-name for the icon, or NULL if its not a
//      themed icon.
//
func (self *IconPaintable) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconPaintable").InvokeMethod("get_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _filename string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}

// IsSymbolic checks if the icon is symbolic or not.
//
// This currently uses only the file name and not the file contents for
// determining this. This behaviour may change in the future.
//
// Note that to render a symbolic GtkIconPaintable properly (with recoloring),
// you have to set its icon name on a GtkImage.
//
// The function returns the following values:
//
//    - ok: TRUE if the icon is symbolic, FALSE otherwise.
//
func (self *IconPaintable) IsSymbolic() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconPaintable").InvokeMethod("is_symbolic", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IconTheme: GtkIconTheme provides a facility for loading themed icons.
//
// The main reason for using a name rather than simply providing a filename is
// to allow different icons to be used depending on what “icon theme” is
// selected by the user. The operation of icon themes on Linux and Unix follows
// the Icon Theme Specification
// (http://www.freedesktop.org/Standards/icon-theme-spec) There is a fallback
// icon theme, named hicolor, where applications should install their icons, but
// additional icon themes can be installed as operating system vendors and users
// choose.
//
// In many cases, named themes are used indirectly, via gtk.Image rather than
// directly, but looking up icons directly is also simple. The GtkIconTheme
// object acts as a database of all the icons in the current theme. You can
// create new GtkIconTheme objects, but it’s much more efficient to use the
// standard icon theme of the GtkWidget so that the icon information is shared
// with other people looking up icons.
//
//    GtkIconTheme *icon_theme;
//    GtkIconPaintable *icon;
//    GdkPaintable *paintable;
//
//    icon_theme = gtk_icon_theme_get_for_display (gtk_widget_get_display (my_widget));
//    icon = gtk_icon_theme_lookup_icon (icon_theme,
//                                       "my-icon-name", // icon name
//                                       48, // icon size
//                                       1,  // scale
//                                       0,  // flags);
//    paintable = GDK_PAINTABLE (icon);
//    // Use the paintable
//    g_object_unref (icon);.
type IconTheme struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*IconTheme)(nil)
)

func wrapIconTheme(obj *coreglib.Object) *IconTheme {
	return &IconTheme{
		Object: obj,
	}
}

func marshalIconTheme(p uintptr) (interface{}, error) {
	return wrapIconTheme(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_IconTheme_ConnectChanged
func _gotk4_gtk4_IconTheme_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted when the icon theme changes.
//
// This can happen becuase current icon theme is switched or because GTK detects
// that a change has occurred in the contents of the current icon theme.
func (self *IconTheme) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "changed", false, unsafe.Pointer(C._gotk4_gtk4_IconTheme_ConnectChanged), f)
}

// NewIconTheme creates a new icon theme object.
//
// Icon theme objects are used to lookup up an icon by name in a particular icon
// theme. Usually, you’ll want to use gtk.IconTheme().GetForDisplay rather than
// creating a new icon theme object for scratch.
//
// The function returns the following values:
//
//    - iconTheme: newly created GtkIconTheme object.
//
func NewIconTheme() *IconTheme {
	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("new_IconTheme", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconTheme
}

// AddResourcePath adds a resource path that will be looked at when looking for
// icons, similar to search paths.
//
// See gtk.IconTheme.SetResourcePath().
//
// This function should be used to make application-specific icons available as
// part of the icon theme.
//
// The function takes the following parameters:
//
//    - path: resource path.
//
func (self *IconTheme) AddResourcePath(path string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "IconTheme").InvokeMethod("add_resource_path", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// AddSearchPath appends a directory to the search path.
//
// See gtk.IconTheme.SetSearchPath().
//
// The function takes the following parameters:
//
//    - path: directory name to append to the icon path.
//
func (self *IconTheme) AddSearchPath(path string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "IconTheme").InvokeMethod("add_search_path", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// Display returns the display that the GtkIconTheme object was created for.
//
// The function returns the following values:
//
//    - display (optional) of icon_theme.
//
func (self *IconTheme) Display() *gdk.Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _display *gdk.Display // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_display = &gdk.Display{
				Object: obj,
			}
		}
	}

	return _display
}

// IconNames lists the names of icons in the current icon theme.
//
// The function returns the following values:
//
//    - utf8s: string array holding the names of all the icons in the theme. You
//      must free the array using g_strfreev().
//
func (self *IconTheme) IconNames() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_icon_names", _args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
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
//
// The function takes the following parameters:
//
//    - iconName: name of an icon.
//
// The function returns the following values:
//
//    - gints: newly allocated array describing the sizes at which the icon is
//      available. The array should be freed with g_free() when it is no longer
//      needed.
//
func (self *IconTheme) IconSizes(iconName string) []int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_icon_sizes", _args[:], nil)
	_cret = *(**C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)

	var _gints []int32 // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z C.int
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_gints = make([]int32, i)
		for i := range src {
			_gints[i] = int32(*(*C.int)(unsafe.Pointer(&src[i])))
		}
	}

	return _gints
}

// ResourcePath gets the current resource path.
//
// See gtk.IconTheme.SetResourcePath().
//
// The function returns the following values:
//
//    - utf8s (optional): A list of resource paths or NULL. The returned value
//      should be freed with g_strfreev().
//
func (self *IconTheme) ResourcePath() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_resource_path", _args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8s []string // out

	if *(***C.void)(unsafe.Pointer(&_cret)) != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.void
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// SearchPath gets the current search path.
//
// See gtk.IconTheme.SetSearchPath().
//
// The function returns the following values:
//
//    - filenames (optional): a list of icon theme path directories or NULL. The
//      returned value should be freed with g_strfreev().
//
func (self *IconTheme) SearchPath() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_search_path", _args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _filenames []string // out

	if *(***C.void)(unsafe.Pointer(&_cret)) != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.void
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_filenames = make([]string, i)
			for i := range src {
				_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _filenames
}

// ThemeName gets the current icon theme name.
//
// Returns (transfer full): the current icon theme name,.
//
// The function returns the following values:
//
func (self *IconTheme) ThemeName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("get_theme_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HasGIcon checks whether an icon theme includes an icon for a particular
// GIcon.
//
// The function takes the following parameters:
//
//    - gicon: GIcon.
//
// The function returns the following values:
//
//    - ok: TRUE if self includes an icon for gicon.
//
func (self *IconTheme) HasGIcon(gicon gio.Iconner) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gicon).Native()))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("has_gicon", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(gicon)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HasIcon checks whether an icon theme includes an icon for a particular name.
//
// The function takes the following parameters:
//
//    - iconName: name of an icon.
//
// The function returns the following values:
//
//    - ok: TRUE if self includes an icon for icon_name.
//
func (self *IconTheme) HasIcon(iconName string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "IconTheme").InvokeMethod("has_icon", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetSearchPath sets the search path for the icon theme object.
//
// When looking for an icon theme, GTK will search for a subdirectory of one or
// more of the directories in path with the same name as the icon theme
// containing an index.theme file. (Themes from multiple of the path elements
// are combined to allow themes to be extended by adding icons in the user’s
// home directory.)
//
// In addition if an icon found isn’t found either in the current icon theme or
// the default icon theme, and an image file with the right name is found
// directly in one of the elements of path, then that image will be used for the
// icon name. (This is legacy feature, and new icons should be put into the
// fallback icon theme, which is called hicolor, rather than directly on the
// icon path.).
//
// The function takes the following parameters:
//
//    - path (optional): NULL-terminated array of directories that are searched
//      for icon themes.
//
func (self *IconTheme) SetSearchPath(path []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(path) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(path)+1)
			var zero *C.void
			out[len(path)] = zero
			for i := range path {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(path[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "IconTheme").InvokeMethod("set_search_path", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// SetThemeName sets the name of the icon theme that the GtkIconTheme object
// uses overriding system configuration.
//
// This function cannot be called on the icon theme objects returned from
// gtk.IconTheme.GetForDisplay.
//
// The function takes the following parameters:
//
//    - themeName (optional): name of icon theme to use instead of configured
//      theme, or NULL to unset a previously set custom theme.
//
func (self *IconTheme) SetThemeName(themeName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if themeName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(themeName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "IconTheme").InvokeMethod("set_theme_name", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(themeName)
}

// IconThemeGetForDisplay gets the icon theme object associated with display.
//
// If this function has not previously been called for the given display, a new
// icon theme object will be created and associated with the display. Icon theme
// objects are fairly expensive to create, so using this function is usually a
// better choice than calling gtk.IconTheme.New and setting the display
// yourself; by using this function a single icon theme object will be shared
// between users.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//
// The function returns the following values:
//
//    - iconTheme: unique GtkIconTheme associated with the given display. This
//      icon theme is associated with the display and can be used as long as the
//      display is open. Do not ref or unref it.
//
func IconThemeGetForDisplay(display *gdk.Display) *IconTheme {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("Gtk", "get_for_display").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(coreglib.Take(unsafe.Pointer(_cret)))

	return _iconTheme
}
