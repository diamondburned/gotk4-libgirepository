// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtknumericsorter.go.
var GTypeNumericSorter = externglib.Type(C.gtk_numeric_sorter_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeNumericSorter, F: marshalNumericSorter},
	})
}

// NumericSorterOverrider contains methods that are overridable.
type NumericSorterOverrider interface {
	externglib.Objector
}

// NumericSorter: GtkNumericSorter is a GtkSorter that compares numbers.
//
// To obtain the numbers to compare, this sorter evaluates a gtk.Expression.
type NumericSorter struct {
	_ [0]func() // equal guard
	Sorter
}

var (
	_ externglib.Objector = (*NumericSorter)(nil)
)

func classInitNumericSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNumericSorter(obj *externglib.Object) *NumericSorter {
	return &NumericSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalNumericSorter(p uintptr) (interface{}, error) {
	return wrapNumericSorter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNumericSorter creates a new numeric sorter using the given expression.
//
// Smaller numbers will be sorted first. You can call
// gtk.NumericSorter.SetSortOrder() to change this.
//
// The function takes the following parameters:
//
//    - expression (optional) to evaluate.
//
// The function returns the following values:
//
//    - numericSorter: new GtkNumericSorter.
//
func NewNumericSorter(expression Expressioner) *NumericSorter {
	var _arg1 *C.GtkExpression    // out
	var _cret *C.GtkNumericSorter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(externglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(expression).Native()))
	}

	_cret = C.gtk_numeric_sorter_new(_arg1)
	runtime.KeepAlive(expression)

	var _numericSorter *NumericSorter // out

	_numericSorter = wrapNumericSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _numericSorter
}

// Expression gets the expression that is evaluated to obtain numbers from
// items.
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression, or NULL.
//
func (self *NumericSorter) Expression() Expressioner {
	var _arg0 *C.GtkNumericSorter // out
	var _cret *C.GtkExpression    // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_numeric_sorter_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Expressioner)
				return ok
			})
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// SortOrder gets whether this sorter will sort smaller numbers first.
//
// The function returns the following values:
//
//    - sortType: order of the numbers.
//
func (self *NumericSorter) SortOrder() SortType {
	var _arg0 *C.GtkNumericSorter // out
	var _cret C.GtkSortType       // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_numeric_sorter_get_sort_order(_arg0)
	runtime.KeepAlive(self)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

// SetExpression sets the expression that is evaluated to obtain numbers from
// items.
//
// Unless an expression is set on self, the sorter will always compare items as
// invalid.
//
// The expression must have a return type that can be compared numerically, such
// as G_TYPE_INT or G_TYPE_DOUBLE.
//
// The function takes the following parameters:
//
//    - expression (optional): GtkExpression, or NULL.
//
func (self *NumericSorter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 *C.GtkExpression    // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(externglib.InternObject(expression).Native()))
	}

	C.gtk_numeric_sorter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetSortOrder sets whether to sort smaller numbers before larger ones.
//
// The function takes the following parameters:
//
//    - sortOrder: whether to sort smaller numbers first.
//
func (self *NumericSorter) SetSortOrder(sortOrder SortType) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 C.GtkSortType       // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.GtkSortType(sortOrder)

	C.gtk_numeric_sorter_set_sort_order(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sortOrder)
}
