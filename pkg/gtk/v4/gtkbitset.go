// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// BitsetIter: an opaque, stack-allocated struct for iterating over the elements
// of a `GtkBitset`.
//
// Before a `GtkBitsetIter` can be used, it needs to be initialized with
// [func@Gtk.BitsetIter.init_first], [func@Gtk.BitsetIter.init_last] or
// [func@Gtk.BitsetIter.init_at].
type BitsetIter struct {
	native C.GtkBitsetIter
}

// WrapBitsetIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBitsetIter(ptr unsafe.Pointer) *BitsetIter {
	if ptr == nil {
		return nil
	}

	return (*BitsetIter)(ptr)
}

func marshalBitsetIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBitsetIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BitsetIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Value gets the current value that @iter points to.
//
// If @iter is not valid and [method@Gtk.BitsetIter.is_valid] returns false,
// this function returns 0.
func (i *BitsetIter) Value() uint {
	var _arg0 *C.GtkBitsetIter

	_arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var _cret C.guint

	cret = C.gtk_bitset_iter_get_value(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// IsValid checks if @iter points to a valid value.
func (i *BitsetIter) IsValid() bool {
	var _arg0 *C.GtkBitsetIter

	_arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.gtk_bitset_iter_is_valid(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Next moves @iter to the next value in the set.
//
// If it was already pointing to the last value in the set, false is returned
// and @iter is invalidated.
func (i *BitsetIter) Next() (uint, bool) {
	var _arg0 *C.GtkBitsetIter

	_arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var _arg1 C.guint
	var _cret C.gboolean

	cret = C.gtk_bitset_iter_next(_arg0, &_arg1)

	var _value uint
	var _ok bool

	_value = (uint)(_arg1)
	if _cret {
		_ok = true
	}

	return _value, _ok
}

// Previous moves @iter to the previous value in the set.
//
// If it was already pointing to the first value in the set, false is returned
// and @iter is invalidated.
func (i *BitsetIter) Previous() (uint, bool) {
	var _arg0 *C.GtkBitsetIter

	_arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var _arg1 C.guint
	var _cret C.gboolean

	cret = C.gtk_bitset_iter_previous(_arg0, &_arg1)

	var _value uint
	var _ok bool

	_value = (uint)(_arg1)
	if _cret {
		_ok = true
	}

	return _value, _ok
}