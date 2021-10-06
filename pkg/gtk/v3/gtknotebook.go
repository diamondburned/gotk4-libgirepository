// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_notebook_get_type()), F: marshalNotebooker},
	})
}

// NotebookOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type NotebookOverrider interface {
	ChangeCurrentPage(offset int) bool
	FocusTab(typ NotebookTab) bool
	InsertPage(child, tabLabel, menuLabel Widgetter, position int) int
	MoveFocusOut(direction DirectionType)
	PageAdded(child Widgetter, pageNum uint)
	PageRemoved(child Widgetter, pageNum uint)
	PageReordered(child Widgetter, pageNum uint)
	ReorderTab(direction DirectionType, moveToLast bool) bool
	SelectPage(moveFocus bool) bool
	SwitchPage(page Widgetter, pageNum uint)
}

// Notebook widget is a Container whose children are pages that can be switched
// between using tab labels along one edge.
//
// There are many configuration options for GtkNotebook. Among other things, you
// can choose on which edge the tabs appear (see gtk_notebook_set_tab_pos()),
// whether, if there are too many tabs to fit the notebook should be made bigger
// or scrolling arrows added (see gtk_notebook_set_scrollable()), and whether
// there will be a popup menu allowing the users to switch pages. (see
// gtk_notebook_popup_enable(), gtk_notebook_popup_disable())
//
//
// GtkNotebook as GtkBuildable
//
// The GtkNotebook implementation of the Buildable interface supports placing
// children into tabs by specifying “tab” as the “type” attribute of a <child>
// element. Note that the content of the tab must be created before the tab can
// be filled. A tab child can be specified without specifying a <child> type
// attribute.
//
// To add a child widget in the notebooks action area, specify "action-start" or
// “action-end” as the “type” attribute of the <child> element.
//
// An example of a UI definition fragment with GtkNotebook:
//
//    <object class="GtkNotebook">
//      <child>
//        <object class="GtkLabel" id="notebook-content">
//          <property name="label">Content</property>
//        </object>
//      </child>
//      <child type="tab">
//        <object class="GtkLabel" id="notebook-tab">
//          <property name="label">Tab</property>
//        </object>
//      </child>
//    </object>
//
// CSS nodes
//
//    notebook
//    ├── header.top
//    │   ├── [<action widget>]
//    │   ├── tabs
//    │   │   ├── [arrow]
//    │   │   ├── tab
//    │   │   │   ╰── <tab label>
//    ┊   ┊   ┊
//    │   │   ├── tab[.reorderable-page]
//    │   │   │   ╰── <tab label>
//    │   │   ╰── [arrow]
//    │   ╰── [<action widget>]
//    │
//    ╰── stack
//        ├── <child>
//        ┊
//        ╰── <child>
//
// GtkNotebook has a main CSS node with name notebook, a subnode with name
// header and below that a subnode with name tabs which contains one subnode per
// tab with name tab.
//
// If action widgets are present, their CSS nodes are placed next to the tabs
// node. If the notebook is scrollable, CSS nodes with name arrow are placed as
// first and last child of the tabs node.
//
// The main node gets the .frame style class when the notebook has a border (see
// gtk_notebook_set_show_border()).
//
// The header node gets one of the style class .top, .bottom, .left or .right,
// depending on where the tabs are placed. For reorderable pages, the tab node
// gets the .reorderable-page class.
//
// A tab node gets the .dnd style class while it is moved with drag-and-drop.
//
// The nodes are always arranged from left-to-right, regarldess of text
// direction.
type Notebook struct {
	Container
}

