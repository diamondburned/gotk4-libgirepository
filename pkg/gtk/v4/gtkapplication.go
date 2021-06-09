// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_application_get_type()), F: marshalApplication},
	})
}

// Application is a class that handles many important aspects of a GTK+
// application in a convenient fashion, without enforcing a one-size-fits-all
// application model.
//
// Currently, GtkApplication handles GTK initialization, application uniqueness,
// session management, provides some basic scriptability and desktop shell
// integration by exporting actions and menus and manages a list of toplevel
// windows whose life-cycle is automatically tied to the life-cycle of your
// application.
//
// While GtkApplication works fine with plain Windows, it is recommended to use
// it together with ApplicationWindow.
//
// When GDK threads are enabled, GtkApplication will acquire the GDK lock when
// invoking actions that arrive from other processes. The GDK lock is not
// touched for local action invocations. In order to have actions invoked in a
// predictable context it is therefore recommended that the GDK lock be held
// while invoking actions locally with g_action_group_activate_action(). The
// same applies to actions associated with ApplicationWindow and to the
// “activate” and “open” #GApplication methods.
//
//
// Automatic resources
//
// Application will automatically load menus from the Builder resource located
// at "gtk/menus.ui", relative to the application's resource base path (see
// g_application_set_resource_base_path()). The menu with the ID "menubar" is
// taken as the application's menubar. Additional menus (most interesting
// submenus) can be named and accessed via gtk_application_get_menu_by_id()
// which allows for dynamic population of a part of the menu structure.
//
// It is also possible to provide the menubar manually using
// gtk_application_set_menubar().
//
// Application will also automatically setup an icon search path for the default
// icon theme by appending "icons" to the resource base path. This allows your
// application to easily store its icons as resources. See
// gtk_icon_theme_add_resource_path() for more information.
//
// If there is a resource located at "gtk/help-overlay.ui" which defines a
// ShortcutsWindow with ID "help_overlay" then GtkApplication associates an
// instance of this shortcuts window with each ApplicationWindow and sets up the
// keyboard accelerator Control-? to open it. To create a menu item that
// displays the shortcuts window, associate the item with the action
// win.show-help-overlay.
//
//
// A simple application
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/bp/bloatpad.c)
//
// GtkApplication optionally registers with a session manager of the users
// session (if you set the Application:register-session property) and offers
// various functionality related to the session life-cycle.
//
// An application can block various ways to end the session with the
// gtk_application_inhibit() function. Typical use cases for this kind of
// inhibiting are long-running, uninterruptible operations, such as burning a CD
// or performing a disk backup. The session manager may not honor the inhibitor,
// but it can be expected to inform the user about the negative consequences of
// ending the session while inhibitors are present.
//
//
// See Also
//
// HowDoI: Using GtkApplication (https://wiki.gnome.org/HowDoI/GtkApplication),
// Getting Started with GTK: Basics
// (https://developer.gnome.org/gtk3/stable/gtk-getting-started.html#id-1.2.3.3)
type Application interface {
	gio.Application
	gio.ActionGroup
	gio.ActionMap

	// AddWindow adds a window to @application.
	//
	// This call can only happen after the @application has started; typically,
	// you should add new application windows in response to the emission of the
	// #GApplication::activate signal.
	//
	// This call is equivalent to setting the Window:application property of
	// @window to @application.
	//
	// Normally, the connection between the application and the window will
	// remain until the window is destroyed, but you can explicitly remove it
	// with gtk_application_remove_window().
	//
	// GTK+ will keep the @application running as long as it has any windows.
	AddWindow(window Window)
	// AccelsForAction gets the accelerators that are currently associated with
	// the given action.
	AccelsForAction(detailedActionName string) []string
	// ActionsForAccel returns the list of actions (possibly empty) that @accel
	// maps to. Each item in the list is a detailed action name in the usual
	// form.
	//
	// This might be useful to discover if an accel already exists in order to
	// prevent installation of a conflicting accelerator (from an accelerator
	// editor or a plugin system, for example). Note that having more than one
	// action per accelerator may not be a bad thing and might make sense in
	// cases where the actions never appear in the same context.
	//
	// In case there are no actions for a given accelerator, an empty array is
	// returned. nil is never returned.
	//
	// It is a programmer error to pass an invalid accelerator string. If you
	// are unsure, check it with gtk_accelerator_parse() first.
	ActionsForAccel(accel string) []string
	// ActiveWindow gets the “active” window for the application.
	//
	// The active window is the one that was most recently focused (within the
	// application). This window may not have the focus at the moment if another
	// application has it — this is just the most recently-focused window within
	// this application.
	ActiveWindow() Window
	// MenuByID gets a menu from automatically loaded resources. See [Automatic
	// resources][automatic-resources] for more information.
	MenuByID(id string) gio.Menu
	// Menubar returns the menu model that has been set with
	// gtk_application_set_menubar().
	Menubar() gio.MenuModel
	// WindowByID returns the ApplicationWindow with the given ID.
	//
	// The ID of a ApplicationWindow can be retrieved with
	// gtk_application_window_get_id().
	WindowByID(id uint) Window
	// Windows gets a list of the Windows associated with @application.
	//
	// The list is sorted by most recently focused window, such that the first
	// element is the currently focused window. (Useful for choosing a parent
	// for a transient window.)
	//
	// The list that is returned should not be modified in any way. It will only
	// remain valid until the next focus change or window creation or deletion.
	Windows() *glib.List
	// Inhibit: inform the session manager that certain types of actions should
	// be inhibited. This is not guaranteed to work on all platforms and for all
	// types of actions.
	//
	// Applications should invoke this method when they begin an operation that
	// should not be interrupted, such as creating a CD or DVD. The types of
	// actions that may be blocked are specified by the @flags parameter. When
	// the application completes the operation it should call
	// gtk_application_uninhibit() to remove the inhibitor. Note that an
	// application can have multiple inhibitors, and all of them must be
	// individually removed. Inhibitors are also cleared when the application
	// exits.
	//
	// Applications should not expect that they will always be able to block the
	// action. In most cases, users will be given the option to force the action
	// to take place.
	//
	// Reasons should be short and to the point.
	//
	// If @window is given, the session manager may point the user to this
	// window to find out more about why the action is inhibited.
	Inhibit(window Window, flags ApplicationInhibitFlags, reason string) uint
	// ListActionDescriptions lists the detailed action names which have
	// associated accelerators. See gtk_application_set_accels_for_action().
	ListActionDescriptions() []string
	// RemoveWindow: remove a window from @application.
	//
	// If @window belongs to @application then this call is equivalent to
	// setting the Window:application property of @window to nil.
	//
	// The application may stop running as a result of a call to this function.
	RemoveWindow(window Window)
	// SetAccelsForAction sets zero or more keyboard accelerators that will
	// trigger the given action. The first item in @accels will be the primary
	// accelerator, which may be displayed in the UI.
	//
	// To remove all accelerators for an action, use an empty, zero-terminated
	// array for @accels.
	//
	// For the @detailed_action_name, see g_action_parse_detailed_name() and
	// g_action_print_detailed_name().
	SetAccelsForAction(detailedActionName string, accels []string)
	// SetMenubar sets or unsets the menubar for windows of @application.
	//
	// This is a menubar in the traditional sense.
	//
	// This can only be done in the primary instance of the application, after
	// it has been registered. #GApplication::startup is a good place to call
	// this.
	//
	// Depending on the desktop environment, this may appear at the top of each
	// window, or at the top of the screen. In some environments, if both the
	// application menu and the menubar are set, the application menu will be
	// presented as if it were the first item of the menubar. Other environments
	// treat the two as completely separate — for example, the application menu
	// may be rendered by the desktop shell while the menubar (if set) remains
	// in each individual window.
	//
	// Use the base Map interface to add actions, to respond to the user
	// selecting these menu items.
	SetMenubar(menubar gio.MenuModel)
	// Uninhibit removes an inhibitor that has been established with
	// gtk_application_inhibit(). Inhibitors are also cleared when the
	// application exits.
	Uninhibit(cookie uint)
}

