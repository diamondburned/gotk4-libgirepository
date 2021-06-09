// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
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
		{T: externglib.Type(C.g_desktop_app_info_lookup_get_type()), F: marshalDesktopAppInfoLookup},
		{T: externglib.Type(C.g_desktop_app_info_get_type()), F: marshalDesktopAppInfo},
	})
}

// DesktopAppLaunchCallback: during invocation,
// g_desktop_app_info_launch_uris_as_manager() may create one or more child
// processes. This callback is invoked once for each, providing the process ID.
type DesktopAppLaunchCallback func()

//export gotk4_DesktopAppLaunchCallback
func gotk4_DesktopAppLaunchCallback(arg0 *C.GDesktopAppInfo, arg1 C.GPid, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(DesktopAppLaunchCallback)
	fn()
}

// DesktopAppInfoLookupOverrider contains methods that are overridable. This
// interface is a subset of the interface DesktopAppInfoLookup.
type DesktopAppInfoLookupOverrider interface {
	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	//
	// The AppInfoLookup interface and this function is used to implement
	// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
	// is no reason for applications to use it directly. Applications should use
	// g_app_info_get_default_for_uri_scheme().
	DefaultForURIScheme(uriScheme string) AppInfo
}

// DesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
type DesktopAppInfoLookup interface {
	gextras.Objector
	DesktopAppInfoLookupOverrider
}

// desktopAppInfoLookup implements the DesktopAppInfoLookup interface.
type desktopAppInfoLookup struct {
	gextras.Objector
}

var _ DesktopAppInfoLookup = (*desktopAppInfoLookup)(nil)

// WrapDesktopAppInfoLookup wraps a GObject to a type that implements interface
// DesktopAppInfoLookup. It is primarily used internally.
func WrapDesktopAppInfoLookup(obj *externglib.Object) DesktopAppInfoLookup {
	return DesktopAppInfoLookup{
		Objector: obj,
	}
}

func marshalDesktopAppInfoLookup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDesktopAppInfoLookup(obj), nil
}

// DefaultForURIScheme gets the default application for launching
// applications using this URI scheme for a particular AppInfoLookup
// implementation.
//
// The AppInfoLookup interface and this function is used to implement
// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
// is no reason for applications to use it directly. Applications should use
// g_app_info_get_default_for_uri_scheme().
func (l desktopAppInfoLookup) DefaultForURIScheme(uriScheme string) AppInfo {
	var _arg0 *C.GDesktopAppInfoLookup
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfoLookup)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.char)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GAppInfo

	cret = C.g_desktop_app_info_lookup_get_default_for_uri_scheme(_arg0, _arg1)

	var _appInfo AppInfo

	_appInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(AppInfo)

	return _appInfo
}

