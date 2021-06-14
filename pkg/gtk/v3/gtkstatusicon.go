// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_status_icon_get_type()), F: marshalStatusIcon},
	})
}

// StatusIcon: the “system tray” or notification area is normally used for
// transient icons that indicate some special state. For example, a system tray
// icon might appear to tell the user that they have new mail, or have an
// incoming instant message, or something along those lines. The basic idea is
// that creating an icon in the notification area is less annoying than popping
// up a dialog.
//
// A StatusIcon object can be used to display an icon in a “system tray”. The
// icon can have a tooltip, and the user can interact with it by activating it
// or popping up a context menu.
//
// It is very important to notice that status icons depend on the existence of a
// notification area being available to the user; you should not use status
// icons as the only way to convey critical information regarding your
// application, as the notification area may not exist on the user's
// environment, or may have been removed. You should always check that a status
// icon has been embedded into a notification area by using
// gtk_status_icon_is_embedded(), and gracefully recover if the function returns
// false.
//
// On X11, the implementation follows the FreeDesktop System Tray Specification
// (http://www.freedesktop.org/wiki/Specifications/systemtray-spec).
// Implementations of the “tray” side of this specification can be found e.g. in
// the GNOME 2 and KDE panel applications.
//
// Note that a GtkStatusIcon is not a widget, but just a #GObject. Making it a
// widget would be impractical, since the system tray on Windows doesn’t allow
// to embed arbitrary widgets.
//
// GtkStatusIcon has been deprecated in 3.14. You should consider using
// notifications or more modern platform-specific APIs instead. GLib provides
// the #GNotification API which works well with Application on multiple
// platforms and environments, and should be the preferred mechanism to notify
// the users of transient status updates. See this HowDoI
// (https://wiki.gnome.org/HowDoI/GNotification) for code examples.
type StatusIcon interface {
	gextras.Objector

	// HasTooltip returns the current value of the has-tooltip property. See
	// StatusIcon:has-tooltip for more information.
	HasTooltip() bool
	// IconName gets the name of the icon being displayed by the StatusIcon. The
	// storage type of the status icon must be GTK_IMAGE_EMPTY or
	// GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()). The
	// returned string is owned by the StatusIcon and should not be freed or
	// modified.
	IconName() string
	// Size gets the size in pixels that is available for the image. Stock icons
	// and named icons adapt their size automatically if the size of the
	// notification area changes. For other storage types, the size-changed
	// signal can be used to react to size changes.
	//
	// Note that the returned size is only meaningful while the status icon is
	// embedded (see gtk_status_icon_is_embedded()).
	Size() int
	// Stock gets the id of the stock icon being displayed by the StatusIcon.
	// The storage type of the status icon must be GTK_IMAGE_EMPTY or
	// GTK_IMAGE_STOCK (see gtk_status_icon_get_storage_type()). The returned
	// string is owned by the StatusIcon and should not be freed or modified.
	Stock() string
	// Title gets the title of this tray icon. See gtk_status_icon_set_title().
	Title() string
	// TooltipMarkup gets the contents of the tooltip for @status_icon.
	TooltipMarkup() string
	// TooltipText gets the contents of the tooltip for @status_icon.
	TooltipText() string
	// Visible returns whether the status icon is visible or not. Note that
	// being visible does not guarantee that the user can actually see the icon,
	// see also gtk_status_icon_is_embedded().
	Visible() bool
	// X11WindowID: this function is only useful on the X11/freedesktop.org
	// platform.
	//
	// It returns a window ID for the widget in the underlying status icon
	// implementation. This is useful for the Galago notification service, which
	// can send a window ID in the protocol in order for the server to position
	// notification windows pointing to a status icon reliably.
	//
	// This function is not intended for other use cases which are more likely
	// to be met by one of the non-X11 specific methods, such as
	// gtk_status_icon_position_menu().
	X11WindowID() uint32
	// IsEmbedded returns whether the status icon is embedded in a notification
	// area.
	IsEmbedded() bool
	// SetFromFile makes @status_icon display the file @filename. See
	// gtk_status_icon_new_from_file() for details.
	SetFromFile(filename string)
	// SetFromGIcon makes @status_icon display the #GIcon. See
	// gtk_status_icon_new_from_gicon() for details.
	SetFromGIcon(icon gio.Icon)
	// SetFromIconName makes @status_icon display the icon named @icon_name from
	// the current icon theme. See gtk_status_icon_new_from_icon_name() for
	// details.
	SetFromIconName(iconName string)
	// SetFromPixbuf makes @status_icon display @pixbuf. See
	// gtk_status_icon_new_from_pixbuf() for details.
	SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf)
	// SetFromStock makes @status_icon display the stock icon with the id
	// @stock_id. See gtk_status_icon_new_from_stock() for details.
	SetFromStock(stockId string)
	// SetHasTooltip sets the has-tooltip property on @status_icon to
	// @has_tooltip. See StatusIcon:has-tooltip for more information.
	SetHasTooltip(hasTooltip bool)
	// SetName sets the name of this tray icon. This should be a string
	// identifying this icon. It is may be used for sorting the icons in the
	// tray and will not be shown to the user.
	SetName(name string)
	// SetScreen sets the Screen where @status_icon is displayed; if the icon is
	// already mapped, it will be unmapped, and then remapped on the new screen.
	SetScreen(screen gdk.Screen)
	// SetTitle sets the title of this tray icon. This should be a short,
	// human-readable, localized string describing the tray icon. It may be used
	// by tools like screen readers to render the tray icon.
	SetTitle(title string)
	// SetTooltipMarkup sets @markup as the contents of the tooltip, which is
	// marked up with the [Pango text markup language][PangoMarkupFormat].
	//
	// This function will take care of setting StatusIcon:has-tooltip to true
	// and of the default handler for the StatusIcon::query-tooltip signal.
	//
	// See also the StatusIcon:tooltip-markup property and
	// gtk_tooltip_set_markup().
	SetTooltipMarkup(markup string)
	// SetTooltipText sets @text as the contents of the tooltip.
	//
	// This function will take care of setting StatusIcon:has-tooltip to true
	// and of the default handler for the StatusIcon::query-tooltip signal.
	//
	// See also the StatusIcon:tooltip-text property and gtk_tooltip_set_text().
	SetTooltipText(text string)
	// SetVisible shows or hides a status icon.
	SetVisible(visible bool)
}

