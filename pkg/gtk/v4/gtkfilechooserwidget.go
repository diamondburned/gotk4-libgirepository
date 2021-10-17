// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidgetter},
	})
}

// FileChooserWidget: GtkFileChooserWidget is a widget for choosing files.
//
// It exposes the gtk.FileChooser interface, and you should use the methods of
// this interface to interact with the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget struct {
	Widget

	FileChooser
	*externglib.Object
}

func wrapFileChooserWidget(obj *externglib.Object) *FileChooserWidget {
	return &FileChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		FileChooser: FileChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFileChooserWidgetter(p uintptr) (interface{}, error) {
	return wrapFileChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileChooserWidget creates a new GtkFileChooserWidget.
//
// This is a file chooser widget that can be embedded in custom windows, and it
// is the same widget that is used by GtkFileChooserDialog.
//
// The function takes the following parameters:
//
//    - action: open or save mode for the widget.
//
func NewFileChooserWidget(action FileChooserAction) *FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_widget_new(_arg1)
	runtime.KeepAlive(action)

	var _fileChooserWidget *FileChooserWidget // out

	_fileChooserWidget = wrapFileChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserWidget
}

func (*FileChooserWidget) privateFileChooserWidget() {}

// ConnectDesktopFolder: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the user's Desktop folder in the
// file list.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>D</kbd>.
func (v *FileChooserWidget) ConnectDesktopFolder(f func()) externglib.SignalHandle {
	return v.Connect("desktop-folder", f)
}

// ConnectDownFolder: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser go to a child of the current folder in
// the file hierarchy. The subfolder that will be used is displayed in the path
// bar widget of the file chooser. For example, if the path bar is showing
// "/foo/bar/baz", with bar currently displayed, then this will cause the file
// chooser to switch to the "baz" subfolder.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Down</kbd>.
func (v *FileChooserWidget) ConnectDownFolder(f func()) externglib.SignalHandle {
	return v.Connect("down-folder", f)
}

// ConnectHomeFolder: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the user's home folder in the file
// list.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Home</kbd>.
func (v *FileChooserWidget) ConnectHomeFolder(f func()) externglib.SignalHandle {
	return v.Connect("home-folder", f)
}

// ConnectLocationPopup: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default bindings for this signal are <kbd>Control</kbd>-<kbd>L</kbd> with
// a path string of "" (the empty string). It is also bound to <kbd>/</kbd> with
// a path string of "/" (a slash): this lets you type / and immediately type a
// path name. On Unix systems, this is bound to <kbd>~</kbd> (tilde) with a path
// string of "~" itself for access to home directories.
func (v *FileChooserWidget) ConnectLocationPopup(f func(path string)) externglib.SignalHandle {
	return v.Connect("location-popup", f)
}

// ConnectLocationPopupOnPaste: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show a "Location" prompt when the user
// pastes into a GtkFileChooserWidget.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>V</kbd>.
func (v *FileChooserWidget) ConnectLocationPopupOnPaste(f func()) externglib.SignalHandle {
	return v.Connect("location-popup-on-paste", f)
}

// ConnectLocationTogglePopup: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to toggle the visibility of a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>L</kbd>.
func (v *FileChooserWidget) ConnectLocationTogglePopup(f func()) externglib.SignalHandle {
	return v.Connect("location-toggle-popup", f)
}

// ConnectPlacesShortcut: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to move the focus to the places sidebar.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>P</kbd>.
func (v *FileChooserWidget) ConnectPlacesShortcut(f func()) externglib.SignalHandle {
	return v.Connect("places-shortcut", f)
}

// ConnectQuickBookmark: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser switch to the bookmark specified in the
// bookmark_index parameter. For example, if you have three bookmarks, you can
// pass 0, 1, 2 to this signal to switch to each of them, respectively.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>1</kbd>,
// <kbd>Alt</kbd>-<kbd>2</kbd>, etc. until <kbd>Alt</kbd>-<kbd>0</kbd>. Note
// that in the default binding, that <kbd>Alt</kbd>-<kbd>1</kbd> is actually
// defined to switch to the bookmark at index 0, and so on successively.
func (v *FileChooserWidget) ConnectQuickBookmark(f func(bookmarkIndex int)) externglib.SignalHandle {
	return v.Connect("quick-bookmark", f)
}

// ConnectRecentShortcut: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the Recent location.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>R</kbd>.
func (v *FileChooserWidget) ConnectRecentShortcut(f func()) externglib.SignalHandle {
	return v.Connect("recent-shortcut", f)
}

// ConnectSearchShortcut: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the search entry.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>S</kbd>.
func (v *FileChooserWidget) ConnectSearchShortcut(f func()) externglib.SignalHandle {
	return v.Connect("search-shortcut", f)
}

// ConnectShowHidden: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser display hidden files.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>H</kbd>.
func (v *FileChooserWidget) ConnectShowHidden(f func()) externglib.SignalHandle {
	return v.Connect("show-hidden", f)
}

// ConnectUpFolder: emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser go to the parent of the current folder
// in the file hierarchy.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Up</kbd>.
func (v *FileChooserWidget) ConnectUpFolder(f func()) externglib.SignalHandle {
	return v.Connect("up-folder", f)
}
