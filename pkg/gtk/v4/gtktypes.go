// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtktypes.go.
var GTypeBitset = coreglib.Type(C.gtk_bitset_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBitset, F: marshalBitset},
	})
}

// Bitset: GtkBitset represents a set of unsigned integers.
//
// Another name for this data structure is "bitmap".
//
// The current implementation is based on roaring bitmaps
// (https://roaringbitmap.org/).
//
// A bitset allows adding a set of integers and provides support for set
// operations like unions, intersections and checks for equality or if a value
// is contained in the set. GtkBitset also contains various functions to query
// metadata about the bitset, such as the minimum or maximum values or its size.
//
// The fastest way to iterate values in a bitset is gtk.BitsetIter.
//
// The main use case for GtkBitset is implementing complex selections for
// gtk.SelectionModel.
//
// An instance of this type is always passed by reference.
type Bitset struct {
	*bitset
}

// bitset is the struct that's finalized.
type bitset struct {
	native unsafe.Pointer
}

func marshalBitset(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Bitset{&bitset{(unsafe.Pointer)(b)}}, nil
}

// NewBitsetEmpty constructs a struct Bitset.
func NewBitsetEmpty() *Bitset {
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _bitset *Bitset // out

	_bitset = (*Bitset)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bitset)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_bitset_unref((*C.GtkBitset)(intern.C))
		},
	)

	return _bitset
}

// NewBitsetRange constructs a struct Bitset.
func NewBitsetRange(start uint32, nItems uint32) *Bitset {
	var _args [2]girepository.Argument

	*(*C.guint)(unsafe.Pointer(&_args[0])) = C.guint(start)
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(nItems)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(start)
	runtime.KeepAlive(nItems)

	var _bitset *Bitset // out

	_bitset = (*Bitset)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bitset)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_bitset_unref((*C.GtkBitset)(intern.C))
		},
	)

	return _bitset
}

// Add adds value to self if it wasn't part of it before.
//
// The function takes the following parameters:
//
//    - value to add.
//
// The function returns the following values:
//
//    - ok: TRUE if value was not part of self and self was changed.
//
func (self *Bitset) Add(value uint32) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(value)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// AddRange adds all values from start (inclusive) to start + n_items
// (exclusive) in self.
//
// The function takes the following parameters:
//
//    - start: first value to add.
//    - nItems: number of consecutive values to add.
//
func (self *Bitset) AddRange(start uint32, nItems uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(start)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(nItems)

	runtime.KeepAlive(self)
	runtime.KeepAlive(start)
	runtime.KeepAlive(nItems)
}

// AddRangeClosed adds the closed range [first, last], so first, last and all
// values in between. first must be smaller than last.
//
// The function takes the following parameters:
//
//    - first value to add.
//    - last value to add.
//
func (self *Bitset) AddRangeClosed(first uint32, last uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(first)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(last)

	runtime.KeepAlive(self)
	runtime.KeepAlive(first)
	runtime.KeepAlive(last)
}

// AddRectangle interprets the values as a 2-dimensional boolean grid with the
// given stride and inside that grid, adds a rectangle with the given width and
// height.
//
// The function takes the following parameters:
//
//    - start: first value to add.
//    - width of the rectangle.
//    - height of the rectangle.
//    - stride: row stride of the grid.
//
func (self *Bitset) AddRectangle(start uint32, width uint32, height uint32, stride uint32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(start)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(width)
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(height)
	*(*C.guint)(unsafe.Pointer(&_args[4])) = C.guint(stride)

	runtime.KeepAlive(self)
	runtime.KeepAlive(start)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(stride)
}

