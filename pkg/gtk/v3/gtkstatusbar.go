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
		{T: externglib.Type(C.gtk_statusbar_get_type()), F: marshalStatusbar},
	})
}

// Statusbar: a Statusbar is usually placed along the bottom of an application's
// main Window. It may provide a regular commentary of the application's status
// (as is usually the case in a web browser, for example), or may be used to
// simply output a message when the status changes, (when an upload is complete
// in an FTP client, for example).
//
// Status bars in GTK+ maintain a stack of messages. The message at the top of
// the each bar’s stack is the one that will currently be displayed.
//
// Any messages added to a statusbar’s stack must specify a context id that is
// used to uniquely identify the source of a message. This context id can be
// generated by gtk_statusbar_get_context_id(), given a message and the
// statusbar that it will be added to. Note that messages are stored in a stack,
// and when choosing which message to display, the stack structure is adhered
// to, regardless of the context identifier of a message.
//
// One could say that a statusbar maintains one stack of messages for display
// purposes, but allows multiple message producers to maintain sub-stacks of the
// messages they produced (via context ids).
//
// Status bars are created using gtk_statusbar_new().
//
// Messages are added to the bar’s stack with gtk_statusbar_push().
//
// The message at the top of the stack can be removed using gtk_statusbar_pop().
// A message can be removed from anywhere in the stack if its message id was
// recorded at the time it was added. This is done using gtk_statusbar_remove().
//
//
// CSS node
//
// GtkStatusbar has a single CSS node with name statusbar.
type Statusbar interface {
	Box
	Buildable
	Orientable

	// ContextID returns a new context identifier, given a description of the
	// actual context. Note that the description is not shown in the UI.
	ContextID(contextDescription string) uint
	// MessageArea retrieves the box containing the label widget.
	MessageArea() Box
	// Pop removes the first message in the Statusbar’s stack with the given
	// context id.
	//
	// Note that this may not change the displayed message, if the message at
	// the top of the stack has a different context id.
	Pop(contextID uint)
	// Push pushes a new message onto a statusbar’s stack.
	Push(contextID uint, text string) uint
	// Remove forces the removal of a message from a statusbar’s stack. The
	// exact @context_id and @message_id must be specified.
	Remove(contextID uint, messageID uint)
	// RemoveAll forces the removal of all messages from a statusbar's stack
	// with the exact @context_id.
	RemoveAll(contextID uint)
}

// statusbar implements the Statusbar interface.
type statusbar struct {
	Box
	Buildable
	Orientable
}

var _ Statusbar = (*statusbar)(nil)

// WrapStatusbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapStatusbar(obj *externglib.Object) Statusbar {
	return Statusbar{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalStatusbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusbar(obj), nil
}

// NewStatusbar constructs a class Statusbar.
func NewStatusbar() Statusbar {
	var cret C.GtkStatusbar
	var goret Statusbar

	cret = C.gtk_statusbar_new()

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Statusbar)

	return goret
}

// ContextID returns a new context identifier, given a description of the
// actual context. Note that the description is not shown in the UI.
func (s statusbar) ContextID(contextDescription string) uint {
	var arg0 *C.GtkStatusbar
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(contextDescription))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.guint
	var goret uint

	cret = C.gtk_statusbar_get_context_id(arg0, arg1)

	goret = uint(cret)

	return goret
}

// MessageArea retrieves the box containing the label widget.
func (s statusbar) MessageArea() Box {
	var arg0 *C.GtkStatusbar

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))

	var cret *C.GtkWidget
	var goret Box

	cret = C.gtk_statusbar_get_message_area(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Box)

	return goret
}

// Pop removes the first message in the Statusbar’s stack with the given
// context id.
//
// Note that this may not change the displayed message, if the message at
// the top of the stack has a different context id.
func (s statusbar) Pop(contextID uint) {
	var arg0 *C.GtkStatusbar
	var arg1 C.guint

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(contextID)

	C.gtk_statusbar_pop(arg0, arg1)
}

// Push pushes a new message onto a statusbar’s stack.
func (s statusbar) Push(contextID uint, text string) uint {
	var arg0 *C.GtkStatusbar
	var arg1 C.guint
	var arg2 *C.gchar

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(contextID)
	arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.guint
	var goret uint

	cret = C.gtk_statusbar_push(arg0, arg1, arg2)

	goret = uint(cret)

	return goret
}

// Remove forces the removal of a message from a statusbar’s stack. The
// exact @context_id and @message_id must be specified.
func (s statusbar) Remove(contextID uint, messageID uint) {
	var arg0 *C.GtkStatusbar
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(contextID)
	arg2 = C.guint(messageID)

	C.gtk_statusbar_remove(arg0, arg1, arg2)
}

// RemoveAll forces the removal of all messages from a statusbar's stack
// with the exact @context_id.
func (s statusbar) RemoveAll(contextID uint) {
	var arg0 *C.GtkStatusbar
	var arg1 C.guint

	arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(contextID)

	C.gtk_statusbar_remove_all(arg0, arg1)
}
