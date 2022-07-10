// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeWindowGroup returns the GType for the type WindowGroup.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeWindowGroup() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "WindowGroup").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalWindowGroup)
	return gtype
}

// WindowGroupOverrider contains methods that are overridable.
type WindowGroupOverrider interface {
}

// WindowGroup: GtkWindowGroup makes group of windows behave like separate
// applications.
//
// It achieves this by limiting the effect of GTK grabs and modality to windows
// in the same group.
//
// A window can be a member in at most one window group at a time. Windows that
// have not been explicitly assigned to a group are implicitly treated like
// windows of the default window group.
//
// GtkWindowGroup objects are referenced by each window in the group, so once
// you have added all windows to a GtkWindowGroup, you can drop the initial
// reference to the window group with g_object_unref(). If the windows in the
// window group are subsequently destroyed, then they will be removed from the
// window group and drop their references on the window group; when all window
// have been removed, the window group will be freed.
type WindowGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WindowGroup)(nil)
)

func classInitWindowGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWindowGroup(obj *coreglib.Object) *WindowGroup {
	return &WindowGroup{
		Object: obj,
	}
}

func marshalWindowGroup(p uintptr) (interface{}, error) {
	return wrapWindowGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowGroup creates a new GtkWindowGroup object.
//
// Modality of windows only affects windows within the same GtkWindowGroup.
//
// The function returns the following values:
//
//    - windowGroup: new GtkWindowGroup.
//
func NewWindowGroup() *WindowGroup {
	_info := girepository.MustFind("Gtk", "WindowGroup")
	_gret := _info.InvokeClassMethod("new_WindowGroup", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _windowGroup *WindowGroup // out

	_windowGroup = wrapWindowGroup(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _windowGroup
}

// AddWindow adds a window to a GtkWindowGroup.
//
// The function takes the following parameters:
//
//    - window: GtkWindow to add.
//
func (windowGroup *WindowGroup) AddWindow(window *Window) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_info := girepository.MustFind("Gtk", "WindowGroup")
	_info.InvokeClassMethod("add_window", _args[:], nil)

	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}

// ListWindows returns a list of the GtkWindows that belong to window_group.
//
// The function returns the following values:
//
//    - list: a newly-allocated list of windows inside the group.
//
func (windowGroup *WindowGroup) ListWindows() []*Window {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))

	_info := girepository.MustFind("Gtk", "WindowGroup")
	_gret := _info.InvokeClassMethod("list_windows", _args[:], nil)
	_cret := *(**C.GList)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(windowGroup)

	var _list []*Window // out

	_list = make([]*Window, 0, gextras.ListSize(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret)))))
	gextras.MoveList(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *Window // out
		dst = wrapWindow(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))))
		_list = append(_list, dst)
	})

	return _list
}

// RemoveWindow removes a window from a GtkWindowGroup.
//
// The function takes the following parameters:
//
//    - window: GtkWindow to remove.
//
func (windowGroup *WindowGroup) RemoveWindow(window *Window) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_info := girepository.MustFind("Gtk", "WindowGroup")
	_info.InvokeClassMethod("remove_window", _args[:], nil)

	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}
