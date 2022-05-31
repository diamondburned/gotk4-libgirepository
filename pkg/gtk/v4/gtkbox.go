// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkbox.go.
var GTypeBox = coreglib.Type(C.gtk_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBox, F: marshalBox},
	})
}

// BoxOverrider contains methods that are overridable.
type BoxOverrider interface {
}

// Box: GtkBox widget arranges child widgets into a single row or column.
//
// !An example GtkBox (box.png)
//
// Whether it is a row or column depends on the value of its
// gtk.Orientable:orientation property. Within the other dimension, all children
// are allocated the same size. Of course, the gtk.Widget:halign and
// gtk.Widget:valign properties can be used on the children to influence their
// allocation.
//
// Use repeated calls to gtk.Box.Append() to pack widgets into a GtkBox from
// start to end. Use gtk.Box.Remove() to remove widgets from the GtkBox.
// gtk.Box.InsertChildAfter() can be used to add a child at a particular
// position.
//
// Use gtk.Box.SetHomogeneous() to specify whether or not all children of the
// GtkBox are forced to get the same amount of space.
//
// Use gtk.Box.SetSpacing() to determine how much space will be minimally placed
// between all children in the GtkBox. Note that spacing is added *between* the
// children.
//
// Use gtk.Box.ReorderChildAfter() to move a child to a different place in the
// box.
//
//
// CSS nodes
//
// GtkBox uses a single CSS node with name box.
//
//
// Accessibility
//
// GtkBox uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Box struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Box)(nil)
	_ coreglib.Objector = (*Box)(nil)
)

func classInitBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBox(obj *coreglib.Object) *Box {
	return &Box{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalBox(p uintptr) (interface{}, error) {
	return wrapBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Append adds child as the last child to box.
//
// The function takes the following parameters:
//
//    - child: GtkWidget to append.
//
func (box *Box) Append(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Box").InvokeMethod("append", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// Homogeneous returns whether the box is homogeneous (all children are the same
// size).
//
// The function returns the following values:
//
//    - ok: TRUE if the box is homogeneous.
//
func (box *Box) Homogeneous() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Box").InvokeMethod("get_homogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_box_set_spacing().
//
// The function returns the following values:
//
//    - gint: spacing between children.
//
func (box *Box) Spacing() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Box").InvokeMethod("get_spacing", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// InsertChildAfter inserts child in the position after sibling in the list of
// box children.
//
// If sibling is NULL, insert child at the first position.
//
// The function takes the following parameters:
//
//    - child: GtkWidget to insert.
//    - sibling (optional) after which to insert child.
//
func (box *Box) InsertChildAfter(child, sibling Widgetter) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if sibling != nil {
		_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sibling).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Box").InvokeMethod("insert_child_after", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// Prepend adds child as the first child to box.
//
// The function takes the following parameters:
//
//    - child: GtkWidget to prepend.
//
func (box *Box) Prepend(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Box").InvokeMethod("prepend", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// Remove removes a child widget from box.
//
// The child must have been added before with gtk.Box.Append(),
// gtk.Box.Prepend(), or gtk.Box.InsertChildAfter().
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (box *Box) Remove(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Box").InvokeMethod("remove", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// box children.
//
// If sibling is NULL, move child to the first position.
//
// The function takes the following parameters:
//
//    - child: GtkWidget to move, must be a child of box.
//    - sibling (optional) to move child after, or NULL.
//
func (box *Box) ReorderChildAfter(child, sibling Widgetter) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if sibling != nil {
		_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sibling).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Box").InvokeMethod("reorder_child_after", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// SetHomogeneous sets whether or not all children of box are given equal space
// in the box.
//
// The function takes the following parameters:
//
//    - homogeneous: boolean value, TRUE to create equal allotments, FALSE for
//      variable allotments.
//
func (box *Box) SetHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Box").InvokeMethod("set_homogeneous", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(homogeneous)
}

// SetSpacing sets the number of pixels to place between children of box.
//
// The function takes the following parameters:
//
//    - spacing: number of pixels to put between children.
//
func (box *Box) SetSpacing(spacing int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = C.int(spacing)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Box").InvokeMethod("set_spacing", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
