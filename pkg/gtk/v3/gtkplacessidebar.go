// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_places_open_flags_get_type()), F: marshalPlacesOpenFlags},
		{T: externglib.Type(C.gtk_places_sidebar_get_type()), F: marshalPlacesSidebarrer},
	})
}

// PlacesOpenFlags: these flags serve two purposes. First, the application can
// call gtk_places_sidebar_set_open_flags() using these flags as a bitmask. This
// tells the sidebar that the application is able to open folders selected from
// the sidebar in various ways, for example, in new tabs or in new windows in
// addition to the normal mode.
//
// Second, when one of these values gets passed back to the application in the
// PlacesSidebar::open-location signal, it means that the application should
// open the selected location in the normal way, in a new tab, or in a new
// window. The sidebar takes care of determining the desired way to open the
// location, based on the modifier keys that the user is pressing at the time
// the selection is made.
//
// If the application never calls gtk_places_sidebar_set_open_flags(), then the
// sidebar will only use K_PLACES_OPEN_NORMAL in the
// PlacesSidebar::open-location signal. This is the default mode of operation.
type PlacesOpenFlags int

const (
	// PlacesOpenFlagsNormal: this is the default mode that PlacesSidebar uses
	// if no other flags are specified. It indicates that the calling
	// application should open the selected location in the normal way, for
	// example, in the folder view beside the sidebar.
	PlacesOpenFlagsNormal PlacesOpenFlags = 0b1
	// PlacesOpenFlagsNewTab: when passed to
	// gtk_places_sidebar_set_open_flags(), this indicates that the application
	// can open folders selected from the sidebar in new tabs. This value will
	// be passed to the PlacesSidebar::open-location signal when the user
	// selects that a location be opened in a new tab instead of in the standard
	// fashion.
	PlacesOpenFlagsNewTab PlacesOpenFlags = 0b10
	// PlacesOpenFlagsNewWindow: similar to @GTK_PLACES_OPEN_NEW_TAB, but
	// indicates that the application can open folders in new windows.
	PlacesOpenFlagsNewWindow PlacesOpenFlags = 0b100
)

