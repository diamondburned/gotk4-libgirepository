// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_get_type()), F: marshalIMContexter},
	})
}

// IMContextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type IMContextOverrider interface {
	Commit(str string)
	// DeleteSurrounding asks the widget that the input context is attached to
	// to delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal. Note that offset and n_chars are
	// in characters not in bytes which differs from the usage other places in
	// IMContext.
	//
	// In order to use this function, you should first call
	// gtk_im_context_get_surrounding() to get the current context, and call
	// this function immediately afterwards to make sure that you know what you
	// are deleting. You should also account for the fact that even if the
	// signal was handled, the input context might not have deleted all the
	// characters that were requested to be deleted.
	//
	// This function is used by an input method that wants to make subsitutions
	// in the existing text in response to new input. It is not useful for
	// applications.
	DeleteSurrounding(offset int, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events. If this function returns TRUE, then no further processing
	// should be done for this key event.
	FilterKeypress(event *gdk.EventKey) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus. The input method may, for example,
	// change the displayed feedback to reflect this change.
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus. The input method may, for example,
	// change the displayed feedback or reset the contexts state to reflect this
	// change.
	FocusOut()
	// PreeditString: retrieve the current preedit string for the input context,
	// and a list of attributes to apply to the string. This string should be
	// displayed inserted at the insertion point.
	PreeditString() (string, *pango.AttrList, int)
	// Surrounding retrieves context around the insertion point. Input methods
	// typically want context in order to constrain input text based on existing
	// text; this is important for languages such as Thai where only some
	// sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// GtkIMContext::retrieve_surrounding signal on the input method; in
	// response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// gtk_im_context_set_surrounding(). Note that there is no obligation for a
	// widget to respond to the ::retrieve_surrounding signal, so input methods
	// must be prepared to function without context.
	Surrounding() (string, int, bool)
	PreeditChanged()
	PreeditEnd()
	PreeditStart()
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made. This will typically cause the input method to
	// clear the preedit state.
	Reset()
	RetrieveSurrounding() bool
	// SetClientWindow: set the client window for the input context; this is the
	// Window in which the input appears. This window is used in order to
	// correctly position status windows, and may also be used for purposes
	// internal to the input method.
	SetClientWindow(window gdk.Windower)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made. The location is relative to the client window.
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string. This function is expected to be called in response to the
	// GtkIMContext::retrieve_surrounding signal, and will likely have no effect
	// if called at other times.
	SetSurrounding(text string, len int, cursorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback. If use_preedit is FALSE (default is TRUE), then the
	// IM context may use some other method to display feedback, such as
	// displaying it in a child of the root window.
	SetUsePreedit(usePreedit bool)
}

// IMContext defines the interface for GTK+ input methods. An input method is
// used by GTK+ text input widgets like Entry to map from key events to Unicode
// character strings.
//
// The default input method can be set programmatically via the
// Settings:gtk-im-module GtkSettings property. Alternatively, you may set the
// GTK_IM_MODULE environment variable as documented in [Running GTK+
// Applications][gtk-running].
//
// The Entry Entry:im-module and TextView TextView:im-module properties may also
// be used to set input methods for specific widget instances. For instance, a
// certain entry widget might be expected to contain certain characters which
// would be easier to input with a certain input method.
//
// An input method may consume multiple key events in sequence and finally
// output the composed result. This is called preediting, and an input method
// may provide feedback about this process by displaying the intermediate
// composition states as preedit text. For instance, the default GTK+ input
// method implements the input of arbitrary Unicode code points by holding down
// the Control and Shift keys and then typing “U” followed by the hexadecimal
// digits of the code point. When releasing the Control and Shift keys,
// preediting ends and the character is inserted as text. Ctrl+Shift+u20AC for
// example results in the € sign.
//
// Additional input methods can be made available for use by GTK+ widgets as
// loadable modules. An input method module is a small shared library which
// implements a subclass of IMContext or IMContextSimple and exports these four
// functions:
//
//    GtkIMContext * im_module_create(const gchar *context_id);
//
// This function should return a pointer to a newly created instance of the
// IMContext subclass identified by context_id. The context ID is the same as
// specified in the IMContextInfo array returned by im_module_list().
//
// After a new loadable input method module has been installed on the system,
// the configuration file gtk.immodules needs to be regenerated by
// [gtk-query-immodules-3.0][gtk-query-immodules-3.0], in order for the new
// input method to become available to GTK+ applications.
type IMContext struct {
	*externglib.Object
}

// IMContexter describes IMContext's abstract methods.
type IMContexter interface {
	gextras.Objector

	// DeleteSurrounding asks the widget that the input context is attached to
	// to delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal.
	DeleteSurrounding(offset int, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events.
	FilterKeypress(event *gdk.EventKey) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus.
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus.
	FocusOut()
	// PreeditString: retrieve the current preedit string for the input context,
	// and a list of attributes to apply to the string.
	PreeditString() (string, *pango.AttrList, int)
	// Surrounding retrieves context around the insertion point.
	Surrounding() (string, int, bool)
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made.
	Reset()
	// SetClientWindow: set the client window for the input context; this is the
	// Window in which the input appears.
	SetClientWindow(window gdk.Windower)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made.
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string.
	SetSurrounding(text string, len int, cursorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback.
	SetUsePreedit(usePreedit bool)
}

var _ IMContexter = (*IMContext)(nil)

func wrapIMContext(obj *externglib.Object) *IMContext {
	return &IMContext{
		Object: obj,
	}
}

func marshalIMContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIMContext(obj), nil
}

// DeleteSurrounding asks the widget that the input context is attached to to
// delete characters around the cursor position by emitting the
// GtkIMContext::delete_surrounding signal. Note that offset and n_chars are in
// characters not in bytes which differs from the usage other places in
// IMContext.
//
// In order to use this function, you should first call
// gtk_im_context_get_surrounding() to get the current context, and call this
// function immediately afterwards to make sure that you know what you are
// deleting. You should also account for the fact that even if the signal was
// handled, the input context might not have deleted all the characters that
// were requested to be deleted.
//
// This function is used by an input method that wants to make subsitutions in
// the existing text in response to new input. It is not useful for
// applications.
func (context *IMContext) DeleteSurrounding(offset int, nChars int) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(offset)
	_arg2 = C.gint(nChars)

	_cret = C.gtk_im_context_delete_surrounding(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FilterKeypress: allow an input method to internally handle key press and
// release events. If this function returns TRUE, then no further processing
// should be done for this key event.
func (context *IMContext) FilterKeypress(event *gdk.EventKey) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkEventKey  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkEventKey)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gtk_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FocusIn: notify the input method that the widget to which this input context
// corresponds has gained focus. The input method may, for example, change the
// displayed feedback to reflect this change.
func (context *IMContext) FocusIn() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_focus_in(_arg0)
}

// FocusOut: notify the input method that the widget to which this input context
// corresponds has lost focus. The input method may, for example, change the
// displayed feedback or reset the contexts state to reflect this change.
func (context *IMContext) FocusOut() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_focus_out(_arg0)
}

