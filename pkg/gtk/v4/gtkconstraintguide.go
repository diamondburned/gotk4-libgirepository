// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_constraint_guide_get_type()), F: marshalConstraintGuide},
	})
}

// ConstraintGuide: `GtkConstraintGuide` is an invisible layout element in a
// `GtkConstraintLayout`.
//
// The `GtkConstraintLayout` treats guides like widgets. They can be used as the
// source or target of a `GtkConstraint`.
//
// Guides have a minimum, maximum and natural size. Depending on the constraints
// that are applied, they can act like a guideline that widgets can be aligned
// to, or like *flexible space*.
//
// Unlike a `GtkWidget`, a `GtkConstraintGuide` will not be drawn.
type ConstraintGuide interface {
	gextras.Objector

	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget

	// MaxSize gets the maximum size of @guide.
	MaxSize(width *int, height *int)
	// MinSize gets the minimum size of @guide.
	MinSize(width *int, height *int)
	// Name retrieves the name set using gtk_constraint_guide_set_name().
	Name() string
	// NatSize gets the natural size of @guide.
	NatSize(width *int, height *int)
	// Strength retrieves the strength set using
	// gtk_constraint_guide_set_strength().
	Strength() ConstraintStrength
	// SetMaxSize sets the maximum size of @guide.
	//
	// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
	// updated to reflect the new size.
	SetMaxSize(width int, height int)
	// SetMinSize sets the minimum size of @guide.
	//
	// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
	// updated to reflect the new size.
	SetMinSize(width int, height int)
	// SetName sets a name for the given `GtkConstraintGuide`.
	//
	// The name is useful for debugging purposes.
	SetName(name string)
	// SetNatSize sets the natural size of @guide.
	//
	// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
	// updated to reflect the new size.
	SetNatSize(width int, height int)
	// SetStrength sets the strength of the constraint on the natural size of
	// the given `GtkConstraintGuide`.
	SetStrength(strength ConstraintStrength)
}

// constraintGuide implements the ConstraintGuide interface.
type constraintGuide struct {
	*externglib.Object
}

var _ ConstraintGuide = (*constraintGuide)(nil)

// WrapConstraintGuide wraps a GObject to a type that implements
// interface ConstraintGuide. It is primarily used internally.
func WrapConstraintGuide(obj *externglib.Object) ConstraintGuide {
	return constraintGuide{obj}
}

func marshalConstraintGuide(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConstraintGuide(obj), nil
}

// NewConstraintGuide creates a new `GtkConstraintGuide` object.
func NewConstraintGuide() ConstraintGuide {
	var _cret *C.GtkConstraintGuide // in

	_cret = C.gtk_constraint_guide_new()

	var _constraintGuide ConstraintGuide // out

	_constraintGuide = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(ConstraintGuide)

	return _constraintGuide
}

func (c constraintGuide) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(c))
}

func (g constraintGuide) MaxSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_max_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) MinSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_min_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) Name() string {
	var _arg0 *C.GtkConstraintGuide // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_constraint_guide_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (g constraintGuide) NatSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_nat_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) Strength() ConstraintStrength {
	var _arg0 *C.GtkConstraintGuide   // out
	var _cret C.GtkConstraintStrength // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_constraint_guide_get_strength(_arg0)

	var _constraintStrength ConstraintStrength // out

	_constraintStrength = ConstraintStrength(_cret)

	return _constraintStrength
}

func (g constraintGuide) SetMaxSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_max_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) SetMinSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_min_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) SetName(name string) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_constraint_guide_set_name(_arg0, _arg1)
}

func (g constraintGuide) SetNatSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_nat_size(_arg0, _arg1, _arg2)
}

func (g constraintGuide) SetStrength(strength ConstraintStrength) {
	var _arg0 *C.GtkConstraintGuide   // out
	var _arg1 C.GtkConstraintStrength // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(g.Native()))
	_arg1 = C.GtkConstraintStrength(strength)

	C.gtk_constraint_guide_set_strength(_arg0, _arg1)
}