// application implements the Application interface.
type application struct {
	gio.Application
	gio.ActionGroup
	gio.ActionMap
}

var _ Application = (*application)(nil)

// WrapApplication wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplication(obj *externglib.Object) Application {
	return Application{
		gio.Application: gio.WrapApplication(obj),
		gio.ActionGroup: gio.WrapActionGroup(obj),
		gio.ActionMap:   gio.WrapActionMap(obj),
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplication(obj), nil
}

// NewApplication constructs a class Application.
func NewApplication(applicationId string, flags gio.ApplicationFlags) Application {
	var _arg1 *C.char
	var _arg2 C.GApplicationFlags

	_arg1 = (*C.char)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GApplicationFlags)(flags)

	var _cret C.GtkApplication

	cret = C.gtk_application_new(_arg1, _arg2)

	var _application Application

	_application = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Application)

	return _application
}

// AddWindow adds a window to @application.
//
// This call can only happen after the @application has started; typically,
// you should add new application windows in response to the emission of the
// #GApplication::activate signal.
//
// This call is equivalent to setting the Window:application property of
// @window to @application.
//
// Normally, the connection between the application and the window will
// remain until the window is destroyed, but you can explicitly remove it
// with gtk_application_remove_window().
//
// GTK+ will keep the @application running as long as it has any windows.
func (a application) AddWindow(window Window) {
	var _arg0 *C.GtkApplication
	var _arg1 *C.GtkWindow

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_add_window(_arg0, _arg1)
}