func marshalPlacesOpenFlags(p uintptr) (interface{}, error) {
	return PlacesOpenFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PlacesSidebarrer describes PlacesSidebar's methods.
type PlacesSidebarrer interface {
	// AddShortcut applications may want to present some folders in the places
	// sidebar if they could be immediately useful to users.
	AddShortcut(location gio.Filer)
	// LocalOnly returns the value previously set with
	// gtk_places_sidebar_set_local_only().
	LocalOnly() bool
	// Location gets the currently selected location in the @sidebar.
	Location() *gio.File
	// NthBookmark: this function queries the bookmarks added by the user to the
	// places sidebar, and returns one of them.
	NthBookmark(n int) *gio.File
	// OpenFlags gets the open flags.
	OpenFlags() PlacesOpenFlags
	// ShowConnectToServer returns the value previously set with
	// gtk_places_sidebar_set_show_connect_to_server() Deprecated: It is
	// recommended to group this functionality with the drives and network
	// location under the new 'Other Location' item.
	ShowConnectToServer() bool
	// ShowDesktop returns the value previously set with
	// gtk_places_sidebar_set_show_desktop()
	ShowDesktop() bool
	// ShowEnterLocation returns the value previously set with
	// gtk_places_sidebar_set_show_enter_location()
	ShowEnterLocation() bool
	// ShowOtherLocations returns the value previously set with
	// gtk_places_sidebar_set_show_other_locations()
	ShowOtherLocations() bool
	// ShowRecent returns the value previously set with
	// gtk_places_sidebar_set_show_recent()
	ShowRecent() bool
	// ShowStarredLocation returns the value previously set with
	// gtk_places_sidebar_set_show_starred_location()
	ShowStarredLocation() bool
	// ShowTrash returns the value previously set with
	// gtk_places_sidebar_set_show_trash()
	ShowTrash() bool
	// RemoveShortcut removes an application-specific shortcut that has been
	// previously been inserted with gtk_places_sidebar_add_shortcut().
	RemoveShortcut(location gio.Filer)
	// SetDropTargetsVisible: make the GtkPlacesSidebar show drop targets, so it
	// can show the available drop targets and a "new bookmark" row.
	SetDropTargetsVisible(visible bool, context gdk.DragContexter)
	// SetLocalOnly sets whether the @sidebar should only show local files.
	SetLocalOnly(localOnly bool)
	// SetLocation sets the location that is being shown in the widgets
	// surrounding the @sidebar, for example, in a folder view in a file
	// manager.
	SetLocation(location gio.Filer)
	// SetShowConnectToServer sets whether the @sidebar should show an item for
	// connecting to a network server; this is off by default.
	SetShowConnectToServer(showConnectToServer bool)
	// SetShowDesktop sets whether the @sidebar should show an item for the
	// Desktop folder.
	SetShowDesktop(showDesktop bool)
	// SetShowEnterLocation sets whether the @sidebar should show an item for
	// entering a location; this is off by default.
	SetShowEnterLocation(showEnterLocation bool)
	// SetShowOtherLocations sets whether the @sidebar should show an item for
	// the application to show an Other Locations view; this is off by default.
	SetShowOtherLocations(showOtherLocations bool)
	// SetShowRecent sets whether the @sidebar should show an item for recent
	// files.
	SetShowRecent(showRecent bool)
	// SetShowStarredLocation: if you enable this, you should connect to the
	// PlacesSidebar::show-starred-location signal.
	SetShowStarredLocation(showStarredLocation bool)
	// SetShowTrash sets whether the @sidebar should show an item for the Trash
	// location.
	SetShowTrash(showTrash bool)
}

// PlacesSidebar is a widget that displays a list of frequently-used places in
// the file system: the user’s home directory, the user’s bookmarks, and volumes
// and drives. This widget is used as a sidebar in FileChooser and may be used
// by file managers and similar programs.
//
// The places sidebar displays drives and volumes, and will automatically mount
// or unmount them when the user selects them.
//
// Applications can hook to various signals in the places sidebar to customize
// its behavior. For example, they can add extra commands to the context menu of
// the sidebar.
//
// While bookmarks are completely in control of the user, the places sidebar
// also allows individual applications to provide extra shortcut folders that
// are unique to each application. For example, a Paint program may want to add
// a shortcut for a Clipart folder. You can do this with
// gtk_places_sidebar_add_shortcut().
//
// To make use of the places sidebar, an application at least needs to connect
// to the PlacesSidebar::open-location signal. This is emitted when the user
// selects in the sidebar a location to open. The application should also call
// gtk_places_sidebar_set_location() when it changes the currently-viewed
// location.
//
//
// CSS nodes
//
// GtkPlacesSidebar uses a single CSS node with name placessidebar and style
// class .sidebar.
//
// Among the children of the places sidebar, the following style classes can be
// used: - .sidebar-new-bookmark-row for the 'Add new bookmark' row -
// .sidebar-placeholder-row for a row that is a placeholder - .has-open-popup
// when a popup is open for a row
type PlacesSidebar struct {
	ScrolledWindow
}

var (
	_ PlacesSidebarrer = (*PlacesSidebar)(nil)
	_ gextras.Nativer  = (*PlacesSidebar)(nil)
)

func wrapPlacesSidebar(obj *externglib.Object) PlacesSidebarrer {
	return &PlacesSidebar{
		ScrolledWindow: ScrolledWindow{
			Bin: Bin{
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
			},
		},
	}
}

func marshalPlacesSidebarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPlacesSidebar(obj), nil
}

// NewPlacesSidebar creates a new PlacesSidebar widget.
//
// The application should connect to at least the PlacesSidebar::open-location
// signal to be notified when the user makes a selection in the sidebar.
func NewPlacesSidebar() *PlacesSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_places_sidebar_new()

	var _placesSidebar *PlacesSidebar // out

	_placesSidebar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PlacesSidebar)

	return _placesSidebar
}

