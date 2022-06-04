// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_IMContextClass_delete_surrounding(void*, gint, gint);
// extern gboolean _gotk4_gtk3_IMContextClass_filter_keypress(void*, void*);
// extern gboolean _gotk4_gtk3_IMContextClass_get_surrounding(void*, void**, void*);
// extern gboolean _gotk4_gtk3_IMContextClass_retrieve_surrounding(void*);
// extern gboolean _gotk4_gtk3_IMContext_ConnectDeleteSurrounding(gpointer, gint, gint, guintptr);
// extern gboolean _gotk4_gtk3_IMContext_ConnectRetrieveSurrounding(gpointer, guintptr);
// extern void _gotk4_gtk3_IMContextClass_commit(void*, void*);
// extern void _gotk4_gtk3_IMContextClass_focus_in(void*);
// extern void _gotk4_gtk3_IMContextClass_focus_out(void*);
// extern void _gotk4_gtk3_IMContextClass_get_preedit_string(void*, void**, void**, void*);
// extern void _gotk4_gtk3_IMContextClass_preedit_changed(void*);
// extern void _gotk4_gtk3_IMContextClass_preedit_end(void*);
// extern void _gotk4_gtk3_IMContextClass_preedit_start(void*);
// extern void _gotk4_gtk3_IMContextClass_reset(void*);
// extern void _gotk4_gtk3_IMContextClass_set_client_window(void*, void*);
// extern void _gotk4_gtk3_IMContextClass_set_cursor_location(void*, void*);
// extern void _gotk4_gtk3_IMContextClass_set_surrounding(void*, void*, gint, gint);
// extern void _gotk4_gtk3_IMContextClass_set_use_preedit(void*, gboolean);
// extern void _gotk4_gtk3_IMContext_ConnectCommit(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_IMContext_ConnectPreeditChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_IMContext_ConnectPreeditEnd(gpointer, guintptr);
// extern void _gotk4_gtk3_IMContext_ConnectPreeditStart(gpointer, guintptr);
import "C"

// glib.Type values for gtkimcontext.go.
var GTypeIMContext = coreglib.Type(C.gtk_im_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeIMContext, F: marshalIMContext},
	})
}

// IMContextOverrider contains methods that are overridable.
type IMContextOverrider interface {
	// The function takes the following parameters:
	//
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
	//
	// The function takes the following parameters:
	//
	//    - offset from cursor position in chars; a negative value means start
	//      before the cursor.
	//    - nChars: number of characters to delete.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the signal was handled.
	//
	DeleteSurrounding(offset, nChars int32) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events. If this function returns TRUE, then no further processing
	// should be done for this key event.
	//
	// The function takes the following parameters:
	//
	//    - event: key event.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the input method handled the key event.
	//
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
	//
	// The function returns the following values:
	//
	//    - str: location to store the retrieved string. The string retrieved
	//      must be freed with g_free().
	//    - attrs: location to store the retrieved attribute list. When you are
	//      done with this list, you must unreference it with
	//      pango_attr_list_unref().
	//    - cursorPos: location to store position of cursor (in characters)
	//      within the preedit string.
	//
	PreeditString() (string, *pango.AttrList, int32)
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
	//
	// The function returns the following values:
	//
	//    - text: location to store a UTF-8 encoded string of text holding
	//      context around the insertion point. If the function returns TRUE,
	//      then you must free the result stored in this location with g_free().
	//    - cursorIndex: location to store byte index of the insertion cursor
	//      within text.
	//    - ok: TRUE if surrounding text was provided; in this case you must free
	//      the result stored in *text.
	//
	Surrounding() (string, int32, bool)
	PreeditChanged()
	PreeditEnd()
	PreeditStart()
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made. This will typically cause the input method to
	// clear the preedit state.
	Reset()
	// The function returns the following values:
	//
	RetrieveSurrounding() bool
	// SetClientWindow: set the client window for the input context; this is the
	// Window in which the input appears. This window is used in order to
	// correctly position status windows, and may also be used for purposes
	// internal to the input method.
	//
	// The function takes the following parameters:
	//
	//    - window (optional): client window. This may be NULL to indicate that
	//      the previous client window no longer exists.
	//
	SetClientWindow(window gdk.Windower)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made. The location is relative to the client window.
	//
	// The function takes the following parameters:
	//
	//    - area: new location.
	//
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string. This function is expected to be called in response to the
	// GtkIMContext::retrieve_surrounding signal, and will likely have no effect
	// if called at other times.
	//
	// The function takes the following parameters:
	//
	//    - text surrounding the insertion point, as UTF-8. the preedit string
	//      should not be included within text.
	//    - len: length of text, or -1 if text is nul-terminated.
	//    - cursorIndex: byte index of the insertion cursor within text.
	//
	SetSurrounding(text string, len, cursorIndex int32)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback. If use_preedit is FALSE (default is TRUE), then the
	// IM context may use some other method to display feedback, such as
	// displaying it in a child of the root window.
	//
	// The function takes the following parameters:
	//
	//    - usePreedit: whether the IM context should use the preedit string.
	//
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*IMContext)(nil)
)

