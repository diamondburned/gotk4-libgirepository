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
import "C"

// GType values.
var (
	GTypeFileChooserAction = coreglib.Type(girepository.MustFind("Gtk", "FileChooserAction").RegisteredGType())
	GTypeFileChooserError  = coreglib.Type(girepository.MustFind("Gtk", "FileChooserError").RegisteredGType())
	GTypeFileChooser       = coreglib.Type(girepository.MustFind("Gtk", "FileChooser").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileChooserAction, F: marshalFileChooserAction},
		coreglib.TypeMarshaler{T: GTypeFileChooserError, F: marshalFileChooserError},
		coreglib.TypeMarshaler{T: GTypeFileChooser, F: marshalFileChooser},
	})
}

// FileChooserAction describes whether a GtkFileChooser is being used to open
// existing files or to save to a possibly new file.
type FileChooserAction C.gint

const (
	// FileChooserActionOpen indicates open mode. The file chooser will only let
	// the user pick an existing file.
	FileChooserActionOpen FileChooserAction = iota
	// FileChooserActionSave indicates save mode. The file chooser will let the
	// user pick an existing file, or type in a new filename.
	FileChooserActionSave
	// FileChooserActionSelectFolder indicates an Open mode for selecting
	// folders. The file chooser will let the user pick an existing folder.
	FileChooserActionSelectFolder
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	return FileChooserAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileChooserAction.
func (f FileChooserAction) String() string {
	switch f {
	case FileChooserActionOpen:
		return "Open"
	case FileChooserActionSave:
		return "Save"
	case FileChooserActionSelectFolder:
		return "SelectFolder"
	default:
		return fmt.Sprintf("FileChooserAction(%d)", f)
	}
}

// FileChooserError: these identify the various errors that can occur while
// calling GtkFileChooser functions.
type FileChooserError C.gint

const (
	// FileChooserErrorNonexistent indicates that a file does not exist.
	FileChooserErrorNonexistent FileChooserError = iota
	// FileChooserErrorBadFilename indicates a malformed filename.
	FileChooserErrorBadFilename
	// FileChooserErrorAlreadyExists indicates a duplicate path (e.g. when
	// adding a bookmark).
	FileChooserErrorAlreadyExists
	// FileChooserErrorIncompleteHostname indicates an incomplete hostname (e.g.
	// "http://foo" without a slash after that).
	FileChooserErrorIncompleteHostname
)

func marshalFileChooserError(p uintptr) (interface{}, error) {
	return FileChooserError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileChooserError.
func (f FileChooserError) String() string {
	switch f {
	case FileChooserErrorNonexistent:
		return "Nonexistent"
	case FileChooserErrorBadFilename:
		return "BadFilename"
	case FileChooserErrorAlreadyExists:
		return "AlreadyExists"
	case FileChooserErrorIncompleteHostname:
		return "IncompleteHostname"
	default:
		return fmt.Sprintf("FileChooserError(%d)", f)
	}
}

// FileChooser: GtkFileChooser is an interface that can be implemented by file
// selection widgets.
//
// In GTK, the main objects that implement this interface are
// gtk.FileChooserWidget and gtk.FileChooserDialog.
//
// You do not need to write an object that implements the GtkFileChooser
// interface unless you are trying to adapt an existing file selector to expose
// a standard programming interface.
//
// GtkFileChooser allows for shortcuts to various places in the filesystem. In
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
// When the user is finished selecting files in a GtkFileChooser, your program
// can get the selected filenames as GFiles.
//
//
// Adding options
//
// You can add extra widgets to a file chooser to provide options that are not
// present in the default design, by using gtk.FileChooser.AddChoice(). Each
// choice has an identifier and a user visible label; additionally, each choice
// can have multiple options. If a choice has no option, it will be rendered as
// a check button with the given label; if a choice has options, it will be
// rendered as a combo box.
//
// FileChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FileChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileChooser)(nil)
)

// FileChooserer describes FileChooser's interface methods.
type FileChooserer interface {
	coreglib.Objector

	baseFileChooser() *FileChooser
}

var _ FileChooserer = (*FileChooser)(nil)

func wrapFileChooser(obj *coreglib.Object) *FileChooser {
	return &FileChooser{
		Object: obj,
	}
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	return wrapFileChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *FileChooser) baseFileChooser() *FileChooser {
	return v
}

// BaseFileChooser returns the underlying base object.
func BaseFileChooser(obj FileChooserer) *FileChooser {
	return obj.baseFileChooser()
}
