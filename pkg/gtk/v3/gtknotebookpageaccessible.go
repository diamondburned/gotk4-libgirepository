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
		{T: externglib.Type(C.gtk_notebook_page_accessible_get_type()), F: marshalNotebookPageAccessibler},
	})
}

// NotebookPageAccessibler describes NotebookPageAccessible's methods.
type NotebookPageAccessibler interface {
	//
	Invalidate()
}

//
type NotebookPageAccessible struct {
	atk.ObjectClass

	atk.Component
}

var (
	_ NotebookPageAccessibler = (*NotebookPageAccessible)(nil)
	_ gextras.Nativer         = (*NotebookPageAccessible)(nil)
)

func wrapNotebookPageAccessible(obj *externglib.Object) NotebookPageAccessibler {
	return &NotebookPageAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalNotebookPageAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNotebookPageAccessible(obj), nil
}

//
func NewNotebookPageAccessible(notebook NotebookAccessibler, child Widgetter) *NotebookPageAccessible {
	var _arg1 *C.GtkNotebookAccessible // out
	var _arg2 *C.GtkWidget             // out
	var _cret *C.AtkObject             // in

	_arg1 = (*C.GtkNotebookAccessible)(unsafe.Pointer((notebook).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_notebook_page_accessible_new(_arg1, _arg2)

	var _notebookPageAccessible *NotebookPageAccessible // out

	_notebookPageAccessible = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NotebookPageAccessible)

	return _notebookPageAccessible
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *NotebookPageAccessible) Native() uintptr {
	return v.ObjectClass.Object.Native()
}

//
func (page *NotebookPageAccessible) Invalidate() {
	var _arg0 *C.GtkNotebookPageAccessible // out

	_arg0 = (*C.GtkNotebookPageAccessible)(unsafe.Pointer(page.Native()))

	C.gtk_notebook_page_accessible_invalidate(_arg0)
}
