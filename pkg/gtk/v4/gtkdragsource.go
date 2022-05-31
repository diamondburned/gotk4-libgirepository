// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GdkContentProvider* _gotk4_gtk4_DragSource_ConnectPrepare(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_DragSource_ConnectDragBegin(gpointer, void*, guintptr);
// extern void _gotk4_gtk4_DragSource_ConnectDragEnd(gpointer, void*, gboolean, guintptr);
import "C"

// glib.Type values for gtkdragsource.go.
var GTypeDragSource = coreglib.Type(C.gtk_drag_source_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDragSource, F: marshalDragSource},
	})
}

// DragSourceOverrider contains methods that are overridable.
type DragSourceOverrider interface {
}

// DragSource: GtkDragSource is an event controller to initiate Drag-And-Drop
// operations.
//
// GtkDragSource can be set up with the necessary ingredients for a DND
// operation ahead of time. This includes the source for the data that is being
// transferred, in the form of a gdk.ContentProvider, the desired action, and
// the icon to use during the drag operation. After setting it up, the drag
// source must be added to a widget as an event controller, using
// gtk.Widget.AddController().
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      GtkDragSource *drag_source = gtk_drag_source_new ();
//
//      g_signal_connect (drag_source, "prepare", G_CALLBACK (on_drag_prepare), self);
//      g_signal_connect (drag_source, "drag-begin", G_CALLBACK (on_drag_begin), self);
//
//      gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (drag_source));
//    }
//
//
// Setting up the content provider and icon ahead of time only makes sense when
// the data does not change. More commonly, you will want to set them up just in
// time. To do so, GtkDragSource has gtk.DragSource::prepare and
// gtk.DragSource::drag-begin signals.
//
// The ::prepare signal is emitted before a drag is started, and can be used to
// set the content provider and actions that the drag should be started with.
//
//    static GdkContentProvider *
//    on_drag_prepare (GtkDragSource *source,
//                     double         x,
//                     double         y,
//                     MyWidget      *self)
//    {
//      // This widget supports two types of content: GFile objects
//      // and GdkPixbuf objects; GTK will handle the serialization
//      // of these types automatically
//      GFile *file = my_widget_get_file (self);
//      GdkPixbuf *pixbuf = my_widget_get_pixbuf (self);
//
//      return gdk_content_provider_new_union ((GdkContentProvider *[2]) {
//          gdk_content_provider_new_typed (G_TYPE_FILE, file),
//          gdk_content_provider_new_typed (GDK_TYPE_PIXBUF, pixbuf),
//        }, 2);
//    }
//
//
// The ::drag-begin signal is emitted after the GdkDrag object has been created,
// and can be used to set up the drag icon.
//
//    static void
//    on_drag_begin (GtkDragSource *source,
//                   GtkDrag       *drag,
//                   MyWidget      *self)
//    {
//      // Set the widget as the drag icon
//      GdkPaintable *paintable = gtk_widget_paintable_new (GTK_WIDGET (self));
//      gtk_drag_source_set_icon (source, paintable, 0, 0);
//      g_object_unref (paintable);
//    }
//
//
// During the DND operation, GtkDragSource emits signals that can be used to
// obtain updates about the status of the operation, but it is not normally
// necessary to connect to any signals, except for one case: when the supported
// actions include GDK_ACTION_MOVE, you need to listen for the
// gtk.DragSource::drag-end signal and delete the data after it has been
// transferred.
type DragSource struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*DragSource)(nil)
)

func classInitDragSourcer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDragSource(obj *coreglib.Object) *DragSource {
	return &DragSource{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalDragSource(p uintptr) (interface{}, error) {
	return wrapDragSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_DragSource_ConnectDragBegin
func _gotk4_gtk4_DragSource_ConnectDragBegin(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(drag gdk.Dragger)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(drag gdk.Dragger))
	}

	var _drag gdk.Dragger // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Dragger is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Dragger)
			return ok
		})
		rv, ok := casted.(gdk.Dragger)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dragger")
		}
		_drag = rv
	}

	f(_drag)
}

// ConnectDragBegin is emitted on the drag source when a drag is started.
//
// It can be used to e.g. set a custom drag icon with gtk.DragSource.SetIcon().
func (source *DragSource) ConnectDragBegin(f func(drag gdk.Dragger)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(source, "drag-begin", false, unsafe.Pointer(C._gotk4_gtk4_DragSource_ConnectDragBegin), f)
}

//export _gotk4_gtk4_DragSource_ConnectDragEnd
func _gotk4_gtk4_DragSource_ConnectDragEnd(arg0 C.gpointer, arg1 *C.void, arg2 C.gboolean, arg3 C.guintptr) {
	var f func(drag gdk.Dragger, deleteData bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(drag gdk.Dragger, deleteData bool))
	}

	var _drag gdk.Dragger // out
	var _deleteData bool  // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Dragger is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Dragger)
			return ok
		})
		rv, ok := casted.(gdk.Dragger)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dragger")
		}
		_drag = rv
	}
	if arg2 != 0 {
		_deleteData = true
	}

	f(_drag, _deleteData)
}

