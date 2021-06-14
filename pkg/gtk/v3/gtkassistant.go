// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_assistant_page_type_get_type()), F: marshalAssistantPageType},
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistant},
	})
}

// AssistantPageType: an enum for determining the page role inside the
// Assistant. It's used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType int

const (
	// AssistantPageTypeContent: the page has regular contents. Both the Back
	// and forward buttons will be shown.
	AssistantPageTypeContent AssistantPageType = 0
	// AssistantPageTypeIntro: the page contains an introduction to the
	// assistant task. Only the Forward button will be shown if there is a next
	// page.
	AssistantPageTypeIntro AssistantPageType = 1
	// AssistantPageTypeConfirm: the page lets the user confirm or deny the
	// changes. The Back and Apply buttons will be shown.
	AssistantPageTypeConfirm AssistantPageType = 2
	// AssistantPageTypeSummary: the page informs the user of the changes done.
	// Only the Close button will be shown.
	AssistantPageTypeSummary AssistantPageType = 3
	// AssistantPageTypeProgress: used for tasks that take a long time to
	// complete, blocks the assistant until the page is marked as complete. Only
	// the back button will be shown.
	AssistantPageTypeProgress AssistantPageType = 4
	// AssistantPageTypeCustom: used for when other page types are not
	// appropriate. No buttons will be shown, and the application must add its
	// own buttons through gtk_assistant_add_action_widget().
	AssistantPageTypeCustom AssistantPageType = 5
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AssistantPageFunc: a function used by gtk_assistant_set_forward_page_func()
// to know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int) (gint int)

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) C.gint {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int // out

	currentPage = (int)(arg0)

	fn := v.(AssistantPageFunc)
	gint := fn(currentPage)

	var cret C.gint // out

	cret = C.gint(gint)

	return cret
}

// Assistant: a Assistant is a widget used to represent a generally complex
// operation splitted in several steps, guiding the user through its pages and
// controlling the page flow to collect the necessary data.
//
// The design of GtkAssistant is that it controls what buttons to show and to
// make sensitive, based on what it knows about the page sequence and the
// [type][GtkAssistantPageType] of each page, in addition to state information
// like the page [completion][gtk-assistant-set-page-complete] and
// [committed][gtk-assistant-commit] status.
//
// If you have a case that doesn’t quite fit in Assistants way of handling
// buttons, you can use the K_ASSISTANT_PAGE_CUSTOM page type and handle buttons
// yourself.
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the Buildable interface exposes the
// @action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant interface {
	Window
	Buildable

	// AddActionWidget adds a widget to the action area of a Assistant.
	AddActionWidget(child Widget)
	// AppendPage appends a page to the @assistant.
	AppendPage(page Widget) int
	// Commit erases the visited page history so the back button is not shown on
	// the current page, and removes the cancel button from subsequent pages.
	//
	// Use this when the information provided up to the current page is
	// hereafter deemed permanent and cannot be modified or undone. For example,
	// showing a progress page to track a long-running, unreversible operation
	// after the user has clicked apply on a confirmation page.
	Commit()
	// CurrentPage returns the page number of the current page.
	CurrentPage() int
	// NPages returns the number of pages in the @assistant
	NPages() int
	// PageComplete gets whether @page is complete.
	PageComplete(page Widget) bool
	// PageHasPadding gets whether page has padding.
	PageHasPadding(page Widget) bool
	// PageTitle gets the title for @page.
	PageTitle(page Widget) string
	// InsertPage inserts a page in the @assistant at a given position.
	InsertPage(page Widget, position int) int
	// NextPage: navigate to the next page.
	//
	// It is a programming error to call this function when there is no next
	// page.
	//
	// This function is for use when creating pages of the
	// K_ASSISTANT_PAGE_CUSTOM type.
	NextPage()
	// PrependPage prepends a page to the @assistant.
	PrependPage(page Widget) int
	// PreviousPage: navigate to the previous visited page.
	//
	// It is a programming error to call this function when no previous page is
	// available.
	//
	// This function is for use when creating pages of the
	// K_ASSISTANT_PAGE_CUSTOM type.
	PreviousPage()
	// RemoveActionWidget removes a widget from the action area of a Assistant.
	RemoveActionWidget(child Widget)
	// RemovePage removes the @page_num’s page from @assistant.
	RemovePage(pageNum int)
	// SetCurrentPage switches the page to @page_num.
	//
	// Note that this will only be necessary in custom buttons, as the
	// @assistant flow can be set with gtk_assistant_set_forward_page_func().
	SetCurrentPage(pageNum int)
	// SetPageComplete sets whether @page contents are complete.
	//
	// This will make @assistant update the buttons state to be able to continue
	// the task.
	SetPageComplete(page Widget, complete bool)
	// SetPageHasPadding sets whether the assistant is adding padding around the
	// page.
	SetPageHasPadding(page Widget, hasPadding bool)
	// SetPageHeaderImage sets a header image for @page.
	SetPageHeaderImage(page Widget, pixbuf gdkpixbuf.Pixbuf)
	// SetPageSideImage sets a side image for @page.
	//
	// This image used to be displayed in the side area of the assistant when
	// @page is the current page.
	SetPageSideImage(page Widget, pixbuf gdkpixbuf.Pixbuf)
	// SetPageTitle sets a title for @page.
	//
	// The title is displayed in the header area of the assistant when @page is
	// the current page.
	SetPageTitle(page Widget, title string)
	// SetPageType sets the page type for @page.
	//
	// The page type determines the page behavior in the @assistant.
	SetPageType(page Widget, typ AssistantPageType)
	// UpdateButtonsState forces @assistant to recompute the buttons state.
	//
	// GTK+ automatically takes care of this in most situations, e.g. when the
	// user goes to a different page, or when the visibility or completeness of
	// a page changes.
	//
	// One situation where it can be necessary to call this function is when
	// changing a value on the current page affects the future page flow of the
	// assistant.
	UpdateButtonsState()
}

