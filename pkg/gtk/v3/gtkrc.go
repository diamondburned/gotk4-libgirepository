// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_rc_style_get_type()), F: marshalRCStyle},
	})
}

// RCAddDefaultFile adds a file to the list of files to be parsed at the end of
// gtk_init().
func RCAddDefaultFile(filename *string) {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_add_default_file(_arg1)
}

// RCFindModuleInPath searches for a theme engine in the GTK+ search path. This
// function is not useful for applications and should not be used.
func RCFindModuleInPath(moduleFile string) *string {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(moduleFile))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.gtk_rc_find_module_in_path(_arg1)

	var _filename *string

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCFindPixmapInPath looks up a file in pixmap path for the specified Settings.
// If the file is not found, it outputs a warning message using g_warning() and
// returns nil.
func RCFindPixmapInPath(settings Settings, scanner *glib.Scanner, pixmapFile string) *string {
	var _arg1 *C.GtkSettings
	var _arg2 *C.GScanner
	var _arg3 *C.gchar

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.GScanner)(unsafe.Pointer(scanner.Native()))
	_arg3 = (*C.gchar)(C.CString(pixmapFile))
	defer C.free(unsafe.Pointer(_arg3))

	var _cret *C.gchar

	cret = C.gtk_rc_find_pixmap_in_path(_arg1, _arg2, _arg3)

	var _filename *string

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetDefaultFiles retrieves the current list of RC files that will be parsed
// at the end of gtk_init().
func RCGetDefaultFiles() []*string {
	var _cret **C.gchar

	cret = C.gtk_rc_get_default_files()

	var _filenames []*string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_filenames = make([]*string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_filenames = C.GoString(_cret)
		}
	}

	return _filenames
}

// RCGetImModuleFile obtains the path to the IM modules file. See the
// documentation of the `GTK_IM_MODULE_FILE` environment variable for more
// details.
func RCGetImModuleFile() *string {
	var _cret *C.gchar

	cret = C.gtk_rc_get_im_module_file()

	var _filename *string

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetImModulePath obtains the path in which to look for IM modules. See the
// documentation of the `GTK_PATH` environment variable for more details about
// looking up modules. This function is useful solely for utilities supplied
// with GTK+ and should not be used by applications under normal circumstances.
func RCGetImModulePath() *string {
	var _cret *C.gchar

	cret = C.gtk_rc_get_im_module_path()

	var _filename *string

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetModuleDir returns a directory in which GTK+ looks for theme engines. For
// full information about the search for theme engines, see the docs for
// `GTK_PATH` in [Running GTK+ Applications][gtk-running].
func RCGetModuleDir() *string {
	var _cret *C.gchar

	cret = C.gtk_rc_get_module_dir()

	var _filename *string

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetStyle finds all matching RC styles for a given widget, composites them
// together, and then creates a Style representing the composite appearance.
// (GTK+ actually keeps a cache of previously created styles, so a new style may
// not be created.)
func RCGetStyle(widget Widget) Style {
	var _arg1 *C.GtkWidget

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret *C.GtkStyle

	cret = C.gtk_rc_get_style(_arg1)

	var _style Style

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Style)

	return _style
}

// RCGetStyleByPaths creates up a Style from styles defined in a RC file by
// providing the raw components used in matching. This function may be useful
// when creating pseudo-widgets that should be themed like widgets but don’t
// actually have corresponding GTK+ widgets. An example of this would be items
// inside a GNOME canvas widget.
//
// The action of gtk_rc_get_style() is similar to:
//
//    gtk_widget_path (widget, NULL, &path, NULL);
//    gtk_widget_class_path (widget, NULL, &class_path, NULL);
//    gtk_rc_get_style_by_paths (gtk_widget_get_settings (widget),
//                               path, class_path,
//                               G_OBJECT_TYPE (widget));
func RCGetStyleByPaths(settings Settings, widgetPath string, classPath string, typ externglib.Type) Style {
	var _arg1 *C.GtkSettings
	var _arg2 *C.char
	var _arg3 *C.char
	var _arg4 C.GType

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.char)(C.CString(widgetPath))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(classPath))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(typ)

	var _cret *C.GtkStyle

	cret = C.gtk_rc_get_style_by_paths(_arg1, _arg2, _arg3, _arg4)

	var _style Style

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Style)

	return _style
}

