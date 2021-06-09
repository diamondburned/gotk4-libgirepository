// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_themed_icon_get_type()), F: marshalThemedIcon},
	})
}

// ThemedIcon is an implementation of #GIcon that supports icon themes. Icon
// contains a list of all of the icons present in an icon theme, so that icons
// can be looked up quickly. Icon does not provide actual pixmaps for icons,
// just the icon names. Ideally something like gtk_icon_theme_choose_icon()
// should be used to resolve the list of names so that fallback icons work
// nicely with themes that inherit other themes.
type ThemedIcon interface {
	gextras.Objector
	Icon

	// AppendName: append a name to the list of icons from within @icon.
	//
	// Note that doing so invalidates the hash computed by prior calls to
	// g_icon_hash().
	AppendName(iconname string)
	// Names gets the names of icons from within @icon.
	Names() []string
	// PrependName: prepend a name to the list of icons from within @icon.
	//
	// Note that doing so invalidates the hash computed by prior calls to
	// g_icon_hash().
	PrependName(iconname string)
}

// themedIcon implements the ThemedIcon interface.
type themedIcon struct {
	gextras.Objector
	Icon
}

var _ ThemedIcon = (*themedIcon)(nil)

// WrapThemedIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapThemedIcon(obj *externglib.Object) ThemedIcon {
	return ThemedIcon{
		Objector: obj,
		Icon:     WrapIcon(obj),
	}
}

func marshalThemedIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapThemedIcon(obj), nil
}

// NewThemedIcon constructs a class ThemedIcon.
func NewThemedIcon(iconname string) ThemedIcon {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GThemedIcon

	cret = C.g_themed_icon_new(_arg1)

	var _themedIcon ThemedIcon

	_themedIcon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ThemedIcon)

	return _themedIcon
}

// NewThemedIconFromNames constructs a class ThemedIcon.
func NewThemedIconFromNames() ThemedIcon {
	var _cret C.GThemedIcon

	cret = C.g_themed_icon_new_from_names()

	var _themedIcon ThemedIcon

	_themedIcon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ThemedIcon)

	return _themedIcon
}

// NewThemedIconWithDefaultFallbacks constructs a class ThemedIcon.
func NewThemedIconWithDefaultFallbacks(iconname string) ThemedIcon {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GThemedIcon

	cret = C.g_themed_icon_new_with_default_fallbacks(_arg1)

	var _themedIcon ThemedIcon

	_themedIcon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ThemedIcon)

	return _themedIcon
}

// AppendName: append a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
func (i themedIcon) AppendName(iconname string) {
	var _arg0 *C.GThemedIcon
	var _arg1 *C.char

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_themed_icon_append_name(_arg0, _arg1)
}

// Names gets the names of icons from within @icon.
func (i themedIcon) Names() []string {
	var _arg0 *C.GThemedIcon

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(i.Native()))

	var _cret **C.gchar

	cret = C.g_themed_icon_get_names(_arg0)

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

// PrependName: prepend a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
func (i themedIcon) PrependName(iconname string) {
	var _arg0 *C.GThemedIcon
	var _arg1 *C.char

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_themed_icon_prepend_name(_arg0, _arg1)
}