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
// extern void _gotk4_gtk4_Expander_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeExpander = coreglib.Type(girepository.MustFind("Gtk", "Expander").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeExpander, F: marshalExpander},
	})
}

// Expander: GtkExpander allows the user to reveal its child by clicking on an
// expander triangle.
//
// !An example GtkExpander (expander.png)
//
// This is similar to the triangles used in a GtkTreeView.
//
// Normally you use an expander as you would use a frame; you create the child
// widget and use gtk.Expander.SetChild() to add it to the expander. When the
// expander is toggled, it will take care of showing and hiding the child
// automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a GtkExpander but do not add a child to
// it. The expander widget has an [propertyGtk.Expander:expanded[ property which
// can be used to monitor its expansion state. You should watch this property
// with a signal connection as follows:
//
//    static void
//    expander_callback (GObject    *object,
//                       GParamSpec *param_spec,
//                       gpointer    user_data)
//    {
//      GtkExpander *expander;
//
//      expander = GTK_EXPANDER (object);
//
//      if (gtk_expander_get_expanded (expander))
//        {
//          // Show or create widgets
//        }
//      else
//        {
//          // Hide or destroy widgets
//        }
//    }
//
//    static void
//    create_expander (void)
//    {
//      GtkWidget *expander = gtk_expander_new_with_mnemonic ("_More Options");
//      g_signal_connect (expander, "notify::expanded",
//                        G_CALLBACK (expander_callback), NULL);
//
//      // ...
//    }
//
//
//
// GtkExpander as GtkBuildable
//
// The GtkExpander implementation of the GtkBuildable interface supports placing
// a child in the label position by specifying “label” as the “type” attribute
// of a <child> element. A normal content child can be specified without
// specifying a <child> type attribute.
//
// An example of a UI definition fragment with GtkExpander:
//
//    <object class="GtkExpander">
//      <child type="label">
//        <object class="GtkLabel" id="expander-label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="expander-content"/>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    expander
//    ╰── box
//        ├── title
//        │   ├── arrow
//        │   ╰── <label widget>
//        ╰── <child>
//
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
//
//
// Accessibility
//
// GtkExpander uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type Expander struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Expander)(nil)
)

func wrapExpander(obj *coreglib.Object) *Expander {
	return &Expander{
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

func marshalExpander(p uintptr) (interface{}, error) {
	return wrapExpander(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate activates the GtkExpander.
func (v *Expander) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate", false, unsafe.Pointer(C._gotk4_gtk4_Expander_ConnectActivate), f)
}
