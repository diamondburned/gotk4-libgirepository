// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_formats_get_type()), F: marshalContentFormats},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// ContentFormats: the `GdkContentFormats` structure is used to advertise and
// negotiate the format of content.
//
// You will encounter `GdkContentFormats` when interacting with objects
// controlling operations that pass data between different widgets, window or
// application, like [class@Gdk.Drag], [class@Gdk.Drop], [class@Gdk.Clipboard]
// or [class@Gdk.ContentProvider].
//
// GDK supports content in 2 forms: `GType` and mime type. Using `GTypes` is
// meant only for in-process content transfers. Mime types are meant to be used
// for data passing both in-process and out-of-process. The details of how data
// is passed is described in the documentation of the actual implementations. To
// transform between the two forms, [class@Gdk.ContentSerializer] and
// [class@Gdk.ContentDeserializer] are used.
//
// A `GdkContentFormats` describes a set of possible formats content can be
// exchanged in. It is assumed that this set is ordered. `GTypes` are more
// important than mime types. Order between different `GTypes` or mime types is
// the order they were added in, most important first. Functions that care about
// order, such as [method@Gdk.ContentFormats.union], will describe in their
// documentation how they interpret that order, though in general the order of
// the first argument is considered the primary order of the result, followed by
// the order of further arguments.
//
// For debugging purposes, the function [method@Gdk.ContentFormats.to_string]
// exists. It will print a comma-separated list of formats from most important
// to least important.
//
// `GdkContentFormats` is an immutable struct. After creation, you cannot change
// the types it represents. Instead, new `GdkContentFormats` have to be created.
// The [struct@Gdk.ContentFormatsBuilder]` structure is meant to help in this
// endeavor.
type ContentFormats struct {
	native C.GdkContentFormats
}

// WrapContentFormats wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapContentFormats(ptr unsafe.Pointer) *ContentFormats {
	if ptr == nil {
		return nil
	}

	return (*ContentFormats)(ptr)
}

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapContentFormats(unsafe.Pointer(b)), nil
}

