// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_get_type()), F: marshalButtonner},
	})
}

// ButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ButtonOverrider interface {
	Activate()
	Clicked()
}

// Button: GtkButton widget is generally used to trigger a callback function
// that is called when the button is pressed.
//
// !An example GtkButton (button.png)
//
// The GtkButton widget can hold any valid child widget. That is, it can hold
// almost any other standard GtkWidget. The most commonly used child is the
// GtkLabel.
//
//
// CSS nodes
//
// GtkButton has a single CSS node with name button. The node will get the style
// classes .image-button or .text-button, if the content is just an image or
// label, respectively. It may also receive the .flat style class. When
// activating a button via the keyboard, the button will temporarily gain the
// .keyboard-activating style class.
//
// Other style classes that are commonly used with GtkButton include
// .suggested-action and .destructive-action. In special cases, buttons can be
// made round by adding the .circular style class.
//
// Button-like widgets like gtk.ToggleButton, gtk.MenuButton, gtk.VolumeButton,
// gtk.LockButton, gtk.ColorButton or gtk.FontButton use style classes such as
// .toggle, .popup, .scale, .lock, .color on the button node to differentiate
// themselves from a plain GtkButton.
//
//
// Accessibility
//
// GtkButton uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type Button struct {
	Widget

	Actionable
	*externglib.Object
}

func wrapButton(obj *externglib.Object) *Button {
	return &Button{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapButton(obj), nil
}

// NewButton creates a new GtkButton widget.
//
// To add a child widget to the button, use gtk.Button.SetChild().
func NewButton() *Button {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_button_new()

	var _button *Button // out

	_button = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonFromIconName creates a new button containing an icon from the
// current icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
func NewButtonFromIconName(iconName string) *Button {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_from_icon_name(_arg1)

	var _button *Button // out

	_button = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonWithLabel creates a GtkButton widget with a GtkLabel child.
func NewButtonWithLabel(label string) *Button {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_with_label(_arg1)

	var _button *Button // out

	_button = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonWithMnemonic creates a new GtkButton containing a label.
//
// If characters in label are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
func NewButtonWithMnemonic(label string) *Button {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_with_mnemonic(_arg1)

	var _button *Button // out

	_button = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// Child gets the child widget of button.
func (button *Button) Child() Widgetter {
	var _arg0 *C.GtkButton // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_child(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// HasFrame returns whether the button has a frame.
func (button *Button) HasFrame() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName returns the icon name of the button.
//
// If the icon name has not been set with gtk.Button.SetIconName() the return
// value will be NULL. This will be the case if you create an empty button with
// gtk.Button.New to use as a container.
func (button *Button) IconName() string {
	var _arg0 *C.GtkButton // out
	var _cret *C.char      // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Label fetches the text from the label of the button.
//
// If the label text has not been set with gtk.Button.SetLabel() the return
// value will be NULL. This will be the case if you create an empty button with
// gtk.Button.New to use as a container.
func (button *Button) Label() string {
	var _arg0 *C.GtkButton // out
	var _cret *C.char      // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether underlines are interpreted as mnemonics.
//
// See gtk.Button.SetUseUnderline().
func (button *Button) UseUnderline() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of button.
func (button *Button) SetChild(child Widgetter) {
	var _arg0 *C.GtkButton // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_button_set_child(_arg0, _arg1)
}

// SetHasFrame sets the style of the button.
//
// Buttons can has a flat appearance or have a frame drawn around them.
func (button *Button) SetHasFrame(hasFrame bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if hasFrame {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_has_frame(_arg0, _arg1)
}

// SetIconName adds a GtkImage with the given icon name as a child.
//
// If button already contains a child widget, that child widget will be removed
// and replaced with the image.
func (button *Button) SetIconName(iconName string) {
	var _arg0 *C.GtkButton // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_button_set_icon_name(_arg0, _arg1)
}

// SetLabel sets the text of the label of the button to label.
//
// This will also clear any previously set labels.
func (button *Button) SetLabel(label string) {
	var _arg0 *C.GtkButton // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_button_set_label(_arg0, _arg1)
}

// SetUseUnderline sets whether to use underlines as mnemonics.
//
// If true, an underline in the text of the button label indicates the next
// character should be used for the mnemonic accelerator key.
func (button *Button) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_use_underline(_arg0, _arg1)
}
