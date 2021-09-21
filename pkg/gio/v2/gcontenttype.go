// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

// ContentTypeCanBeExecutable checks if a content type can be executable. Note
// that for instance things like text files can be executables (i.e. scripts and
// batch files).
func ContentTypeCanBeExecutable(typ string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_can_be_executable(_arg1)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeEquals compares two content types for equality.
func ContentTypeEquals(type1 string, type2 string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(type1)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(type2)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_content_type_equals(_arg1, _arg2)
	runtime.KeepAlive(type1)
	runtime.KeepAlive(type2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeFromMIMEType tries to find a content type based on the mime type
// name.
func ContentTypeFromMIMEType(mimeType string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_from_mime_type(_arg1)
	runtime.KeepAlive(mimeType)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ContentTypeGetDescription gets the human readable description of the content
// type.
func ContentTypeGetDescription(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_description(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetGenericIconName gets the generic icon name for a content type.
//
// See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on the generic icon name.
func ContentTypeGetGenericIconName(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_generic_icon_name(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ContentTypeGetIcon gets the icon for a content type.
func ContentTypeGetIcon(typ string) Iconner {
	var _arg1 *C.gchar // out
	var _cret *C.GIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_icon(_arg1)
	runtime.KeepAlive(typ)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// ContentTypeGetMIMEDirs: get the list of directories which MIME data is loaded
// from. See g_content_type_set_mime_dirs() for details.
func ContentTypeGetMIMEDirs() []string {
	var _cret **C.gchar // in

	_cret = C.g_content_type_get_mime_dirs()

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// ContentTypeGetMIMEType gets the mime type for the content type, if one is
// registered.
func ContentTypeGetMIMEType(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_mime_type(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ContentTypeGetSymbolicIcon gets the symbolic icon for a content type.
func ContentTypeGetSymbolicIcon(typ string) Iconner {
	var _arg1 *C.gchar // out
	var _cret *C.GIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_symbolic_icon(_arg1)
	runtime.KeepAlive(typ)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// ContentTypeGuess guesses the content type based on example data. If the
// function is uncertain, result_uncertain will be set to TRUE. Either filename
// or data may be NULL, in which case the guess will be based solely on the
// other argument.
func ContentTypeGuess(filename string, data []byte) (bool, string) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.guchar // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // in
	var _cret *C.gchar   // in

	if filename != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_content_type_guess(_arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(data)

	var _resultUncertain bool // out
	var _utf8 string          // out

	if _arg4 != 0 {
		_resultUncertain = true
	}
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _resultUncertain, _utf8
}

// ContentTypeGuessForTree tries to guess the type of the tree with root root,
// by looking at the files it contains. The result is an array of content types,
// with the best guess coming first.
//
// The types returned all have the form x-content/foo, e.g. x-content/audio-cdda
// (for audio CDs) or x-content/image-dcf (for a camera memory card). See the
// shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This function is useful in the implementation of
// g_mount_guess_content_type().
func ContentTypeGuessForTree(root Filer) []string {
	var _arg1 *C.GFile  // out
	var _cret **C.gchar // in

	_arg1 = (*C.GFile)(unsafe.Pointer(root.Native()))

	_cret = C.g_content_type_guess_for_tree(_arg1)
	runtime.KeepAlive(root)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ContentTypeIsA determines if type is a subset of supertype.
func ContentTypeIsA(typ string, supertype string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(supertype)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_content_type_is_a(_arg1, _arg2)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(supertype)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeIsMIMEType determines if type is a subset of mime_type.
// Convenience wrapper around g_content_type_is_a().
func ContentTypeIsMIMEType(typ string, mimeType string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_content_type_is_mime_type(_arg1, _arg2)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(mimeType)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeIsUnknown checks if the content type is the generic "unknown"
// type. On UNIX this is the "application/octet-stream" mimetype, while on win32
// it is "*" and on OSX it is a dynamic type or octet-stream.
func ContentTypeIsUnknown(typ string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_is_unknown(_arg1)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeSetMIMEDirs: set the list of directories used by GIO to load the
// MIME database. If dirs is NULL, the directories used are the default:
//
//    - the mime subdirectory of the directory in $XDG_DATA_HOME
//    - the mime subdirectory of every directory in $XDG_DATA_DIRS
//
// This function is intended to be used when writing tests that depend on
// information stored in the MIME database, in order to control the data.
//
// Typically, in case your tests use G_TEST_OPTION_ISOLATE_DIRS, but they depend
// on the system’s MIME database, you should call this function with dirs set to
// NULL before calling g_test_init(), for instance:
//
//      // Load MIME data from the system
//      g_content_type_set_mime_dirs (NULL);
//      // Isolate the environment
//      g_test_init (&argc, &argv, G_TEST_OPTION_ISOLATE_DIRS, NULL);
//
//      …
//
//      return g_test_run ();
func ContentTypeSetMIMEDirs(dirs []string) {
	var _arg1 **C.gchar // out

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(dirs)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(dirs)+1)
			var zero *C.gchar
			out[len(dirs)] = zero
			for i := range dirs {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(dirs[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.g_content_type_set_mime_dirs(_arg1)
	runtime.KeepAlive(dirs)
}

// ContentTypesGetRegistered gets a list of strings containing all the
// registered content types known to the system. The list and its data should be
// freed using g_list_free_full (list, g_free).
func ContentTypesGetRegistered() []string {
	var _cret *C.GList // in

	_cret = C.g_content_types_get_registered()

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		defer C.free(unsafe.Pointer(src))
		_list = append(_list, dst)
	})

	return _list
}
