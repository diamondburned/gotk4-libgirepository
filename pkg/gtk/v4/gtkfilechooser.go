// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_action_get_type()), F: marshalFileChooserAction},
		{T: externglib.Type(C.gtk_file_chooser_error_get_type()), F: marshalFileChooserError},
		{T: externglib.Type(C.gtk_file_chooser_get_type()), F: marshalFileChooser},
	})
}

// FileChooserAction describes whether a `GtkFileChooser` is being used to open
// existing files or to save to a possibly new file.
type FileChooserAction int

const (
	// Open indicates open mode. The file chooser will only let the user pick an
	// existing file.
	FileChooserActionOpen FileChooserAction = iota
	// Save indicates save mode. The file chooser will let the user pick an
	// existing file, or type in a new filename.
	FileChooserActionSave
	// SelectFolder indicates an Open mode for selecting folders. The file
	// chooser will let the user pick an existing folder.
	FileChooserActionSelectFolder
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	return FileChooserAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FileChooserError: these identify the various errors that can occur while
// calling `GtkFileChooser` functions.
type FileChooserError int

const (
	// Nonexistent indicates that a file does not exist.
	FileChooserErrorNonexistent FileChooserError = iota
	// BadFilename indicates a malformed filename.
	FileChooserErrorBadFilename
	// AlreadyExists indicates a duplicate path (e.g. when adding a bookmark).
	FileChooserErrorAlreadyExists
	// IncompleteHostname indicates an incomplete hostname (e.g. "http://foo"
	// without a slash after that).
	FileChooserErrorIncompleteHostname
)

func marshalFileChooserError(p uintptr) (interface{}, error) {
	return FileChooserError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FileChooser: `GtkFileChooser` is an interface that can be implemented by file
// selection widgets.
//
// In GTK, the main objects that implement this interface are
// [class@Gtk.FileChooserWidget] and [class@Gtk.FileChooserDialog].
//
// You do not need to write an object that implements the `GtkFileChooser`
// interface unless you are trying to adapt an existing file selector to expose
// a standard programming interface.
//
// `GtkFileChooser` allows for shortcuts to various places in the filesystem. In
// the default implementation these are displayed in the left pane. It may be a
// bit confusing at first that these shortcuts come from various sources and in
// various flavours, so lets explain the terminology here:
//
// - Bookmarks: are created by the user, by dragging folders from the right pane
// to the left pane, or by using the “Add”. Bookmarks can be renamed and deleted
// by the user.
//
// - Shortcuts: can be provided by the application. For example, a Paint program
// may want to add a shortcut for a Clipart folder. Shortcuts cannot be modified
// by the user.
//
// - Volumes: are provided by the underlying filesystem abstraction. They are
// the “roots” of the filesystem.
//
//
// File Names and Encodings
//
// When the user is finished selecting files in a `GtkFileChooser`, your program
// can get the selected filenames as `GFile`s.
//
//
// Adding options
//
// You can add extra widgets to a file chooser to provide options that are not
// present in the default design, by using [method@Gtk.FileChooser.add_choice].
// Each choice has an identifier and a user visible label; additionally, each
// choice can have multiple options. If a choice has no option, it will be
// rendered as a check button with the given label; if a choice has options, it
// will be rendered as a combo box.
type FileChooser interface {
	gextras.Objector

	// AddChoice adds a 'choice' to the file chooser.
	//
	// This is typically implemented as a combobox or, for boolean choices, as a
	// checkbutton. You can select a value using
	// [method@Gtk.FileChooser.set_choice] before the dialog is shown, and you
	// can obtain the user-selected value in the [signal@Gtk.Dialog::response]
	// signal handler using [method@Gtk.FileChooser.get_choice].
	AddChoice(id string, label string, options []string, optionLabels []string)
	// AddFilter adds @filter to the list of filters that the user can select
	// between.
	//
	// When a filter is selected, only files that are passed by that filter are
	// displayed.
	//
	// Note that the @chooser takes ownership of the filter if it is floating,
	// so you have to ref and sink it if you want to keep a reference.
	AddFilter(filter FileFilter)
	// Action gets the type of operation that the file chooser is performing.
	Action() FileChooserAction
	// Choice gets the currently selected option in the 'choice' with the given
	// ID.
	Choice(id string) string
	// CreateFolders gets whether file chooser will offer to create new folders.
	CreateFolders() bool
	// CurrentName gets the current name in the file selector, as entered by the
	// user.
	//
	// This is meant to be used in save dialogs, to get the currently typed
	// filename when the file itself does not exist yet.
	CurrentName() string
	// Filter gets the current filter.
	Filter() FileFilter
	// SelectMultiple gets whether multiple files can be selected in the file
	// chooser.
	SelectMultiple() bool
	// RemoveChoice removes a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	RemoveChoice(id string)
	// RemoveFilter removes @filter from the list of filters that the user can
	// select between.
	RemoveFilter(filter FileFilter)
	// SetAction sets the type of operation that the chooser is performing.
	//
	// The user interface is adapted to suit the selected action.
	//
	// For example, an option to create a new folder might be shown if the
	// action is GTK_FILE_CHOOSER_ACTION_SAVE but not if the action is
	// GTK_FILE_CHOOSER_ACTION_OPEN.
	SetAction(action FileChooserAction)
	// SetChoice selects an option in a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	//
	// For a boolean choice, the possible options are "true" and "false".
	SetChoice(id string, option string)
	// SetCreateFolders sets whether file chooser will offer to create new
	// folders.
	//
	// This is only relevant if the action is not set to be
	// GTK_FILE_CHOOSER_ACTION_OPEN.
	SetCreateFolders(createFolders bool)
	// SetCurrentName sets the current name in the file selector, as if entered
	// by the user.
	//
	// Note that the name passed in here is a UTF-8 string rather than a
	// filename. This function is meant for such uses as a suggested name in a
	// “Save As...” dialog. You can pass “Untitled.doc” or a similarly suitable
	// suggestion for the @name.
	//
	// If you want to preselect a particular existing file, you should use
	// [method@Gtk.FileChooser.set_file] instead.
	//
	// Please see the documentation for those functions for an example of using
	// [method@Gtk.FileChooser.set_current_name] as well.
	SetCurrentName(name string)
	// SetFilter sets the current filter.
	//
	// Only the files that pass the filter will be displayed. If the
	// user-selectable list of filters is non-empty, then the filter should be
	// one of the filters in that list.
	//
	// Setting the current filter when the list of filters is empty is useful if
	// you want to restrict the displayed set of files without letting the user
	// change it.
	SetFilter(filter FileFilter)
	// SetSelectMultiple sets whether multiple files can be selected in the file
	// chooser.
	//
	// This is only relevant if the action is set to be
	// GTK_FILE_CHOOSER_ACTION_OPEN or GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
	SetSelectMultiple(selectMultiple bool)
}

// fileChooser implements the FileChooser interface.
type fileChooser struct {
	*externglib.Object
}

var _ FileChooser = (*fileChooser)(nil)

// WrapFileChooser wraps a GObject to a type that implements
// interface FileChooser. It is primarily used internally.
func WrapFileChooser(obj *externglib.Object) FileChooser {
	return fileChooser{obj}
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooser(obj), nil
}

func (c fileChooser) AddChoice(id string, label string, options []string, optionLabels []string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out
	var _arg3 **C.char
	var _arg4 **C.char

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (**C.char)(C.malloc(C.ulong(len(options)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(options))
		for i := range options {
			out[i] = (*C.char)(C.CString(options[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg4 = (**C.char)(C.malloc(C.ulong(len(optionLabels)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(optionLabels))
		for i := range optionLabels {
			out[i] = (*C.char)(C.CString(optionLabels[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_file_chooser_add_choice(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (c fileChooser) AddFilter(filter FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_add_filter(_arg0, _arg1)
}

func (c fileChooser) Action() FileChooserAction {
	var _arg0 *C.GtkFileChooser      // out
	var _cret C.GtkFileChooserAction // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_file_chooser_get_action(_arg0)

	var _fileChooserAction FileChooserAction // out

	_fileChooserAction = FileChooserAction(_cret)

	return _fileChooserAction
}

func (c fileChooser) Choice(id string) string {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_file_chooser_get_choice(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c fileChooser) CreateFolders() bool {
	var _arg0 *C.GtkFileChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_file_chooser_get_create_folders(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c fileChooser) CurrentName() string {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_file_chooser_get_current_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (c fileChooser) Filter() FileFilter {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GtkFileFilter  // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_file_chooser_get_filter(_arg0)

	var _fileFilter FileFilter // out

	_fileFilter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FileFilter)

	return _fileFilter
}

func (c fileChooser) SelectMultiple() bool {
	var _arg0 *C.GtkFileChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_file_chooser_get_select_multiple(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c fileChooser) RemoveChoice(id string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_remove_choice(_arg0, _arg1)
}

func (c fileChooser) RemoveFilter(filter FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_remove_filter(_arg0, _arg1)
}

func (c fileChooser) SetAction(action FileChooserAction) {
	var _arg0 *C.GtkFileChooser      // out
	var _arg1 C.GtkFileChooserAction // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkFileChooserAction(action)

	C.gtk_file_chooser_set_action(_arg0, _arg1)
}

func (c fileChooser) SetChoice(id string, option string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(option))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_file_chooser_set_choice(_arg0, _arg1, _arg2)
}

func (c fileChooser) SetCreateFolders(createFolders bool) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	if createFolders {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_set_create_folders(_arg0, _arg1)
}

func (c fileChooser) SetCurrentName(name string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_set_current_name(_arg0, _arg1)
}

func (c fileChooser) SetFilter(filter FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_set_filter(_arg0, _arg1)
}

func (c fileChooser) SetSelectMultiple(selectMultiple bool) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(c.Native()))
	if selectMultiple {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_set_select_multiple(_arg0, _arg1)
}
