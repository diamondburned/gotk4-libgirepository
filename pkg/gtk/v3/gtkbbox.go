// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_box_get_type()), F: marshalButtonBox},
	})
}

type ButtonBox interface {
	Box
	Buildable
	Orientable

	// ChildNonHomogeneous returns whether the child is exempted from homogenous
	// sizing.
	ChildNonHomogeneous(child Widget) bool
	// ChildSecondary returns whether @child should appear in a secondary group
	// of children.
	ChildSecondary(child Widget) bool
	// Layout retrieves the method being used to arrange the buttons in a button
	// box.
	Layout() ButtonBoxStyle
	// SetChildNonHomogeneous sets whether the child is exempted from homogeous
	// sizing.
	SetChildNonHomogeneous(child Widget, nonHomogeneous bool)
	// SetChildSecondary sets whether @child should appear in a secondary group
	// of children. A typical use of a secondary child is the help button in a
	// dialog.
	//
	// This group appears after the other children if the style is
	// GTK_BUTTONBOX_START, GTK_BUTTONBOX_SPREAD or GTK_BUTTONBOX_EDGE, and
	// before the other children if the style is GTK_BUTTONBOX_END. For
	// horizontal button boxes, the definition of before/after depends on
	// direction of the widget (see gtk_widget_set_direction()). If the style is
	// GTK_BUTTONBOX_START or GTK_BUTTONBOX_END, then the secondary children are
	// aligned at the other end of the button box from the main children. For
	// the other styles, they appear immediately next to the main children.
	SetChildSecondary(child Widget, isSecondary bool)
	// SetLayout changes the way buttons are arranged in their container.
	SetLayout(layoutStyle ButtonBoxStyle)
}

// buttonBox implements the ButtonBox interface.
type buttonBox struct {
	Box
	Buildable
	Orientable
}

var _ ButtonBox = (*buttonBox)(nil)

// WrapButtonBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapButtonBox(obj *externglib.Object) ButtonBox {
	return ButtonBox{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButtonBox(obj), nil
}

// NewButtonBox constructs a class ButtonBox.
func NewButtonBox(orientation Orientation) ButtonBox {
	var _arg1 C.GtkOrientation

	_arg1 = (C.GtkOrientation)(orientation)

	var _cret C.GtkButtonBox

	cret = C.gtk_button_box_new(_arg1)

	var _buttonBox ButtonBox

	_buttonBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ButtonBox)

	return _buttonBox
}

// ChildNonHomogeneous returns whether the child is exempted from homogenous
// sizing.
func (w buttonBox) ChildNonHomogeneous(child Widget) bool {
	var _arg0 *C.GtkButtonBox
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var _cret C.gboolean

	cret = C.gtk_button_box_get_child_non_homogeneous(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ChildSecondary returns whether @child should appear in a secondary group
// of children.
func (w buttonBox) ChildSecondary(child Widget) bool {
	var _arg0 *C.GtkButtonBox
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var _cret C.gboolean

	cret = C.gtk_button_box_get_child_secondary(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Layout retrieves the method being used to arrange the buttons in a button
// box.
func (w buttonBox) Layout() ButtonBoxStyle {
	var _arg0 *C.GtkButtonBox

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))

	var _cret C.GtkButtonBoxStyle

	cret = C.gtk_button_box_get_layout(_arg0)

	var _buttonBoxStyle ButtonBoxStyle

	_buttonBoxStyle = ButtonBoxStyle(_cret)

	return _buttonBoxStyle
}

// SetChildNonHomogeneous sets whether the child is exempted from homogeous
// sizing.
func (w buttonBox) SetChildNonHomogeneous(child Widget, nonHomogeneous bool) {
	var _arg0 *C.GtkButtonBox
	var _arg1 *C.GtkWidget
	var _arg2 C.gboolean

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if nonHomogeneous {
		_arg2 = C.gboolean(1)
	}

	C.gtk_button_box_set_child_non_homogeneous(_arg0, _arg1, _arg2)
}

// SetChildSecondary sets whether @child should appear in a secondary group
// of children. A typical use of a secondary child is the help button in a
// dialog.
//
// This group appears after the other children if the style is
// GTK_BUTTONBOX_START, GTK_BUTTONBOX_SPREAD or GTK_BUTTONBOX_EDGE, and
// before the other children if the style is GTK_BUTTONBOX_END. For
// horizontal button boxes, the definition of before/after depends on
// direction of the widget (see gtk_widget_set_direction()). If the style is
// GTK_BUTTONBOX_START or GTK_BUTTONBOX_END, then the secondary children are
// aligned at the other end of the button box from the main children. For
// the other styles, they appear immediately next to the main children.
func (w buttonBox) SetChildSecondary(child Widget, isSecondary bool) {
	var _arg0 *C.GtkButtonBox
	var _arg1 *C.GtkWidget
	var _arg2 C.gboolean

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if isSecondary {
		_arg2 = C.gboolean(1)
	}

	C.gtk_button_box_set_child_secondary(_arg0, _arg1, _arg2)
}

// SetLayout changes the way buttons are arranged in their container.
func (w buttonBox) SetLayout(layoutStyle ButtonBoxStyle) {
	var _arg0 *C.GtkButtonBox
	var _arg1 C.GtkButtonBoxStyle

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (C.GtkButtonBoxStyle)(layoutStyle)

	C.gtk_button_box_set_layout(_arg0, _arg1)
}