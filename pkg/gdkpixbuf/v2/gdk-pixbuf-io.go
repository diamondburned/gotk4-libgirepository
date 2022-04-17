// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for gdk-pixbuf-io.go.
var GTypePixbufFormat = externglib.Type(C.gdk_pixbuf_format_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePixbufFormat, F: marshalPixbufFormat},
	})
}

// PixbufFormatFlags flags which allow a module to specify further details about
// the supported operations.
type PixbufFormatFlags C.guint

const (
	// PixbufFormatWritable: module can write out images in the format.
	PixbufFormatWritable PixbufFormatFlags = 0b1
	// PixbufFormatScalable: image format is scalable.
	PixbufFormatScalable PixbufFormatFlags = 0b10
	// PixbufFormatThreadsafe: module is threadsafe. gdk-pixbuf ignores modules
	// that are not marked as threadsafe. (Since 2.28).
	PixbufFormatThreadsafe PixbufFormatFlags = 0b100
)

// String returns the names in string for PixbufFormatFlags.
func (p PixbufFormatFlags) String() string {
	if p == 0 {
		return "PixbufFormatFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(64)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PixbufFormatWritable:
			builder.WriteString("Writable|")
		case PixbufFormatScalable:
			builder.WriteString("Scalable|")
		case PixbufFormatThreadsafe:
			builder.WriteString("Threadsafe|")
		default:
			builder.WriteString(fmt.Sprintf("PixbufFormatFlags(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PixbufFormatFlags) Has(other PixbufFormatFlags) bool {
	return (p & other) == other
}

// PixbufModulePreparedFunc defines the type of the function that gets called
// once the initial setup of pixbuf is done.
//
// PixbufLoader uses a function of this type to emit the "<link
// linkend="GdkPixbufLoader-area-prepared">area_prepared</link>" signal.
type PixbufModulePreparedFunc func(pixbuf *Pixbuf, anim *PixbufAnimation)

//export _gotk4_gdkpixbuf2_PixbufModulePreparedFunc
func _gotk4_gdkpixbuf2_PixbufModulePreparedFunc(arg1 *C.GdkPixbuf, arg2 *C.GdkPixbufAnimation, arg3 C.gpointer) {
	var fn PixbufModulePreparedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PixbufModulePreparedFunc)
	}

	var _pixbuf *Pixbuf        // out
	var _anim *PixbufAnimation // out

	_pixbuf = wrapPixbuf(externglib.Take(unsafe.Pointer(arg1)))
	_anim = wrapPixbufAnimation(externglib.Take(unsafe.Pointer(arg2)))

	fn(_pixbuf, _anim)
}

// PixbufModuleSizeFunc defines the type of the function that gets called once
// the size of the loaded image is known.
//
// The function is expected to set width and height to the desired size to which
// the image should be scaled. If a module has no efficient way to achieve the
// desired scaling during the loading of the image, it may either ignore the
// size request, or only approximate it - gdk-pixbuf will then perform the
// required scaling on the completely loaded image.
//
// If the function sets width or height to zero, the module should interpret
// this as a hint that it will be closed soon and shouldn't allocate further
// resources. This convention is used to implement gdk_pixbuf_get_file_info()
// efficiently.
type PixbufModuleSizeFunc func(width, height *int)

//export _gotk4_gdkpixbuf2_PixbufModuleSizeFunc
func _gotk4_gdkpixbuf2_PixbufModuleSizeFunc(arg1 *C.gint, arg2 *C.gint, arg3 C.gpointer) {
	var fn PixbufModuleSizeFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PixbufModuleSizeFunc)
	}

	var _width *int  // out
	var _height *int // out

	_width = (*int)(unsafe.Pointer(arg1))
	_height = (*int)(unsafe.Pointer(arg2))

	fn(_width, _height)
}

// PixbufModuleUpdatedFunc defines the type of the function that gets called
// every time a region of pixbuf is updated.
//
// PixbufLoader uses a function of this type to emit the "<link
// linkend="GdkPixbufLoader-area-updated">area_updated</link>" signal.
type PixbufModuleUpdatedFunc func(pixbuf *Pixbuf, x, y, width, height int)

//export _gotk4_gdkpixbuf2_PixbufModuleUpdatedFunc
func _gotk4_gdkpixbuf2_PixbufModuleUpdatedFunc(arg1 *C.GdkPixbuf, arg2 C.int, arg3 C.int, arg4 C.int, arg5 C.int, arg6 C.gpointer) {
	var fn PixbufModuleUpdatedFunc
	{
		v := gbox.Get(uintptr(arg6))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PixbufModuleUpdatedFunc)
	}

	var _pixbuf *Pixbuf // out
	var _x int          // out
	var _y int          // out
	var _width int      // out
	var _height int     // out

	_pixbuf = wrapPixbuf(externglib.Take(unsafe.Pointer(arg1)))
	_x = int(arg2)
	_y = int(arg3)
	_width = int(arg4)
	_height = int(arg5)

	fn(_pixbuf, _x, _y, _width, _height)
}

