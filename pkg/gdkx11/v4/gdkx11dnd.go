// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_drag_get_type()), F: marshalX11Drag},
	})
}

type X11Drag interface {
	gextras.Objector

	// AsDrag casts the class to the gdk.Drag interface.
	AsDrag() gdk.Drag

	// DropDone informs GDK that the drop ended.
	//
	// Passing false for @success may trigger a drag cancellation animation.
	//
	// This function is called by the drag source, and should be the last call
	// before dropping the reference to the @drag.
	//
	// The `GdkDrag` will only take the first [method@Gdk.Drag.drop_done] call
	// as effective, if this function is called multiple times, all subsequent
	// calls will be ignored.
	//
	// This method is inherited from gdk.Drag
	DropDone(success bool)
	// GetActions determines the bitmask of possible actions proposed by the
	// source.
	//
	// This method is inherited from gdk.Drag
	GetActions() gdk.DragAction
	// GetContent returns the `GdkContentProvider` associated to the `GdkDrag`
	// object.
	//
	// This method is inherited from gdk.Drag
	GetContent() gdk.ContentProvider
	// GetDevice returns the `GdkDevice` associated to the `GdkDrag` object.
	//
	// This method is inherited from gdk.Drag
	GetDevice() gdk.Device
	// GetDisplay gets the `GdkDisplay` that the drag object was created for.
	//
	// This method is inherited from gdk.Drag
	GetDisplay() gdk.Display
	// GetDragSurface returns the surface on which the drag icon should be
	// rendered during the drag operation.
	//
	// Note that the surface may not be available until the drag operation has
	// begun. GDK will move the surface in accordance with the ongoing drag
	// operation. The surface is owned by @drag and will be destroyed when the
	// drag operation is over.
	//
	// This method is inherited from gdk.Drag
	GetDragSurface() gdk.Surface
	// GetFormats retrieves the formats supported by this `GdkDrag` object.
	//
	// This method is inherited from gdk.Drag
	GetFormats() *gdk.ContentFormats
	// GetSelectedAction determines the action chosen by the drag destination.
	//
	// This method is inherited from gdk.Drag
	GetSelectedAction() gdk.DragAction
	// GetSurface returns the `GdkSurface` where the drag originates.
	//
	// This method is inherited from gdk.Drag
	GetSurface() gdk.Surface
	// SetHotspot sets the position of the drag surface that will be kept under
	// the cursor hotspot.
	//
	// Initially, the hotspot is at the top left corner of the drag surface.
	//
	// This method is inherited from gdk.Drag
	SetHotspot(hotX int, hotY int)
}

// x11Drag implements the X11Drag interface.
type x11Drag struct {
	*externglib.Object
}

var _ X11Drag = (*x11Drag)(nil)

// WrapX11Drag wraps a GObject to a type that implements
// interface X11Drag. It is primarily used internally.
func WrapX11Drag(obj *externglib.Object) X11Drag {
	return x11Drag{obj}
}

func marshalX11Drag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Drag(obj), nil
}

func (x x11Drag) AsDrag() gdk.Drag {
	return gdk.WrapDrag(gextras.InternObject(x))
}

func (d x11Drag) DropDone(success bool) {
	gdk.WrapDrag(gextras.InternObject(d)).DropDone(success)
}

func (d x11Drag) GetActions() gdk.DragAction {
	return gdk.WrapDrag(gextras.InternObject(d)).GetActions()
}

func (d x11Drag) GetContent() gdk.ContentProvider {
	return gdk.WrapDrag(gextras.InternObject(d)).GetContent()
}

func (d x11Drag) GetDevice() gdk.Device {
	return gdk.WrapDrag(gextras.InternObject(d)).GetDevice()
}

func (d x11Drag) GetDisplay() gdk.Display {
	return gdk.WrapDrag(gextras.InternObject(d)).GetDisplay()
}

func (d x11Drag) GetDragSurface() gdk.Surface {
	return gdk.WrapDrag(gextras.InternObject(d)).GetDragSurface()
}

func (d x11Drag) GetFormats() *gdk.ContentFormats {
	return gdk.WrapDrag(gextras.InternObject(d)).GetFormats()
}

func (d x11Drag) GetSelectedAction() gdk.DragAction {
	return gdk.WrapDrag(gextras.InternObject(d)).GetSelectedAction()
}

func (d x11Drag) GetSurface() gdk.Surface {
	return gdk.WrapDrag(gextras.InternObject(d)).GetSurface()
}

func (d x11Drag) SetHotspot(hotX int, hotY int) {
	gdk.WrapDrag(gextras.InternObject(d)).SetHotspot(hotX, hotY)
}
