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
		{T: externglib.Type(C.gtk_notebook_get_type()), F: marshalNotebook},
	})
}

// Notebook: the Notebook widget is a Container whose children are pages that
// can be switched between using tab labels along one edge.
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
type Notebook interface {
	Container
	Buildable

	// AppendPage appends a page to @notebook.
	AppendPage(child Widget, tabLabel Widget) int
	// AppendPageMenu appends a page to @notebook, specifying the widget to use
	// as the label in the popup menu.
	AppendPageMenu(child Widget, tabLabel Widget, menuLabel Widget) int
	// DetachTab removes the child from the notebook.
	//
	// This function is very similar to gtk_container_remove(), but additionally
	// informs the notebook that the removal is happening as part of a tab DND
	// operation, which should not be cancelled.
	DetachTab(child Widget)
	// ActionWidget gets one of the action widgets. See
	// gtk_notebook_set_action_widget().
	ActionWidget(packType PackType) Widget
	// CurrentPage returns the page number of the current page.
	CurrentPage() int
	// GroupName gets the current group name for @notebook.
	GroupName() string
	// MenuLabel retrieves the menu label widget of the page containing @child.
	MenuLabel(child Widget) Widget
	// MenuLabelText retrieves the text of the menu label for the page
	// containing @child.
	MenuLabelText(child Widget) string
	// NPages gets the number of pages in a notebook.
	NPages() int
	// NthPage returns the child widget contained in page number @page_num.
	NthPage(pageNum int) Widget
	// Scrollable returns whether the tab label area has arrows for scrolling.
	// See gtk_notebook_set_scrollable().
	Scrollable() bool
	// ShowBorder returns whether a bevel will be drawn around the notebook
	// pages. See gtk_notebook_set_show_border().
	ShowBorder() bool
	// ShowTabs returns whether the tabs of the notebook are shown. See
	// gtk_notebook_set_show_tabs().
	ShowTabs() bool
	// TabDetachable returns whether the tab contents can be detached from
	// @notebook.
	TabDetachable(child Widget) bool
	// TabHborder returns the horizontal width of a tab border.
	TabHborder() uint16
	// TabLabel returns the tab label widget for the page @child. nil is
	// returned if @child is not in @notebook or if no tab label has
	// specifically been set for @child.
	TabLabel(child Widget) Widget
	// TabLabelText retrieves the text of the tab label for the page containing
	// @child.
	TabLabelText(child Widget) string
	// TabPos gets the edge at which the tabs for switching pages in the
	// notebook are drawn.
	TabPos() PositionType
	// TabReorderable gets whether the tab can be reordered via drag and drop or
	// not.
	TabReorderable(child Widget) bool
	// TabVborder returns the vertical width of a tab border.
	TabVborder() uint16
	// InsertPage: insert a page into @notebook at the given position.
	InsertPage(child Widget, tabLabel Widget, position int) int
	// InsertPageMenu: insert a page into @notebook at the given position,
	// specifying the widget to use as the label in the popup menu.
	InsertPageMenu(child Widget, tabLabel Widget, menuLabel Widget, position int) int
	// NextPage switches to the next page. Nothing happens if the current page
	// is the last page.
	NextPage()
	// PageNum finds the index of the page which contains the given child
	// widget.
	PageNum(child Widget) int
	// PopupDisable disables the popup menu.
	PopupDisable()
	// PopupEnable enables the popup menu: if the user clicks with the right
	// mouse button on the tab labels, a menu with all the pages will be popped
	// up.
	PopupEnable()
	// PrependPage prepends a page to @notebook.
	PrependPage(child Widget, tabLabel Widget) int
	// PrependPageMenu prepends a page to @notebook, specifying the widget to
	// use as the label in the popup menu.
	PrependPageMenu(child Widget, tabLabel Widget, menuLabel Widget) int
	// PrevPage switches to the previous page. Nothing happens if the current
	// page is the first page.
	PrevPage()
	// RemovePage removes a page from the notebook given its index in the
	// notebook.
	RemovePage(pageNum int)
	// ReorderChild reorders the page containing @child, so that it appears in
	// position @position. If @position is greater than or equal to the number
	// of children in the list or negative, @child will be moved to the end of
	// the list.
	ReorderChild(child Widget, position int)
	// SetActionWidget sets @widget as one of the action widgets. Depending on
	// the pack type the widget will be placed before or after the tabs. You can
	// use a Box if you need to pack more than one widget on the same side.
	//
	// Note that action widgets are “internal” children of the notebook and thus
	// not included in the list returned from gtk_container_foreach().
	SetActionWidget(widget Widget, packType PackType)
	// SetCurrentPage switches to the page number @page_num.
	//
	// Note that due to historical reasons, GtkNotebook refuses to switch to a
	// page unless the child widget is visible. Therefore, it is recommended to
	// show child widgets before adding them to a notebook.
	SetCurrentPage(pageNum int)
	// SetGroupName sets a group name for @notebook.
	//
	// Notebooks with the same name will be able to exchange tabs via drag and
	// drop. A notebook with a nil group name will not be able to exchange tabs
	// with any other notebook.
	SetGroupName(groupName string)
	// SetMenuLabel changes the menu label for the page containing @child.
	SetMenuLabel(child Widget, menuLabel Widget)
	// SetMenuLabelText creates a new label and sets it as the menu label of
	// @child.
	SetMenuLabelText(child Widget, menuText string)
	// SetScrollable sets whether the tab label area will have arrows for
	// scrolling if there are too many tabs to fit in the area.
	SetScrollable(scrollable bool)
	// SetShowBorder sets whether a bevel will be drawn around the notebook
	// pages. This only has a visual effect when the tabs are not shown. See
	// gtk_notebook_set_show_tabs().
	SetShowBorder(showBorder bool)
	// SetShowTabs sets whether to show the tabs for the notebook or not.
	SetShowTabs(showTabs bool)
	// SetTabDetachable sets whether the tab can be detached from @notebook to
	// another notebook or widget.
	//
	// Note that 2 notebooks must share a common group identificator (see
	// gtk_notebook_set_group_name()) to allow automatic tabs interchange
	// between them.
	//
	// If you want a widget to interact with a notebook through DnD (i.e.:
	// accept dragged tabs from it) it must be set as a drop destination and
	// accept the target “GTK_NOTEBOOK_TAB”. The notebook will fill the
	// selection with a GtkWidget** pointing to the child widget that
	// corresponds to the dropped tab.
	//
	// Note that you should use gtk_notebook_detach_tab() instead of
	// gtk_container_remove() if you want to remove the tab from the source
	// notebook as part of accepting a drop. Otherwise, the source notebook will
	// think that the dragged tab was removed from underneath the ongoing drag
	// operation, and will initiate a drag cancel animation.
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
	// If you want a notebook to accept drags from other widgets, you will have
	// to set your own DnD code to do it.
	SetTabDetachable(child Widget, detachable bool)
	// SetTabLabel changes the tab label for @child. If nil is specified for
	// @tab_label, then the page will have the label “page N”.
	SetTabLabel(child Widget, tabLabel Widget)
	// SetTabLabelText creates a new label and sets it as the tab label for the
	// page containing @child.
	SetTabLabelText(child Widget, tabText string)
	// SetTabPos sets the edge at which the tabs for switching pages in the
	// notebook are drawn.
	SetTabPos(pos PositionType)
	// SetTabReorderable sets whether the notebook tab can be reordered via drag
	// and drop or not.
	SetTabReorderable(child Widget, reorderable bool)
}

