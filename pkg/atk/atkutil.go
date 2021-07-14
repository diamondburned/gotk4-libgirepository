// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_coord_type_get_type()), F: marshalCoordType},
		{T: externglib.Type(C.atk_key_event_type_get_type()), F: marshalKeyEventType},
		{T: externglib.Type(C.atk_util_get_type()), F: marshalUtiler},
	})
}

// CoordType specifies how xy coordinates are to be interpreted. Used by
// functions such as atk_component_get_position() and
// atk_text_get_character_extents()
type CoordType int

const (
	// Screen specifies xy coordinates relative to the screen
	CoordTypeScreen CoordType = iota
	// Window specifies xy coordinates relative to the widget's top-level window
	CoordTypeWindow
	// Parent specifies xy coordinates relative to the widget's immediate
	// parent. Since: 2.30
	CoordTypeParent
)

func marshalCoordType(p uintptr) (interface{}, error) {
	return CoordType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// KeyEventType specifies the type of a keyboard evemt.
type KeyEventType int

const (
	// Press specifies a key press event
	KeyEventTypePress KeyEventType = iota
	// Release specifies a key release event
	KeyEventTypeRelease
	// LastDefined: not a valid value; specifies end of enumeration
	KeyEventTypeLastDefined
)

func marshalKeyEventType(p uintptr) (interface{}, error) {
	return KeyEventType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// KeySnoopFunc is a type of callback which is called whenever a key event
// occurs, if registered via atk_add_key_event_listener. It allows for
// pre-emptive interception of key events via the return code as described
// below.
type KeySnoopFunc func(event *KeyEventStruct, userData cgo.Handle) (gint int)

//export _gotk4_atk1_KeySnoopFunc
func _gotk4_atk1_KeySnoopFunc(arg0 *C.AtkKeyEventStruct, arg1 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var event *KeyEventStruct // out
	var userData cgo.Handle   // out

	event = (*KeyEventStruct)(unsafe.Pointer(arg0))
	runtime.SetFinalizer(event, func(v *KeyEventStruct) {
		C.free(unsafe.Pointer(v))
	})
	userData = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(KeySnoopFunc)
	gint := fn(event, userData)

	cret = C.gint(gint)

	return cret
}

// FocusTrackerNotify: cause the focus tracker functions which have been
// specified to be executed for the object.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. As Object::focus-event was deprecated in favor of a
// Object::state-change signal, in order to notify a focus change on your
// implementation, you can use atk_object_notify_state_change() instead.
func FocusTrackerNotify(object *ObjectClass) {
	var _arg1 *C.AtkObject // out

	_arg1 = (*C.AtkObject)(unsafe.Pointer(object.Native()))

	C.atk_focus_tracker_notify(_arg1)
}

// GetFocusObject gets the currently focused object.
func GetFocusObject() *ObjectClass {
	var _cret *C.AtkObject // in

	_cret = C.atk_get_focus_object()

	var _object *ObjectClass // out

	_object = wrapObject(externglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// GetRoot gets the root accessible container for the current application.
func GetRoot() *ObjectClass {
	var _cret *C.AtkObject // in

	_cret = C.atk_get_root()

	var _object *ObjectClass // out

	_object = wrapObject(externglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// GetToolkitName gets name string for the GUI toolkit implementing ATK for this
// application.
func GetToolkitName() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_toolkit_name()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetToolkitVersion gets version string for the GUI toolkit implementing ATK
// for this application.
func GetToolkitVersion() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_toolkit_version()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetVersion gets the current version for ATK.
func GetVersion() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_version()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RemoveFocusTracker removes the specified focus tracker from the list of
// functions to be called when any object receives focus.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. If you need focus tracking on your implementation, subscribe to
// the Object::state-change "focused" signal.
func RemoveFocusTracker(trackerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(trackerId)

	C.atk_remove_focus_tracker(_arg1)
}

// RemoveGlobalEventListener: listener_id is the value returned by
// #atk_add_global_event_listener when you registered that event listener.
//
// Toolkit implementor note: ATK provides a default implementation for this
// virtual method. ATK implementors are discouraged from reimplementing this
// method.
//
// Toolkit implementor note: this method is not intended to be used by ATK
// implementors but by ATK consumers.
//
// Removes the specified event listener
func RemoveGlobalEventListener(listenerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(listenerId)

	C.atk_remove_global_event_listener(_arg1)
}

// RemoveKeyEventListener: listener_id is the value returned by
// #atk_add_key_event_listener when you registered that event listener.
//
// Removes the specified event listener.
func RemoveKeyEventListener(listenerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(listenerId)

	C.atk_remove_key_event_listener(_arg1)
}

// Utiler describes Util's methods.
type Utiler interface {
	privateUtil()
}

// Util: set of ATK utility functions which are used to support event
// registration of various types, and obtaining the 'root' accessible of a
// process and information about the current ATK implementation and toolkit
// version.
type Util struct {
	*externglib.Object
}

var (
	_ Utiler          = (*Util)(nil)
	_ gextras.Nativer = (*Util)(nil)
)

func wrapUtil(obj *externglib.Object) *Util {
	return &Util{
		Object: obj,
	}
}

func marshalUtiler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUtil(obj), nil
}

func (*Util) privateUtil() {}

// KeyEventStruct encapsulates information about a key event.
type KeyEventStruct struct {
	native C.AtkKeyEventStruct
}

// Native returns the underlying C source pointer.
func (k *KeyEventStruct) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Type: atkKeyEventType, generally one of ATK_KEY_EVENT_PRESS or
// ATK_KEY_EVENT_RELEASE
func (k *KeyEventStruct) Type() int {
	var v int // out
	v = int(k.native._type)
	return v
}

// State: bitmask representing the state of the modifier keys immediately after
// the event takes place. The meaning of the bits is currently defined to match
// the bitmask used by GDK in GdkEventType.state, see
// http://developer.gnome.org/doc/API/2.0/gdk/gdk-Event-Structures.htmlEventKey
func (k *KeyEventStruct) State() uint {
	var v uint // out
	v = uint(k.native.state)
	return v
}

// Keyval: guint representing a keysym value corresponding to those used by GDK
// and X11: see /usr/X11/include/keysymdef.h.
func (k *KeyEventStruct) Keyval() uint {
	var v uint // out
	v = uint(k.native.keyval)
	return v
}

// Length: length of member #string.
func (k *KeyEventStruct) Length() int {
	var v int // out
	v = int(k.native.length)
	return v
}

// String: string containing one of the following: either a string approximating
// the text that would result from this keypress, if the key is a control or
// graphic character, or a symbolic name for this keypress. Alphanumeric and
// printable keys will have the symbolic key name in this string member, for
// instance "A". "0", "semicolon", "aacute". Keypad keys have the prefix "KP".
func (k *KeyEventStruct) String() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(k.native.string)))
	return v
}

// Keycode: raw hardware code that generated the key event. This field is raraly
// useful.
func (k *KeyEventStruct) Keycode() uint16 {
	var v uint16 // out
	v = uint16(k.native.keycode)
	return v
}

// Timestamp: timestamp in milliseconds indicating when the event occurred.
// These timestamps are relative to a starting point which should be considered
// arbitrary, and only used to compare the dispatch times of events to one
// another.
func (k *KeyEventStruct) Timestamp() uint32 {
	var v uint32 // out
	v = uint32(k.native.timestamp)
	return v
}