// Contains checks if the given value has been added to self.
//
// The function takes the following parameters:
//
//    - value to check.
//
// The function returns the following values:
//
//    - ok: TRUE if self contains value.
//
func (self *Bitset) Contains(value uint32) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(value)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Copy creates a copy of self.
//
// The function returns the following values:
//
//    - bitset: new bitset that contains the same values as self.
//
func (self *Bitset) Copy() *Bitset {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _bitset *Bitset // out

	_bitset = (*Bitset)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bitset)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_bitset_unref((*C.GtkBitset)(intern.C))
		},
	)

	return _bitset
}

// Difference sets self to be the symmetric difference of self and other.
//
// The symmetric difference is set self to contain all values that were either
// contained in self or in other, but not in both. This operation is also called
// an XOR.
//
// It is allowed for self and other to be the same bitset. The bitset will be
// emptied in that case.
//
// The function takes the following parameters:
//
//    - other: GtkBitset to compute the difference from.
//
func (self *Bitset) Difference(other *Bitset) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	runtime.KeepAlive(self)
	runtime.KeepAlive(other)
}

// Equals returns TRUE if self and other contain the same values.
//
// The function takes the following parameters:
//
//    - other GtkBitset.
//
// The function returns the following values:
//
//    - ok: TRUE if self and other contain the same values.
//
func (self *Bitset) Equals(other *Bitset) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(other)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Maximum returns the largest value in self.
//
// If self is empty, 0 is returned.
//
// The function returns the following values:
//
//    - guint: largest value in self.
//
func (self *Bitset) Maximum() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Minimum returns the smallest value in self.
//
// If self is empty, G_MAXUINT is returned.
//
// The function returns the following values:
//
//    - guint: smallest value in self.
//
func (self *Bitset) Minimum() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Nth returns the value of the nth item in self.
//
// If nth is >= the size of self, 0 is returned.
//
// The function takes the following parameters:
//
//    - nth: index of the item to get.
//
// The function returns the following values:
//
//    - guint: value of the nth item in self.
//
func (self *Bitset) Nth(nth uint32) uint32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(nth)

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(nth)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Size gets the number of values that were added to the set.
//
// For example, if the set is empty, 0 is returned.
//
// Note that this function returns a guint64, because when all values are set,
// the return value is G_MAXUINT + 1. Unless you are sure this cannot happen (it
// can't with GListModel), be sure to use a 64bit type.
//
// The function returns the following values:
//
//    - guint64: number of values in the set.
//
func (self *Bitset) Size() uint64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = *(*C.guint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(*(*C.guint64)(unsafe.Pointer(&_cret)))

	return _guint64
}

// SizeInRange gets the number of values that are part of the set from first to
// last (inclusive).
//
// Note that this function returns a guint64, because when all values are set,
// the return value is G_MAXUINT + 1. Unless you are sure this cannot happen (it
// can't with GListModel), be sure to use a 64bit type.
//
// The function takes the following parameters:
//
//    - first element to include.
//    - last element to include.
//
// The function returns the following values:
//
//    - guint64: number of values in the set from first to last.
//
func (self *Bitset) SizeInRange(first uint32, last uint32) uint64 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(first)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(last)

	_cret = *(*C.guint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(first)
	runtime.KeepAlive(last)

	var _guint64 uint64 // out

	_guint64 = uint64(*(*C.guint64)(unsafe.Pointer(&_cret)))

	return _guint64
}

// Intersect sets self to be the intersection of self and other.
//
// In other words, remove all values from self that are not part of other.
//
// It is allowed for self and other to be the same bitset. Nothing will happen
// in that case.
//
// The function takes the following parameters:
//
//    - other: GtkBitset to intersect with.
//
func (self *Bitset) Intersect(other *Bitset) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	runtime.KeepAlive(self)
	runtime.KeepAlive(other)
}

// IsEmpty: check if no value is contained in bitset.
//
// The function returns the following values:
//
//    - ok: TRUE if self is empty.
//
func (self *Bitset) IsEmpty() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes value from self if it was part of it before.
//
// The function takes the following parameters:
//
//    - value to add.
//
// The function returns the following values:
//
//    - ok: TRUE if value was part of self and self was changed.
//
func (self *Bitset) Remove(value uint32) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(value)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RemoveAll removes all values from the bitset so that it is empty again.
func (self *Bitset) RemoveAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))

	runtime.KeepAlive(self)
}

