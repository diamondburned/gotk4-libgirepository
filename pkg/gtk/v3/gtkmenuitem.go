// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gchar* _gotk4_gtk3_MenuItemClass_get_label(GtkMenuItem*);
// extern void _gotk4_gtk3_MenuItemClass_activate(GtkMenuItem*);
// extern void _gotk4_gtk3_MenuItemClass_activate_item(GtkMenuItem*);
// extern void _gotk4_gtk3_MenuItemClass_deselect(GtkMenuItem*);
// extern void _gotk4_gtk3_MenuItemClass_select(GtkMenuItem*);
// extern void _gotk4_gtk3_MenuItemClass_set_label(GtkMenuItem*, gchar*);
// extern void _gotk4_gtk3_MenuItemClass_toggle_size_allocate(GtkMenuItem*, gint);
// extern void _gotk4_gtk3_MenuItem_ConnectActivate(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuItem_ConnectActivateItem(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuItem_ConnectDeselect(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuItem_ConnectSelect(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuItem_ConnectToggleSizeAllocate(gpointer, gint, guintptr);
// extern void _gotk4_gtk3_MenuItem_ConnectToggleSizeRequest(gpointer, gpointer, guintptr);
import "C"

// GTypeMenuItem returns the GType for the type MenuItem.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMenuItem() coreglib.Type {
	gtype := coreglib.Type(C.gtk_menu_item_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalMenuItem)
	return gtype
}

// MenuItemOverrider contains methods that are overridable.
type MenuItemOverrider interface {
	// Activate emits the MenuItem::activate signal on the given item.
	Activate()
	ActivateItem()
	// Deselect emits the MenuItem::deselect signal on the given item.
	Deselect()
	// Label sets text on the menu_item label.
	//
	// The function returns the following values:
	//
	//    - utf8: text in the menu_item label. This is the internal string used
	//      by the label, and must not be modified.
	//
	Label() string
	// Select emits the MenuItem::select signal on the given item.
	Select()
	// SetLabel sets text on the menu_item label.
	//
	// The function takes the following parameters:
	//
	//    - label: text you want to set.
	//
	SetLabel(label string)
	// ToggleSizeAllocate emits the MenuItem::toggle-size-allocate signal on the
	// given item.
	//
	// The function takes the following parameters:
	//
	//    - allocation to use as signal data.
	//
	ToggleSizeAllocate(allocation int)
}

// MenuItem widget and the derived widgets are the only valid children for
// menus. Their function is to correctly handle highlighting, alignment, events
// and submenus.
//
// As a GtkMenuItem derives from Bin it can hold any valid child widget,
// although only a few are really useful.
//
// By default, a GtkMenuItem sets a AccelLabel as its child. GtkMenuItem has
// direct functions to set the label and its mnemonic. For more advanced label
// settings, you can fetch the child widget from the GtkBin.
//
// An example for setting markup and accelerator on a MenuItem:
//
//    menuitem
//    ├── <child>
//    ╰── [arrow.right]
//
// GtkMenuItem has a single CSS node with name menuitem. If the menuitem has a
// submenu, it gets another CSS node with name arrow, which has the .left or
// .right style class.
type MenuItem struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Actionable
	Activatable
}

var (
	_ Binner            = (*MenuItem)(nil)
	_ coreglib.Objector = (*MenuItem)(nil)
)

func classInitMenuItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkMenuItemClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Activate() }); ok {
		pclass.activate = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_activate)
	}

	if _, ok := goval.(interface{ ActivateItem() }); ok {
		pclass.activate_item = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_activate_item)
	}

	if _, ok := goval.(interface{ Deselect() }); ok {
		pclass.deselect = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_deselect)
	}

	if _, ok := goval.(interface{ Label() string }); ok {
		pclass.get_label = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_get_label)
	}

	if _, ok := goval.(interface{ Select() }); ok {
		pclass._select = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_select)
	}

	if _, ok := goval.(interface{ SetLabel(label string) }); ok {
		pclass.set_label = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_set_label)
	}

	if _, ok := goval.(interface{ ToggleSizeAllocate(allocation int) }); ok {
		pclass.toggle_size_allocate = (*[0]byte)(C._gotk4_gtk3_MenuItemClass_toggle_size_allocate)
	}
}

//export _gotk4_gtk3_MenuItemClass_activate
func _gotk4_gtk3_MenuItemClass_activate(arg0 *C.GtkMenuItem) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

//export _gotk4_gtk3_MenuItemClass_activate_item
func _gotk4_gtk3_MenuItemClass_activate_item(arg0 *C.GtkMenuItem) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateItem() })

	iface.ActivateItem()
}

//export _gotk4_gtk3_MenuItemClass_deselect
func _gotk4_gtk3_MenuItemClass_deselect(arg0 *C.GtkMenuItem) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Deselect() })

	iface.Deselect()
}

//export _gotk4_gtk3_MenuItemClass_get_label
func _gotk4_gtk3_MenuItemClass_get_label(arg0 *C.GtkMenuItem) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Label() string })

	utf8 := iface.Label()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk3_MenuItemClass_select
