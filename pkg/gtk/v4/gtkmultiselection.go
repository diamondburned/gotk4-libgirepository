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
		{T: externglib.Type(C.gtk_multi_selection_get_type()), F: marshalMultiSelection},
	})
}

// MultiSelection: `GtkMultiSelection` is a `GtkSelectionModel` that allows
// selecting multiple elements.
type MultiSelection interface {
	gextras.Objector
	gio.ListModel
	SelectionModel

	// Model returns the underlying model of @self.
	Model() gio.ListModel
	// SetModel sets the model that @self should wrap.
	//
	// If @model is nil, @self will be empty.
	SetModel(model gio.ListModel)
}

// multiSelection implements the MultiSelection interface.
type multiSelection struct {
	gextras.Objector
	gio.ListModel
	SelectionModel
}

var _ MultiSelection = (*multiSelection)(nil)

// WrapMultiSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapMultiSelection(obj *externglib.Object) MultiSelection {
	return MultiSelection{
		Objector:       obj,
		gio.ListModel:  gio.WrapListModel(obj),
		SelectionModel: WrapSelectionModel(obj),
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiSelection(obj), nil
}

// NewMultiSelection constructs a class MultiSelection.
func NewMultiSelection(model gio.ListModel) MultiSelection {
	var _arg1 *C.GListModel

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkMultiSelection

	cret = C.gtk_multi_selection_new(_arg1)

	var _multiSelection MultiSelection

	_multiSelection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MultiSelection)

	return _multiSelection
}

// Model returns the underlying model of @self.
func (s multiSelection) Model() gio.ListModel {
	var _arg0 *C.GtkMultiSelection

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel

	cret = C.gtk_multi_selection_get_model(_arg0)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, @self will be empty.
func (s multiSelection) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkMultiSelection
	var _arg1 *C.GListModel

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_multi_selection_set_model(_arg0, _arg1)
}