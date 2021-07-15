// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_axis_use_get_type()), F: marshalAxisUse},
		{T: externglib.Type(C.gdk_gl_error_get_type()), F: marshalGLError},
		{T: externglib.Type(C.gdk_gravity_get_type()), F: marshalGravity},
		{T: externglib.Type(C.gdk_vulkan_error_get_type()), F: marshalVulkanError},
		{T: externglib.Type(C.gdk_axis_flags_get_type()), F: marshalAxisFlags},
		{T: externglib.Type(C.gdk_drag_action_get_type()), F: marshalDragAction},
		{T: externglib.Type(C.gdk_modifier_type_get_type()), F: marshalModifierType},
		{T: externglib.Type(C.gdk_content_formats_get_type()), F: marshalContentFormats},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// AxisUse defines how device axes are interpreted by GTK.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse int

const (
	// Ignore axis is ignored.
	AxisUseIgnore AxisUse = iota
	// X axis is used as the x axis.
	AxisUseX
	// Y axis is used as the y axis.
	AxisUseY
	// DeltaX axis is used as the scroll x delta
	AxisUseDeltaX
	// DeltaY axis is used as the scroll y delta
	AxisUseDeltaY
	// Pressure axis is used for pressure information.
	AxisUsePressure
	// Xtilt axis is used for x tilt information.
	AxisUseXtilt
	// Ytilt axis is used for y tilt information.
	AxisUseYtilt
	// Wheel axis is used for wheel information.
	AxisUseWheel
	// Distance axis is used for pen/tablet distance information
	AxisUseDistance
	// Rotation axis is used for pen rotation information
	AxisUseRotation
	// Slider axis is used for pen slider information
	AxisUseSlider
	// Last: constant equal to the numerically highest axis value.
	AxisUseLast
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GLError: error enumeration for GdkGLContext.
type GLError int

const (
	// NotAvailable: openGL support is not available
	GLErrorNotAvailable GLError = iota
	// UnsupportedFormat: requested visual format is not supported
	GLErrorUnsupportedFormat
	// UnsupportedProfile: requested profile is not supported
	GLErrorUnsupportedProfile
	// CompilationFailed: shader compilation failed
	GLErrorCompilationFailed
	// LinkFailed: shader linking failed
	GLErrorLinkFailed
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Gravity defines the reference point of a surface and is used in PopupLayout.
type Gravity int

const (
	// NorthWest: reference point is at the top left corner.
	GravityNorthWest Gravity = 1
	// North: reference point is in the middle of the top edge.
	GravityNorth Gravity = 2
	// NorthEast: reference point is at the top right corner.
	GravityNorthEast Gravity = 3
	// West: reference point is at the middle of the left edge.
	GravityWest Gravity = 4
	// Center: reference point is at the center of the surface.
	GravityCenter Gravity = 5
	// East: reference point is at the middle of the right edge.
	GravityEast Gravity = 6
	// SouthWest: reference point is at the lower left corner.
	GravitySouthWest Gravity = 7
	// South: reference point is at the middle of the lower edge.
	GravitySouth Gravity = 8
	// SouthEast: reference point is at the lower right corner.
	GravitySouthEast Gravity = 9
	// Static: reference point is at the top left corner of the surface itself,
	// ignoring window manager decorations.
	GravityStatic Gravity = 10
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// VulkanError: error enumeration for VulkanContext.
type VulkanError int

const (
	// Unsupported: vulkan is not supported on this backend or has not been
	// compiled in.
	VulkanErrorUnsupported VulkanError = iota
	// NotAvailable: vulkan support is not available on this Surface
	VulkanErrorNotAvailable
)

func marshalVulkanError(p uintptr) (interface{}, error) {
	return VulkanError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags int

const (
	// AxisFlagsX: x axis is present
	AxisFlagsX AxisFlags = 0b10
	// AxisFlagsY: y axis is present
	AxisFlagsY AxisFlags = 0b100
	// AxisFlagsDeltaX: scroll X delta axis is present
	AxisFlagsDeltaX AxisFlags = 0b1000
	// AxisFlagsDeltaY: scroll Y delta axis is present
	AxisFlagsDeltaY AxisFlags = 0b10000
	// AxisFlagsPressure: pressure axis is present
	AxisFlagsPressure AxisFlags = 0b100000
	// AxisFlagsXtilt: x tilt axis is present
	AxisFlagsXtilt AxisFlags = 0b1000000
	// AxisFlagsYtilt: y tilt axis is present
	AxisFlagsYtilt AxisFlags = 0b10000000
	// AxisFlagsWheel: wheel axis is present
	AxisFlagsWheel AxisFlags = 0b100000000
	// AxisFlagsDistance: distance axis is present
	AxisFlagsDistance AxisFlags = 0b1000000000
	// AxisFlagsRotation z-axis rotation is present
	AxisFlagsRotation AxisFlags = 0b10000000000
	// AxisFlagsSlider: slider axis is present
	AxisFlagsSlider AxisFlags = 0b100000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragAction: used in GdkDrop and GdkDrag to indicate the actions that the
// destination can and should do with the dropped data.
type DragAction int

const (
	// DragActionCopy: copy the data.
	DragActionCopy DragAction = 0b1
	// DragActionMove: move the data, i.e. first copy it, then delete it from
	// the source using the DELETE target of the X selection protocol.
	DragActionMove DragAction = 0b10
	// DragActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means, and is not supported on
	// all platforms.
	DragActionLink DragAction = 0b100
	// DragActionAsk: ask the user what to do with the data.
	DragActionAsk DragAction = 0b1000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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
type ModifierType int

const (
	// ModifierTypeShiftMask: shift key.
	ModifierTypeShiftMask ModifierType = 0b1
	// ModifierTypeLockMask: lock key (depending on the modifier mapping of the
	// X server this may either be CapsLock or ShiftLock).
	ModifierTypeLockMask ModifierType = 0b10
	// ModifierTypeControlMask: control key.
	ModifierTypeControlMask ModifierType = 0b100
	// ModifierTypeAltMask: fourth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier, but
	// normally it is the Alt key).
	ModifierTypeAltMask ModifierType = 0b1000
	// ModifierTypeButton1Mask: first mouse button.
	ModifierTypeButton1Mask ModifierType = 0b100000000
	// ModifierTypeButton2Mask: second mouse button.
	ModifierTypeButton2Mask ModifierType = 0b1000000000
	// ModifierTypeButton3Mask: third mouse button.
	ModifierTypeButton3Mask ModifierType = 0b10000000000
	// ModifierTypeButton4Mask: fourth mouse button.
	ModifierTypeButton4Mask ModifierType = 0b100000000000
	// ModifierTypeButton5Mask: fifth mouse button.
	ModifierTypeButton5Mask ModifierType = 0b1000000000000
	// ModifierTypeSuperMask: super modifier
	ModifierTypeSuperMask ModifierType = 0b100000000000000000000000000
	// ModifierTypeHyperMask: hyper modifier
	ModifierTypeHyperMask ModifierType = 0b1000000000000000000000000000
	// ModifierTypeMetaMask: meta modifier
	ModifierTypeMetaMask ModifierType = 0b10000000000000000000000000000
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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
type ContentFormats struct {
	nocopy gextras.NoCopy
	native *C.GdkContentFormats
}

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &ContentFormats{native: (*C.GdkContentFormats)(unsafe.Pointer(b))}, nil
}

// NewContentFormats constructs a struct ContentFormats.
func NewContentFormats(mimeTypes []string) *ContentFormats {
	var _arg1 **C.char
	var _arg2 C.guint
	var _cret *C.GdkContentFormats // in

	_arg2 = (C.guint)(len(mimeTypes))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(mimeTypes)) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice((**C.char)(_arg1), len(mimeTypes))
		for i := range mimeTypes {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(mimeTypes[i])))
		}
	}

	_cret = C.gdk_content_formats_new(_arg1, _arg2)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// NewContentFormatsForGType constructs a struct ContentFormats.
func NewContentFormatsForGType(typ externglib.Type) *ContentFormats {
	var _arg1 C.GType              // out
	var _cret *C.GdkContentFormats // in

	_arg1 = C.GType(typ)

	_cret = C.gdk_content_formats_new_for_gtype(_arg1)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// ContainGType checks if a given GType is part of the given formats.
func (formats *ContentFormats) ContainGType(typ externglib.Type) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 C.GType              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	_arg1 = C.GType(typ)

	_cret = C.gdk_content_formats_contain_gtype(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainMIMEType checks if a given mime type is part of the given formats.
func (formats *ContentFormats) ContainMIMEType(mimeType string) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.char              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))

	_cret = C.gdk_content_formats_contain_mime_type(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Match checks if first and second have any matching formats.
func (first *ContentFormats) Match(second *ContentFormats) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchGType finds the first GType from first that is also contained in second.
//
// If no matching GType is found, G_TYPE_INVALID is returned.
func (first *ContentFormats) MatchGType(second *ContentFormats) externglib.Type {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.GType              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match_gtype(_arg0, _arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// MatchMIMEType finds the first mime type from first that is also contained in
// second.
//
// If no matching mime type is found, NULL is returned.
func (first *ContentFormats) MatchMIMEType(second *ContentFormats) string {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match_mime_type(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Ref increases the reference count of a GdkContentFormats by one.
func (formats *ContentFormats) ref() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_ref(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// String prints the given formats into a human-readable string.
//
// This is a small wrapper around gdk.ContentFormats.Print() to help when
// debugging.
func (formats *ContentFormats) String() string {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Union: append all missing types from second to first, in the order they had
// in second.
func (first *ContentFormats) Union(second *ContentFormats) *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_union(_arg0, _arg1)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// UnionDeserializeGTypes: add GTypes for mime types in formats for which
// deserializers are registered.
func (formats *ContentFormats) UnionDeserializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_deserialize_gtypes(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// UnionDeserializeMIMETypes: add mime types for GTypes in formats for which
// deserializers are registered.
func (formats *ContentFormats) UnionDeserializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_deserialize_mime_types(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// UnionSerializeGTypes: add GTypes for the mime types in formats for which
// serializers are registered.
func (formats *ContentFormats) UnionSerializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_serialize_gtypes(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// UnionSerializeMIMETypes: add mime types for GTypes in formats for which
// serializers are registered.
func (formats *ContentFormats) UnionSerializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_serialize_mime_types(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// Unref decreases the reference count of a GdkContentFormats by one.
//
// If the resulting reference count is zero, frees the formats.
func (formats *ContentFormats) unref() {
	var _arg0 *C.GdkContentFormats // out

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	C.gdk_content_formats_unref(_arg0)
}

// KeymapKey: GdkKeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	nocopy gextras.NoCopy
	native *C.GdkKeymapKey
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) Keycode() uint {
	var v uint // out
	v = uint(k.native.keycode)
	return v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	var v int // out
	v = int(k.native.group)
	return v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	var v int // out
	v = int(k.native.level)
	return v
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
type Rectangle struct {
	nocopy gextras.NoCopy
	native *C.GdkRectangle
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Rectangle{native: (*C.GdkRectangle)(unsafe.Pointer(b))}, nil
}

// X: x coordinate of the top left corner
func (r *Rectangle) X() int {
	var v int // out
	v = int(r.native.x)
	return v
}

// Y: y coordinate of the top left corner
func (r *Rectangle) Y() int {
	var v int // out
	v = int(r.native.y)
	return v
}

// Width: width of the rectangle
func (r *Rectangle) Width() int {
	var v int // out
	v = int(r.native.width)
	return v
}

// Height: height of the rectangle
func (r *Rectangle) Height() int {
	var v int // out
	v = int(r.native.height)
	return v
}

// ContainsPoint returns UE if rect contains the point described by x and y.
func (rect *Rectangle) ContainsPoint(x int, y int) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect)))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gdk_rectangle_contains_point(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Equal checks if the two given rectangles are equal.
func (rect1 *Rectangle) Equal(rect2 *Rectangle) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect2)))

	_cret = C.gdk_rectangle_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Intersect calculates the intersection of two rectangles.
//
// It is allowed for dest to be the same as either src1 or src2. If the
// rectangles do not intersect, dest’s width and height is set to 0 and its x
// and y values are undefined. If you are only interested in whether the
// rectangles intersect, but not in the intersecting area itself, pass NULL for
// dest.
func (src1 *Rectangle) Intersect(src2 *Rectangle) (Rectangle, bool) {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	_cret = C.gdk_rectangle_intersect(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out
	var _ok bool        // out

	_dest = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _dest, _ok
}

// Union calculates the union of two rectangles.
//
// The union of rectangles src1 and src2 is the smallest rectangle which
// includes both src1 and src2 within it. It is allowed for dest to be the same
// as either src1 or src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (src1 *Rectangle) Union(src2 *Rectangle) Rectangle {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	C.gdk_rectangle_union(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out

	_dest = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _dest
}
