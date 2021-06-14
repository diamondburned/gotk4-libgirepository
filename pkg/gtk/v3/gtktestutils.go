// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TestListAllTypes: return the type ids that have been registered after calling
// gtk_test_register_all_types().
func TestListAllTypes() []externglib.Type {
	var _cret *C.GType
	var _arg1 C.guint // in

	_cret = C.gtk_test_list_all_types(&_arg1)

	var _gTypes []externglib.Type

	{
		src := unsafe.Slice(_cret, _arg1)
		_gTypes = make([]externglib.Type, _arg1)
		for i := 0; i < int(_arg1); i++ {
			_gTypes[i] = externglib.Type(src[i])
		}
	}

	return _gTypes
}

// TestRegisterAllTypes: force registration of all core Gtk+ and Gdk object
// types. This allowes to refer to any of those object types via
// g_type_from_name() after calling this function.
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()
}

// TestSliderGetValue: retrive the literal adjustment value for GtkRange based
// widgets and spin buttons. Note that the value returned by this function is
// anything between the lower and upper bounds of the adjustment belonging to
// @widget, and is not a percentage as passed in to gtk_test_slider_set_perc().
func TestSliderGetValue(widget Widget) float64 {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.double // in

	_cret = C.gtk_test_slider_get_value(_arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// TestSliderSetPerc: this function will adjust the slider position of all
// GtkRange based widgets, such as scrollbars or scales, it’ll also adjust spin
// buttons. The adjustment value of these widgets is set to a value between the
// lower and upper limits, according to the @percentage argument.
func TestSliderSetPerc(widget Widget, percentage float64) {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.double(percentage)

	C.gtk_test_slider_set_perc(_arg1, _arg2)
}

// TestSpinButtonClick: this function will generate a @button click in the
// upwards or downwards spin button arrow areas, usually leading to an increase
// or decrease of spin button’s value.
func TestSpinButtonClick(spinner SpinButton, button uint, upwards bool) bool {
	var _arg1 *C.GtkSpinButton // out
	var _arg2 C.guint          // out
	var _arg3 C.gboolean       // out

	_arg1 = (*C.GtkSpinButton)(unsafe.Pointer(spinner.Native()))
	_arg2 = C.guint(button)
	if upwards {
		_arg3 = C.TRUE
	}

	var _cret C.gboolean // in

	_cret = C.gtk_test_spin_button_click(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestTextGet: retrive the text string of @widget if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
func TestTextGet(widget Widget) string {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_test_text_get(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// TestTextSet: set the text string of @widget to @string if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
func TestTextSet(widget Widget, string string) {
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_test_text_set(_arg1, _arg2)
}

// TestWidgetClick: this function will generate a @button click (button press
// and button release event) in the middle of the first GdkWindow found that
// belongs to @widget. For windowless widgets like Button (which returns false
// from gtk_widget_get_has_window()), this will often be an input-only event
// window. For other widgets, this is usually widget->window. Certain caveats
// should be considered when using this function, in particular because the
// mouse pointer is warped to the button click location, see
// gdk_test_simulate_button() for details.
func TestWidgetClick(widget Widget, button uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(button)
	_arg3 = (C.GdkModifierType)(modifiers)

	var _cret C.gboolean // in

	_cret = C.gtk_test_widget_click(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestWidgetSendKey: this function will generate keyboard press and release
// events in the middle of the first GdkWindow found that belongs to @widget.
// For windowless widgets like Button (which returns false from
// gtk_widget_get_has_window()), this will often be an input-only event window.
// For other widgets, this is usually widget->window. Certain caveats should be
// considered when using this function, in particular because the mouse pointer
// is warped to the key press location, see gdk_test_simulate_key() for details.
func TestWidgetSendKey(widget Widget, keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(keyval)
	_arg3 = (C.GdkModifierType)(modifiers)

	var _cret C.gboolean // in

	_cret = C.gtk_test_widget_send_key(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestWidgetWaitForDraw enters the main loop and waits for @widget to be
// “drawn”. In this context that means it waits for the frame clock of @widget
// to have run a full styling, layout and drawing cycle.
//
// This function is intended to be used for syncing with actions that depend on
// @widget relayouting or on interaction with the display server.
func TestWidgetWaitForDraw(widget Widget) {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_test_widget_wait_for_draw(_arg1)
}