// RemoveRange removes all values from start (inclusive) to start + n_items
// (exclusive) in self.
//
// The function takes the following parameters:
//
//    - start: first value to remove.
//    - nItems: number of consecutive values to remove.
//
func (self *Bitset) RemoveRange(start uint32, nItems uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(start)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(nItems)

	runtime.KeepAlive(self)
	runtime.KeepAlive(start)
	runtime.KeepAlive(nItems)
}

// RemoveRangeClosed removes the closed range [first, last], so first, last and
// all values in between. first must be smaller than last.
//
// The function takes the following parameters:
//
//    - first value to remove.
//    - last value to remove.
//
func (self *Bitset) RemoveRangeClosed(first uint32, last uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(first)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(last)

	runtime.KeepAlive(self)
	runtime.KeepAlive(first)
	runtime.KeepAlive(last)
}

// RemoveRectangle interprets the values as a 2-dimensional boolean grid with
// the given stride and inside that grid, removes a rectangle with the given
// width and height.
//
// The function takes the following parameters:
//
//    - start: first value to remove.
//    - width of the rectangle.
//    - height of the rectangle.
//    - stride: row stride of the grid.
//
func (self *Bitset) RemoveRectangle(start uint32, width uint32, height uint32, stride uint32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(start)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(width)
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(height)
	*(*C.guint)(unsafe.Pointer(&_args[4])) = C.guint(stride)

	runtime.KeepAlive(self)
	runtime.KeepAlive(start)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(stride)
}

// ShiftLeft shifts all values in self to the left by amount.
//
// Values smaller than amount are discarded.
//
// The function takes the following parameters:
//
//    - amount to shift all values to the left.
//
func (self *Bitset) ShiftLeft(amount uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(amount)

	runtime.KeepAlive(self)
	runtime.KeepAlive(amount)
}

// ShiftRight shifts all values in self to the right by amount.
//
// Values that end up too large to be held in a #guint are discarded.
//
// The function takes the following parameters:
//
//    - amount to shift all values to the right.
//
func (self *Bitset) ShiftRight(amount uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(amount)

	runtime.KeepAlive(self)
	runtime.KeepAlive(amount)
}

// Splice: this is a support function for GListModel handling, by mirroring the
// GlistModel::items-changed signal.
//
// First, it "cuts" the values from position to removed from the bitset. That
// is, it removes all those values and shifts all larger values to the left by
// removed places.
//
// Then, it "pastes" new room into the bitset by shifting all values larger than
// position by added spaces to the right. This frees up space that can then be
// filled.
//
// The function takes the following parameters:
//
//    - position at which to slice.
//    - removed: number of values to remove.
//    - added: number of values to add.
//
func (self *Bitset) Splice(position uint32, removed uint32, added uint32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(removed)
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(added)

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
	runtime.KeepAlive(removed)
	runtime.KeepAlive(added)
}

// Subtract sets self to be the subtraction of other from self.
//
// In other words, remove all values from self that are part of other.
//
// It is allowed for self and other to be the same bitset. The bitset will be
// emptied in that case.
//
// The function takes the following parameters:
//
//    - other: GtkBitset to subtract.
//
func (self *Bitset) Subtract(other *Bitset) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	runtime.KeepAlive(self)
	runtime.KeepAlive(other)
}

// Union sets self to be the union of self and other.
//
// That is, add all values from other into self that weren't part of it.
//
// It is allowed for self and other to be the same bitset. Nothing will happen
// in that case.
//
// The function takes the following parameters:
//
//    - other: GtkBitset to union with.
//
func (self *Bitset) Union(other *Bitset) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	runtime.KeepAlive(self)
	runtime.KeepAlive(other)
}
