// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_CheckButtonClass_draw_indicator(void*, void*);
import "C"

// glib.Type values for gtkcheckbutton.go.
var GTypeCheckButton = coreglib.Type(C.gtk_check_button_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCheckButton, F: marshalCheckButton},
	})
}

// CheckButtonOverrider contains methods that are overridable.
type CheckButtonOverrider interface {
	// The function takes the following parameters:
	//
	DrawIndicator(cr *cairo.Context)
}

// CheckButton places a discrete ToggleButton next to a widget, (usually a
// Label). See the section on ToggleButton widgets for more information about
// toggle/check buttons.
//
// The important signal ( ToggleButton::toggled ) is also inherited from
// ToggleButton.
//
// CSS nodes
//
//    button.check
//    ├── check
//    ╰── <child>
//
// A GtkCheckButton without indicator changes the name of its main node to
// button and adds a .check style class to it. The subnode is invisible in this
// case.
type CheckButton struct {
	_ [0]func() // equal guard
	ToggleButton
}

var (
	_ Binner            = (*CheckButton)(nil)
	_ coreglib.Objector = (*CheckButton)(nil)
)

func classInitCheckButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCheckButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCheckButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ DrawIndicator(cr *cairo.Context) }); ok {
		pclass.draw_indicator = (*[0]byte)(C._gotk4_gtk3_CheckButtonClass_draw_indicator)
	}
}

//export _gotk4_gtk3_CheckButtonClass_draw_indicator
func _gotk4_gtk3_CheckButtonClass_draw_indicator(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DrawIndicator(cr *cairo.Context) })

	var _cr *cairo.Context // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	iface.DrawIndicator(_cr)
}

func wrapCheckButton(obj *coreglib.Object) *CheckButton {
	return &CheckButton{
		ToggleButton: ToggleButton{
			Button: Button{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: coreglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
						},
					},
				},
				Object: obj,
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
				Activatable: Activatable{
					Object: obj,
				},
			},
		},
	}
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	return wrapCheckButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCheckButton creates a new CheckButton.
//
// The function returns the following values:
//
//    - checkButton: Widget.
//
func NewCheckButton() *CheckButton {
	_gret := girepository.MustFind("Gtk", "CheckButton").InvokeMethod("new_CheckButton", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithLabel creates a new CheckButton with a Label to the right
// of it.
//
// The function takes the following parameters:
//
//    - label: text for the check button.
//
// The function returns the following values:
//
//    - checkButton: Widget.
//
func NewCheckButtonWithLabel(label string) *CheckButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "CheckButton").InvokeMethod("new_CheckButton_with_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithMnemonic creates a new CheckButton containing a label. The
// label will be created using gtk_label_new_with_mnemonic(), so underscores in
// label indicate the mnemonic for the check button.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - checkButton: new CheckButton.
//
func NewCheckButtonWithMnemonic(label string) *CheckButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "CheckButton").InvokeMethod("new_CheckButton_with_mnemonic", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}