// DesktopAppInfo is an implementation of Info based on desktop files.
//
// Note that `<gio/gdesktopappinfo.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type DesktopAppInfo interface {
	gextras.Objector
	AppInfo

	// ActionName gets the user-visible display name of the "additional
	// application action" specified by @action_name.
	//
	// This corresponds to the "Name" key within the keyfile group for the
	// action.
	ActionName(actionName string) string
	// Boolean looks up a boolean value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	Boolean(key string) bool
	// Categories gets the categories from the desktop file.
	Categories() string
	// Filename: when @info was created from a known filename, return it. In
	// some situations such as the AppInfo returned from
	// g_desktop_app_info_new_from_keyfile(), this function will return nil.
	Filename() *string
	// GenericName gets the generic name from the destkop file.
	GenericName() string
	// IsHidden: a desktop file is hidden if the Hidden key in it is set to
	// True.
	IsHidden() bool
	// Keywords gets the keywords from the desktop file.
	Keywords() []string
	// LocaleString looks up a localized string value in the keyfile backing
	// @info translated to the current locale.
	//
	// The @key is looked up in the "Desktop Entry" group.
	LocaleString(key string) string
	// Nodisplay gets the value of the NoDisplay key, which helps determine if
	// the application info should be shown in menus. See
	// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
	Nodisplay() bool
	// ShowIn checks if the application info should be shown in menus that list
	// available applications for a specific name of the desktop, based on the
	// `OnlyShowIn` and `NotShowIn` keys.
	//
	// @desktop_env should typically be given as nil, in which case the
	// `XDG_CURRENT_DESKTOP` environment variable is consulted. If you want to
	// override the default mechanism then you may specify @desktop_env, but
	// this is not recommended.
	//
	// Note that g_app_info_should_show() for @info will include this check
	// (with nil for @desktop_env) as well as additional checks.
	ShowIn(desktopEnv string) bool
	// StartupWmClass retrieves the StartupWMClass field from @info. This
	// represents the WM_CLASS property of the main window of the application,
	// if launched through @info.
	StartupWmClass() string
	// String looks up a string value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	String(key string) string
	// StringList looks up a string list value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	StringList(key string) []string
	// HasKey returns whether @key exists in the "Desktop Entry" group of the
	// keyfile backing @info.
	HasKey(key string) bool
	// LaunchAction activates the named application action.
	//
	// You may only call this function on action names that were returned from
	// g_desktop_app_info_list_actions().
	//
	// Note that if the main entry of the desktop file indicates that the
	// application supports startup notification, and @launch_context is
	// non-nil, then startup notification will be used when activating the
	// action (and as such, invocation of the action on the receiving side must
	// signal the end of startup notification when it is completed). This is the
	// expected behaviour of applications declaring additional actions, as per
	// the desktop file specification.
	//
	// As with g_app_info_launch() there is no way to detect failures that occur
	// while using this function.
	LaunchAction(actionName string, launchContext AppLaunchContext)
	// LaunchUrisAsManager: this function performs the equivalent of
	// g_app_info_launch_uris(), but is intended primarily for operating system
	// components that launch applications. Ordinary applications should use
	// g_app_info_launch_uris().
	//
	// If the application is launched via GSpawn, then @spawn_flags, @user_setup
	// and @user_setup_data are used for the call to g_spawn_async().
	// Additionally, @pid_callback (with @pid_callback_data) will be called to
	// inform about the PID of the created process. See
	// g_spawn_async_with_pipes() for information on certain parameter
	// conditions that can enable an optimized posix_spawn() codepath to be
	// used.
	//
	// If application launching occurs via some other mechanism (eg: D-Bus
	// activation) then @spawn_flags, @user_setup, @user_setup_data,
	// @pid_callback and @pid_callback_data are ignored.
	LaunchUrisAsManager() error
	// ListActions returns the list of "additional application actions"
	// supported on the desktop file, as per the desktop file specification.
	//
	// As per the specification, this is the list of actions that are explicitly
	// listed in the "Actions" key of the [Desktop Entry] group.
	ListActions() []string
}

// desktopAppInfo implements the DesktopAppInfo interface.
type desktopAppInfo struct {
	gextras.Objector
	AppInfo
}

var _ DesktopAppInfo = (*desktopAppInfo)(nil)

// WrapDesktopAppInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapDesktopAppInfo(obj *externglib.Object) DesktopAppInfo {
	return DesktopAppInfo{
		Objector: obj,
		AppInfo:  WrapAppInfo(obj),
	}
}

func marshalDesktopAppInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDesktopAppInfo(obj), nil
}