// IMContexter describes types inherited from class IMContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type IMContexter interface {
	coreglib.Objector
	baseIMContext() *IMContext
}

var _ IMContexter = (*IMContext)(nil)

func classInitIMContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkIMContextClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkIMContextClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Commit(str string) }); ok {
		pclass.commit = (*[0]byte)(C._gotk4_gtk3_IMContextClass_commit)
	}

	if _, ok := goval.(interface {
		DeleteSurrounding(offset, nChars int32) bool
	}); ok {
		pclass.delete_surrounding = (*[0]byte)(C._gotk4_gtk3_IMContextClass_delete_surrounding)
	}

	if _, ok := goval.(interface {
		FilterKeypress(event *gdk.EventKey) bool
	}); ok {
		pclass.filter_keypress = (*[0]byte)(C._gotk4_gtk3_IMContextClass_filter_keypress)
	}

	if _, ok := goval.(interface{ FocusIn() }); ok {
		pclass.focus_in = (*[0]byte)(C._gotk4_gtk3_IMContextClass_focus_in)
	}

	if _, ok := goval.(interface{ FocusOut() }); ok {
		pclass.focus_out = (*[0]byte)(C._gotk4_gtk3_IMContextClass_focus_out)
	}

	if _, ok := goval.(interface {
		PreeditString() (string, *pango.AttrList, int32)
	}); ok {
		pclass.get_preedit_string = (*[0]byte)(C._gotk4_gtk3_IMContextClass_get_preedit_string)
	}

	if _, ok := goval.(interface{ Surrounding() (string, int32, bool) }); ok {
		pclass.get_surrounding = (*[0]byte)(C._gotk4_gtk3_IMContextClass_get_surrounding)
	}

	if _, ok := goval.(interface{ PreeditChanged() }); ok {
		pclass.preedit_changed = (*[0]byte)(C._gotk4_gtk3_IMContextClass_preedit_changed)
	}

	if _, ok := goval.(interface{ PreeditEnd() }); ok {
		pclass.preedit_end = (*[0]byte)(C._gotk4_gtk3_IMContextClass_preedit_end)
	}

	if _, ok := goval.(interface{ PreeditStart() }); ok {
		pclass.preedit_start = (*[0]byte)(C._gotk4_gtk3_IMContextClass_preedit_start)
	}

	if _, ok := goval.(interface{ Reset() }); ok {
		pclass.reset = (*[0]byte)(C._gotk4_gtk3_IMContextClass_reset)
	}

	if _, ok := goval.(interface{ RetrieveSurrounding() bool }); ok {
		pclass.retrieve_surrounding = (*[0]byte)(C._gotk4_gtk3_IMContextClass_retrieve_surrounding)
	}

	if _, ok := goval.(interface{ SetClientWindow(window gdk.Windower) }); ok {
		pclass.set_client_window = (*[0]byte)(C._gotk4_gtk3_IMContextClass_set_client_window)
	}

	if _, ok := goval.(interface{ SetCursorLocation(area *gdk.Rectangle) }); ok {
		pclass.set_cursor_location = (*[0]byte)(C._gotk4_gtk3_IMContextClass_set_cursor_location)
	}

	if _, ok := goval.(interface {
		SetSurrounding(text string, len, cursorIndex int32)
	}); ok {
		pclass.set_surrounding = (*[0]byte)(C._gotk4_gtk3_IMContextClass_set_surrounding)
	}

	if _, ok := goval.(interface{ SetUsePreedit(usePreedit bool) }); ok {
		pclass.set_use_preedit = (*[0]byte)(C._gotk4_gtk3_IMContextClass_set_use_preedit)
	}
}

