// Code generated by girgen. DO NOT EDIT.

package pangoot

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoot pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango-ot.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_ot_buffer_get_type()), F: marshalBuffer},
		{T: externglib.Type(C.pango_ot_ruleset_description_get_type()), F: marshalRulesetDescription},
	})
}

// Tag: the OTTag typedef is used to represent TrueType and OpenType four letter
// tags inside Pango. Use PANGO_OT_TAG_MAKE() or PANGO_OT_TAG_MAKE_FROM_STRING()
// macros to create PangoOTTags manually.
type Tag uint32

// TableType: the PangoOTTableType enumeration values are used to identify the
// various OpenType tables in the pango_ot_info_… functions.
type TableType int

const (
	// TableTypeGsub: the GSUB table.
	TableTypeGsub TableType = 0
	// TableTypeGpos: the GPOS table.
	TableTypeGpos TableType = 1
)

type Buffer struct {
	native C.PangoOTBuffer
}

// WrapBuffer wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBuffer(ptr unsafe.Pointer) *Buffer {
	if ptr == nil {
		return nil
	}

	return (*Buffer)(ptr)
}

func marshalBuffer(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBuffer(unsafe.Pointer(b)), nil
}

// NewBuffer constructs a struct Buffer.
func NewBuffer(font pangofc.Font) *Buffer {
	var arg1 *C.PangoFcFont

	arg1 = (*C.PangoFcFont)(unsafe.Pointer(font.Native()))

	cret := new(C.PangoOTBuffer)
	var goret *Buffer

	cret = C.pango_ot_buffer_new(arg1)

	goret = WrapBuffer(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Buffer) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (b *Buffer) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// AddGlyph appends a glyph to a OTBuffer, with @properties identifying which
// features should be applied on this glyph. See pango_ot_ruleset_add_feature().
func (b *Buffer) AddGlyph(glyph uint, properties uint, cluster uint) {
	var arg0 *C.PangoOTBuffer
	var arg1 C.guint
	var arg2 C.guint
	var arg3 C.guint

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.guint(glyph)
	arg2 = C.guint(properties)
	arg3 = C.guint(cluster)

	C.pango_ot_buffer_add_glyph(arg0, arg1, arg2, arg3)
}

// Clear empties a OTBuffer, make it ready to add glyphs to.
func (b *Buffer) Clear() {
	var arg0 *C.PangoOTBuffer

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))

	C.pango_ot_buffer_clear(arg0)
}

// Destroy destroys a OTBuffer and free all associated memory.
func (b *Buffer) Destroy() {
	var arg0 *C.PangoOTBuffer

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))

	C.pango_ot_buffer_destroy(arg0)
}

// Glyphs gets the glyph array contained in a OTBuffer. The glyphs are owned by
// the buffer and should not be freed, and are only valid as long as buffer is
// not modified.
func (b *Buffer) Glyphs() {
	var arg0 *C.PangoOTBuffer

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))

	C.pango_ot_buffer_get_glyphs(arg0, arg1, arg2)

	return ret1, ret2
}

// Output exports the glyphs in a OTBuffer into a GlyphString. This is typically
// used after the OpenType layout processing is over, to convert the resulting
// glyphs into a generic Pango glyph string.
func (b *Buffer) Output(glyphs *pango.GlyphString) {
	var arg0 *C.PangoOTBuffer
	var arg1 *C.PangoGlyphString

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))
	arg1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_ot_buffer_output(arg0, arg1)
}

// SetRTL sets whether glyphs will be rendered right-to-left. This setting is
// needed for proper horizontal positioning of right-to-left scripts.
func (b *Buffer) SetRTL(rtL bool) {
	var arg0 *C.PangoOTBuffer
	var arg1 C.gboolean

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))
	if rtL {
		arg1 = C.gboolean(1)
	}

	C.pango_ot_buffer_set_rtl(arg0, arg1)
}

// SetZeroWidthMarks sets whether characters with a mark class should be forced
// to zero width. This setting is needed for proper positioning of Arabic
// accents, but will produce incorrect results with standard OpenType Indic
// fonts.
func (b *Buffer) SetZeroWidthMarks(zeroWidthMarks bool) {
	var arg0 *C.PangoOTBuffer
	var arg1 C.gboolean

	arg0 = (*C.PangoOTBuffer)(unsafe.Pointer(b.Native()))
	if zeroWidthMarks {
		arg1 = C.gboolean(1)
	}

	C.pango_ot_buffer_set_zero_width_marks(arg0, arg1)
}

// FeatureMap: the OTFeatureMap typedef is used to represent an OpenType feature
// with the property bit associated with it. The feature tag is represented as a
// char array instead of a OTTag for convenience.
type FeatureMap struct {
	native C.PangoOTFeatureMap
}

// WrapFeatureMap wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFeatureMap(ptr unsafe.Pointer) *FeatureMap {
	if ptr == nil {
		return nil
	}

	return (*FeatureMap)(ptr)
}

func marshalFeatureMap(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFeatureMap(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FeatureMap) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// FeatureName gets the field inside the struct.
func (f *FeatureMap) FeatureName() [5]byte {
	var v [5]byte
	v = *(*[5]byte)(unsafe.Pointer(f.native.feature_name))
	return v
}

// PropertyBit gets the field inside the struct.
func (f *FeatureMap) PropertyBit() uint32 {
	var v uint32
	v = uint32(f.native.property_bit)
	return v
}

// Glyph: the OTGlyph structure represents a single glyph together with
// information used for OpenType layout processing of the glyph. It contains the
// following fields.
type Glyph struct {
	native C.PangoOTGlyph
}

// WrapGlyph wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyph(ptr unsafe.Pointer) *Glyph {
	if ptr == nil {
		return nil
	}

	return (*Glyph)(ptr)
}

