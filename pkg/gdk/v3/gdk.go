// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_status_get_type()), F: marshalStatus},
		{T: externglib.Type(C.gdk_device_tool_get_type()), F: marshalDeviceTooler},
		{T: externglib.Type(C.gdk_drag_context_get_type()), F: marshalDragContexter},
	})
}

func GLErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_gl_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type Status int

const (
	OK         Status = 0
	Error      Status = -1
	ErrorParam Status = -2
	ErrorFile  Status = -3
	ErrorMem   Status = -4
)

func marshalStatus(p uintptr) (interface{}, error) {
	return Status(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for Status.
func (s Status) String() string {
	switch s {
	case OK:
		return "OK"
	case Error:
		return "Error"
	case ErrorParam:
		return "ErrorParam"
	case ErrorFile:
		return "ErrorFile"
	case ErrorMem:
		return "ErrorMem"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}

type DeviceTool struct {
	*externglib.Object
}

func wrapDeviceTool(obj *externglib.Object) *DeviceTool {
	return &DeviceTool{
		Object: obj,
	}
}

func marshalDeviceTooler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDeviceTool(obj), nil
}

// HardwareID gets the hardware ID of this tool, or 0 if it's not known. When
// non-zero, the identificator is unique for the given tool model, meaning that
// two identical tools will share the same hardware_id, but will have different
// serial numbers (see gdk_device_tool_get_serial()).
//
// This is a more concrete (and device specific) method to identify a DeviceTool
// than gdk_device_tool_get_tool_type(), as a tablet may support multiple
// devices with the same DeviceToolType, but having different hardware
// identificators.
func (tool *DeviceTool) HardwareID() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_hardware_id(_arg0)
	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Serial gets the serial of this tool, this value can be used to identify a
// physical tool (eg. a tablet pen) across program executions.
func (tool *DeviceTool) Serial() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_serial(_arg0)
	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// ToolType gets the DeviceToolType of the tool.
func (tool *DeviceTool) ToolType() DeviceToolType {
	var _arg0 *C.GdkDeviceTool    // out
	var _cret C.GdkDeviceToolType // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_tool_type(_arg0)
	runtime.KeepAlive(tool)

	var _deviceToolType DeviceToolType // out

	_deviceToolType = DeviceToolType(_cret)

	return _deviceToolType
}

type DragContext struct {
	*externglib.Object
}

func wrapDragContext(obj *externglib.Object) *DragContext {
	return &DragContext{
		Object: obj,
	}
}

func marshalDragContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDragContext(obj), nil
}

// Actions determines the bitmask of actions proposed by the source if
// gdk_drag_context_get_suggested_action() returns GDK_ACTION_ASK.
func (context *DragContext) Actions() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_actions(_arg0)
	runtime.KeepAlive(context)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// DestWindow returns the destination window for the DND operation.
func (context *DragContext) DestWindow() Windower {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_dest_window(_arg0)
	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// Device returns the Device associated to the drag context.
func (context *DragContext) Device() Devicer {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkDevice      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_device(_arg0)
	runtime.KeepAlive(context)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Devicer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// DragWindow returns the window on which the drag icon should be rendered
// during the drag operation. Note that the window may not be available until
// the drag operation has begun. GDK will move the window in accordance with the
// ongoing drag operation. The window is owned by context and will be destroyed
// when the drag operation is over.
func (context *DragContext) DragWindow() Windower {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_drag_window(_arg0)
	runtime.KeepAlive(context)

	var _window Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// Protocol returns the drag protocol that is used by this context.
func (context *DragContext) Protocol() DragProtocol {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragProtocol // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_protocol(_arg0)
	runtime.KeepAlive(context)

	var _dragProtocol DragProtocol // out

	_dragProtocol = DragProtocol(_cret)

	return _dragProtocol
}

// SelectedAction determines the action chosen by the drag destination.
func (context *DragContext) SelectedAction() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_selected_action(_arg0)
	runtime.KeepAlive(context)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// SourceWindow returns the Window where the DND operation started.
func (context *DragContext) SourceWindow() Windower {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_source_window(_arg0)
	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// SuggestedAction determines the suggested drag action of the context.
func (context *DragContext) SuggestedAction() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_context_get_suggested_action(_arg0)
	runtime.KeepAlive(context)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// ManageDnd requests the drag and drop operation to be managed by context. When
// a drag and drop operation becomes managed, the DragContext will internally
// handle all input and source-side EventDND events as required by the windowing
// system.
//
// Once the drag and drop operation is managed, the drag context will emit the
// following signals:
//
// - The DragContext::action-changed signal whenever the final action to be
// performed by the drag and drop operation changes.
//
// - The DragContext::drop-performed signal after the user performs the drag and
// drop gesture (typically by releasing the mouse button).
//
// - The DragContext::dnd-finished signal after the drag and drop operation
// concludes (after all Selection transfers happen).
//
// - The DragContext::cancel signal if the drag and drop operation is finished
// but doesn't happen over an accepting destination, or is cancelled through
// other means.
//
// The function takes the following parameters:
//
//    - ipcWindow: window to use for IPC messaging/events.
//    - actions supported by the drag source.
//
func (context *DragContext) ManageDnd(ipcWindow Windower, actions DragAction) bool {
	var _arg0 *C.GdkDragContext // out
	var _arg1 *C.GdkWindow      // out
	var _arg2 C.GdkDragAction   // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(ipcWindow.Native()))
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gdk_drag_context_manage_dnd(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(ipcWindow)
	runtime.KeepAlive(actions)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDevice associates a Device to context, so all Drag and Drop events for
// context are emitted as if they came from this device.
//
// The function takes the following parameters:
//
//    - device: Device.
//
func (context *DragContext) SetDevice(device Devicer) {
	var _arg0 *C.GdkDragContext // out
	var _arg1 *C.GdkDevice      // out

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gdk_drag_context_set_device(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(device)
}

// SetHotspot sets the position of the drag window that will be kept under the
// cursor hotspot. Initially, the hotspot is at the top left corner of the drag
// window.
//
// The function takes the following parameters:
//
//    - hotX: x coordinate of the drag window hotspot.
//    - hotY: y coordinate of the drag window hotspot.
//
func (context *DragContext) SetHotspot(hotX, hotY int) {
	var _arg0 *C.GdkDragContext // out
	var _arg1 C.gint            // out
	var _arg2 C.gint            // out

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(hotX)
	_arg2 = C.gint(hotY)

	C.gdk_drag_context_set_hotspot(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// ConnectActionChanged: new action is being chosen for the drag and drop
// operation.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectActionChanged(f func(action DragAction)) externglib.SignalHandle {
	return context.Connect("action-changed", f)
}

// ConnectCancel: drag and drop operation was cancelled.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectCancel(f func(reason DragCancelReason)) externglib.SignalHandle {
	return context.Connect("cancel", f)
}

// ConnectDndFinished: drag and drop operation was finished, the drag
// destination finished reading all data. The drag source can now free all
// miscellaneous data.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectDndFinished(f func()) externglib.SignalHandle {
	return context.Connect("dnd-finished", f)
}

// ConnectDropPerformed: drag and drop operation was performed on an accepting
// client.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectDropPerformed(f func(time int)) externglib.SignalHandle {
	return context.Connect("drop-performed", f)
}
