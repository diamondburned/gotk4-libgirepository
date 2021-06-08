// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TestCreateSimpleWindow: create a simple window with window title
// @window_title and text contents @dialog_text. The window will quit any
// running gtk_main()-loop when destroyed, and it will automatically be
// destroyed upon test function teardown.
func TestCreateSimpleWindow(windowTitle string, dialogText string) Widget {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(windowTitle))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(dialogText))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_test_create_simple_window(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// TestFindLabel: this function will search @widget and all its descendants for
// a GtkLabel widget with a text string matching @label_pattern. The
// @label_pattern may contain asterisks “*” and question marks “?” as
// placeholders, g_pattern_match() is used for the matching. Note that locales
// other than "C“ tend to alter (translate” label strings, so this function is
// genrally only useful in test programs with predetermined locales, see
// gtk_test_init() for more details.
func TestFindLabel(widget Widget, labelPattern string) Widget {
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (*C.gchar)(C.CString(labelPattern))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_test_find_label(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// TestFindSibling: this function will search siblings of @base_widget and
// siblings of its ancestors for all widgets matching @widget_type. Of the
// matching widgets, the one that is geometrically closest to @base_widget will
// be returned. The general purpose of this function is to find the most likely
// “action” widget, relative to another labeling widget. Such as finding a
// button or text entry widget, given its corresponding label widget.
func TestFindSibling(baseWidget Widget, widgetType externglib.Type) Widget {
	var arg1 *C.GtkWidget
	var arg2 C.GType

	arg1 = (*C.GtkWidget)(unsafe.Pointer(baseWidget.Native()))
	arg2 = C.GType(widgetType)

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_test_find_sibling(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// TestFindWidget: this function will search the descendants of @widget for a
// widget of type @widget_type that has a label matching @label_pattern next to
// it. This is most useful for automated GUI testing, e.g. to find the “OK”
// button in a dialog and synthesize clicks on it. However see
// gtk_test_find_label(), gtk_test_find_sibling() and gtk_test_widget_click()
// for possible caveats involving the search of such widgets and synthesizing
// widget events.
func TestFindWidget(widget Widget, labelPattern string, widgetType externglib.Type) Widget {
	var arg1 *C.GtkWidget
	var arg2 *C.gchar
	var arg3 C.GType

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (*C.gchar)(C.CString(labelPattern))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.GType(widgetType)

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_test_find_widget(arg1, arg2, arg3)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// TestListAllTypes: return the type ids that have been registered after calling
// gtk_test_register_all_types().
func TestListAllTypes() []externglib.Type {
	var cret *C.GType
	var arg1 *C.guint
	var goret []externglib.Type

	cret = C.gtk_test_list_all_types(arg1)

	goret = make([]externglib.Type, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (C.GType)(ptr.Add(unsafe.Pointer(cret), i))
		goret[i] = externglib.Type(src)
	}

	return ret1, goret
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
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var cret C.double
	var goret float64

	cret = C.gtk_test_slider_get_value(arg1)

	goret = float64(cret)

	return goret
}

// TestSliderSetPerc: this function will adjust the slider position of all
// GtkRange based widgets, such as scrollbars or scales, it’ll also adjust spin
// buttons. The adjustment value of these widgets is set to a value between the
// lower and upper limits, according to the @percentage argument.
func TestSliderSetPerc(widget Widget, percentage float64) {
	var arg1 *C.GtkWidget
	var arg2 C.double

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.double(percentage)

	C.gtk_test_slider_set_perc(arg1, arg2)
}

// TestSpinButtonClick: this function will generate a @button click in the
// upwards or downwards spin button arrow areas, usually leading to an increase
// or decrease of spin button’s value.
func TestSpinButtonClick(spinner SpinButton, button uint, upwards bool) bool {
	var arg1 *C.GtkSpinButton
	var arg2 C.guint
	var arg3 C.gboolean

	arg1 = (*C.GtkSpinButton)(unsafe.Pointer(spinner.Native()))
	arg2 = C.guint(button)
	if upwards {
		arg3 = C.gboolean(1)
	}

	var cret C.gboolean
	var goret bool

	cret = C.gtk_test_spin_button_click(arg1, arg2, arg3)

	if cret {
		goret = true
	}

	return goret
}

// TestTextGet: retrive the text string of @widget if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
func TestTextGet(widget Widget) string {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	cret := new(C.gchar)
	var goret string

	cret = C.gtk_test_text_get(arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// TestTextSet: set the text string of @widget to @string if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
func TestTextSet(widget Widget, string string) {
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_test_text_set(arg1, arg2)
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
	var arg1 *C.GtkWidget
	var arg2 C.guint
	var arg3 C.GdkModifierType

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.guint(button)
	arg3 = (C.GdkModifierType)(modifiers)

	var cret C.gboolean
	var goret bool

	cret = C.gtk_test_widget_click(arg1, arg2, arg3)

	if cret {
		goret = true
	}

	return goret
}

// TestWidgetSendKey: this function will generate keyboard press and release
// events in the middle of the first GdkWindow found that belongs to @widget.
// For windowless widgets like Button (which returns false from
// gtk_widget_get_has_window()), this will often be an input-only event window.
// For other widgets, this is usually widget->window. Certain caveats should be
// considered when using this function, in particular because the mouse pointer
// is warped to the key press location, see gdk_test_simulate_key() for details.
func TestWidgetSendKey(widget Widget, keyval uint, modifiers gdk.ModifierType) bool {
	var arg1 *C.GtkWidget
	var arg2 C.guint
	var arg3 C.GdkModifierType

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.guint(keyval)
	arg3 = (C.GdkModifierType)(modifiers)

	var cret C.gboolean
	var goret bool

	cret = C.gtk_test_widget_send_key(arg1, arg2, arg3)

	if cret {
		goret = true
	}

	return goret
}

// TestWidgetWaitForDraw enters the main loop and waits for @widget to be
// “drawn”. In this context that means it waits for the frame clock of @widget
// to have run a full styling, layout and drawing cycle.
//
// This function is intended to be used for syncing with actions that depend on
// @widget relayouting or on interaction with the display server.
func TestWidgetWaitForDraw(widget Widget) {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_test_widget_wait_for_draw(arg1)
}
