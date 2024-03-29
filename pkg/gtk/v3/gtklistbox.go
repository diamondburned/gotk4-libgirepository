// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ListBox_ConnectUnselectAll(gpointer, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectToggleCursorRow(gpointer, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectSelectedRowsChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectSelectAll(gpointer, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectRowSelected(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectRowActivated(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_ListBox_ConnectActivateCursorRow(gpointer, guintptr);
// extern void _gotk4_gtk3_ListBoxRow_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeListBox    = coreglib.Type(girepository.MustFind("Gtk", "ListBox").RegisteredGType())
	GTypeListBoxRow = coreglib.Type(girepository.MustFind("Gtk", "ListBoxRow").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeListBox, F: marshalListBox},
		coreglib.TypeMarshaler{T: GTypeListBoxRow, F: marshalListBoxRow},
	})
}

// ListBoxOverrides contains methods that are overridable.
type ListBoxOverrides struct {
}

func defaultListBoxOverrides(v *ListBox) ListBoxOverrides {
	return ListBoxOverrides{}
}

// ListBox is a vertical container that contains GtkListBoxRow children. These
// rows can by dynamically sorted and filtered, and headers can be added
// dynamically depending on the row content. It also allows keyboard and mouse
// navigation and selection like a typical list.
//
// Using GtkListBox is often an alternative to TreeView, especially when the
// list contents has a more complicated layout than what is allowed by a
// CellRenderer, or when the contents is interactive (i.e. has a button in it).
//
// Although a ListBox must have only ListBoxRow children you can add any kind of
// widget to it via gtk_container_add(), and a ListBoxRow widget will
// automatically be inserted between the list and the widget.
//
// ListBoxRows can be marked as activatable or selectable. If a row is
// activatable, ListBox::row-activated will be emitted for it when the user
// tries to activate it. If it is selectable, the row will be marked as selected
// when the user tries to select it.
//
// The GtkListBox widget was added in GTK+ 3.10.
//
//
// GtkListBox as GtkBuildable
//
// The GtkListBox implementation of the Buildable interface supports setting a
// child as the placeholder by specifying “placeholder” as the “type” attribute
// of a <child> element. See gtk_list_box_set_placeholder() for info.
//
// CSS nodes
//
//    list
//    ╰── row[.activatable]
//
// GtkListBox uses a single CSS node named list. Each GtkListBoxRow uses a
// single CSS node named row. The row nodes get the .activatable style class
// added when appropriate.
type ListBox struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*ListBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ListBox, *ListBoxClass, ListBoxOverrides](
		GTypeListBox,
		initListBoxClass,
		wrapListBox,
		defaultListBoxOverrides,
	)
}

func initListBoxClass(gclass unsafe.Pointer, overrides ListBoxOverrides, classInitFunc func(*ListBoxClass)) {
	if classInitFunc != nil {
		class := (*ListBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapListBox(obj *coreglib.Object) *ListBox {
	return &ListBox{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
		},
	}
}

func marshalListBox(p uintptr) (interface{}, error) {
	return wrapListBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ListBox) ConnectActivateCursorRow(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate-cursor-row", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectActivateCursorRow), f)
}

// ConnectRowActivated signal is emitted when a row has been activated by the
// user.
func (v *ListBox) ConnectRowActivated(f func(row *ListBoxRow)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "row-activated", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectRowActivated), f)
}

// ConnectRowSelected signal is emitted when a new row is selected, or (with a
// NULL row) when the selection is cleared.
//
// When the box is using K_SELECTION_MULTIPLE, this signal will not give you the
// full picture of selection changes, and you should use the
// ListBox::selected-rows-changed signal instead.
func (v *ListBox) ConnectRowSelected(f func(row *ListBoxRow)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "row-selected", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectRowSelected), f)
}

// ConnectSelectAll signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to select all children of the box, if the selection mode permits it.
//
// The default bindings for this signal is Ctrl-a.
func (v *ListBox) ConnectSelectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "select-all", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectSelectAll), f)
}

// ConnectSelectedRowsChanged signal is emitted when the set of selected rows
// changes.
func (v *ListBox) ConnectSelectedRowsChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "selected-rows-changed", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectSelectedRowsChanged), f)
}

func (v *ListBox) ConnectToggleCursorRow(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "toggle-cursor-row", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectToggleCursorRow), f)
}

// ConnectUnselectAll signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to unselect all children of the box, if the selection mode
// permits it.
//
// The default bindings for this signal is Ctrl-Shift-a.
func (v *ListBox) ConnectUnselectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "unselect-all", false, unsafe.Pointer(C._gotk4_gtk3_ListBox_ConnectUnselectAll), f)
}

// ListBoxRowOverrides contains methods that are overridable.
type ListBoxRowOverrides struct {
}

func defaultListBoxRowOverrides(v *ListBoxRow) ListBoxRowOverrides {
	return ListBoxRowOverrides{}
}

type ListBoxRow struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Actionable
}

var (
	_ Binner            = (*ListBoxRow)(nil)
	_ coreglib.Objector = (*ListBoxRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ListBoxRow, *ListBoxRowClass, ListBoxRowOverrides](
		GTypeListBoxRow,
		initListBoxRowClass,
		wrapListBoxRow,
		defaultListBoxRowOverrides,
	)
}

func initListBoxRowClass(gclass unsafe.Pointer, overrides ListBoxRowOverrides, classInitFunc func(*ListBoxRowClass)) {
	if classInitFunc != nil {
		class := (*ListBoxRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapListBoxRow(obj *coreglib.Object) *ListBoxRow {
	return &ListBoxRow{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			},
		},
		Object: obj,
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
		},
	}
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	return wrapListBoxRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate: this is a keybinding signal, which will cause this row to be
// activated.
//
// If you want to be notified when the user activates a row (by key or not), use
// the ListBox::row-activated signal on the row’s parent ListBox.
func (v *ListBoxRow) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate", false, unsafe.Pointer(C._gotk4_gtk3_ListBoxRow_ConnectActivate), f)
}

// ListBoxClass: instance of this type is always passed by reference.
type ListBoxClass struct {
	*listBoxClass
}

// listBoxClass is the struct that's finalized.
type listBoxClass struct {
	native unsafe.Pointer
}

var GIRInfoListBoxClass = girepository.MustFind("Gtk", "ListBoxClass")

// ListBoxRowClass: instance of this type is always passed by reference.
type ListBoxRowClass struct {
	*listBoxRowClass
}

// listBoxRowClass is the struct that's finalized.
type listBoxRowClass struct {
	native unsafe.Pointer
}

var GIRInfoListBoxRowClass = girepository.MustFind("Gtk", "ListBoxRowClass")
