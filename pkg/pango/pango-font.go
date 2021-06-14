// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_stretch_get_type()), F: marshalStretch},
		{T: externglib.Type(C.pango_style_get_type()), F: marshalStyle},
		{T: externglib.Type(C.pango_variant_get_type()), F: marshalVariant},
		{T: externglib.Type(C.pango_weight_get_type()), F: marshalWeight},
		{T: externglib.Type(C.pango_font_mask_get_type()), F: marshalFontMask},
		{T: externglib.Type(C.pango_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_font_face_get_type()), F: marshalFontFace},
		{T: externglib.Type(C.pango_font_family_get_type()), F: marshalFontFamily},
		{T: externglib.Type(C.pango_font_description_get_type()), F: marshalFontDescription},
		{T: externglib.Type(C.pango_font_metrics_get_type()), F: marshalFontMetrics},
	})
}

// Stretch: an enumeration specifying the width of the font relative to other
// designs within a family.
type Stretch int

const (
	// StretchUltraCondensed: ultra condensed width
	StretchUltraCondensed Stretch = 0
	// StretchExtraCondensed: extra condensed width
	StretchExtraCondensed Stretch = 1
	// StretchCondensed: condensed width
	StretchCondensed Stretch = 2
	// StretchSemiCondensed: semi condensed width
	StretchSemiCondensed Stretch = 3
	// StretchNormal: the normal width
	StretchNormal Stretch = 4
	// StretchSemiExpanded: semi expanded width
	StretchSemiExpanded Stretch = 5
	// StretchExpanded: expanded width
	StretchExpanded Stretch = 6
	// StretchExtraExpanded: extra expanded width
	StretchExtraExpanded Stretch = 7
	// StretchUltraExpanded: ultra expanded width
	StretchUltraExpanded Stretch = 8
)