// notebook implements the Notebook interface.
type notebook struct {
	Container
	Buildable
}

var _ Notebook = (*notebook)(nil)

// WrapNotebook wraps a GObject to the right type. It is
// primarily used internally.
func WrapNotebook(obj *externglib.Object) Notebook {
	return Notebook{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalNotebook(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNotebook(obj), nil
}

// NewNotebook constructs a class Notebook.
func NewNotebook() Notebook {
	var cret C.GtkNotebook
	var goret Notebook

	cret = C.gtk_notebook_new()

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Notebook)

	return goret
}

// AppendPage appends a page to @notebook.
func (n notebook) AppendPage(child Widget, tabLabel Widget) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_append_page(arg0, arg1, arg2)

	goret = int(cret)

	return goret
}

// AppendPageMenu appends a page to @notebook, specifying the widget to use
// as the label in the popup menu.
func (n notebook) AppendPageMenu(child Widget, tabLabel Widget, menuLabel Widget) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget
	var arg3 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_append_page_menu(arg0, arg1, arg2, arg3)

	goret = int(cret)

	return goret
}

// DetachTab removes the child from the notebook.
//
// This function is very similar to gtk_container_remove(), but additionally
// informs the notebook that the removal is happening as part of a tab DND
// operation, which should not be cancelled.
func (n notebook) DetachTab(child Widget) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_notebook_detach_tab(arg0, arg1)
}

