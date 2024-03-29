// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"strings"
	"unsafe"

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
	GTypeAnchorHints = coreglib.Type(girepository.MustFind("Gdk", "AnchorHints").RegisteredGType())
	GTypePopupLayout = coreglib.Type(girepository.MustFind("Gdk", "PopupLayout").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAnchorHints, F: marshalAnchorHints},
		coreglib.TypeMarshaler{T: GTypePopupLayout, F: marshalPopupLayout},
	})
}

// AnchorHints: positioning hints for aligning a surface relative to a
// rectangle.
//
// These hints determine how the surface should be positioned in the case that
// the surface would fall off-screen if placed in its ideal position.
//
// For example, GDK_ANCHOR_FLIP_X will replace GDK_GRAVITY_NORTH_WEST with
// GDK_GRAVITY_NORTH_EAST and vice versa if the surface extends beyond the left
// or right edges of the monitor.
//
// If GDK_ANCHOR_SLIDE_X is set, the surface can be shifted horizontally to fit
// on-screen. If GDK_ANCHOR_RESIZE_X is set, the surface can be shrunken
// horizontally to fit.
//
// In general, when multiple flags are set, flipping should take precedence over
// sliding, which should take precedence over resizing.
type AnchorHints C.guint

const (
	// AnchorFlipX: allow flipping anchors horizontally.
	AnchorFlipX AnchorHints = 0b1
	// AnchorFlipY: allow flipping anchors vertically.
	AnchorFlipY AnchorHints = 0b10
	// AnchorSlideX: allow sliding surface horizontally.
	AnchorSlideX AnchorHints = 0b100
	// AnchorSlideY: allow sliding surface vertically.
	AnchorSlideY AnchorHints = 0b1000
	// AnchorResizeX: allow resizing surface horizontally.
	AnchorResizeX AnchorHints = 0b10000
	// AnchorResizeY: allow resizing surface vertically.
	AnchorResizeY AnchorHints = 0b100000
	// AnchorFlip: allow flipping anchors on both axes.
	AnchorFlip AnchorHints = 0b11
	// AnchorSlide: allow sliding surface on both axes.
	AnchorSlide AnchorHints = 0b1100
	// AnchorResize: allow resizing surface on both axes.
	AnchorResize AnchorHints = 0b110000
)

func marshalAnchorHints(p uintptr) (interface{}, error) {
	return AnchorHints(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AnchorHints.
func (a AnchorHints) String() string {
	if a == 0 {
		return "AnchorHints(0)"
	}

	var builder strings.Builder
	builder.Grow(113)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AnchorFlipX:
			builder.WriteString("FlipX|")
		case AnchorFlipY:
			builder.WriteString("FlipY|")
		case AnchorSlideX:
			builder.WriteString("SlideX|")
		case AnchorSlideY:
			builder.WriteString("SlideY|")
		case AnchorResizeX:
			builder.WriteString("ResizeX|")
		case AnchorResizeY:
			builder.WriteString("ResizeY|")
		case AnchorFlip:
			builder.WriteString("Flip|")
		case AnchorSlide:
			builder.WriteString("Slide|")
		case AnchorResize:
			builder.WriteString("Resize|")
		default:
			builder.WriteString(fmt.Sprintf("AnchorHints(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AnchorHints) Has(other AnchorHints) bool {
	return (a & other) == other
}

// PopupLayout: GdkPopupLayout struct contains information that is necessary
// position a gdk.Popup relative to its parent.
//
// The positioning requires a negotiation with the windowing system, since it
// depends on external constraints, such as the position of the parent surface,
// and the screen dimensions.
//
// The basic ingredients are a rectangle on the parent surface, and the anchor
// on both that rectangle and the popup. The anchors specify a side or corner to
// place next to each other.
//
// !Popup anchors (popup-anchors.png)
//
// For cases where placing the anchors next to each other would make the popup
// extend offscreen, the layout includes some hints for how to resolve this
// problem. The hints may suggest to flip the anchor position to the other side,
// or to 'slide' the popup along a side, or to resize it.
//
// !Flipping popups (popup-flip.png)
//
// !Sliding popups (popup-slide.png)
//
// These hints may be combined.
//
// Ultimatively, it is up to the windowing system to determine the position and
// size of the popup. You can learn about the result by calling
// gdk.Popup.GetPositionX(), gdk.Popup.GetPositionY(), gdk.Popup.GetRectAnchor()
// and gdk.Popup.GetSurfaceAnchor() after the popup has been presented. This can
// be used to adjust the rendering. For example, gtk.Popover changes its arrow
// position accordingly. But you have to be careful avoid changing the size of
// the popover, or it has to be presented again.
//
// An instance of this type is always passed by reference.
type PopupLayout struct {
	*popupLayout
}

// popupLayout is the struct that's finalized.
type popupLayout struct {
	native unsafe.Pointer
}

var GIRInfoPopupLayout = girepository.MustFind("Gdk", "PopupLayout")

func marshalPopupLayout(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PopupLayout{&popupLayout{(unsafe.Pointer)(b)}}, nil
}
