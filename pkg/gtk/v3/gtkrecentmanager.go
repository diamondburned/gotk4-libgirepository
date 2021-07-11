// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_recent_manager_error_get_type()), F: marshalRecentManagerError},
		{T: externglib.Type(C.gtk_recent_manager_get_type()), F: marshalRecentManagerer},
		{T: externglib.Type(C.gtk_recent_info_get_type()), F: marshalRecentInfo},
	})
}

// RecentManagerError: error codes for RecentManager operations
type RecentManagerError int

const (
	// NotFound: URI specified does not exists in the recently used resources
	// list.
	RecentManagerErrorNotFound RecentManagerError = iota
	// InvalidURI: URI specified is not valid.
	RecentManagerErrorInvalidURI
	// InvalidEncoding: supplied string is not UTF-8 encoded.
	RecentManagerErrorInvalidEncoding
	// NotRegistered: no application has registered the specified item.
	RecentManagerErrorNotRegistered
	// Read: failure while reading the recently used resources file.
	RecentManagerErrorRead
	// Write: failure while writing the recently used resources file.
	RecentManagerErrorWrite
	// Unknown: unspecified error.
	RecentManagerErrorUnknown
)

func marshalRecentManagerError(p uintptr) (interface{}, error) {
	return RecentManagerError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RecentManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RecentManagerOverrider interface {
	//
	Changed()
}

// RecentManagerer describes RecentManager's methods.
type RecentManagerer interface {
	// AddFull adds a new resource, pointed by @uri, into the recently used
	// resources list, using the metadata specified inside the RecentData-struct
	// passed in @recent_data.
	AddFull(uri string, recentData *RecentData) bool
	// AddItem adds a new resource, pointed by @uri, into the recently used
	// resources list.
	AddItem(uri string) bool
	// HasItem checks whether there is a recently used resource registered with
	// @uri inside the recent manager.
	HasItem(uri string) bool
	// LookupItem searches for a URI inside the recently used resources list,
	// and returns a RecentInfo-struct containing informations about the
	// resource like its MIME type, or its display name.
	LookupItem(uri string) (*RecentInfo, error)
	// MoveItem changes the location of a recently used resource from @uri to
	// @new_uri.
	MoveItem(uri string, newUri string) error
	// PurgeItems purges every item from the recently used resources list.
	PurgeItems() (int, error)
	// RemoveItem removes a resource pointed by @uri from the recently used
	// resources list handled by a recent manager.
	RemoveItem(uri string) error
}

// RecentManager provides a facility for adding, removing and looking up
// recently used files. Each recently used file is identified by its URI, and
// has meta-data associated to it, like the names and command lines of the
// applications that have registered it, the number of time each application has
// registered the same file, the mime type of the file and whether the file
// should be displayed only by the applications that have registered it.
//
// The recently used files list is per user.
//
// The RecentManager acts like a database of all the recently used files. You
// can create new RecentManager objects, but it is more efficient to use the
// default manager created by GTK+.
//
// Adding a new recently used file is as simple as:
//
//    GtkRecentManager *manager;
//    GtkRecentInfo *info;
//    GError *error = NULL;
//
//    manager = gtk_recent_manager_get_default ();
//    info = gtk_recent_manager_lookup_item (manager, file_uri, &error);
//    if (error)
//      {
//        g_warning ("Could not find the file: s", error->message);
//        g_error_free (error);
//      }
//    else
//     {
//       // Use the info object
//       gtk_recent_info_unref (info);
//     }
//
// In order to retrieve the list of recently used files, you can use
// gtk_recent_manager_get_items(), which returns a list of RecentInfo-structs.
//
// A RecentManager is the model used to populate the contents of one, or more
// RecentChooser implementations.
//
// Note that the maximum age of the recently used files list is controllable
// through the Settings:gtk-recent-files-max-age property.
//
// Recently used files are supported since GTK+ 2.10.
type RecentManager struct {
	*externglib.Object
}

var (
	_ RecentManagerer = (*RecentManager)(nil)
	_ gextras.Nativer = (*RecentManager)(nil)
)

func wrapRecentManager(obj *externglib.Object) RecentManagerer {
	return &RecentManager{
		Object: obj,
	}
}

func marshalRecentManagerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRecentManager(obj), nil
}

