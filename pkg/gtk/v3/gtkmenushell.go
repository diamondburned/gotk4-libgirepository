// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_MenuShellClass_move_selected(void*, gint);
// extern gboolean _gotk4_gtk3_MenuShell_ConnectMoveSelected(gpointer, gint, guintptr);
// extern gint _gotk4_gtk3_MenuShellClass_get_popup_delay(void*);
// extern void _gotk4_gtk3_MenuShellClass_activate_current(void*, gboolean);
// extern void _gotk4_gtk3_MenuShellClass_cancel(void*);
// extern void _gotk4_gtk3_MenuShellClass_deactivate(void*);
// extern void _gotk4_gtk3_MenuShellClass_insert(void*, void*, gint);
// extern void _gotk4_gtk3_MenuShellClass_select_item(void*, void*);
// extern void _gotk4_gtk3_MenuShellClass_selection_done(void*);
// extern void _gotk4_gtk3_MenuShell_ConnectActivateCurrent(gpointer, gboolean, guintptr);
// extern void _gotk4_gtk3_MenuShell_ConnectCancel(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuShell_ConnectDeactivate(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuShell_ConnectInsert(gpointer, void*, gint, guintptr);
// extern void _gotk4_gtk3_MenuShell_ConnectSelectionDone(gpointer, guintptr);
import "C"

// glib.Type values for gtkmenushell.go.
var GTypeMenuShell = coreglib.Type(C.gtk_menu_shell_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeMenuShell, F: marshalMenuShell},
	})
}

// MenuShellOverrider contains methods that are overridable.
type MenuShellOverrider interface {
	// The function takes the following parameters:
	//
	ActivateCurrent(forceHide bool)
	// Cancel cancels the selection within the menu shell.
	Cancel()
	// Deactivate deactivates the menu shell.
	//
	// Typically this results in the menu shell being erased from the screen.
	Deactivate()
	// The function returns the following values:
	//
	PopupDelay() int32
	// Insert adds a new MenuItem to the menu shell’s item list at the position
	// indicated by position.
	//
	// The function takes the following parameters:
	//
	//    - child to add.
	//    - position in the item list where child is added. Positions are
	//      numbered from 0 to n-1.
	//
	Insert(child Widgetter, position int32)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	MoveSelected(distance int32) bool
	// SelectItem selects the menu item from the menu shell.
	//
	// The function takes the following parameters:
	//
	//    - menuItem to select.
	//
	SelectItem(menuItem Widgetter)
	SelectionDone()
}

// MenuShell is the abstract base class used to derive the Menu and MenuBar
// subclasses.
//
// A MenuShell is a container of MenuItem objects arranged in a list which can
// be navigated, selected, and activated by the user to perform application
// functions. A MenuItem can have a submenu associated with it, allowing for
// nested hierarchical menus.
//
//
// Terminology
//
// A menu item can be “selected”, this means that it is displayed in the
// prelight state, and if it has a submenu, that submenu will be popped up.
//
// A menu is “active” when it is visible onscreen and the user is selecting from
// it. A menubar is not active until the user clicks on one of its menuitems.
// When a menu is active, passing the mouse over a submenu will pop it up.
//
// There is also is a concept of the current menu and a current menu item. The
// current menu item is the selected menu item that is furthest down in the
// hierarchy. (Every active menu shell does not necessarily contain a selected
// menu item, but if it does, then the parent menu shell must also contain a
// selected menu item.) The current menu is the menu that contains the current
// menu item. It will always have a GTK grab and receive all key presses.
type MenuShell struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*MenuShell)(nil)
)

// MenuSheller describes types inherited from class MenuShell.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MenuSheller interface {
	coreglib.Objector
	baseMenuShell() *MenuShell
}

var _ MenuSheller = (*MenuShell)(nil)

