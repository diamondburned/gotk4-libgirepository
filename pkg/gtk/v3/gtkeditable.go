// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_Editable_ConnectDeleteText(gpointer, gint, gint, guintptr);
// extern void _gotk4_gtk3_Editable_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeEditable = coreglib.Type(girepository.MustFind("Gtk", "Editable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEditable, F: marshalEditable},
	})
}

// EditableOverrider contains methods that are overridable.
type EditableOverrider interface {
}

// Editable interface is an interface which should be implemented by text
// editing widgets, such as Entry and SpinButton. It contains functions for
// generically manipulating an editable widget, a large number of action signals
// used for key bindings, and several signals that an application can connect to
// to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>;
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const gchar *text,
//                         gint         length,
//                         gint        *position,
//                         gpointer     data)
//    {
//      gchar *result = g_utf8_strup (text, length);
//
//      g_signal_handlers_block_by_func (editable,
//                                   (gpointer) insert_text_handler, data);
//      gtk_editable_insert_text (editable, result, length, position);
//      g_signal_handlers_unblock_by_func (editable,
//                                         (gpointer) insert_text_handler, data);
//
//      g_signal_stop_emission_by_name (editable, "insert_text");
//
//      g_free (result);
//    }.
//
// Editable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Editable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Editable)(nil)
)

// Editabler describes Editable's interface methods.
type Editabler interface {
	coreglib.Objector

	baseEditable() *Editable
}

var _ Editabler = (*Editable)(nil)

func ifaceInitEditabler(gifacePtr, data C.gpointer) {
}

func wrapEditable(obj *coreglib.Object) *Editable {
	return &Editable{
		Object: obj,
	}
}

func marshalEditable(p uintptr) (interface{}, error) {
	return wrapEditable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Editable) baseEditable() *Editable {
	return v
}

// BaseEditable returns the underlying base object.
func BaseEditable(obj Editabler) *Editable {
	return obj.baseEditable()
}

// ConnectChanged signal is emitted at the end of a single user-visible
// operation on the contents of the Editable.
//
// E.g., a paste operation that replaces the contents of the selection will
// cause only one signal emission (even though it is implemented by first
// deleting the selection, then inserting the new content, and may cause
// multiple ::notify::text signals to be emitted).
func (v *Editable) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk3_Editable_ConnectChanged), f)
}

// ConnectDeleteText: this signal is emitted when text is deleted from the
// widget by the user. The default handler for this signal will normally be
// responsible for deleting the text, so by connecting to this signal and then
// stopping the signal with g_signal_stop_emission(), it is possible to modify
// the range of deleted text, or prevent it from being deleted entirely. The
// start_pos and end_pos parameters are interpreted as for
// gtk_editable_delete_text().
func (v *Editable) ConnectDeleteText(f func(startPos, endPos int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "delete-text", false, unsafe.Pointer(C._gotk4_gtk3_Editable_ConnectDeleteText), f)
}

// EditableInterface: instance of this type is always passed by reference.
type EditableInterface struct {
	*editableInterface
}

// editableInterface is the struct that's finalized.
type editableInterface struct {
	native unsafe.Pointer
}

var GIRInfoEditableInterface = girepository.MustFind("Gtk", "EditableInterface")
