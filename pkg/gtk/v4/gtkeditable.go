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
// extern void _gotk4_gtk4_Editable_ConnectDeleteText(gpointer, gint, gint, guintptr);
// extern void _gotk4_gtk4_Editable_ConnectChanged(gpointer, guintptr);
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

// Editable: GtkEditable is an interface for text editing widgets.
//
// Typical examples of editable widgets are gtk.Entry and gtk.SpinButton. It
// contains functions for generically manipulating an editable widget, a large
// number of action signals used for key bindings, and several signals that an
// application can connect to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// gtk.Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const char  *text,
//                         int          length,
//                         int         *position,
//                         gpointer     data)
//    {
//      char *result = g_utf8_strup (text, length);
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
//    }
//
//
//
// Implementing GtkEditable
//
// The most likely scenario for implementing GtkEditable on your own widget is
// that you will embed a Text inside a complex widget, and want to delegate the
// editable functionality to that text widget. GtkEditable provides some utility
// functions to make this easy.
//
// In your class_init function, call gtk.Editable().InstallProperties, passing
// the first available property ID:
//
//    static void
//    my_class_init (MyClass *class)
//    {
//      ...
//      g_object_class_install_properties (object_class, NUM_PROPERTIES, props);
//      gtk_editable_install_properties (object_clas, NUM_PROPERTIES);
//      ...
//    }
//
//
// In your interface_init function for the GtkEditable interface, provide an
// implementation for the get_delegate vfunc that returns your text widget:
//
//    GtkEditable *
//    get_editable_delegate (GtkEditable *editable)
//    {
//      return GTK_EDITABLE (MY_WIDGET (editable)->text_widget);
//    }
//
//    static void
//    my_editable_init (GtkEditableInterface *iface)
//    {
//      iface->get_delegate = get_editable_delegate;
//    }
//
//
// You don't need to provide any other vfuncs. The default implementations work
// by forwarding to the delegate that the GtkEditableInterface.get_delegate()
// vfunc returns.
//
// In your instance_init function, create your text widget, and then call
// gtk.Editable.InitDelegate():
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      ...
//      self->text_widget = gtk_text_new ();
//      gtk_editable_init_delegate (GTK_EDITABLE (self));
//      ...
//    }
//
//
// In your dispose function, call gtk.Editable.FinishDelegate() before
// destroying your text widget:
//
//    static void
//    my_widget_dispose (GObject *object)
//    {
//      ...
//      gtk_editable_finish_delegate (GTK_EDITABLE (self));
//      g_clear_pointer (&self->text_widget, gtk_widget_unparent);
//      ...
//    }
//
//
// Finally, use gtk.Editable().DelegateSetProperty in your set_property function
// (and similar for get_property), to set the editable properties:
//
//      ...
//      if (gtk_editable_delegate_set_property (object, prop_id, value, pspec))
//        return;
//
//      switch (prop_id)
//      ...
//
//
// It is important to note that if you create a GtkEditable that uses a
// delegate, the low level gtk.Editable::insert-text and
// gtk.Editable::delete-text signals will be propagated from the "wrapper"
// editable to the delegate, but they will not be propagated from the delegate
// to the "wrapper" editable, as they would cause an infinite recursion. If you
// wish to connect to the gtk.Editable::insert-text and
// gtk.Editable::delete-text signals, you will need to connect to them on the
// delegate obtained via gtk.Editable.GetDelegate().
//
// Editable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Editable struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Editable)(nil)
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

// ConnectChanged is emitted at the end of a single user-visible operation on
// the contents.
//
// E.g., a paste operation that replaces the contents of the selection will
// cause only one signal emission (even though it is implemented by first
// deleting the selection, then inserting the new content, and may cause
// multiple ::notify::text signals to be emitted).
func (v *Editable) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk4_Editable_ConnectChanged), f)
}

// ConnectDeleteText is emitted when text is deleted from the widget by the
// user.
//
// The default handler for this signal will normally be responsible for deleting
// the text, so by connecting to this signal and then stopping the signal with
// g_signal_stop_emission(), it is possible to modify the range of deleted text,
// or prevent it from being deleted entirely.
//
// The start_pos and end_pos parameters are interpreted as for
// gtk.Editable.DeleteText().
func (v *Editable) ConnectDeleteText(f func(startPos, endPos int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "delete-text", false, unsafe.Pointer(C._gotk4_gtk4_Editable_ConnectDeleteText), f)
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