// ConnectDragEnd is emitted on the drag source when a drag is finished.
//
// A typical reason to connect to this signal is to undo things done in
// gtk.DragSource::prepare or gtk.DragSource::drag-begin handlers.
func (source *DragSource) ConnectDragEnd(f func(drag gdk.Dragger, deleteData bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(source, "drag-end", false, unsafe.Pointer(C._gotk4_gtk4_DragSource_ConnectDragEnd), f)
}

//export _gotk4_gtk4_DragSource_ConnectPrepare
func _gotk4_gtk4_DragSource_ConnectPrepare(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) (cret *C.GdkContentProvider) {
	var f func(x, y float64) (contentProvider *gdk.ContentProvider)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64) (contentProvider *gdk.ContentProvider))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	contentProvider := f(_x, _y)

	if contentProvider != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(contentProvider).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(contentProvider).Native()))
	}

	return cret
}

// ConnectPrepare is emitted when a drag is about to be initiated.
//
// It returns the GdkContentProvider to use for the drag that is about to start.
// The default handler for this signal returns the value of the
// gtk.DragSource:content property, so if you set up that property ahead of
// time, you don't need to connect to this signal.
func (source *DragSource) ConnectPrepare(f func(x, y float64) (contentProvider *gdk.ContentProvider)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(source, "prepare", false, unsafe.Pointer(C._gotk4_gtk4_DragSource_ConnectPrepare), f)
}

// NewDragSource creates a new GtkDragSource object.
//
// The function returns the following values:
//
//    - dragSource: new GtkDragSource.
//
func NewDragSource() *DragSource {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "DragSource").InvokeMethod("new_DragSource", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _dragSource *DragSource // out

	_dragSource = wrapDragSource(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dragSource
}

// DragCancel cancels a currently ongoing drag operation.
func (source *DragSource) DragCancel() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "DragSource").InvokeMethod("drag_cancel", _args[:], nil)

	runtime.KeepAlive(source)
}

// Content gets the current content provider of a GtkDragSource.
//
// The function returns the following values:
//
//    - contentProvider (optional): GdkContentProvider of source.
//
func (source *DragSource) Content() *gdk.ContentProvider {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DragSource").InvokeMethod("get_content", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(source)

	var _contentProvider *gdk.ContentProvider // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_contentProvider = &gdk.ContentProvider{
				Object: obj,
			}
		}
	}

	return _contentProvider
}

// Drag returns the underlying GdkDrag object for an ongoing drag.
//
// The function returns the following values:
//
//    - drag (optional): GdkDrag of the current drag operation, or NULL.
//
func (source *DragSource) Drag() gdk.Dragger {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DragSource").InvokeMethod("get_drag", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(source)

	var _drag gdk.Dragger // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Dragger)
				return ok
			})
			rv, ok := casted.(gdk.Dragger)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dragger")
			}
			_drag = rv
		}
	}

	return _drag
}

// SetContent sets a content provider on a GtkDragSource.
//
// When the data is requested in the cause of a DND operation, it will be
// obtained from the content provider.
//
// This function can be called before a drag is started, or in a handler for the
// gtk.DragSource::prepare signal.
//
// You may consider setting the content provider back to NULL in a
// gtk.DragSource::drag-end signal handler.
//
// The function takes the following parameters:
//
//    - content (optional): GdkContentProvider, or NULL.
//
func (source *DragSource) SetContent(content *gdk.ContentProvider) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))
	if content != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "DragSource").InvokeMethod("set_content", _args[:], nil)

	runtime.KeepAlive(source)
	runtime.KeepAlive(content)
}

// SetIcon sets a paintable to use as icon during DND operations.
//
// The hotspot coordinates determine the point on the icon that gets aligned
// with the hotspot of the cursor.
//
// If paintable is NULL, a default icon is used.
//
// This function can be called before a drag is started, or in a
// gtk.DragSource::prepare or gtk.DragSource::drag-begin signal handler.
//
// The function takes the following parameters:
//
//    - paintable (optional) to use as icon, or NULL.
//    - hotX: hotspot X coordinate on the icon.
//    - hotY: hotspot Y coordinate on the icon.
//
func (source *DragSource) SetIcon(paintable gdk.Paintabler, hotX, hotY int32) {
	var _args [4]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.int   // out
	var _arg3 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))
	if paintable != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}
	_arg2 = C.int(hotX)
	_arg3 = C.int(hotY)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.int)(unsafe.Pointer(&_args[3])) = _arg3

	girepository.MustFind("Gtk", "DragSource").InvokeMethod("set_icon", _args[:], nil)

	runtime.KeepAlive(source)
	runtime.KeepAlive(paintable)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragCheckThreshold checks to see if a drag movement has passed the GTK drag
// threshold.
//
// The function takes the following parameters:
//
//    - startX: x coordinate of start of drag.
//    - startY: y coordinate of start of drag.
//    - currentX: current X coordinate.
//    - currentY: current Y coordinate.
//
// The function returns the following values:
//
//    - ok: TRUE if the drag threshold has been passed.
//
func (widget *Widget) DragCheckThreshold(startX, startY, currentX, currentY int32) bool {
	var _args [5]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out
	var _arg3 C.int      // out
	var _arg4 C.int      // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = C.int(startX)
	_arg2 = C.int(startY)
	_arg3 = C.int(currentX)
	_arg4 = C.int(currentY)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.int)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.int)(unsafe.Pointer(&_args[4])) = _arg4

	_gret := girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_check_threshold", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)
	runtime.KeepAlive(startX)
	runtime.KeepAlive(startY)
	runtime.KeepAlive(currentX)
	runtime.KeepAlive(currentY)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
