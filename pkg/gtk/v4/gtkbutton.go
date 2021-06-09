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
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_get_type()), F: marshalButton},
	})
}

// Button: the Button widget is generally used to trigger a callback function
// that is called when the button is pressed. The various signals and how to use
// them are outlined below.
//
// The Button widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is the Label.
//
//
// CSS nodes
//
// GtkButton has a single CSS node with name button. The node will get the style
// classes .image-button or .text-button, if the content is just an image or
// label, respectively. It may also receive the .flat style class.
//
// Other style classes that are commonly used with GtkButton include
// .suggested-action and .destructive-action. In special cases, buttons can be
// made round by adding the .circular style class.
//
// Button-like widgets like ToggleButton, MenuButton, VolumeButton, LockButton,
// ColorButton or FontButton use style classes such as .toggle, .popup, .scale,
// .lock, .color on the button node to differentiate themselves from a plain
// GtkButton.
//
//
// Accessibility
//
// GtkButton uses the K_ACCESSIBLE_ROLE_BUTTON role.
type Button interface {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget

	// Child gets the child widget of @button.
	Child() Widget
	// HasFrame returns whether the button has a frame.
	HasFrame() bool
	// IconName returns the icon name set via gtk_button_set_icon_name().
	IconName() string
	// Label fetches the text from the label of the button, as set by
	// gtk_button_set_label(). If the label text has not been set the return
	// value will be nil. This will be the case if you create an empty button
	// with gtk_button_new() to use as a container.
	Label() string
	// UseUnderline returns whether an embedded underline in the button label
	// indicates a mnemonic. See gtk_button_set_use_underline().
	UseUnderline() bool
	// SetChild sets the child widget of @button.
	SetChild(child Widget)
	// SetHasFrame sets the style of the button. Buttons can has a flat
	// appearance or have a frame drawn around them.
	SetHasFrame(hasFrame bool)
	// SetIconName adds a Image with the given icon name as a child. If @button
	// already contains a child widget, that child widget will be removed and
	// replaced with the image.
	SetIconName(iconName string)
	// SetLabel sets the text of the label of the button to @label.
	//
	// This will also clear any previously set labels.
	SetLabel(label string)
	// SetUseUnderline: if true, an underline in the text of the button label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(useUnderline bool)
}

// button implements the Button interface.
type button struct {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget
}

var _ Button = (*button)(nil)

// WrapButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapButton(obj *externglib.Object) Button {
	return Button{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Actionable:       WrapActionable(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButton(obj), nil
}

// NewButton constructs a class Button.
func NewButton() Button {
	var _cret C.GtkButton

	cret = C.gtk_button_new()

	var _button Button

	_button = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Button)

	return _button
}

// NewButtonFromIconName constructs a class Button.
func NewButtonFromIconName(iconName string) Button {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkButton

	cret = C.gtk_button_new_from_icon_name(_arg1)

	var _button Button

	_button = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Button)

	return _button
}

// NewButtonWithLabel constructs a class Button.
func NewButtonWithLabel(label string) Button {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkButton

	cret = C.gtk_button_new_with_label(_arg1)

	var _button Button

	_button = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Button)

	return _button
}

// NewButtonWithMnemonic constructs a class Button.
func NewButtonWithMnemonic(label string) Button {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkButton

	cret = C.gtk_button_new_with_mnemonic(_arg1)

	var _button Button

	_button = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Button)

	return _button
}

// Child gets the child widget of @button.
func (b button) Child() Widget {
	var _arg0 *C.GtkButton

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_button_get_child(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// HasFrame returns whether the button has a frame.
func (b button) HasFrame() bool {
	var _arg0 *C.GtkButton

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_button_get_has_frame(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IconName returns the icon name set via gtk_button_set_icon_name().
func (b button) IconName() string {
	var _arg0 *C.GtkButton

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var _cret *C.char

	cret = C.gtk_button_get_icon_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Label fetches the text from the label of the button, as set by
// gtk_button_set_label(). If the label text has not been set the return
// value will be nil. This will be the case if you create an empty button
// with gtk_button_new() to use as a container.
func (b button) Label() string {
	var _arg0 *C.GtkButton

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var _cret *C.char

	cret = C.gtk_button_get_label(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UseUnderline returns whether an embedded underline in the button label
// indicates a mnemonic. See gtk_button_set_use_underline().
func (b button) UseUnderline() bool {
	var _arg0 *C.GtkButton

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_button_get_use_underline(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of @button.
func (b button) SetChild(child Widget) {
	var _arg0 *C.GtkButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_button_set_child(_arg0, _arg1)
}

// SetHasFrame sets the style of the button. Buttons can has a flat
// appearance or have a frame drawn around them.
func (b button) SetHasFrame(hasFrame bool) {
	var _arg0 *C.GtkButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if hasFrame {
		_arg1 = C.gboolean(1)
	}

	C.gtk_button_set_has_frame(_arg0, _arg1)
}

// SetIconName adds a Image with the given icon name as a child. If @button
// already contains a child widget, that child widget will be removed and
// replaced with the image.
func (b button) SetIconName(iconName string) {
	var _arg0 *C.GtkButton
	var _arg1 *C.char

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_button_set_icon_name(_arg0, _arg1)
}

// SetLabel sets the text of the label of the button to @label.
//
// This will also clear any previously set labels.
func (b button) SetLabel(label string) {
	var _arg0 *C.GtkButton
	var _arg1 *C.char

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_button_set_label(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text of the button label
// indicates the next character should be used for the mnemonic accelerator
// key.
func (b button) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if useUnderline {
		_arg1 = C.gboolean(1)
	}

	C.gtk_button_set_use_underline(_arg0, _arg1)
}
