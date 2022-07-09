// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
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
// extern gboolean _gotk4_atk1_ComponentIface_grab_focus(void*);
// extern gboolean _gotk4_atk1_ComponentIface_set_size(void*, gint, gint);
// extern gdouble _gotk4_atk1_ComponentIface_get_alpha(void*);
// extern gint _gotk4_atk1_ComponentIface_get_mdi_zorder(void*);
// extern void _gotk4_atk1_ComponentIface_bounds_changed(void*, void*);
// extern void _gotk4_atk1_ComponentIface_get_size(void*, void*, void*);
// extern void _gotk4_atk1_ComponentIface_remove_focus_handler(void*, guint);
// extern void _gotk4_atk1_Component_ConnectBoundsChanged(gpointer, void*, guintptr);
import "C"

// GTypeScrollType returns the GType for the type ScrollType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScrollType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "ScrollType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalScrollType)
	return gtype
}

// GTypeComponent returns the GType for the type Component.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeComponent() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "Component").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalComponent)
	return gtype
}

// GTypeRectangle returns the GType for the type Rectangle.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRectangle() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "Rectangle").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRectangle)
	return gtype
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType C.gint

const (
	// ScrollTopLeft: scroll the object vertically and horizontally to bring its
	// top left corner to the top left corner of the window.
	ScrollTopLeft ScrollType = iota
	// ScrollBottomRight: scroll the object vertically and horizontally to bring
	// its bottom right corner to the bottom right corner of the window.
	ScrollBottomRight
	// ScrollTopEdge: scroll the object vertically to bring its top edge to the
	// top edge of the window.
	ScrollTopEdge
	// ScrollBottomEdge: scroll the object vertically to bring its bottom edge
	// to the bottom edge of the window.
	ScrollBottomEdge
	// ScrollLeftEdge: scroll the object vertically and horizontally to bring
	// its left edge to the left edge of the window.
	ScrollLeftEdge
	// ScrollRightEdge: scroll the object vertically and horizontally to bring
	// its right edge to the right edge of the window.
	ScrollRightEdge
	// ScrollAnywhere: scroll the object vertically and horizontally so that as
	// much as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollAnywhere
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScrollType.
func (s ScrollType) String() string {
	switch s {
	case ScrollTopLeft:
		return "TopLeft"
	case ScrollBottomRight:
		return "BottomRight"
	case ScrollTopEdge:
		return "TopEdge"
	case ScrollBottomEdge:
		return "BottomEdge"
	case ScrollLeftEdge:
		return "LeftEdge"
	case ScrollRightEdge:
		return "RightEdge"
	case ScrollAnywhere:
		return "Anywhere"
	default:
		return fmt.Sprintf("ScrollType(%d)", s)
	}
}

// ComponentOverrider contains methods that are overridable.
type ComponentOverrider interface {
	// The function takes the following parameters:
	//
	BoundsChanged(bounds *Rectangle)
	// Alpha returns the alpha value (i.e. the opacity) for this component, on a
	// scale from 0 (fully transparent) to 1.0 (fully opaque).
	//
	// The function returns the following values:
	//
	//    - gdouble: alpha value from 0 to 1.0, inclusive.
	//
	Alpha() float64
	// MDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	//
	// The function returns the following values:
	//
	//    - gint which is the zorder of the component, i.e. the depth at which
	//      the component is shown in relation to other components in the same
	//      container.
	//
	MDIZOrder() int32
	// Size gets the size of the component in terms of width and height.
	//
	// If the size can not be obtained (e.g. a non-embedded plug or missing
	// support), width and height are set to -1.
	//
	// Deprecated: Since 2.12. Use atk_component_get_extents() instead.
	//
	// The function returns the following values:
	//
	//    - width (optional) address of #gint to put width of component.
	//    - height (optional) address of #gint to put height of component.
	//
	Size() (width, height int32)
	// GrabFocus grabs focus for this component.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if successful, FALSE otherwise.
	//
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: If you need to track when an object gains or lose the focus,
	// use the Object::state-change "focused" notification instead.
	//
	// The function takes the following parameters:
	//
	//    - handlerId: handler id of the focus handler to be removed from
	//      component.
	//
	RemoveFocusHandler(handlerId uint32)
	// SetSize: set the size of the component in terms of width and height.
	//
	// The function takes the following parameters:
	//
	//    - width to set for component.
	//    - height to set for component.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE or FALSE whether the size was set or not.
	//
	SetSize(width, height int32) bool
}

// Component should be implemented by most if not all UI elements with an actual
// on-screen presence, i.e. components which can be said to have a
// screen-coordinate bounding box. Virtually all widgets will need to have
// Component implementations provided for their corresponding Object class. In
// short, only UI elements which are *not* GUI elements will omit this ATK
// interface.
//
// A possible exception might be textual information with a transparent
// background, in which case text glyph bounding box information is provided by
// Text.
//
// Component wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Component struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Component)(nil)
)