// AccelsForAction gets the accelerators that are currently associated with
// the given action.
func (a application) AccelsForAction(detailedActionName string) []string {
	var _arg0 *C.GtkApplication
	var _arg1 *C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.char

	cret = C.gtk_application_get_accels_for_action(_arg0, _arg1)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// ActionsForAccel returns the list of actions (possibly empty) that @accel
// maps to. Each item in the list is a detailed action name in the usual
// form.
//
// This might be useful to discover if an accel already exists in order to
// prevent installation of a conflicting accelerator (from an accelerator
// editor or a plugin system, for example). Note that having more than one
// action per accelerator may not be a bad thing and might make sense in
// cases where the actions never appear in the same context.
//
// In case there are no actions for a given accelerator, an empty array is
// returned. nil is never returned.
//
// It is a programmer error to pass an invalid accelerator string. If you
// are unsure, check it with gtk_accelerator_parse() first.
func (a application) ActionsForAccel(accel string) []string {
	var _arg0 *C.GtkApplication
	var _arg1 *C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(accel))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.char

	cret = C.gtk_application_get_actions_for_accel(_arg0, _arg1)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// ActiveWindow gets the “active” window for the application.
//
// The active window is the one that was most recently focused (within the
// application). This window may not have the focus at the moment if another
// application has it — this is just the most recently-focused window within
// this application.
func (a application) ActiveWindow() Window {
	var _arg0 *C.GtkApplication

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	var _cret *C.GtkWindow

	cret = C.gtk_application_get_active_window(_arg0)

	var _window Window

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Window)

	return _window
}

// MenuByID gets a menu from automatically loaded resources. See [Automatic
// resources][automatic-resources] for more information.
func (a application) MenuByID(id string) gio.Menu {
	var _arg0 *C.GtkApplication
	var _arg1 *C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GMenu

	cret = C.gtk_application_get_menu_by_id(_arg0, _arg1)

	var _menu gio.Menu

	_menu = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.Menu)

	return _menu
}

// Menubar returns the menu model that has been set with
// gtk_application_set_menubar().
func (a application) Menubar() gio.MenuModel {
	var _arg0 *C.GtkApplication

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	var _cret *C.GMenuModel

	cret = C.gtk_application_get_menubar(_arg0)

	var _menuModel gio.MenuModel

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.MenuModel)

	return _menuModel
}

// WindowByID returns the ApplicationWindow with the given ID.
//
// The ID of a ApplicationWindow can be retrieved with
// gtk_application_window_get_id().
func (a application) WindowByID(id uint) Window {
	var _arg0 *C.GtkApplication
	var _arg1 C.guint

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(id)

	var _cret *C.GtkWindow

	cret = C.gtk_application_get_window_by_id(_arg0, _arg1)

	var _window Window

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Window)

	return _window
}

// Windows gets a list of the Windows associated with @application.
//
// The list is sorted by most recently focused window, such that the first
// element is the currently focused window. (Useful for choosing a parent
// for a transient window.)
//
// The list that is returned should not be modified in any way. It will only
// remain valid until the next focus change or window creation or deletion.
func (a application) Windows() *glib.List {
	var _arg0 *C.GtkApplication

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	var _cret *C.GList

	cret = C.gtk_application_get_windows(_arg0)

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))

	return _list
}