// PreeditString: retrieve the current preedit string for the input context, and
// a list of attributes to apply to the string. This string should be displayed
// inserted at the insertion point.
func (context *IMContext) PreeditString() (string, *pango.AttrList, int) {
	var _arg0 *C.GtkIMContext  // out
	var _arg1 *C.gchar         // in
	var _arg2 *C.PangoAttrList // in
	var _arg3 C.gint           // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_get_preedit_string(_arg0, &_arg1, &_arg2, &_arg3)

	var _str string            // out
	var _attrs *pango.AttrList // out
	var _cursorPos int         // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_attrs = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
	C.pango_attr_list_ref(_arg2)
	runtime.SetFinalizer(_attrs, func(v *pango.AttrList) {
		C.pango_attr_list_unref((*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_cursorPos = int(_arg3)

	return _str, _attrs, _cursorPos
}

// Surrounding retrieves context around the insertion point. Input methods
// typically want context in order to constrain input text based on existing
// text; this is important for languages such as Thai where only some sequences
// of characters are allowed.
//
// This function is implemented by emitting the
// GtkIMContext::retrieve_surrounding signal on the input method; in response to
// this signal, a widget should provide as much context as is available, up to
// an entire paragraph, by calling gtk_im_context_set_surrounding(). Note that
// there is no obligation for a widget to respond to the ::retrieve_surrounding
// signal, so input methods must be prepared to function without context.
func (context *IMContext) Surrounding() (string, int, bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.gchar        // in
	var _arg2 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_context_get_surrounding(_arg0, &_arg1, &_arg2)

	var _text string     // out
	var _cursorIndex int // out
	var _ok bool         // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _text, _cursorIndex, _ok
}

// Reset: notify the input method that a change such as a change in cursor
// position has been made. This will typically cause the input method to clear
// the preedit state.
func (context *IMContext) Reset() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_reset(_arg0)
}

// SetClientWindow: set the client window for the input context; this is the
// Window in which the input appears. This window is used in order to correctly
// position status windows, and may also be used for purposes internal to the
// input method.
func (context *IMContext) SetClientWindow(window gdk.Windower) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkWindow    // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_im_context_set_client_window(_arg0, _arg1)
}

// SetCursorLocation: notify the input method that a change in cursor position
// has been made. The location is relative to the client window.
func (context *IMContext) SetCursorLocation(area *gdk.Rectangle) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(area)))

	C.gtk_im_context_set_cursor_location(_arg0, _arg1)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string. This function is expected to be called in response to the
// GtkIMContext::retrieve_surrounding signal, and will likely have no effect if
// called at other times.
func (context *IMContext) SetSurrounding(text string, len int, cursorIndex int) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(len)
	_arg3 = C.gint(cursorIndex)

	C.gtk_im_context_set_surrounding(_arg0, _arg1, _arg2, _arg3)
}

// SetUsePreedit sets whether the IM context should use the preedit string to
// display feedback. If use_preedit is FALSE (default is TRUE), then the IM
// context may use some other method to display feedback, such as displaying it
// in a child of the root window.
func (context *IMContext) SetUsePreedit(usePreedit bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	if usePreedit {
		_arg1 = C.TRUE
	}

	C.gtk_im_context_set_use_preedit(_arg0, _arg1)
}