// PixbufGetFileInfo parses an image file far enough to determine its format and
// size.
//
// The function takes the following parameters:
//
//    - filename: name of the file to identify.
//
// The function returns the following values:
//
//    - width (optional): return location for the width of the image.
//    - height (optional): return location for the height of the image.
//    - pixbufFormat (optional): GdkPixbufFormat describing the image format of
//      the file.
//
func PixbufGetFileInfo(filename string) (width int, height int, pixbufFormat *PixbufFormat) {
	var _arg1 *C.gchar           // out
	var _arg2 C.gint             // in
	var _arg3 C.gint             // in
	var _cret *C.GdkPixbufFormat // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_get_file_info(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(filename)

	var _width int                  // out
	var _height int                 // out
	var _pixbufFormat *PixbufFormat // out

	_width = int(_arg2)
	_height = int(_arg3)
	if _cret != nil {
		_pixbufFormat = (*PixbufFormat)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _width, _height, _pixbufFormat
}

// PixbufGetFileInfoAsync: asynchronously parses an image file far enough to
// determine its format and size.
//
// For more details see gdk_pixbuf_get_file_info(), which is the synchronous
// version of this function.
//
// When the operation is finished, callback will be called in the main thread.
// You can then call gdk_pixbuf_get_file_info_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - filename: name of the file to identify.
//    - callback (optional): GAsyncReadyCallback to call when the file info is
//      available.
//
func PixbufGetFileInfoAsync(ctx context.Context, filename string, callback gio.AsyncReadyCallback) {
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_pixbuf_get_file_info_async(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(callback)
}

// PixbufGetFileInfoFinish finishes an asynchronous pixbuf parsing operation
// started with gdk_pixbuf_get_file_info_async().
//
// The function takes the following parameters:
//
//    - asyncResult: GAsyncResult.
//
// The function returns the following values:
//
//    - width: return location for the width of the image, or NULL.
//    - height: return location for the height of the image, or NULL.
//    - pixbufFormat (optional): GdkPixbufFormat describing the image format of
//      the file.
//
func PixbufGetFileInfoFinish(asyncResult gio.AsyncResultOverrider) (width int, height int, pixbufFormat *PixbufFormat, goerr error) {
	var _arg1 *C.GAsyncResult    // out
	var _arg2 C.gint             // in
	var _arg3 C.gint             // in
	var _cret *C.GdkPixbufFormat // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(asyncResult).Native()))

	_cret = C.gdk_pixbuf_get_file_info_finish(_arg1, &_arg2, &_arg3, &_cerr)
	runtime.KeepAlive(asyncResult)

	var _width int                  // out
	var _height int                 // out
	var _pixbufFormat *PixbufFormat // out
	var _goerr error                // out

	_width = int(_arg2)
	_height = int(_arg3)
	if _cret != nil {
		_pixbufFormat = (*PixbufFormat)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _width, _height, _pixbufFormat, _goerr
}

// PixbufGetFormats obtains the available information about the image formats
// supported by GdkPixbuf.
//
// The function returns the following values:
//
//    - sList: list of support image formats.
//
func PixbufGetFormats() []*PixbufFormat {
	var _cret *C.GSList // in

	_cret = C.gdk_pixbuf_get_formats()

	var _sList []*PixbufFormat // out

	_sList = make([]*PixbufFormat, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkPixbufFormat)(v)
		var dst *PixbufFormat // out
		dst = (*PixbufFormat)(gextras.NewStructNative(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// PixbufInitModules initalizes the gdk-pixbuf loader modules referenced by the
// loaders.cache file present inside that directory.
//
// This is to be used by applications that want to ship certain loaders in a
// different location from the system ones.
//
// This is needed when the OS or runtime ships a minimal number of loaders so as
// to reduce the potential attack surface of carefully crafted image files,
// especially for uncommon file types. Applications that require broader image
// file types coverage, such as image viewers, would be expected to ship the
// gdk-pixbuf modules in a separate location, bundled with the application in a
// separate directory from the OS or runtime- provided modules.
//
// The function takes the following parameters:
//
//    - path: path to directory where the loaders.cache is installed.
//
func PixbufInitModules(path string) error {
	var _arg1 *C.char   // out
	var _cerr *C.GError // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_pixbuf_init_modules(_arg1, &_cerr)
	runtime.KeepAlive(path)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PixbufFormat: GdkPixbufFormat contains information about the image format
// accepted by a module.
//
// Only modules should access the fields directly, applications should use the
// gdk_pixbuf_format_* family of functions.
//
// An instance of this type is always passed by reference.
type PixbufFormat struct {
	*pixbufFormat
}

// pixbufFormat is the struct that's finalized.
type pixbufFormat struct {
	native *C.GdkPixbufFormat
}

func marshalPixbufFormat(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PixbufFormat{&pixbufFormat{(*C.GdkPixbufFormat)(b)}}, nil
}

// Copy creates a copy of format.
//
// The function returns the following values:
//
//    - pixbufFormat: newly allocated copy of a GdkPixbufFormat. Use
//      gdk_pixbuf_format_free() to free the resources when done.
//
func (format *PixbufFormat) Copy() *PixbufFormat {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.GdkPixbufFormat // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_copy(_arg0)
	runtime.KeepAlive(format)

	var _pixbufFormat *PixbufFormat // out

	_pixbufFormat = (*PixbufFormat)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_pixbufFormat)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_pixbuf_format_free((*C.GdkPixbufFormat)(intern.C))
		},
	)

	return _pixbufFormat
}

// Description returns a description of the format.
//
// The function returns the following values:
//
//    - utf8: description of the format.
//
func (format *PixbufFormat) Description() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_get_description(_arg0)
	runtime.KeepAlive(format)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Extensions returns the filename extensions typically used for files in the
// given format.
//
// The function returns the following values:
//
//    - utf8s: array of filename extensions.
//
func (format *PixbufFormat) Extensions() []string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_get_extensions(_arg0)
	runtime.KeepAlive(format)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
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

// License returns information about the license of the image loader for the
// format.
//
// The returned string should be a shorthand for a well known license, e.g.
// "LGPL", "GPL", "QPL", "GPL/QPL", or "other" to indicate some other license.
//
// The function returns the following values:
//
//    - utf8: string describing the license of the pixbuf format.
//
func (format *PixbufFormat) License() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_get_license(_arg0)
	runtime.KeepAlive(format)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MIMETypes returns the mime types supported by the format.
//
// The function returns the following values:
//
//    - utf8s: array of mime types.
//
func (format *PixbufFormat) MIMETypes() []string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_get_mime_types(_arg0)
	runtime.KeepAlive(format)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
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

// Name returns the name of the format.
//
// The function returns the following values:
//
//    - utf8: name of the format.
//
func (format *PixbufFormat) Name() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_get_name(_arg0)
	runtime.KeepAlive(format)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// IsDisabled returns whether this image format is disabled.
//
// See gdk_pixbuf_format_set_disabled().
//
// The function returns the following values:
//
//    - ok: whether this image format is disabled.
//
func (format *PixbufFormat) IsDisabled() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_is_disabled(_arg0)
	runtime.KeepAlive(format)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSaveOptionSupported returns TRUE if the save option specified by option_key
// is supported when saving a pixbuf using the module implementing format.
//
// See gdk_pixbuf_save() for more information about option keys.
//
// The function takes the following parameters:
//
//    - optionKey: name of an option.
//
// The function returns the following values:
//
//    - ok: TRUE if the specified option is supported.
//
func (format *PixbufFormat) IsSaveOptionSupported(optionKey string) bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(optionKey)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_format_is_save_option_supported(_arg0, _arg1)
	runtime.KeepAlive(format)
	runtime.KeepAlive(optionKey)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsScalable returns whether this image format is scalable.
//
// If a file is in a scalable format, it is preferable to load it at the desired
// size, rather than loading it at the default size and scaling the resulting
// pixbuf to the desired size.
//
// The function returns the following values:
//
//    - ok: whether this image format is scalable.
//
func (format *PixbufFormat) IsScalable() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_is_scalable(_arg0)
	runtime.KeepAlive(format)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsWritable returns whether pixbufs can be saved in the given format.
//
// The function returns the following values:
//
//    - ok: whether pixbufs can be saved in the given format.
//
func (format *PixbufFormat) IsWritable() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))

	_cret = C.gdk_pixbuf_format_is_writable(_arg0)
	runtime.KeepAlive(format)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDisabled disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
//
// The function takes the following parameters:
//
//    - disabled: TRUE to disable the format format.
//
func (format *PixbufFormat) SetDisabled(disabled bool) {
	var _arg0 *C.GdkPixbufFormat // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GdkPixbufFormat)(gextras.StructNative(unsafe.Pointer(format)))
	if disabled {
		_arg1 = C.TRUE
	}

	C.gdk_pixbuf_format_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(format)
	runtime.KeepAlive(disabled)
}