// NewRecentManager creates a new recent manager object. Recent manager objects
// are used to handle the list of recently used resources. A RecentManager
// object monitors the recently used resources list, and emits the “changed”
// signal each time something inside the list changes.
//
// RecentManager objects are expensive: be sure to create them only when needed.
// You should use gtk_recent_manager_get_default() instead.
func NewRecentManager() *RecentManager {
	var _cret *C.GtkRecentManager // in

	_cret = C.gtk_recent_manager_new()

	var _recentManager *RecentManager // out

	_recentManager = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RecentManager)

	return _recentManager
}

// AddFull adds a new resource, pointed by @uri, into the recently used
// resources list, using the metadata specified inside the RecentData-struct
// passed in @recent_data.
//
// The passed URI will be used to identify this resource inside the list.
//
// In order to register the new recently used resource, metadata about the
// resource must be passed as well as the URI; the metadata is stored in a
// RecentData-struct, which must contain the MIME type of the resource pointed
// by the URI; the name of the application that is registering the item, and a
// command line to be used when launching the item.
//
// Optionally, a RecentData-struct might contain a UTF-8 string to be used when
// viewing the item instead of the last component of the URI; a short
// description of the item; whether the item should be considered private - that
// is, should be displayed only by the applications that have registered it.
func (manager *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _arg2 *C.GtkRecentData    // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkRecentData)(unsafe.Pointer(recentData))

	_cret = C.gtk_recent_manager_add_full(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddItem adds a new resource, pointed by @uri, into the recently used
// resources list.
//
// This function automatically retrieves some of the needed metadata and setting
// other metadata to common default values; it then feeds the data to
// gtk_recent_manager_add_full().
//
// See gtk_recent_manager_add_full() if you want to explicitly define the
// metadata for the resource pointed by @uri.
func (manager *RecentManager) AddItem(uri string) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_add_item(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasItem checks whether there is a recently used resource registered with @uri
// inside the recent manager.
func (manager *RecentManager) HasItem(uri string) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_has_item(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupItem searches for a URI inside the recently used resources list, and
// returns a RecentInfo-struct containing informations about the resource like
// its MIME type, or its display name.
func (manager *RecentManager) LookupItem(uri string) (*RecentInfo, error) {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret *C.GtkRecentInfo    // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_lookup_item(_arg0, _arg1, &_cerr)

	var _recentInfo *RecentInfo // out
	var _goerr error            // out

	_recentInfo = (*RecentInfo)(unsafe.Pointer(_cret))
	C.gtk_recent_info_ref(_cret)
	runtime.SetFinalizer(_recentInfo, func(v *RecentInfo) {
		C.gtk_recent_info_unref((*C.GtkRecentInfo)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _recentInfo, _goerr
}

// MoveItem changes the location of a recently used resource from @uri to
// @new_uri.
//
// Please note that this function will not affect the resource pointed by the
// URIs, but only the URI used in the recently used resources list.
func (manager *RecentManager) MoveItem(uri string, newUri string) error {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(newUri))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_recent_manager_move_item(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PurgeItems purges every item from the recently used resources list.
func (manager *RecentManager) PurgeItems() (int, error) {
	var _arg0 *C.GtkRecentManager // out
	var _cret C.gint              // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_manager_purge_items(_arg0, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// RemoveItem removes a resource pointed by @uri from the recently used
// resources list handled by a recent manager.
func (manager *RecentManager) RemoveItem(uri string) error {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_manager_remove_item(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RecentData: meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
type RecentData struct {
	native C.GtkRecentData
}

// Native returns the underlying C source pointer.
func (r *RecentData) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// RecentInfo contains private data only, and should be accessed using the
// provided API.
//
// RecentInfo constains all the meta-data associated with an entry in the
// recently used files list.
type RecentInfo struct {
	native C.GtkRecentInfo
}

func marshalRecentInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*RecentInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// CreateAppInfo creates a Info for the specified RecentInfo
func (info *RecentInfo) CreateAppInfo(appName string) (*gio.AppInfo, error) {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret *C.GAppInfo      // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_create_app_info(_arg0, _arg1, &_cerr)

	var _appInfo *gio.AppInfo // out
	var _goerr error          // out

	_appInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.AppInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _appInfo, _goerr
}

// Exists checks whether the resource pointed by @info still exists. At the
// moment this check is done only on resources pointing to local files.
func (info *RecentInfo) Exists() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_exists(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Added gets the timestamp (seconds from system’s Epoch) when the resource was
// added to the recently used resources list.
func (info *RecentInfo) Added() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_added(_arg0)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// Age gets the number of days elapsed since the last update of the resource
// pointed by @info.
func (info *RecentInfo) Age() int {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_age(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ApplicationInfo gets the data regarding the application that has registered
// the resource pointed by @info.
//
// If the command line contains any escape characters defined inside the storage
// specification, they will be expanded.
func (info *RecentInfo) ApplicationInfo(appName string) (string, uint, int32, bool) {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // in
	var _arg3 C.guint          // in
	var _arg4 C.time_t         // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_get_application_info(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _appExec string // out
	var _count uint     // out
	var _time_ int32    // out
	var _ok bool        // out

	_appExec = C.GoString((*C.gchar)(_arg2))
	_count = uint(_arg3)
	_time_ = int32(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _appExec, _count, _time_, _ok
}

// Description gets the (short) description of the resource.
func (info *RecentInfo) Description() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// DisplayName gets the name of the resource. If none has been defined, the
// basename of the resource is obtained.
func (info *RecentInfo) DisplayName() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// GIcon retrieves the icon associated to the resource MIME type.
func (info *RecentInfo) GIcon() *gio.Icon {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_gicon(_arg0)

	var _icon *gio.Icon // out

	_icon = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.Icon)

	return _icon
}

// Icon retrieves the icon of size @size associated to the resource MIME type.
func (info *RecentInfo) Icon(size int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 C.gint           // out
	var _cret *C.GdkPixbuf     // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))
	_arg1 = C.gint(size)

	_cret = C.gtk_recent_info_get_icon(_arg0, _arg1)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// MIMEType gets the MIME type of the resource.
func (info *RecentInfo) MIMEType() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_mime_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// Modified gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last modified.
func (info *RecentInfo) Modified() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_modified(_arg0)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// PrivateHint gets the value of the “private” flag. Resources in the recently
// used list that have this flag set to true should only be displayed by the
// applications that have registered them.
func (info *RecentInfo) PrivateHint() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_private_hint(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShortName computes a valid UTF-8 string that can be used as the name of the
// item in a menu or list. For example, calling this function on an item that
// refers to “file:///foo/bar.txt” will yield “bar.txt”.
func (info *RecentInfo) ShortName() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_short_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// URI gets the URI of the resource.
func (info *RecentInfo) URI() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// URIDisplay gets a displayable version of the resource’s URI. If the resource
// is local, it returns a local path; if the resource is not local, it returns
// the UTF-8 encoded content of gtk_recent_info_get_uri().
func (info *RecentInfo) URIDisplay() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_uri_display(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Visited gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last visited.
func (info *RecentInfo) Visited() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_get_visited(_arg0)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// HasApplication checks whether an application registered this resource using
// @app_name.
func (info *RecentInfo) HasApplication(appName string) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_has_application(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasGroup checks whether @group_name appears inside the groups registered for
// the recently used item @info.
func (info *RecentInfo) HasGroup(groupName string) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_has_group(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLocal checks whether the resource is local or not by looking at the scheme
// of its URI.
func (info *RecentInfo) IsLocal() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_is_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LastApplication gets the name of the last application that have registered
// the recently used resource represented by @info.
func (info *RecentInfo) LastApplication() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_last_application(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Match checks whether two RecentInfo-struct point to the same resource.
func (infoA *RecentInfo) Match(infoB *RecentInfo) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(infoA))
	_arg1 = (*C.GtkRecentInfo)(unsafe.Pointer(infoB))

	_cret = C.gtk_recent_info_match(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref increases the reference count of @recent_info by one.
func (info *RecentInfo) ref() *RecentInfo {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.GtkRecentInfo // in

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	_cret = C.gtk_recent_info_ref(_arg0)

	var _recentInfo *RecentInfo // out

	_recentInfo = (*RecentInfo)(unsafe.Pointer(_cret))
	C.gtk_recent_info_ref(_cret)
	runtime.SetFinalizer(_recentInfo, func(v *RecentInfo) {
		C.gtk_recent_info_unref((*C.GtkRecentInfo)(unsafe.Pointer(v)))
	})

	return _recentInfo
}

// Unref decreases the reference count of @info by one. If the reference count
// reaches zero, @info is deallocated, and the memory freed.
func (info *RecentInfo) unref() {
	var _arg0 *C.GtkRecentInfo // out

	_arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(info))

	C.gtk_recent_info_unref(_arg0)
}
