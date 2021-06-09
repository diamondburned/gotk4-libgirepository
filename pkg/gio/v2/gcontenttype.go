// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
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
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_content_type_can_be_executable(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContentTypeEquals compares two content types for equality.
func ContentTypeEquals(type1 string, type2 string) bool {
	var _arg1 *C.gchar
	var _arg2 *C.gchar

	_arg1 = (*C.gchar)(C.CString(type1))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(type2))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.gboolean

	cret = C.g_content_type_equals(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContentTypeFromMIMEType tries to find a content type based on the mime type
// name.
func ContentTypeFromMIMEType(mimeType string) string {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.g_content_type_from_mime_type(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetDescription gets the human readable description of the content
// type.
func ContentTypeGetDescription(typ string) string {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.g_content_type_get_description(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetGenericIconName gets the generic icon name for a content type.
//
// See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on the generic icon name.
func ContentTypeGetGenericIconName(typ string) string {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.g_content_type_get_generic_icon_name(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetIcon gets the icon for a content type.
func ContentTypeGetIcon(typ string) Icon {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GIcon

	cret = C.g_content_type_get_icon(_arg1)

	var _icon Icon

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Icon)

	return _icon
}

// ContentTypeGetMIMEDirs: get the list of directories which MIME data is loaded
// from. See g_content_type_set_mime_dirs() for details.
func ContentTypeGetMIMEDirs() []string {
	var _cret **C.gchar

	cret = C.g_content_type_get_mime_dirs()

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
		}
	}

	return _utf8s
}

// ContentTypeGetMIMEType gets the mime type for the content type, if one is
// registered.
func ContentTypeGetMIMEType(typ string) string {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.g_content_type_get_mime_type(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetSymbolicIcon gets the symbolic icon for a content type.
func ContentTypeGetSymbolicIcon(typ string) Icon {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GIcon

	cret = C.g_content_type_get_symbolic_icon(_arg1)

	var _icon Icon

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Icon)

	return _icon
}

// ContentTypeGuessForTree tries to guess the type of the tree with root @root,
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
func ContentTypeGuessForTree(root File) []string {
	var _arg1 *C.GFile

	_arg1 = (*C.GFile)(unsafe.Pointer(root.Native()))

	var _cret **C.gchar

	cret = C.g_content_type_guess_for_tree(_arg1)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// ContentTypeIsA determines if @type is a subset of @supertype.
func ContentTypeIsA(typ string, supertype string) bool {
	var _arg1 *C.gchar
	var _arg2 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(supertype))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.gboolean

	cret = C.g_content_type_is_a(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContentTypeIsMIMEType determines if @type is a subset of @mime_type.
// Convenience wrapper around g_content_type_is_a().
func ContentTypeIsMIMEType(typ string, mimeType string) bool {
	var _arg1 *C.gchar
	var _arg2 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.gboolean

	cret = C.g_content_type_is_mime_type(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContentTypeIsUnknown checks if the content type is the generic "unknown"
// type. On UNIX this is the "application/octet-stream" mimetype, while on win32
// it is "*" and on OSX it is a dynamic type or octet-stream.
func ContentTypeIsUnknown(typ string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_content_type_is_unknown(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContentTypeSetMIMEDirs: set the list of directories used by GIO to load the
// MIME database. If @dirs is nil, the directories used are the default:
//
//    - the `mime` subdirectory of the directory in `$XDG_DATA_HOME`
//    - the `mime` subdirectory of every directory in `$XDG_DATA_DIRS`
//
// This function is intended to be used when writing tests that depend on
// information stored in the MIME database, in order to control the data.
//
// Typically, in case your tests use G_TEST_OPTION_ISOLATE_DIRS, but they depend
// on the system’s MIME database, you should call this function with @dirs set
// to nil before calling g_test_init(), for instance:
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
	var _arg1 **C.gchar

	_arg1 = (**C.gchar)(C.malloc((len(dirs) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(_arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(_arg1), int(len(dirs)))

		for i := range dirs {
			_arg1 = (*C.gchar)(C.CString(dirs))
			defer C.free(unsafe.Pointer(_arg1))
		}
	}

	C.g_content_type_set_mime_dirs(_arg1)
}

// ContentTypesGetRegistered gets a list of strings containing all the
// registered content types known to the system. The list and its data should be
// freed using `g_list_free_full (list, g_free)`.
func ContentTypesGetRegistered() *glib.List {
	var _cret *C.GList

	cret = C.g_content_types_get_registered()

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _list
}