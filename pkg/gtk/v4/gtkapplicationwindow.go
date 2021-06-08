// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_application_window_get_type()), F: marshalApplicationWindow},
	})
}

// ApplicationWindow: `GtkApplicationWindow` is a `GtkWindow` subclass that
// integrates with `GtkApplication`.
//
// Notably, `GtkApplicationWindow` can handle an application menubar.
//
// This class implements the `GActionGroup` and `GActionMap` interfaces, to let
// you add window-specific actions that will be exported by the associated
// [class@Gtk.Application], together with its application-wide actions.
// Window-specific actions are prefixed with the “win.” prefix and
// application-wide actions are prefixed with the “app.” prefix. Actions must be
// addressed with the prefixed name when referring to them from a `GMenuModel`.
//
// Note that widgets that are placed inside a `GtkApplicationWindow` can also
// activate these actions, if they implement the [iface@Gtk.Actionable]
// interface.
//
// The settings [property@Gtk.Settings:gtk-shell-shows-app-menu] and
// [property@Gtk.Settings:gtk-shell-shows-menubar] tell GTK whether the desktop
// environment is showing the application menu and menubar models outside the
// application as part of the desktop shell. For instance, on OS X, both menus
// will be displayed remotely; on Windows neither will be.
//
// If the desktop environment does not display the menubar, then
// `GtkApplicationWindow` will automatically show a menubar for it. This
// behaviour can be overridden with the
// [property@Gtk.ApplicationWindow:show-menubar] property. If the desktop
// environment does not display the application menu, then it will automatically
// be included in the menubar or in the windows client-side decorations.
//
// See [class@Gtk.PopoverMenu] for information about the XML language used by
// `GtkBuilder` for menu models.
//
// See also: [method@Gtk.Application.set_menubar].
//
//
// A GtkApplicationWindow with a menubar
//
// The code sample below shows how to set up a `GtkApplicationWindow` with a
// menu bar defined on the [class@Gtk.Application]:
//
// “`c GtkApplication *app = gtk_application_new ("org.gtk.test", 0);
//
// GtkBuilder *builder = gtk_builder_new_from_string ( "<interface>" " <menu
// id='menubar'>" " <submenu>" " <attribute name='label'
// translatable='yes'>_Edit</attribute>" " <item>" " <attribute name='label'
// translatable='yes'>_Copy</attribute>" " <attribute
// name='action'>win.copy</attribute>" " </item>" " <item>" " <attribute
// name='label' translatable='yes'>_Paste</attribute>" " <attribute
// name='action'>win.paste</attribute>" " </item>" " </submenu>" " </menu>"
// "</interface>", -1);
//
// GMenuModel *menubar = G_MENU_MODEL (gtk_builder_get_object (builder,
// "menubar")); gtk_application_set_menubar (GTK_APPLICATION (app), menubar);
// g_object_unref (builder);
//
// // ...
//
// GtkWidget *window = gtk_application_window_new (app); “`
type ApplicationWindow interface {
	Window
	gio.ActionGroup
	gio.ActionMap
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager

	// HelpOverlay gets the `GtkShortcutsWindow` that is associated with
	// @window.
	//
	// See [method@Gtk.ApplicationWindow.set_help_overlay].
	HelpOverlay() ShortcutsWindow
	// ID returns the unique ID of the window.
	//
	//    If the window has not yet been added to a `GtkApplication`, returns `0`.
	ID() uint
	// ShowMenubar returns whether the window will display a menubar for the app
	// menu and menubar as needed.
	ShowMenubar() bool
	// SetHelpOverlay associates a shortcuts window with the application window.
	//
	// Additionally, sets up an action with the name `win.show-help-overlay` to
	// present it.
	//
	// @window takes responsibility for destroying @help_overlay.
	SetHelpOverlay(helpOverlay ShortcutsWindow)
	// SetShowMenubar sets whether the window will display a menubar for the app
	// menu and menubar as needed.
	SetShowMenubar(showMenubar bool)
}

// applicationWindow implements the ApplicationWindow interface.
type applicationWindow struct {
	Window
	gio.ActionGroup
	gio.ActionMap
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager
}

var _ ApplicationWindow = (*applicationWindow)(nil)

// WrapApplicationWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplicationWindow(obj *externglib.Object) ApplicationWindow {
	return ApplicationWindow{
		Window:           WrapWindow(obj),
		gio.ActionGroup:  gio.WrapActionGroup(obj),
		gio.ActionMap:    gio.WrapActionMap(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		Root:             WrapRoot(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplicationWindow(obj), nil
}

// NewApplicationWindow constructs a class ApplicationWindow.
func NewApplicationWindow(application Application) ApplicationWindow {
	var arg1 *C.GtkApplication

	arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	var cret C.GtkApplicationWindow
	var goret ApplicationWindow

	cret = C.gtk_application_window_new(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ApplicationWindow)

	return goret
}

// HelpOverlay gets the `GtkShortcutsWindow` that is associated with
// @window.
//
// See [method@Gtk.ApplicationWindow.set_help_overlay].
func (w applicationWindow) HelpOverlay() ShortcutsWindow {
	var arg0 *C.GtkApplicationWindow

	arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var cret *C.GtkShortcutsWindow
	var goret ShortcutsWindow

	cret = C.gtk_application_window_get_help_overlay(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ShortcutsWindow)

	return goret
}

// ID returns the unique ID of the window.
//
//    If the window has not yet been added to a `GtkApplication`, returns `0`.
func (w applicationWindow) ID() uint {
	var arg0 *C.GtkApplicationWindow

	arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_application_window_get_id(arg0)

	goret = uint(cret)

	return goret
}

// ShowMenubar returns whether the window will display a menubar for the app
// menu and menubar as needed.
func (w applicationWindow) ShowMenubar() bool {
	var arg0 *C.GtkApplicationWindow

	arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_application_window_get_show_menubar(arg0)

	if cret {
		goret = true
	}

	return goret
}

// SetHelpOverlay associates a shortcuts window with the application window.
//
// Additionally, sets up an action with the name `win.show-help-overlay` to
// present it.
//
// @window takes responsibility for destroying @help_overlay.
func (w applicationWindow) SetHelpOverlay(helpOverlay ShortcutsWindow) {
	var arg0 *C.GtkApplicationWindow
	var arg1 *C.GtkShortcutsWindow

	arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))
	arg1 = (*C.GtkShortcutsWindow)(unsafe.Pointer(helpOverlay.Native()))

	C.gtk_application_window_set_help_overlay(arg0, arg1)
}

// SetShowMenubar sets whether the window will display a menubar for the app
// menu and menubar as needed.
func (w applicationWindow) SetShowMenubar(showMenubar bool) {
	var arg0 *C.GtkApplicationWindow
	var arg1 C.gboolean

	arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))
	if showMenubar {
		arg1 = C.gboolean(1)
	}

	C.gtk_application_window_set_show_menubar(arg0, arg1)
}
