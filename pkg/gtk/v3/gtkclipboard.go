// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// void _gotk4_gtk3_ClipboardImageReceivedFunc(GtkClipboard*, GdkPixbuf*, gpointer);
// void _gotk4_gtk3_ClipboardTextReceivedFunc(GtkClipboard*, gchar*, gpointer);
// void _gotk4_gtk3_ClipboardURIReceivedFunc(GtkClipboard*, gchar**, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_clipboard_get_type()), F: marshalClipboarder},
	})
}

// ClipboardImageReceivedFunc: function to be called when the results of
// gtk_clipboard_request_image() are received, or when the request fails.
type ClipboardImageReceivedFunc func(clipboard *Clipboard, pixbuf *gdkpixbuf.Pixbuf)

//export _gotk4_gtk3_ClipboardImageReceivedFunc
func _gotk4_gtk3_ClipboardImageReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GdkPixbuf, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard *Clipboard     // out
	var pixbuf *gdkpixbuf.Pixbuf // out

	clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg0)))
	{
		obj := externglib.Take(unsafe.Pointer(arg1))
		pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	fn := v.(ClipboardImageReceivedFunc)
	fn(clipboard, pixbuf)
}

// ClipboardReceivedFunc: function to be called when the results of
// gtk_clipboard_request_contents() are received, or when the request fails.
type ClipboardReceivedFunc func(clipboard *Clipboard, selectionData *SelectionData)

//export _gotk4_gtk3_ClipboardReceivedFunc
func _gotk4_gtk3_ClipboardReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GtkSelectionData, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard *Clipboard         // out
	var selectionData *SelectionData // out

	clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg0)))
	selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(selectionData, func(v *SelectionData) {
		C.gtk_selection_data_free((*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(v))))
	})

	fn := v.(ClipboardReceivedFunc)
	fn(clipboard, selectionData)
}

// ClipboardTextReceivedFunc: function to be called when the results of
// gtk_clipboard_request_text() are received, or when the request fails.
type ClipboardTextReceivedFunc func(clipboard *Clipboard, text string)

//export _gotk4_gtk3_ClipboardTextReceivedFunc
func _gotk4_gtk3_ClipboardTextReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard *Clipboard // out
	var text string          // out

	clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg0)))
	text = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))

	fn := v.(ClipboardTextReceivedFunc)
	fn(clipboard, text)
}

// ClipboardURIReceivedFunc: function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard *Clipboard, uris []string)

//export _gotk4_gtk3_ClipboardURIReceivedFunc
func _gotk4_gtk3_ClipboardURIReceivedFunc(arg0 *C.GtkClipboard, arg1 **C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard *Clipboard // out
	var uris []string        // out

	clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg0)))
	{
		var i int
		var z *C.gchar
		for p := arg1; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(arg1, i)
		uris = make([]string, i)
		for i := range src {
			uris[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	fn := v.(ClipboardURIReceivedFunc)
	fn(clipboard, uris)
}

// Clipboard object represents a clipboard of data shared between different
// processes or between different widgets in the same process. Each clipboard is
// identified by a name encoded as a Atom. (Conversion to and from strings can
// be done with gdk_atom_intern() and gdk_atom_name().) The default clipboard
// corresponds to the “CLIPBOARD” atom; another commonly used clipboard is the
// “PRIMARY” clipboard, which, in X, traditionally contains the currently
// selected text.
//
// To support having a number of different formats on the clipboard at the same
// time, the clipboard mechanism allows providing callbacks instead of the
// actual data. When you set the contents of the clipboard, you can either
// supply the data directly (via functions like gtk_clipboard_set_text()), or
// you can supply a callback to be called at a later time when the data is
// needed (via gtk_clipboard_set_with_data() or gtk_clipboard_set_with_owner().)
// Providing a callback also avoids having to make copies of the data when it is
// not needed.
//
// gtk_clipboard_set_with_data() and gtk_clipboard_set_with_owner() are quite
// similar; the choice between the two depends mostly on which is more
// convenient in a particular situation. The former is most useful when you want
// to have a blob of data with callbacks to convert it into the various data
// types that you advertise. When the clear_func you provided is called, you
// simply free the data blob. The latter is more useful when the contents of
// clipboard reflect the internal state of a #GObject (As an example, for the
// PRIMARY clipboard, when an entry widget provides the clipboard’s contents the
// contents are simply the text within the selected region.) If the contents
// change, the entry widget can call gtk_clipboard_set_with_owner() to update
// the timestamp for clipboard ownership, without having to worry about
// clear_func being called.
//
// Requesting the data from the clipboard is essentially asynchronous. If the
// contents of the clipboard are provided within the same process, then a direct
// function call will be made to retrieve the data, but if they are provided by
// another process, then the data needs to be retrieved from the other process,
// which may take some time. To avoid blocking the user interface, the call to
// request the selection, gtk_clipboard_request_contents() takes a callback that
// will be called when the contents are received (or when the request fails.) If
// you don’t want to deal with providing a separate callback, you can also use
// gtk_clipboard_wait_for_contents(). What this does is run the GLib main loop
// recursively waiting for the contents. This can simplify the code flow, but
// you still have to be aware that other callbacks in your program can be called
// while this recursive mainloop is running.
//
// Along with the functions to get the clipboard contents as an arbitrary data
// chunk, there are also functions to retrieve it as text,
// gtk_clipboard_request_text() and gtk_clipboard_wait_for_text(). These
// functions take care of determining which formats are advertised by the
// clipboard provider, asking for the clipboard in the best available format and
// converting the results into the UTF-8 encoding. (The standard form for
// representing strings in GTK+.)
type Clipboard struct {
	*externglib.Object
}

func wrapClipboard(obj *externglib.Object) *Clipboard {
	return &Clipboard{
		Object: obj,
	}
}

func marshalClipboarder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapClipboard(obj), nil
}

// Clear clears the contents of the clipboard. Generally this should only be
// called between the time you call gtk_clipboard_set_with_owner() or
// gtk_clipboard_set_with_data(), and when the clear_func you supplied is
// called. Otherwise, the clipboard may be owned by someone else.
func (clipboard *Clipboard) Clear() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	C.gtk_clipboard_clear(_arg0)
}

