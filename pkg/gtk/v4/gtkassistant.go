// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern int _gotk4_gtk4_AssistantPageFunc(int, gpointer);
// extern void _gotk4_gtk4_Assistant_ConnectApply(gpointer, guintptr);
// extern void _gotk4_gtk4_Assistant_ConnectCancel(gpointer, guintptr);
// extern void _gotk4_gtk4_Assistant_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk4_Assistant_ConnectEscape(gpointer, guintptr);
// extern void _gotk4_gtk4_Assistant_ConnectPrepare(gpointer, void*, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkassistant.go.
var (
	GTypeAssistantPageType = coreglib.Type(C.gtk_assistant_page_type_get_type())
	GTypeAssistant         = coreglib.Type(C.gtk_assistant_get_type())
	GTypeAssistantPage     = coreglib.Type(C.gtk_assistant_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAssistantPageType, F: marshalAssistantPageType},
		{T: GTypeAssistant, F: marshalAssistant},
		{T: GTypeAssistantPage, F: marshalAssistantPage},
	})
}

// AssistantPageType determines the page role inside a GtkAssistant.
//
// The role is used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType C.gint

const (
	// AssistantPageContent: page has regular contents. Both the Back and
	// forward buttons will be shown.
	AssistantPageContent AssistantPageType = iota
	// AssistantPageIntro: page contains an introduction to the assistant task.
	// Only the Forward button will be shown if there is a next page.
	AssistantPageIntro
	// AssistantPageConfirm: page lets the user confirm or deny the changes. The
	// Back and Apply buttons will be shown.
	AssistantPageConfirm
	// AssistantPageSummary: page informs the user of the changes done. Only the
	// Close button will be shown.
	AssistantPageSummary
	// AssistantPageProgress: used for tasks that take a long time to complete,
	// blocks the assistant until the page is marked as complete. Only the back
	// button will be shown.
	AssistantPageProgress
	// AssistantPageCustom: used for when other page types are not appropriate.
	// No buttons will be shown, and the application must add its own buttons
	// through gtk_assistant_add_action_widget().
	AssistantPageCustom
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AssistantPageType.
func (a AssistantPageType) String() string {
	switch a {
	case AssistantPageContent:
		return "Content"
	case AssistantPageIntro:
		return "Intro"
	case AssistantPageConfirm:
		return "Confirm"
	case AssistantPageSummary:
		return "Summary"
	case AssistantPageProgress:
		return "Progress"
	case AssistantPageCustom:
		return "Custom"
	default:
		return fmt.Sprintf("AssistantPageType(%d)", a)
	}
}

// AssistantPageFunc: type of callback used to calculate the next page in a
// GtkAssistant.
//
// It’s called both for computing the next page when the user presses the
// “forward” button and for handling the behavior of the “last” button.
//
// See gtk.Assistant.SetForwardPageFunc().
type AssistantPageFunc func(currentPage int32) (gint int32)

//export _gotk4_gtk4_AssistantPageFunc
func _gotk4_gtk4_AssistantPageFunc(arg1 C.int, arg2 C.gpointer) (cret C.int) {
	var fn AssistantPageFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AssistantPageFunc)
	}

	var _currentPage int32 // out

	_currentPage = int32(arg1)

	gint := fn(_currentPage)

	cret = C.int(gint)

	return cret
}

// Assistant: GtkAssistant is used to represent a complex as a series of steps.
//
// !An example GtkAssistant (assistant.png)
//
// Each step consists of one or more pages. GtkAssistant guides the user through
// the pages, and controls the page flow to collect the data needed for the
// operation.
//
// GtkAssistant handles which buttons to show and to make sensitive based on
// page sequence knowledge and the gtk.AssistantPageType of each page in
// addition to state information like the *completed* and *committed* page
// statuses.
//
// If you have a case that doesn’t quite fit in GtkAssistants way of handling
// buttons, you can use the GTK_ASSISTANT_PAGE_CUSTOM page type and handle
// buttons yourself.
//
// GtkAssistant maintains a GtkAssistantPage object for each added child, which
// holds additional per-child properties. You obtain the GtkAssistantPage for a
// child with gtk.Assistant.GetPage().
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the GtkBuildable interface exposes the
// action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in GtkBuilder, simply add it as a child to the
// GtkAssistant object. If you need to set per-object properties, create a
// GtkAssistantPage object explicitly, and set the child widget as a property on
// it.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name window and style class
// .assistant.
type Assistant struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Widgetter         = (*Assistant)(nil)
	_ coreglib.Objector = (*Assistant)(nil)
)

