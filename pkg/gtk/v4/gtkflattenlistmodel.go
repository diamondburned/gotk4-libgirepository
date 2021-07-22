// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flatten_list_model_get_type()), F: marshalFlattenListModeller},
	})
}

// FlattenListModel: GtkFlattenListModel is a list model that concatenates other
// list models.
//
// GtkFlattenListModel takes a list model containing list models, and flattens
// it into a single model.
type FlattenListModel struct {
	*externglib.Object

	gio.ListModel
}

func wrapFlattenListModel(obj *externglib.Object) *FlattenListModel {
	return &FlattenListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalFlattenListModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFlattenListModel(obj), nil
}

// NewFlattenListModel creates a new GtkFlattenListModel that flattens list.
func NewFlattenListModel(model gio.ListModeller) *FlattenListModel {
	var _arg1 *C.GListModel          // out
	var _cret *C.GtkFlattenListModel // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_flatten_list_model_new(_arg1)

	var _flattenListModel *FlattenListModel // out

	_flattenListModel = wrapFlattenListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _flattenListModel
}

// Model gets the model set via gtk_flatten_list_model_set_model().
func (self *FlattenListModel) Model() gio.ListModeller {
	var _arg0 *C.GtkFlattenListModel // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_flatten_list_model_get_model(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// ModelForItem returns the model containing the item at the given position.
func (self *FlattenListModel) ModelForItem(position uint) gio.ListModeller {
	var _arg0 *C.GtkFlattenListModel // out
	var _arg1 C.guint                // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_flatten_list_model_get_model_for_item(_arg0, _arg1)

	var _listModel gio.ListModeller // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// SetModel sets a new model to be flattened.
func (self *FlattenListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkFlattenListModel // out
	var _arg1 *C.GListModel          // out

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_flatten_list_model_set_model(_arg0, _arg1)
}
