// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_button_get_type()), F: marshalToolButton},
	})
}

// ToolButton ToolButtons are ToolItems containing buttons.
//
// Use gtk_tool_button_new() to create a new ToolButton.
//
// The label of a ToolButton is determined by the properties
// ToolButton:label-widget, ToolButton:label, and ToolButton:stock-id. If
// ToolButton:label-widget is non-nil, then that widget is used as the label.
// Otherwise, if ToolButton:label is non-nil, that string is used as the label.
// Otherwise, if ToolButton:stock-id is non-nil, the label is determined by the
// stock item. Otherwise, the button does not have a label.
//
// The icon of a ToolButton is determined by the properties
// ToolButton:icon-widget and ToolButton:stock-id. If ToolButton:icon-widget is
// non-nil, then that widget is used as the icon. Otherwise, if
// ToolButton:stock-id is non-nil, the icon is determined by the stock item.
// Otherwise, the button does not have a icon.
//
//
// CSS nodes
//
// GtkToolButton has a single CSS node with name toolbutton.
type ToolButton interface {
	ToolItem
	Actionable
	Activatable
	Buildable

	// IconName returns the name of the themed icon for the tool button, see
	// gtk_tool_button_set_icon_name().
	IconName() string
	// IconWidget: return the widget used as icon widget on @button. See
	// gtk_tool_button_set_icon_widget().
	IconWidget() Widget
	// Label returns the label used by the tool button, or nil if the tool
	// button doesn’t have a label. or uses a the label from a stock item. The
	// returned string is owned by GTK+, and must not be modified or freed.
	Label() string
	// LabelWidget returns the widget used as label on @button. See
	// gtk_tool_button_set_label_widget().
	LabelWidget() Widget
	// StockID returns the name of the stock item. See
	// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and
	// must not be freed or modifed.
	StockID() string
	// UseUnderline returns whether underscores in the label property are used
	// as mnemonics on menu items on the overflow menu. See
	// gtk_tool_button_set_use_underline().
	UseUnderline() bool
	// SetIconName sets the icon for the tool button from a named themed icon.
	// See the docs for IconTheme for more details. The ToolButton:icon-name
	// property only has an effect if not overridden by non-nil
	// ToolButton:label-widget, ToolButton:icon-widget and ToolButton:stock-id
	// properties.
	SetIconName(iconName string)
	// SetIconWidget sets @icon as the widget used as icon on @button. If
	// @icon_widget is nil the icon is determined by the ToolButton:stock-id
	// property. If the ToolButton:stock-id property is also nil, @button will
	// not have an icon.
	SetIconWidget(iconWidget Widget)
	// SetLabel sets @label as the label used for the tool button. The
	// ToolButton:label property only has an effect if not overridden by a
	// non-nil ToolButton:label-widget property. If both the
	// ToolButton:label-widget and ToolButton:label properties are nil, the
	// label is determined by the ToolButton:stock-id property. If the
	// ToolButton:stock-id property is also nil, @button will not have a label.
	SetLabel(label string)
	// SetLabelWidget sets @label_widget as the widget that will be used as the
	// label for @button. If @label_widget is nil the ToolButton:label property
	// is used as label. If ToolButton:label is also nil, the label in the stock
	// item determined by the ToolButton:stock-id property is used as label. If
	// ToolButton:stock-id is also nil, @button does not have a label.
	SetLabelWidget(labelWidget Widget)
	// SetStockID sets the name of the stock item. See
	// gtk_tool_button_new_from_stock(). The stock_id property only has an
	// effect if not overridden by non-nil ToolButton:label-widget and
	// ToolButton:icon-widget properties.
	SetStockID(stockId string)
	// SetUseUnderline: if set, an underline in the label property indicates
	// that the next character should be used for the mnemonic accelerator key
	// in the overflow menu. For example, if the label property is “_Open” and
	// @use_underline is true, the label on the tool button will be “Open” and
	// the item on the overflow menu will have an underlined “O”.
	//
	// Labels shown on tool buttons never have mnemonics on them; this property
	// only affects the menu item on the overflow menu.
	SetUseUnderline(useUnderline bool)
}

// toolButton implements the ToolButton interface.
type toolButton struct {
	ToolItem
	Actionable
	Activatable
	Buildable
}

var _ ToolButton = (*toolButton)(nil)

// WrapToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolButton(obj *externglib.Object) ToolButton {
	return ToolButton{
		ToolItem:    WrapToolItem(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolButton(obj), nil
}

// NewToolButton constructs a class ToolButton.
func NewToolButton(iconWidget Widget, label string) ToolButton {
	var _arg1 *C.GtkWidget
	var _arg2 *C.gchar

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.GtkToolButton

	cret = C.gtk_tool_button_new(_arg1, _arg2)

	var _toolButton ToolButton

	_toolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToolButton)

	return _toolButton
}