// ActionWidget gets one of the action widgets. See
// gtk_notebook_set_action_widget().
func (n notebook) ActionWidget(packType PackType) Widget {
	var arg0 *C.GtkNotebook
	var arg1 C.GtkPackType

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (C.GtkPackType)(packType)

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_notebook_get_action_widget(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// CurrentPage returns the page number of the current page.
func (n notebook) CurrentPage() int {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_get_current_page(arg0)

	goret = int(cret)

	return goret
}

// GroupName gets the current group name for @notebook.
func (n notebook) GroupName() string {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret *C.gchar
	var goret string

	cret = C.gtk_notebook_get_group_name(arg0)

	goret = C.GoString(cret)

	return goret
}

// MenuLabel retrieves the menu label widget of the page containing @child.
func (n notebook) MenuLabel(child Widget) Widget {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_notebook_get_menu_label(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// MenuLabelText retrieves the text of the menu label for the page
// containing @child.
func (n notebook) MenuLabelText(child Widget) string {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret *C.gchar
	var goret string

	cret = C.gtk_notebook_get_menu_label_text(arg0, arg1)

	goret = C.GoString(cret)

	return goret
}

// NPages gets the number of pages in a notebook.
func (n notebook) NPages() int {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_get_n_pages(arg0)

	goret = int(cret)

	return goret
}

// NthPage returns the child widget contained in page number @page_num.
func (n notebook) NthPage(pageNum int) Widget {
	var arg0 *C.GtkNotebook
	var arg1 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = C.gint(pageNum)

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_notebook_get_nth_page(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// Scrollable returns whether the tab label area has arrows for scrolling.
// See gtk_notebook_set_scrollable().
func (n notebook) Scrollable() bool {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_notebook_get_scrollable(arg0)

	if cret {
		goret = true
	}

	return goret
}

// ShowBorder returns whether a bevel will be drawn around the notebook
// pages. See gtk_notebook_set_show_border().
func (n notebook) ShowBorder() bool {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_notebook_get_show_border(arg0)

	if cret {
		goret = true
	}

	return goret
}

// ShowTabs returns whether the tabs of the notebook are shown. See
// gtk_notebook_set_show_tabs().
func (n notebook) ShowTabs() bool {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_notebook_get_show_tabs(arg0)

	if cret {
		goret = true
	}

	return goret
}

// TabDetachable returns whether the tab contents can be detached from
// @notebook.
func (n notebook) TabDetachable(child Widget) bool {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_notebook_get_tab_detachable(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// TabHborder returns the horizontal width of a tab border.
func (n notebook) TabHborder() uint16 {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.guint16
	var goret uint16

	cret = C.gtk_notebook_get_tab_hborder(arg0)

	goret = uint16(cret)

	return goret
}

// TabLabel returns the tab label widget for the page @child. nil is
// returned if @child is not in @notebook or if no tab label has
// specifically been set for @child.
func (n notebook) TabLabel(child Widget) Widget {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret *C.GtkWidget
	var goret Widget

	cret = C.gtk_notebook_get_tab_label(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret
}

// TabLabelText retrieves the text of the tab label for the page containing
// @child.
func (n notebook) TabLabelText(child Widget) string {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret *C.gchar
	var goret string

	cret = C.gtk_notebook_get_tab_label_text(arg0, arg1)

	goret = C.GoString(cret)

	return goret
}

// TabPos gets the edge at which the tabs for switching pages in the
// notebook are drawn.
func (n notebook) TabPos() PositionType {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.GtkPositionType
	var goret PositionType

	cret = C.gtk_notebook_get_tab_pos(arg0)

	goret = PositionType(cret)

	return goret
}

// TabReorderable gets whether the tab can be reordered via drag and drop or
// not.
func (n notebook) TabReorderable(child Widget) bool {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_notebook_get_tab_reorderable(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// TabVborder returns the vertical width of a tab border.
func (n notebook) TabVborder() uint16 {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	var cret C.guint16
	var goret uint16

	cret = C.gtk_notebook_get_tab_vborder(arg0)

	goret = uint16(cret)

	return goret
}

// InsertPage: insert a page into @notebook at the given position.
func (n notebook) InsertPage(child Widget, tabLabel Widget, position int) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget
	var arg3 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	arg3 = C.gint(position)

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_insert_page(arg0, arg1, arg2, arg3)

	goret = int(cret)

	return goret
}

// InsertPageMenu: insert a page into @notebook at the given position,
// specifying the widget to use as the label in the popup menu.
func (n notebook) InsertPageMenu(child Widget, tabLabel Widget, menuLabel Widget, position int) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget
	var arg3 *C.GtkWidget
	var arg4 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	arg4 = C.gint(position)

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_insert_page_menu(arg0, arg1, arg2, arg3, arg4)

	goret = int(cret)

	return goret
}

// NextPage switches to the next page. Nothing happens if the current page
// is the last page.
func (n notebook) NextPage() {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_next_page(arg0)
}

// PageNum finds the index of the page which contains the given child
// widget.
func (n notebook) PageNum(child Widget) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_page_num(arg0, arg1)

	goret = int(cret)

	return goret
}

// PopupDisable disables the popup menu.
func (n notebook) PopupDisable() {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_popup_disable(arg0)
}

// PopupEnable enables the popup menu: if the user clicks with the right
// mouse button on the tab labels, a menu with all the pages will be popped
// up.
func (n notebook) PopupEnable() {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_popup_enable(arg0)
}

// PrependPage prepends a page to @notebook.
func (n notebook) PrependPage(child Widget, tabLabel Widget) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_prepend_page(arg0, arg1, arg2)

	goret = int(cret)

	return goret
}

// PrependPageMenu prepends a page to @notebook, specifying the widget to
// use as the label in the popup menu.
func (n notebook) PrependPageMenu(child Widget, tabLabel Widget, menuLabel Widget) int {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget
	var arg3 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_notebook_prepend_page_menu(arg0, arg1, arg2, arg3)

	goret = int(cret)

	return goret
}

// PrevPage switches to the previous page. Nothing happens if the current
// page is the first page.
func (n notebook) PrevPage() {
	var arg0 *C.GtkNotebook

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_prev_page(arg0)
}

// RemovePage removes a page from the notebook given its index in the
// notebook.
func (n notebook) RemovePage(pageNum int) {
	var arg0 *C.GtkNotebook
	var arg1 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = C.gint(pageNum)

	C.gtk_notebook_remove_page(arg0, arg1)
}

// ReorderChild reorders the page containing @child, so that it appears in
// position @position. If @position is greater than or equal to the number
// of children in the list or negative, @child will be moved to the end of
// the list.
func (n notebook) ReorderChild(child Widget, position int) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = C.gint(position)

	C.gtk_notebook_reorder_child(arg0, arg1, arg2)
}

// SetActionWidget sets @widget as one of the action widgets. Depending on
// the pack type the widget will be placed before or after the tabs. You can
// use a Box if you need to pack more than one widget on the same side.
//
// Note that action widgets are “internal” children of the notebook and thus
// not included in the list returned from gtk_container_foreach().
func (n notebook) SetActionWidget(widget Widget, packType PackType) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 C.GtkPackType

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GtkPackType)(packType)

	C.gtk_notebook_set_action_widget(arg0, arg1, arg2)
}

// SetCurrentPage switches to the page number @page_num.
//
// Note that due to historical reasons, GtkNotebook refuses to switch to a
// page unless the child widget is visible. Therefore, it is recommended to
// show child widgets before adding them to a notebook.
func (n notebook) SetCurrentPage(pageNum int) {
	var arg0 *C.GtkNotebook
	var arg1 C.gint

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = C.gint(pageNum)

	C.gtk_notebook_set_current_page(arg0, arg1)
}

// SetGroupName sets a group name for @notebook.
//
// Notebooks with the same name will be able to exchange tabs via drag and
// drop. A notebook with a nil group name will not be able to exchange tabs
// with any other notebook.
func (n notebook) SetGroupName(groupName string) {
	var arg0 *C.GtkNotebook
	var arg1 *C.gchar

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_notebook_set_group_name(arg0, arg1)
}

// SetMenuLabel changes the menu label for the page containing @child.
func (n notebook) SetMenuLabel(child Widget, menuLabel Widget) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	C.gtk_notebook_set_menu_label(arg0, arg1, arg2)
}

// SetMenuLabelText creates a new label and sets it as the menu label of
// @child.
func (n notebook) SetMenuLabelText(child Widget, menuText string) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(menuText))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_notebook_set_menu_label_text(arg0, arg1, arg2)
}

// SetScrollable sets whether the tab label area will have arrows for
// scrolling if there are too many tabs to fit in the area.
func (n notebook) SetScrollable(scrollable bool) {
	var arg0 *C.GtkNotebook
	var arg1 C.gboolean

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if scrollable {
		arg1 = C.gboolean(1)
	}

	C.gtk_notebook_set_scrollable(arg0, arg1)
}

// SetShowBorder sets whether a bevel will be drawn around the notebook
// pages. This only has a visual effect when the tabs are not shown. See
// gtk_notebook_set_show_tabs().
func (n notebook) SetShowBorder(showBorder bool) {
	var arg0 *C.GtkNotebook
	var arg1 C.gboolean

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if showBorder {
		arg1 = C.gboolean(1)
	}

	C.gtk_notebook_set_show_border(arg0, arg1)
}

// SetShowTabs sets whether to show the tabs for the notebook or not.
func (n notebook) SetShowTabs(showTabs bool) {
	var arg0 *C.GtkNotebook
	var arg1 C.gboolean

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if showTabs {
		arg1 = C.gboolean(1)
	}

	C.gtk_notebook_set_show_tabs(arg0, arg1)
}

// SetTabDetachable sets whether the tab can be detached from @notebook to
// another notebook or widget.
//
// Note that 2 notebooks must share a common group identificator (see
// gtk_notebook_set_group_name()) to allow automatic tabs interchange
// between them.
//
// If you want a widget to interact with a notebook through DnD (i.e.:
// accept dragged tabs from it) it must be set as a drop destination and
// accept the target “GTK_NOTEBOOK_TAB”. The notebook will fill the
// selection with a GtkWidget** pointing to the child widget that
// corresponds to the dropped tab.
//
// Note that you should use gtk_notebook_detach_tab() instead of
// gtk_container_remove() if you want to remove the tab from the source
// notebook as part of accepting a drop. Otherwise, the source notebook will
// think that the dragged tab was removed from underneath the ongoing drag
// operation, and will initiate a drag cancel animation.
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
// If you want a notebook to accept drags from other widgets, you will have
// to set your own DnD code to do it.
func (n notebook) SetTabDetachable(child Widget, detachable bool) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 C.gboolean

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if detachable {
		arg2 = C.gboolean(1)
	}

	C.gtk_notebook_set_tab_detachable(arg0, arg1, arg2)
}

