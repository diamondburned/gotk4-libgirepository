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
// extern gboolean _gotk4_gtk4_Paned_ConnectToggleHandleFocus(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCycleHandleFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCycleChildFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCancelPosition(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectAcceptPosition(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypePaned = coreglib.Type(girepository.MustFind("Gtk", "Paned").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePaned, F: marshalPaned},
	})
}

// Paned: GtkPaned has two panes, arranged either horizontally or vertically.
//
// !An example GtkPaned (panes.png)
//
// The division between the two panes is adjustable by the user by dragging a
// handle.
//
// Child widgets are added to the panes of the widget with
// gtk.Paned.SetStartChild() and gtk.Paned.SetEndChild(). The division between
// the two children is set by default from the size requests of the children,
// but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a gtk.Frame so that the gutter appears as a ridge. No separator is drawn if
// one of the children is missing.
//
// Each child has two options that can be set, resize and shrink. If resize is
// true, then when the GtkPaned is resized, that child will expand or shrink
// along with the paned widget. If shrink is true, then that child can be made
// smaller than its requisition by the user. Setting shrink to FALSE allows the
// application to set a minimum size. If resize is false for both children, then
// this is treated as if resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling gtk.Paned.SetPosition().
//
// CSS nodes
//
//    paned
//    ├── <child>
//    ├── separator[.wide]
//    ╰── <child>
//
//
// GtkPaned has a main CSS node with name paned, and a subnode for the separator
// with name separator. The subnode gets a .wide style class when the paned is
// supposed to be wide.
//
// In horizontal orientation, the nodes are arranged based on the text
// direction, so in left-to-right mode, :first-child will select the leftmost
// child, while it will select the rightmost child in RTL layouts.
//
// Creating a paned widget with minimum sizes.
//
//    GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
//    GtkWidget *frame1 = gtk_frame_new (NULL);
//    GtkWidget *frame2 = gtk_frame_new (NULL);
//
//    gtk_widget_set_size_request (hpaned, 200, -1);
//
//    gtk_paned_set_start_child (GTK_PANED (hpaned), frame1);
//    gtk_paned_set_start_child_resize (GTK_PANED (hpaned), TRUE);
//    gtk_paned_set_start_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame1, 50, -1);
//
//    gtk_paned_set_end_child (GTK_PANED (hpaned), frame2);
//    gtk_paned_set_end_child_resize (GTK_PANED (hpaned), FALSE);
//    gtk_paned_set_end_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame2, 50, -1);.
type Paned struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Paned)(nil)
	_ coreglib.Objector = (*Paned)(nil)
)

func wrapPaned(obj *coreglib.Object) *Paned {
	return &Paned{
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

func marshalPaned(p uintptr) (interface{}, error) {
	return wrapPaned(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAcceptPosition is emitted to accept the current position of the handle
// when moving it using key bindings.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Return or Space.
func (v *Paned) ConnectAcceptPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "accept-position", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectAcceptPosition), f)
}

// ConnectCancelPosition is emitted to cancel moving the position of the handle
// using key bindings.
//
// The position of the handle will be reset to the value prior to moving it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Escape.
func (v *Paned) ConnectCancelPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "cancel-position", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCancelPosition), f)
}

// ConnectCycleChildFocus is emitted to cycle the focus between the children of
// the paned.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding is F6.
func (v *Paned) ConnectCycleChildFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "cycle-child-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCycleChildFocus), f)
}

// ConnectCycleHandleFocus is emitted to cycle whether the paned should grab
// focus to allow the user to change position of the handle by using key
// bindings.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is F8.
func (v *Paned) ConnectCycleHandleFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "cycle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCycleHandleFocus), f)
}

// ConnectToggleHandleFocus is emitted to accept the current position of the
// handle and then move focus to the next widget in the focus chain.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding is Tab.
func (v *Paned) ConnectToggleHandleFocus(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "toggle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectToggleHandleFocus), f)
}