//export _gotk4_gtk3_IMContextClass_commit
func _gotk4_gtk3_IMContextClass_commit(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Commit(str string) })

	var _str string // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.Commit(_str)
}

//export _gotk4_gtk3_IMContextClass_delete_surrounding
func _gotk4_gtk3_IMContextClass_delete_surrounding(arg0 *C.void, arg1 C.gint, arg2 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DeleteSurrounding(offset, nChars int32) bool
	})

	var _offset int32 // out
	var _nChars int32 // out

	_offset = int32(arg1)
	_nChars = int32(arg2)

	ok := iface.DeleteSurrounding(_offset, _nChars)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_IMContextClass_filter_keypress
func _gotk4_gtk3_IMContextClass_filter_keypress(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		FilterKeypress(event *gdk.EventKey) bool
	})

	var _event *gdk.EventKey // out

	_event = (*gdk.EventKey)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.FilterKeypress(_event)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_IMContextClass_focus_in
func _gotk4_gtk3_IMContextClass_focus_in(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FocusIn() })

	iface.FocusIn()
}

//export _gotk4_gtk3_IMContextClass_focus_out
func _gotk4_gtk3_IMContextClass_focus_out(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FocusOut() })

	iface.FocusOut()
}

//export _gotk4_gtk3_IMContextClass_get_preedit_string
func _gotk4_gtk3_IMContextClass_get_preedit_string(arg0 *C.void, arg1 **C.void, arg2 **C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		PreeditString() (string, *pango.AttrList, int32)
	})

	str, attrs, cursorPos := iface.PreeditString()

	*arg1 = (*C.void)(unsafe.Pointer(C.CString(str)))
	*arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(attrs)))
	*arg3 = (*C.void)(unsafe.Pointer(cursorPos))
}

//export _gotk4_gtk3_IMContextClass_get_surrounding
func _gotk4_gtk3_IMContextClass_get_surrounding(arg0 *C.void, arg1 **C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Surrounding() (string, int32, bool) })

	text, cursorIndex, ok := iface.Surrounding()

	*arg1 = (*C.void)(unsafe.Pointer(C.CString(text)))
	*arg2 = (*C.void)(unsafe.Pointer(cursorIndex))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_IMContextClass_preedit_changed
func _gotk4_gtk3_IMContextClass_preedit_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PreeditChanged() })

	iface.PreeditChanged()
}

//export _gotk4_gtk3_IMContextClass_preedit_end
func _gotk4_gtk3_IMContextClass_preedit_end(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PreeditEnd() })

	iface.PreeditEnd()
}

//export _gotk4_gtk3_IMContextClass_preedit_start
func _gotk4_gtk3_IMContextClass_preedit_start(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PreeditStart() })

	iface.PreeditStart()
}

//export _gotk4_gtk3_IMContextClass_reset
func _gotk4_gtk3_IMContextClass_reset(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Reset() })

	iface.Reset()
}

//export _gotk4_gtk3_IMContextClass_retrieve_surrounding
func _gotk4_gtk3_IMContextClass_retrieve_surrounding(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ RetrieveSurrounding() bool })

	ok := iface.RetrieveSurrounding()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_IMContextClass_set_client_window
func _gotk4_gtk3_IMContextClass_set_client_window(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SetClientWindow(window gdk.Windower) })

	var _window gdk.Windower // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Windower)
				return ok
			})
			rv, ok := casted.(gdk.Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	iface.SetClientWindow(_window)
}

//export _gotk4_gtk3_IMContextClass_set_cursor_location
func _gotk4_gtk3_IMContextClass_set_cursor_location(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SetCursorLocation(area *gdk.Rectangle) })

	var _area *gdk.Rectangle // out

	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.SetCursorLocation(_area)
}

//export _gotk4_gtk3_IMContextClass_set_surrounding
func _gotk4_gtk3_IMContextClass_set_surrounding(arg0 *C.void, arg1 *C.void, arg2 C.gint, arg3 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		SetSurrounding(text string, len, cursorIndex int32)
	})

	var _text string       // out
	var _len int32         // out
	var _cursorIndex int32 // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_len = int32(arg2)
	_cursorIndex = int32(arg3)

	iface.SetSurrounding(_text, _len, _cursorIndex)
}