// assistant implements the Assistant class.
type assistant struct {
	Window
	Buildable
}

var _ Assistant = (*assistant)(nil)

// WrapAssistant wraps a GObject to the right type. It is
// primarily used internally.
func WrapAssistant(obj *externglib.Object) Assistant {
	return assistant{
		Window:    WrapWindow(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAssistant(obj), nil
}

// AddActionWidget adds a widget to the action area of a Assistant.
func (a assistant) AddActionWidget(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
}

// AppendPage appends a page to the @assistant.
func (a assistant) AppendPage(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var _cret C.gint // in

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Commit erases the visited page history so the back button is not shown on
// the current page, and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is
// hereafter deemed permanent and cannot be modified or undone. For example,
// showing a progress page to track a long-running, unreversible operation
// after the user has clicked apply on a confirmation page.
func (a assistant) Commit() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_commit(_arg0)
}

// CurrentPage returns the page number of the current page.
func (a assistant) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	var _cret C.gint // in

	_cret = C.gtk_assistant_get_current_page(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// NPages returns the number of pages in the @assistant
func (a assistant) NPages() int {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	var _cret C.gint // in

	_cret = C.gtk_assistant_get_n_pages(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// PageComplete gets whether @page is complete.
func (a assistant) PageComplete(page Widget) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHasPadding gets whether page has padding.
func (a assistant) PageHasPadding(page Widget) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_assistant_get_page_has_padding(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageTitle gets the title for @page.
func (a assistant) PageTitle(page Widget) string {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// InsertPage inserts a page in the @assistant at a given position.
func (a assistant) InsertPage(page Widget, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = C.gint(position)

	var _cret C.gint // in

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next
// page.
//
// This function is for use when creating pages of the
// K_ASSISTANT_PAGE_CUSTOM type.
func (a assistant) NextPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_next_page(_arg0)
}

// PrependPage prepends a page to the @assistant.
func (a assistant) PrependPage(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var _cret C.gint // in

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the
// K_ASSISTANT_PAGE_CUSTOM type.
func (a assistant) PreviousPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_previous_page(_arg0)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
func (a assistant) RemoveActionWidget(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
}

// RemovePage removes the @page_num’s page from @assistant.
func (a assistant) RemovePage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_remove_page(_arg0, _arg1)
}

// SetCurrentPage switches the page to @page_num.
//
// Note that this will only be necessary in custom buttons, as the
// @assistant flow can be set with gtk_assistant_set_forward_page_func().
func (a assistant) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
}

// SetPageComplete sets whether @page contents are complete.
//
// This will make @assistant update the buttons state to be able to continue
// the task.
func (a assistant) SetPageComplete(page Widget, complete bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
}

// SetPageHasPadding sets whether the assistant is adding padding around the
// page.
func (a assistant) SetPageHasPadding(page Widget, hasPadding bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if hasPadding {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_has_padding(_arg0, _arg1, _arg2)
}

// SetPageHeaderImage sets a header image for @page.
func (a assistant) SetPageHeaderImage(page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_header_image(_arg0, _arg1, _arg2)
}

// SetPageSideImage sets a side image for @page.
//
// This image used to be displayed in the side area of the assistant when
// @page is the current page.
func (a assistant) SetPageSideImage(page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_side_image(_arg0, _arg1, _arg2)
}

// SetPageTitle sets a title for @page.
//
// The title is displayed in the header area of the assistant when @page is
// the current page.
func (a assistant) SetPageTitle(page Widget, title string) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
}

// SetPageType sets the page type for @page.
//
// The page type determines the page behavior in the @assistant.
func (a assistant) SetPageType(page Widget, typ AssistantPageType) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkAssistantPageType // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (C.GtkAssistantPageType)(typ)

	C.gtk_assistant_set_page_type(_arg0, _arg1, _arg2)
}

// UpdateButtonsState forces @assistant to recompute the buttons state.
//
// GTK+ automatically takes care of this in most situations, e.g. when the
// user goes to a different page, or when the visibility or completeness of
// a page changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (a assistant) UpdateButtonsState() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
}