// NewContentFormats constructs a struct ContentFormats.
func NewContentFormats() *ContentFormats {
	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_new(arg1, arg2)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// NewContentFormatsForGType constructs a struct ContentFormats.
func NewContentFormatsForGType(typ externglib.Type) *ContentFormats {
	var arg1 C.GType

	arg1 = C.GType(typ)

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_new_for_gtype(arg1)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (c *ContentFormats) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// ContainGType checks if a given `GType` is part of the given @formats.
func (f *ContentFormats) ContainGType(typ externglib.Type) bool {
	var arg0 *C.GdkContentFormats
	var arg1 C.GType

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = C.GType(typ)

	var cret C.gboolean
	var goret bool

	cret = C.gdk_content_formats_contain_gtype(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// ContainMIMEType checks if a given mime type is part of the given @formats.
func (f *ContentFormats) ContainMIMEType(mimeType string) bool {
	var arg0 *C.GdkContentFormats
	var arg1 *C.char

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.gdk_content_formats_contain_mime_type(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// GTypes gets the `GTypes` included in @formats.
//
// Note that @formats may not contain any #GTypes, in particular when they are
// empty. In that case nil will be returned.
func (f *ContentFormats) GTypes() []externglib.Type {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	var cret *C.GType
	var arg1 *C.gsize
	var goret []externglib.Type

	cret = C.gdk_content_formats_get_gtypes(arg0, arg1)

	goret = make([]externglib.Type, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (C.GType)(ptr.Add(unsafe.Pointer(cret), i))
		goret[i] = externglib.Type(src)
	}

	return ret1, goret
}

// MIMETypes gets the mime types included in @formats.
//
// Note that @formats may not contain any mime types, in particular when they
// are empty. In that case nil will be returned.
func (f *ContentFormats) MIMETypes() []string {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	var cret **C.char
	var arg1 *C.gsize
	var goret []string

	cret = C.gdk_content_formats_get_mime_types(arg0, arg1)

	goret = make([]string, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		goret[i] = C.GoString(src)
	}

	return ret1, goret
}

// Match checks if @first and @second have any matching formats.
func (f *ContentFormats) Match(second *ContentFormats) bool {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gdk_content_formats_match(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// MatchGType finds the first `GType` from @first that is also contained in
// @second.
//
// If no matching `GType` is found, G_TYPE_INVALID is returned.
func (f *ContentFormats) MatchGType(second *ContentFormats) externglib.Type {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	var cret C.GType
	var goret externglib.Type

	cret = C.gdk_content_formats_match_gtype(arg0, arg1)

	goret = externglib.Type(cret)

	return goret
}

// MatchMIMEType finds the first mime type from @first that is also contained in
// @second.
//
// If no matching mime type is found, nil is returned.
func (f *ContentFormats) MatchMIMEType(second *ContentFormats) string {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	var cret *C.char
	var goret string

	cret = C.gdk_content_formats_match_mime_type(arg0, arg1)

	goret = C.GoString(cret)

	return goret
}

// Print prints the given @formats into a string for human consumption.
//
// This is meant for debugging and logging.
//
// The form of the representation may change at any time and is not guaranteed
// to stay identical.
func (f *ContentFormats) Print(string *glib.String) {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GString

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gdk_content_formats_print(arg0, arg1)
}

// Ref increases the reference count of a `GdkContentFormats` by one.
func (f *ContentFormats) Ref() *ContentFormats {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_ref(arg0)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// String prints the given @formats into a human-readable string.
//
// This is a small wrapper around [method@Gdk.ContentFormats.print] to help when
// debugging.
func (f *ContentFormats) String() string {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.char)
	var goret string

	cret = C.gdk_content_formats_to_string(arg0)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// Union: append all missing types from @second to @first, in the order they had
// in @second.
func (f *ContentFormats) Union(second *ContentFormats) *ContentFormats {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_union(arg0, arg1)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// UnionDeserializeGTypes: add GTypes for mime types in @formats for which
// deserializers are registered.
func (f *ContentFormats) UnionDeserializeGTypes() *ContentFormats {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_union_deserialize_gtypes(arg0)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// UnionDeserializeMIMETypes: add mime types for GTypes in @formats for which
// deserializers are registered.
func (f *ContentFormats) UnionDeserializeMIMETypes() *ContentFormats {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_union_deserialize_mime_types(arg0)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// UnionSerializeGTypes: add GTypes for the mime types in @formats for which
// serializers are registered.
func (f *ContentFormats) UnionSerializeGTypes() *ContentFormats {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_union_serialize_gtypes(arg0)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// UnionSerializeMIMETypes: add mime types for GTypes in @formats for which
// serializers are registered.
func (f *ContentFormats) UnionSerializeMIMETypes() *ContentFormats {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	cret := new(C.GdkContentFormats)
	var goret *ContentFormats

	cret = C.gdk_content_formats_union_serialize_mime_types(arg0)

	goret = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Unref decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) Unref() {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_unref(arg0)
}

// KeymapKey: a `GdkKeymapKey` is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	native C.GdkKeymapKey
}

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	if ptr == nil {
		return nil
	}

	return (*KeymapKey)(ptr)
}

func marshalKeymapKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapKeymapKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Keycode gets the field inside the struct.
func (k *KeymapKey) Keycode() uint {
	var v uint
	v = uint(k.native.keycode)
	return v
}

// Group gets the field inside the struct.
func (k *KeymapKey) Group() int {
	var v int
	v = int(k.native.group)
	return v
}

// Level gets the field inside the struct.
func (k *KeymapKey) Level() int {
	var v int
	v = int(k.native.level)
	return v
}

// Rectangle: a `GdkRectangle` data type for representing rectangles.
//
// `GdkRectangle` is identical to `cairo_rectangle_t`. Together with Cairo’s
// `cairo_region_t` data type, these are the central types for representing sets
// of pixels.
//
// The intersection of two rectangles can be computed with
// [method@Gdk.Rectangle.intersect]; to find the union of two rectangles use
// [method@Gdk.Rectangle.union].
//
// The `cairo_region_t` type provided by Cairo is usually used for managing
// non-rectangular clipping of graphical operations.
//
// The Graphene library has a number of other data types for regions and volumes
// in 2D and 3D.
type Rectangle struct {
	native C.GdkRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	if ptr == nil {
		return nil
	}

	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRectangle(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// X gets the field inside the struct.
func (r *Rectangle) X() int {
	var v int
	v = int(r.native.x)
	return v
}

// Y gets the field inside the struct.
func (r *Rectangle) Y() int {
	var v int
	v = int(r.native.y)
	return v
}

// Width gets the field inside the struct.
func (r *Rectangle) Width() int {
	var v int
	v = int(r.native.width)
	return v
}

// Height gets the field inside the struct.
func (r *Rectangle) Height() int {
	var v int
	v = int(r.native.height)
	return v
}

// ContainsPoint returns UE if @rect contains the point described by @x and @y.
func (r *Rectangle) ContainsPoint(x int, y int) bool {
	var arg0 *C.GdkRectangle
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	arg1 = C.int(x)
	arg2 = C.int(y)

	var cret C.gboolean
	var goret bool

	cret = C.gdk_rectangle_contains_point(arg0, arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// Equal checks if the two given rectangles are equal.
func (r *Rectangle) Equal(rect2 *Rectangle) bool {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect2.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gdk_rectangle_equal(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Intersect calculates the intersection of two rectangles.
//
// It is allowed for @dest to be the same as either @src1 or @src2. If the
// rectangles do not intersect, @dest’s width and height is set to 0 and its x
// and y values are undefined. If you are only interested in whether the
// rectangles intersect, but not in the intersecting area itself, pass nil for
// @dest.
func (s *Rectangle) Intersect(src2 *Rectangle) (dest *Rectangle, ok bool) {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	arg2 := new(C.GdkRectangle)
	var ret2 *Rectangle
	var cret C.gboolean
	var goret bool

	cret = C.gdk_rectangle_intersect(arg0, arg1, arg2)

	ret2 = WrapRectangle(unsafe.Pointer(arg2))
	if cret {
		goret = true
	}

	return ret2, goret
}

// Union calculates the union of two rectangles.
//
// The union of rectangles @src1 and @src2 is the smallest rectangle which
// includes both @src1 and @src2 within it. It is allowed for @dest to be the
// same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Union(src2 *Rectangle) *Rectangle {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	arg2 := new(C.GdkRectangle)
	var ret2 *Rectangle

	C.gdk_rectangle_union(arg0, arg1, arg2)

	ret2 = WrapRectangle(unsafe.Pointer(arg2))

	return ret2
}