// Componenter describes Component's interface methods.
type Componenter interface {
	coreglib.Objector

	// Alpha returns the alpha value (i.e.
	Alpha() float64
	// MDIZOrder gets the zorder of the component.
	MDIZOrder() int32
	// Size gets the size of the component in terms of width and height.
	Size() (width, height int32)
	// GrabFocus grabs focus for this component.
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	RemoveFocusHandler(handlerId uint32)
	// SetSize: set the size of the component in terms of width and height.
	SetSize(width, height int32) bool

	// Bounds-changed: 'bounds-changed" signal is emitted when the bposition or
	// size of the component changes.
	ConnectBoundsChanged(func(arg1 *Rectangle)) coreglib.SignalHandle
}

var _ Componenter = (*Component)(nil)

func ifaceInitComponenter(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Atk", "ComponentIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("bounds_changed"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_bounds_changed)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_alpha"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_get_alpha)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_mdi_zorder"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_get_mdi_zorder)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_size"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_get_size)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("grab_focus"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_grab_focus)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("remove_focus_handler"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_remove_focus_handler)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_size"))) = unsafe.Pointer(C._gotk4_atk1_ComponentIface_set_size)
}

//export _gotk4_atk1_ComponentIface_bounds_changed
func _gotk4_atk1_ComponentIface_bounds_changed(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _bounds *Rectangle // out

	_bounds = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.BoundsChanged(_bounds)
}

//export _gotk4_atk1_ComponentIface_get_alpha
func _gotk4_atk1_ComponentIface_get_alpha(arg0 *C.void) (cret C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	gdouble := iface.Alpha()

	cret = C.gdouble(gdouble)

	return cret
}

//export _gotk4_atk1_ComponentIface_get_mdi_zorder
func _gotk4_atk1_ComponentIface_get_mdi_zorder(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	gint := iface.MDIZOrder()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ComponentIface_get_size
func _gotk4_atk1_ComponentIface_get_size(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	width, height := iface.Size()

	*arg1 = (*C.void)(unsafe.Pointer(width))
	*arg2 = (*C.void)(unsafe.Pointer(height))
}

//export _gotk4_atk1_ComponentIface_grab_focus
func _gotk4_atk1_ComponentIface_grab_focus(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	ok := iface.GrabFocus()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_ComponentIface_remove_focus_handler
func _gotk4_atk1_ComponentIface_remove_focus_handler(arg0 *C.void, arg1 C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _handlerId uint32 // out

	_handlerId = uint32(arg1)

	iface.RemoveFocusHandler(_handlerId)
}

//export _gotk4_atk1_ComponentIface_set_size
func _gotk4_atk1_ComponentIface_set_size(arg0 *C.void, arg1 C.gint, arg2 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _width int32  // out
	var _height int32 // out

	_width = int32(arg1)
	_height = int32(arg2)

	ok := iface.SetSize(_width, _height)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapComponent(obj *coreglib.Object) *Component {
	return &Component{
		Object: obj,
	}
}

func marshalComponent(p uintptr) (interface{}, error) {
	return wrapComponent(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Component_ConnectBoundsChanged
func _gotk4_atk1_Component_ConnectBoundsChanged(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(arg1 *Rectangle)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *Rectangle))
	}

	var _arg1 *Rectangle // out

	_arg1 = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_arg1)
}

// ConnectBoundsChanged: 'bounds-changed" signal is emitted when the bposition
// or size of the component changes.
func (component *Component) ConnectBoundsChanged(f func(arg1 *Rectangle)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(component, "bounds-changed", false, unsafe.Pointer(C._gotk4_atk1_Component_ConnectBoundsChanged), f)
}

// Alpha returns the alpha value (i.e. the opacity) for this component, on a
// scale from 0 (fully transparent) to 1.0 (fully opaque).
//
// The function returns the following values:
//
//    - gdouble: alpha value from 0 to 1.0, inclusive.
//
func (component *Component) Alpha() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))

	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// MDIZOrder gets the zorder of the component. The value G_MININT will be
