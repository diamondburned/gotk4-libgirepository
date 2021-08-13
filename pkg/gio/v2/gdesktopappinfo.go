// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
// void _gotk4_gio2_DesktopAppLaunchCallback(GDesktopAppInfo*, GPid, gpointer);
// void _gotk4_glib2_SpawnChildSetupFunc(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_desktop_app_info_lookup_get_type()), F: marshalDesktopAppInfoLookupper},
		{T: externglib.Type(C.g_desktop_app_info_get_type()), F: marshalDesktopAppInfor},
	})
}

// DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME: extension point for default
// handler to URI association. See [Extending GIO][extending-gio].
//
// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME = "gio-desktop-app-info-lookup"

// DesktopAppLaunchCallback: during invocation,
// g_desktop_app_info_launch_uris_as_manager() may create one or more child
// processes. This callback is invoked once for each, providing the process ID.
type DesktopAppLaunchCallback func(appinfo *DesktopAppInfo, pid glib.Pid)

//export _gotk4_gio2_DesktopAppLaunchCallback
func _gotk4_gio2_DesktopAppLaunchCallback(arg0 *C.GDesktopAppInfo, arg1 C.GPid, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var appinfo *DesktopAppInfo // out
	var pid glib.Pid            // out

	appinfo = wrapDesktopAppInfo(externglib.Take(unsafe.Pointer(arg0)))
	pid = int(arg1)

	fn := v.(DesktopAppLaunchCallback)
	fn(appinfo, pid)
}

// DesktopAppInfoLookupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DesktopAppInfoLookupOverrider interface {
	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	//
	// The AppInfoLookup interface and this function is used to implement
	// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
	// is no reason for applications to use it directly. Applications should use
	// g_app_info_get_default_for_uri_scheme().
	//
	// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
	DefaultForURIScheme(uriScheme string) AppInfor
}

// DesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
//
// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
type DesktopAppInfoLookup struct {
	*externglib.Object
}

// DesktopAppInfoLookupper describes DesktopAppInfoLookup's abstract methods.
type DesktopAppInfoLookupper interface {
	externglib.Objector

	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	DefaultForURIScheme(uriScheme string) AppInfor
}

var _ DesktopAppInfoLookupper = (*DesktopAppInfoLookup)(nil)

func wrapDesktopAppInfoLookup(obj *externglib.Object) *DesktopAppInfoLookup {
	return &DesktopAppInfoLookup{
		Object: obj,
	}
}

func marshalDesktopAppInfoLookupper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDesktopAppInfoLookup(obj), nil
}

// DefaultForURIScheme gets the default application for launching applications
// using this URI scheme for a particular AppInfoLookup implementation.
//
// The AppInfoLookup interface and this function is used to implement
// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There is no
// reason for applications to use it directly. Applications should use
// g_app_info_get_default_for_uri_scheme().
//
// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
func (lookup *DesktopAppInfoLookup) DefaultForURIScheme(uriScheme string) AppInfor {
	var _arg0 *C.GDesktopAppInfoLookup // out
	var _arg1 *C.char                  // out
	var _cret *C.GAppInfo              // in

	_arg0 = (*C.GDesktopAppInfoLookup)(unsafe.Pointer(lookup.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uriScheme)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_lookup_get_default_for_uri_scheme(_arg0, _arg1)

	var _appInfo AppInfor // out

	if _cret != nil {
		_appInfo = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(AppInfor)
	}

	return _appInfo
}

// DesktopAppInfo is an implementation of Info based on desktop files.
//
// Note that <gio/gdesktopappinfo.h> belongs to the UNIX-specific GIO
// interfaces, thus you have to use the gio-unix-2.0.pc pkg-config file when
// using it.
type DesktopAppInfo struct {
	*externglib.Object

	AppInfo
}

func wrapDesktopAppInfo(obj *externglib.Object) *DesktopAppInfo {
	return &DesktopAppInfo{
		Object: obj,
		AppInfo: AppInfo{
			Object: obj,
		},
	}
}

func marshalDesktopAppInfor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDesktopAppInfo(obj), nil
}

// NewDesktopAppInfo creates a new AppInfo based on a desktop file id.
//
// A desktop file id is the basename of the desktop file, including the .desktop
// extension. GIO is looking for a desktop file with this name in the
// applications subdirectories of the XDG data directories (i.e. the directories
// specified in the XDG_DATA_HOME and XDG_DATA_DIRS environment variables). GIO
// also supports the prefix-to-subdirectory mapping that is described in the
// Menu Spec (http://standards.freedesktop.org/menu-spec/latest/) (i.e. a
// desktop id of kde-foo.desktop will match
// /usr/share/applications/kde/foo.desktop).
func NewDesktopAppInfo(desktopId string) *DesktopAppInfo {
	var _arg1 *C.char            // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(desktopId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_new(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	if _cret != nil {
		_desktopAppInfo = wrapDesktopAppInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _desktopAppInfo
}