// NewToolButtonFromStock constructs a class ToolButton.
func NewToolButtonFromStock(stockId string) ToolButton {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkToolButton

	cret = C.gtk_tool_button_new_from_stock(_arg1)

	var _toolButton ToolButton

	_toolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToolButton)

	return _toolButton
}

// IconName returns the name of the themed icon for the tool button, see
// gtk_tool_button_set_icon_name().
func (b toolButton) IconName() string {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar

	cret = C.gtk_tool_button_get_icon_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IconWidget: return the widget used as icon widget on @button. See
// gtk_tool_button_set_icon_widget().
func (b toolButton) IconWidget() Widget {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_tool_button_get_icon_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Label returns the label used by the tool button, or nil if the tool
// button doesn’t have a label. or uses a the label from a stock item. The
// returned string is owned by GTK+, and must not be modified or freed.
func (b toolButton) Label() string {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar

	cret = C.gtk_tool_button_get_label(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// LabelWidget returns the widget used as label on @button. See
// gtk_tool_button_set_label_widget().
func (b toolButton) LabelWidget() Widget {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_tool_button_get_label_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// StockID returns the name of the stock item. See
// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and
// must not be freed or modifed.
func (b toolButton) StockID() string {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar

	cret = C.gtk_tool_button_get_stock_id(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UseUnderline returns whether underscores in the label property are used
// as mnemonics on menu items on the overflow menu. See
// gtk_tool_button_set_use_underline().
func (b toolButton) UseUnderline() bool {
	var _arg0 *C.GtkToolButton

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_tool_button_get_use_underline(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon for the tool button from a named themed icon.
// See the docs for IconTheme for more details. The ToolButton:icon-name
// property only has an effect if not overridden by non-nil
// ToolButton:label-widget, ToolButton:icon-widget and ToolButton:stock-id
// properties.
func (b toolButton) SetIconName(iconName string) {
	var _arg0 *C.GtkToolButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_icon_name(_arg0, _arg1)
}

// SetIconWidget sets @icon as the widget used as icon on @button. If
// @icon_widget is nil the icon is determined by the ToolButton:stock-id
// property. If the ToolButton:stock-id property is also nil, @button will
// not have an icon.
func (b toolButton) SetIconWidget(iconWidget Widget) {
	var _arg0 *C.GtkToolButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))

	C.gtk_tool_button_set_icon_widget(_arg0, _arg1)
}

// SetLabel sets @label as the label used for the tool button. The
// ToolButton:label property only has an effect if not overridden by a
// non-nil ToolButton:label-widget property. If both the
// ToolButton:label-widget and ToolButton:label properties are nil, the
// label is determined by the ToolButton:stock-id property. If the
// ToolButton:stock-id property is also nil, @button will not have a label.
func (b toolButton) SetLabel(label string) {
	var _arg0 *C.GtkToolButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_label(_arg0, _arg1)
}

// SetLabelWidget sets @label_widget as the widget that will be used as the
// label for @button. If @label_widget is nil the ToolButton:label property
// is used as label. If ToolButton:label is also nil, the label in the stock
// item determined by the ToolButton:stock-id property is used as label. If
// ToolButton:stock-id is also nil, @button does not have a label.
func (b toolButton) SetLabelWidget(labelWidget Widget) {
	var _arg0 *C.GtkToolButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_button_set_label_widget(_arg0, _arg1)
}

// SetStockID sets the name of the stock item. See
// gtk_tool_button_new_from_stock(). The stock_id property only has an
// effect if not overridden by non-nil ToolButton:label-widget and
// ToolButton:icon-widget properties.
func (b toolButton) SetStockID(stockId string) {
	var _arg0 *C.GtkToolButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_stock_id(_arg0, _arg1)
}

// SetUseUnderline: if set, an underline in the label property indicates
// that the next character should be used for the mnemonic accelerator key
// in the overflow menu. For example, if the label property is “_Open” and
// @use_underline is true, the label on the tool button will be “Open” and
// the item on the overflow menu will have an underlined “O”.
//
// Labels shown on tool buttons never have mnemonics on them; this property
// only affects the menu item on the overflow menu.
func (b toolButton) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkToolButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	if useUnderline {
		_arg1 = C.gboolean(1)
	}

	C.gtk_tool_button_set_use_underline(_arg0, _arg1)
}
