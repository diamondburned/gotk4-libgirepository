// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeButtonBoxStyle = coreglib.Type(girepository.MustFind("Gtk", "ButtonBoxStyle").RegisteredGType())
	GTypeButtonBox      = coreglib.Type(girepository.MustFind("Gtk", "ButtonBox").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeButtonBoxStyle, F: marshalButtonBoxStyle},
		coreglib.TypeMarshaler{T: GTypeButtonBox, F: marshalButtonBox},
	})
}

// ButtonBoxStyle: used to dictate the style that a ButtonBox uses to layout the
// buttons it contains.
type ButtonBoxStyle C.gint

const (
	// ButtonboxSpread buttons are evenly spread across the box.
	ButtonboxSpread ButtonBoxStyle = 1
	// ButtonboxEdge buttons are placed at the edges of the box.
	ButtonboxEdge ButtonBoxStyle = 2
	// ButtonboxStart buttons are grouped towards the start of the box, (on the
	// left for a HBox, or the top for a VBox).
	ButtonboxStart ButtonBoxStyle = 3
	// ButtonboxEnd buttons are grouped towards the end of the box, (on the
	// right for a HBox, or the bottom for a VBox).
	ButtonboxEnd ButtonBoxStyle = 4
	// ButtonboxCenter buttons are centered in the box. Since 2.12.
	ButtonboxCenter ButtonBoxStyle = 5
	// ButtonboxExpand buttons expand to fill the box. This entails giving
	// buttons a "linked" appearance, making button sizes homogeneous, and
	// setting spacing to 0 (same as calling gtk_box_set_homogeneous() and
	// gtk_box_set_spacing() manually). Since 3.12.
	ButtonboxExpand ButtonBoxStyle = 6
)

func marshalButtonBoxStyle(p uintptr) (interface{}, error) {
	return ButtonBoxStyle(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ButtonBoxStyle.
func (b ButtonBoxStyle) String() string {
	switch b {
	case ButtonboxSpread:
		return "Spread"
	case ButtonboxEdge:
		return "Edge"
	case ButtonboxStart:
		return "Start"
	case ButtonboxEnd:
		return "End"
	case ButtonboxCenter:
		return "Center"
	case ButtonboxExpand:
		return "Expand"
	default:
		return fmt.Sprintf("ButtonBoxStyle(%d)", b)
	}
}

// ButtonBoxOverrides contains methods that are overridable.
type ButtonBoxOverrides struct {
}

func defaultButtonBoxOverrides(v *ButtonBox) ButtonBoxOverrides {
	return ButtonBoxOverrides{}
}

type ButtonBox struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ButtonBox)(nil)
	_ coreglib.Objector = (*ButtonBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ButtonBox, *ButtonBoxClass, ButtonBoxOverrides](
		GTypeButtonBox,
		initButtonBoxClass,
		wrapButtonBox,
		defaultButtonBoxOverrides,
	)
}

func initButtonBoxClass(gclass unsafe.Pointer, overrides ButtonBoxOverrides, classInitFunc func(*ButtonBoxClass)) {
	if classInitFunc != nil {
		class := (*ButtonBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapButtonBox(obj *coreglib.Object) *ButtonBox {
	return &ButtonBox{
		Box: Box{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalButtonBox(p uintptr) (interface{}, error) {
	return wrapButtonBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ButtonBoxClass: instance of this type is always passed by reference.
type ButtonBoxClass struct {
	*buttonBoxClass
}

// buttonBoxClass is the struct that's finalized.
type buttonBoxClass struct {
	native unsafe.Pointer
}

var GIRInfoButtonBoxClass = girepository.MustFind("Gtk", "ButtonBoxClass")
