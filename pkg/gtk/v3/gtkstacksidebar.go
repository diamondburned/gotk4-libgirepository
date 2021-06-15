// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebar},
	})
}

// StackSidebar: a GtkStackSidebar enables you to quickly and easily provide a
// consistent "sidebar" object for your user interface.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk_stack_sidebar_set_stack() to connect the StackSidebar to the Stack.
//
//
// CSS nodes
//
// GtkStackSidebar has a single CSS node with name stacksidebar and style class
// .sidebar.
//
// When circumstances require it, GtkStackSidebar adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar interface {
	Bin
	Buildable

	// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
	Stack() Stack
	// SetStack: set the Stack associated with this StackSidebar.
	//
	// The sidebar widget will automatically update according to the order
	// (packing) and items within the given Stack.
	SetStack(stack Stack)
}

// stackSidebar implements the StackSidebar class.
type stackSidebar struct {
	Bin
	Buildable
}

var _ StackSidebar = (*stackSidebar)(nil)

// WrapStackSidebar wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackSidebar(obj *externglib.Object) StackSidebar {
	return stackSidebar{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
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

// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
func (s stackSidebar) Stack() Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)

	var _stack Stack // out

	_stack = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Stack)

	return _stack
}

// SetStack: set the Stack associated with this StackSidebar.
//
// The sidebar widget will automatically update according to the order
// (packing) and items within the given Stack.
func (s stackSidebar) SetStack(stack Stack) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
}