// Display gets the Display associated with clipboard
func (clipboard *Clipboard) Display() *gdk.Display {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_get_display(_arg0)

	var _display *gdk.Display // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Owner: if the clipboard contents callbacks were set with
// gtk_clipboard_set_with_owner(), and the gtk_clipboard_set_with_data() or
// gtk_clipboard_clear() has not subsequently called, returns the owner set by
// gtk_clipboard_set_with_owner().
func (clipboard *Clipboard) Owner() *externglib.Object {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_get_owner(_arg0)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// RequestImage requests the contents of the clipboard as image. When the image
// is later received, it will be converted to a Pixbuf, and callback will be
// called.
//
// The pixbuf parameter to callback will contain the resulting Pixbuf if the
// request succeeded, or NULL if it failed. This could happen for various
// reasons, in particular if the clipboard was empty or if the contents of the
// clipboard could not be converted into an image.
func (clipboard *Clipboard) RequestImage(callback ClipboardImageReceivedFunc) {
	var _arg0 *C.GtkClipboard                 // out
	var _arg1 C.GtkClipboardImageReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardImageReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_image(_arg0, _arg1, _arg2)
}

// RequestText requests the contents of the clipboard as text. When the text is
// later received, it will be converted to UTF-8 if necessary, and callback will
// be called.
//
// The text parameter to callback will contain the resulting text if the request
// succeeded, or NULL if it failed. This could happen for various reasons, in
// particular if the clipboard was empty or if the contents of the clipboard
// could not be converted into text form.
func (clipboard *Clipboard) RequestText(callback ClipboardTextReceivedFunc) {
	var _arg0 *C.GtkClipboard                // out
	var _arg1 C.GtkClipboardTextReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardTextReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_text(_arg0, _arg1, _arg2)
}

// RequestURIs requests the contents of the clipboard as URIs. When the URIs are
// later received callback will be called.
//
// The uris parameter to callback will contain the resulting array of URIs if
// the request succeeded, or NULL if it failed. This could happen for various
// reasons, in particular if the clipboard was empty or if the contents of the
// clipboard could not be converted into URI form.
func (clipboard *Clipboard) RequestURIs(callback ClipboardURIReceivedFunc) {
	var _arg0 *C.GtkClipboard               // out
	var _arg1 C.GtkClipboardURIReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardURIReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_uris(_arg0, _arg1, _arg2)
}

// SetCanStore hints that the clipboard data should be stored somewhere when the
// application exits or when gtk_clipboard_store () is called.
//
// This value is reset when the clipboard owner changes. Where the clipboard
// data is stored is platform dependent, see gdk_display_store_clipboard () for
// more information.
func (clipboard *Clipboard) SetCanStore(targets []TargetEntry) {
	var _arg0 *C.GtkClipboard   // out
	var _arg1 *C.GtkTargetEntry // out
	var _arg2 C.gint

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg2 = (C.gint)(len(targets))
	if len(targets) > 0 {
		_arg1 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))
	}

	C.gtk_clipboard_set_can_store(_arg0, _arg1, _arg2)
}

