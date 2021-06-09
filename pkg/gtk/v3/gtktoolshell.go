// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_tool_shell_get_type()), F: marshalToolShell},
	})
}

// ToolShellOverrider contains methods that are overridable. This
// interface is a subset of the interface ToolShell.
type ToolShellOverrider interface {
	// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_ellipsize_mode() instead.
	EllipsizeMode() pango.EllipsizeMode

	IconSize() IconSize
	// Orientation retrieves the current orientation for the tool shell. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_orientation() instead.
	Orientation() Orientation
	// ReliefStyle returns the relief style of buttons on @shell. Tool items
	// must not call this function directly, but rely on
	// gtk_tool_item_get_relief_style() instead.
	ReliefStyle() ReliefStyle
	// Style retrieves whether the tool shell has text, icons, or both. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_toolbar_style() instead.
	Style() ToolbarStyle
	// TextAlignment retrieves the current text alignment for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_alignment() instead.
	TextAlignment() float32
	// TextOrientation retrieves the current text orientation for the tool
	// shell. Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_orientation() instead.
	TextOrientation() Orientation
	// TextSizeGroup retrieves the current text size group for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_size_group() instead.
	TextSizeGroup() SizeGroup
	// RebuildMenu: calling this function signals the tool shell that the
	// overflow menu item for tool items have changed. If there is an overflow
	// menu and if it is visible when this function it called, the menu will be
	// rebuilt.
	//
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_rebuild_menu() instead.
	RebuildMenu()
}

// ToolShell: the ToolShell interface allows container widgets to provide
// additional information when embedding ToolItem widgets.
type ToolShell interface {
	Widget
	ToolShellOverrider
}

// toolShell implements the ToolShell interface.
type toolShell struct {
	Widget
}

var _ ToolShell = (*toolShell)(nil)

// WrapToolShell wraps a GObject to a type that implements interface
// ToolShell. It is primarily used internally.
func WrapToolShell(obj *externglib.Object) ToolShell {
	return ToolShell{
		Widget: WrapWidget(obj),
	}
}

func marshalToolShell(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolShell(obj), nil
}

// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
// Tool items must not call this function directly, but rely on
// gtk_tool_item_get_ellipsize_mode() instead.
func (s toolShell) EllipsizeMode() pango.EllipsizeMode {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.PangoEllipsizeMode

	cret = C.gtk_tool_shell_get_ellipsize_mode(_arg0)

	var _ellipsizeMode pango.EllipsizeMode

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// IconSize retrieves the icon size for the tool shell. Tool items must not
// call this function directly, but rely on gtk_tool_item_get_icon_size()
// instead.
func (s toolShell) IconSize() int {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.GtkIconSize

	cret = C.gtk_tool_shell_get_icon_size(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Orientation retrieves the current orientation for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_orientation() instead.
func (s toolShell) Orientation() Orientation {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.GtkOrientation

	cret = C.gtk_tool_shell_get_orientation(_arg0)

	var _orientation Orientation

	_orientation = Orientation(_cret)

	return _orientation
}

// ReliefStyle returns the relief style of buttons on @shell. Tool items
// must not call this function directly, but rely on
// gtk_tool_item_get_relief_style() instead.
func (s toolShell) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.GtkReliefStyle

	cret = C.gtk_tool_shell_get_relief_style(_arg0)

	var _reliefStyle ReliefStyle

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// Style retrieves whether the tool shell has text, icons, or both. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_toolbar_style() instead.
func (s toolShell) Style() ToolbarStyle {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.GtkToolbarStyle

	cret = C.gtk_tool_shell_get_style(_arg0)

	var _toolbarStyle ToolbarStyle

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// TextAlignment retrieves the current text alignment for the tool shell.
// Tool items must not call this function directly, but rely on
// gtk_tool_item_get_text_alignment() instead.
func (s toolShell) TextAlignment() float32 {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.gfloat

	cret = C.gtk_tool_shell_get_text_alignment(_arg0)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// TextOrientation retrieves the current text orientation for the tool
// shell. Tool items must not call this function directly, but rely on
// gtk_tool_item_get_text_orientation() instead.
func (s toolShell) TextOrientation() Orientation {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret C.GtkOrientation

	cret = C.gtk_tool_shell_get_text_orientation(_arg0)

	var _orientation Orientation

	_orientation = Orientation(_cret)

	return _orientation
}

// TextSizeGroup retrieves the current text size group for the tool shell.
// Tool items must not call this function directly, but rely on
// gtk_tool_item_get_text_size_group() instead.
func (s toolShell) TextSizeGroup() SizeGroup {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkSizeGroup

	cret = C.gtk_tool_shell_get_text_size_group(_arg0)

	var _sizeGroup SizeGroup

	_sizeGroup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(SizeGroup)

	return _sizeGroup
}

// RebuildMenu: calling this function signals the tool shell that the
// overflow menu item for tool items have changed. If there is an overflow
// menu and if it is visible when this function it called, the menu will be
// rebuilt.
//
// Tool items must not call this function directly, but rely on
// gtk_tool_item_rebuild_menu() instead.
func (s toolShell) RebuildMenu() {
	var _arg0 *C.GtkToolShell

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(s.Native()))

	C.gtk_tool_shell_rebuild_menu(_arg0)
}