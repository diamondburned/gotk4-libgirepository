// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// SpacedPrimesClosest gets the smallest prime number from a built-in array of
// primes which is larger than @num. This is used within GLib to calculate the
// optimum size of a Table.
//
// The built-in array of primes ranges from 11 to 13845163 such that each prime
// is approximately 1.5-2 times the previous prime.
func SpacedPrimesClosest(num uint) uint {
	var _arg1 C.guint

	_arg1 = C.guint(num)

	var _cret C.guint

	cret = C.g_spaced_primes_closest(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}