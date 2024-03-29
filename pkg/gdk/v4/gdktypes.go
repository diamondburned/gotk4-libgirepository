// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"strings"
	"unsafe"

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
	GTypeAxisUse        = coreglib.Type(girepository.MustFind("Gdk", "AxisUse").RegisteredGType())
	GTypeGLError        = coreglib.Type(girepository.MustFind("Gdk", "GLError").RegisteredGType())
	GTypeGravity        = coreglib.Type(girepository.MustFind("Gdk", "Gravity").RegisteredGType())
	GTypeVulkanError    = coreglib.Type(girepository.MustFind("Gdk", "VulkanError").RegisteredGType())
	GTypeAxisFlags      = coreglib.Type(girepository.MustFind("Gdk", "AxisFlags").RegisteredGType())
	GTypeDragAction     = coreglib.Type(girepository.MustFind("Gdk", "DragAction").RegisteredGType())
	GTypeModifierType   = coreglib.Type(girepository.MustFind("Gdk", "ModifierType").RegisteredGType())
	GTypeContentFormats = coreglib.Type(girepository.MustFind("Gdk", "ContentFormats").RegisteredGType())
	GTypeRectangle      = coreglib.Type(girepository.MustFind("Gdk", "Rectangle").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAxisUse, F: marshalAxisUse},
		coreglib.TypeMarshaler{T: GTypeGLError, F: marshalGLError},
		coreglib.TypeMarshaler{T: GTypeGravity, F: marshalGravity},
		coreglib.TypeMarshaler{T: GTypeVulkanError, F: marshalVulkanError},
		coreglib.TypeMarshaler{T: GTypeAxisFlags, F: marshalAxisFlags},
		coreglib.TypeMarshaler{T: GTypeDragAction, F: marshalDragAction},
		coreglib.TypeMarshaler{T: GTypeModifierType, F: marshalModifierType},
		coreglib.TypeMarshaler{T: GTypeContentFormats, F: marshalContentFormats},
		coreglib.TypeMarshaler{T: GTypeRectangle, F: marshalRectangle},
	})
}

// ACTION_ALL defines all possible DND actions.
//
// This can be used in gdk.Drop.Status() messages when any drop can be accepted
// or a more specific drop method is not yet known.
const ACTION_ALL = 7

// CURRENT_TIME represents the current time, and can be used anywhere a time is
// expected.
const CURRENT_TIME = 0

// MODIFIER_MASK: mask covering all entries in GdkModifierType.
const MODIFIER_MASK = 469769999

// AxisUse defines how device axes are interpreted by GTK.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse C.gint

