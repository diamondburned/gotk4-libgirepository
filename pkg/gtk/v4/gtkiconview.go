// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_IconView_ConnectUnselectAll(gpointer, guintptr);
// extern void _gotk4_gtk4_IconView_ConnectToggleCursorItem(gpointer, guintptr);
// extern void _gotk4_gtk4_IconView_ConnectSelectionChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_IconView_ConnectSelectCursorItem(gpointer, guintptr);
// extern void _gotk4_gtk4_IconView_ConnectSelectAll(gpointer, guintptr);
// extern void _gotk4_gtk4_IconView_ConnectItemActivated(gpointer, void*, guintptr);
// extern gboolean _gotk4_gtk4_IconView_ConnectActivateCursorItem(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeIconViewDropPosition = coreglib.Type(girepository.MustFind("Gtk", "IconViewDropPosition").RegisteredGType())
	GTypeIconView             = coreglib.Type(girepository.MustFind("Gtk", "IconView").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIconViewDropPosition, F: marshalIconViewDropPosition},
		coreglib.TypeMarshaler{T: GTypeIconView, F: marshalIconView},
	})
}

// IconViewDropPosition: enum for determining where a dropped item goes.
type IconViewDropPosition C.gint

const (
	// IconViewNoDrop: no drop possible.
	IconViewNoDrop IconViewDropPosition = iota
	// IconViewDropInto: dropped item replaces the item.
	IconViewDropInto
	// IconViewDropLeft: dropped item is inserted to the left.
	IconViewDropLeft
	// IconViewDropRight: dropped item is inserted to the right.
	IconViewDropRight
	// IconViewDropAbove: dropped item is inserted above.
	IconViewDropAbove
	// IconViewDropBelow: dropped item is inserted below.
	IconViewDropBelow
)

func marshalIconViewDropPosition(p uintptr) (interface{}, error) {
	return IconViewDropPosition(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for IconViewDropPosition.
func (i IconViewDropPosition) String() string {
	switch i {
	case IconViewNoDrop:
		return "NoDrop"
	case IconViewDropInto:
		return "DropInto"
	case IconViewDropLeft:
		return "DropLeft"
	case IconViewDropRight:
		return "DropRight"
	case IconViewDropAbove:
		return "DropAbove"
	case IconViewDropBelow:
		return "DropBelow"
	default:
		return fmt.Sprintf("IconViewDropPosition(%d)", i)
	}
}

// IconViewForEachFunc: function used by gtk_icon_view_selected_foreach() to map
// all selected rows.
//
// It will be called on every selected row in the view.
type IconViewForEachFunc func(iconView *IconView, path *TreePath)

// IconView: GtkIconView is a widget which displays data in a grid of icons.
//
// GtkIconView provides an alternative view on a GtkTreeModel. It displays the
// model as a grid of icons with labels. Like gtk.TreeView, it allows to select
// one or multiple items (depending on the selection mode, see
// gtk.IconView.SetSelectionMode()). In addition to selection with the arrow
// keys, GtkIconView supports rubberband selection, which is controlled by
// dragging the pointer.
//
// Note that if the tree model is backed by an actual tree store (as opposed to
// a flat list where the mapping to icons is obvious), IconView will only
// display the first level of the tree and ignore the tree’s branches.
//
// CSS nodes
//
//    iconview.view
//    ╰── [rubberband]
//
//
// GtkIconView has a single CSS node with name iconview and style class .view.
// For rubberband selection, a subnode with name rubberband is used.
type IconView struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	CellLayout
	Scrollable
}

var (
	_ Widgetter         = (*IconView)(nil)
	_ coreglib.Objector = (*IconView)(nil)
)

func wrapIconView(obj *coreglib.Object) *IconView {
	return &IconView{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		CellLayout: CellLayout{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalIconView(p uintptr) (interface{}, error) {
	return wrapIconView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivateCursorItem: [keybinding signal][GtkSignalAction] which gets
// emitted when the user activates the currently focused item.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control activation programmatically.
//
// The default bindings for this signal are Space, Return and Enter.
func (v *IconView) ConnectActivateCursorItem(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate-cursor-item", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectActivateCursorItem), f)
}

// ConnectItemActivated signal is emitted when the method
// gtk_icon_view_item_activated() is called, when the user double clicks an item
// with the "activate-on-single-click" property set to FALSE, or when the user
// single clicks an item when the "activate-on-single-click" property set to
// TRUE. It is also emitted when a non-editable item is selected and one of the
// keys: Space, Return or Enter is pressed.
func (v *IconView) ConnectItemActivated(f func(path *TreePath)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "item-activated", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectItemActivated), f)
}

// ConnectSelectAll: [keybinding signal][GtkSignalAction] which gets emitted
// when the user selects all items.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control selection programmatically.
//
// The default binding for this signal is Ctrl-a.
func (v *IconView) ConnectSelectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "select-all", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectSelectAll), f)
}

// ConnectSelectCursorItem: [keybinding signal][GtkSignalAction] which gets
// emitted when the user selects the item that is currently focused.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control selection programmatically.
//
// There is no default binding for this signal.
func (v *IconView) ConnectSelectCursorItem(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "select-cursor-item", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectSelectCursorItem), f)
}

// ConnectSelectionChanged signal is emitted when the selection (i.e. the set of
// selected items) changes.
func (v *IconView) ConnectSelectionChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "selection-changed", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectSelectionChanged), f)
}

// ConnectToggleCursorItem: [keybinding signal][GtkSignalAction] which gets
// emitted when the user toggles whether the currently focused item is selected
// or not. The exact effect of this depend on the selection mode.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control selection programmatically.
//
// There is no default binding for this signal is Ctrl-Space.
func (v *IconView) ConnectToggleCursorItem(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "toggle-cursor-item", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectToggleCursorItem), f)
}

// ConnectUnselectAll: [keybinding signal][GtkSignalAction] which gets emitted
// when the user unselects all items.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control selection programmatically.
//
// The default binding for this signal is Ctrl-Shift-a.
func (v *IconView) ConnectUnselectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "unselect-all", false, unsafe.Pointer(C._gotk4_gtk4_IconView_ConnectUnselectAll), f)
}
