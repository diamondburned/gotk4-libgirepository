// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_option_group_get_type()), F: marshalOptionGroup},
	})
}

// OptionArg: the Arg enum values determine which type of extra argument the
// options expect to find. If an option expects an extra argument, it can be
// specified in several ways; with a short option: `-x arg`, with a long option:
// `--name arg` or combined in a single argument: `--name=arg`.
type OptionArg int

const (
	// None: no extra argument. This is useful for simple flags.
	OptionArgNone OptionArg = iota
	// String: the option takes a UTF-8 string argument.
	OptionArgString
	// Int: the option takes an integer argument.
	OptionArgInt
	// Callback: the option provides a callback (of type ArgFunc) to parse the
	// extra argument.
	OptionArgCallback
	// Filename: the option takes a filename as argument, which will be in the
	// GLib filename encoding rather than UTF-8.
	OptionArgFilename
	// StringArray: the option takes a string argument, multiple uses of the
	// option are collected into an array of strings.
	OptionArgStringArray
	// FilenameArray: the option takes a filename as argument, multiple uses of
	// the option are collected into an array of strings.
	OptionArgFilenameArray
	// Double: the option takes a double argument. The argument can be formatted
	// either for the user's locale or for the "C" locale. Since 2.12
	OptionArgDouble
	// Int64: the option takes a 64-bit integer. Like G_OPTION_ARG_INT but for
	// larger numbers. The number can be in decimal base, or in hexadecimal
	// (when prefixed with `0x`, for example, `0xffffffff`). Since 2.12
	OptionArgInt64
)

// OptionError: error codes returned by option parsing.
type OptionError int

const (
	// UnknownOption: option was not known to the parser. This error will only
	// be reported, if the parser hasn't been instructed to ignore unknown
	// options, see g_option_context_set_ignore_unknown_options().
	OptionErrorUnknownOption OptionError = iota
	// BadValue: value couldn't be parsed.
	OptionErrorBadValue
	// Failed callback failed.
	OptionErrorFailed
)

// OptionFlags flags which modify individual options.
type OptionFlags int

const (
	// OptionFlagsNone: no flags. Since: 2.42.
	OptionFlagsNone OptionFlags = 0b0
	// OptionFlagsHidden: the option doesn't appear in `--help` output.
	OptionFlagsHidden OptionFlags = 0b1
	// OptionFlagsInMain: the option appears in the main section of the `--help`
	// output, even if it is defined in a group.
	OptionFlagsInMain OptionFlags = 0b10
	// OptionFlagsReverse: for options of the G_OPTION_ARG_NONE kind, this flag
	// indicates that the sense of the option is reversed.
	OptionFlagsReverse OptionFlags = 0b100
	// OptionFlagsNoArg: for options of the G_OPTION_ARG_CALLBACK kind, this
	// flag indicates that the callback does not take any argument (like a
	// G_OPTION_ARG_NONE option). Since 2.8
	OptionFlagsNoArg OptionFlags = 0b1000
	// OptionFlagsFilename: for options of the G_OPTION_ARG_CALLBACK kind, this
	// flag indicates that the argument should be passed to the callback in the
	// GLib filename encoding rather than UTF-8. Since 2.8
	OptionFlagsFilename OptionFlags = 0b10000
	// OptionFlagsOptionalArg: for options of the G_OPTION_ARG_CALLBACK kind,
	// this flag indicates that the argument supply is optional. If no argument
	// is given then data of GOptionParseFunc will be set to NULL. Since 2.8
	OptionFlagsOptionalArg OptionFlags = 0b100000
	// OptionFlagsNoalias: this flag turns off the automatic conflict resolution
	// which prefixes long option names with `groupname-` if there is a
	// conflict. This option should only be used in situations where aliasing is
	// necessary to model some legacy commandline interface. It is not safe to
	// use this option, unless all option groups are under your direct control.
	// Since 2.8.
	OptionFlagsNoalias OptionFlags = 0b1000000
)

// OptionEntry struct defines a single option. To have an effect, they must be
// added to a Group with g_option_context_add_main_entries() or
// g_option_group_add_entries().
type OptionEntry struct {
	native C.GOptionEntry
}

// WrapOptionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOptionEntry(ptr unsafe.Pointer) *OptionEntry {
	return (*OptionEntry)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OptionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// LongName: the long name of an option can be used to specify it in a
// commandline as `--long_name`. Every option must have a long name. To resolve
// conflicts if multiple option groups contain the same long name, it is also
// possible to specify the option as `--groupname-long_name`.
func (o *OptionEntry) LongName() string {
	var v string // out
	v = C.GoString(o.long_name)
	return v
}