func _gotk4_gtk3_MenuItemClass_select(arg0 *C.GtkMenuItem) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Select() })

	iface.Select()
}

//export _gotk4_gtk3_MenuItemClass_set_label
func _gotk4_gtk3_MenuItemClass_set_label(arg0 *C.GtkMenuItem, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SetLabel(label string) })

	var _label string // out

	_label = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetLabel(_label)
}

//export _gotk4_gtk3_MenuItemClass_toggle_size_allocate
func _gotk4_gtk3_MenuItemClass_toggle_size_allocate(arg0 *C.GtkMenuItem, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ToggleSizeAllocate(allocation int) })

	var _allocation int // out

	_allocation = int(arg1)

	iface.ToggleSizeAllocate(_allocation)
}

func wrapMenuItem(obj *coreglib.Object) *MenuItem {
	return &MenuItem{
		Bin: Bin{
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
		},
		Object: obj,
		Actionable: Actionable{
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
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	return wrapMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_MenuItem_ConnectActivate
func _gotk4_gtk3_MenuItem_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectActivate is emitted when the item is activated.
func (menuItem *MenuItem) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "activate", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectActivate), f)
}

//export _gotk4_gtk3_MenuItem_ConnectActivateItem
func _gotk4_gtk3_MenuItem_ConnectActivateItem(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectActivateItem is emitted when the item is activated, but also if the
// menu item has a submenu. For normal applications, the relevant signal is
// MenuItem::activate.
func (menuItem *MenuItem) ConnectActivateItem(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "activate-item", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectActivateItem), f)
}

//export _gotk4_gtk3_MenuItem_ConnectDeselect
func _gotk4_gtk3_MenuItem_ConnectDeselect(arg0 C.gpointer, arg1 C.guintptr) {
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

func (menuItem *MenuItem) ConnectDeselect(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "deselect", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectDeselect), f)
}

//export _gotk4_gtk3_MenuItem_ConnectSelect
func _gotk4_gtk3_MenuItem_ConnectSelect(arg0 C.gpointer, arg1 C.guintptr) {
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

func (menuItem *MenuItem) ConnectSelect(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "select", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectSelect), f)
}

//export _gotk4_gtk3_MenuItem_ConnectToggleSizeAllocate
func _gotk4_gtk3_MenuItem_ConnectToggleSizeAllocate(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(object int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object int))
	}

	var _object int // out

	_object = int(arg1)

	f(_object)
}

func (menuItem *MenuItem) ConnectToggleSizeAllocate(f func(object int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "toggle-size-allocate", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectToggleSizeAllocate), f)
}

//export _gotk4_gtk3_MenuItem_ConnectToggleSizeRequest
func _gotk4_gtk3_MenuItem_ConnectToggleSizeRequest(arg0 C.gpointer, arg1 C.gpointer, arg2 C.guintptr) {
	var f func(object unsafe.Pointer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object unsafe.Pointer))
	}

	var _object unsafe.Pointer // out

	_object = (unsafe.Pointer)(unsafe.Pointer(arg1))

	f(_object)
}

func (menuItem *MenuItem) ConnectToggleSizeRequest(f func(object unsafe.Pointer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menuItem, "toggle-size-request", false, unsafe.Pointer(C._gotk4_gtk3_MenuItem_ConnectToggleSizeRequest), f)
}

// NewMenuItem creates a new MenuItem.
//
// The function returns the following values:
//
//    - menuItem: newly created MenuItem.
//
func NewMenuItem() *MenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_item_new()

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuItem
}

// NewMenuItemWithLabel creates a new MenuItem whose child is a Label.
//
// The function takes the following parameters:
//
//    - label: text for the label.
//
// The function returns the following values:
//
//    - menuItem: newly created MenuItem.
//
func NewMenuItemWithLabel(label string) *MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_label(_arg1)
	runtime.KeepAlive(label)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuItem
}

// NewMenuItemWithMnemonic creates a new MenuItem containing a label.
//
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - menuItem: new MenuItem.
//
func NewMenuItemWithMnemonic(label string) *MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuItem
}

// Activate emits the MenuItem::activate signal on the given item.
func (menuItem *MenuItem) Activate() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	C.gtk_menu_item_activate(_arg0)
	runtime.KeepAlive(menuItem)
}

// Deselect emits the MenuItem::deselect signal on the given item.
func (menuItem *MenuItem) Deselect() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	C.gtk_menu_item_deselect(_arg0)
	runtime.KeepAlive(menuItem)
}

// AccelPath: retrieve the accelerator path that was previously set on
// menu_item.
//
// See gtk_menu_item_set_accel_path() for details.
//
// The function returns the following values:
//
//    - utf8 (optional): accelerator path corresponding to this menu item’s
//      functionality, or NULL if not set.
//
func (menuItem *MenuItem) AccelPath() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_accel_path(_arg0)
	runtime.KeepAlive(menuItem)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Label sets text on the menu_item label.
//
// The function returns the following values:
//
//    - utf8: text in the menu_item label. This is the internal string used by
//      the label, and must not be modified.
//
func (menuItem *MenuItem) Label() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_label(_arg0)
	runtime.KeepAlive(menuItem)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ReserveIndicator returns whether the menu_item reserves space for the submenu