func wrapNotebook(obj *externglib.Object) *Notebook {
	return &Notebook{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalNotebooker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNotebook(obj), nil
}

// NewNotebook creates a new Notebook widget with no pages.
func NewNotebook() *Notebook {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_notebook_new()

	var _notebook *Notebook // out

	_notebook = wrapNotebook(externglib.Take(unsafe.Pointer(_cret)))

	return _notebook
}

// AppendPage appends a page to notebook.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//
func (notebook *Notebook) AppendPage(child, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}

	_cret = C.gtk_notebook_append_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// AppendPageMenu appends a page to notebook, specifying the widget to use as
// the label in the popup menu.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//    - menuLabel: widget to use as a label for the page-switch menu, if that
//    is enabled. If NULL, and tab_label is a Label or NULL, then the menu
//    label will be a newly created label with the same text as tab_label; if
//    tab_label is not a Label, menu_label must be specified if the page-switch
//    menu is to be used.
//
func (notebook *Notebook) AppendPageMenu(child, tabLabel, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}
	if menuLabel != nil {
		_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	}

	_cret = C.gtk_notebook_append_page_menu(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)
	runtime.KeepAlive(menuLabel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DetachTab removes the child from the notebook.
//
// This function is very similar to gtk_container_remove(), but additionally
// informs the notebook that the removal is happening as part of a tab DND
// operation, which should not be cancelled.
//
// The function takes the following parameters:
//
//    - child: child.
//
func (notebook *Notebook) DetachTab(child Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_notebook_detach_tab(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
}

// ActionWidget gets one of the action widgets. See
// gtk_notebook_set_action_widget().
//
// The function takes the following parameters:
//
//    - packType: pack type of the action widget to receive.
//
func (notebook *Notebook) ActionWidget(packType PackType) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.GtkPackType  // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.GtkPackType(packType)

	_cret = C.gtk_notebook_get_action_widget(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(packType)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// CurrentPage returns the page number of the current page.
func (notebook *Notebook) CurrentPage() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_current_page(_arg0)
	runtime.KeepAlive(notebook)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GroupName gets the current group name for notebook.
func (notebook *Notebook) GroupName() string {
	var _arg0 *C.GtkNotebook // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_group_name(_arg0)
	runtime.KeepAlive(notebook)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MenuLabel retrieves the menu label widget of the page containing child.
//
// The function takes the following parameters:
//
//    - child: widget contained in a page of notebook.
//
func (notebook *Notebook) MenuLabel(child Widgetter) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// MenuLabelText retrieves the text of the menu label for the page containing
// child.
//
// The function takes the following parameters:
//
//    - child widget of a page of the notebook.
//
func (notebook *Notebook) MenuLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label_text(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NPages gets the number of pages in a notebook.
func (notebook *Notebook) NPages() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_n_pages(_arg0)
	runtime.KeepAlive(notebook)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number page_num.
//
// The function takes the following parameters:
//
//    - pageNum: index of a page in the notebook, or -1 to get the last page.
//
func (notebook *Notebook) NthPage(pageNum int) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_notebook_get_nth_page(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(pageNum)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Scrollable returns whether the tab label area has arrows for scrolling. See
// gtk_notebook_set_scrollable().
func (notebook *Notebook) Scrollable() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_scrollable(_arg0)
	runtime.KeepAlive(notebook)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowBorder returns whether a bevel will be drawn around the notebook pages.
// See gtk_notebook_set_show_border().
func (notebook *Notebook) ShowBorder() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_show_border(_arg0)
	runtime.KeepAlive(notebook)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowTabs returns whether the tabs of the notebook are shown. See
// gtk_notebook_set_show_tabs().
func (notebook *Notebook) ShowTabs() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_show_tabs(_arg0)
	runtime.KeepAlive(notebook)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabDetachable returns whether the tab contents can be detached from notebook.
//
// The function takes the following parameters:
//
//    - child Widget.
//
func (notebook *Notebook) TabDetachable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_detachable(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabHborder returns the horizontal width of a tab border.
//
// Deprecated: this function returns zero.
func (notebook *Notebook) TabHborder() uint16 {
	var _arg0 *C.GtkNotebook // out
	var _cret C.guint16      // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_tab_hborder(_arg0)
	runtime.KeepAlive(notebook)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// TabLabel returns the tab label widget for the page child. NULL is returned if
// child is not in notebook or if no tab label has specifically been set for
// child.
//
// The function takes the following parameters:
//
//    - child: page.
//
func (notebook *Notebook) TabLabel(child Widgetter) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// TabLabelText retrieves the text of the tab label for the page containing
// child.
//
// The function takes the following parameters:
//
//    - child: widget contained in a page of notebook.
//
func (notebook *Notebook) TabLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label_text(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TabPos gets the edge at which the tabs for switching pages in the notebook
// are drawn.
func (notebook *Notebook) TabPos() PositionType {
	var _arg0 *C.GtkNotebook    // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_tab_pos(_arg0)
	runtime.KeepAlive(notebook)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// TabReorderable gets whether the tab can be reordered via drag and drop or
// not.
//
// The function takes the following parameters:
//
//    - child Widget.
//
func (notebook *Notebook) TabReorderable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_reorderable(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabVborder returns the vertical width of a tab border.
//
// Deprecated: this function returns zero.
func (notebook *Notebook) TabVborder() uint16 {
	var _arg0 *C.GtkNotebook // out
	var _cret C.guint16      // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_tab_vborder(_arg0)
	runtime.KeepAlive(notebook)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// InsertPage: insert a page into notebook at the given position.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//    - position: index (starting at 0) at which to insert the page, or -1 to
//    append the page after all other pages.
//
func (notebook *Notebook) InsertPage(child, tabLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}
	_arg3 = C.gint(position)

	_cret = C.gtk_notebook_insert_page(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)
	runtime.KeepAlive(position)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertPageMenu: insert a page into notebook at the given position, specifying
// the widget to use as the label in the popup menu.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//    - menuLabel: widget to use as a label for the page-switch menu, if that
//    is enabled. If NULL, and tab_label is a Label or NULL, then the menu
//    label will be a newly created label with the same text as tab_label; if
//    tab_label is not a Label, menu_label must be specified if the page-switch
//    menu is to be used.
//    - position: index (starting at 0) at which to insert the page, or -1 to
//    append the page after all other pages.
//
func (notebook *Notebook) InsertPageMenu(child, tabLabel, menuLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _arg4 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}
	if menuLabel != nil {
		_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	}
	_arg4 = C.gint(position)

	_cret = C.gtk_notebook_insert_page_menu(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)
	runtime.KeepAlive(menuLabel)
	runtime.KeepAlive(position)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NextPage switches to the next page. Nothing happens if the current page is
// the last page.
func (notebook *Notebook) NextPage() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_next_page(_arg0)
	runtime.KeepAlive(notebook)
}

// PageNum finds the index of the page which contains the given child widget.
//
// The function takes the following parameters:
//
//    - child: Widget.
//
func (notebook *Notebook) PageNum(child Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_page_num(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PopupDisable disables the popup menu.
func (notebook *Notebook) PopupDisable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_disable(_arg0)
	runtime.KeepAlive(notebook)
}

// PopupEnable enables the popup menu: if the user clicks with the right mouse
// button on the tab labels, a menu with all the pages will be popped up.
func (notebook *Notebook) PopupEnable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_enable(_arg0)
	runtime.KeepAlive(notebook)
}

// PrependPage prepends a page to notebook.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//
func (notebook *Notebook) PrependPage(child, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}

	_cret = C.gtk_notebook_prepend_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrependPageMenu prepends a page to notebook, specifying the widget to use as
// the label in the popup menu.
//
// The function takes the following parameters:
//
//    - child to use as the contents of the page.
//    - tabLabel to be used as the label for the page, or NULL to use the
//    default label, “page N”.
//    - menuLabel: widget to use as a label for the page-switch menu, if that
//    is enabled. If NULL, and tab_label is a Label or NULL, then the menu
//    label will be a newly created label with the same text as tab_label; if
//    tab_label is not a Label, menu_label must be specified if the page-switch
//    menu is to be used.
//
func (notebook *Notebook) PrependPageMenu(child, tabLabel, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}
	if menuLabel != nil {
		_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	}

	_cret = C.gtk_notebook_prepend_page_menu(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)
	runtime.KeepAlive(menuLabel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrevPage switches to the previous page. Nothing happens if the current page
// is the first page.
func (notebook *Notebook) PrevPage() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_prev_page(_arg0)
	runtime.KeepAlive(notebook)
}

// RemovePage removes a page from the notebook given its index in the notebook.
//
// The function takes the following parameters:
//
//    - pageNum: index of a notebook page, starting from 0. If -1, the last
//    page will be removed.
//
func (notebook *Notebook) RemovePage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_remove_page(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(pageNum)
}

// ReorderChild reorders the page containing child, so that it appears in
// position position. If position is greater than or equal to the number of
// children in the list or negative, child will be moved to the end of the list.
//
// The function takes the following parameters:
//
//    - child to move.
//    - position: new position, or -1 to move to the end.
//
func (notebook *Notebook) ReorderChild(child Widgetter, position int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_notebook_reorder_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// SetActionWidget sets widget as one of the action widgets. Depending on the
// pack type the widget will be placed before or after the tabs. You can use a
// Box if you need to pack more than one widget on the same side.
//
// Note that action widgets are “internal” children of the notebook and thus not
// included in the list returned from gtk_container_foreach().
//
// The function takes the following parameters:
//
//    - widget: Widget.
//    - packType: pack type of the action widget.
//
func (notebook *Notebook) SetActionWidget(widget Widgetter, packType PackType) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.GtkPackType  // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkPackType(packType)

	C.gtk_notebook_set_action_widget(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(packType)
}

// SetCurrentPage switches to the page number page_num.
//
// Note that due to historical reasons, GtkNotebook refuses to switch to a page
// unless the child widget is visible. Therefore, it is recommended to show
// child widgets before adding them to a notebook.
//
// The function takes the following parameters:
//
//    - pageNum: index of the page to switch to, starting from 0. If negative,
//    the last page will be used. If greater than the number of pages in the
//    notebook, nothing will be done.
//
func (notebook *Notebook) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_set_current_page(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(pageNum)
}

// SetGroupName sets a group name for notebook.
//
// Notebooks with the same name will be able to exchange tabs via drag and drop.
// A notebook with a NULL group name will not be able to exchange tabs with any
// other notebook.
//
// The function takes the following parameters:
//
//    - groupName: name of the notebook group, or NULL to unset it.
//
func (notebook *Notebook) SetGroupName(groupName string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if groupName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_notebook_set_group_name(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(groupName)
}

// SetMenuLabel changes the menu label for the page containing child.
//
// The function takes the following parameters:
//
//    - child widget.
//    - menuLabel: menu label, or NULL for default.
//
func (notebook *Notebook) SetMenuLabel(child, menuLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if menuLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	}

	C.gtk_notebook_set_menu_label(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(menuLabel)
}

// SetMenuLabelText creates a new label and sets it as the menu label of child.
//
// The function takes the following parameters:
//
//    - child widget.
//    - menuText: label text.
//
func (notebook *Notebook) SetMenuLabelText(child Widgetter, menuText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(menuText)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_menu_label_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(menuText)
}

// SetScrollable sets whether the tab label area will have arrows for scrolling
// if there are too many tabs to fit in the area.
//
// The function takes the following parameters:
//
//    - scrollable: TRUE if scroll arrows should be added.
//
func (notebook *Notebook) SetScrollable(scrollable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if scrollable {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_scrollable(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(scrollable)
}

// SetShowBorder sets whether a bevel will be drawn around the notebook pages.
// This only has a visual effect when the tabs are not shown. See
// gtk_notebook_set_show_tabs().
//
// The function takes the following parameters:
//
//    - showBorder: TRUE if a bevel should be drawn around the notebook.
//
func (notebook *Notebook) SetShowBorder(showBorder bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showBorder {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_border(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(showBorder)
}

// SetShowTabs sets whether to show the tabs for the notebook or not.
//
// The function takes the following parameters:
//
//    - showTabs: TRUE if the tabs should be shown.
//
func (notebook *Notebook) SetShowTabs(showTabs bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showTabs {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_tabs(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(showTabs)
}

// SetTabDetachable sets whether the tab can be detached from notebook to
// another notebook or widget.
//
// Note that 2 notebooks must share a common group identificator (see
// gtk_notebook_set_group_name()) to allow automatic tabs interchange between
// them.
//
// If you want a widget to interact with a notebook through DnD (i.e.: accept
// dragged tabs from it) it must be set as a drop destination and accept the
// target “GTK_NOTEBOOK_TAB”. The notebook will fill the selection with a
// GtkWidget** pointing to the child widget that corresponds to the dropped tab.
//
// Note that you should use gtk_notebook_detach_tab() instead of
// gtk_container_remove() if you want to remove the tab from the source notebook
// as part of accepting a drop. Otherwise, the source notebook will think that
// the dragged tab was removed from underneath the ongoing drag operation, and
// will initiate a drag cancel animation.
//
//     static void
//     on_drag_data_received (GtkWidget        *widget,
//                            GdkDragContext   *context,
//                            gint              x,
//                            gint              y,
//                            GtkSelectionData *data,
//                            guint             info,
//                            guint             time,
//                            gpointer          user_data)
//     {
//       GtkWidget *notebook;
//       GtkWidget **child;
//
//       notebook = gtk_drag_get_source_widget (context);
//       child = (void*) gtk_selection_data_get_data (data);
//
//       // process_widget (*child);
//
//       gtk_notebook_detach_tab (GTK_NOTEBOOK (notebook), *child);
//     }
//
// If you want a notebook to accept drags from other widgets, you will have to
// set your own DnD code to do it.
//
// The function takes the following parameters:
//
//    - child Widget.
//    - detachable: whether the tab is detachable or not.
//
func (notebook *Notebook) SetTabDetachable(child Widgetter, detachable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if detachable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_detachable(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(detachable)
}

// SetTabLabel changes the tab label for child. If NULL is specified for
// tab_label, then the page will have the label “page N”.
//
// The function takes the following parameters:
//
//    - child: page.
//    - tabLabel: tab label widget to use, or NULL for default tab label.
//
func (notebook *Notebook) SetTabLabel(child, tabLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if tabLabel != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	}

	C.gtk_notebook_set_tab_label(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabLabel)
}

// SetTabLabelText creates a new label and sets it as the tab label for the page
// containing child.
//
// The function takes the following parameters:
//
//    - child: page.
//    - tabText: label text.
//
func (notebook *Notebook) SetTabLabelText(child Widgetter, tabText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(tabText)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_tab_label_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tabText)
}

// SetTabPos sets the edge at which the tabs for switching pages in the notebook
// are drawn.
//
// The function takes the following parameters:
//
//    - pos: edge to draw the tabs at.
//
func (notebook *Notebook) SetTabPos(pos PositionType) {
	var _arg0 *C.GtkNotebook    // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.GtkPositionType(pos)

	C.gtk_notebook_set_tab_pos(_arg0, _arg1)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(pos)
}

// SetTabReorderable sets whether the notebook tab can be reordered via drag and
// drop or not.
//
// The function takes the following parameters:
//
//    - child Widget.
//    - reorderable: whether the tab is reorderable or not.
//
func (notebook *Notebook) SetTabReorderable(child Widgetter, reorderable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if reorderable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_reorderable(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)
	runtime.KeepAlive(reorderable)
}

func (notebook *Notebook) ConnectChangeCurrentPage(f func(object int) bool) externglib.SignalHandle {
	return notebook.Connect("change-current-page", f)
}

// ConnectCreateWindow signal is emitted when a detachable tab is dropped on the
// root window.
//
// A handler for this signal can create a window containing a notebook where the
// tab will be attached. It is also responsible for moving/resizing the window
// and adding the necessary properties to the notebook (e.g. the
// Notebook:group-name ).
func (notebook *Notebook) ConnectCreateWindow(f func(page Widgetter, x, y int) Notebook) externglib.SignalHandle {
	return notebook.Connect("create-window", f)
}

func (notebook *Notebook) ConnectFocusTab(f func(object NotebookTab) bool) externglib.SignalHandle {
	return notebook.Connect("focus-tab", f)
}

func (notebook *Notebook) ConnectMoveFocusOut(f func(object DirectionType)) externglib.SignalHandle {
	return notebook.Connect("move-focus-out", f)
}

// ConnectPageAdded signal is emitted in the notebook right after a page is
// added to the notebook.
func (notebook *Notebook) ConnectPageAdded(f func(child Widgetter, pageNum uint)) externglib.SignalHandle {
	return notebook.Connect("page-added", f)
}

// ConnectPageRemoved signal is emitted in the notebook right after a page is
// removed from the notebook.
func (notebook *Notebook) ConnectPageRemoved(f func(child Widgetter, pageNum uint)) externglib.SignalHandle {
	return notebook.Connect("page-removed", f)
}

// ConnectPageReordered signal is emitted in the notebook right after a page has
// been reordered.
func (notebook *Notebook) ConnectPageReordered(f func(child Widgetter, pageNum uint)) externglib.SignalHandle {
	return notebook.Connect("page-reordered", f)
}

func (notebook *Notebook) ConnectReorderTab(f func(object DirectionType, p0 bool) bool) externglib.SignalHandle {
	return notebook.Connect("reorder-tab", f)
}

func (notebook *Notebook) ConnectSelectPage(f func(object bool) bool) externglib.SignalHandle {
	return notebook.Connect("select-page", f)
}

// ConnectSwitchPage: emitted when the user or a function changes the current
// page.
func (notebook *Notebook) ConnectSwitchPage(f func(page Widgetter, pageNum uint)) externglib.SignalHandle {
	return notebook.Connect("switch-page", f)
}
