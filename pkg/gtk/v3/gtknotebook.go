// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
	//
	ChangeCurrentPage(offset int) bool
	//
	InsertPage(child Widgetter, tabLabel Widgetter, menuLabel Widgetter, position int) int
	//
	PageAdded(child Widgetter, pageNum uint)
	//
	PageRemoved(child Widgetter, pageNum uint)
	//
	PageReordered(child Widgetter, pageNum uint)
	//
	SelectPage(moveFocus bool) bool
	//
	SwitchPage(page Widgetter, pageNum uint)
}

// Notebooker describes Notebook's methods.
type Notebooker interface {
	// AppendPage appends a page to @notebook.
	AppendPage(child Widgetter, tabLabel Widgetter) int
	// AppendPageMenu appends a page to @notebook, specifying the widget to use
	// as the label in the popup menu.
	AppendPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int
	// DetachTab removes the child from the notebook.
	DetachTab(child Widgetter)
	// CurrentPage returns the page number of the current page.
	CurrentPage() int
	// GroupName gets the current group name for @notebook.
	GroupName() string
	// MenuLabel retrieves the menu label widget of the page containing @child.
	MenuLabel(child Widgetter) *Widget
	// MenuLabelText retrieves the text of the menu label for the page
	// containing @child.
	MenuLabelText(child Widgetter) string
	// NPages gets the number of pages in a notebook.
	NPages() int
	// NthPage returns the child widget contained in page number @page_num.
	NthPage(pageNum int) *Widget
	// Scrollable returns whether the tab label area has arrows for scrolling.
	Scrollable() bool
	// ShowBorder returns whether a bevel will be drawn around the notebook
	// pages.
	ShowBorder() bool
	// ShowTabs returns whether the tabs of the notebook are shown.
	ShowTabs() bool
	// TabDetachable returns whether the tab contents can be detached from
	// @notebook.
	TabDetachable(child Widgetter) bool
	// TabHborder returns the horizontal width of a tab border.
	TabHborder() uint16
	// TabLabel returns the tab label widget for the page @child.
	TabLabel(child Widgetter) *Widget
	// TabLabelText retrieves the text of the tab label for the page containing
	// @child.
	TabLabelText(child Widgetter) string
	// TabPos gets the edge at which the tabs for switching pages in the
	// notebook are drawn.
	TabPos() PositionType
	// TabReorderable gets whether the tab can be reordered via drag and drop or
	// not.
	TabReorderable(child Widgetter) bool
	// TabVborder returns the vertical width of a tab border.
	TabVborder() uint16
	// InsertPage: insert a page into @notebook at the given position.
	InsertPage(child Widgetter, tabLabel Widgetter, position int) int
	// InsertPageMenu: insert a page into @notebook at the given position,
	// specifying the widget to use as the label in the popup menu.
	InsertPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter, position int) int
	// NextPage switches to the next page.
	NextPage()
	// PageNum finds the index of the page which contains the given child
	// widget.
	PageNum(child Widgetter) int
	// PopupDisable disables the popup menu.
	PopupDisable()
	// PopupEnable enables the popup menu: if the user clicks with the right
	// mouse button on the tab labels, a menu with all the pages will be popped
	// up.
	PopupEnable()
	// PrependPage prepends a page to @notebook.
	PrependPage(child Widgetter, tabLabel Widgetter) int
	// PrependPageMenu prepends a page to @notebook, specifying the widget to
	// use as the label in the popup menu.
	PrependPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int
	// PrevPage switches to the previous page.
	PrevPage()
	// RemovePage removes a page from the notebook given its index in the
	// notebook.
	RemovePage(pageNum int)
	// ReorderChild reorders the page containing @child, so that it appears in
	// position @position.
	ReorderChild(child Widgetter, position int)
	// SetCurrentPage switches to the page number @page_num.
	SetCurrentPage(pageNum int)
	// SetGroupName sets a group name for @notebook.
	SetGroupName(groupName string)
	// SetMenuLabel changes the menu label for the page containing @child.
	SetMenuLabel(child Widgetter, menuLabel Widgetter)
	// SetMenuLabelText creates a new label and sets it as the menu label of
	// @child.
	SetMenuLabelText(child Widgetter, menuText string)
	// SetScrollable sets whether the tab label area will have arrows for
	// scrolling if there are too many tabs to fit in the area.
	SetScrollable(scrollable bool)
	// SetShowBorder sets whether a bevel will be drawn around the notebook
	// pages.
	SetShowBorder(showBorder bool)
	// SetShowTabs sets whether to show the tabs for the notebook or not.
	SetShowTabs(showTabs bool)
	// SetTabDetachable sets whether the tab can be detached from @notebook to
	// another notebook or widget.
	SetTabDetachable(child Widgetter, detachable bool)
	// SetTabLabel changes the tab label for @child.
	SetTabLabel(child Widgetter, tabLabel Widgetter)
	// SetTabLabelText creates a new label and sets it as the tab label for the
	// page containing @child.
	SetTabLabelText(child Widgetter, tabText string)
	// SetTabReorderable sets whether the notebook tab can be reordered via drag
	// and drop or not.
	SetTabReorderable(child Widgetter, reorderable bool)
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