func wrapAssistant(obj *coreglib.Object) *Assistant {
	return &Assistant{
		Window: Window{
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
			Root: Root{
				NativeSurface: NativeSurface{
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
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	return wrapAssistant(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Assistant_ConnectApply
func _gotk4_gtk4_Assistant_ConnectApply(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectApply is emitted when the apply button is clicked.
//
// The default behavior of the GtkAssistant is to switch to the page after the
// current page, unless the current page is the last one.
//
// A handler for the ::apply signal should carry out the actions for which the
// wizard has collected data. If the action takes a long time to complete, you
// might consider putting a page of type GTK_ASSISTANT_PAGE_PROGRESS after the
// confirmation page and handle this operation within the gtk.Assistant::prepare
// signal of the progress page.
func (assistant *Assistant) ConnectApply(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "apply", false, unsafe.Pointer(C._gotk4_gtk4_Assistant_ConnectApply), f)
}

//export _gotk4_gtk4_Assistant_ConnectCancel
func _gotk4_gtk4_Assistant_ConnectCancel(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectCancel is emitted when then the cancel button is clicked.
func (assistant *Assistant) ConnectCancel(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "cancel", false, unsafe.Pointer(C._gotk4_gtk4_Assistant_ConnectCancel), f)
}

//export _gotk4_gtk4_Assistant_ConnectClose
func _gotk4_gtk4_Assistant_ConnectClose(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClose is emitted either when the close button of a summary page is
// clicked, or when the apply button in the last page in the flow (of type
// GTK_ASSISTANT_PAGE_CONFIRM) is clicked.
func (assistant *Assistant) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "close", false, unsafe.Pointer(C._gotk4_gtk4_Assistant_ConnectClose), f)
}

//export _gotk4_gtk4_Assistant_ConnectEscape
func _gotk4_gtk4_Assistant_ConnectEscape(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectEscape: action signal for the Escape binding.
func (assistant *Assistant) ConnectEscape(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "escape", false, unsafe.Pointer(C._gotk4_gtk4_Assistant_ConnectEscape), f)
}

//export _gotk4_gtk4_Assistant_ConnectPrepare
func _gotk4_gtk4_Assistant_ConnectPrepare(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(page Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(page Widgetter))
	}

	var _page Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_page = rv
	}

	f(_page)
}

// ConnectPrepare is emitted when a new page is set as the assistant's current
// page, before making the new page visible.
//
// A handler for this signal can do any preparations which are necessary before
// showing page.
func (assistant *Assistant) ConnectPrepare(f func(page Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "prepare", false, unsafe.Pointer(C._gotk4_gtk4_Assistant_ConnectPrepare), f)
}

// NewAssistant creates a new GtkAssistant.
//
// The function returns the following values:
//
//    - assistant: newly created GtkAssistant.
//
func NewAssistant() *Assistant {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("new_Assistant", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _assistant *Assistant // out

	_assistant = wrapAssistant(coreglib.Take(unsafe.Pointer(_cret)))

	return _assistant
}

// AddActionWidget adds a widget to the action area of a GtkAssistant.
//
// The function takes the following parameters:
//
//    - child: GtkWidget.
//
func (assistant *Assistant) AddActionWidget(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("add_action_widget", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// AppendPage appends a page to the assistant.
//
// The function takes the following parameters:
//
//    - page: GtkWidget.
//
// The function returns the following values:
//
//    - gint: index (starting at 0) of the inserted page.
//
func (assistant *Assistant) AppendPage(page Widgetter) int32 {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("append_page", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Commit erases the visited page history.
//
// GTK will then hide the back button on the current page, and removes the
// cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is hereafter
// deemed permanent and cannot be modified or undone. For example, showing a
// progress page to track a long-running, unreversible operation after the user
// has clicked apply on a confirmation page.
func (assistant *Assistant) Commit() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("commit", _args[:], nil)

	runtime.KeepAlive(assistant)
}

// CurrentPage returns the page number of the current page.
//
// The function returns the following values:
//
//    - gint: index (starting from 0) of the current page in the assistant, or -1
//      if the assistant has no pages, or no current page.
//
func (assistant *Assistant) CurrentPage() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_current_page", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NPages returns the number of pages in the assistant.
//
// The function returns the following values:
//
//    - gint: number of pages in the assistant.
//
func (assistant *Assistant) NPages() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_n_pages", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number page_num.
//
// The function takes the following parameters:
//
//    - pageNum: index of a page in the assistant, or -1 to get the last page.
//
// The function returns the following values:
//
//    - widget (optional): child widget, or NULL if page_num is out of bounds.
//
func (assistant *Assistant) NthPage(pageNum int32) Widgetter {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.int(pageNum)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_nth_page", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Page returns the GtkAssistantPage object for child.
//
// The function takes the following parameters:
//
//    - child of assistant.
//
// The function returns the following values:
//
//    - assistantPage: GtkAssistantPage for child.
//
func (assistant *Assistant) Page(child Widgetter) *AssistantPage {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)

	var _assistantPage *AssistantPage // out

	_assistantPage = wrapAssistantPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _assistantPage
}

// PageComplete gets whether page is complete.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - ok: TRUE if page is complete.
//
func (assistant *Assistant) PageComplete(page Widgetter) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_complete", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageTitle gets the title for page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - utf8: title for page.
//
func (assistant *Assistant) PageTitle(page Widgetter) string {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_title", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Pages gets a list model of the assistant pages.
//
// The function returns the following values:
//
//    - listModel: list model of the pages.
//
func (assistant *Assistant) Pages() *gio.ListModel {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_pages", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// InsertPage inserts a page in the assistant at a given position.
//
// The function takes the following parameters:
//
//    - page: GtkWidget.
//    - position: index (starting at 0) at which to insert the page, or -1 to
//      append the page to the assistant.
//
// The function returns the following values:
//
//    - gint: index (starting from 0) of the inserted page.
//
func (assistant *Assistant) InsertPage(page Widgetter, position int32) int32 {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.int   // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = C.int(position)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("insert_page", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(position)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the GTK_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) NextPage() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("next_page", _args[:], nil)

	runtime.KeepAlive(assistant)
}

// PrependPage prepends a page to the assistant.
//
// The function takes the following parameters:
//
//    - page: GtkWidget.
//
// The function returns the following values:
//
//    - gint: index (starting at 0) of the inserted page.
//
func (assistant *Assistant) PrependPage(page Widgetter) int32 {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("prepend_page", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the GTK_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) PreviousPage() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("previous_page", _args[:], nil)

	runtime.KeepAlive(assistant)
}

// RemoveActionWidget removes a widget from the action area of a GtkAssistant.
//
// The function takes the following parameters:
//
//    - child: GtkWidget.
//
func (assistant *Assistant) RemoveActionWidget(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("remove_action_widget", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// RemovePage removes the page_num’s page from assistant.
//
// The function takes the following parameters:
//
//    - pageNum: index of a page in the assistant, or -1 to remove the last page.
//
func (assistant *Assistant) RemovePage(pageNum int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.int(pageNum)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("remove_page", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetCurrentPage switches the page to page_num.
//
// Note that this will only be necessary in custom buttons, as the assistant
// flow can be set with gtk_assistant_set_forward_page_func().
//
// The function takes the following parameters:
//
//    - pageNum: index of the page to switch to, starting from 0. If negative,
//      the last page will be used. If greater than the number of pages in the
//      assistant, nothing will be done.
//
func (assistant *Assistant) SetCurrentPage(pageNum int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.int(pageNum)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_current_page", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetForwardPageFunc sets the page forwarding function to be page_func.
//
// This function will be used to determine what will be the next page when the
// user presses the forward button. Setting page_func to NULL will make the
// assistant to use the default forward function, which just goes to the next
// visible page.
//
// The function takes the following parameters:
//
//    - pageFunc (optional): GtkAssistantPageFunc, or NULL to use the default
//      one.
//
func (assistant *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc) {
	var _args [4]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gpointer // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	if pageFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_AssistantPageFunc)
		_arg2 = C.gpointer(gbox.Assign(pageFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_forward_page_func", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageFunc)
}

// SetPageComplete sets whether page contents are complete.
//
// This will make assistant update the buttons state to be able to continue the
// task.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - complete completeness status of the page.
//
func (assistant *Assistant) SetPageComplete(page Widgetter, complete bool) {
	var _args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if complete {
		_arg2 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gboolean)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_complete", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(complete)
}

// SetPageTitle sets a title for page.
//
// The title is displayed in the header area of the assistant when page is the
// current page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - title: new title for page.
//
func (assistant *Assistant) SetPageTitle(page Widgetter, title string) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_title", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(title)
}

// UpdateButtonsState forces assistant to recompute the buttons state.
//
// GTK automatically takes care of this in most situations, e.g. when the user
// goes to a different page, or when the visibility or completeness of a page
// changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (assistant *Assistant) UpdateButtonsState() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("update_buttons_state", _args[:], nil)

	runtime.KeepAlive(assistant)
}

// AssistantPage: GtkAssistantPage is an auxiliary object used by `GtkAssistant.
type AssistantPage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AssistantPage)(nil)
)

func wrapAssistantPage(obj *coreglib.Object) *AssistantPage {
	return &AssistantPage{
		Object: obj,
	}
}

func marshalAssistantPage(p uintptr) (interface{}, error) {
	return wrapAssistantPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child returns the child to which page belongs.
//
// The function returns the following values:
//
//    - widget: child to which page belongs.
//
func (page *AssistantPage) Child() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AssistantPage").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(page)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}
