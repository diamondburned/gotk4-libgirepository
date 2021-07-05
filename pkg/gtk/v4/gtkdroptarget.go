// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_target_get_type()), F: marshalDropTarget},
	})
}

// DropTarget: `GtkDropTarget` is an event controller to receive Drag-and-Drop
// operations.
//
// The most basic way to use a `GtkDropTarget` to receive drops on a widget is
// to create it via [ctor@Gtk.DropTarget.new], passing in the `GType` of the
// data you want to receive and connect to the [signal@Gtk.DropTarget::drop]
// signal to receive the data:
//
// “`c static gboolean on_drop (GtkDropTarget *target, const GValue *value,
// double x, double y, gpointer data) { MyWidget *self = data;
//
//    // Call the appropriate setter depending on the type of data
//    // that we received
//    if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//      my_widget_set_file (self, g_value_get_object (value));
//    else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//      my_widget_set_pixbuf (self, g_value_get_object (value));
//    else
//      return FALSE;
//
//    return TRUE;
//
// }
//
// static void my_widget_init (MyWidget *self) { GtkDropTarget *target =
// gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//    // This widget accepts two types of drop types: GFile objects
//    // and GdkPixbuf objects
//    gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//      G_TYPE_FILE,
//      GDK_TYPE_PIXBUF,
//    }, 2);
//
//    gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//
// } “`
//
// `GtkDropTarget` supports more options, such as:
//
//    * rejecting potential drops via the [signal@Gtk.DropTarget::accept] signal
//      and the [method@Gtk.DropTarget.reject] function to let other drop
//      targets handle the drop
//    * tracking an ongoing drag operation before the drop via the
//      [signal@Gtk.DropTarget::enter], [signal@Gtk.DropTarget::motion] and
//      [signal@Gtk.DropTarget::leave] signals
//    * configuring how to receive data by setting the
//      [property@Gtk.DropTarget:preload] property and listening for its
//      availability via the [property@Gtk.DropTarget:value] property
//
// However, `GtkDropTarget` is ultimately modeled in a synchronous way and only
// supports data transferred via `GType`. If you want full control over an
// ongoing drop, the [class@Gtk.DropTargetAsync] object gives you this ability.
//
// While a pointer is dragged over the drop target's widget and the drop has not
// been rejected, that widget will receive the GTK_STATE_FLAG_DROP_ACTIVE state,
// which can be used to style the widget.
//
// If you are not interested in receiving the drop, but just want to update UI
// state during a Drag-and-Drop operation (e.g. switching tabs), you can use
// [class@Gtk.DropControllerMotion].
type DropTarget interface {
	EventController

	// Actions gets the actions that this drop target supports.
	Actions() gdk.DragAction
	// Drop gets the currently handled drop operation.
	//
	// If no drop operation is going on, nil is returned.
	Drop() gdk.Drop
	// Formats gets the data formats that this drop target accepts.
	//
	// If the result is nil, all formats are expected to be supported.
	Formats() *gdk.ContentFormats
	// Preload gets whether data should be preloaded on hover.
	Preload() bool
	// Value gets the current drop data, as a `GValue`.
	Value() externglib.Value
	// RejectDropTarget rejects the ongoing drop operation.
	//
	// If no drop operation is ongoing, i.e when [property@Gtk.DropTarget:drop]
	// is nil, this function does nothing.
	//
	// This function should be used when delaying the decision on whether to
	// accept a drag or not until after reading the data.
	RejectDropTarget()
	// SetActionsDropTarget sets the actions that this drop target supports.
	SetActionsDropTarget(actions gdk.DragAction)
	// SetGTypesDropTarget sets the supported `GTypes` for this drop target.
	SetGTypesDropTarget(types []externglib.Type)
	// SetPreloadDropTarget sets whether data should be preloaded on hover.
	SetPreloadDropTarget(preload bool)
}

// dropTarget implements the DropTarget class.
type dropTarget struct {
	EventController
}

// WrapDropTarget wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropTarget(obj *externglib.Object) DropTarget {
	return dropTarget{
		EventController: WrapEventController(obj),
	}
}

func marshalDropTarget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropTarget(obj), nil
}

// NewDropTarget creates a new `GtkDropTarget` object.
//
// If the drop target should support more than 1 type, pass G_TYPE_INVALID for
// @type and then call [method@Gtk.DropTarget.set_gtypes].
func NewDropTarget(typ externglib.Type, actions gdk.DragAction) DropTarget {
	var _arg1 C.GType          // out
	var _arg2 C.GdkDragAction  // out
	var _cret *C.GtkDropTarget // in

	_arg1 = (C.GType)(typ)
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gtk_drop_target_new(_arg1, _arg2)

	var _dropTarget DropTarget // out

	_dropTarget = WrapDropTarget(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dropTarget
}

func (s dropTarget) Actions() gdk.DragAction {
	var _arg0 *C.GtkDropTarget // out
	var _cret C.GdkDragAction  // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_get_actions(_arg0)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

func (s dropTarget) Drop() gdk.Drop {
	var _arg0 *C.GtkDropTarget // out
	var _cret *C.GdkDrop       // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_get_drop(_arg0)

	var _drop gdk.Drop // out

	_drop = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Drop)

	return _drop
}

func (s dropTarget) Formats() *gdk.ContentFormats {
	var _arg0 *C.GtkDropTarget     // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_get_formats(_arg0)

	var _contentFormats *gdk.ContentFormats // out

	_contentFormats = (*gdk.ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *gdk.ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(unsafe.Pointer(v)))
	})

	return _contentFormats
}

func (s dropTarget) Preload() bool {
	var _arg0 *C.GtkDropTarget // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_get_preload(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s dropTarget) Value() externglib.Value {
	var _arg0 *C.GtkDropTarget // out
	var _cret *C.GValue        // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_get_value(_arg0)

	var _value externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

func (s dropTarget) RejectDropTarget() {
	var _arg0 *C.GtkDropTarget // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	C.gtk_drop_target_reject(_arg0)
}

func (s dropTarget) SetActionsDropTarget(actions gdk.DragAction) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 C.GdkDragAction  // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drop_target_set_actions(_arg0, _arg1)
}

func (s dropTarget) SetGTypesDropTarget(types []externglib.Type) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 *C.GType
	var _arg2 C.gsize

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	_arg2 = C.gsize(len(types))
	_arg1 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	C.gtk_drop_target_set_gtypes(_arg0, _arg1, _arg2)
}

func (s dropTarget) SetPreloadDropTarget(preload bool) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	if preload {
		_arg1 = C.TRUE
	}

	C.gtk_drop_target_set_preload(_arg0, _arg1)
}
