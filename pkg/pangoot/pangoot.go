// Code generated by girgen. DO NOT EDIT.

package pangoot

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	"github.com/gotk3/gotk3/glib"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoot
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango-ot.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		// Enums
		// Skipped TableType.

		// Objects/Classes
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

// TagFromLanguage: finds the OpenType language-system tag best describing
// @language.
func TagFromLanguage(language *pango.Language) Tag

// TagFromScript: finds the OpenType script tag corresponding to @script.
//
// The PANGO_SCRIPT_COMMON, PANGO_SCRIPT_INHERITED, and PANGO_SCRIPT_UNKNOWN
// scripts are mapped to the OpenType 'DFLT' script tag that is also defined as
// PANGO_OT_TAG_DEFAULT_SCRIPT.
//
// Note that multiple Script values may map to the same OpenType script tag. In
// particular, PANGO_SCRIPT_HIRAGANA and PANGO_SCRIPT_KATAKANA both map to the
// OT tag 'kana'.
func TagFromScript(script pango.Script) Tag

// TagToLanguage: finds a Language corresponding to @language_tag.
func TagToLanguage(languageTag Tag) *pango.Language

// TagToScript: finds the Script corresponding to @script_tag.
//
// The 'DFLT' script tag is mapped to PANGO_SCRIPT_COMMON.
//
// Note that an OpenType script tag may correspond to multiple Script values. In
// such cases, the Script value with the smallest value is returned. In
// particular, PANGO_SCRIPT_HIRAGANA and PANGO_SCRIPT_KATAKANA both map to the
// OT tag 'kana'. This function will return PANGO_SCRIPT_HIRAGANA for 'kana'.
func TagToScript(scriptTag Tag) pango.Script

type Buffer struct {
	native *C.PangoOTBuffer
}

func wrapBuffer(p *C.PangoOTBuffer) *Buffer {
	v := Buffer{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*Buffer).free)

	return &v
}

func marshalBuffer(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.PangoOTBuffer)(unsafe.Pointer(b))

	return wrapBuffer(c)
}

func (b *Buffer) free() {}

// Native returns the pointer to *C.PangoOTBuffer. The caller is expected to
// cast.
func (b *Buffer) Native() unsafe.Pointer {
	return unsafe.Pointer(b.native)
}

func NewBuffer(font *pangofc.Font) *Buffer

// FeatureMap: the OTFeatureMap typedef is used to represent an OpenType feature
// with the property bit associated with it. The feature tag is represented as a
// char array instead of a OTTag for convenience.
type FeatureMap struct {
	// FeatureName: feature tag in represented as four-letter ASCII string.
	FeatureName [5]byte
	// PropertyBit: the property bit to use for this feature. See
	// pango_ot_ruleset_add_feature() for details.
	PropertyBit uint32

	native *C.PangoOTFeatureMap
}

func wrapFeatureMap(p *C.PangoOTFeatureMap) *FeatureMap {
	var v FeatureMap

	v.PropertyBit = uint32(p.property_bit)

	return &v
}

func marshalFeatureMap(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.PangoOTFeatureMap)(unsafe.Pointer(b))

	return wrapFeatureMap(c)
}

// Native returns the pointer to *C.PangoOTFeatureMap. The caller is expected to
// cast.
func (f *FeatureMap) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

// Glyph: the OTGlyph structure represents a single glyph together with
// information used for OpenType layout processing of the glyph. It contains the
// following fields.
type Glyph struct {
	// Glyph: the glyph itself.
	Glyph uint32
	// Properties: the properties value, identifying which features should be
	// applied on this glyph. See pango_ot_ruleset_add_feature().
	Properties uint
	// Cluster: the cluster that this glyph belongs to.
	Cluster uint
	// Component: a component value, set by the OpenType layout engine.
	Component uint16
	// LigID: a ligature index value, set by the OpenType layout engine.
	LigID uint16
	// Internal: for Pango internal use
	Internal uint

	native *C.PangoOTGlyph
}

func wrapGlyph(p *C.PangoOTGlyph) *Glyph {
	var v Glyph

	v.Glyph = uint32(p.glyph)
	v.Properties = uint(p.properties)
	v.Cluster = uint(p.cluster)
	v.Component = uint16(p.component)
	v.LigID = uint16(p.ligID)
	v.Internal = uint(p.internal)

	return &v
}

func marshalGlyph(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.PangoOTGlyph)(unsafe.Pointer(b))

	return wrapGlyph(c)
}

// Native returns the pointer to *C.PangoOTGlyph. The caller is expected to
// cast.
func (g *Glyph) Native() unsafe.Pointer {
	return unsafe.Pointer(g.native)
}

// RulesetDescription: the OTRuleset structure holds all the information needed
// to build a complete OTRuleset from an OpenType font. The main use of this
// struct is to act as the key for a per-font hash of rulesets. The user
// populates a ruleset description and gets the ruleset using
// pango_ot_ruleset_get_for_description() or create a new one using
// pango_ot_ruleset_new_from_description().
type RulesetDescription struct {
	// Script: a Script.
	Script pango.Script
	// Language: a Language.
	Language *pango.Language
	// StaticGsubFeatures: static map of GSUB features, or nil.
	StaticGsubFeatures *FeatureMap
	// NStaticGsubFeatures: length of @static_gsub_features, or 0.
	NStaticGsubFeatures uint
	// StaticGposFeatures: static map of GPOS features, or nil.
	StaticGposFeatures *FeatureMap
	// NStaticGposFeatures: length of @static_gpos_features, or 0.
	NStaticGposFeatures uint
	// OtherFeatures: map of extra features to add to both GSUB and GPOS, or
	// nil. Unlike the static maps, this pointer need not live beyond the life
	// of function calls taking this struct.
	OtherFeatures *FeatureMap
	// NOtherFeatures: length of @other_features, or 0.
	NOtherFeatures uint

	native *C.PangoOTRulesetDescription
}

func wrapRulesetDescription(p *C.PangoOTRulesetDescription) *RulesetDescription {
	var v RulesetDescription

	v.Script = Script(p.script)
	v.Language = wrap * Language(p.language)
	v.StaticGsubFeatures = wrap * FeatureMap(p.static_gsub_features)
	v.NStaticGsubFeatures = uint(p.n_static_gsub_features)
	v.StaticGposFeatures = wrap * FeatureMap(p.static_gpos_features)
	v.NStaticGposFeatures = uint(p.n_static_gpos_features)
	v.OtherFeatures = wrap * FeatureMap(p.other_features)
	v.NOtherFeatures = uint(p.n_other_features)

	return &v
}

func marshalRulesetDescription(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.PangoOTRulesetDescription)(unsafe.Pointer(b))

	return wrapRulesetDescription(c)
}

// Native returns the pointer to *C.PangoOTRulesetDescription. The caller is expected to
// cast.
func (r *RulesetDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

type Info struct {
	*externglib.Object
}

func wrapInfo(obj *glib.Object) *Info {
	return &Info{*externglib.Object{obj}}
}

func marshalInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// Ruleset: the OTRuleset structure holds a set of features selected from the
// tables in an OpenType font. (A feature is an operation such as adjusting
// glyph positioning that should be applied to a text feature such as a certain
// type of accent.) A OTRuleset is created with pango_ot_ruleset_new(), features
// are added to it with pango_ot_ruleset_add_feature(), then it is applied to a
// GlyphString with pango_ot_ruleset_shape().
type Ruleset struct {
	*externglib.Object
}

func wrapRuleset(obj *glib.Object) *Ruleset {
	return &Ruleset{*externglib.Object{obj}}
}

func marshalRuleset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}
