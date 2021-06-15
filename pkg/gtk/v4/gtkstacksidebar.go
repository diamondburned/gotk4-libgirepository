// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebar},
	})
}

// StackSidebar: a `GtkStackSidebar` uses a sidebar to switch between `GtkStack`
// pages.
//
// In order to use a `GtkStackSidebar`, you simply use a `GtkStack` to organize
// your UI flow, and add the sidebar to your sidebar area. You can use
// [method@Gtk.StackSidebar.set_stack] to connect the `GtkStackSidebar` to the
// `GtkStack`.
//
//
// CSS nodes
//
// `GtkStackSidebar` has a single CSS node with name stacksidebar and style
// class .sidebar.
//
// When circumstances require it, `GtkStackSidebar` adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Stack retrieves the stack.
	Stack() Stack
	// SetStack: set the `GtkStack` associated with this `GtkStackSidebar`.
	//
	// The sidebar widget will automatically update according to the order and
	// items within the given `GtkStack`.
	SetStack(stack Stack)
}

// stackSidebar implements the StackSidebar class.
type stackSidebar struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ StackSidebar = (*stackSidebar)(nil)

// WrapStackSidebar wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackSidebar(obj *externglib.Object) StackSidebar {
	return stackSidebar{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalStackSidebar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStackSidebar(obj), nil
}

// NewStackSidebar constructs a class StackSidebar.
func NewStackSidebar() StackSidebar {
	var _cret C.GtkStackSidebar // in

	_cret = C.gtk_stack_sidebar_new()

	var _stackSidebar StackSidebar // out

	_stackSidebar = WrapStackSidebar(externglib.Take(unsafe.Pointer(_cret)))

	return _stackSidebar
}

// Stack retrieves the stack.
func (s stackSidebar) Stack() Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)

	var _stack Stack // out

	_stack = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Stack)

	return _stack
}

// SetStack: set the `GtkStack` associated with this `GtkStackSidebar`.
//
// The sidebar widget will automatically update according to the order and
// items within the given `GtkStack`.
func (s stackSidebar) SetStack(stack Stack) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
}
