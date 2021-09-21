// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_string_sorter_get_type()), F: marshalStringSorterer},
	})
}

// StringSorter: GtkStringSorter is a GtkSorter that compares strings.
//
// It does the comparison in a linguistically correct way using the current
// locale by normalizing Unicode strings and possibly case-folding them before
// performing the comparison.
//
// To obtain the strings to compare, this sorter evaluates a gtk.Expression.
type StringSorter struct {
	Sorter
}

func wrapStringSorter(obj *externglib.Object) *StringSorter {
	return &StringSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalStringSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStringSorter(obj), nil
}

// NewStringSorter creates a new string sorter that compares items using the
// given expression.
//
// Unless an expression is set on it, this sorter will always compare items as
// invalid.
func NewStringSorter(expression Expressioner) *StringSorter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringSorter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}

	_cret = C.gtk_string_sorter_new(_arg1)
	runtime.KeepAlive(expression)

	var _stringSorter *StringSorter // out

	_stringSorter = wrapStringSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringSorter
}

// Expression gets the expression that is evaluated to obtain strings from
// items.
func (self *StringSorter) Expression() Expressioner {
	var _arg0 *C.GtkStringSorter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_sorter_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Expressioner)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// IgnoreCase gets whether the sorter ignores case differences.
func (self *StringSorter) IgnoreCase() bool {
	var _arg0 *C.GtkStringSorter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_sorter_get_ignore_case(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpression sets the expression that is evaluated to obtain strings from
// items.
//
// The expression must have the type G_TYPE_STRING.
func (self *StringSorter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
	}

	C.gtk_string_sorter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetIgnoreCase sets whether the sorter will ignore case differences.
func (self *StringSorter) SetIgnoreCase(ignoreCase bool) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_sorter_set_ignore_case(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ignoreCase)
}