func classInitMenuSheller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkMenuShellClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkMenuShellClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ActivateCurrent(forceHide bool) }); ok {
		pclass.activate_current = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_activate_current)
	}

	if _, ok := goval.(interface{ Cancel() }); ok {
		pclass.cancel = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_cancel)
	}

	if _, ok := goval.(interface{ Deactivate() }); ok {
		pclass.deactivate = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_deactivate)
	}

	if _, ok := goval.(interface{ PopupDelay() int32 }); ok {
		pclass.get_popup_delay = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_get_popup_delay)
	}

	if _, ok := goval.(interface {
		Insert(child Widgetter, position int32)
	}); ok {
		pclass.insert = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_insert)
	}

	if _, ok := goval.(interface{ MoveSelected(distance int32) bool }); ok {
		pclass.move_selected = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_move_selected)
	}

	if _, ok := goval.(interface{ SelectItem(menuItem Widgetter) }); ok {
		pclass.select_item = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_select_item)
	}

	if _, ok := goval.(interface{ SelectionDone() }); ok {
		pclass.selection_done = (*[0]byte)(C._gotk4_gtk3_MenuShellClass_selection_done)
	}
}

//export _gotk4_gtk3_MenuShellClass_activate_current
func _gotk4_gtk3_MenuShellClass_activate_current(arg0 *C.void, arg1 C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateCurrent(forceHide bool) })

	var _forceHide bool // out

	if arg1 != 0 {
		_forceHide = true
	}

	iface.ActivateCurrent(_forceHide)
}

//export _gotk4_gtk3_MenuShellClass_cancel
func _gotk4_gtk3_MenuShellClass_cancel(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancel() })

	iface.Cancel()
}

//export _gotk4_gtk3_MenuShellClass_deactivate
func _gotk4_gtk3_MenuShellClass_deactivate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Deactivate() })

	iface.Deactivate()
}

//export _gotk4_gtk3_MenuShellClass_get_popup_delay
func _gotk4_gtk3_MenuShellClass_get_popup_delay(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PopupDelay() int32 })

	gint := iface.PopupDelay()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_MenuShellClass_insert
func _gotk4_gtk3_MenuShellClass_insert(arg0 *C.void, arg1 *C.void, arg2 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Insert(child Widgetter, position int32)
	})

	var _child Widgetter // out
	var _position int32  // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}
	_position = int32(arg2)

	iface.Insert(_child, _position)
}

//export _gotk4_gtk3_MenuShellClass_move_selected
func _gotk4_gtk3_MenuShellClass_move_selected(arg0 *C.void, arg1 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ MoveSelected(distance int32) bool })

	var _distance int32 // out

	_distance = int32(arg1)

	ok := iface.MoveSelected(_distance)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_MenuShellClass_select_item
func _gotk4_gtk3_MenuShellClass_select_item(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SelectItem(menuItem Widgetter) })

	var _menuItem Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_menuItem = rv
	}

	iface.SelectItem(_menuItem)
}

//export _gotk4_gtk3_MenuShellClass_selection_done
func _gotk4_gtk3_MenuShellClass_selection_done(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SelectionDone() })

	iface.SelectionDone()
}

func wrapMenuShell(obj *coreglib.Object) *MenuShell {
	return &MenuShell{
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

func marshalMenuShell(p uintptr) (interface{}, error) {
	return wrapMenuShell(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (menuShell *MenuShell) baseMenuShell() *MenuShell {
	return menuShell
}

// BaseMenuShell returns the underlying base object.
func BaseMenuShell(obj MenuSheller) *MenuShell {
	return obj.baseMenuShell()
}

//export _gotk4_gtk3_MenuShell_ConnectActivateCurrent
func _gotk4_gtk3_MenuShell_ConnectActivateCurrent(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(forceHide bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(forceHide bool))
	}

	var _forceHide bool // out

	if arg1 != 0 {
		_forceHide = true
	}

	f(_forceHide)
}

// ConnectActivateCurrent: action signal that activates the current menu item
// within the menu shell.
func (menuShell *MenuShell) ConnectActivateCurrent(f func(forceHide bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "activate-current", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectActivateCurrent), f)
}

//export _gotk4_gtk3_MenuShell_ConnectCancel
func _gotk4_gtk3_MenuShell_ConnectCancel(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectCancel: action signal which cancels the selection within the menu
// shell. Causes the MenuShell::selection-done signal to be emitted.
func (menuShell *MenuShell) ConnectCancel(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "cancel", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectCancel), f)
}

//export _gotk4_gtk3_MenuShell_ConnectDeactivate
func _gotk4_gtk3_MenuShell_ConnectDeactivate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectDeactivate: this signal is emitted when a menu shell is deactivated.
func (menuShell *MenuShell) ConnectDeactivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "deactivate", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectDeactivate), f)
}