// SetTabLabel changes the tab label for @child. If nil is specified for
// @tab_label, then the page will have the label “page N”.
func (n notebook) SetTabLabel(child Widget, tabLabel Widget) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	C.gtk_notebook_set_tab_label(arg0, arg1, arg2)
}

// SetTabLabelText creates a new label and sets it as the tab label for the
// page containing @child.
func (n notebook) SetTabLabelText(child Widget, tabText string) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(tabText))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_notebook_set_tab_label_text(arg0, arg1, arg2)
}

// SetTabPos sets the edge at which the tabs for switching pages in the
// notebook are drawn.
func (n notebook) SetTabPos(pos PositionType) {
	var arg0 *C.GtkNotebook
	var arg1 C.GtkPositionType

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (C.GtkPositionType)(pos)

	C.gtk_notebook_set_tab_pos(arg0, arg1)
}

// SetTabReorderable sets whether the notebook tab can be reordered via drag
// and drop or not.
func (n notebook) SetTabReorderable(child Widget, reorderable bool) {
	var arg0 *C.GtkNotebook
	var arg1 *C.GtkWidget
	var arg2 C.gboolean

	arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if reorderable {
		arg2 = C.gboolean(1)
	}

	C.gtk_notebook_set_tab_reorderable(arg0, arg1, arg2)
}
