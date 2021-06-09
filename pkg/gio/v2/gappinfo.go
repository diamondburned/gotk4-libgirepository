// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_app_info_get_type()), F: marshalAppInfo},
		{T: externglib.Type(C.g_app_launch_context_get_type()), F: marshalAppLaunchContext},
	})
}

// AppInfoOverrider contains methods that are overridable. This
// interface is a subset of the interface AppInfo.
type AppInfoOverrider interface {
	// AddSupportsType adds a content type to the application information to
	// indicate the application is capable of opening files with the given
	// content type.
	AddSupportsType(contentType string) error
	// CanDelete obtains the information whether the Info can be deleted. See
	// g_app_info_delete().
	CanDelete() bool
	// CanRemoveSupportsType checks if a supported content type can be removed
	// from an application.
	CanRemoveSupportsType() bool
	// DoDelete tries to delete a Info.
	//
	// On some platforms, there may be a difference between user-defined Infos
	// which can be deleted, and system-wide ones which cannot. See
	// g_app_info_can_delete().
	DoDelete() bool
	// Dup creates a duplicate of a Info.
	Dup() AppInfo
	// Equal checks if two Infos are equal.
	//
	// Note that the check *may not* compare each individual field, and only
	// does an identity check. In case detecting changes in the contents is
	// needed, program code must additionally compare relevant fields.
	Equal(appinfo2 AppInfo) bool
	// Commandline gets the commandline with which the application will be
	// started.
	Commandline() *string
	// Description gets a human-readable description of an installed
	// application.
	Description() string
	// DisplayName gets the display name of the application. The display name is
	// often more descriptive to the user than the name itself.
	DisplayName() string
	// Executable gets the executable's name for the installed application.
	Executable() *string
	// Icon gets the icon for the application.
	Icon() Icon
	// ID gets the ID of an application. An id is a string that identifies the
	// application. The exact format of the id is platform dependent. For
	// instance, on Unix this is the desktop file id from the xdg menu
	// specification.
	//
	// Note that the returned ID may be nil, depending on how the @appinfo has
	// been constructed.
	ID() string
	// Name gets the installed name of the application.
	Name() string
	// SupportedTypes retrieves the list of content types that @app_info claims
	// to support. If this information is not provided by the environment, this
	// function will return nil. This function does not take in consideration
	// associations added with g_app_info_add_supports_type(), but only those
	// exported directly by the application.
	SupportedTypes() []string
	// Launch launches the application. Passes @files to the launched
	// application as arguments, using the optional @context to get information
	// about the details of the launcher (like what screen it is on). On error,
	// @error will be set accordingly.
	//
	// To launch the application without arguments pass a nil @files list.
	//
	// Note that even if the launch is successful the application launched can
	// fail to start if it runs into problems during startup. There is no way to
	// detect this.
	//
	// Some URIs can be changed when passed through a GFile (for instance
	// unsupported URIs with strange formats like mailto:), so if you have a
	// textual URI you want to pass in as argument, consider using
	// g_app_info_launch_uris() instead.
	//
	// The launched application inherits the environment of the launching
	// process, but it can be modified with g_app_launch_context_setenv() and
	// g_app_launch_context_unsetenv().
	//
	// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE` environment
	// variable with the path of the launched desktop file and
	// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
	// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`, should
	// it be inherited by further processes. The `DISPLAY` and
	// `DESKTOP_STARTUP_ID` environment variables are also set, based on
	// information provided in @context.
	Launch(files *glib.List, context AppLaunchContext) error
	// LaunchUris launches the application. This passes the @uris to the
	// launched application as arguments, using the optional @context to get
	// information about the details of the launcher (like what screen it is
	// on). On error, @error will be set accordingly.
	//
	// To launch the application without arguments pass a nil @uris list.
	//
	// Note that even if the launch is successful the application launched can
	// fail to start if it runs into problems during startup. There is no way to
	// detect this.
	LaunchUris(uris *glib.List, context AppLaunchContext) error
	// LaunchUrisAsync: async version of g_app_info_launch_uris().
	//
	// The @callback is invoked immediately after the application launch, but it
	// waits for activation in case of D-Bus–activated applications and also
	// provides extended error information for sandboxed applications, see notes
	// for g_app_info_launch_default_for_uri_async().
	LaunchUrisAsync()
	// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
	LaunchUrisFinish(result AsyncResult) error
	// RemoveSupportsType removes a supported type from an application, if
	// possible.
	RemoveSupportsType(contentType string) error
	// SetAsDefaultForExtension sets the application as the default handler for
	// the given file extension.
	SetAsDefaultForExtension(extension *string) error
	// SetAsDefaultForType sets the application as the default handler for a
	// given type.
	SetAsDefaultForType(contentType string) error
	// SetAsLastUsedForType sets the application as the last used application
	// for a given type. This will make the application appear as first in the
	// list returned by g_app_info_get_recommended_for_type(), regardless of the
	// default application for that content type.
	SetAsLastUsedForType(contentType string) error
	// ShouldShow checks if the application info should be shown in menus that
	// list available applications.
	ShouldShow() bool
	// SupportsFiles checks if the application accepts files as arguments.
	SupportsFiles() bool
	// SupportsUris checks if the application supports reading files and
	// directories from URIs.
	SupportsUris() bool
}

