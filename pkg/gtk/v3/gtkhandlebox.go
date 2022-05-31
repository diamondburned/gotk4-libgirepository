// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_HandleBoxClass_child_attached(void*, void*);
// extern void _gotk4_gtk3_HandleBoxClass_child_detached(void*, void*);
// extern void _gotk4_gtk3_HandleBox_ConnectChildAttached(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_HandleBox_ConnectChildDetached(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtkhandlebox.go.
var GTypeHandleBox = coreglib.Type(C.gtk_handle_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeHandleBox, F: marshalHandleBox},
	})
}

// HandleBoxOverrider contains methods that are overridable.
type HandleBoxOverrider interface {
	// The function takes the following parameters:
	//
	ChildAttached(child Widgetter)
	// The function takes the following parameters:
	//
	ChildDetached(child Widgetter)
}

// HandleBox widget allows a portion of a window to be "torn off". It is a bin
// widget which displays its child and a handle that the user can drag to tear
// off a separate window (the “float window”) containing the child widget. A
// thin “ghost” is drawn in the original location of the handlebox. By dragging
// the separate window back to its original location, it can be reattached.
//
// When reattaching, the ghost and float window, must be aligned along one of
// the edges, the “snap edge”. This either can be specified by the application
// programmer explicitly, or GTK+ will pick a reasonable default based on the
// handle position.
//
// To make detaching and reattaching the handlebox as minimally confusing as
// possible to the user, it is important to set the snap edge so that the snap
// edge does not move when the handlebox is deattached. For instance, if the
// handlebox is packed at the bottom of a VBox, then when the handlebox is
// detached, the bottom edge of the handlebox's allocation will remain fixed as
// the height of the handlebox shrinks, so the snap edge should be set to
// GTK_POS_BOTTOM.
//
// > HandleBox has been deprecated. It is very specialized, lacks features > to
// make it useful and most importantly does not fit well into modern >
// application design. Do not use it. There is no replacement.
type HandleBox struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*HandleBox)(nil)
)

func classInitHandleBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkHandleBoxClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkHandleBoxClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ChildAttached(child Widgetter) }); ok {
		pclass.child_attached = (*[0]byte)(C._gotk4_gtk3_HandleBoxClass_child_attached)
	}

	if _, ok := goval.(interface{ ChildDetached(child Widgetter) }); ok {
		pclass.child_detached = (*[0]byte)(C._gotk4_gtk3_HandleBoxClass_child_detached)
	}
}

//export _gotk4_gtk3_HandleBoxClass_child_attached
func _gotk4_gtk3_HandleBoxClass_child_attached(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChildAttached(child Widgetter) })

	var _child Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}

	iface.ChildAttached(_child)
}

//export _gotk4_gtk3_HandleBoxClass_child_detached
func _gotk4_gtk3_HandleBoxClass_child_detached(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChildDetached(child Widgetter) })

	var _child Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}

	iface.ChildDetached(_child)
}

func wrapHandleBox(obj *coreglib.Object) *HandleBox {
	return &HandleBox{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalHandleBox(p uintptr) (interface{}, error) {
	return wrapHandleBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_HandleBox_ConnectChildAttached
func _gotk4_gtk3_HandleBox_ConnectChildAttached(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

// ConnectChildAttached: this signal is emitted when the contents of the
// handlebox are reattached to the main window.
func (handleBox *HandleBox) ConnectChildAttached(f func(widget Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(handleBox, "child-attached", false, unsafe.Pointer(C._gotk4_gtk3_HandleBox_ConnectChildAttached), f)
}

//export _gotk4_gtk3_HandleBox_ConnectChildDetached
func _gotk4_gtk3_HandleBox_ConnectChildDetached(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

// ConnectChildDetached: this signal is emitted when the contents of the
// handlebox are detached from the main window.
func (handleBox *HandleBox) ConnectChildDetached(f func(widget Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(handleBox, "child-detached", false, unsafe.Pointer(C._gotk4_gtk3_HandleBox_ConnectChildDetached), f)
}

// NewHandleBox: create a new handle box.
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - handleBox: new HandleBox.
//
func NewHandleBox() *HandleBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "HandleBox").InvokeMethod("new_HandleBox", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _handleBox *HandleBox // out

	_handleBox = wrapHandleBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _handleBox
}

// ChildDetached: whether the handlebox’s child is currently detached.
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - ok: TRUE if the child is currently detached, otherwise FALSE.
//
func (handleBox *HandleBox) ChildDetached() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(handleBox).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "HandleBox").InvokeMethod("get_child_detached", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(handleBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
