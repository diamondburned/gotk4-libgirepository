// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hscale_get_type()), F: marshalHScaler},
	})
}

// HScaleOverrider contains methods that are overridable.
type HScaleOverrider interface {
}

// HScale widget is used to allow the user to select a value using a horizontal
// slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkHScale has been deprecated, use Scale instead.
type HScale struct {
	_ [0]func() // equal guard
	Scale
}

var (
	_ Ranger = (*HScale)(nil)
)

func classInitHScaler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHScale(obj *externglib.Object) *HScale {
	return &HScale{
		Scale: Scale{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
				Object: obj,
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScaler(p uintptr) (interface{}, error) {
	return wrapHScale(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHScale creates a new HScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function takes the following parameters:
//
//    - adjustment (optional) which sets the range of the scale.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScale(adjustment *Adjustment) *HScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	_cret = C.gtk_hscale_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _hScale *HScale // out

	_hScale = wrapHScale(externglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}

// NewHScaleWithRange creates a new horizontal scale widget that lets the user
// input a number between min and max (including min and max) with the increment
// step. step must be nonzero; it’s the distance the slider moves when using the
// arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_HORIZONTAL
// instead.
//
// The function takes the following parameters:
//
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScaleWithRange(min, max, step float64) *HScale {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _arg3 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	_arg3 = C.gdouble(step)

	_cret = C.gtk_hscale_new_with_range(_arg1, _arg2, _arg3)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _hScale *HScale // out

	_hScale = wrapHScale(externglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}
