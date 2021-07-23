// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TestCreateSimpleWindow: create a simple window with window title window_title
// and text contents dialog_text. The window will quit any running
// gtk_main()-loop when destroyed, and it will automatically be destroyed upon
// test function teardown.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestCreateSimpleWindow(windowTitle string, dialogText string) Widgetter {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(windowTitle)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(dialogText)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_test_create_simple_window(_arg1, _arg2)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// TestFindLabel: this function will search widget and all its descendants for a
// GtkLabel widget with a text string matching label_pattern. The label_pattern
// may contain asterisks “*” and question marks “?” as placeholders,
// g_pattern_match() is used for the matching. Note that locales other than "C“
// tend to alter (translate” label strings, so this function is genrally only
// useful in test programs with predetermined locales, see gtk_test_init() for
// more details.
func TestFindLabel(widget Widgetter, labelPattern string) Widgetter {
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(labelPattern)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_test_find_label(_arg1, _arg2)

	var _ret Widgetter // out

	_ret = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _ret
}

// TestFindSibling: this function will search siblings of base_widget and
// siblings of its ancestors for all widgets matching widget_type. Of the
// matching widgets, the one that is geometrically closest to base_widget will
// be returned. The general purpose of this function is to find the most likely
// “action” widget, relative to another labeling widget. Such as finding a
// button or text entry widget, given its corresponding label widget.
func TestFindSibling(baseWidget Widgetter, widgetType externglib.Type) Widgetter {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GType      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(baseWidget.Native()))
	_arg2 = C.GType(widgetType)

	_cret = C.gtk_test_find_sibling(_arg1, _arg2)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// TestFindWidget: this function will search the descendants of widget for a
// widget of type widget_type that has a label matching label_pattern next to
// it. This is most useful for automated GUI testing, e.g. to find the “OK”
// button in a dialog and synthesize clicks on it. However see
// gtk_test_find_label(), gtk_test_find_sibling() and gtk_test_widget_click()
// for possible caveats involving the search of such widgets and synthesizing
// widget events.
func TestFindWidget(widget Widgetter, labelPattern string, widgetType externglib.Type) Widgetter {
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out
	var _arg3 C.GType      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(labelPattern)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.GType(widgetType)

	_cret = C.gtk_test_find_widget(_arg1, _arg2, _arg3)

	var _ret Widgetter // out

	if _cret != nil {
		_ret = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _ret
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
// widget, and is not a percentage as passed in to gtk_test_slider_set_perc().
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestSliderGetValue(widget Widgetter) float64 {
	var _arg1 *C.GtkWidget // out
	var _cret C.double     // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_test_slider_get_value(_arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// TestSliderSetPerc: this function will adjust the slider position of all
// GtkRange based widgets, such as scrollbars or scales, it’ll also adjust spin
// buttons. The adjustment value of these widgets is set to a value between the
// lower and upper limits, according to the percentage argument.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestSliderSetPerc(widget Widgetter, percentage float64) {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.double(percentage)

	C.gtk_test_slider_set_perc(_arg1, _arg2)
}

// TestSpinButtonClick: this function will generate a button click in the
// upwards or downwards spin button arrow areas, usually leading to an increase
// or decrease of spin button’s value.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestSpinButtonClick(spinner *SpinButton, button uint, upwards bool) bool {
	var _arg1 *C.GtkSpinButton // out
	var _arg2 C.guint          // out
	var _arg3 C.gboolean       // out
	var _cret C.gboolean       // in

	_arg1 = (*C.GtkSpinButton)(unsafe.Pointer(spinner.Native()))
	_arg2 = C.guint(button)
	if upwards {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_test_spin_button_click(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestTextGet: retrive the text string of widget if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestTextGet(widget Widgetter) string {
	var _arg1 *C.GtkWidget // out
	var _cret *C.gchar     // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_test_text_get(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// TestTextSet: set the text string of widget to string if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestTextSet(widget Widgetter, _string string) {
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_test_text_set(_arg1, _arg2)
}

// TestWidgetClick: this function will generate a button click (button press and
// button release event) in the middle of the first GdkWindow found that belongs
// to widget. For windowless widgets like Button (which returns FALSE from
// gtk_widget_get_has_window()), this will often be an input-only event window.
// For other widgets, this is usually widget->window. Certain caveats should be
// considered when using this function, in particular because the mouse pointer
// is warped to the button click location, see gdk_test_simulate_button() for
// details.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
func TestWidgetClick(widget Widgetter, button uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(button)
	_arg3 = C.GdkModifierType(modifiers)

	_cret = C.gtk_test_widget_click(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestWidgetSendKey: this function will generate keyboard press and release
// events in the middle of the first GdkWindow found that belongs to widget. For
// windowless widgets like Button (which returns FALSE from
// gtk_widget_get_has_window()), this will often be an input-only event window.
// For other widgets, this is usually widget->window. Certain caveats should be
// considered when using this function, in particular because the mouse pointer
// is warped to the key press location, see gdk_test_simulate_key() for details.
func TestWidgetSendKey(widget Widgetter, keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	_cret = C.gtk_test_widget_send_key(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestWidgetWaitForDraw enters the main loop and waits for widget to be
// “drawn”. In this context that means it waits for the frame clock of widget to
// have run a full styling, layout and drawing cycle.
//
// This function is intended to be used for syncing with actions that depend on
// widget relayouting or on interaction with the display server.
func TestWidgetWaitForDraw(widget Widgetter) {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_test_widget_wait_for_draw(_arg1)
}