const (
	// AxisIgnore axis is ignored.
	AxisIgnore AxisUse = iota
	// AxisX axis is used as the x axis.
	AxisX
	// AxisY axis is used as the y axis.
	AxisY
	// AxisDeltaX axis is used as the scroll x delta.
	AxisDeltaX
	// AxisDeltaY axis is used as the scroll y delta.
	AxisDeltaY
	// AxisPressure axis is used for pressure information.
	AxisPressure
	// AxisXtilt axis is used for x tilt information.
	AxisXtilt
	// AxisYtilt axis is used for y tilt information.
	AxisYtilt
	// AxisWheel axis is used for wheel information.
	AxisWheel
	// AxisDistance axis is used for pen/tablet distance information.
	AxisDistance
	// AxisRotation axis is used for pen rotation information.
	AxisRotation
	// AxisSlider axis is used for pen slider information.
	AxisSlider
	// AxisLast: constant equal to the numerically highest axis value.
	AxisLast
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AxisUse.
func (a AxisUse) String() string {
	switch a {
	case AxisIgnore:
		return "Ignore"
	case AxisX:
		return "X"
	case AxisY:
		return "Y"
	case AxisDeltaX:
		return "DeltaX"
	case AxisDeltaY:
		return "DeltaY"
	case AxisPressure:
		return "Pressure"
	case AxisXtilt:
		return "Xtilt"
	case AxisYtilt:
		return "Ytilt"
	case AxisWheel:
		return "Wheel"
	case AxisDistance:
		return "Distance"
	case AxisRotation:
		return "Rotation"
	case AxisSlider:
		return "Slider"
	case AxisLast:
		return "Last"
	default:
		return fmt.Sprintf("AxisUse(%d)", a)
	}
}

// GLError: error enumeration for GdkGLContext.
type GLError C.gint

const (
	// GLErrorNotAvailable: openGL support is not available.
	GLErrorNotAvailable GLError = iota
	// GLErrorUnsupportedFormat: requested visual format is not supported.
	GLErrorUnsupportedFormat
	// GLErrorUnsupportedProfile: requested profile is not supported.
	GLErrorUnsupportedProfile
	// GLErrorCompilationFailed: shader compilation failed.
	GLErrorCompilationFailed
	// GLErrorLinkFailed: shader linking failed.
	GLErrorLinkFailed
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLError.
func (g GLError) String() string {
	switch g {
	case GLErrorNotAvailable:
		return "NotAvailable"
	case GLErrorUnsupportedFormat:
		return "UnsupportedFormat"
	case GLErrorUnsupportedProfile:
		return "UnsupportedProfile"
	case GLErrorCompilationFailed:
		return "CompilationFailed"
	case GLErrorLinkFailed:
		return "LinkFailed"
	default:
		return fmt.Sprintf("GLError(%d)", g)
	}
}

// Gravity defines the reference point of a surface and is used in PopupLayout.
type Gravity C.gint

const (
	// GravityNorthWest: reference point is at the top left corner.
	GravityNorthWest Gravity = 1
	// GravityNorth: reference point is in the middle of the top edge.
	GravityNorth Gravity = 2
	// GravityNorthEast: reference point is at the top right corner.
	GravityNorthEast Gravity = 3
	// GravityWest: reference point is at the middle of the left edge.
	GravityWest Gravity = 4
	// GravityCenter: reference point is at the center of the surface.
	GravityCenter Gravity = 5
	// GravityEast: reference point is at the middle of the right edge.
	GravityEast Gravity = 6
	// GravitySouthWest: reference point is at the lower left corner.
	GravitySouthWest Gravity = 7
	// GravitySouth: reference point is at the middle of the lower edge.
	GravitySouth Gravity = 8
	// GravitySouthEast: reference point is at the lower right corner.
	GravitySouthEast Gravity = 9
	// GravityStatic: reference point is at the top left corner of the surface
	// itself, ignoring window manager decorations.
	GravityStatic Gravity = 10
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Gravity.
func (g Gravity) String() string {
	switch g {
	case GravityNorthWest:
		return "NorthWest"
	case GravityNorth:
		return "North"
	case GravityNorthEast:
		return "NorthEast"
	case GravityWest:
		return "West"
	case GravityCenter:
		return "Center"
	case GravityEast:
		return "East"
	case GravitySouthWest:
		return "SouthWest"
	case GravitySouth:
		return "South"
	case GravitySouthEast:
		return "SouthEast"
	case GravityStatic:
		return "Static"
	default:
		return fmt.Sprintf("Gravity(%d)", g)
	}
}

// VulkanError: error enumeration for VulkanContext.
type VulkanError C.gint

const (
	// VulkanErrorUnsupported: vulkan is not supported on this backend or has
	// not been compiled in.
	VulkanErrorUnsupported VulkanError = iota
	// VulkanErrorNotAvailable: vulkan support is not available on this Surface.
	VulkanErrorNotAvailable
)

func marshalVulkanError(p uintptr) (interface{}, error) {
	return VulkanError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VulkanError.
func (v VulkanError) String() string {
	switch v {
	case VulkanErrorUnsupported:
		return "Unsupported"
	case VulkanErrorNotAvailable:
		return "NotAvailable"
	default:
		return fmt.Sprintf("VulkanError(%d)", v)
	}
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags C.guint

const (
	// AxisFlagX: x axis is present.
	AxisFlagX AxisFlags = 0b10
	// AxisFlagY: y axis is present.
	AxisFlagY AxisFlags = 0b100
	// AxisFlagDeltaX: scroll X delta axis is present.
	AxisFlagDeltaX AxisFlags = 0b1000
	// AxisFlagDeltaY: scroll Y delta axis is present.
	AxisFlagDeltaY AxisFlags = 0b10000
	// AxisFlagPressure: pressure axis is present.
	AxisFlagPressure AxisFlags = 0b100000
	// AxisFlagXtilt: x tilt axis is present.
	AxisFlagXtilt AxisFlags = 0b1000000
	// AxisFlagYtilt: y tilt axis is present.
	AxisFlagYtilt AxisFlags = 0b10000000
	// AxisFlagWheel: wheel axis is present.
	AxisFlagWheel AxisFlags = 0b100000000
	// AxisFlagDistance: distance axis is present.
	AxisFlagDistance AxisFlags = 0b1000000000
	// AxisFlagRotation z-axis rotation is present.
	AxisFlagRotation AxisFlags = 0b10000000000
	// AxisFlagSlider: slider axis is present.
	AxisFlagSlider AxisFlags = 0b100000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AxisFlags.
func (a AxisFlags) String() string {
	if a == 0 {
		return "AxisFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(157)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AxisFlagX:
			builder.WriteString("X|")
		case AxisFlagY:
			builder.WriteString("Y|")
		case AxisFlagDeltaX:
			builder.WriteString("DeltaX|")
		case AxisFlagDeltaY:
			builder.WriteString("DeltaY|")
		case AxisFlagPressure:
			builder.WriteString("Pressure|")
		case AxisFlagXtilt:
			builder.WriteString("Xtilt|")
		case AxisFlagYtilt:
			builder.WriteString("Ytilt|")
		case AxisFlagWheel:
			builder.WriteString("Wheel|")
		case AxisFlagDistance:
			builder.WriteString("Distance|")
		case AxisFlagRotation:
			builder.WriteString("Rotation|")
		case AxisFlagSlider:
			builder.WriteString("Slider|")
		default:
			builder.WriteString(fmt.Sprintf("AxisFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AxisFlags) Has(other AxisFlags) bool {
	return (a & other) == other
}

// DragAction: used in GdkDrop and GdkDrag to indicate the actions that the
// destination can and should do with the dropped data.
type DragAction C.guint

const (
	// ActionCopy: copy the data.
	ActionCopy DragAction = 0b1
	// ActionMove: move the data, i.e. first copy it, then delete it from the
	// source using the DELETE target of the X selection protocol.
	ActionMove DragAction = 0b10
	// ActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means, and is not supported on
	// all platforms.
	ActionLink DragAction = 0b100
	// ActionAsk: ask the user what to do with the data.
	ActionAsk DragAction = 0b1000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DragAction.
func (d DragAction) String() string {
	if d == 0 {
		return "DragAction(0)"
	}

	var builder strings.Builder
	builder.Grow(42)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case ActionCopy:
			builder.WriteString("Copy|")
		case ActionMove:
			builder.WriteString("Move|")
		case ActionLink:
			builder.WriteString("Link|")
		case ActionAsk:
			builder.WriteString("Ask|")
		default:
			builder.WriteString(fmt.Sprintf("DragAction(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DragAction) Has(other DragAction) bool {
	return (d & other) == other
}

// ModifierType flags to indicate the state of modifier keys and mouse buttons
// in events.
//
// Typical modifier keys are Shift, Control, Meta, Super, Hyper, Alt, Compose,
// Apple, CapsLock or ShiftLock.
//
// Note that GDK may add internal values to events which include values outside
// of this enumeration. Your code should preserve and ignore them. You can use
// GDK_MODIFIER_MASK to remove all private values.
type ModifierType C.guint

const (
	// ShiftMask: shift key.
	ShiftMask ModifierType = 0b1
	// LockMask: lock key (depending on the modifier mapping of the X server
	// this may either be CapsLock or ShiftLock).
	LockMask ModifierType = 0b10
	// ControlMask: control key.
	ControlMask ModifierType = 0b100
	// AltMask: fourth modifier key (it depends on the modifier mapping of the X
	// server which key is interpreted as this modifier, but normally it is the
	// Alt key).
	AltMask ModifierType = 0b1000
	// Button1Mask: first mouse button.
	Button1Mask ModifierType = 0b100000000
	// Button2Mask: second mouse button.
	Button2Mask ModifierType = 0b1000000000
	// Button3Mask: third mouse button.
	Button3Mask ModifierType = 0b10000000000
	// Button4Mask: fourth mouse button.
	Button4Mask ModifierType = 0b100000000000
	// Button5Mask: fifth mouse button.
	Button5Mask ModifierType = 0b1000000000000
	// SuperMask: super modifier.
	SuperMask ModifierType = 0b100000000000000000000000000
	// HyperMask: hyper modifier.
	HyperMask ModifierType = 0b1000000000000000000000000000
	// MetaMask: meta modifier.
	MetaMask ModifierType = 0b10000000000000000000000000000
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ModifierType.
func (m ModifierType) String() string {
	if m == 0 {
		return "ModifierType(0)"
	}

	var builder strings.Builder
	builder.Grow(127)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case ShiftMask:
			builder.WriteString("ShiftMask|")
		case LockMask:
			builder.WriteString("LockMask|")
		case ControlMask:
			builder.WriteString("ControlMask|")
		case AltMask:
			builder.WriteString("AltMask|")
		case Button1Mask:
			builder.WriteString("Button1Mask|")
		case Button2Mask:
			builder.WriteString("Button2Mask|")
		case Button3Mask:
			builder.WriteString("Button3Mask|")
		case Button4Mask:
			builder.WriteString("Button4Mask|")
		case Button5Mask:
			builder.WriteString("Button5Mask|")
		case SuperMask:
			builder.WriteString("SuperMask|")
		case HyperMask:
			builder.WriteString("HyperMask|")
		case MetaMask:
			builder.WriteString("MetaMask|")
		default:
			builder.WriteString(fmt.Sprintf("ModifierType(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m ModifierType) Has(other ModifierType) bool {
	return (m & other) == other
}

// ContentFormats: GdkContentFormats structure is used to advertise and
// negotiate the format of content.
//
// You will encounter GdkContentFormats when interacting with objects
// controlling operations that pass data between different widgets, window or
// application, like gdk.Drag, gdk.Drop, gdk.Clipboard or gdk.ContentProvider.
//
// GDK supports content in 2 forms: GType and mime type. Using GTypes is meant
// only for in-process content transfers. Mime types are meant to be used for
// data passing both in-process and out-of-process. The details of how data is
// passed is described in the documentation of the actual implementations. To
// transform between the two forms, gdk.ContentSerializer and
// gdk.ContentDeserializer are used.
//
// A GdkContentFormats describes a set of possible formats content can be
// exchanged in. It is assumed that this set is ordered. GTypes are more
// important than mime types. Order between different GTypes or mime types is
// the order they were added in, most important first. Functions that care about
// order, such as gdk.ContentFormats.Union(), will describe in their
// documentation how they interpret that order, though in general the order of
// the first argument is considered the primary order of the result, followed by
// the order of further arguments.
//
// For debugging purposes, the function gdk.ContentFormats.ToString() exists. It
// will print a comma-separated list of formats from most important to least
// important.
//
// GdkContentFormats is an immutable struct. After creation, you cannot change
// the types it represents. Instead, new GdkContentFormats have to be created.
// The gdk.ContentFormatsBuilder` structure is meant to help in this endeavor.
//
// An instance of this type is always passed by reference.
type ContentFormats struct {
	*contentFormats
}

// contentFormats is the struct that's finalized.
type contentFormats struct {
	native unsafe.Pointer
}

var GIRInfoContentFormats = girepository.MustFind("Gdk", "ContentFormats")

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ContentFormats{&contentFormats{(unsafe.Pointer)(b)}}, nil
}

// KeymapKey: GdkKeymapKey is a hardware key that can be mapped to a keyval.
//
// An instance of this type is always passed by reference.
type KeymapKey struct {
	*keymapKey
}

// keymapKey is the struct that's finalized.
type keymapKey struct {
	native unsafe.Pointer
}

var GIRInfoKeymapKey = girepository.MustFind("Gdk", "KeymapKey")

// NewKeymapKey creates a new KeymapKey instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewKeymapKey(keycode uint, group, level int) KeymapKey {
	var f0 C.guint // out
	f0 = C.guint(keycode)
	var f1 C.int // out
	f1 = C.int(group)
	var f2 C.int // out
	f2 = C.int(level)

	size := GIRInfoKeymapKey.StructSize()
	native := make([]byte, size)
	gextras.Sink(&native[0])

	offset0 := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr0 := (*C.guint)(unsafe.Add(unsafe.Pointer(&native[0]), offset0))
	*valptr0 = f0

	offset1 := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr1 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset1))
	*valptr1 = f1

	offset2 := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr2 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset2))
	*valptr2 = f2

	return *(*KeymapKey)(gextras.NewStructNative(unsafe.Pointer(&native[0])))
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) Keycode() uint {
	offset := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr := (*uint)(unsafe.Add(k.native, offset))
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	offset := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr := (*int)(unsafe.Add(k.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	offset := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr := (*int)(unsafe.Add(k.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) SetKeycode(keycode uint) {
	offset := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr := (*C.guint)(unsafe.Add(k.native, offset))
	*valptr = C.guint(keycode)
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) SetGroup(group int) {
	offset := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr := (*C.int)(unsafe.Add(k.native, offset))
	*valptr = C.int(group)
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) SetLevel(level int) {
	offset := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr := (*C.int)(unsafe.Add(k.native, offset))
	*valptr = C.int(level)
}

// Rectangle: GdkRectangle data type for representing rectangles.
//
// GdkRectangle is identical to cairo_rectangle_t. Together with Cairo’s
// cairo_region_t data type, these are the central types for representing sets
// of pixels.
//
// The intersection of two rectangles can be computed with
// gdk.Rectangle.Intersect(); to find the union of two rectangles use
// gdk.Rectangle.Union().
//
// The cairo_region_t type provided by Cairo is usually used for managing
// non-rectangular clipping of graphical operations.
//
// The Graphene library has a number of other data types for regions and volumes
// in 2D and 3D.
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native unsafe.Pointer
}

var GIRInfoRectangle = girepository.MustFind("Gdk", "Rectangle")

func marshalRectangle(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rectangle{&rectangle{(unsafe.Pointer)(b)}}, nil
}

// NewRectangle creates a new Rectangle instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewRectangle(x, y, width, height int) Rectangle {
	var f0 C.int // out
	f0 = C.int(x)
	var f1 C.int // out
	f1 = C.int(y)
	var f2 C.int // out
	f2 = C.int(width)
	var f3 C.int // out
	f3 = C.int(height)

	size := GIRInfoRectangle.StructSize()
	native := make([]byte, size)
	gextras.Sink(&native[0])

	offset0 := GIRInfoRectangle.StructFieldOffset("x")
	valptr0 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset0))
	*valptr0 = f0

	offset1 := GIRInfoRectangle.StructFieldOffset("y")
	valptr1 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset1))
	*valptr1 = f1

	offset2 := GIRInfoRectangle.StructFieldOffset("width")
	valptr2 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset2))
	*valptr2 = f2

	offset3 := GIRInfoRectangle.StructFieldOffset("height")
	valptr3 := (*C.int)(unsafe.Add(unsafe.Pointer(&native[0]), offset3))
	*valptr3 = f3

	return *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer(&native[0])))
}

// X: x coordinate of the top left corner.
func (r *Rectangle) X() int {
	offset := GIRInfoRectangle.StructFieldOffset("x")
	valptr := (*int)(unsafe.Add(r.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Y: y coordinate of the top left corner.
func (r *Rectangle) Y() int {
	offset := GIRInfoRectangle.StructFieldOffset("y")
	valptr := (*int)(unsafe.Add(r.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Width: width of the rectangle.
func (r *Rectangle) Width() int {
	offset := GIRInfoRectangle.StructFieldOffset("width")
	valptr := (*int)(unsafe.Add(r.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Height: height of the rectangle.
func (r *Rectangle) Height() int {
	offset := GIRInfoRectangle.StructFieldOffset("height")
	valptr := (*int)(unsafe.Add(r.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// X: x coordinate of the top left corner.
func (r *Rectangle) SetX(x int) {
	offset := GIRInfoRectangle.StructFieldOffset("x")
	valptr := (*C.int)(unsafe.Add(r.native, offset))
	*valptr = C.int(x)
}

// Y: y coordinate of the top left corner.
func (r *Rectangle) SetY(y int) {
	offset := GIRInfoRectangle.StructFieldOffset("y")
	valptr := (*C.int)(unsafe.Add(r.native, offset))
	*valptr = C.int(y)
}

// Width: width of the rectangle.
func (r *Rectangle) SetWidth(width int) {
	offset := GIRInfoRectangle.StructFieldOffset("width")
	valptr := (*C.int)(unsafe.Add(r.native, offset))
	*valptr = C.int(width)
}

// Height: height of the rectangle.
func (r *Rectangle) SetHeight(height int) {
	offset := GIRInfoRectangle.StructFieldOffset("height")
	valptr := (*C.int)(unsafe.Add(r.native, offset))
	*valptr = C.int(height)
}