// SetImage sets the contents of the clipboard to the given Pixbuf. GTK+ will
// take responsibility for responding for requests for the image, and for
// converting the image into the requested format.
func (clipboard *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_clipboard_set_image(_arg0, _arg1)
}

// SetText sets the contents of the clipboard to the given UTF-8 string. GTK+
// will make a copy of the text and take responsibility for responding for
// requests for the text, and for converting the text into the requested format.
func (clipboard *Clipboard) SetText(text string, len int) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(len)

	C.gtk_clipboard_set_text(_arg0, _arg1, _arg2)
}

// Store stores the current clipboard data somewhere so that it will stay around
// after the application has quit.
func (clipboard *Clipboard) Store() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	C.gtk_clipboard_store(_arg0)
}

// WaitForImage requests the contents of the clipboard as image and converts the
// result to a Pixbuf. This function waits for the data to be received using the
// main loop, so events, timeouts, etc, may be dispatched during the wait.
func (clipboard *Clipboard) WaitForImage() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_for_image(_arg0)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// WaitForText requests the contents of the clipboard as text and converts the
// result to UTF-8 if necessary. This function waits for the data to be received
// using the main loop, so events, timeouts, etc, may be dispatched during the
// wait.
func (clipboard *Clipboard) WaitForText() string {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_for_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// WaitForURIs requests the contents of the clipboard as URIs. This function
// waits for the data to be received using the main loop, so events, timeouts,
// etc, may be dispatched during the wait.
func (clipboard *Clipboard) WaitForURIs() []string {
	var _arg0 *C.GtkClipboard // out
	var _cret **C.gchar       // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_for_uris(_arg0)

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
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// WaitIsImageAvailable: test to see if there is an image available to be pasted
// This is done by requesting the TARGETS atom and checking if it contains any
// of the supported image targets. This function waits for the data to be
// received using the main loop, so events, timeouts, etc, may be dispatched
// during the wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_image()
// since it doesn’t need to retrieve the actual image data.
func (clipboard *Clipboard) WaitIsImageAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_is_image_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsRichTextAvailable: test to see if there is rich text available to be
// pasted This is done by requesting the TARGETS atom and checking if it
// contains any of the supported rich text targets. This function waits for the
// data to be received using the main loop, so events, timeouts, etc, may be
// dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_rich_text() since it doesn’t need to retrieve the
// actual text.
func (clipboard *Clipboard) WaitIsRichTextAvailable(buffer *TextBuffer) bool {
	var _arg0 *C.GtkClipboard  // out
	var _arg1 *C.GtkTextBuffer // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))
	_arg1 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_clipboard_wait_is_rich_text_available(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsTextAvailable: test to see if there is text available to be pasted This
// is done by requesting the TARGETS atom and checking if it contains any of the
// supported text targets. This function waits for the data to be received using
// the main loop, so events, timeouts, etc, may be dispatched during the wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_text()
// since it doesn’t need to retrieve the actual text.
func (clipboard *Clipboard) WaitIsTextAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_is_text_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsURIsAvailable: test to see if there is a list of URIs available to be
// pasted This is done by requesting the TARGETS atom and checking if it
// contains the URI targets. This function waits for the data to be received
// using the main loop, so events, timeouts, etc, may be dispatched during the
// wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_uris()
// since it doesn’t need to retrieve the actual URI data.
func (clipboard *Clipboard) WaitIsURIsAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(clipboard.Native()))

	_cret = C.gtk_clipboard_wait_is_uris_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ClipboardGetDefault returns the default clipboard object for use with
// cut/copy/paste menu items and keyboard shortcuts.
func ClipboardGetDefault(display *gdk.Display) *Clipboard {
	var _arg1 *C.GdkDisplay   // out
	var _cret *C.GtkClipboard // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_clipboard_get_default(_arg1)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}
