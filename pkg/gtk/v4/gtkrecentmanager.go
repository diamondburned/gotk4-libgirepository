// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_RecentManager_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeRecentManagerError = coreglib.Type(girepository.MustFind("Gtk", "RecentManagerError").RegisteredGType())
	GTypeRecentManager      = coreglib.Type(girepository.MustFind("Gtk", "RecentManager").RegisteredGType())
	GTypeRecentInfo         = coreglib.Type(girepository.MustFind("Gtk", "RecentInfo").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentManagerError, F: marshalRecentManagerError},
		coreglib.TypeMarshaler{T: GTypeRecentManager, F: marshalRecentManager},
		coreglib.TypeMarshaler{T: GTypeRecentInfo, F: marshalRecentInfo},
	})
}

// RecentManagerError: error codes for RecentManager operations.
type RecentManagerError C.gint

const (
	// RecentManagerErrorNotFound: URI specified does not exists in the recently
	// used resources list.
	RecentManagerErrorNotFound RecentManagerError = iota
	// RecentManagerErrorInvalidURI: URI specified is not valid.
	RecentManagerErrorInvalidURI
	// RecentManagerErrorInvalidEncoding: supplied string is not UTF-8 encoded.
	RecentManagerErrorInvalidEncoding
	// RecentManagerErrorNotRegistered: no application has registered the
	// specified item.
	RecentManagerErrorNotRegistered
	// RecentManagerErrorRead: failure while reading the recently used resources
	// file.
	RecentManagerErrorRead
	// RecentManagerErrorWrite: failure while writing the recently used
	// resources file.
	RecentManagerErrorWrite
	// RecentManagerErrorUnknown: unspecified error.
	RecentManagerErrorUnknown
)

func marshalRecentManagerError(p uintptr) (interface{}, error) {
	return RecentManagerError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentManagerError.
func (r RecentManagerError) String() string {
	switch r {
	case RecentManagerErrorNotFound:
		return "NotFound"
	case RecentManagerErrorInvalidURI:
		return "InvalidURI"
	case RecentManagerErrorInvalidEncoding:
		return "InvalidEncoding"
	case RecentManagerErrorNotRegistered:
		return "NotRegistered"
	case RecentManagerErrorRead:
		return "Read"
	case RecentManagerErrorWrite:
		return "Write"
	case RecentManagerErrorUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("RecentManagerError(%d)", r)
	}
}

// RecentManagerOverrides contains methods that are overridable.
type RecentManagerOverrides struct {
}

func defaultRecentManagerOverrides(v *RecentManager) RecentManagerOverrides {
	return RecentManagerOverrides{}
}

// RecentManager: GtkRecentManager manages and looks up recently used files.
//
// Each recently used file is identified by its URI, and has meta-data
// associated to it, like the names and command lines of the applications that
// have registered it, the number of time each application has registered the
// same file, the mime type of the file and whether the file should be displayed
// only by the applications that have registered it.
//
// The recently used files list is per user.
//
// GtkRecentManager acts like a database of all the recently used files. You can
// create new GtkRecentManager objects, but it is more efficient to use the
// default manager created by GTK.
//
// Adding a new recently used file is as simple as:
//
//    GtkRecentManager *manager;
//
//    manager = gtk_recent_manager_get_default ();
//    gtk_recent_manager_add_item (manager, file_uri);
//
//
// The GtkRecentManager will try to gather all the needed information from the
// file itself through GIO.
//
// Looking up the meta-data associated with a recently used file given its URI
// requires calling gtk.RecentManager.LookupItem():
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
//
// In order to retrieve the list of recently used files, you can use
// gtk.RecentManager.GetItems(), which returns a list of gtk.RecentInfo.
//
// Note that the maximum age of the recently used files list is controllable
// through the gtk.Settings:gtk-recent-files-max-age property.
type RecentManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*RecentManager)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RecentManager, *RecentManagerClass, RecentManagerOverrides](
		GTypeRecentManager,
		initRecentManagerClass,
		wrapRecentManager,
		defaultRecentManagerOverrides,
	)
}

