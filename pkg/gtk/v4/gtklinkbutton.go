// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButtonner},
	})
}

// LinkButton: GtkLinkButton is a button with a hyperlink.
//
// !An example GtkLinkButton (link-button.png)
//
// It is useful to show quick links to resources.
//
// A link button is created by calling either gtk.LinkButton.New or
// gtk.LinkButton.NewWithLabel. If using the former, the URI you pass to the
// constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk.LinkButton.SetURI().
//
// By default, GtkLinkButton calls gtk.ShowURI() when the button is clicked.
// This behaviour can be overridden by connecting to the
// gtk.LinkButton::activate-link signal and returning TRUE from the signal
// handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain GtkButton, it gets the .link style class.
//
//
// Accessibility
//
// GtkLinkButton uses the K_ACCESSIBLE_ROLE_LINK role.
type LinkButton struct {
	Button
}

var (
	_ Widgetter           = (*LinkButton)(nil)
	_ externglib.Objector = (*LinkButton)(nil)
)

func wrapLinkButton(obj *externglib.Object) *LinkButton {
	return &LinkButton{
		Button: Button{
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
		},
	}
}

func marshalLinkButtonner(p uintptr) (interface{}, error) {
	return wrapLinkButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLinkButton creates a new GtkLinkButton with the URI as its text.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
func NewLinkButton(uri string) *LinkButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_link_button_new(_arg1)
	runtime.KeepAlive(uri)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(externglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// NewLinkButtonWithLabel creates a new GtkLinkButton containing a label.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//    - label: text of the button.
//
func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(label)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(externglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// URI retrieves the URI of the GtkLinkButton.
func (linkButton *LinkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_uri(_arg0)
	runtime.KeepAlive(linkButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Visited retrieves the “visited” state of the GtkLinkButton.
//
// The button becomes visited when it is clicked. If the URI is changed on the
// button, the “visited” state is unset again.
//
// The state may also be changed using gtk.LinkButton.SetVisited().
func (linkButton *LinkButton) Visited() bool {
	var _arg0 *C.GtkLinkButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_visited(_arg0)
	runtime.KeepAlive(linkButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetURI sets uri as the URI where the GtkLinkButton points.
//
// As a side-effect this unsets the “visited” state of the button.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
func (linkButton *LinkButton) SetURI(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(uri)
}

// SetVisited sets the “visited” state of the GtkLinkButton.
//
// See gtk.LinkButton.GetVisited() for more details.
//
// The function takes the following parameters:
//
//    - visited: new “visited” state.
//
func (linkButton *LinkButton) SetVisited(visited bool) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	if visited {
		_arg1 = C.TRUE
	}

	C.gtk_link_button_set_visited(_arg0, _arg1)
	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(visited)
}

// ConnectActivateLink: emitted each time the GtkLinkButton is clicked.
//
// The default handler will call gtk.ShowURI() with the URI stored inside the
// gtk.LinkButton:uri property.
//
// To override the default behavior, you can connect to the ::activate-link
// signal and stop the propagation of the signal by returning TRUE from your
// handler.
func (linkButton *LinkButton) ConnectActivateLink(f func() bool) externglib.SignalHandle {
	return linkButton.Connect("activate-link", f)
}
