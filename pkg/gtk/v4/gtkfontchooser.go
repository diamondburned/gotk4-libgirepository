// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk4_FontFilterFunc(PangoFontFamily*, PangoFontFace*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_level_get_type()), F: marshalFontChooserLevel},
		{T: externglib.Type(C.gtk_font_chooser_get_type()), F: marshalFontChooserer},
	})
}

// FontChooserLevel specifies the granularity of font selection that is desired
// in a GtkFontChooser.
//
// This enumeration may be extended in the future; applications should ignore
// unknown values.
type FontChooserLevel int

const (
	// FontChooserLevelFamily: allow selecting a font family
	FontChooserLevelFamily FontChooserLevel = 0b0
	// FontChooserLevelStyle: allow selecting a specific font face
	FontChooserLevelStyle FontChooserLevel = 0b1
	// FontChooserLevelSize: allow selecting a specific font size
	FontChooserLevelSize FontChooserLevel = 0b10
	// FontChooserLevelVariations: allow changing OpenType font variation axes
	FontChooserLevelVariations FontChooserLevel = 0b100
	// FontChooserLevelFeatures: allow selecting specific OpenType font features
	FontChooserLevelFeatures FontChooserLevel = 0b1000
)

func marshalFontChooserLevel(p uintptr) (interface{}, error) {
	return FontChooserLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for FontChooserLevel.
func (f FontChooserLevel) String() string {
	if f == 0 {
		return "FontChooserLevel(0)"
	}

	var builder strings.Builder
	builder.Grow(117)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FontChooserLevelFamily:
			builder.WriteString("Family|")
		case FontChooserLevelStyle:
			builder.WriteString("Style|")
		case FontChooserLevelSize:
			builder.WriteString("Size|")
		case FontChooserLevelVariations:
			builder.WriteString("Variations|")
		case FontChooserLevelFeatures:
			builder.WriteString("Features|")
		default:
			builder.WriteString(fmt.Sprintf("FontChooserLevel(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// FontFilterFunc: type of function that is used for deciding what fonts get
// shown in a GtkFontChooser.
//
// See gtk.FontChooser.SetFilterFunc().
type FontFilterFunc func(family pango.FontFamilier, face pango.FontFacer) (ok bool)

//export _gotk4_gtk4_FontFilterFunc
func _gotk4_gtk4_FontFilterFunc(arg0 *C.PangoFontFamily, arg1 *C.PangoFontFace, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var family pango.FontFamilier // out
	var face pango.FontFacer      // out

	family = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(pango.FontFamilier)
	face = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(pango.FontFacer)

	fn := v.(FontFilterFunc)
	ok := fn(family, face)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontChooserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontChooserOverrider interface {
	FontActivated(fontname string)
	// FontFace gets the PangoFontFace representing the selected font group
	// details (i.e. family, slant, weight, width, etc).
	//
	// If the selected font is not installed, returns NULL.
	FontFace() pango.FontFacer
	// FontFamily gets the PangoFontFamily representing the selected font
	// family.
	//
	// Font families are a collection of font faces.
	//
	// If the selected font is not installed, returns NULL.
	FontFamily() pango.FontFamilier
	// FontMap gets the custom font map of this font chooser widget, or NULL if
	// it does not have one.
	FontMap() pango.FontMapper
	// FontSize: selected font size.
	FontSize() int
	// SetFilterFunc adds a filter function that decides which fonts to display
	// in the font chooser.
	SetFilterFunc(filter FontFilterFunc)
	// SetFontMap sets a custom font map to use for this font chooser widget.
	//
	// A custom font map can be used to present application-specific fonts
	// instead of or in addition to the normal system fonts.
	//
	//    FcConfig *config;
	//    PangoFontMap *fontmap;
	//
	//    config = FcInitLoadConfigAndFonts ();
	//    FcConfigAppFontAddFile (config, my_app_font_file);
	//
	//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
	//
	//
	// Note that other GTK widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	//    context = gtk_widget_get_pango_context (label);
	//    pango_context_set_font_map (context, fontmap);
	SetFontMap(fontmap pango.FontMapper)
}

// FontChooser: GtkFontChooser is an interface that can be implemented by
// widgets for choosing fonts.
//
// In GTK, the main objects that implement this interface are
// gtk.FontChooserWidget, gtk.FontChooserDialog and gtk.FontButton.
type FontChooser struct {
	*externglib.Object
}

// FontChooserer describes FontChooser's abstract methods.
type FontChooserer interface {
	gextras.Objector

	// Font gets the currently-selected font name.
	Font() string
	// FontDesc gets the currently-selected font.
	FontDesc() *pango.FontDescription
	// FontFace gets the PangoFontFace representing the selected font group
	// details (i.e.
	FontFace() pango.FontFacer
	// FontFamily gets the PangoFontFamily representing the selected font
	// family.
	FontFamily() pango.FontFamilier
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// FontMap gets the custom font map of this font chooser widget, or NULL if
	// it does not have one.
	FontMap() pango.FontMapper
	// FontSize: selected font size.
	FontSize() int
	// Language gets the language that is used for font features.
	Language() string
	// Level returns the current level of granularity for selecting fonts.
	Level() FontChooserLevel
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFilterFunc adds a filter function that decides which fonts to display
	// in the font chooser.
	SetFilterFunc(filter FontFilterFunc)
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetFontMap sets a custom font map to use for this font chooser widget.
	SetFontMap(fontmap pango.FontMapper)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetLevel sets the desired level of granularity for selecting fonts.
	SetLevel(level FontChooserLevel)
	// SetPreviewText sets the text displayed in the preview area.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)
}

var _ FontChooserer = (*FontChooser)(nil)

func wrapFontChooser(obj *externglib.Object) *FontChooser {
	return &FontChooser{
		Object: obj,
	}
}

func marshalFontChooserer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooser(obj), nil
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk.FontChooser.SetFont(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango.FontDescription.Equal() if you want to compare two font
// descriptions.
func (fontchooser *FontChooser) Font() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// FontDesc gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// gtk.FontChooser.SetFont(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango.FontDescription.Equal() if you want to compare two font
// descriptions.
func (fontchooser *FontChooser) FontDesc() *pango.FontDescription {
	var _arg0 *C.GtkFontChooser       // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_desc(_arg0)

	var _fontDescription *pango.FontDescription // out

	if _cret != nil {
		_fontDescription = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_fontDescription, func(v *pango.FontDescription) {
			C.pango_font_description_free((*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _fontDescription
}

// FontFace gets the PangoFontFace representing the selected font group details
// (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns NULL.
func (fontchooser *FontChooser) FontFace() pango.FontFacer {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontFace  // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_face(_arg0)

	var _fontFace pango.FontFacer // out

	if _cret != nil {
		_fontFace = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(pango.FontFacer)
	}

	return _fontFace
}

// FontFamily gets the PangoFontFamily representing the selected font family.
//
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns NULL.
func (fontchooser *FontChooser) FontFamily() pango.FontFamilier {
	var _arg0 *C.GtkFontChooser  // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_family(_arg0)

	var _fontFamily pango.FontFamilier // out

	if _cret != nil {
		_fontFamily = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(pango.FontFamilier)
	}

	return _fontFamily
}

// FontFeatures gets the currently-selected font features.
func (fontchooser *FontChooser) FontFeatures() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_features(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontMap gets the custom font map of this font chooser widget, or NULL if it
// does not have one.
func (fontchooser *FontChooser) FontMap() pango.FontMapper {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontMap   // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_map(_arg0)

	var _fontMap pango.FontMapper // out

	if _cret != nil {
		_fontMap = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(pango.FontMapper)
	}

	return _fontMap
}

// FontSize: selected font size.
func (fontchooser *FontChooser) FontSize() int {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.int             // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_font_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Language gets the language that is used for font features.
func (fontchooser *FontChooser) Language() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_language(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Level returns the current level of granularity for selecting fonts.
func (fontchooser *FontChooser) Level() FontChooserLevel {
	var _arg0 *C.GtkFontChooser     // out
	var _cret C.GtkFontChooserLevel // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_level(_arg0)

	var _fontChooserLevel FontChooserLevel // out

	_fontChooserLevel = FontChooserLevel(_cret)

	return _fontChooserLevel
}

// PreviewText gets the text displayed in the preview area.
func (fontchooser *FontChooser) PreviewText() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
func (fontchooser *FontChooser) ShowPreviewEntry() bool {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))

	_cret = C.gtk_font_chooser_get_show_preview_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFilterFunc adds a filter function that decides which fonts to display in
// the font chooser.
func (fontchooser *FontChooser) SetFilterFunc(filter FontFilterFunc) {
	var _arg0 *C.GtkFontChooser   // out
	var _arg1 C.GtkFontFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	if filter != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_FontFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(filter))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_font_chooser_set_filter_func(_arg0, _arg1, _arg2, _arg3)
}

// SetFont sets the currently-selected font.
func (fontchooser *FontChooser) SetFont(fontname string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_font(_arg0, _arg1)
}

// SetFontDesc sets the currently-selected font from font_desc.
func (fontchooser *FontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var _arg0 *C.GtkFontChooser       // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(fontDesc)))

	C.gtk_font_chooser_set_font_desc(_arg0, _arg1)
}

// SetFontMap sets a custom font map to use for this font chooser widget.
//
// A custom font map can be used to present application-specific fonts instead
// of or in addition to the normal system fonts.
//
//    FcConfig *config;
//    PangoFontMap *fontmap;
//
//    config = FcInitLoadConfigAndFonts ();
//    FcConfigAppFontAddFile (config, my_app_font_file);
//
//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
//
//
// Note that other GTK widgets will only be able to use the application-specific
// font if it is present in the font map they use:
//
//    context = gtk_widget_get_pango_context (label);
//    pango_context_set_font_map (context, fontmap);
func (fontchooser *FontChooser) SetFontMap(fontmap pango.FontMapper) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	if fontmap != nil {
		_arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))
	}

	C.gtk_font_chooser_set_font_map(_arg0, _arg1)
}

// SetLanguage sets the language to use for font features.
func (fontchooser *FontChooser) SetLanguage(language string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(language)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_language(_arg0, _arg1)
}

// SetLevel sets the desired level of granularity for selecting fonts.
func (fontchooser *FontChooser) SetLevel(level FontChooserLevel) {
	var _arg0 *C.GtkFontChooser     // out
	var _arg1 C.GtkFontChooserLevel // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	_arg1 = C.GtkFontChooserLevel(level)

	C.gtk_font_chooser_set_level(_arg0, _arg1)
}

// SetPreviewText sets the text displayed in the preview area.
//
// The text is used to show how the selected font looks.
func (fontchooser *FontChooser) SetPreviewText(text string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_preview_text(_arg0, _arg1)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
func (fontchooser *FontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(fontchooser.Native()))
	if showPreviewEntry {
		_arg1 = C.TRUE
	}

	C.gtk_font_chooser_set_show_preview_entry(_arg0, _arg1)
}