func initRecentManagerClass(gclass unsafe.Pointer, overrides RecentManagerOverrides, classInitFunc func(*RecentManagerClass)) {
	if classInitFunc != nil {
		class := (*RecentManagerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRecentManager(obj *coreglib.Object) *RecentManager {
	return &RecentManager{
		Object: obj,
	}
}

func marshalRecentManager(p uintptr) (interface{}, error) {
	return wrapRecentManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when the current recently used resources manager
// changes its contents.
//
// This can happen either by calling gtk.RecentManager.AddItem() or by another
// application.
func (v *RecentManager) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk4_RecentManager_ConnectChanged), f)
}

// RecentData: meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
//
// An instance of this type is always passed by reference.
type RecentData struct {
	*recentData
}

// recentData is the struct that's finalized.
type recentData struct {
	native unsafe.Pointer
}

var GIRInfoRecentData = girepository.MustFind("Gtk", "RecentData")

// DisplayName: UTF-8 encoded string, containing the name of the recently used
// resource to be displayed, or NULL;.
func (r *RecentData) DisplayName() string {
	offset := GIRInfoRecentData.StructFieldOffset("display_name")
	valptr := (*string)(unsafe.Add(r.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// Description: UTF-8 encoded string, containing a short description of the
// resource, or NULL;.
func (r *RecentData) Description() string {
	offset := GIRInfoRecentData.StructFieldOffset("description")
	valptr := (*string)(unsafe.Add(r.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// MIMEType: MIME type of the resource;.
func (r *RecentData) MIMEType() string {
	offset := GIRInfoRecentData.StructFieldOffset("mime_type")
	valptr := (*string)(unsafe.Add(r.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// AppName: name of the application that is registering this recently used
// resource;.
func (r *RecentData) AppName() string {
	offset := GIRInfoRecentData.StructFieldOffset("app_name")
	valptr := (*string)(unsafe.Add(r.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// AppExec: command line used to launch this resource; may contain the “\f” and
// “\u” escape characters which will be expanded to the resource file path and
// URI respectively when the command line is retrieved;.
func (r *RecentData) AppExec() string {
	offset := GIRInfoRecentData.StructFieldOffset("app_exec")
	valptr := (*string)(unsafe.Add(r.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// Groups: vector of strings containing groups names;.
func (r *RecentData) Groups() []string {
	offset := GIRInfoRecentData.StructFieldOffset("groups")
	valptr := (*[]string)(unsafe.Add(r.native, offset))
	var _v []string // out
	{
		var i int
		var z *C.char
		for p := *valptr; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*valptr, i)
		_v = make([]string, i)
		for i := range src {
			_v[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}
	return _v
}

// IsPrivate: whether this resource should be displayed only by the applications
// that have registered it or not.
func (r *RecentData) IsPrivate() bool {
	offset := GIRInfoRecentData.StructFieldOffset("is_private")
	valptr := (*bool)(unsafe.Add(r.native, offset))
	var _v bool // out
	if *valptr != 0 {
		_v = true
	}
	return _v
}

// IsPrivate: whether this resource should be displayed only by the applications
// that have registered it or not.
func (r *RecentData) SetIsPrivate(isPrivate bool) {
	offset := GIRInfoRecentData.StructFieldOffset("is_private")
	valptr := (*C.gboolean)(unsafe.Add(r.native, offset))
	if isPrivate {
		*valptr = C.TRUE
	}
}

// RecentInfo: GtkRecentInfo contains the metadata associated with an item in
// the recently used files list.
//
// An instance of this type is always passed by reference.
type RecentInfo struct {
	*recentInfo
}

// recentInfo is the struct that's finalized.
type recentInfo struct {
	native unsafe.Pointer
}

var GIRInfoRecentInfo = girepository.MustFind("Gtk", "RecentInfo")

func marshalRecentInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &RecentInfo{&recentInfo{(unsafe.Pointer)(b)}}, nil
}

// RecentManagerClass contains only private data.
//
// An instance of this type is always passed by reference.
type RecentManagerClass struct {
	*recentManagerClass
}

// recentManagerClass is the struct that's finalized.
type recentManagerClass struct {
	native unsafe.Pointer
}

var GIRInfoRecentManagerClass = girepository.MustFind("Gtk", "RecentManagerClass")