//export _gotk4_gtk3_IMContextClass_set_use_preedit
func _gotk4_gtk3_IMContextClass_set_use_preedit(arg0 *C.void, arg1 C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SetUsePreedit(usePreedit bool) })

	var _usePreedit bool // out

	if arg1 != 0 {
		_usePreedit = true
	}

	iface.SetUsePreedit(_usePreedit)
}

func wrapIMContext(obj *coreglib.Object) *IMContext {
	return &IMContext{
		Object: obj,
	}
}

func marshalIMContext(p uintptr) (interface{}, error) {
	return wrapIMContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (context *IMContext) baseIMContext() *IMContext {
	return context
}

// BaseIMContext returns the underlying base object.
func BaseIMContext(obj IMContexter) *IMContext {
	return obj.baseIMContext()
}

//export _gotk4_gtk3_IMContext_ConnectCommit
func _gotk4_gtk3_IMContext_ConnectCommit(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(str string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(str string))
	}

	var _str string // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_str)
}

// ConnectCommit signal is emitted when a complete input sequence has been
// entered by the user. This can be a single character immediately after a key
// press or the final result of preediting.
func (context *IMContext) ConnectCommit(f func(str string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "commit", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectCommit), f)
}

//export _gotk4_gtk3_IMContext_ConnectDeleteSurrounding
func _gotk4_gtk3_IMContext_ConnectDeleteSurrounding(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) (cret C.gboolean) {
	var f func(offset, nChars int32) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offset, nChars int32) (ok bool))
	}

	var _offset int32 // out
	var _nChars int32 // out

	_offset = int32(arg1)
	_nChars = int32(arg2)

	ok := f(_offset, _nChars)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectDeleteSurrounding signal is emitted when the input method needs to
// delete all or part of the context surrounding the cursor.
func (context *IMContext) ConnectDeleteSurrounding(f func(offset, nChars int32) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "delete-surrounding", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectDeleteSurrounding), f)
}

//export _gotk4_gtk3_IMContext_ConnectPreeditChanged
func _gotk4_gtk3_IMContext_ConnectPreeditChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPreeditChanged signal is emitted whenever the preedit sequence
// currently being entered has changed. It is also emitted at the end of a
// preedit sequence, in which case gtk_im_context_get_preedit_string() returns
// the empty string.
func (context *IMContext) ConnectPreeditChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "preedit-changed", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectPreeditChanged), f)
}

//export _gotk4_gtk3_IMContext_ConnectPreeditEnd
func _gotk4_gtk3_IMContext_ConnectPreeditEnd(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPreeditEnd signal is emitted when a preediting sequence has been
// completed or canceled.
func (context *IMContext) ConnectPreeditEnd(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "preedit-end", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectPreeditEnd), f)
}

//export _gotk4_gtk3_IMContext_ConnectPreeditStart
func _gotk4_gtk3_IMContext_ConnectPreeditStart(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPreeditStart signal is emitted when a new preediting sequence starts.
func (context *IMContext) ConnectPreeditStart(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "preedit-start", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectPreeditStart), f)
}

//export _gotk4_gtk3_IMContext_ConnectRetrieveSurrounding
func _gotk4_gtk3_IMContext_ConnectRetrieveSurrounding(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectRetrieveSurrounding signal is emitted when the input method requires
// the context surrounding the cursor. The callback should set the input method
// surrounding context by calling the gtk_im_context_set_surrounding() method.
func (context *IMContext) ConnectRetrieveSurrounding(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(context, "retrieve-surrounding", false, unsafe.Pointer(C._gotk4_gtk3_IMContext_ConnectRetrieveSurrounding), f)
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
//
// The function takes the following parameters:
//
//    - offset from cursor position in chars; a negative value means start before
//      the cursor.
//    - nChars: number of characters to delete.
//
// The function returns the following values:
//
//    - ok: TRUE if the signal was handled.
//
func (context *IMContext) DeleteSurrounding(offset, nChars int32) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(offset)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(nChars)

	_gret := girepository.MustFind("Gtk", "IMContext").InvokeMethod("delete_surrounding", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(nChars)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// FilterKeypress: allow an input method to internally handle key press and
// release events. If this function returns TRUE, then no further processing
// should be done for this key event.
//
// The function takes the following parameters:
//
//    - event: key event.
//
// The function returns the following values:
//
//    - ok: TRUE if the input method handled the key event.
//
func (context *IMContext) FilterKeypress(event *gdk.EventKey) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(event)))

	_gret := girepository.MustFind("Gtk", "IMContext").InvokeMethod("filter_keypress", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)
	runtime.KeepAlive(event)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// FocusIn: notify the input method that the widget to which this input context
// corresponds has gained focus. The input method may, for example, change the
// displayed feedback to reflect this change.
func (context *IMContext) FocusIn() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("focus_in", _args[:], nil)

	runtime.KeepAlive(context)
}