// RCGetThemeDir returns the standard directory in which themes should be
// installed. (GTK+ does not actually use this directory itself.)
func RCGetThemeDir() string {
	var _cret *C.gchar

	cret = C.gtk_rc_get_theme_dir()

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RCParse parses a given resource file.
func RCParse(filename string) {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse(_arg1)
}

// RCParseColor parses a color in the format expected in a RC file.
//
// Note that theme engines should use gtk_rc_parse_color_full() in order to
// support symbolic colors.
func RCParseColor(scanner *glib.Scanner) (gdk.Color, uint) {
	var _arg1 *C.GScanner

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner.Native()))

	var _color gdk.Color
	var _cret C.guint

	cret = C.gtk_rc_parse_color(_arg1, (*C.GdkColor)(unsafe.Pointer(&_color)))

	var _guint uint

	_guint = (uint)(_cret)

	return _color, _guint
}

// RCParseColorFull parses a color in the format expected in a RC file. If
// @style is not nil, it will be consulted to resolve references to symbolic
// colors.
func RCParseColorFull(scanner *glib.Scanner, style RCStyle) (gdk.Color, uint) {
	var _arg1 *C.GScanner
	var _arg2 *C.GtkRcStyle

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner.Native()))
	_arg2 = (*C.GtkRcStyle)(unsafe.Pointer(style.Native()))

	var _color gdk.Color
	var _cret C.guint

	cret = C.gtk_rc_parse_color_full(_arg1, _arg2, (*C.GdkColor)(unsafe.Pointer(&_color)))

	var _guint uint

	_guint = (uint)(_cret)

	return _color, _guint
}

// RCParsePriority parses a PathPriorityType variable from the format expected
// in a RC file.
func RCParsePriority(scanner *glib.Scanner, priority *PathPriorityType) uint {
	var _arg1 *C.GScanner
	var _arg2 *C.GtkPathPriorityType

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner.Native()))
	_arg2 = (*C.GtkPathPriorityType)(priority)

	var _cret C.guint

	cret = C.gtk_rc_parse_priority(_arg1, _arg2)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// RCParseState parses a StateType variable from the format expected in a RC
// file.
func RCParseState(scanner *glib.Scanner) (StateType, uint) {
	var _arg1 *C.GScanner

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner.Native()))

	var _arg2 C.GtkStateType
	var _cret C.guint

	cret = C.gtk_rc_parse_state(_arg1, &_arg2)

	var _state StateType
	var _guint uint

	_state = StateType(_arg2)
	_guint = (uint)(_cret)

	return _state, _guint
}

// RCParseString parses resource information directly from a string.
func RCParseString(rcString string) {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(rcString))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse_string(_arg1)
}