//export _gotk4_gtk3_MenuShell_ConnectInsert
func _gotk4_gtk3_MenuShell_ConnectInsert(arg0 C.gpointer, arg1 *C.void, arg2 C.gint, arg3 C.guintptr) {
	var f func(child Widgetter, position int32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(child Widgetter, position int32))
	}

	var _child Widgetter // out
	var _position int32  // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}
	_position = int32(arg2)

	f(_child, _position)
}

// ConnectInsert signal is emitted when a new MenuItem is added to a MenuShell.
// A separate signal is used instead of GtkContainer::add because of the need
// for an additional position parameter.
//
// The inverse of this signal is the GtkContainer::removed signal.
func (menuShell *MenuShell) ConnectInsert(f func(child Widgetter, position int32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "insert", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectInsert), f)
}

//export _gotk4_gtk3_MenuShell_ConnectMoveSelected
func _gotk4_gtk3_MenuShell_ConnectMoveSelected(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) (cret C.gboolean) {
	var f func(distance int32) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(distance int32) (ok bool))
	}

	var _distance int32 // out

	_distance = int32(arg1)

	ok := f(_distance)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectMoveSelected signal is emitted to move the selection to another item.
func (menuShell *MenuShell) ConnectMoveSelected(f func(distance int32) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "move-selected", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectMoveSelected), f)
}

//export _gotk4_gtk3_MenuShell_ConnectSelectionDone
func _gotk4_gtk3_MenuShell_ConnectSelectionDone(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectSelectionDone: this signal is emitted when a selection has been
// completed within a menu shell.
func (menuShell *MenuShell) ConnectSelectionDone(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuShell, "selection-done", false, unsafe.Pointer(C._gotk4_gtk3_MenuShell_ConnectSelectionDone), f)
}

// ActivateItem activates the menu item within the menu shell.
//
// The function takes the following parameters:
//
//    - menuItem to activate.
//    - forceDeactivate: if TRUE, force the deactivation of the menu shell after
//      the menu item is activated.
//
func (menuShell *MenuShell) ActivateItem(menuItem Widgetter, forceDeactivate bool) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if forceDeactivate {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("activate_item", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(forceDeactivate)
}

// Append adds a new MenuItem to the end of the menu shell's item list.
//
// The function takes the following parameters:
//
//    - child to add.
//
func (menuShell *MenuShell) Append(child *MenuItem) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("append", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
}

// BindModel establishes a binding between a MenuShell and a Model.
//
// The contents of shell are removed and then refilled with menu items according
// to model. When model changes, shell is updated. Calling this function twice
// on shell with different model will cause the first binding to be replaced
// with a binding to the new model. If model is NULL then any previous binding
// is undone and all children are removed.
//
// with_separators determines if toplevel items (eg: sections) have separators
// inserted between them. This is typically desired for menus but doesn’t make
// sense for menubars.
//
// If action_namespace is non-NULL then the effect is as if all actions
// mentioned in the model have their names prefixed with the namespace, plus a
// dot. For example, if the action “quit” is mentioned and action_namespace is
// “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values on
// the created menu items. If you want to use an action group other than “app”
// and “win”, or if you want to use a MenuShell outside of a ApplicationWindow,
// then you will need to attach your own action group to the widget hierarchy
// using gtk_widget_insert_action_group(). As an example, if you created a group
// with a “quit” action and inserted it with the name “mygroup” then you would
// use the action name “mygroup.quit” in your Model.
//
// For most cases you are probably better off using gtk_menu_new_from_model() or
// gtk_menu_bar_new_from_model() or just directly passing the Model to
// gtk_application_set_app_menu() or gtk_application_set_menubar().
//
// The function takes the following parameters:
//
//    - model (optional) to bind to or NULL to remove binding.
//    - actionNamespace (optional): namespace for actions in model.
//    - withSeparators: TRUE if toplevel items in shell should have separators
//      between them.
//
func (menuShell *MenuShell) BindModel(model gio.MenuModeller, actionNamespace string, withSeparators bool) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	if actionNamespace != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(actionNamespace)))
		defer C.free(unsafe.Pointer(_args[2]))
	}
	if withSeparators {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("bind_model", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(model)
	runtime.KeepAlive(actionNamespace)
	runtime.KeepAlive(withSeparators)
}

// Cancel cancels the selection within the menu shell.
func (menuShell *MenuShell) Cancel() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("cancel", _args[:], nil)

	runtime.KeepAlive(menuShell)
}

