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

// ApplicationWindow is a Window subclass that offers some extra functionality
// for better integration with Application features. Notably, it can handle an
// application menubar. See gtk_application_set_menubar().
//
// This class implements the Group and Map interfaces, to let you add
// window-specific actions that will be exported by the associated Application,
// together with its application-wide actions. Window-specific actions are
// prefixed with the “win.” prefix and application-wide actions are prefixed
// with the “app.” prefix. Actions must be addressed with the prefixed name when
// referring to them from a Model.
//
// Note that widgets that are placed inside a ApplicationWindow can also
// activate these actions, if they implement the Actionable interface.
//
// As with Application, the GDK lock will be acquired when processing actions
// arriving from other processes and should therefore be held when activating
// actions locally (if GDK threads are enabled).
//
// The settings Settings:gtk-shell-shows-app-menu and
// Settings:gtk-shell-shows-menubar tell GTK+ whether the desktop environment is
// showing the application menu and menubar models outside the application as
// part of the desktop shell. For instance, on OS X, both menus will be
// displayed remotely; on Windows neither will be. gnome-shell (starting with
// version 3.4) will display the application menu, but not the menubar.
//
// If the desktop environment does not display the menubar, then
// ApplicationWindow will automatically show a menubar for it. This behaviour
// can be overridden with the ApplicationWindow:show-menubar property. If the
// desktop environment does not display the application menu, then it will
// automatically be included in the menubar or in the windows client-side
// decorations.
//
// See PopoverMenu for information about the XML language used by Builder for
// menu models.
//
// A GtkApplicationWindow with a menubar
//
//    GtkApplication *app = gtk_application_new ("org.gtk.test", 0);
//
//    GtkBuilder *builder = gtk_builder_new_from_string (
//        "<interface>"
//        "  <menu id='menubar'>"
//        "    <submenu>"
//        "      <attribute name='label' translatable='yes'>_Edit</attribute>"
//        "      <item>"
//        "        <attribute name='label' translatable='yes'>_Copy</attribute>"
//        "        <attribute name='action'>win.copy</attribute>"
//        "      </item>"
//        "      <item>"
//        "        <attribute name='label' translatable='yes'>_Paste</attribute>"
//        "        <attribute name='action'>win.paste</attribute>"
//        "      </item>"
//        "    </submenu>"
//        "  </menu>"
//        "</interface>",
//        -1);
//
//    GMenuModel *menubar = G_MENU_MODEL (gtk_builder_get_object (builder,
//                                                               "menubar"));
//    gtk_application_set_menubar (GTK_APPLICATION (app), menubar);
//    g_object_unref (builder);
//
//    // ...
//
//    GtkWidget *window = gtk_application_window_new (app);
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

	// HelpOverlay gets the ShortcutsWindow that has been set up with a prior
	// call to gtk_application_window_set_help_overlay().
	HelpOverlay() ShortcutsWindow
	// ID returns the unique ID of the window. If the window has not yet been
	// added to a Application, returns `0`.
	ID() uint
	// ShowMenubar returns whether the window will display a menubar for the app
	// menu and menubar as needed.
	ShowMenubar() bool
	// SetHelpOverlay associates a shortcuts window with the application window,
	// and sets up an action with the name win.show-help-overlay to present it.
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
	var _arg1 *C.GtkApplication

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	var _cret C.GtkApplicationWindow

	cret = C.gtk_application_window_new(_arg1)

	var _applicationWindow ApplicationWindow

	_applicationWindow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ApplicationWindow)

	return _applicationWindow
}

// HelpOverlay gets the ShortcutsWindow that has been set up with a prior
// call to gtk_application_window_set_help_overlay().
func (w applicationWindow) HelpOverlay() ShortcutsWindow {
	var _arg0 *C.GtkApplicationWindow

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var _cret *C.GtkShortcutsWindow

	cret = C.gtk_application_window_get_help_overlay(_arg0)

	var _shortcutsWindow ShortcutsWindow

	_shortcutsWindow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ShortcutsWindow)

	return _shortcutsWindow
}

// ID returns the unique ID of the window. If the window has not yet been
// added to a Application, returns `0`.
func (w applicationWindow) ID() uint {
	var _arg0 *C.GtkApplicationWindow

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var _cret C.guint

	cret = C.gtk_application_window_get_id(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// ShowMenubar returns whether the window will display a menubar for the app
// menu and menubar as needed.
func (w applicationWindow) ShowMenubar() bool {
	var _arg0 *C.GtkApplicationWindow

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))

	var _cret C.gboolean

	cret = C.gtk_application_window_get_show_menubar(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetHelpOverlay associates a shortcuts window with the application window,
// and sets up an action with the name win.show-help-overlay to present it.
//
// @window takes responsibility for destroying @help_overlay.
func (w applicationWindow) SetHelpOverlay(helpOverlay ShortcutsWindow) {
	var _arg0 *C.GtkApplicationWindow
	var _arg1 *C.GtkShortcutsWindow

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkShortcutsWindow)(unsafe.Pointer(helpOverlay.Native()))

	C.gtk_application_window_set_help_overlay(_arg0, _arg1)
}

// SetShowMenubar sets whether the window will display a menubar for the app
// menu and menubar as needed.
func (w applicationWindow) SetShowMenubar(showMenubar bool) {
	var _arg0 *C.GtkApplicationWindow
	var _arg1 C.gboolean

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(w.Native()))
	if showMenubar {
		_arg1 = C.gboolean(1)
	}

	C.gtk_application_window_set_show_menubar(_arg0, _arg1)
}
