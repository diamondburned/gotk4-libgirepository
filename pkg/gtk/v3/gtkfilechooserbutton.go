// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_FileChooserButtonClass_file_set(GtkFileChooserButton*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_button_get_type()), F: marshalFileChooserButtonner},
	})
}

// FileChooserButtonOverrider contains methods that are overridable.
type FileChooserButtonOverrider interface {
	FileSet()
}

// FileChooserButton is a widget that lets the user select a file. It implements
// the FileChooser interface. Visually, it is a file name with a button to bring
// up a FileChooserDialog. The user can then use that dialog to change the file
// associated with that button. This widget does not support setting the
// FileChooser:select-multiple property to TRUE.
//
// Create a button to let the user select a file in /etc
//
//    {
//      GtkWidget *button;
//
//      button = gtk_file_chooser_button_new (_("Select a file"),
//                                            GTK_FILE_CHOOSER_ACTION_OPEN);
//      gtk_file_chooser_set_current_folder (GTK_FILE_CHOOSER (button),
//                                           "/etc");
//    }
//
// The FileChooserButton supports the FileChooserActions
// GTK_FILE_CHOOSER_ACTION_OPEN and GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
//
// > The FileChooserButton will ellipsize the label, and will thus > request
// little horizontal space. To give the button more space, > you should call
// gtk_widget_get_preferred_size(), > gtk_file_chooser_button_set_width_chars(),
// or pack the button in > such a way that other interface elements give space
// to the > widget.
//
//
// CSS nodes
//
// GtkFileChooserButton has a CSS node with name “filechooserbutton”, containing
// a subnode for the internal button with name “button” and style class “.file”.
type FileChooserButton struct {
	_ [0]func() // equal guard
	Box

	*externglib.Object
	FileChooser
}

var (
	_ externglib.Objector = (*FileChooserButton)(nil)
	_ Containerer         = (*FileChooserButton)(nil)
)

func classInitFileChooserButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkFileChooserButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkFileChooserButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ FileSet() }); ok {
		pclass.file_set = (*[0]byte)(C._gotk4_gtk3_FileChooserButtonClass_file_set)
	}
}

//export _gotk4_gtk3_FileChooserButtonClass_file_set
func _gotk4_gtk3_FileChooserButtonClass_file_set(arg0 *C.GtkFileChooserButton) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FileSet() })

	iface.FileSet()
}

func wrapFileChooserButton(obj *externglib.Object) *FileChooserButton {
	return &FileChooserButton{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
		Object: obj,
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserButtonner(p uintptr) (interface{}, error) {
	return wrapFileChooserButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectFileSet signal is emitted when the user selects a file.
//
// Note that this signal is only emitted when the user changes the file.
func (button *FileChooserButton) ConnectFileSet(f func()) externglib.SignalHandle {
	return button.Connect("file-set", externglib.GeneratedClosure{Func: f})
}

// NewFileChooserButton creates a new file-selecting button widget.
//
// The function takes the following parameters:
//
//    - title of the browse dialog.
//    - action: open mode for the widget.
//
// The function returns the following values:
//
//    - fileChooserButton: new button widget.
//
func NewFileChooserButton(title string, action FileChooserAction) *FileChooserButton {
	var _arg1 *C.gchar               // out
	var _arg2 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_button_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(action)

	var _fileChooserButton *FileChooserButton // out

	_fileChooserButton = wrapFileChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserButton
}

// NewFileChooserButtonWithDialog creates a FileChooserButton widget which uses
// dialog as its file-picking window.
//
// Note that dialog must be a Dialog (or subclass) which implements the
// FileChooser interface and must not have GTK_DIALOG_DESTROY_WITH_PARENT set.
//
// Also note that the dialog needs to have its confirmative button added with
// response GTK_RESPONSE_ACCEPT or GTK_RESPONSE_OK in order for the button to
// take over the file selected in the dialog.
//
// The function takes the following parameters:
//
//    - dialog: widget to use as dialog.
//
// The function returns the following values:
//
//    - fileChooserButton: new button widget.
//
func NewFileChooserButtonWithDialog(dialog *Dialog) *FileChooserButton {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_file_chooser_button_new_with_dialog(_arg1)
	runtime.KeepAlive(dialog)

	var _fileChooserButton *FileChooserButton // out

	_fileChooserButton = wrapFileChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserButton
}

// FocusOnClick returns whether the button grabs focus when it is clicked with
// the mouse. See gtk_file_chooser_button_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
//
// The function returns the following values:
//
//    - ok: TRUE if the button grabs focus when it is clicked with the mouse.
//
func (button *FileChooserButton) FocusOnClick() bool {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_focus_on_click(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the browse dialog used by button. The returned
// value should not be modified or freed.
//
// The function returns the following values:
//
//    - utf8: pointer to the browse dialog’s title.
//
func (button *FileChooserButton) Title() string {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret *C.gchar                // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_title(_arg0)
	runtime.KeepAlive(button)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WidthChars retrieves the width in characters of the button widget’s entry
// and/or label.
//
// The function returns the following values:
//
//    - gint: integer width (in characters) that the button will use to size
//      itself.
//
func (button *FileChooserButton) WidthChars() int {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gint                  // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_width_chars(_arg0)
	runtime.KeepAlive(button)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetFocusOnClick sets whether the button will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
//
// The function takes the following parameters:
//
//    - focusOnClick: whether the button grabs focus when clicked with the mouse.
//
func (button *FileChooserButton) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_button_set_focus_on_click(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(focusOnClick)
}

// SetTitle modifies the title of the browse dialog used by button.
//
// The function takes the following parameters:
//
//    - title: new browse dialog title.
//
func (button *FileChooserButton) SetTitle(title string) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_button_set_title(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(title)
}

// SetWidthChars sets the width (in characters) that button will use to n_chars.
//
// The function takes the following parameters:
//
//    - nChars: new width, in characters.
//
func (button *FileChooserButton) SetWidthChars(nChars int) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gint                  // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_file_chooser_button_set_width_chars(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(nChars)
}
