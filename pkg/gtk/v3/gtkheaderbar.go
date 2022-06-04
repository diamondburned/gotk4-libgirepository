// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkheaderbar.go.
var GTypeHeaderBar = coreglib.Type(C.gtk_header_bar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeHeaderBar, F: marshalHeaderBar},
	})
}

// HeaderBarOverrider contains methods that are overridable.
type HeaderBarOverrider interface {
}

// HeaderBar is similar to a horizontal Box. It allows children to be placed at
// the start or the end. In addition, it allows a title and subtitle to be
// displayed. The title will be centered with respect to the width of the box,
// even if the children at either side take up different amounts of space. The
// height of the titlebar will be set to provide sufficient space for the
// subtitle, even if none is currently set. If a subtitle is not needed, the
// space reservation can be turned off with gtk_header_bar_set_has_subtitle().
//
// GtkHeaderBar can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, GtkHeaderBar is the natural choice for use as the custom
// titlebar widget of a Window (see gtk_window_set_titlebar()), as it gives
// features typical of titlebars while allowing the addition of child widgets.
type HeaderBar struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*HeaderBar)(nil)
)

func classInitHeaderBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHeaderBar(obj *coreglib.Object) *HeaderBar {
	return &HeaderBar{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	return wrapHeaderBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHeaderBar creates a new HeaderBar widget.
//
// The function returns the following values:
//
//    - headerBar: new HeaderBar.
//
func NewHeaderBar() *HeaderBar {
	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("new_HeaderBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// CustomTitle retrieves the custom title widget of the header. See
// gtk_header_bar_set_custom_title().
//
// The function returns the following values:
//
//    - widget (optional): custom title widget of the header, or NULL if none has
//      been set explicitly.
//
func (bar *HeaderBar) CustomTitle() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_custom_title", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// DecorationLayout gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
//
// The function returns the following values:
//
//    - utf8: decoration layout.
//
func (bar *HeaderBar) DecorationLayout() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_decoration_layout", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasSubtitle retrieves whether the header bar reserves space for a subtitle,
// regardless if one is currently set or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the header bar reserves space for a subtitle.
//
func (bar *HeaderBar) HasSubtitle() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_has_subtitle", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether this header bar shows the standard window
// decorations.
//
// The function returns the following values:
//
//    - ok: TRUE if the decorations are shown.
//
func (bar *HeaderBar) ShowCloseButton() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_show_close_button", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Subtitle retrieves the subtitle of the header. See
// gtk_header_bar_set_subtitle().
//
// The function returns the following values:
//
//    - utf8 (optional): subtitle of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (bar *HeaderBar) Subtitle() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_subtitle", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title retrieves the title of the header. See gtk_header_bar_set_title().
//
// The function returns the following values:
//
//    - utf8 (optional): title of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (bar *HeaderBar) Title() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("get_title", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PackEnd adds child to bar, packed with reference to the end of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackEnd(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("pack_end", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// PackStart adds child to bar, packed with reference to the start of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackStart(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("pack_start", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes any
// title set by gtk_header_bar_set_title() or gtk_header_bar_set_subtitle(). To
// achieve the same style as the builtin title and subtitle, use the “title” and
// “subtitle” style classes.
//
// You should set the custom title to NULL, for the header title label to be
// visible again.
//
// The function takes the following parameters:
//
//    - titleWidget (optional): custom widget to use for a title.
//
func (bar *HeaderBar) SetCustomTitle(titleWidget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if titleWidget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(titleWidget).Native()))
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_custom_title", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(titleWidget)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close, icon (the window icon)
// and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left, and
// minimize, maximize and close buttons on the right.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (bar *HeaderBar) SetDecorationLayout(layout string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if layout != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_decoration_layout", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(layout)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
//
// The function takes the following parameters:
//
//    - setting: TRUE to reserve space for a subtitle.
//
func (bar *HeaderBar) SetHasSubtitle(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_has_subtitle", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
//
// The function takes the following parameters:
//
//    - setting: TRUE to show standard window decorations.
//
func (bar *HeaderBar) SetShowCloseButton(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_show_close_button", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a user
// an additional detail to help him identify the current view.
//
// Note that GtkHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the HeaderBar:has-subtitle
// property to FALSE.
//
// The function takes the following parameters:
//
//    - subtitle (optional): subtitle, or NULL.
//
func (bar *HeaderBar) SetSubtitle(subtitle string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if subtitle != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(subtitle)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_subtitle", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the application
// name.
//
// The function takes the following parameters:
//
//    - title (optional): title, or NULL.
//
func (bar *HeaderBar) SetTitle(title string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if title != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "HeaderBar").InvokeMethod("set_title", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(title)
}
