// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeDestDefaults returns the GType for the type DestDefaults.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDestDefaults() coreglib.Type {
	gtype := coreglib.Type(C.gtk_dest_defaults_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalDestDefaults)
	return gtype
}

// DestDefaults enumeration specifies the various types of action that will be
// taken on behalf of the user for a drag destination site.
type DestDefaults C.guint

const (
	// DestDefaultMotion: if set for a widget, GTK+, during a drag over this
	// widget will check if the drag matches this widget’s list of possible
	// targets and actions. GTK+ will then call gdk_drag_status() as
	// appropriate.
	DestDefaultMotion DestDefaults = 0b1
	// DestDefaultHighlight: if set for a widget, GTK+ will draw a highlight on
	// this widget as long as a drag is over this widget and the widget drag
	// format and action are acceptable.
	DestDefaultHighlight DestDefaults = 0b10
	// DestDefaultDrop: if set for a widget, when a drop occurs, GTK+ will will
	// check if the drag matches this widget’s list of possible targets and
	// actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the
	// widget. Whether or not the drop is successful, GTK+ will call
	// gtk_drag_finish(). If the action was a move, then if the drag was
	// successful, then TRUE will be passed for the delete parameter to
	// gtk_drag_finish().
	DestDefaultDrop DestDefaults = 0b100
	// DestDefaultAll: if set, specifies that all default actions should be
	// taken.
	DestDefaultAll DestDefaults = 0b111
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	return DestDefaults(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DestDefaults.
func (d DestDefaults) String() string {
	if d == 0 {
		return "DestDefaults(0)"
	}

	var builder strings.Builder
	builder.Grow(69)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DestDefaultMotion:
			builder.WriteString("Motion|")
		case DestDefaultHighlight:
			builder.WriteString("Highlight|")
		case DestDefaultDrop:
			builder.WriteString("Drop|")
		case DestDefaultAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("DestDefaults(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DestDefaults) Has(other DestDefaults) bool {
	return (d & other) == other
}

// DragDestAddImageTargets: add the image targets supported by SelectionData to
// the target list of the drag destination. The targets are added with info = 0.
// If you need another value, use gtk_target_list_add_image_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddImageTargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_drag_dest_add_image_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragDestAddTextTargets: add the text targets supported by SelectionData to
// the target list of the drag destination. The targets are added with info = 0.
// If you need another value, use gtk_target_list_add_text_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddTextTargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_drag_dest_add_text_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragDestAddURITargets: add the URI targets supported by SelectionData to the
// target list of the drag destination. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_uri_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddURITargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_drag_dest_add_uri_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragDestGetTargetList returns the list of targets this widget can accept from
// drag-and-drop.
//
// The function returns the following values:
//
//    - targetList (optional) or NULL if none.
//
func (widget *Widget) DragDestGetTargetList() *TargetList {
	var _arg0 *C.GtkWidget     // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_drag_dest_get_target_list(_arg0)
	runtime.KeepAlive(widget)

	var _targetList *TargetList // out

	if _cret != nil {
		_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_target_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_targetList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _targetList
}

// DragDestGetTrackMotion returns whether the widget has been configured to
// always emit Widget::drag-motion signals.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget always emits Widget::drag-motion events.
//
func (widget *Widget) DragDestGetTrackMotion() bool {
	var _arg0 *C.GtkWidget // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_drag_dest_get_track_motion(_arg0)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDestSet sets a widget as a potential drop destination, and adds default
// behaviors.
//
// The default behaviors listed in flags have an effect similar to installing
// default handlers for the widget’s drag-and-drop signals (Widget::drag-motion,
// Widget::drag-drop, ...). They all exist for convenience. When passing
// K_DEST_DEFAULT_ALL for instance it is sufficient to connect to the widget’s
// Widget::drag-data-received signal to get primitive, but consistent
// drag-and-drop support.
//
// Things become more complicated when you try to preview the dragged data, as
// described in the documentation for Widget::drag-motion. The default behaviors
// described by flags make some assumptions, that can conflict with your own
// signal handlers. For instance K_DEST_DEFAULT_DROP causes invokations of
// gdk_drag_status() in the context of Widget::drag-motion, and invokations of
// gtk_drag_finish() in Widget::drag-data-received. Especially the later is
// dramatic, when your own Widget::drag-motion handler calls gtk_drag_get_data()
// to inspect the dragged data.
//
// There’s no way to set a default action here, you can use the
// Widget::drag-motion callback for that. Here’s an example which selects the
// action to use depending on whether the control key is pressed or not:
//
//    static void
//    drag_motion (GtkWidget *widget,
//                 GdkDragContext *context,
//                 gint x,
//                 gint y,
//                 guint time)
//    {
//      GdkModifierType mask;
//
//      gdk_window_get_pointer (gtk_widget_get_window (widget),
//                              NULL, NULL, &mask);
//      if (mask & GDK_CONTROL_MASK)
//        gdk_drag_status (context, GDK_ACTION_COPY, time);
//      else
//        gdk_drag_status (context, GDK_ACTION_MOVE, time);
//    }.
//
// The function takes the following parameters:
//
//    - flags: which types of default drag behavior to use.
//    - targets (optional): pointer to an array of TargetEntrys indicating the
//      drop types that this widget will accept, or NULL. Later you can access
//      the list with gtk_drag_dest_get_target_list() and
//      gtk_drag_dest_find_target().
//    - actions: bitmask of possible actions for a drop onto this widget.
//
func (widget *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	var _arg0 *C.GtkWidget      // out
	var _arg1 C.GtkDestDefaults // out
	var _arg2 *C.GtkTargetEntry // out
	var _arg3 C.gint
	var _arg4 C.GdkDragAction // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = C.GtkDestDefaults(flags)
	_arg3 = (C.gint)(len(targets))
	_arg2 = (*C.GtkTargetEntry)(C.calloc(C.size_t(len(targets)), C.size_t(C.sizeof_GtkTargetEntry)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GtkTargetEntry)(_arg2), len(targets))
		for i := range targets {
			out[i] = *(*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer((&targets[i]))))
		}
	}
	_arg4 = C.GdkDragAction(actions)

	C.gtk_drag_dest_set(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(actions)
}

// DragDestSetProxy sets this widget as a proxy for drops to another window.
//
// Deprecated: since version 3.22.
//
// The function takes the following parameters:
//
//    - proxyWindow: window to which to forward drag events.
//    - protocol: drag protocol which the proxy_window accepts (You can use
//      gdk_drag_get_protocol() to determine this).
//    - useCoordinates: if TRUE, send the same coordinates to the destination,
//      because it is an embedded subwindow.
//
func (widget *Widget) DragDestSetProxy(proxyWindow gdk.Windower, protocol gdk.DragProtocol, useCoordinates bool) {
	var _arg0 *C.GtkWidget      // out
	var _arg1 *C.GdkWindow      // out
	var _arg2 C.GdkDragProtocol // out
	var _arg3 C.gboolean        // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(proxyWindow).Native()))
	_arg2 = C.GdkDragProtocol(protocol)
	if useCoordinates {
		_arg3 = C.TRUE
	}

	C.gtk_drag_dest_set_proxy(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(proxyWindow)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(useCoordinates)
}

// DragDestSetTargetList sets the target types that this widget can accept from
// drag-and-drop. The widget must first be made into a drag destination with
// gtk_drag_dest_set().
//
// The function takes the following parameters:
//
//    - targetList (optional): list of droppable targets, or NULL for none.
//
func (widget *Widget) DragDestSetTargetList(targetList *TargetList) {
	var _arg0 *C.GtkWidget     // out
	var _arg1 *C.GtkTargetList // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if targetList != nil {
		_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(targetList)))
	}

	C.gtk_drag_dest_set_target_list(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(targetList)
}

// DragDestSetTrackMotion tells the widget to emit Widget::drag-motion and
// Widget::drag-leave events regardless of the targets and the
// GTK_DEST_DEFAULT_MOTION flag.
//
// This may be used when a widget wants to do generic actions regardless of the
// targets that the source offers.
//
// The function takes the following parameters:
//
//    - trackMotion: whether to accept all targets.
//
func (widget *Widget) DragDestSetTrackMotion(trackMotion bool) {
	var _arg0 *C.GtkWidget // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if trackMotion {
		_arg1 = C.TRUE
	}

	C.gtk_drag_dest_set_track_motion(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(trackMotion)
}

// DragDestUnset clears information about a drop destination set with
// gtk_drag_dest_set(). The widget will no longer receive notification of drags.
func (widget *Widget) DragDestUnset() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_drag_dest_unset(_arg0)
	runtime.KeepAlive(widget)
}