// AddShortcut applications may want to present some folders in the places
// sidebar if they could be immediately useful to users. For example, a drawing
// program could add a “/usr/share/clipart” location when the sidebar is being
// used in an “Insert Clipart” dialog box.
//
// This function adds the specified @location to a special place for immutable
// shortcuts. The shortcuts are application-specific; they are not shared across
// applications, and they are not persistent. If this function is called
// multiple times with different locations, then they are added to the sidebar’s
// list in the same order as the function is called.
func (sidebar *PlacesSidebar) AddShortcut(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer((location).(gextras.Nativer).Native()))

	C.gtk_places_sidebar_add_shortcut(_arg0, _arg1)
}

// LocalOnly returns the value previously set with
// gtk_places_sidebar_set_local_only().
func (sidebar *PlacesSidebar) LocalOnly() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_local_only(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Location gets the currently selected location in the @sidebar. This can be
// nil when nothing is selected, for example, when
// gtk_places_sidebar_set_location() has been called with a location that is not
// among the sidebar’s list of places to show.
//
// You can use this function to get the selection in the @sidebar. Also, if you
// connect to the PlacesSidebar::populate-popup signal, you can use this
// function to get the location that is being referred to during the callbacks
// for your menu items.
func (sidebar *PlacesSidebar) Location() *gio.File {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_location(_arg0)

	var _file *gio.File // out

	_file = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.File)

	return _file
}

// NthBookmark: this function queries the bookmarks added by the user to the
// places sidebar, and returns one of them. This function is used by FileChooser
// to implement the “Alt-1”, “Alt-2”, etc. shortcuts, which activate the
// cooresponding bookmark.
func (sidebar *PlacesSidebar) NthBookmark(n int) *gio.File {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gint              // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = C.gint(n)

	_cret = C.gtk_places_sidebar_get_nth_bookmark(_arg0, _arg1)

	var _file *gio.File // out

	_file = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.File)

	return _file
}

// OpenFlags gets the open flags.
func (sidebar *PlacesSidebar) OpenFlags() PlacesOpenFlags {
	var _arg0 *C.GtkPlacesSidebar  // out
	var _cret C.GtkPlacesOpenFlags // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_open_flags(_arg0)

	var _placesOpenFlags PlacesOpenFlags // out

	_placesOpenFlags = PlacesOpenFlags(_cret)

	return _placesOpenFlags
}

