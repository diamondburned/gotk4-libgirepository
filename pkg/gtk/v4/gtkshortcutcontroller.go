// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkshortcutcontroller.go.
var GTypeShortcutController = coreglib.Type(C.gtk_shortcut_controller_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeShortcutController, F: marshalShortcutController},
	})
}

// ShortcutControllerOverrider contains methods that are overridable.
type ShortcutControllerOverrider interface {
}

// ShortcutController: GtkShortcutController is an event controller that manages
// shortcuts.
//
// Most common shortcuts are using this controller implicitly, e.g. by adding a
// mnemonic underline to a GtkLabel, or by installing a key binding using
// gtk_widget_class_add_binding(), or by adding accelerators to global actions
// using gtk_application_set_accels_for_action().
//
// But it is possible to create your own shortcut controller, and add shortcuts
// to it.
//
// GtkShortcutController implements GListModel for querying the shortcuts that
// have been added to it.
//
//
// GtkShortcutController as a GtkBuildable
//
// GtkShortcutControllers can be creates in ui files to set up shortcuts in the
// same place as the widgets.
//
// An example of a UI definition fragment with GtkShortcutController:
//
//      <object class='GtkButton'>
//        <child>
//          <object class='GtkShortcutController'>
//            <property name='scope'>managed</property>
//            <child>
//              <object class='GtkShortcut'>
//                <property name='trigger'>&amp;lt;Control&amp;gt;k</property>
//                <property name='action'>activate</property>
//              </object>
//            </child>
//          </object>
//        </child>
//      </object>
//
//
// This example creates a gtk.ActivateAction for triggering the activate signal
// of the GtkButton. See gtk.ShortcutAction.ParseString for the syntax for other
// kinds of GtkShortcutAction. See gtk.ShortcutTrigger.ParseString to learn more
// about the syntax for triggers.
type ShortcutController struct {
	_ [0]func() // equal guard
	EventController

	*coreglib.Object
	gio.ListModel
	Buildable
}

var (
	_ EventControllerer = (*ShortcutController)(nil)
	_ coreglib.Objector = (*ShortcutController)(nil)
)

func classInitShortcutControllerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapShortcutController(obj *coreglib.Object) *ShortcutController {
	return &ShortcutController{
		EventController: EventController{
			Object: obj,
		},
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalShortcutController(p uintptr) (interface{}, error) {
	return wrapShortcutController(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewShortcutController creates a new shortcut controller.
//
// The function returns the following values:
//
//    - shortcutController: newly created shortcut controller.
//
func NewShortcutController() *ShortcutController {
	_gret := girepository.MustFind("Gtk", "ShortcutController").InvokeMethod("new_ShortcutController", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _shortcutController *ShortcutController // out

	_shortcutController = wrapShortcutController(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutController
}

// NewShortcutControllerForModel creates a new shortcut controller that takes
// its shortcuts from the given list model.
//
// A controller created by this function does not let you add or remove
// individual shortcuts using the shortcut controller api, but you can change
// the contents of the model.
//
// The function takes the following parameters:
//
//    - model: GListModel containing shortcuts.
//
// The function returns the following values:
//
//    - shortcutController: newly created shortcut controller.
//
func NewShortcutControllerForModel(model gio.ListModeller) *ShortcutController {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutController").InvokeMethod("new_ShortcutController_for_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _shortcutController *ShortcutController // out

	_shortcutController = wrapShortcutController(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutController
}

// AddShortcut adds shortcut to the list of shortcuts handled by self.
//
// If this controller uses an external shortcut list, this function does
// nothing.
//
// The function takes the following parameters:
//
//    - shortcut: GtkShortcut.
//
func (self *ShortcutController) AddShortcut(shortcut *Shortcut) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shortcut).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(shortcut).Native()))

	girepository.MustFind("Gtk", "ShortcutController").InvokeMethod("add_shortcut", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(shortcut)
}

// RemoveShortcut removes shortcut from the list of shortcuts handled by self.
//
// If shortcut had not been added to controller or this controller uses an
// external shortcut list, this function does nothing.
//
// The function takes the following parameters:
//
//    - shortcut: GtkShortcut.
//
func (self *ShortcutController) RemoveShortcut(shortcut *Shortcut) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shortcut).Native()))

	girepository.MustFind("Gtk", "ShortcutController").InvokeMethod("remove_shortcut", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(shortcut)
}
