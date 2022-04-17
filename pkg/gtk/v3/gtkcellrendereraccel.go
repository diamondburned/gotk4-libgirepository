// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellRendererAccelClass_accel_cleared(GtkCellRendererAccel*, gchar*);
// extern void _gotk4_gtk3_CellRendererAccelClass_accel_edited(GtkCellRendererAccel*, gchar*, guint, GdkModifierType, guint);
// extern void _gotk4_gtk3_CellRendererAccel_ConnectAccelCleared(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk3_CellRendererAccel_ConnectAccelEdited(gpointer, gchar*, guint, GdkModifierType, guint, guintptr);
import "C"

// glib.Type values for gtkcellrendereraccel.go.
var (
	GTypeCellRendererAccelMode = externglib.Type(C.gtk_cell_renderer_accel_mode_get_type())
	GTypeCellRendererAccel     = externglib.Type(C.gtk_cell_renderer_accel_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCellRendererAccelMode, F: marshalCellRendererAccelMode},
		{T: GTypeCellRendererAccel, F: marshalCellRendererAccel},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK+
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK+ are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode C.gint

const (
	// CellRendererAccelModeGTK: GTK+ accelerators mode.
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// CellRendererAccelModeOther: other accelerator mode.
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CellRendererAccelMode.
func (c CellRendererAccelMode) String() string {
	switch c {
	case CellRendererAccelModeGTK:
		return "GTK"
	case CellRendererAccelModeOther:
		return "Other"
	default:
		return fmt.Sprintf("CellRendererAccelMode(%d)", c)
	}
}

// CellRendererAccelOverrider contains methods that are overridable.
type CellRendererAccelOverrider interface {
	externglib.Objector
	// The function takes the following parameters:
	//
	AccelCleared(pathString string)
	// The function takes the following parameters:
	//
	//    - pathString
	//    - accelKey
	//    - accelMods
	//    - hardwareKeycode
	//
	AccelEdited(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
}

// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like Control + a). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
//
// The CellRendererAccel cell renderer was added in GTK+ 2.10.
type CellRendererAccel struct {
	_ [0]func() // equal guard
	CellRendererText
}

var (
	_ CellRendererer = (*CellRendererAccel)(nil)
)

func classInitCellRendererAcceller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellRendererAccelClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCellRendererAccelClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ AccelCleared(pathString string) }); ok {
		pclass.accel_cleared = (*[0]byte)(C._gotk4_gtk3_CellRendererAccelClass_accel_cleared)
	}

	if _, ok := goval.(interface {
		AccelEdited(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
	}); ok {
		pclass.accel_edited = (*[0]byte)(C._gotk4_gtk3_CellRendererAccelClass_accel_edited)
	}
}

//export _gotk4_gtk3_CellRendererAccelClass_accel_cleared
func _gotk4_gtk3_CellRendererAccelClass_accel_cleared(arg0 *C.GtkCellRendererAccel, arg1 *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ AccelCleared(pathString string) })

	var _pathString string // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.AccelCleared(_pathString)
}

//export _gotk4_gtk3_CellRendererAccelClass_accel_edited
func _gotk4_gtk3_CellRendererAccelClass_accel_edited(arg0 *C.GtkCellRendererAccel, arg1 *C.gchar, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		AccelEdited(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
	})

	var _pathString string          // out
	var _accelKey uint              // out
	var _accelMods gdk.ModifierType // out
	var _hardwareKeycode uint       // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_accelKey = uint(arg2)
	_accelMods = gdk.ModifierType(arg3)
	_hardwareKeycode = uint(arg4)

	iface.AccelEdited(_pathString, _accelKey, _accelMods, _hardwareKeycode)
}

func wrapCellRendererAccel(obj *externglib.Object) *CellRendererAccel {
	return &CellRendererAccel{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	return wrapCellRendererAccel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_CellRendererAccel_ConnectAccelCleared
func _gotk4_gtk3_CellRendererAccel_ConnectAccelCleared(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(pathString string)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pathString string))
	}

	var _pathString string // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_pathString)
}

// ConnectAccelCleared gets emitted when the user has removed the accelerator.
func (accel *CellRendererAccel) ConnectAccelCleared(f func(pathString string)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(accel, "accel-cleared", false, unsafe.Pointer(C._gotk4_gtk3_CellRendererAccel_ConnectAccelCleared), f)
}

//export _gotk4_gtk3_CellRendererAccel_ConnectAccelEdited
func _gotk4_gtk3_CellRendererAccel_ConnectAccelEdited(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guint, arg5 C.guintptr) {
	var f func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg5))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint))
	}

	var _pathString string          // out
	var _accelKey uint              // out
	var _accelMods gdk.ModifierType // out
	var _hardwareKeycode uint       // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_accelKey = uint(arg2)
	_accelMods = gdk.ModifierType(arg3)
	_hardwareKeycode = uint(arg4)

	f(_pathString, _accelKey, _accelMods, _hardwareKeycode)
}

// ConnectAccelEdited gets emitted when the user has selected a new accelerator.
func (accel *CellRendererAccel) ConnectAccelEdited(f func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(accel, "accel-edited", false, unsafe.Pointer(C._gotk4_gtk3_CellRendererAccel_ConnectAccelEdited), f)
}

// NewCellRendererAccel creates a new CellRendererAccel.
//
// The function returns the following values:
//
//    - cellRendererAccel: new cell renderer.
//
func NewCellRendererAccel() *CellRendererAccel {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_accel_new()

	var _cellRendererAccel *CellRendererAccel // out

	_cellRendererAccel = wrapCellRendererAccel(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererAccel
}