// ShortName: if an option has a short name, it can be specified `-short_name`
// in a commandline. @short_name must be a printable ASCII character different
// from '-', or zero if the option has no short name.
func (o *OptionEntry) ShortName() byte {
	var v byte // out
	v = byte(o.short_name)
	return v
}

// Flags from Flags
func (o *OptionEntry) Flags() int {
	var v int // out
	v = int(o.flags)
	return v
}

// Arg: the type of the option, as a Arg
func (o *OptionEntry) Arg() OptionArg {
	var v OptionArg // out
	v = OptionArg(o.arg)
	return v
}

// ArgData: if the @arg type is G_OPTION_ARG_CALLBACK, then @arg_data must point
// to a ArgFunc callback function, which will be called to handle the extra
// argument. Otherwise, @arg_data is a pointer to a location to store the value,
// the required type of the location depends on the @arg type: -
// G_OPTION_ARG_NONE: gboolean - G_OPTION_ARG_STRING: gchar* - G_OPTION_ARG_INT:
// gint - G_OPTION_ARG_FILENAME: gchar* - G_OPTION_ARG_STRING_ARRAY: gchar** -
// G_OPTION_ARG_FILENAME_ARRAY: gchar** - G_OPTION_ARG_DOUBLE: gdouble If @arg
// type is G_OPTION_ARG_STRING or G_OPTION_ARG_FILENAME, the location will
// contain a newly allocated string if the option was given. That string needs
// to be freed by the callee using g_free(). Likewise if @arg type is
// G_OPTION_ARG_STRING_ARRAY or G_OPTION_ARG_FILENAME_ARRAY, the data should be
// freed using g_strfreev().
func (o *OptionEntry) ArgData() interface{} {
	var v interface{} // out
	v = box.Get(uintptr(o.arg_data))
	return v
}

// Description: the description for the option in `--help` output. The
// @description is translated using the @translate_func of the group, see
// g_option_group_set_translation_domain().
func (o *OptionEntry) Description() string {
	var v string // out
	v = C.GoString(o.description)
	return v
}

// ArgDescription: the placeholder to use for the extra argument parsed by the
// option in `--help` output. The @arg_description is translated using the
// @translate_func of the group, see g_option_group_set_translation_domain().
func (o *OptionEntry) ArgDescription() string {
	var v string // out
	v = C.GoString(o.arg_description)
	return v
}

// OptionGroup: `GOptionGroup` struct defines the options in a single group. The
// struct has only private fields and should not be directly accessed.
//
// All options in a group share the same translation function. Libraries which
// need to parse commandline options are expected to provide a function for
// getting a `GOptionGroup` holding their options, which the application can
// then add to its Context.
type OptionGroup struct {
	native C.GOptionGroup
}

// WrapOptionGroup wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOptionGroup(ptr unsafe.Pointer) *OptionGroup {
	return (*OptionGroup)(ptr)
}

func marshalOptionGroup(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*OptionGroup)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (o *OptionGroup) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// AddEntries adds the options specified in @entries to @group.
func (g *OptionGroup) AddEntries(entries []OptionEntry) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.GOptionEntry

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g))
	{
		var zero OptionEntry
		entries = append(entries, zero)
	}
	_arg1 = (*C.GOptionEntry)(unsafe.Pointer(&entries[0]))

	C.g_option_group_add_entries(_arg0, _arg1)
}

// Free frees a Group. Note that you must not free groups which have been added
// to a Context.
//
// Deprecated: since version 2.44.
func (g *OptionGroup) free() {
	var _arg0 *C.GOptionGroup // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g))

	C.g_option_group_free(_arg0)
}

// Ref increments the reference count of @group by one.
func (g *OptionGroup) ref() *OptionGroup {
	var _arg0 *C.GOptionGroup // out
	var _cret *C.GOptionGroup // in

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g))

	_cret = C.g_option_group_ref(_arg0)

	var _optionGroup *OptionGroup // out

	_optionGroup = (*OptionGroup)(unsafe.Pointer(_cret))
	C.g_option_group_ref(_cret)
	runtime.SetFinalizer(_optionGroup, func(v *OptionGroup) {
		C.g_option_group_unref((*C.GOptionGroup)(unsafe.Pointer(v)))
	})

	return _optionGroup
}

// SetTranslationDomain: convenience function to use gettext() for translating
// user-visible strings.
func (g *OptionGroup) SetTranslationDomain(domain string) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g))
	_arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_group_set_translation_domain(_arg0, _arg1)
}

// Unref decrements the reference count of @group by one. If the reference count
// drops to 0, the @group will be freed. and all memory allocated by the @group
// is released.
func (g *OptionGroup) unref() {
	var _arg0 *C.GOptionGroup // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g))

	C.g_option_group_unref(_arg0)
}