// FocusOut: notify the input method that the widget to which this input context
// corresponds has lost focus. The input method may, for example, change the
// displayed feedback or reset the contexts state to reflect this change.
func (context *IMContext) FocusOut() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("focus_out", _args[:], nil)

	runtime.KeepAlive(context)
}

// PreeditString: retrieve the current preedit string for the input context, and
// a list of attributes to apply to the string. This string should be displayed
// inserted at the insertion point.
//
// The function returns the following values:
//
//    - str: location to store the retrieved string. The string retrieved must be
//      freed with g_free().
//    - attrs: location to store the retrieved attribute list. When you are done
//      with this list, you must unreference it with pango_attr_list_unref().
//    - cursorPos: location to store position of cursor (in characters) within
//      the preedit string.
//
func (context *IMContext) PreeditString() (string, *pango.AttrList, int32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("get_preedit_string", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _str string            // out
	var _attrs *pango.AttrList // out
	var _cursorPos int32       // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(_outs[0])))
	defer C.free(unsafe.Pointer(_outs[0]))
	_attrs = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_outs[1])))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attrs)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
		},
	)
	_cursorPos = *(*int32)(unsafe.Pointer(_outs[2]))

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
//
// The function returns the following values:
//
//    - text: location to store a UTF-8 encoded string of text holding context
//      around the insertion point. If the function returns TRUE, then you must
//      free the result stored in this location with g_free().
//    - cursorIndex: location to store byte index of the insertion cursor within
//      text.
//    - ok: TRUE if surrounding text was provided; in this case you must free the
//      result stored in *text.
//
func (context *IMContext) Surrounding() (string, int32, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gtk", "IMContext").InvokeMethod("get_surrounding", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _text string       // out
	var _cursorIndex int32 // out
	var _ok bool           // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(_outs[0])))
	defer C.free(unsafe.Pointer(_outs[0]))
	_cursorIndex = *(*int32)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _text, _cursorIndex, _ok
}

// Reset: notify the input method that a change such as a change in cursor
// position has been made. This will typically cause the input method to clear
// the preedit state.
func (context *IMContext) Reset() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("reset", _args[:], nil)

	runtime.KeepAlive(context)
}

// SetClientWindow: set the client window for the input context; this is the
// Window in which the input appears. This window is used in order to correctly
// position status windows, and may also be used for purposes internal to the
// input method.
//
// The function takes the following parameters:
//
//    - window (optional): client window. This may be NULL to indicate that the
//      previous client window no longer exists.
//
func (context *IMContext) SetClientWindow(window gdk.Windower) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if window != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	}

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("set_client_window", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(window)
}

// SetCursorLocation: notify the input method that a change in cursor position
// has been made. The location is relative to the client window.
//
// The function takes the following parameters:
//
//    - area: new location.
//
func (context *IMContext) SetCursorLocation(area *gdk.Rectangle) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(area)))

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("set_cursor_location", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(area)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string. This function is expected to be called in response to the
// GtkIMContext::retrieve_surrounding signal, and will likely have no effect if
// called at other times.
//
// The function takes the following parameters:
//
//    - text surrounding the insertion point, as UTF-8. the preedit string should
//      not be included within text.
//    - len: length of text, or -1 if text is nul-terminated.
//    - cursorIndex: byte index of the insertion cursor within text.
//
func (context *IMContext) SetSurrounding(text string, len, cursorIndex int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(len)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(cursorIndex)

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("set_surrounding", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(len)
	runtime.KeepAlive(cursorIndex)
}

// SetUsePreedit sets whether the IM context should use the preedit string to
// display feedback. If use_preedit is FALSE (default is TRUE), then the IM
// context may use some other method to display feedback, such as displaying it
// in a child of the root window.
//
// The function takes the following parameters:
//
//    - usePreedit: whether the IM context should use the preedit string.
//
func (context *IMContext) SetUsePreedit(usePreedit bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if usePreedit {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "IMContext").InvokeMethod("set_use_preedit", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(usePreedit)
}