var (
	_ Notebooker      = (*Notebook)(nil)
	_ gextras.Nativer = (*Notebook)(nil)
)

func wrapNotebook(obj *externglib.Object) Notebooker {
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

	_notebook = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Notebook)

	return _notebook
}

// AppendPage appends a page to @notebook.
func (notebook *Notebook) AppendPage(child Widgetter, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_append_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// AppendPageMenu appends a page to @notebook, specifying the widget to use as
// the label in the popup menu.
func (notebook *Notebook) AppendPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer((menuLabel).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_append_page_menu(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DetachTab removes the child from the notebook.
//
// This function is very similar to gtk_container_remove(), but additionally
// informs the notebook that the removal is happening as part of a tab DND
// operation, which should not be cancelled.
func (notebook *Notebook) DetachTab(child Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_notebook_detach_tab(_arg0, _arg1)
}

// CurrentPage returns the page number of the current page.
func (notebook *Notebook) CurrentPage() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GroupName gets the current group name for @notebook.
func (notebook *Notebook) GroupName() string {
	var _arg0 *C.GtkNotebook // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_group_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// MenuLabel retrieves the menu label widget of the page containing @child.
func (notebook *Notebook) MenuLabel(child Widgetter) *Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_menu_label(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// MenuLabelText retrieves the text of the menu label for the page containing
// @child.
func (notebook *Notebook) MenuLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_menu_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// NPages gets the number of pages in a notebook.
func (notebook *Notebook) NPages() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number @page_num.
func (notebook *Notebook) NthPage(pageNum int) *Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_notebook_get_nth_page(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// Scrollable returns whether the tab label area has arrows for scrolling. See
// gtk_notebook_set_scrollable().
func (notebook *Notebook) Scrollable() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_scrollable(_arg0)

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

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabDetachable returns whether the tab contents can be detached from
// @notebook.
func (notebook *Notebook) TabDetachable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_tab_detachable(_arg0, _arg1)

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

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// TabLabel returns the tab label widget for the page @child. nil is returned if
// @child is not in @notebook or if no tab label has specifically been set for
// @child.
func (notebook *Notebook) TabLabel(child Widgetter) *Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_tab_label(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// TabLabelText retrieves the text of the tab label for the page containing
// @child.
func (notebook *Notebook) TabLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_tab_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// TabPos gets the edge at which the tabs for switching pages in the notebook
// are drawn.
func (notebook *Notebook) TabPos() PositionType {
	var _arg0 *C.GtkNotebook    // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_tab_pos(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// TabReorderable gets whether the tab can be reordered via drag and drop or
// not.
func (notebook *Notebook) TabReorderable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_get_tab_reorderable(_arg0, _arg1)

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

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// InsertPage: insert a page into @notebook at the given position.
func (notebook *Notebook) InsertPage(child Widgetter, tabLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))
	_arg3 = C.gint(position)

	_cret = C.gtk_notebook_insert_page(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertPageMenu: insert a page into @notebook at the given position,
// specifying the widget to use as the label in the popup menu.
func (notebook *Notebook) InsertPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _arg4 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer((menuLabel).(gextras.Nativer).Native()))
	_arg4 = C.gint(position)

	_cret = C.gtk_notebook_insert_page_menu(_arg0, _arg1, _arg2, _arg3, _arg4)

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
}

// PageNum finds the index of the page which contains the given child widget.
func (notebook *Notebook) PageNum(child Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_page_num(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PopupDisable disables the popup menu.
func (notebook *Notebook) PopupDisable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_disable(_arg0)
}

// PopupEnable enables the popup menu: if the user clicks with the right mouse
// button on the tab labels, a menu with all the pages will be popped up.
func (notebook *Notebook) PopupEnable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_enable(_arg0)
}

// PrependPage prepends a page to @notebook.
func (notebook *Notebook) PrependPage(child Widgetter, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_prepend_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrependPageMenu prepends a page to @notebook, specifying the widget to use as
// the label in the popup menu.
func (notebook *Notebook) PrependPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer((menuLabel).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_prepend_page_menu(_arg0, _arg1, _arg2, _arg3)

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
}

// RemovePage removes a page from the notebook given its index in the notebook.
func (notebook *Notebook) RemovePage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_remove_page(_arg0, _arg1)
}

// ReorderChild reorders the page containing @child, so that it appears in
// position @position. If @position is greater than or equal to the number of
// children in the list or negative, @child will be moved to the end of the
// list.
func (notebook *Notebook) ReorderChild(child Widgetter, position int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = C.gint(position)

	C.gtk_notebook_reorder_child(_arg0, _arg1, _arg2)
}

// SetCurrentPage switches to the page number @page_num.
//
// Note that due to historical reasons, GtkNotebook refuses to switch to a page
// unless the child widget is visible. Therefore, it is recommended to show
// child widgets before adding them to a notebook.
func (notebook *Notebook) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_set_current_page(_arg0, _arg1)
}

// SetGroupName sets a group name for @notebook.
//
// Notebooks with the same name will be able to exchange tabs via drag and drop.
// A notebook with a nil group name will not be able to exchange tabs with any
// other notebook.
func (notebook *Notebook) SetGroupName(groupName string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_notebook_set_group_name(_arg0, _arg1)
}

// SetMenuLabel changes the menu label for the page containing @child.
func (notebook *Notebook) SetMenuLabel(child Widgetter, menuLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((menuLabel).(gextras.Nativer).Native()))

	C.gtk_notebook_set_menu_label(_arg0, _arg1, _arg2)
}

// SetMenuLabelText creates a new label and sets it as the menu label of @child.
func (notebook *Notebook) SetMenuLabelText(child Widgetter, menuText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(C.CString(menuText))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_menu_label_text(_arg0, _arg1, _arg2)
}

// SetScrollable sets whether the tab label area will have arrows for scrolling
// if there are too many tabs to fit in the area.
func (notebook *Notebook) SetScrollable(scrollable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if scrollable {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_scrollable(_arg0, _arg1)
}

// SetShowBorder sets whether a bevel will be drawn around the notebook pages.
// This only has a visual effect when the tabs are not shown. See
// gtk_notebook_set_show_tabs().
func (notebook *Notebook) SetShowBorder(showBorder bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showBorder {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_border(_arg0, _arg1)
}

// SetShowTabs sets whether to show the tabs for the notebook or not.
func (notebook *Notebook) SetShowTabs(showTabs bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showTabs {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_tabs(_arg0, _arg1)
}

// SetTabDetachable sets whether the tab can be detached from @notebook to
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
func (notebook *Notebook) SetTabDetachable(child Widgetter, detachable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	if detachable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_detachable(_arg0, _arg1, _arg2)
}

// SetTabLabel changes the tab label for @child. If nil is specified for
// @tab_label, then the page will have the label “page N”.
func (notebook *Notebook) SetTabLabel(child Widgetter, tabLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((tabLabel).(gextras.Nativer).Native()))

	C.gtk_notebook_set_tab_label(_arg0, _arg1, _arg2)
}

// SetTabLabelText creates a new label and sets it as the tab label for the page
// containing @child.
func (notebook *Notebook) SetTabLabelText(child Widgetter, tabText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(C.CString(tabText))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_tab_label_text(_arg0, _arg1, _arg2)
}

// SetTabReorderable sets whether the notebook tab can be reordered via drag and
// drop or not.
func (notebook *Notebook) SetTabReorderable(child Widgetter, reorderable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	if reorderable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_reorderable(_arg0, _arg1, _arg2)
}