// Inhibit: inform the session manager that certain types of actions should
// be inhibited. This is not guaranteed to work on all platforms and for all
// types of actions.
//
// Applications should invoke this method when they begin an operation that
// should not be interrupted, such as creating a CD or DVD. The types of
// actions that may be blocked are specified by the @flags parameter. When
// the application completes the operation it should call
// gtk_application_uninhibit() to remove the inhibitor. Note that an
// application can have multiple inhibitors, and all of them must be
// individually removed. Inhibitors are also cleared when the application
// exits.
//
// Applications should not expect that they will always be able to block the
// action. In most cases, users will be given the option to force the action
// to take place.
//
// Reasons should be short and to the point.
//
// If @window is given, the session manager may point the user to this
// window to find out more about why the action is inhibited.
func (a application) Inhibit(window Window, flags ApplicationInhibitFlags, reason string) uint {
	var _arg0 *C.GtkApplication
	var _arg1 *C.GtkWindow
	var _arg2 C.GtkApplicationInhibitFlags
	var _arg3 *C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (C.GtkApplicationInhibitFlags)(flags)
	_arg3 = (*C.char)(C.CString(reason))
	defer C.free(unsafe.Pointer(_arg3))

	var _cret C.guint

	cret = C.gtk_application_inhibit(_arg0, _arg1, _arg2, _arg3)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// ListActionDescriptions lists the detailed action names which have
// associated accelerators. See gtk_application_set_accels_for_action().
func (a application) ListActionDescriptions() []string {
	var _arg0 *C.GtkApplication

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	var _cret **C.char

	cret = C.gtk_application_list_action_descriptions(_arg0)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// RemoveWindow: remove a window from @application.
//
// If @window belongs to @application then this call is equivalent to
// setting the Window:application property of @window to nil.
//
// The application may stop running as a result of a call to this function.
func (a application) RemoveWindow(window Window) {
	var _arg0 *C.GtkApplication
	var _arg1 *C.GtkWindow

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_remove_window(_arg0, _arg1)
}

// SetAccelsForAction sets zero or more keyboard accelerators that will
// trigger the given action. The first item in @accels will be the primary
// accelerator, which may be displayed in the UI.
//
// To remove all accelerators for an action, use an empty, zero-terminated
// array for @accels.
//
// For the @detailed_action_name, see g_action_parse_detailed_name() and
// g_action_print_detailed_name().
func (a application) SetAccelsForAction(detailedActionName string, accels []string) {
	var _arg0 *C.GtkApplication
	var _arg1 *C.char
	var _arg2 **C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc((len(accels) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(_arg2))

	{
		var out []*C.char
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(_arg2), int(len(accels)))

		for i := range accels {
			_arg2 = (*C.char)(C.CString(accels))
			defer C.free(unsafe.Pointer(_arg2))
		}
	}

	C.gtk_application_set_accels_for_action(_arg0, _arg1, _arg2)
}

// SetMenubar sets or unsets the menubar for windows of @application.
//
// This is a menubar in the traditional sense.
//
// This can only be done in the primary instance of the application, after
// it has been registered. #GApplication::startup is a good place to call
// this.
//
// Depending on the desktop environment, this may appear at the top of each
// window, or at the top of the screen. In some environments, if both the
// application menu and the menubar are set, the application menu will be
// presented as if it were the first item of the menubar. Other environments
// treat the two as completely separate — for example, the application menu
// may be rendered by the desktop shell while the menubar (if set) remains
// in each individual window.
//
// Use the base Map interface to add actions, to respond to the user
// selecting these menu items.
func (a application) SetMenubar(menubar gio.MenuModel) {
	var _arg0 *C.GtkApplication
	var _arg1 *C.GMenuModel

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menubar.Native()))

	C.gtk_application_set_menubar(_arg0, _arg1)
}

// Uninhibit removes an inhibitor that has been established with
// gtk_application_inhibit(). Inhibitors are also cleared when the
// application exits.
func (a application) Uninhibit(cookie uint) {
	var _arg0 *C.GtkApplication
	var _arg1 C.guint

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(cookie)

	C.gtk_application_uninhibit(_arg0, _arg1)
}