// RCReparseAll: if the modification time on any previously read file for the
// default Settings has changed, discard all style information and then reread
// all previously read RC files.
func RCReparseAll() bool {
	var _cret C.gboolean

	cret = C.gtk_rc_reparse_all()

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// RCReparseAllForSettings: if the modification time on any previously read file
// for the given Settings has changed, discard all style information and then
// reread all previously read RC files.
func RCReparseAllForSettings(settings Settings, forceLoad bool) bool {
	var _arg1 *C.GtkSettings
	var _arg2 C.gboolean

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	if forceLoad {
		_arg2 = C.gboolean(1)
	}

	var _cret C.gboolean

	cret = C.gtk_rc_reparse_all_for_settings(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// RCResetStyles: this function recomputes the styles for all widgets that use a
// particular Settings object. (There is one Settings object per Screen, see
// gtk_settings_get_for_screen()); It is useful when some global parameter has
// changed that affects the appearance of all widgets, because when a widget
// gets a new style, it will both redraw and recompute any cached information
// about its appearance. As an example, it is used when the default font size
// set by the operating system changes. Note that this function doesn’t affect
// widgets that have a style set explicitly on them with gtk_widget_set_style().
func RCResetStyles(settings Settings) {
	var _arg1 *C.GtkSettings

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))

	C.gtk_rc_reset_styles(_arg1)
}

func NewRCScanner() *glib.Scanner {
	var _cret *C.GScanner

	cret = C.gtk_rc_scanner_new()

	var _scanner *glib.Scanner

	_scanner = glib.WrapScanner(unsafe.Pointer(_cret))

	return _scanner
}

// RCSetDefaultFiles sets the list of files that GTK+ will read at the end of
// gtk_init().
func RCSetDefaultFiles(filenames []*string) {
	var _arg1 **C.gchar

	_arg1 = (**C.gchar)(C.malloc((len(filenames) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(_arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(_arg1), int(len(filenames)))

		for i := range filenames {
			_arg1 = (*C.gchar)(C.CString(filenames))
			defer C.free(unsafe.Pointer(_arg1))
		}
	}

	C.gtk_rc_set_default_files(_arg1)
}

// RCStyle: the RcStyle-struct is used to represent a set of information about
// the appearance of a widget. This can later be composited together with other
// RcStyle-struct<!-- -->s to form a Style.
type RCStyle interface {
	gextras.Objector

	// Copy makes a copy of the specified RcStyle. This function will correctly
	// copy an RC style that is a member of a class derived from RcStyle.
	Copy() RCStyle
}

// rcStyle implements the RCStyle interface.
type rcStyle struct {
	gextras.Objector
}

var _ RCStyle = (*rcStyle)(nil)

// WrapRCStyle wraps a GObject to the right type. It is
// primarily used internally.
func WrapRCStyle(obj *externglib.Object) RCStyle {
	return RCStyle{
		Objector: obj,
	}
}

func marshalRCStyle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRCStyle(obj), nil
}

// NewRCStyle constructs a class RCStyle.
func NewRCStyle() RCStyle {
	var _cret C.GtkRcStyle

	cret = C.gtk_rc_style_new()

	var _rcStyle RCStyle

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(RCStyle)

	return _rcStyle
}

// Copy makes a copy of the specified RcStyle. This function will correctly
// copy an RC style that is a member of a class derived from RcStyle.
func (o rcStyle) Copy() RCStyle {
	var _arg0 *C.GtkRcStyle

	_arg0 = (*C.GtkRcStyle)(unsafe.Pointer(o.Native()))

	var _cret *C.GtkRcStyle

	cret = C.gtk_rc_style_copy(_arg0)

	var _rcStyle RCStyle

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(RCStyle)

	return _rcStyle
}

type RCContext struct {
	native C.GtkRcContext
}

// WrapRCContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRCContext(ptr unsafe.Pointer) *RCContext {
	if ptr == nil {
		return nil
	}

	return (*RCContext)(ptr)
}

func marshalRCContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRCContext(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RCContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// RCProperty: deprecated
type RCProperty struct {
	native C.GtkRcProperty
}

// WrapRCProperty wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRCProperty(ptr unsafe.Pointer) *RCProperty {
	if ptr == nil {
		return nil
	}

	return (*RCProperty)(ptr)
}

func marshalRCProperty(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRCProperty(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RCProperty) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Origin gets the field inside the struct.
func (r *RCProperty) Origin() string {
	var v string
	v = C.GoString(r.native.origin)
	return v
}

// Value gets the field inside the struct.
func (r *RCProperty) Value() **externglib.Value {
	var v **externglib.Value
	v = externglib.ValueFromNative(unsafe.Pointer(r.native.value))
	return v
}