// returned if the layer of the component is not ATK_LAYER_MDI or
// ATK_LAYER_WINDOW.
//
// The function returns the following values:
//
//    - gint which is the zorder of the component, i.e. the depth at which the
//      component is shown in relation to other components in the same container.
//
func (component *Component) MDIZOrder() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Size gets the size of the component in terms of width and height.
//
// If the size can not be obtained (e.g. a non-embedded plug or missing
// support), width and height are set to -1.
//
// Deprecated: Since 2.12. Use atk_component_get_extents() instead.
//
// The function returns the following values:
//
//    - width (optional) address of #gint to put width of component.
//    - height (optional) address of #gint to put height of component.
//
func (component *Component) Size() (width, height int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))

	runtime.KeepAlive(component)

	var _width int32  // out
	var _height int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_width = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_height = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _width, _height
}

// GrabFocus grabs focus for this component.
//
// The function returns the following values:
//
//    - ok: TRUE if successful, FALSE otherwise.
//
func (component *Component) GrabFocus() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RemoveFocusHandler: remove the handler specified by handler_id from the list
// of functions to be executed when this object receives focus events (in or
// out).
//
// Deprecated: If you need to track when an object gains or lose the focus, use
// the Object::state-change "focused" notification instead.
//
// The function takes the following parameters:
//
//    - handlerId: handler id of the focus handler to be removed from component.
//
func (component *Component) RemoveFocusHandler(handlerId uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(handlerId)

	runtime.KeepAlive(component)
	runtime.KeepAlive(handlerId)
}

// SetSize: set the size of the component in terms of width and height.
//
// The function takes the following parameters:
//
//    - width to set for component.
//    - height to set for component.
//
// The function returns the following values:
//
//    - ok: TRUE or FALSE whether the size was set or not.
//
func (component *Component) SetSize(width, height int32) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(width)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(height)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Rectangle: data structure for holding a rectangle. Those coordinates are
// relative to the component top-level parent.
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native unsafe.Pointer
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rectangle{&rectangle{(unsafe.Pointer)(b)}}, nil
}

// X coordinate of the left side of the rectangle.
func (r *Rectangle) X() int32 {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("x")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Y coordinate of the top side of the rectangle.
func (r *Rectangle) Y() int32 {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("y")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Width: width of the rectangle.
func (r *Rectangle) Width() int32 {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("width")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Height: height of the rectangle.
func (r *Rectangle) Height() int32 {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("height")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// X coordinate of the left side of the rectangle.
func (r *Rectangle) SetX(x int32) {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("x")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(x)
}

// Y coordinate of the top side of the rectangle.
func (r *Rectangle) SetY(y int32) {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("y")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(y)
}

// Width: width of the rectangle.
func (r *Rectangle) SetWidth(width int32) {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("width")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(width)
}

// Height: height of the rectangle.
func (r *Rectangle) SetHeight(height int32) {
	offset := girepository.MustFind("Atk", "Rectangle").StructFieldOffset("height")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(height)
}
