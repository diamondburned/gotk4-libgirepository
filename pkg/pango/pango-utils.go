// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// IsZeroWidth checks if a character that should not be normally rendered.
//
// This includes all Unicode characters with "ZERO WIDTH" in their name, as well
// as *bidi* formatting characters, and a few other ones. This is totally
// different from g_unichar_iszerowidth() and is at best misnamed.
func IsZeroWidth(ch uint32) bool {
	var arg1 C.gunichar

	arg1 = C.gunichar(ch)

	var cret C.gboolean
	var goret bool

	cret = C.pango_is_zero_width(arg1)

	if cret {
		goret = true
	}

	return goret
}

// Log2VisGetEmbeddingLevels: return the bidirectional embedding levels of the
// input paragraph.
//
// The bidirectional embedding levels are defined by the Unicode Bidirectional
// Algorithm available at:
//
//    http://www.unicode.org/reports/tr9/
//
// If the input base direction is a weak direction, the direction of the
// characters in the text will determine the final resolved direction.
func Log2VisGetEmbeddingLevels(text string, length int, pbaseDir *Direction) byte {
	var arg1 *C.gchar
	var arg2 C.int
	var arg3 *C.PangoDirection

	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)
	arg3 = (*C.PangoDirection)(pbaseDir)

	var cret *C.guint8
	var goret byte

	cret = C.pango_log2vis_get_embedding_levels(arg1, arg2, arg3)

	goret = byte(cret)

	return goret
}

// ParseEnum parses an enum type and stores the result in @value.
//
// If @str does not match the nick name of any of the possible values for the
// enum and is not an integer, false is returned, a warning is issued if @warn
// is true, and a string representing the list of possible values is stored in
// @possible_values. The list is slash-separated, eg. "none/start/middle/end".
// If failed and @possible_values is not nil, returned string should be freed
// using g_free().
func ParseEnum(typ externglib.Type, str string, warn bool) (value int, possibleValues string, ok bool) {
	var arg1 C.GType
	var arg2 *C.char
	var arg4 C.gboolean

	arg1 = C.GType(typ)
	arg2 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg2))
	if warn {
		arg4 = C.gboolean(1)
	}

	arg3 := new(C.int)
	var ret3 int
	arg5 := new(*C.char)
	var ret5 string
	var cret C.gboolean
	var goret bool

	cret = C.pango_parse_enum(arg1, arg2, arg3, arg4, arg5)

	ret3 = int(*arg3)
	ret5 = C.GoString(*arg5)
	defer C.free(unsafe.Pointer(*arg5))
	if cret {
		goret = true
	}

	return ret3, ret5, goret
}

