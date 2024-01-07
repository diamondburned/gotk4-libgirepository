// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_SpinButton_ConnectWrapped(gpointer, guintptr);
// extern void _gotk4_gtk4_SpinButton_ConnectValueChanged(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_SpinButton_ConnectOutput(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSpinButtonUpdatePolicy = coreglib.Type(girepository.MustFind("Gtk", "SpinButtonUpdatePolicy").RegisteredGType())
	GTypeSpinType               = coreglib.Type(girepository.MustFind("Gtk", "SpinType").RegisteredGType())
	GTypeSpinButton             = coreglib.Type(girepository.MustFind("Gtk", "SpinButton").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinButtonUpdatePolicy, F: marshalSpinButtonUpdatePolicy},
		coreglib.TypeMarshaler{T: GTypeSpinType, F: marshalSpinType},
		coreglib.TypeMarshaler{T: GTypeSpinButton, F: marshalSpinButton},
	})
}

// INPUT_ERROR: constant to return from a signal handler for the ::input signal
// in case of conversion failure.
//
// See gtk.SpinButton::input.
const INPUT_ERROR = -1

// SpinButtonUpdatePolicy determines whether the spin button displays values
// outside the adjustment bounds.
//
// See gtk.SpinButton.SetUpdatePolicy().
type SpinButtonUpdatePolicy C.gint

const (
	// UpdateAlways: when refreshing your SpinButton, the value is always
	// displayed.
	UpdateAlways SpinButtonUpdatePolicy = iota
	// UpdateIfValid: when refreshing your SpinButton, the value is only
	// displayed if it is valid within the bounds of the spin button's
	// adjustment.
	UpdateIfValid
)

func marshalSpinButtonUpdatePolicy(p uintptr) (interface{}, error) {
	return SpinButtonUpdatePolicy(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinButtonUpdatePolicy.
func (s SpinButtonUpdatePolicy) String() string {
	switch s {
	case UpdateAlways:
		return "Always"
	case UpdateIfValid:
		return "IfValid"
	default:
		return fmt.Sprintf("SpinButtonUpdatePolicy(%d)", s)
	}
}

// SpinType values of the GtkSpinType enumeration are used to specify the change
// to make in gtk_spin_button_spin().
type SpinType C.gint

const (
	// SpinStepForward: increment by the adjustments step increment.
	SpinStepForward SpinType = iota
	// SpinStepBackward: decrement by the adjustments step increment.
	SpinStepBackward
	// SpinPageForward: increment by the adjustments page increment.
	SpinPageForward
	// SpinPageBackward: decrement by the adjustments page increment.
	SpinPageBackward
	// SpinHome: go to the adjustments lower bound.
	SpinHome
	// SpinEnd: go to the adjustments upper bound.
	SpinEnd
	// SpinUserDefined: change by a specified amount.
	SpinUserDefined
)

func marshalSpinType(p uintptr) (interface{}, error) {
	return SpinType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinType.
func (s SpinType) String() string {
	switch s {
	case SpinStepForward:
		return "StepForward"
	case SpinStepBackward:
		return "StepBackward"
	case SpinPageForward:
		return "PageForward"
	case SpinPageBackward:
		return "PageBackward"
	case SpinHome:
		return "Home"
	case SpinEnd:
		return "End"
	case SpinUserDefined:
		return "UserDefined"
	default:
		return fmt.Sprintf("SpinType(%d)", s)
	}
}

// SpinButton: GtkSpinButton is an ideal way to allow the user to set the value
// of some attribute.
//
// !An example GtkSpinButton (spinbutton.png)
//
// Rather than having to directly type a number into a GtkEntry, GtkSpinButton
// allows the user to click on one of two arrows to increment or decrement the
// displayed value. A value can still be typed in, with the bonus that it can be
// checked to ensure it is in a given range.
//
// The main properties of a GtkSpinButton are through an adjustment. See the
// gtk.Adjustment documentation for more details about an adjustment's
// properties.
//
// Note that GtkSpinButton will by default make its entry large enough to
// accommodate the lower and upper bounds of the adjustment. If this is not
// desired, the automatic sizing can be turned off by explicitly setting
// gtk.Editable:width-chars to a value != -1.
//
// Using a GtkSpinButton to get an integer
//
//    // Provides a function to retrieve an integer value from a GtkSpinButton
//    // and creates a spin button to model percentage values.
//
//    int
//    grab_int_value (GtkSpinButton *button,
//                    gpointer       user_data)
//    {
//      return gtk_spin_button_get_value_as_int (button);
//    }
//
//    void
//    create_integer_spin_button (void)
//    {
//
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (50.0, 0.0, 100.0, 1.0, 5.0, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with no decimal places
//      button = gtk_spin_button_new (adjustment, 1.0, 0);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// Using a GtkSpinButton to get a floating point value
//
//    // Provides a function to retrieve a floating point value from a
//    // GtkSpinButton, and creates a high precision spin button.
//
//    float
//    grab_float_value (GtkSpinButton *button,
//                      gpointer       user_data)
//    {
//      return gtk_spin_button_get_value (button);
//    }
//
//    void
//    create_floating_spin_button (void)
//    {
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (2.500, 0.0, 5.0, 0.001, 0.1, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with three decimal places
//      button = gtk_spin_button_new (adjustment, 0.001, 3);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// CSS nodes
//
//    spinbutton.horizontal
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ├── button.down
//    ╰── button.up
//
//
//
//
//    spinbutton.vertical
//    ├── button.up
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ╰── button.down
//
//
// GtkSpinButtons main CSS node has the name spinbutton. It creates subnodes for
// the entry and the two buttons, with these names. The button nodes have the
// style classes .up and .down. The GtkText subnodes (if present) are put below
// the text node. The orientation of the spin button is reflected in the
// .vertical or .horizontal style class on the main node.
//
//
// Accessiblity
//
// GtkSpinButton uses the GTK_ACCESSIBLE_ROLE_SPIN_BUTTON role.
type SpinButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	coreglib.InitiallyUnowned
	Accessible
	Buildable
	CellEditable
	ConstraintTarget
	Editable
	Orientable
}

var (
	_ Widgetter         = (*SpinButton)(nil)
	_ coreglib.Objector = (*SpinButton)(nil)
)

func wrapSpinButton(obj *coreglib.Object) *SpinButton {
	return &SpinButton{
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
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		CellEditable: CellEditable{
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
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		Editable: Editable{
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
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	return wrapSpinButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectOutput is emitted to tweak the formatting of the value for display.
//
//    // show leading zeros
//    static gboolean
//    on_output (GtkSpinButton *spin,
//               gpointer       data)
//    {
//       GtkAdjustment *adjustment;
//       char *text;
//       int value;
//
//       adjustment = gtk_spin_button_get_adjustment (spin);
//       value = (int)gtk_adjustment_get_value (adjustment);
//       text = g_strdup_printf ("02d", value);
//       gtk_spin_button_set_text (spin, text):
//       g_free (text);
//
//       return TRUE;
//    }.
func (v *SpinButton) ConnectOutput(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "output", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectOutput), f)
}

// ConnectValueChanged is emitted when the value is changed.
//
// Also see the gtk.SpinButton::output signal.
func (v *SpinButton) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectValueChanged), f)
}

// ConnectWrapped is emitted right after the spinbutton wraps from its maximum
// to its minimum value or vice-versa.
func (v *SpinButton) ConnectWrapped(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "wrapped", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectWrapped), f)
}