func marshalStretch(p uintptr) (interface{}, error) {
	return Stretch(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Style: an enumeration specifying the various slant styles possible for a
// font.
type Style int

const (
	// StyleNormal: the font is upright.
	StyleNormal Style = 0
	// StyleOblique: the font is slanted, but in a roman style.
	StyleOblique Style = 1
	// StyleItalic: the font is slanted in an italic style.
	StyleItalic Style = 2
)

func marshalStyle(p uintptr) (interface{}, error) {
	return Style(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Variant: an enumeration specifying capitalization variant of the font.
type Variant int

const (
	// VariantNormal: a normal font.
	VariantNormal Variant = 0
	// VariantSmallCaps: a font with the lower case characters replaced by
	// smaller variants of the capital characters.
	VariantSmallCaps Variant = 1
)

func marshalVariant(p uintptr) (interface{}, error) {
	return Variant(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Weight: an enumeration specifying the weight (boldness) of a font.
//
// This is a numerical value ranging from 100 to 1000, but there are some
// predefined values.
type Weight int

const (
	// WeightThin: the thin weight (= 100; Since: 1.24)
	WeightThin Weight = 100
	// WeightUltralight: the ultralight weight (= 200)
	WeightUltralight Weight = 200
	// WeightLight: the light weight (= 300)
	WeightLight Weight = 300
	// WeightSemilight: the semilight weight (= 350; Since: 1.36.7)
	WeightSemilight Weight = 350
	// WeightBook: the book weight (= 380; Since: 1.24)
	WeightBook Weight = 380
	// WeightNormal: the default weight (= 400)
	WeightNormal Weight = 400
	// WeightMedium: the normal weight (= 500; Since: 1.24)
	WeightMedium Weight = 500
	// WeightSemibold: the semibold weight (= 600)
	WeightSemibold Weight = 600
	// WeightBold: the bold weight (= 700)
	WeightBold Weight = 700
	// WeightUltrabold: the ultrabold weight (= 800)
	WeightUltrabold Weight = 800
	// WeightHeavy: the heavy weight (= 900)
	WeightHeavy Weight = 900
	// WeightUltraheavy: the ultraheavy weight (= 1000; Since: 1.24)
	WeightUltraheavy Weight = 1000
)

func marshalWeight(p uintptr) (interface{}, error) {
	return Weight(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontMask: the bits in a `PangoFontMask` correspond to the set fields in a
// `PangoFontDescription`.
type FontMask int

const (
	// FontMaskFamily: the font family is specified.
	FontMaskFamily FontMask = 1
	// FontMaskStyle: the font style is specified.
	FontMaskStyle FontMask = 2
	// FontMaskVariant: the font variant is specified.
	FontMaskVariant FontMask = 4
	// FontMaskWeight: the font weight is specified.
	FontMaskWeight FontMask = 8
	// FontMaskStretch: the font stretch is specified.
	FontMaskStretch FontMask = 16
	// FontMaskSize: the font size is specified.
	FontMaskSize FontMask = 32
	// FontMaskGravity: the font gravity is specified (Since: 1.16.)
	FontMaskGravity FontMask = 64
	// FontMaskVariations: openType font variations are specified (Since: 1.42)
	FontMaskVariations FontMask = 128
)

func marshalFontMask(p uintptr) (interface{}, error) {
	return FontMask(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Font: a `PangoFont` is used to represent a font in a
// rendering-system-independent manner.
type Font interface {
	gextras.Objector

	// HasChar returns whether the font provides a glyph for this character.
	//
	// Returns true if @font can render @wc
	HasChar(wc uint32) bool
}

// font implements the Font class.
type font struct {
	gextras.Objector
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return font{
		Objector: obj,
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// HasChar returns whether the font provides a glyph for this character.
//
// Returns true if @font can render @wc
func (f font) HasChar(wc uint32) bool {
	var _arg0 *C.PangoFont // out
	var _arg1 C.gunichar   // out

	_arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	_arg1 = C.gunichar(wc)

	var _cret C.gboolean // in

	_cret = C.pango_font_has_char(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FontFace: a `PangoFontFace` is used to represent a group of fonts with the
// same family, slant, weight, and width, but varying sizes.
type FontFace interface {
	gextras.Objector

	// FaceName gets a name representing the style of this face among the
	// different faces in the `PangoFontFamily` for the face. The name is
	// suitable for displaying to users.
	FaceName() string
	// IsSynthesized returns whether a `PangoFontFace` is synthesized by the
	// underlying font rendering engine from another face, perhaps by shearing,
	// emboldening, or lightening it.
	IsSynthesized() bool
	// ListSizes: list the available sizes for a font.
	//
	// This is only applicable to bitmap fonts. For scalable fonts, stores nil
	// at the location pointed to by @sizes and 0 at the location pointed to by
	// @n_sizes. The sizes returned are in Pango units and are sorted in
	// ascending order.
	ListSizes() []int
}

// fontFace implements the FontFace class.
type fontFace struct {
	gextras.Objector
}

var _ FontFace = (*fontFace)(nil)

// WrapFontFace wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontFace(obj *externglib.Object) FontFace {
	return fontFace{
		Objector: obj,
	}
}

func marshalFontFace(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontFace(obj), nil
}

// FaceName gets a name representing the style of this face among the
// different faces in the `PangoFontFamily` for the face. The name is
// suitable for displaying to users.
func (f fontFace) FaceName() string {
	var _arg0 *C.PangoFontFace // out

	_arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_face_get_face_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsSynthesized returns whether a `PangoFontFace` is synthesized by the
// underlying font rendering engine from another face, perhaps by shearing,
// emboldening, or lightening it.
func (f fontFace) IsSynthesized() bool {
	var _arg0 *C.PangoFontFace // out

	_arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_face_is_synthesized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListSizes: list the available sizes for a font.
//
// This is only applicable to bitmap fonts. For scalable fonts, stores nil
// at the location pointed to by @sizes and 0 at the location pointed to by
// @n_sizes. The sizes returned are in Pango units and are sorted in
// ascending order.
func (f fontFace) ListSizes() []int {
	var _arg0 *C.PangoFontFace // out

	_arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	var _arg1 *C.int
	var _arg2 C.int // in

	C.pango_font_face_list_sizes(_arg0, &_arg1, &_arg2)

	var _sizes []int

	_sizes = unsafe.Slice((*int)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_sizes, func(v *[]int) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})

	return _sizes
}

// FontFamily: a `PangoFontFamily` is used to represent a family of related font
// faces.
//
// The font faces in a family share a common design, but differ in slant,
// weight, width or other aspects.
type FontFamily interface {
	gextras.Objector

	// Name gets the name of the family.
	//
	// The name is unique among all fonts for the font backend and can be used
	// in a `PangoFontDescription` to specify that a face from this family is
	// desired.
	Name() string
	// IsMonospace: a monospace font is a font designed for text display where
	// the the characters form a regular grid.
	//
	// For Western languages this would mean that the advance width of all
	// characters are the same, but this categorization also includes Asian
	// fonts which include double-width characters: characters that occupy two
	// grid cells. g_unichar_iswide() returns a result that indicates whether a
	// character is typically double-width in a monospace font.
	//
	// The best way to find out the grid-cell size is to call
	// [method@Pango.FontMetrics.get_approximate_digit_width], since the results
	// of [method@Pango.FontMetrics.get_approximate_char_width] may be affected
	// by double-width characters.
	IsMonospace() bool
	// IsVariable: a variable font is a font which has axes that can be modified
	// to produce different faces.
	IsVariable() bool
}

// fontFamily implements the FontFamily class.
type fontFamily struct {
	gextras.Objector
}

var _ FontFamily = (*fontFamily)(nil)

// WrapFontFamily wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontFamily(obj *externglib.Object) FontFamily {
	return fontFamily{
		Objector: obj,
	}
}

func marshalFontFamily(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontFamily(obj), nil
}

// Name gets the name of the family.
//
// The name is unique among all fonts for the font backend and can be used
// in a `PangoFontDescription` to specify that a face from this family is
// desired.
func (f fontFamily) Name() string {
	var _arg0 *C.PangoFontFamily // out

	_arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_family_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsMonospace: a monospace font is a font designed for text display where
// the the characters form a regular grid.
//
// For Western languages this would mean that the advance width of all
// characters are the same, but this categorization also includes Asian
// fonts which include double-width characters: characters that occupy two
// grid cells. g_unichar_iswide() returns a result that indicates whether a
// character is typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// [method@Pango.FontMetrics.get_approximate_digit_width], since the results
// of [method@Pango.FontMetrics.get_approximate_char_width] may be affected
// by double-width characters.
func (f fontFamily) IsMonospace() bool {
	var _arg0 *C.PangoFontFamily // out

	_arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_family_is_monospace(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsVariable: a variable font is a font which has axes that can be modified
// to produce different faces.
func (f fontFamily) IsVariable() bool {
	var _arg0 *C.PangoFontFamily // out

	_arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_family_is_variable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FontDescription: a `PangoFontDescription` describes a font in an
// implementation-independent manner.
//
// `PangoFontDescription` structures are used both to list what fonts are
// available on the system and also for specifying the characteristics of a font
// to load.
type FontDescription struct {
	native C.PangoFontDescription
}

// WrapFontDescription wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontDescription(ptr unsafe.Pointer) *FontDescription {
	if ptr == nil {
		return nil
	}

	return (*FontDescription)(ptr)
}

func marshalFontDescription(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFontDescription(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// BetterMatch determines if the style attributes of @new_match are a closer
// match for @desc than those of @old_match are, or if @old_match is nil,
// determines if @new_match is a match at all.
//
// Approximate matching is done for weight and style; other style attributes
// must match exactly. Style attributes are all attributes other than family and
// size-related attributes. Approximate matching for style considers
// PANGO_STYLE_OBLIQUE and PANGO_STYLE_ITALIC as matches, but not as good a
// match as when the styles are equal.
//
// Note that @old_match must match @desc.
func (d *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(oldMatch.Native()))
	_arg2 = (*C.PangoFontDescription)(unsafe.Pointer(newMatch.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_description_better_match(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Equal compares two font descriptions for equality.
//
// Two font descriptions are considered equal if the fonts they describe are
// provably identical. This means that their masks do not have to match, as long
// as other fields are all the same. (Two font descriptions may result in
// identical fonts being loaded, but still compare false.)
func (d *FontDescription) Equal(desc2 *FontDescription) bool {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc2.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_description_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Free frees a font description.
func (d *FontDescription) Free() {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_free(_arg0)
}

// Family gets the family name field of a font description.
//
// See [method@Pango.FontDescription.set_family].
func (d *FontDescription) Family() string {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_description_get_family(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Size gets the size field of a font description.
//
// See [method@Pango.FontDescription.set_size].
func (d *FontDescription) Size() int {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret C.gint // in

	_cret = C.pango_font_description_get_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// SizeIsAbsolute determines whether the size of the font is in points (not
// absolute) or device units (absolute).
//
// See [method@Pango.FontDescription.set_size] and
// [method@Pango.FontDescription.set_absolute_size].
func (d *FontDescription) SizeIsAbsolute() bool {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_font_description_get_size_is_absolute(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Variations gets the variations field of a font description.
//
// See [method@Pango.FontDescription.set_variations].
func (d *FontDescription) Variations() string {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_description_get_variations(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Hash computes a hash of a `PangoFontDescription` structure.
//
// This is suitable to be used, for example, as an argument to
// g_hash_table_new(). The hash value is independent of @desc->mask.
func (d *FontDescription) Hash() uint {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret C.guint // in

	_cret = C.pango_font_description_hash(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Merge merges the fields that are set in @desc_to_merge into the fields in
// @desc.
//
// If @replace_existing is false, only fields in @desc that are not already set
// are affected. If true, then fields that are already set will be replaced as
// well.
//
// If @desc_to_merge is nil, this function performs nothing.
func (d *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 C.gboolean              // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(descToMerge.Native()))
	if replaceExisting {
		_arg2 = C.TRUE
	}

	C.pango_font_description_merge(_arg0, _arg1, _arg2)
}

// MergeStatic merges the fields that are set in @desc_to_merge into the fields
// in @desc, without copying allocated fields.
//
// This is like [method@Pango.FontDescription.merge], but only a shallow copy is
// made of the family name and other allocated fields. @desc can only be used
// until @desc_to_merge is modified or freed. This is meant to be used when the
// merged font description is only needed temporarily.
func (d *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 C.gboolean              // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(descToMerge.Native()))
	if replaceExisting {
		_arg2 = C.TRUE
	}

	C.pango_font_description_merge_static(_arg0, _arg1, _arg2)
}

// SetAbsoluteSize sets the size field of a font description, in device units.
//
// This is mutually exclusive with [method@Pango.FontDescription.set_size] which
// sets the font size in points.
func (d *FontDescription) SetAbsoluteSize(size float64) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.double                // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = C.double(size)

	C.pango_font_description_set_absolute_size(_arg0, _arg1)
}

// SetFamily sets the family name field of a font description.
//
// The family name represents a family of related font styles, and will resolve
// to a particular `PangoFontFamily`. In some uses of `PangoFontDescription`, it
// is also possible to use a comma separated list of family names for this
// field.
func (d *FontDescription) SetFamily(family string) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(_arg1))

	C.pango_font_description_set_family(_arg0, _arg1)
}

// SetFamilyStatic sets the family name field of a font description, without
// copying the string.
//
// This is like [method@Pango.FontDescription.set_family], except that no copy
// of @family is made. The caller must make sure that the string passed in stays
// around until @desc has been freed or the name is set again. This function can
// be used if @family is a static string such as a C string literal, or if @desc
// is only needed temporarily.
func (d *FontDescription) SetFamilyStatic(family string) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(_arg1))

	C.pango_font_description_set_family_static(_arg0, _arg1)
}

// SetGravity sets the gravity field of a font description.
//
// The gravity field specifies how the glyphs should be rotated. If @gravity is
// PANGO_GRAVITY_AUTO, this actually unsets the gravity mask on the font
// description.
//
// This function is seldom useful to the user. Gravity should normally be set on
// a `PangoContext`.
func (d *FontDescription) SetGravity(gravity Gravity) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoGravity          // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity(_arg0, _arg1)
}

// SetSize sets the size field of a font description in fractional points.
//
// This is mutually exclusive with
// [method@Pango.FontDescription.set_absolute_size].
func (d *FontDescription) SetSize(size int) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.gint                  // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(size)

	C.pango_font_description_set_size(_arg0, _arg1)
}

// SetStretch sets the stretch field of a font description.
//
// The [enum@Pango.Stretch] field specifies how narrow or wide the font should
// be.
func (d *FontDescription) SetStretch(stretch Stretch) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoStretch          // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch(_arg0, _arg1)
}

// SetStyle sets the style field of a `PangoFontDescription`.
//
// The [enum@Pango.Style] enumeration describes whether the font is slanted and
// the manner in which it is slanted; it can be either NGO_STYLE_NORMAL,
// NGO_STYLE_ITALIC, or NGO_STYLE_OBLIQUE.
//
// Most fonts will either have a italic style or an oblique style, but not both,
// and font matching in Pango will match italic specifications with oblique
// fonts and vice-versa if an exact match is not found.
func (d *FontDescription) SetStyle(style Style) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoStyle            // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoStyle)(style)

	C.pango_font_description_set_style(_arg0, _arg1)
}

// SetVariant sets the variant field of a font description.
//
// The [enum@Pango.Variant] can either be PANGO_VARIANT_NORMAL or
// PANGO_VARIANT_SMALL_CAPS.
func (d *FontDescription) SetVariant(variant Variant) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoVariant          // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoVariant)(variant)

	C.pango_font_description_set_variant(_arg0, _arg1)
}

// SetVariations sets the variations field of a font description.
//
// OpenType font variations allow to select a font instance by specifying values
// for a number of axes, such as width or weight.
//
// The format of the variations string is
//
//    AXIS1=VALUE,AXIS2=VALUE...
//
// with each AXIS a 4 character tag that identifies a font axis, and each VALUE
// a floating point number. Unknown axes are ignored, and values are clamped to
// their allowed range.
//
// Pango does not currently have a way to find supported axes of a font. Both
// harfbuzz or freetype have API for this.
func (d *FontDescription) SetVariations(variations string) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(variations))
	defer C.free(unsafe.Pointer(_arg1))

	C.pango_font_description_set_variations(_arg0, _arg1)
}

// SetVariationsStatic sets the variations field of a font description.
//
// This is like [method@Pango.FontDescription.set_variations], except that no
// copy of @variations is made. The caller must make sure that the string passed
// in stays around until @desc has been freed or the name is set again. This
// function can be used if @variations is a static string such as a C string
// literal, or if @desc is only needed temporarily.
func (d *FontDescription) SetVariationsStatic(variations string) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(variations))
	defer C.free(unsafe.Pointer(_arg1))

	C.pango_font_description_set_variations_static(_arg0, _arg1)
}

// SetWeight sets the weight field of a font description.
//
// The weight field specifies how bold or light the font should be. In addition
// to the values of the [enum@Pango.Weight] enumeration, other intermediate
// numeric values are possible.
func (d *FontDescription) SetWeight(weight Weight) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoWeight           // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoWeight)(weight)

	C.pango_font_description_set_weight(_arg0, _arg1)
}

// ToFilename creates a filename representation of a font description.
//
// The filename is identical to the result from calling
// [method@Pango.FontDescription.to_string], but with underscores instead of
// characters that are untypical in filenames, and in lower case only.
func (d *FontDescription) ToFilename() string {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_description_to_filename(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// String creates a string representation of a font description.
//
// See [type_func@Pango.FontDescription.from_string] for a description of the
// format of the string representation. The family list in the string
// description will only have a terminating comma if the last word of the list
// is a valid style option.
func (d *FontDescription) String() string {
	var _arg0 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var _cret *C.char // in

	_cret = C.pango_font_description_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UnsetFields unsets some of the fields in a `PangoFontDescription`.
//
// The unset fields will get back to their default values.
func (d *FontDescription) UnsetFields(toUnset FontMask) {
	var _arg0 *C.PangoFontDescription // out
	var _arg1 C.PangoFontMask         // out

	_arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	_arg1 = (C.PangoFontMask)(toUnset)

	C.pango_font_description_unset_fields(_arg0, _arg1)
}

// FontMetrics: a `PangoFontMetrics` structure holds the overall metric
// information for a font.
//
// The information in a `PangoFontMetrics` structure may be restricted to a
// script. The fields of this structure are private to implementations of a font
// backend. See the documentation of the corresponding getters for documentation
// of their meaning.
type FontMetrics struct {
	native C.PangoFontMetrics
}

// WrapFontMetrics wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontMetrics(ptr unsafe.Pointer) *FontMetrics {
	if ptr == nil {
		return nil
	}

	return (*FontMetrics)(ptr)
}

func marshalFontMetrics(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFontMetrics(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontMetrics) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// ApproximateCharWidth gets the approximate character width for a font metrics
// structure.
//
// This is merely a representative value useful, for example, for determining
// the initial size for a window. Actual characters in text will be wider and
// narrower than this.
func (m *FontMetrics) ApproximateCharWidth() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_approximate_char_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// ApproximateDigitWidth gets the approximate digit width for a font metrics
// structure.
//
// This is merely a representative value useful, for example, for determining
// the initial size for a window. Actual digits in text can be wider or narrower
// than this, though this value is generally somewhat more accurate than the
// result of pango_font_metrics_get_approximate_char_width() for digits.
func (m *FontMetrics) ApproximateDigitWidth() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_approximate_digit_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Ascent gets the ascent from a font metrics structure.
//
// The ascent is the distance from the baseline to the logical top of a line of
// text. (The logical top may be above or below the top of the actual drawn ink.
// It is necessary to lay out the text to figure where the ink will be.)
func (m *FontMetrics) Ascent() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_ascent(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Descent gets the descent from a font metrics structure.
//
// The descent is the distance from the baseline to the logical bottom of a line
// of text. (The logical bottom may be above or below the bottom of the actual
// drawn ink. It is necessary to lay out the text to figure where the ink will
// be.)
func (m *FontMetrics) Descent() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_descent(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Height gets the line height from a font metrics structure.
//
// The line height is the distance between successive baselines in wrapped text.
//
// If the line height is not available, 0 is returned.
func (m *FontMetrics) Height() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// StrikethroughPosition gets the suggested position to draw the strikethrough.
//
// The value returned is the distance *above* the baseline of the top of the
// strikethrough.
func (m *FontMetrics) StrikethroughPosition() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_strikethrough_position(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// StrikethroughThickness gets the suggested thickness to draw for the
// strikethrough.
func (m *FontMetrics) StrikethroughThickness() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_strikethrough_thickness(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// UnderlinePosition gets the suggested position to draw the underline.
//
// The value returned is the distance *above* the baseline of the top of the
// underline. Since most fonts have underline positions beneath the baseline,
// this value is typically negative.
func (m *FontMetrics) UnderlinePosition() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_underline_position(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// UnderlineThickness gets the suggested thickness to draw for the underline.
func (m *FontMetrics) UnderlineThickness() int {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.pango_font_metrics_get_underline_thickness(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Unref: decrease the reference count of a font metrics structure by one. If
// the result is zero, frees the structure and any associated memory.
func (m *FontMetrics) Unref() {
	var _arg0 *C.PangoFontMetrics // out

	_arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_unref(_arg0)
}
