// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModel},
	})
}

// SelectionFilterModel: `GtkSelectionFilterModel` is a list model that presents
// the selection from a `GtkSelectionModel`.
type SelectionFilterModel interface {
	gextras.Objector
	gio.ListModel

	// SetModel sets the model to be filtered.
	//
	// Note that GTK makes no effort to ensure that @model conforms to the item
	// type of @self. It assumes that the caller knows what they are doing and
	// have set up an appropriate filter to ensure that item types match.
	SetModel(model SelectionModel)
}

// selectionFilterModel implements the SelectionFilterModel class.
type selectionFilterModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SelectionFilterModel = (*selectionFilterModel)(nil)

// WrapSelectionFilterModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSelectionFilterModel(obj *externglib.Object) SelectionFilterModel {
	return selectionFilterModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSelectionFilterModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSelectionFilterModel(obj), nil
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that @model conforms to the item
// type of @self. It assumes that the caller knows what they are doing and
// have set up an appropriate filter to ensure that item types match.
func (s selectionFilterModel) SetModel(model SelectionModel) {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _arg1 *C.GtkSelectionModel       // out

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_selection_filter_model_set_model(_arg0, _arg1)
}