// Deactivate deactivates the menu shell.
//
// Typically this results in the menu shell being erased from the screen.
func (menuShell *MenuShell) Deactivate() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("deactivate", _args[:], nil)

	runtime.KeepAlive(menuShell)
}

// Deselect deselects the currently selected item from the menu shell, if any.
func (menuShell *MenuShell) Deselect() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("deselect", _args[:], nil)

	runtime.KeepAlive(menuShell)
}

// ParentShell gets the parent menu shell.
//
// The parent menu shell of a submenu is the Menu or MenuBar from which it was
// opened up.
//
// The function returns the following values:
//
//    - widget: parent MenuShell.
//
func (menuShell *MenuShell) ParentShell() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	_gret := girepository.MustFind("Gtk", "MenuShell").InvokeMethod("get_parent_shell", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuShell)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// SelectedItem gets the currently selected item.
//
// The function returns the following values:
//
//    - widget: currently selected item.
//
func (menuShell *MenuShell) SelectedItem() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	_gret := girepository.MustFind("Gtk", "MenuShell").InvokeMethod("get_selected_item", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuShell)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// TakeFocus returns TRUE if the menu shell will take the keyboard focus on
// popup.
//
// The function returns the following values:
//
//    - ok: TRUE if the menu shell will take the keyboard focus on popup.
//
func (menuShell *MenuShell) TakeFocus() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))

	_gret := girepository.MustFind("Gtk", "MenuShell").InvokeMethod("get_take_focus", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuShell)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Insert adds a new MenuItem to the menu shell’s item list at the position
// indicated by position.
//
// The function takes the following parameters:
//
//    - child to add.
//    - position in the item list where child is added. Positions are numbered
//      from 0 to n-1.
//
func (menuShell *MenuShell) Insert(child Widgetter, position int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(position)

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("insert", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// Prepend adds a new MenuItem to the beginning of the menu shell's item list.
//
// The function takes the following parameters:
//
//    - child to add.
//
func (menuShell *MenuShell) Prepend(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("prepend", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
}

// SelectFirst: select the first visible or selectable child of the menu shell;
// don’t select tearoff items unless the only item is a tearoff item.
//
// The function takes the following parameters:
//
//    - searchSensitive: if TRUE, search for the first selectable menu item,
//      otherwise select nothing if the first item isn’t sensitive. This should
//      be FALSE if the menu is being popped up initially.
//
func (menuShell *MenuShell) SelectFirst(searchSensitive bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	if searchSensitive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("select_first", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(searchSensitive)
}

// SelectItem selects the menu item from the menu shell.
//
// The function takes the following parameters:
//
//    - menuItem to select.
//
func (menuShell *MenuShell) SelectItem(menuItem Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("select_item", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(menuItem)
}

// SetTakeFocus: if take_focus is TRUE (the default) the menu shell will take
// the keyboard focus so that it will receive all keyboard events which is
// needed to enable keyboard navigation in menus.
//
// Setting take_focus to FALSE is useful only for special applications like
// virtual keyboard implementations which should not take keyboard focus.
//
// The take_focus state of a menu or menu bar is automatically propagated to
// submenus whenever a submenu is popped up, so you don’t have to worry about
// recursively setting it for your entire menu hierarchy. Only when
// programmatically picking a submenu and popping it up manually, the take_focus
// property of the submenu needs to be set explicitly.
//
// Note that setting it to FALSE has side-effects:
//
// If the focus is in some other app, it keeps the focus and keynav in the menu
// doesn’t work. Consequently, keynav on the menu will only work if the focus is
// on some toplevel owned by the onscreen keyboard.
//
// To avoid confusing the user, menus with take_focus set to FALSE should not
// display mnemonics or accelerators, since it cannot be guaranteed that they
// will work.
//
// See also gdk_keyboard_grab().
//
// The function takes the following parameters:
//
//    - takeFocus: TRUE if the menu shell should take the keyboard focus on
//      popup.
//
func (menuShell *MenuShell) SetTakeFocus(takeFocus bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	if takeFocus {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MenuShell").InvokeMethod("set_take_focus", _args[:], nil)

	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(takeFocus)
}