// ParseStretch parses a font stretch.
//
// The allowed values are "ultra_condensed", "extra_condensed", "condensed",
// "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" and
// "ultra_expanded". Case variations are ignored and the '_' characters may be
// omitted.
func ParseStretch(str string, warn bool) (stretch *Stretch, ok bool) {
	var arg1 *C.char
	var arg3 C.gboolean

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))
	if warn {
		arg3 = C.gboolean(1)
	}

	arg2 := new(C.PangoStretch)
	var ret2 *Stretch
	var cret C.gboolean
	var goret bool

	cret = C.pango_parse_stretch(arg1, arg2, arg3)

	ret2 = *Stretch(arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// ParseStyle parses a font style.
//
// The allowed values are "normal", "italic" and "oblique", case variations
// being ignored.
func ParseStyle(str string, warn bool) (style *Style, ok bool) {
	var arg1 *C.char
	var arg3 C.gboolean

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))
	if warn {
		arg3 = C.gboolean(1)
	}

	arg2 := new(C.PangoStyle)
	var ret2 *Style
	var cret C.gboolean
	var goret bool

	cret = C.pango_parse_style(arg1, arg2, arg3)

	ret2 = *Style(arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// ParseVariant parses a font variant.
//
// The allowed values are "normal" and "smallcaps" or "small_caps", case
// variations being ignored.
func ParseVariant(str string, warn bool) (variant *Variant, ok bool) {
	var arg1 *C.char
	var arg3 C.gboolean

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))
	if warn {
		arg3 = C.gboolean(1)
	}

	arg2 := new(C.PangoVariant)
	var ret2 *Variant
	var cret C.gboolean
	var goret bool

	cret = C.pango_parse_variant(arg1, arg2, arg3)

	ret2 = *Variant(arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// ParseWeight parses a font weight.
//
// The allowed values are "heavy", "ultrabold", "bold", "normal", "light",
// "ultraleight" and integers. Case variations are ignored.
func ParseWeight(str string, warn bool) (weight *Weight, ok bool) {
	var arg1 *C.char
	var arg3 C.gboolean

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))
	if warn {
		arg3 = C.gboolean(1)
	}

	arg2 := new(C.PangoWeight)
	var ret2 *Weight
	var cret C.gboolean
	var goret bool

	cret = C.pango_parse_weight(arg1, arg2, arg3)

	ret2 = *Weight(arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// QuantizeLineGeometry quantizes the thickness and position of a line to whole
// device pixels.
//
// This is typically used for underline or strikethrough. The purpose of this
// function is to avoid such lines looking blurry.
//
// Care is taken to make sure @thickness is at least one pixel when this
// function returns, but returned @position may become zero as a result of
// rounding.
func QuantizeLineGeometry(thickness int, position int) {
	var arg1 *C.int
	var arg2 *C.int

	arg1 = *C.int(thickness)
	arg2 = *C.int(position)

	C.pango_quantize_line_geometry(arg1, arg2)
}

// ReadLine reads an entire line from a file into a buffer.
//
// Lines may be delimited with '\n', '\r', '\n\r', or '\r\n'. The delimiter is
// not written into the buffer. Text after a '#' character is treated as a
// comment and skipped. '\' can be used to escape a
//
// character.
//
// '\' proceeding a line delimiter combines adjacent lines. A '\' proceeding any
// other character is ignored and written into the output buffer unmodified.
func ReadLine(stream interface{}, str *glib.String) int {
	var arg1 *C.FILE
	var arg2 *C.GString

	arg1 = *C.FILE(stream)
	arg2 = (*C.GString)(unsafe.Pointer(str.Native()))

	var cret C.gint
	var goret int

	cret = C.pango_read_line(arg1, arg2)

	goret = int(cret)

	return goret
}

// ScanInt scans an integer.
//
// Leading white space is skipped.
func ScanInt(pos string) (out int, ok bool) {
	var arg1 **C.char

	arg1 = (**C.char)(C.CString(pos))

	arg2 := new(C.int)
	var ret2 int
	var cret C.gboolean
	var goret bool

	cret = C.pango_scan_int(arg1, arg2)

	ret2 = int(*arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// ScanString scans a string into a #GString buffer.
//
// The string may either be a sequence of non-white-space characters, or a
// quoted string with '"'. Instead a quoted string, '\"' represents a literal
// quote. Leading white space outside of quotes is skipped.
func ScanString(pos string, out *glib.String) bool {
	var arg1 **C.char
	var arg2 *C.GString

	arg1 = (**C.char)(C.CString(pos))
	arg2 = (*C.GString)(unsafe.Pointer(out.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.pango_scan_string(arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// ScanWord scans a word into a #GString buffer.
//
// A word consists of [A-Za-z_] followed by zero or more [A-Za-z_0-9]. Leading
// white space is skipped.
func ScanWord(pos string, out *glib.String) bool {
	var arg1 **C.char
	var arg2 *C.GString

	arg1 = (**C.char)(C.CString(pos))
	arg2 = (*C.GString)(unsafe.Pointer(out.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.pango_scan_word(arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// SkipSpace skips 0 or more characters of white space.
func SkipSpace(pos string) bool {
	var arg1 **C.char

	arg1 = (**C.char)(C.CString(pos))

	var cret C.gboolean
	var goret bool

	cret = C.pango_skip_space(arg1)

	if cret {
		goret = true
	}

	return goret
}

// SplitFileList splits a G_SEARCHPATH_SEPARATOR-separated list of files,
// stripping white space and substituting ~/ with $HOME/.
func SplitFileList(str string) []string {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret **C.char
	var goret []string

	cret = C.pango_split_file_list(arg1)

	{
		var length int
		for p := cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return goret
}

// TrimString trims leading and trailing whitespace from a string.
func TrimString(str string) string {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.char)
	var goret string

	cret = C.pango_trim_string(arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// Version returns the encoded version of Pango available at run-time.
//
// This is similar to the macro PANGO_VERSION except that the macro returns the
// encoded version available at compile-time. A version number can be encoded
// into an integer using PANGO_VERSION_ENCODE().
func Version() int {
	var cret C.int
	var goret int

	cret = C.pango_version()

	goret = int(cret)

	return goret
}

// VersionCheck checks that the Pango library in use is compatible with the
// given version.
//
// Generally you would pass in the constants PANGO_VERSION_MAJOR,
// PANGO_VERSION_MINOR, PANGO_VERSION_MICRO as the three arguments to this
// function; that produces a check that the library in use at run-time is
// compatible with the version of Pango the application or module was compiled
// against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// @required_major.required_minor.@required_micro. Second the running library
// must be binary compatible with the version
// @required_major.required_minor.@required_micro (same major version.)
//
// For compile-time version checking use PANGO_VERSION_CHECK().
func VersionCheck(requiredMajor int, requiredMinor int, requiredMicro int) string {
	var arg1 C.int
	var arg2 C.int
	var arg3 C.int

	arg1 = C.int(requiredMajor)
	arg2 = C.int(requiredMinor)
	arg3 = C.int(requiredMicro)

	var cret *C.char
	var goret string

	cret = C.pango_version_check(arg1, arg2, arg3)

	goret = C.GoString(cret)

	return goret
}

// VersionString returns the version of Pango available at run-time.
//
// This is similar to the macro PANGO_VERSION_STRING except that the macro
// returns the version available at compile-time.
func VersionString() string {
	var cret *C.char
	var goret string

	cret = C.pango_version_string()

	goret = C.GoString(cret)

	return goret
}
