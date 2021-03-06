// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ToggleActionClass_toggled(GtkToggleAction*);
// extern void _gotk4_gtk3_ToggleAction_ConnectToggled(gpointer, guintptr);
import "C"

// GTypeToggleAction returns the GType for the type ToggleAction.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToggleAction() coreglib.Type {
	gtype := coreglib.Type(C.gtk_toggle_action_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalToggleAction)
	return gtype
}

// ToggleActionOverrider contains methods that are overridable.
type ToggleActionOverrider interface {
	// Toggled emits the “toggled” signal on the toggle action.
	//
	// Deprecated: since version 3.10.
	Toggled()
}

// ToggleAction corresponds roughly to a CheckMenuItem. It has an “active” state
// specifying whether the action has been checked or not.
type ToggleAction struct {
	_ [0]func() // equal guard
	Action
}

var (
	_ coreglib.Objector = (*ToggleAction)(nil)
)

func classInitToggleActioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkToggleActionClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Toggled() }); ok {
		pclass.toggled = (*[0]byte)(C._gotk4_gtk3_ToggleActionClass_toggled)
	}
}

//export _gotk4_gtk3_ToggleActionClass_toggled
func _gotk4_gtk3_ToggleActionClass_toggled(arg0 *C.GtkToggleAction) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapToggleAction(obj *coreglib.Object) *ToggleAction {
	return &ToggleAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToggleAction(p uintptr) (interface{}, error) {
	return wrapToggleAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToggleAction_ConnectToggled
func _gotk4_gtk3_ToggleAction_ConnectToggled(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectToggled: should be connected if you wish to perform an action whenever
// the ToggleAction state is changed.
func (action *ToggleAction) ConnectToggled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(action, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_ToggleAction_ConnectToggled), f)
}

// NewToggleAction creates a new ToggleAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//
// The function returns the following values:
//
//    - toggleAction: new ToggleAction.
//
func NewToggleAction(name, label, tooltip, stockId string) *ToggleAction {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.gchar           // out
	var _arg4 *C.gchar           // out
	var _cret *C.GtkToggleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if tooltip != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	if stockId != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_toggle_action_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)

	var _toggleAction *ToggleAction // out

	_toggleAction = wrapToggleAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _toggleAction
}

// Active returns the checked state of the toggle action.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: checked state of the toggle action.
//
func (action *ToggleAction) Active() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_toggle_action_get_active(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: whether the action should have proxies like a radio action.
//
func (action *ToggleAction) DrawAsRadio() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_toggle_action_get_draw_as_radio(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the checked state on the toggle action.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - isActive: whether the action should be checked or not.
//
func (action *ToggleAction) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_active(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(isActive)
}

// SetDrawAsRadio sets whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - drawAsRadio: whether the action should have proxies like a radio action.
//
func (action *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if drawAsRadio {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_draw_as_radio(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(drawAsRadio)
}

// Toggled emits the “toggled” signal on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Toggled() {
	var _arg0 *C.GtkToggleAction // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	C.gtk_toggle_action_toggled(_arg0)
	runtime.KeepAlive(action)
}