// indicator, regardless if it has a submenu or not.
//
// The function returns the following values:
//
//    - ok: TRUE if menu_item always reserves space for the submenu indicator.
//
func (menuItem *MenuItem) ReserveIndicator() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_reserve_indicator(_arg0)
	runtime.KeepAlive(menuItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RightJustified gets whether the menu item appears justified at the right side
// of the menu bar.
//
// Deprecated: See gtk_menu_item_set_right_justified().
//
// The function returns the following values:
//
//    - ok: TRUE if the menu item will appear at the far right if added to a menu
//      bar.
//
func (menuItem *MenuItem) RightJustified() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_right_justified(_arg0)
	runtime.KeepAlive(menuItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Submenu gets the submenu underneath this menu item, if any. See
// gtk_menu_item_set_submenu().
//
// The function returns the following values:
//
//    - widget (optional): submenu for this menu item, or NULL if none.
//
func (menuItem *MenuItem) Submenu() Widgetter {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_submenu(_arg0)
	runtime.KeepAlive(menuItem)

	var _widget Widgetter // out

	if _cret != nil {
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

// UseUnderline checks if an underline in the text indicates the next character
// should be used for the mnemonic accelerator key.
//
// The function returns the following values:
//
//    - ok: TRUE if an embedded underline in the label indicates the mnemonic
//      accelerator key.
//
func (menuItem *MenuItem) UseUnderline() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	_cret = C.gtk_menu_item_get_use_underline(_arg0)
	runtime.KeepAlive(menuItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Select emits the MenuItem::select signal on the given item.
func (menuItem *MenuItem) Select() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))

	C.gtk_menu_item_select(_arg0)
	runtime.KeepAlive(menuItem)
}

// SetAccelPath: set the accelerator path on menu_item, through which runtime
// changes of the menu item’s accelerator caused by the user can be identified
// and saved to persistent storage (see gtk_accel_map_save() on this). To set up
// a default accelerator for this menu item, call gtk_accel_map_add_entry() with
// the same accel_path. See also gtk_accel_map_add_entry() on the specifics of
// accelerator paths, and gtk_menu_set_accel_path() for a more convenient
// variant of this function.
//
// This function is basically a convenience wrapper that handles calling
// gtk_widget_set_accel_path() with the appropriate accelerator group for the
// menu item.
//
// Note that you do need to set an accelerator on the parent menu with
// gtk_menu_set_accel_group() for this to work.
//
// Note that accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
//
// The function takes the following parameters:
//
//    - accelPath (optional): accelerator path, corresponding to this menu item’s
//      functionality, or NULL to unset the current path.
//
func (menuItem *MenuItem) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if accelPath != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_menu_item_set_accel_path(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(accelPath)
}

// SetLabel sets text on the menu_item label.
//
// The function takes the following parameters:
//
//    - label: text you want to set.
//
func (menuItem *MenuItem) SetLabel(label string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_item_set_label(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(label)
}

// SetReserveIndicator sets whether the menu_item should reserve space for the
// submenu indicator, regardless if it actually has a submenu or not.
//
// There should be little need for applications to call this functions.
//
// The function takes the following parameters:
//
//    - reserve: new value.
//
func (menuItem *MenuItem) SetReserveIndicator(reserve bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if reserve {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_reserve_indicator(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(reserve)
}

// SetRightJustified sets whether the menu item appears justified at the right
// side of a menu bar. This was traditionally done for “Help” menu items, but is
// now considered a bad idea. (If the widget layout is reversed for a
// right-to-left language like Hebrew or Arabic, right-justified-menu-items
// appear at the left.)
//
// Deprecated: If you insist on using it, use gtk_widget_set_hexpand() and
// gtk_widget_set_halign().
//
// The function takes the following parameters:
//
//    - rightJustified: if TRUE the menu item will appear at the far right if
//      added to a menu bar.
//
func (menuItem *MenuItem) SetRightJustified(rightJustified bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if rightJustified {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_right_justified(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(rightJustified)
}

// SetSubmenu sets or replaces the menu item’s submenu, or removes it when a
// NULL submenu is passed.
//
// The function takes the following parameters:
//
//    - submenu (optional): submenu, or NULL.
//
func (menuItem *MenuItem) SetSubmenu(submenu *Menu) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if submenu != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))
	}

	C.gtk_menu_item_set_submenu(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(submenu)
}

// SetUseUnderline: if true, an underline in the text indicates the next
// character should be used for the mnemonic accelerator key.
//
// The function takes the following parameters:
//
//    - setting: TRUE if underlines in the text indicate mnemonics.
//
func (menuItem *MenuItem) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(setting)
}

// ToggleSizeAllocate emits the MenuItem::toggle-size-allocate signal on the
// given item.
//
// The function takes the following parameters:
//
//    - allocation to use as signal data.
//
func (menuItem *MenuItem) ToggleSizeAllocate(allocation int) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = C.gint(allocation)

	C.gtk_menu_item_toggle_size_allocate(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(allocation)
}
