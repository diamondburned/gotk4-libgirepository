// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_ButtonClass_activate(void*);
// extern void _gotk4_gtk4_ButtonClass_clicked(void*);
// extern void _gotk4_gtk4_Button_ConnectActivate(gpointer, guintptr);
// extern void _gotk4_gtk4_Button_ConnectClicked(gpointer, guintptr);
import "C"

// glib.Type values for gtkbutton.go.
var GTypeButton = coreglib.Type(C.gtk_button_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeButton, F: marshalButton},
	})
}

// ButtonOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Actionable
}

var (
	_ Widgetter         = (*Button)(nil)
	_ coreglib.Objector = (*Button)(nil)
)

func classInitButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ButtonClass")

	if _, ok := goval.(interface{ Activate() }); ok {
		o := pclass.StructFieldOffset("activate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_ButtonClass_activate)
	}

	if _, ok := goval.(interface{ Clicked() }); ok {
		o := pclass.StructFieldOffset("clicked")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_ButtonClass_clicked)
	}
}

//export _gotk4_gtk4_ButtonClass_activate
func _gotk4_gtk4_ButtonClass_activate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

//export _gotk4_gtk4_ButtonClass_clicked
func _gotk4_gtk4_ButtonClass_clicked(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Clicked() })

	iface.Clicked()
}

func wrapButton(obj *coreglib.Object) *Button {
	return &Button{
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
		Actionable: Actionable{
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
	}
}

func marshalButton(p uintptr) (interface{}, error) {
	return wrapButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Button_ConnectActivate
func _gotk4_gtk4_Button_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate is emitted to animate press then release.
//
// This is an action signal. Applications should never connect to this signal,
// but use the gtk.Button::clicked signal.
func (button *Button) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "activate", false, unsafe.Pointer(C._gotk4_gtk4_Button_ConnectActivate), f)
}

//export _gotk4_gtk4_Button_ConnectClicked
func _gotk4_gtk4_Button_ConnectClicked(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClicked is emitted when the button has been activated (pressed and
// released).
func (button *Button) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "clicked", false, unsafe.Pointer(C._gotk4_gtk4_Button_ConnectClicked), f)
}

// NewButton creates a new GtkButton widget.
//
// To add a child widget to the button, use gtk.Button.SetChild().
//
// The function returns the following values:
//
//    - button: newly created GtkButton widget.
//
func NewButton() *Button {
	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("new_Button", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonFromIconName creates a new button containing an icon from the
// current icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name or NULL.
//
// The function returns the following values:
//
//    - button: new GtkButton displaying the themed icon.
//
func NewButtonFromIconName(iconName string) *Button {
	var _args [1]girepository.Argument

	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("new_Button_from_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iconName)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonWithLabel creates a GtkButton widget with a GtkLabel child.
//
// The function takes the following parameters:
//
//    - label: text you want the GtkLabel to hold.
//
// The function returns the following values:
//
//    - button: newly created GtkButton widget.
//
func NewButtonWithLabel(label string) *Button {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("new_Button_with_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// NewButtonWithMnemonic creates a new GtkButton containing a label.
//
// If characters in label are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - button: new GtkButton.
//
func NewButtonWithMnemonic(label string) *Button {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("new_Button_with_mnemonic", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// Child gets the child widget of button.
//
// The function returns the following values:
//
//    - widget (optional): child widget of button.
//
func (button *Button) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// HasFrame returns whether the button has a frame.
//
// The function returns the following values:
//
//    - ok: TRUE if the button has a frame.
//
func (button *Button) HasFrame() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("get_has_frame", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IconName returns the icon name of the button.
//
// If the icon name has not been set with gtk.Button.SetIconName() the return
// value will be NULL. This will be the case if you create an empty button with
// gtk.Button.New to use as a container.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name set via gtk.Button.SetIconName().
//
func (button *Button) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("get_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Label fetches the text from the label of the button.
//
// If the label text has not been set with gtk.Button.SetLabel() the return
// value will be NULL. This will be the case if you create an empty button with
// gtk.Button.New to use as a container.
//
// The function returns the following values:
//
//    - utf8 (optional): text of the label widget. This string is owned by the
//      widget and must not be modified or freed.
//
func (button *Button) Label() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("get_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline gets whether underlines are interpreted as mnemonics.
//
// See gtk.Button.SetUseUnderline().
//
// The function returns the following values:
//
//    - ok: TRUE if an embedded underline in the button label indicates the
//      mnemonic accelerator keys.
//
func (button *Button) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "Button").InvokeMethod("get_use_underline", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of button.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (button *Button) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "Button").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(child)
}

// SetHasFrame sets the style of the button.
//
// Buttons can has a flat appearance or have a frame drawn around them.
//
// The function takes the following parameters:
//
//    - hasFrame: whether the button should have a visible frame.
//
func (button *Button) SetHasFrame(hasFrame bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if hasFrame {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Button").InvokeMethod("set_has_frame", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(hasFrame)
}

// SetIconName adds a GtkImage with the given icon name as a child.
//
// If button already contains a child widget, that child widget will be removed
// and replaced with the image.
//
// The function takes the following parameters:
//
//    - iconName: icon name.
//
func (button *Button) SetIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "Button").InvokeMethod("set_icon_name", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the text of the label of the button to label.
//
// This will also clear any previously set labels.
//
// The function takes the following parameters:
//
//    - label: string.
//
func (button *Button) SetLabel(label string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "Button").InvokeMethod("set_label", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(label)
}

// SetUseUnderline sets whether to use underlines as mnemonics.
//
// If true, an underline in the text of the button label indicates the next
// character should be used for the mnemonic accelerator key.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (button *Button) SetUseUnderline(useUnderline bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if useUnderline {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Button").InvokeMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(useUnderline)
}