// statusIcon implements the StatusIcon class.
type statusIcon struct {
	gextras.Objector
}

var _ StatusIcon = (*statusIcon)(nil)

// WrapStatusIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapStatusIcon(obj *externglib.Object) StatusIcon {
	return statusIcon{
		Objector: obj,
	}
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusIcon(obj), nil
}

// HasTooltip returns the current value of the has-tooltip property. See
// StatusIcon:has-tooltip for more information.
func (s statusIcon) HasTooltip() bool {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_status_icon_get_has_tooltip(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the icon being displayed by the StatusIcon. The
// storage type of the status icon must be GTK_IMAGE_EMPTY or
// GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()). The
// returned string is owned by the StatusIcon and should not be freed or
// modified.
func (s statusIcon) IconName() string {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_status_icon_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Size gets the size in pixels that is available for the image. Stock icons
// and named icons adapt their size automatically if the size of the
// notification area changes. For other storage types, the size-changed
// signal can be used to react to size changes.
//
// Note that the returned size is only meaningful while the status icon is
// embedded (see gtk_status_icon_is_embedded()).
func (s statusIcon) Size() int {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret C.gint // in

	_cret = C.gtk_status_icon_get_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Stock gets the id of the stock icon being displayed by the StatusIcon.
// The storage type of the status icon must be GTK_IMAGE_EMPTY or
// GTK_IMAGE_STOCK (see gtk_status_icon_get_storage_type()). The returned
// string is owned by the StatusIcon and should not be freed or modified.
func (s statusIcon) Stock() string {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_status_icon_get_stock(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Title gets the title of this tray icon. See gtk_status_icon_set_title().
func (s statusIcon) Title() string {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_status_icon_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// TooltipMarkup gets the contents of the tooltip for @status_icon.
func (s statusIcon) TooltipMarkup() string {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_status_icon_get_tooltip_markup(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// TooltipText gets the contents of the tooltip for @status_icon.
func (s statusIcon) TooltipText() string {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_status_icon_get_tooltip_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Visible returns whether the status icon is visible or not. Note that
// being visible does not guarantee that the user can actually see the icon,
// see also gtk_status_icon_is_embedded().
func (s statusIcon) Visible() bool {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_status_icon_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// X11WindowID: this function is only useful on the X11/freedesktop.org
// platform.
//
// It returns a window ID for the widget in the underlying status icon
// implementation. This is useful for the Galago notification service, which
// can send a window ID in the protocol in order for the server to position
// notification windows pointing to a status icon reliably.
//
// This function is not intended for other use cases which are more likely
// to be met by one of the non-X11 specific methods, such as
// gtk_status_icon_position_menu().
func (s statusIcon) X11WindowID() uint32 {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret C.guint32 // in

	_cret = C.gtk_status_icon_get_x11_window_id(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// IsEmbedded returns whether the status icon is embedded in a notification
// area.
func (s statusIcon) IsEmbedded() bool {
	var _arg0 *C.GtkStatusIcon // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_status_icon_is_embedded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFromFile makes @status_icon display the file @filename. See
// gtk_status_icon_new_from_file() for details.
func (s statusIcon) SetFromFile(filename string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_file(_arg0, _arg1)
}

// SetFromGIcon makes @status_icon display the #GIcon. See
// gtk_status_icon_new_from_gicon() for details.
func (s statusIcon) SetFromGIcon(icon gio.Icon) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GIcon         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_status_icon_set_from_gicon(_arg0, _arg1)
}

// SetFromIconName makes @status_icon display the icon named @icon_name from
// the current icon theme. See gtk_status_icon_new_from_icon_name() for
// details.
func (s statusIcon) SetFromIconName(iconName string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_icon_name(_arg0, _arg1)
}

// SetFromPixbuf makes @status_icon display @pixbuf. See
// gtk_status_icon_new_from_pixbuf() for details.
func (s statusIcon) SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GdkPixbuf     // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_status_icon_set_from_pixbuf(_arg0, _arg1)
}

// SetFromStock makes @status_icon display the stock icon with the id
// @stock_id. See gtk_status_icon_new_from_stock() for details.
func (s statusIcon) SetFromStock(stockId string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_stock(_arg0, _arg1)
}

// SetHasTooltip sets the has-tooltip property on @status_icon to
// @has_tooltip. See StatusIcon:has-tooltip for more information.
func (s statusIcon) SetHasTooltip(hasTooltip bool) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	if hasTooltip {
		_arg1 = C.TRUE
	}

	C.gtk_status_icon_set_has_tooltip(_arg0, _arg1)
}

// SetName sets the name of this tray icon. This should be a string
// identifying this icon. It is may be used for sorting the icons in the
// tray and will not be shown to the user.
func (s statusIcon) SetName(name string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_name(_arg0, _arg1)
}

// SetScreen sets the Screen where @status_icon is displayed; if the icon is
// already mapped, it will be unmapped, and then remapped on the new screen.
func (s statusIcon) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GdkScreen     // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_status_icon_set_screen(_arg0, _arg1)
}

// SetTitle sets the title of this tray icon. This should be a short,
// human-readable, localized string describing the tray icon. It may be used
// by tools like screen readers to render the tray icon.
func (s statusIcon) SetTitle(title string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_title(_arg0, _arg1)
}

// SetTooltipMarkup sets @markup as the contents of the tooltip, which is
// marked up with the [Pango text markup language][PangoMarkupFormat].
//
// This function will take care of setting StatusIcon:has-tooltip to true
// and of the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-markup property and
// gtk_tooltip_set_markup().
func (s statusIcon) SetTooltipMarkup(markup string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_tooltip_markup(_arg0, _arg1)
}

// SetTooltipText sets @text as the contents of the tooltip.
//
// This function will take care of setting StatusIcon:has-tooltip to true
// and of the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-text property and gtk_tooltip_set_text().
func (s statusIcon) SetTooltipText(text string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_tooltip_text(_arg0, _arg1)
}

// SetVisible shows or hides a status icon.
func (s statusIcon) SetVisible(visible bool) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(s.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_status_icon_set_visible(_arg0, _arg1)
}
