// Code generated by girgen. DO NOT EDIT.

package pangoot

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangoot
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango-ot.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_ot_ruleset_get_type()), F: marshalRuleset},
	})
}

// Ruleset: the OTRuleset structure holds a set of features selected from the
// tables in an OpenType font. (A feature is an operation such as adjusting
// glyph positioning that should be applied to a text feature such as a certain
// type of accent.) A OTRuleset is created with pango_ot_ruleset_new(), features
// are added to it with pango_ot_ruleset_add_feature(), then it is applied to a
// GlyphString with pango_ot_ruleset_shape().
type Ruleset interface {
	gextras.Objector

	// AddFeature adds a feature to the ruleset.
	AddFeature(tableType TableType, featureIndex uint, propertyBit uint32)
	// FeatureCount gets the number of GSUB and GPOS features in the ruleset.
	FeatureCount() (nGsubFeatures uint, nGposFeatures uint, guint uint)
	// MaybeAddFeatures: this is a convenience function that for each feature in
	// the feature map array @features converts the feature name to a OTTag
	// feature tag using PANGO_OT_TAG_MAKE() and calls
	// pango_ot_ruleset_maybe_add_feature() on it.
	MaybeAddFeatures(tableType TableType, features *FeatureMap, nFeatures uint) uint
	// Position performs the OpenType GPOS positioning on @buffer using the
	// features in @ruleset
	Position(buffer *Buffer)
	// Substitute performs the OpenType GSUB substitution on @buffer using the
	// features in @ruleset
	Substitute(buffer *Buffer)
}

// ruleset implements the Ruleset class.
type ruleset struct {
	gextras.Objector
}

var _ Ruleset = (*ruleset)(nil)

// WrapRuleset wraps a GObject to the right type. It is
// primarily used internally.
func WrapRuleset(obj *externglib.Object) Ruleset {
	return ruleset{
		Objector: obj,
	}
}

func marshalRuleset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRuleset(obj), nil
}

// AddFeature adds a feature to the ruleset.
func (r ruleset) AddFeature(tableType TableType, featureIndex uint, propertyBit uint32) {
	var _arg0 *C.PangoOTRuleset  // out
	var _arg1 C.PangoOTTableType // out
	var _arg2 C.guint            // out
	var _arg3 C.gulong           // out

	_arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoOTTableType)(tableType)
	_arg2 = C.guint(featureIndex)
	_arg3 = C.gulong(propertyBit)

	C.pango_ot_ruleset_add_feature(_arg0, _arg1, _arg2, _arg3)
}

// FeatureCount gets the number of GSUB and GPOS features in the ruleset.
func (r ruleset) FeatureCount() (nGsubFeatures uint, nGposFeatures uint, guint uint) {
	var _arg0 *C.PangoOTRuleset // out

	_arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))

	var _arg1 C.guint // in
	var _arg2 C.guint // in
	var _cret C.guint // in

	_cret = C.pango_ot_ruleset_get_feature_count(_arg0, &_arg1, &_arg2)

	var _nGsubFeatures uint // out
	var _nGposFeatures uint // out
	var _guint uint         // out

	_nGsubFeatures = (uint)(_arg1)
	_nGposFeatures = (uint)(_arg2)
	_guint = (uint)(_cret)

	return _nGsubFeatures, _nGposFeatures, _guint
}

// MaybeAddFeatures: this is a convenience function that for each feature in
// the feature map array @features converts the feature name to a OTTag
// feature tag using PANGO_OT_TAG_MAKE() and calls
// pango_ot_ruleset_maybe_add_feature() on it.
func (r ruleset) MaybeAddFeatures(tableType TableType, features *FeatureMap, nFeatures uint) uint {
	var _arg0 *C.PangoOTRuleset    // out
	var _arg1 C.PangoOTTableType   // out
	var _arg2 *C.PangoOTFeatureMap // out
	var _arg3 C.guint              // out

	_arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoOTTableType)(tableType)
	_arg2 = (*C.PangoOTFeatureMap)(unsafe.Pointer(features.Native()))
	_arg3 = C.guint(nFeatures)

	var _cret C.guint // in

	_cret = C.pango_ot_ruleset_maybe_add_features(_arg0, _arg1, _arg2, _arg3)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Position performs the OpenType GPOS positioning on @buffer using the
// features in @ruleset
func (r ruleset) Position(buffer *Buffer) {
	var _arg0 *C.PangoOTRuleset // out
	var _arg1 *C.PangoOTBuffer  // out

	_arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoOTBuffer)(unsafe.Pointer(buffer.Native()))

	C.pango_ot_ruleset_position(_arg0, _arg1)
}

// Substitute performs the OpenType GSUB substitution on @buffer using the
// features in @ruleset
func (r ruleset) Substitute(buffer *Buffer) {
	var _arg0 *C.PangoOTRuleset // out
	var _arg1 *C.PangoOTBuffer  // out

	_arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoOTBuffer)(unsafe.Pointer(buffer.Native()))

	C.pango_ot_ruleset_substitute(_arg0, _arg1)
}