// NewDesktopAppInfoFromFilename creates a new AppInfo.
func NewDesktopAppInfoFromFilename(filename string) *DesktopAppInfo {
	var _arg1 *C.char            // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_new_from_filename(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	if _cret != nil {
		_desktopAppInfo = wrapDesktopAppInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _desktopAppInfo
}

// NewDesktopAppInfoFromKeyfile creates a new AppInfo.
func NewDesktopAppInfoFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	var _arg1 *C.GKeyFile        // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))

	_cret = C.g_desktop_app_info_new_from_keyfile(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	if _cret != nil {
		_desktopAppInfo = wrapDesktopAppInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _desktopAppInfo
}

// ActionName gets the user-visible display name of the "additional application
// action" specified by action_name.
//
// This corresponds to the "Name" key within the keyfile group for the action.
func (info *DesktopAppInfo) ActionName(actionName string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_action_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Boolean looks up a boolean value in the keyfile backing info.
//
// The key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) Boolean(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Categories gets the categories from the desktop file.
func (info *DesktopAppInfo) Categories() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_categories(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Filename: when info was created from a known filename, return it. In some
// situations such as the AppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return NULL.
func (info *DesktopAppInfo) Filename() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_filename(_arg0)

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}

// GenericName gets the generic name from the desktop file.
func (info *DesktopAppInfo) GenericName() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_generic_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IsHidden: desktop file is hidden if the Hidden key in it is set to True.
func (info *DesktopAppInfo) IsHidden() bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_is_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Keywords gets the keywords from the desktop file.
func (info *DesktopAppInfo) Keywords() []string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret **C.char           // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_keywords(_arg0)

	var _utf8s []string // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// LocaleString looks up a localized string value in the keyfile backing info
// translated to the current locale.
//
// The key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) LocaleString(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_locale_string(_arg0, _arg1)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Nodisplay gets the value of the NoDisplay key, which helps determine if the
// application info should be shown in menus. See
// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
func (info *DesktopAppInfo) Nodisplay() bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_nodisplay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowIn checks if the application info should be shown in menus that list
// available applications for a specific name of the desktop, based on the
// OnlyShowIn and NotShowIn keys.
//
// desktop_env should typically be given as NULL, in which case the
// XDG_CURRENT_DESKTOP environment variable is consulted. If you want to
// override the default mechanism then you may specify desktop_env, but this is
// not recommended.
//
// Note that g_app_info_should_show() for info will include this check (with
// NULL for desktop_env) as well as additional checks.
func (info *DesktopAppInfo) ShowIn(desktopEnv string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	if desktopEnv != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(desktopEnv)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_desktop_app_info_get_show_in(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartupWmClass retrieves the StartupWMClass field from info. This represents
// the WM_CLASS property of the main window of the application, if launched
// through info.
func (info *DesktopAppInfo) StartupWmClass() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_startup_wm_class(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// String looks up a string value in the keyfile backing info.
//
// The key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) String(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_string(_arg0, _arg1)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// HasKey returns whether key exists in the "Desktop Entry" group of the keyfile
// backing info.
func (info *DesktopAppInfo) HasKey(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LaunchAction activates the named application action.
//
// You may only call this function on action names that were returned from
// g_desktop_app_info_list_actions().
//
// Note that if the main entry of the desktop file indicates that the
// application supports startup notification, and launch_context is non-NULL,
// then startup notification will be used when activating the action (and as
// such, invocation of the action on the receiving side must signal the end of
// startup notification when it is completed). This is the expected behaviour of
// applications declaring additional actions, as per the desktop file
// specification.
//
// As with g_app_info_launch() there is no way to detect failures that occur
// while using this function.
func (info *DesktopAppInfo) LaunchAction(actionName string, launchContext *AppLaunchContext) {
	var _arg0 *C.GDesktopAppInfo   // out
	var _arg1 *C.gchar             // out
	var _arg2 *C.GAppLaunchContext // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if launchContext != nil {
		_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(launchContext.Native()))
	}

	C.g_desktop_app_info_launch_action(_arg0, _arg1, _arg2)
}

// LaunchURIsAsManager: this function performs the equivalent of
// g_app_info_launch_uris(), but is intended primarily for operating system
// components that launch applications. Ordinary applications should use
// g_app_info_launch_uris().
//
// If the application is launched via GSpawn, then spawn_flags, user_setup and
// user_setup_data are used for the call to g_spawn_async(). Additionally,
// pid_callback (with pid_callback_data) will be called to inform about the PID
// of the created process. See g_spawn_async_with_pipes() for information on
// certain parameter conditions that can enable an optimized posix_spawn()
// codepath to be used.
//
// If application launching occurs via some other mechanism (eg: D-Bus
// activation) then spawn_flags, user_setup, user_setup_data, pid_callback and
// pid_callback_data are ignored.
func (appinfo *DesktopAppInfo) LaunchURIsAsManager(uris []string, launchContext *AppLaunchContext, spawnFlags glib.SpawnFlags, userSetup glib.SpawnChildSetupFunc, pidCallback DesktopAppLaunchCallback) error {
	var _arg0 *C.GDesktopAppInfo     // out
	var _arg1 *C.GList               // out
	var _arg2 *C.GAppLaunchContext   // out
	var _arg3 C.GSpawnFlags          // out
	var _arg4 C.GSpawnChildSetupFunc // out
	var _arg5 C.gpointer
	var _arg6 C.GDesktopAppLaunchCallback // out
	var _arg7 C.gpointer
	var _cerr *C.GError // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(appinfo.Native()))
	for i := len(uris) - 1; i >= 0; i-- {
		src := uris[i]
		var dst *C.gchar // out
		dst = (*C.gchar)(unsafe.Pointer(C.CString(src)))
		defer C.free(unsafe.Pointer(dst))
		_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg1)
	if launchContext != nil {
		_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(launchContext.Native()))
	}
	_arg3 = C.GSpawnFlags(spawnFlags)
	if userSetup != nil {
		_arg4 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg5 = C.gpointer(gbox.AssignOnce(userSetup))
	}
	if pidCallback != nil {
		_arg6 = (*[0]byte)(C._gotk4_gio2_DesktopAppLaunchCallback)
		_arg7 = C.gpointer(gbox.Assign(pidCallback))
		defer gbox.Delete(uintptr(_arg7))
	}

	C.g_desktop_app_info_launch_uris_as_manager(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LaunchURIsAsManagerWithFds: equivalent to
// g_desktop_app_info_launch_uris_as_manager() but allows you to pass in file
// descriptors for the stdin, stdout and stderr streams of the launched process.
//
// If application launching occurs via some non-spawn mechanism (e.g. D-Bus
// activation) then stdin_fd, stdout_fd and stderr_fd are ignored.
func (appinfo *DesktopAppInfo) LaunchURIsAsManagerWithFds(uris []string, launchContext *AppLaunchContext, spawnFlags glib.SpawnFlags, userSetup glib.SpawnChildSetupFunc, pidCallback DesktopAppLaunchCallback, stdinFd int, stdoutFd int, stderrFd int) error {
	var _arg0 *C.GDesktopAppInfo     // out
	var _arg1 *C.GList               // out
	var _arg2 *C.GAppLaunchContext   // out
	var _arg3 C.GSpawnFlags          // out
	var _arg4 C.GSpawnChildSetupFunc // out
	var _arg5 C.gpointer
	var _arg6 C.GDesktopAppLaunchCallback // out
	var _arg7 C.gpointer
	var _arg8 C.gint    // out
	var _arg9 C.gint    // out
	var _arg10 C.gint   // out
	var _cerr *C.GError // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(appinfo.Native()))
	for i := len(uris) - 1; i >= 0; i-- {
		src := uris[i]
		var dst *C.gchar // out
		dst = (*C.gchar)(unsafe.Pointer(C.CString(src)))
		defer C.free(unsafe.Pointer(dst))
		_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg1)
	if launchContext != nil {
		_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(launchContext.Native()))
	}
	_arg3 = C.GSpawnFlags(spawnFlags)
	if userSetup != nil {
		_arg4 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg5 = C.gpointer(gbox.AssignOnce(userSetup))
	}
	if pidCallback != nil {
		_arg6 = (*[0]byte)(C._gotk4_gio2_DesktopAppLaunchCallback)
		_arg7 = C.gpointer(gbox.Assign(pidCallback))
		defer gbox.Delete(uintptr(_arg7))
	}
	_arg8 = C.gint(stdinFd)
	_arg9 = C.gint(stdoutFd)
	_arg10 = C.gint(stderrFd)

	C.g_desktop_app_info_launch_uris_as_manager_with_fds(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ListActions returns the list of "additional application actions" supported on
// the desktop file, as per the desktop file specification.
//
// As per the specification, this is the list of actions that are explicitly
// listed in the "Actions" key of the [Desktop Entry] group.
func (info *DesktopAppInfo) ListActions() []string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_list_actions(_arg0)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// DesktopAppInfoGetImplementations gets all applications that implement
// interface.
//
// An application implements an interface if that interface is listed in the
// Implements= line of the desktop file of the application.
func DesktopAppInfoGetImplementations(_interface string) []DesktopAppInfo {
	var _arg1 *C.gchar // out
	var _cret *C.GList // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_interface)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_implementations(_arg1)

	var _list []DesktopAppInfo // out

	_list = make([]DesktopAppInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDesktopAppInfo)(v)
		var dst DesktopAppInfo // out
		dst = *wrapDesktopAppInfo(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// DesktopAppInfoSetDesktopEnv sets the name of the desktop that the application
// is running in. This is used by g_app_info_should_show() and
// g_desktop_app_info_get_show_in() to evaluate the OnlyShowIn and NotShowIn
// desktop entry fields.
//
// Should be called only once; subsequent calls are ignored.
//
// Deprecated: do not use this API. Since 2.42 the value of the
// XDG_CURRENT_DESKTOP environment variable will be used.
func DesktopAppInfoSetDesktopEnv(desktopEnv string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(desktopEnv)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_desktop_app_info_set_desktop_env(_arg1)
}