// NewDesktopAppInfo constructs a class DesktopAppInfo.
func NewDesktopAppInfo(desktopId string) DesktopAppInfo {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(desktopId))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GDesktopAppInfo

	cret = C.g_desktop_app_info_new(_arg1)

	var _desktopAppInfo DesktopAppInfo

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromFilename constructs a class DesktopAppInfo.
func NewDesktopAppInfoFromFilename(filename *string) DesktopAppInfo {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GDesktopAppInfo

	cret = C.g_desktop_app_info_new_from_filename(_arg1)

	var _desktopAppInfo DesktopAppInfo

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromKeyfile constructs a class DesktopAppInfo.
func NewDesktopAppInfoFromKeyfile(keyFile *glib.KeyFile) DesktopAppInfo {
	var _arg1 *C.GKeyFile

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))

	var _cret C.GDesktopAppInfo

	cret = C.g_desktop_app_info_new_from_keyfile(_arg1)

	var _desktopAppInfo DesktopAppInfo

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// ActionName gets the user-visible display name of the "additional
// application action" specified by @action_name.
//
// This corresponds to the "Name" key within the keyfile group for the
// action.
func (i desktopAppInfo) ActionName(actionName string) string {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.gchar

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	cret = C.g_desktop_app_info_get_action_name(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Boolean looks up a boolean value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) Boolean(key string) bool {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_desktop_app_info_get_boolean(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Categories gets the categories from the desktop file.
func (i desktopAppInfo) Categories() string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_categories(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Filename: when @info was created from a known filename, return it. In
// some situations such as the AppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return nil.
func (i desktopAppInfo) Filename() *string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_filename(_arg0)

	var _filename *string

	_filename = C.GoString(_cret)

	return _filename
}

// GenericName gets the generic name from the destkop file.
func (i desktopAppInfo) GenericName() string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_generic_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsHidden: a desktop file is hidden if the Hidden key in it is set to
// True.
func (i desktopAppInfo) IsHidden() bool {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_desktop_app_info_get_is_hidden(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Keywords gets the keywords from the desktop file.
func (i desktopAppInfo) Keywords() []string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret **C.char

	cret = C.g_desktop_app_info_get_keywords(_arg0)

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

// LocaleString looks up a localized string value in the keyfile backing
// @info translated to the current locale.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) LocaleString(key string) string {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_locale_string(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Nodisplay gets the value of the NoDisplay key, which helps determine if
// the application info should be shown in menus. See
// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
func (i desktopAppInfo) Nodisplay() bool {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_desktop_app_info_get_nodisplay(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowIn checks if the application info should be shown in menus that list
// available applications for a specific name of the desktop, based on the
// `OnlyShowIn` and `NotShowIn` keys.
//
// @desktop_env should typically be given as nil, in which case the
// `XDG_CURRENT_DESKTOP` environment variable is consulted. If you want to
// override the default mechanism then you may specify @desktop_env, but
// this is not recommended.
//
// Note that g_app_info_should_show() for @info will include this check
// (with nil for @desktop_env) as well as additional checks.
func (i desktopAppInfo) ShowIn(desktopEnv string) bool {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.gchar

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(desktopEnv))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_desktop_app_info_get_show_in(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// StartupWmClass retrieves the StartupWMClass field from @info. This
// represents the WM_CLASS property of the main window of the application,
// if launched through @info.
func (i desktopAppInfo) StartupWmClass() string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_startup_wm_class(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// String looks up a string value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) String(key string) string {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char

	cret = C.g_desktop_app_info_get_string(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// StringList looks up a string list value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) StringList(key string) []string {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.gchar
	var _arg2 *C.gsize

	cret = C.g_desktop_app_info_get_string_list(_arg0, _arg1)

	var _utf8s []string

	{
		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg2))

		_utf8s = make([]string, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// HasKey returns whether @key exists in the "Desktop Entry" group of the
// keyfile backing @info.
func (i desktopAppInfo) HasKey(key string) bool {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_desktop_app_info_has_key(_arg0, _arg1)

	var _ok bool

	if _cret {
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
// application supports startup notification, and @launch_context is
// non-nil, then startup notification will be used when activating the
// action (and as such, invocation of the action on the receiving side must
// signal the end of startup notification when it is completed). This is the
// expected behaviour of applications declaring additional actions, as per
// the desktop file specification.
//
// As with g_app_info_launch() there is no way to detect failures that occur
// while using this function.
func (i desktopAppInfo) LaunchAction(actionName string, launchContext AppLaunchContext) {
	var _arg0 *C.GDesktopAppInfo
	var _arg1 *C.gchar
	var _arg2 *C.GAppLaunchContext

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(launchContext.Native()))

	C.g_desktop_app_info_launch_action(_arg0, _arg1, _arg2)
}

// LaunchUrisAsManager: this function performs the equivalent of
// g_app_info_launch_uris(), but is intended primarily for operating system
// components that launch applications. Ordinary applications should use
// g_app_info_launch_uris().
//
// If the application is launched via GSpawn, then @spawn_flags, @user_setup
// and @user_setup_data are used for the call to g_spawn_async().
// Additionally, @pid_callback (with @pid_callback_data) will be called to
// inform about the PID of the created process. See
// g_spawn_async_with_pipes() for information on certain parameter
// conditions that can enable an optimized posix_spawn() codepath to be
// used.
//
// If application launching occurs via some other mechanism (eg: D-Bus
// activation) then @spawn_flags, @user_setup, @user_setup_data,
// @pid_callback and @pid_callback_data are ignored.
func (a desktopAppInfo) LaunchUrisAsManager() error {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(a.Native()))

	var _cerr *C.GError

	C.g_desktop_app_info_launch_uris_as_manager(_arg0, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ListActions returns the list of "additional application actions"
// supported on the desktop file, as per the desktop file specification.
//
// As per the specification, this is the list of actions that are explicitly
// listed in the "Actions" key of the [Desktop Entry] group.
func (i desktopAppInfo) ListActions() []string {
	var _arg0 *C.GDesktopAppInfo

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret **C.gchar

	cret = C.g_desktop_app_info_list_actions(_arg0)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
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