// ShowConnectToServer returns the value previously set with
// gtk_places_sidebar_set_show_connect_to_server()
//
// Deprecated: It is recommended to group this functionality with the drives and
// network location under the new 'Other Location' item.
func (sidebar *PlacesSidebar) ShowConnectToServer() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_connect_to_server(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDesktop returns the value previously set with
// gtk_places_sidebar_set_show_desktop()
func (sidebar *PlacesSidebar) ShowDesktop() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_desktop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowEnterLocation returns the value previously set with
// gtk_places_sidebar_set_show_enter_location()
func (sidebar *PlacesSidebar) ShowEnterLocation() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_enter_location(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowOtherLocations returns the value previously set with
// gtk_places_sidebar_set_show_other_locations()
func (sidebar *PlacesSidebar) ShowOtherLocations() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_other_locations(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecent returns the value previously set with
// gtk_places_sidebar_set_show_recent()
func (sidebar *PlacesSidebar) ShowRecent() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_recent(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowStarredLocation returns the value previously set with
// gtk_places_sidebar_set_show_starred_location()
func (sidebar *PlacesSidebar) ShowStarredLocation() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_starred_location(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowTrash returns the value previously set with
// gtk_places_sidebar_set_show_trash()
func (sidebar *PlacesSidebar) ShowTrash() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_trash(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveShortcut removes an application-specific shortcut that has been
// previously been inserted with gtk_places_sidebar_add_shortcut(). If the
// @location is not a shortcut in the sidebar, then nothing is done.
func (sidebar *PlacesSidebar) RemoveShortcut(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer((location).(gextras.Nativer).Native()))

	C.gtk_places_sidebar_remove_shortcut(_arg0, _arg1)
}

// SetDropTargetsVisible: make the GtkPlacesSidebar show drop targets, so it can
// show the available drop targets and a "new bookmark" row. This improves the
// Drag-and-Drop experience of the user and allows applications to show all
// available drop targets at once.
//
// This needs to be called when the application is aware of an ongoing drag that
// might target the sidebar. The drop-targets-visible state will be unset
// automatically if the drag finishes in the GtkPlacesSidebar. You only need to
// unset the state when the drag ends on some other widget on your application.
func (sidebar *PlacesSidebar) SetDropTargetsVisible(visible bool, context gdk.DragContexter) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out
	var _arg2 *C.GdkDragContext   // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if visible {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GdkDragContext)(unsafe.Pointer((context).(gextras.Nativer).Native()))

	C.gtk_places_sidebar_set_drop_targets_visible(_arg0, _arg1, _arg2)
}

// SetLocalOnly sets whether the @sidebar should only show local files.
func (sidebar *PlacesSidebar) SetLocalOnly(localOnly bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if localOnly {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_local_only(_arg0, _arg1)
}

// SetLocation sets the location that is being shown in the widgets surrounding
// the @sidebar, for example, in a folder view in a file manager. In turn, the
// @sidebar will highlight that location if it is being shown in the list of
// places, or it will unhighlight everything if the @location is not among the
// places in the list.
func (sidebar *PlacesSidebar) SetLocation(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer((location).(gextras.Nativer).Native()))

	C.gtk_places_sidebar_set_location(_arg0, _arg1)
}

// SetShowConnectToServer sets whether the @sidebar should show an item for
// connecting to a network server; this is off by default. An application may
// want to turn this on if it implements a way for the user to connect to
// network servers directly.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-connect-to-server signal.
//
// Deprecated: It is recommended to group this functionality with the drives and
// network location under the new 'Other Location' item.
func (sidebar *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showConnectToServer {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_connect_to_server(_arg0, _arg1)
}

// SetShowDesktop sets whether the @sidebar should show an item for the Desktop
// folder. The default value for this option is determined by the desktop
// environment and the user’s configuration, but this function can be used to
// override it on a per-application basis.
func (sidebar *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showDesktop {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_desktop(_arg0, _arg1)
}

// SetShowEnterLocation sets whether the @sidebar should show an item for
// entering a location; this is off by default. An application may want to turn
// this on if manually entering URLs is an expected user action.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-enter-location signal.
func (sidebar *PlacesSidebar) SetShowEnterLocation(showEnterLocation bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showEnterLocation {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_enter_location(_arg0, _arg1)
}

// SetShowOtherLocations sets whether the @sidebar should show an item for the
// application to show an Other Locations view; this is off by default. When set
// to true, persistent devices such as hard drives are hidden, otherwise they
// are shown in the sidebar. An application may want to turn this on if it
// implements a way for the user to see and interact with drives and network
// servers directly.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-other-locations signal.
func (sidebar *PlacesSidebar) SetShowOtherLocations(showOtherLocations bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showOtherLocations {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_other_locations(_arg0, _arg1)
}

// SetShowRecent sets whether the @sidebar should show an item for recent files.
// The default value for this option is determined by the desktop environment,
// but this function can be used to override it on a per-application basis.
func (sidebar *PlacesSidebar) SetShowRecent(showRecent bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showRecent {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_recent(_arg0, _arg1)
}

// SetShowStarredLocation: if you enable this, you should connect to the
// PlacesSidebar::show-starred-location signal.
func (sidebar *PlacesSidebar) SetShowStarredLocation(showStarredLocation bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showStarredLocation {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_starred_location(_arg0, _arg1)
}

// SetShowTrash sets whether the @sidebar should show an item for the Trash
// location.
func (sidebar *PlacesSidebar) SetShowTrash(showTrash bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showTrash {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_trash(_arg0, _arg1)
}