// AppInfo: Info and LaunchContext are used for describing and launching
// applications installed on the system.
//
// As of GLib 2.20, URIs will always be converted to POSIX paths (using
// g_file_get_path()) when using g_app_info_launch() even if the application
// requested an URI and not a POSIX path. For example for a desktop-file based
// application with Exec key `totem U` and a single URI, `sftp://foo/file.avi`,
// then `/home/user/.gvfs/sftp on foo/file.avi` will be passed. This will only
// work if a set of suitable GIO extensions (such as gvfs 2.26 compiled with
// FUSE support), is available and operational; if this is not the case, the URI
// will be passed unmodified to the application. Some URIs, such as `mailto:`,
// of course cannot be mapped to a POSIX path (in gvfs there's no FUSE mount for
// it); such URIs will be passed unmodified to the application.
//
// Specifically for gvfs 2.26 and later, the POSIX URI will be mapped back to
// the GIO URI in the #GFile constructors (since gvfs implements the #GVfs
// extension point). As such, if the application needs to examine the URI, it
// needs to use g_file_get_uri() or similar on #GFile. In other words, an
// application cannot assume that the URI passed to e.g.
// g_file_new_for_commandline_arg() is equal to the result of g_file_get_uri().
// The following snippet illustrates this:
//
//    GFile *f;
//    char *uri;
//
//    file = g_file_new_for_commandline_arg (uri_from_commandline);
//
//    uri = g_file_get_uri (file);
//    strcmp (uri, uri_from_commandline) == 0;
//    g_free (uri);
//
//    if (g_file_has_uri_scheme (file, "cdda"))
//      {
//        // do something special with uri
//      }
//    g_object_unref (file);
//
// This code will work when both `cdda://sr0/Track 1.wav` and
// `/home/user/.gvfs/cdda on sr0/Track 1.wav` is passed to the application. It
// should be noted that it's generally not safe for applications to rely on the
// format of a particular URIs. Different launcher applications (e.g. file
// managers) may have different ideas of what a given URI means.
type AppInfo interface {
	gextras.Objector
	AppInfoOverrider

	// Delete tries to delete a Info.
	//
	// On some platforms, there may be a difference between user-defined Infos
	// which can be deleted, and system-wide ones which cannot. See
	// g_app_info_can_delete().
	Delete() bool
}

// appInfo implements the AppInfo interface.
type appInfo struct {
	gextras.Objector
}

var _ AppInfo = (*appInfo)(nil)

// WrapAppInfo wraps a GObject to a type that implements interface
// AppInfo. It is primarily used internally.
func WrapAppInfo(obj *externglib.Object) AppInfo {
	return AppInfo{
		Objector: obj,
	}
}

func marshalAppInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppInfo(obj), nil
}

// AddSupportsType adds a content type to the application information to
// indicate the application is capable of opening files with the given
// content type.
func (a appInfo) AddSupportsType(contentType string) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError

	C.g_app_info_add_supports_type(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// CanDelete obtains the information whether the Info can be deleted. See
// g_app_info_delete().
func (a appInfo) CanDelete() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_can_delete(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// CanRemoveSupportsType checks if a supported content type can be removed
// from an application.
func (a appInfo) CanRemoveSupportsType() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_can_remove_supports_type(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Delete tries to delete a Info.
//
// On some platforms, there may be a difference between user-defined Infos
// which can be deleted, and system-wide ones which cannot. See
// g_app_info_can_delete().
func (a appInfo) Delete() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_delete(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Dup creates a duplicate of a Info.
func (a appInfo) Dup() AppInfo {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.GAppInfo

	cret = C.g_app_info_dup(_arg0)

	var _appInfo AppInfo

	_appInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(AppInfo)

	return _appInfo
}

// Equal checks if two Infos are equal.
//
// Note that the check *may not* compare each individual field, and only
// does an identity check. In case detecting changes in the contents is
// needed, program code must additionally compare relevant fields.
func (a appInfo) Equal(appinfo2 AppInfo) bool {
	var _arg0 *C.GAppInfo
	var _arg1 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GAppInfo)(unsafe.Pointer(appinfo2.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Commandline gets the commandline with which the application will be
// started.
func (a appInfo) Commandline() *string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_commandline(_arg0)

	var _filename *string

	_filename = C.GoString(_cret)

	return _filename
}

// Description gets a human-readable description of an installed
// application.
func (a appInfo) Description() string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_description(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// DisplayName gets the display name of the application. The display name is
// often more descriptive to the user than the name itself.
func (a appInfo) DisplayName() string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_display_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Executable gets the executable's name for the installed application.
func (a appInfo) Executable() *string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_executable(_arg0)

	var _filename *string

	_filename = C.GoString(_cret)

	return _filename
}

// Icon gets the icon for the application.
func (a appInfo) Icon() Icon {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.GIcon

	cret = C.g_app_info_get_icon(_arg0)

	var _icon Icon

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Icon)

	return _icon
}

// ID gets the ID of an application. An id is a string that identifies the
// application. The exact format of the id is platform dependent. For
// instance, on Unix this is the desktop file id from the xdg menu
// specification.
//
// Note that the returned ID may be nil, depending on how the @appinfo has
// been constructed.
func (a appInfo) ID() string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_id(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Name gets the installed name of the application.
func (a appInfo) Name() string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret *C.char

	cret = C.g_app_info_get_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SupportedTypes retrieves the list of content types that @app_info claims
// to support. If this information is not provided by the environment, this
// function will return nil. This function does not take in consideration
// associations added with g_app_info_add_supports_type(), but only those
// exported directly by the application.
func (a appInfo) SupportedTypes() []string {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret **C.char

	cret = C.g_app_info_get_supported_types(_arg0)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_utf8s = C.GoString(_cret)
		}
	}

	return _utf8s
}

// Launch launches the application. Passes @files to the launched
// application as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on). On error,
// @error will be set accordingly.
//
// To launch the application without arguments pass a nil @files list.
//
// Note that even if the launch is successful the application launched can
// fail to start if it runs into problems during startup. There is no way to
// detect this.
//
// Some URIs can be changed when passed through a GFile (for instance
// unsupported URIs with strange formats like mailto:), so if you have a
// textual URI you want to pass in as argument, consider using
// g_app_info_launch_uris() instead.
//
// The launched application inherits the environment of the launching
// process, but it can be modified with g_app_launch_context_setenv() and
// g_app_launch_context_unsetenv().
//
// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE` environment
// variable with the path of the launched desktop file and
// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`, should
// it be inherited by further processes. The `DISPLAY` and
// `DESKTOP_STARTUP_ID` environment variables are also set, based on
// information provided in @context.
func (a appInfo) Launch(files *glib.List, context AppLaunchContext) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.GList
	var _arg2 *C.GAppLaunchContext

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GList)(unsafe.Pointer(files.Native()))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	var _cerr *C.GError

	C.g_app_info_launch(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LaunchUris launches the application. This passes the @uris to the
// launched application as arguments, using the optional @context to get
// information about the details of the launcher (like what screen it is
// on). On error, @error will be set accordingly.
//
// To launch the application without arguments pass a nil @uris list.
//
// Note that even if the launch is successful the application launched can
// fail to start if it runs into problems during startup. There is no way to
// detect this.
func (a appInfo) LaunchUris(uris *glib.List, context AppLaunchContext) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.GList
	var _arg2 *C.GAppLaunchContext

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GList)(unsafe.Pointer(uris.Native()))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	var _cerr *C.GError

	C.g_app_info_launch_uris(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LaunchUrisAsync: async version of g_app_info_launch_uris().
//
// The @callback is invoked immediately after the application launch, but it
// waits for activation in case of D-Bus–activated applications and also
// provides extended error information for sandboxed applications, see notes
// for g_app_info_launch_default_for_uri_async().
func (a appInfo) LaunchUrisAsync() {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_launch_uris_async(_arg0)
}

// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
func (a appInfo) LaunchUrisFinish(result AsyncResult) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cerr *C.GError

	C.g_app_info_launch_uris_finish(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RemoveSupportsType removes a supported type from an application, if
// possible.
func (a appInfo) RemoveSupportsType(contentType string) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError

	C.g_app_info_remove_supports_type(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsDefaultForExtension sets the application as the default handler for
// the given file extension.
func (a appInfo) SetAsDefaultForExtension(extension *string) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(extension))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError

	C.g_app_info_set_as_default_for_extension(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsDefaultForType sets the application as the default handler for a
// given type.
func (a appInfo) SetAsDefaultForType(contentType string) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError

	C.g_app_info_set_as_default_for_type(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsLastUsedForType sets the application as the last used application
// for a given type. This will make the application appear as first in the
// list returned by g_app_info_get_recommended_for_type(), regardless of the
// default application for that content type.
func (a appInfo) SetAsLastUsedForType(contentType string) error {
	var _arg0 *C.GAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError

	C.g_app_info_set_as_last_used_for_type(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ShouldShow checks if the application info should be shown in menus that
// list available applications.
func (a appInfo) ShouldShow() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_should_show(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SupportsFiles checks if the application accepts files as arguments.
func (a appInfo) SupportsFiles() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_supports_files(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SupportsUris checks if the application supports reading files and
// directories from URIs.
func (a appInfo) SupportsUris() bool {
	var _arg0 *C.GAppInfo

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_app_info_supports_uris(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// AppLaunchContext: integrating the launch with the launching application. This
// is used to handle for instance startup notification and launching the new
// application on the same screen as the launching window.
type AppLaunchContext interface {
	gextras.Objector

	// Display gets the display string for the @context. This is used to ensure
	// new applications are started on the same display as the launching
	// application, by setting the `DISPLAY` environment variable.
	Display(info AppInfo, files *glib.List) string
	// Environment gets the complete environment variable list to be passed to
	// the child process when @context is used to launch an application. This is
	// a nil-terminated array of strings, where each string has the form
	// `KEY=VALUE`.
	Environment() []*string
	// StartupNotifyID initiates startup notification for the application and
	// returns the `DESKTOP_STARTUP_ID` for the launched operation, if
	// supported.
	//
	// Startup notification IDs are defined in the FreeDesktop.Org Startup
	// Notifications standard
	// (http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
	StartupNotifyID(info AppInfo, files *glib.List) string
	// LaunchFailed: called when an application has failed to launch, so that it
	// can cancel the application startup notification started in
	// g_app_launch_context_get_startup_notify_id().
	LaunchFailed(startupNotifyId string)
	// Setenv arranges for @variable to be set to @value in the child's
	// environment when @context is used to launch an application.
	Setenv(variable *string, value *string)
	// Unsetenv arranges for @variable to be unset in the child's environment
	// when @context is used to launch an application.
	Unsetenv(variable *string)
}

// appLaunchContext implements the AppLaunchContext interface.
type appLaunchContext struct {
	gextras.Objector
}

var _ AppLaunchContext = (*appLaunchContext)(nil)

// WrapAppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppLaunchContext(obj *externglib.Object) AppLaunchContext {
	return AppLaunchContext{
		Objector: obj,
	}
}

func marshalAppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppLaunchContext(obj), nil
}

// NewAppLaunchContext constructs a class AppLaunchContext.
func NewAppLaunchContext() AppLaunchContext {
	var _cret C.GAppLaunchContext

	cret = C.g_app_launch_context_new()

	var _appLaunchContext AppLaunchContext

	_appLaunchContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(AppLaunchContext)

	return _appLaunchContext
}

// Display gets the display string for the @context. This is used to ensure
// new applications are started on the same display as the launching
// application, by setting the `DISPLAY` environment variable.
func (c appLaunchContext) Display(info AppInfo, files *glib.List) string {
	var _arg0 *C.GAppLaunchContext
	var _arg1 *C.GAppInfo
	var _arg2 *C.GList

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAppInfo)(unsafe.Pointer(info.Native()))
	_arg2 = (*C.GList)(unsafe.Pointer(files.Native()))

	var _cret *C.char

	cret = C.g_app_launch_context_get_display(_arg0, _arg1, _arg2)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Environment gets the complete environment variable list to be passed to
// the child process when @context is used to launch an application. This is
// a nil-terminated array of strings, where each string has the form
// `KEY=VALUE`.
func (c appLaunchContext) Environment() []*string {
	var _arg0 *C.GAppLaunchContext

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))

	var _cret **C.char

	cret = C.g_app_launch_context_get_environment(_arg0)

	var _filenames []*string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_filenames = make([]*string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			_filenames = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _filenames
}

// StartupNotifyID initiates startup notification for the application and
// returns the `DESKTOP_STARTUP_ID` for the launched operation, if
// supported.
//
// Startup notification IDs are defined in the FreeDesktop.Org Startup
// Notifications standard
// (http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
func (c appLaunchContext) StartupNotifyID(info AppInfo, files *glib.List) string {
	var _arg0 *C.GAppLaunchContext
	var _arg1 *C.GAppInfo
	var _arg2 *C.GList

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAppInfo)(unsafe.Pointer(info.Native()))
	_arg2 = (*C.GList)(unsafe.Pointer(files.Native()))

	var _cret *C.char

	cret = C.g_app_launch_context_get_startup_notify_id(_arg0, _arg1, _arg2)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LaunchFailed: called when an application has failed to launch, so that it
// can cancel the application startup notification started in
// g_app_launch_context_get_startup_notify_id().
func (c appLaunchContext) LaunchFailed(startupNotifyId string) {
	var _arg0 *C.GAppLaunchContext
	var _arg1 *C.char

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(startupNotifyId))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_app_launch_context_launch_failed(_arg0, _arg1)
}

// Setenv arranges for @variable to be set to @value in the child's
// environment when @context is used to launch an application.
func (c appLaunchContext) Setenv(variable *string, value *string) {
	var _arg0 *C.GAppLaunchContext
	var _arg1 *C.char
	var _arg2 *C.char

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(variable))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_app_launch_context_setenv(_arg0, _arg1, _arg2)
}

// Unsetenv arranges for @variable to be unset in the child's environment
// when @context is used to launch an application.
func (c appLaunchContext) Unsetenv(variable *string) {
	var _arg0 *C.GAppLaunchContext
	var _arg1 *C.char

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(variable))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_app_launch_context_unsetenv(_arg0, _arg1)
}