func marshalGlyph(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyph(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *Glyph) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// Glyph gets the field inside the struct.
func (g *Glyph) Glyph() uint32 {
	var v uint32
	v = uint32(g.native.glyph)
	return v
}

// Properties gets the field inside the struct.
func (g *Glyph) Properties() uint {
	var v uint
	v = uint(g.native.properties)
	return v
}

// Cluster gets the field inside the struct.
func (g *Glyph) Cluster() uint {
	var v uint
	v = uint(g.native.cluster)
	return v
}

// Component gets the field inside the struct.
func (g *Glyph) Component() uint16 {
	var v uint16
	v = uint16(g.native.component)
	return v
}

// LigID gets the field inside the struct.
func (g *Glyph) LigID() uint16 {
	var v uint16
	v = uint16(g.native.ligID)
	return v
}

// Internal gets the field inside the struct.
func (g *Glyph) Internal() uint {
	var v uint
	v = uint(g.native.internal)
	return v
}

// RulesetDescription: the OTRuleset structure holds all the information needed
// to build a complete OTRuleset from an OpenType font. The main use of this
// struct is to act as the key for a per-font hash of rulesets. The user
// populates a ruleset description and gets the ruleset using
// pango_ot_ruleset_get_for_description() or create a new one using
// pango_ot_ruleset_new_from_description().
type RulesetDescription struct {
	native C.PangoOTRulesetDescription
}

// WrapRulesetDescription wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRulesetDescription(ptr unsafe.Pointer) *RulesetDescription {
	if ptr == nil {
		return nil
	}

	return (*RulesetDescription)(ptr)
}

func marshalRulesetDescription(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRulesetDescription(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RulesetDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Script gets the field inside the struct.
func (r *RulesetDescription) Script() pango.Script {
	var v pango.Script
	v = pango.Script(r.native.script)
	return v
}

// Language gets the field inside the struct.
func (r *RulesetDescription) Language() *pango.Language {
	var v *pango.Language
	v = pango.WrapLanguage(unsafe.Pointer(r.native.language))
	return v
}

// StaticGsubFeatures gets the field inside the struct.
func (r *RulesetDescription) StaticGsubFeatures() *FeatureMap {
	var v *FeatureMap
	v = WrapFeatureMap(unsafe.Pointer(r.native.static_gsub_features))
	return v
}

// NStaticGsubFeatures gets the field inside the struct.
func (r *RulesetDescription) NStaticGsubFeatures() uint {
	var v uint
	v = uint(r.native.n_static_gsub_features)
	return v
}

// StaticGposFeatures gets the field inside the struct.
func (r *RulesetDescription) StaticGposFeatures() *FeatureMap {
	var v *FeatureMap
	v = WrapFeatureMap(unsafe.Pointer(r.native.static_gpos_features))
	return v
}

// NStaticGposFeatures gets the field inside the struct.
func (r *RulesetDescription) NStaticGposFeatures() uint {
	var v uint
	v = uint(r.native.n_static_gpos_features)
	return v
}

// OtherFeatures gets the field inside the struct.
func (r *RulesetDescription) OtherFeatures() *FeatureMap {
	var v *FeatureMap
	v = WrapFeatureMap(unsafe.Pointer(r.native.other_features))
	return v
}

// NOtherFeatures gets the field inside the struct.
func (r *RulesetDescription) NOtherFeatures() uint {
	var v uint
	v = uint(r.native.n_other_features)
	return v
}

// Copy creates a copy of @desc, which should be freed with
// pango_ot_ruleset_description_free(). Primarily used internally by
// pango_ot_ruleset_get_for_description() to cache rulesets for ruleset
// descriptions.
func (d *RulesetDescription) Copy() *RulesetDescription {
	var arg0 *C.PangoOTRulesetDescription

	arg0 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(d.Native()))

	cret := new(C.PangoOTRulesetDescription)
	var goret *RulesetDescription

	cret = C.pango_ot_ruleset_description_copy(arg0)

	goret = WrapRulesetDescription(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *RulesetDescription) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Equal compares two ruleset descriptions for equality. Two ruleset
// descriptions are considered equal if the rulesets they describe are provably
// identical. This means that their script, language, and all feature sets
// should be equal. For static feature sets, the array addresses are compared
// directly, while for other features, the list of features is compared one by
// one. (Two ruleset descriptions may result in identical rulesets being
// created, but still compare false.)
func (d *RulesetDescription) Equal(desc2 *RulesetDescription) bool {
	var arg0 *C.PangoOTRulesetDescription
	var arg1 *C.PangoOTRulesetDescription

	arg0 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(desc2.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.pango_ot_ruleset_description_equal(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Free frees a ruleset description allocated by
// pango_ot_ruleset_description_copy().
func (d *RulesetDescription) Free() {
	var arg0 *C.PangoOTRulesetDescription

	arg0 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(d.Native()))

	C.pango_ot_ruleset_description_free(arg0)
}

// Hash computes a hash of a OTRulesetDescription structure suitable to be used,
// for example, as an argument to g_hash_table_new().
func (d *RulesetDescription) Hash() uint {
	var arg0 *C.PangoOTRulesetDescription

	arg0 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(d.Native()))

	var cret C.guint
	var goret uint

	cret = C.pango_ot_ruleset_description_hash(arg0)

	goret = uint(cret)

	return goret
}
