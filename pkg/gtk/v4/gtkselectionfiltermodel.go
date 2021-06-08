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
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModel},
	})
}

// SelectionFilterModel: `GtkSelectionFilterModel` is a list model that presents
// the selection from a `GtkSelectionModel`.
type SelectionFilterModel interface {
	gextras.Objector
	gio.ListModel

	// Model gets the model currently filtered or nil if none.
	Model() SelectionModel
	// SetModel sets the model to be filtered.
	//
	// Note that GTK makes no effort to ensure that @model conforms to the item
	// type of @self. It assumes that the caller knows what they are doing and
	// have set up an appropriate filter to ensure that item types match.
	SetModel(model SelectionModel)
}

// selectionFilterModel implements the SelectionFilterModel interface.
type selectionFilterModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SelectionFilterModel = (*selectionFilterModel)(nil)

// WrapSelectionFilterModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSelectionFilterModel(obj *externglib.Object) SelectionFilterModel {
	return SelectionFilterModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSelectionFilterModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSelectionFilterModel(obj), nil
}

// NewSelectionFilterModel constructs a class SelectionFilterModel.
func NewSelectionFilterModel(model SelectionModel) SelectionFilterModel {
	var arg1 *C.GtkSelectionModel

	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	cret := new(C.GtkSelectionFilterModel)
	var goret SelectionFilterModel

	cret = C.gtk_selection_filter_model_new(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SelectionFilterModel)

	return goret
}

// Model gets the model currently filtered or nil if none.
func (s selectionFilterModel) Model() SelectionModel {
	var arg0 *C.GtkSelectionFilterModel

	arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(s.Native()))

	var cret *C.GtkSelectionModel
	var goret SelectionModel

	cret = C.gtk_selection_filter_model_get_model(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(SelectionModel)

	return goret
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that @model conforms to the item
// type of @self. It assumes that the caller knows what they are doing and
// have set up an appropriate filter to ensure that item types match.
func (s selectionFilterModel) SetModel(model SelectionModel) {
	var arg0 *C.GtkSelectionFilterModel
	var arg1 *C.GtkSelectionModel

	arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_selection_filter_model_set_model(arg0, arg1)
}
