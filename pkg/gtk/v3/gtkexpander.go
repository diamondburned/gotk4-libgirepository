// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_ExpanderClass_activate(void*);
// extern void _gotk4_gtk3_Expander_ConnectActivate(gpointer, guintptr);
import "C"

// glib.Type values for gtkexpander.go.
var GTypeExpander = coreglib.Type(C.gtk_expander_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeExpander, F: marshalExpander},
	})
}

// ExpanderOverrider contains methods that are overridable.
type ExpanderOverrider interface {
	Activate()
}

// Expander allows the user to hide or show its child by clicking on an expander
// triangle similar to the triangles used in a TreeView.
//
// Normally you use an expander as you would use any other descendant of Bin;
// you create the child widget and use gtk_container_add() to add it to the
// expander. When the expander is toggled, it will take care of showing and
// hiding the child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a Expander but do not add a child to it.
// The expander widget has an Expander:expanded property which can be used to
// monitor its expansion state. You should watch this property with a signal
// connection as follows:
//
//    expander
//    ├── title
//    │   ├── arrow
//    │   ╰── <label widget>
//    ╰── <child>
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
type Expander struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Expander)(nil)
)

func classInitExpanderer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ExpanderClass")

	if _, ok := goval.(interface{ Activate() }); ok {
		o := pclass.StructFieldOffset("activate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ExpanderClass_activate)
	}
}

//export _gotk4_gtk3_ExpanderClass_activate
func _gotk4_gtk3_ExpanderClass_activate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

func wrapExpander(obj *coreglib.Object) *Expander {
	return &Expander{
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
	}
}

func marshalExpander(p uintptr) (interface{}, error) {
	return wrapExpander(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Expander_ConnectActivate
func _gotk4_gtk3_Expander_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

func (expander *Expander) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(expander, "activate", false, unsafe.Pointer(C._gotk4_gtk3_Expander_ConnectActivate), f)
}

// NewExpander creates a new expander using label as the text of the label.
//
// The function takes the following parameters:
//
//    - label (optional): text of the label.
//
// The function returns the following values:
//
//    - expander: new Expander widget.
//
func NewExpander(label string) *Expander {
	var _args [1]girepository.Argument

	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("new_Expander", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(coreglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// NewExpanderWithMnemonic creates a new expander using label as the text of the
// label. If characters in label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use “__”
// (two underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
//
// The function takes the following parameters:
//
//    - label (optional): text of the label with an underscore in front of the
//      mnemonic character.
//
// The function returns the following values:
//
//    - expander: new Expander widget.
//
func NewExpanderWithMnemonic(label string) *Expander {
	var _args [1]girepository.Argument

	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("new_Expander_with_mnemonic", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(coreglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// Expanded queries a Expander and returns its current state. Returns TRUE if
// the child widget is revealed.
//
// See gtk_expander_set_expanded().
//
// The function returns the following values:
//
//    - ok: current state of the expander.
//
func (expander *Expander) Expanded() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_expanded", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Label fetches the text from a label widget including any embedded underlines
// indicating mnemonics and Pango markup, as set by gtk_expander_set_label(). If
// the label text has not been set the return value will be NULL. This will be
// the case if you create an empty button with gtk_button_new() to use as a
// container.
//
// Note that this function behaved differently in versions prior to 2.14 and
// used to return the label text stripped of embedded underlines indicating
// mnemonics and Pango markup. This problem can be avoided by fetching the label
// text directly from the label widget.
//
// The function returns the following values:
//
//    - utf8 (optional): text of the label widget. This string is owned by the
//      widget and must not be modified or freed.
//
func (expander *Expander) Label() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelFill returns whether the label widget will fill all available horizontal
// space allocated to expander.
//
// The function returns the following values:
//
//    - ok: TRUE if the label widget will fill all available horizontal space.
//
func (expander *Expander) LabelFill() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_label_fill", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
//
// The function returns the following values:
//
//    - widget (optional): label widget, or NULL if there is none.
//
func (expander *Expander) LabelWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_label_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

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

// ResizeToplevel returns whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
//
// The function returns the following values:
//
//    - ok: “resize toplevel” setting.
//
func (expander *Expander) ResizeToplevel() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_resize_toplevel", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_expander_set_spacing().
//
// Deprecated: Use margins on the child instead.
//
// The function returns the following values:
//
//    - gint: spacing between the expander and child.
//
func (expander *Expander) Spacing() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_spacing", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// UseMarkup returns whether the label’s text is interpreted as marked up with
// the [Pango text markup language][PangoMarkupFormat]. See
// gtk_expander_set_use_markup().
//
// The function returns the following values:
//
//    - ok: TRUE if the label’s text will be parsed for markup.
//
func (expander *Expander) UseMarkup() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_use_markup", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
//
// The function returns the following values:
//
//    - ok: TRUE if an embedded underline in the expander label indicates the
//      mnemonic accelerator keys.
//
func (expander *Expander) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_gret := girepository.MustFind("Gtk", "Expander").InvokeMethod("get_use_underline", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expander)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded sets the state of the expander. Set to TRUE, if you want the
// child widget to be revealed, and FALSE if you want the child widget to be
// hidden.
//
// The function takes the following parameters:
//
//    - expanded: whether the child widget is revealed.
//
func (expander *Expander) SetExpanded(expanded bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if expanded {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_expanded", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(expanded)
}

// SetLabel sets the text of the label of the expander to label.
//
// This will also clear any previously set labels.
//
// The function takes the following parameters:
//
//    - label (optional): string.
//
func (expander *Expander) SetLabel(label string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_label", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(label)
}

// SetLabelFill sets whether the label widget should fill all available
// horizontal space allocated to expander.
//
// Note that this function has no effect since 3.20.
//
// The function takes the following parameters:
//
//    - labelFill: TRUE if the label should should fill all available horizontal
//      space.
//
func (expander *Expander) SetLabelFill(labelFill bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if labelFill {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_label_fill", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(labelFill)
}

// SetLabelWidget: set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
//
// The function takes the following parameters:
//
//    - labelWidget (optional): new label widget.
//
func (expander *Expander) SetLabelWidget(labelWidget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if labelWidget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(labelWidget).Native()))
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_label_widget", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(labelWidget)
}

// SetResizeToplevel sets whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
//
// The function takes the following parameters:
//
//    - resizeToplevel: whether to resize the toplevel.
//
func (expander *Expander) SetResizeToplevel(resizeToplevel bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if resizeToplevel {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_resize_toplevel", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(resizeToplevel)
}

// SetSpacing sets the spacing field of expander, which is the number of pixels
// to place between expander and the child.
//
// Deprecated: Use margins on the child instead.
//
// The function takes the following parameters:
//
//    - spacing: distance between the expander and child in pixels.
//
func (expander *Expander) SetSpacing(spacing int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(spacing)

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_spacing", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(spacing)
}

// SetUseMarkup sets whether the text of the label contains markup in [Pango’s
// text markup language][PangoMarkupFormat]. See gtk_label_set_markup().
//
// The function takes the following parameters:
//
//    - useMarkup: TRUE if the label’s text should be parsed for markup.
//
func (expander *Expander) SetUseMarkup(useMarkup bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if useMarkup {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_use_markup", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(useMarkup)
}

// SetUseUnderline: if true, an underline in the text of the expander label
// indicates the next character should be used for the mnemonic accelerator key.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (expander *Expander) SetUseUnderline(useUnderline bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if useUnderline {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Expander").InvokeMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(expander)
	runtime.KeepAlive(useUnderline)
}
