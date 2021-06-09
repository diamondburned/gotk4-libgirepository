// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_radio_action_get_type()), F: marshalRadioAction},
	})
}

// RadioAction: a RadioAction is similar to RadioMenuItem. A number of radio
// actions can be linked together so that only one may be active at any one
// time.
type RadioAction interface {
	ToggleAction
	Buildable

	// CurrentValue obtains the value property of the currently active member of
	// the group to which @action belongs.
	CurrentValue() int
	// Group returns the list representing the radio group for this object. Note
	// that the returned list is only valid until the next change to the group.
	//
	// A common way to set up a group of radio group is the following:
	//
	//     GSList *group = NULL;
	//     GtkRadioAction *action;
	//
	//     while ( ...more actions to add... /)
	//       {
	//          action = gtk_radio_action_new (...);
	//
	//          gtk_radio_action_set_group (action, group);
	//          group = gtk_radio_action_get_group (action);
	//       }
	Group() *glib.SList
	// JoinGroup joins a radio action object to the group of another radio
	// action object.
	//
	// Use this in language bindings instead of the gtk_radio_action_get_group()
	// and gtk_radio_action_set_group() methods
	//
	// A common way to set up a group of radio actions is the following:
	//
	//     GtkRadioAction *action;
	//     GtkRadioAction *last_action;
	//
	//     while ( ...more actions to add... /)
	//       {
	//          action = gtk_radio_action_new (...);
	//
	//          gtk_radio_action_join_group (action, last_action);
	//          last_action = action;
	//       }
	JoinGroup(groupSource RadioAction)
	// SetCurrentValue sets the currently active group member to the member with
	// value property @current_value.
	SetCurrentValue(currentValue int)
	// SetGroup sets the radio group for the radio action object.
	SetGroup(group *glib.SList)
}

// radioAction implements the RadioAction interface.
type radioAction struct {
	ToggleAction
	Buildable
}

var _ RadioAction = (*radioAction)(nil)

// WrapRadioAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioAction(obj *externglib.Object) RadioAction {
	return RadioAction{
		ToggleAction: WrapToggleAction(obj),
		Buildable:    WrapBuildable(obj),
	}
}

func marshalRadioAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioAction(obj), nil
}

// NewRadioAction constructs a class RadioAction.
func NewRadioAction(name string, label string, tooltip string, stockId string, value int) RadioAction {
	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _arg3 *C.gchar
	var _arg4 *C.gchar
	var _arg5 C.gint

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gint(value)

	var _cret C.GtkRadioAction

	cret = C.gtk_radio_action_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _radioAction RadioAction

	_radioAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(RadioAction)

	return _radioAction
}

// CurrentValue obtains the value property of the currently active member of
// the group to which @action belongs.
func (a radioAction) CurrentValue() int {
	var _arg0 *C.GtkRadioAction

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))

	var _cret C.gint

	cret = C.gtk_radio_action_get_current_value(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Group returns the list representing the radio group for this object. Note
// that the returned list is only valid until the next change to the group.
//
// A common way to set up a group of radio group is the following:
//
//     GSList *group = NULL;
//     GtkRadioAction *action;
//
//     while ( ...more actions to add... /)
//       {
//          action = gtk_radio_action_new (...);
//
//          gtk_radio_action_set_group (action, group);
//          group = gtk_radio_action_get_group (action);
//       }
func (a radioAction) Group() *glib.SList {
	var _arg0 *C.GtkRadioAction

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))

	var _cret *C.GSList

	cret = C.gtk_radio_action_get_group(_arg0)

	var _sList *glib.SList

	_sList = glib.WrapSList(unsafe.Pointer(_cret))

	return _sList
}

// JoinGroup joins a radio action object to the group of another radio
// action object.
//
// Use this in language bindings instead of the gtk_radio_action_get_group()
// and gtk_radio_action_set_group() methods
//
// A common way to set up a group of radio actions is the following:
//
//     GtkRadioAction *action;
//     GtkRadioAction *last_action;
//
//     while ( ...more actions to add... /)
//       {
//          action = gtk_radio_action_new (...);
//
//          gtk_radio_action_join_group (action, last_action);
//          last_action = action;
//       }
func (a radioAction) JoinGroup(groupSource RadioAction) {
	var _arg0 *C.GtkRadioAction
	var _arg1 *C.GtkRadioAction

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkRadioAction)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_action_join_group(_arg0, _arg1)
}

// SetCurrentValue sets the currently active group member to the member with
// value property @current_value.
func (a radioAction) SetCurrentValue(currentValue int) {
	var _arg0 *C.GtkRadioAction
	var _arg1 C.gint

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(currentValue)

	C.gtk_radio_action_set_current_value(_arg0, _arg1)
}

// SetGroup sets the radio group for the radio action object.
func (a radioAction) SetGroup(group *glib.SList) {
	var _arg0 *C.GtkRadioAction
	var _arg1 *C.GSList

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	C.gtk_radio_action_set_group(_arg0, _arg1)
}