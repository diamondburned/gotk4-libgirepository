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
		{T: externglib.Type(C.gtk_slice_list_model_get_type()), F: marshalSliceListModel},
	})
}

// SliceListModel: `GtkSliceListModel` is a list model that presents a slice of
// another model.
//
// This is useful when implementing paging by setting the size to the number of
// elements per page and updating the offset whenever a different page is
// opened.
type SliceListModel interface {
	gextras.Objector
	gio.ListModel

	// Offset gets the offset set via gtk_slice_list_model_set_offset().
	Offset() uint
	// Size gets the size set via gtk_slice_list_model_set_size().
	Size() uint
	// SetModel sets the model to show a slice of.
	//
	// The model's item type must conform to @self's item type.
	SetModel(model gio.ListModel)
	// SetOffset sets the offset into the original model for this slice.
	//
	// If the offset is too large for the sliced model, @self will end up empty.
	SetOffset(offset uint)
	// SetSize sets the maximum size. @self will never have more items than
	// @size.
	//
	// It can however have fewer items if the offset is too large or the model
	// sliced from doesn't have enough items.
	SetSize(size uint)
}

// sliceListModel implements the SliceListModel class.
type sliceListModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SliceListModel = (*sliceListModel)(nil)

// WrapSliceListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSliceListModel(obj *externglib.Object) SliceListModel {
	return sliceListModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSliceListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSliceListModel(obj), nil
}

// Offset gets the offset set via gtk_slice_list_model_set_offset().
func (s sliceListModel) Offset() uint {
	var _arg0 *C.GtkSliceListModel // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_slice_list_model_get_offset(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Size gets the size set via gtk_slice_list_model_set_size().
func (s sliceListModel) Size() uint {
	var _arg0 *C.GtkSliceListModel // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_slice_list_model_get_size(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// SetModel sets the model to show a slice of.
//
// The model's item type must conform to @self's item type.
func (s sliceListModel) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_slice_list_model_set_model(_arg0, _arg1)
}

// SetOffset sets the offset into the original model for this slice.
//
// If the offset is too large for the sliced model, @self will end up empty.
func (s sliceListModel) SetOffset(offset uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(offset)

	C.gtk_slice_list_model_set_offset(_arg0, _arg1)
}

// SetSize sets the maximum size. @self will never have more items than
// @size.
//
// It can however have fewer items if the offset is too large or the model
// sliced from doesn't have enough items.
func (s sliceListModel) SetSize(size uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(size)

	C.gtk_slice_list_model_set_size(_arg0, _arg1)
}
