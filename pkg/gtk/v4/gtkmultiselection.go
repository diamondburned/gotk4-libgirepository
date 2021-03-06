// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeMultiSelection returns the GType for the type MultiSelection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMultiSelection() coreglib.Type {
	gtype := coreglib.Type(C.gtk_multi_selection_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalMultiSelection)
	return gtype
}

// MultiSelectionOverrider contains methods that are overridable.
type MultiSelectionOverrider interface {
}

// MultiSelection: GtkMultiSelection is a GtkSelectionModel that allows
// selecting multiple elements.
type MultiSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SelectionModel
}

var (
	_ coreglib.Objector = (*MultiSelection)(nil)
)

func classInitMultiSelectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMultiSelection(obj *coreglib.Object) *MultiSelection {
	return &MultiSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	return wrapMultiSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultiSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to manage, or NULL.
//
// The function returns the following values:
//
//    - multiSelection: new GtkMultiSelection.
//
func NewMultiSelection(model gio.ListModeller) *MultiSelection {
	var _arg1 *C.GListModel        // out
	var _cret *C.GtkMultiSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_cret = C.gtk_multi_selection_new(_arg1)
	runtime.KeepAlive(model)

	var _multiSelection *MultiSelection // out

	_multiSelection = wrapMultiSelection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multiSelection
}

// Model returns the underlying model of self.
//
// The function returns the following values:
//
//    - listModel: underlying model.
//
func (self *MultiSelection) Model() *gio.ListModel {
	var _arg0 *C.GtkMultiSelection // out
	var _cret *C.GListModel        // in

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_multi_selection_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to wrap.
//
func (self *MultiSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkMultiSelection // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_multi_selection_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
