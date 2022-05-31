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

// glib.Type values for gtkselectionfiltermodel.go.
var GTypeSelectionFilterModel = coreglib.Type(C.gtk_selection_filter_model_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSelectionFilterModel, F: marshalSelectionFilterModel},
	})
}

// SelectionFilterModelOverrider contains methods that are overridable.
type SelectionFilterModelOverrider interface {
}

// SelectionFilterModel: GtkSelectionFilterModel is a list model that presents
// the selection from a GtkSelectionModel.
type SelectionFilterModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*SelectionFilterModel)(nil)
)

func classInitSelectionFilterModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSelectionFilterModel(obj *coreglib.Object) *SelectionFilterModel {
	return &SelectionFilterModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSelectionFilterModel(p uintptr) (interface{}, error) {
	return wrapSelectionFilterModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSelectionFilterModel creates a new GtkSelectionFilterModel that will
// include the selected items from the underlying selection model.
//
// The function takes the following parameters:
//
//    - model (optional): selection model to filter, or NULL.
//
// The function returns the following values:
//
//    - selectionFilterModel: new GtkSelectionFilterModel.
//
func NewSelectionFilterModel(model SelectionModeller) *SelectionFilterModel {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if model != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SelectionFilterModel").InvokeMethod("new_SelectionFilterModel", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _selectionFilterModel *SelectionFilterModel // out

	_selectionFilterModel = wrapSelectionFilterModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _selectionFilterModel
}

// Model gets the model currently filtered or NULL if none.
//
// The function returns the following values:
//
//    - selectionModel (optional): model that gets filtered.
//
func (self *SelectionFilterModel) Model() *SelectionModel {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SelectionFilterModel").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _selectionModel *SelectionModel // out

	if _cret != nil {
		_selectionModel = wrapSelectionModel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _selectionModel
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that model conforms to the item type
// of self. It assumes that the caller knows what they are doing and have set up
// an appropriate filter to ensure that item types match.
//
// The function takes the following parameters:
//
//    - model (optional) to be filtered.
//
func (self *SelectionFilterModel) SetModel(model SelectionModeller) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "SelectionFilterModel").InvokeMethod("set_model", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
