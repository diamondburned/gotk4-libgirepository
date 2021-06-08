// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.g_application_command_line_get_type()), F: marshalApplicationCommandLine},
	})
}

// ApplicationCommandLine represents a command-line invocation of an
// application. It is created by #GApplication and emitted in the
// #GApplication::command-line signal and virtual function.
//
// The class contains the list of arguments that the program was invoked with.
// It is also possible to query if the commandline invocation was local (ie: the
// current process is running in direct response to the invocation) or remote
// (ie: some other process forwarded the commandline to this process).
//
// The GApplicationCommandLine object can provide the @argc and @argv parameters
// for use with the Context command-line parsing API, with the
// g_application_command_line_get_arguments() function. See
// [gapplication-example-cmdline3.c][gapplication-example-cmdline3] for an
// example.
//
// The exit status of the originally-invoked process may be set and messages can
// be printed to stdout or stderr of that process. The lifecycle of the
// originally-invoked process is tied to the lifecycle of this object (ie: the
// process exits when the last reference is dropped).
//
// The main use for CommandLine (and the #GApplication::command-line signal) is
// 'Emacs server' like use cases: You can set the `EDITOR` environment variable
// to have e.g. git use your favourite editor to edit commit messages, and if
// you already have an instance of the editor running, the editing will happen
// in the running instance, instead of opening a new one. An important aspect of
// this use case is that the process that gets started by git does not return
// until the editing is done.
//
// Normally, the commandline is completely handled in the
// #GApplication::command-line handler. The launching instance exits once the
// signal handler in the primary instance has returned, and the return value of
// the signal handler becomes the exit status of the launching instance.
//
//    static gboolean
//    my_cmdline_handler (gpointer data)
//    {
//      GApplicationCommandLine *cmdline = data;
//
//      // do the heavy lifting in an idle
//
//      g_application_command_line_set_exit_status (cmdline, 0);
//      g_object_unref (cmdline); // this releases the application
//
//      return G_SOURCE_REMOVE;
//    }
//
//    static int
//    command_line (GApplication            *application,
//                  GApplicationCommandLine *cmdline)
//    {
//      // keep the application running until we are done with this commandline
//      g_application_hold (application);
//
//      g_object_set_data_full (G_OBJECT (cmdline),
//                              "application", application,
//                              (GDestroyNotify)g_application_release);
//
//      g_object_ref (cmdline);
//      g_idle_add (my_cmdline_handler, cmdline);
//
//      return 0;
//    }
//
// In this example the commandline is not completely handled before the
// #GApplication::command-line handler returns. Instead, we keep a reference to
// the CommandLine object and handle it later (in this example, in an idle).
// Note that it is necessary to hold the application until you are done with the
// commandline.
//
// The complete example can be found here: gapplication-example-cmdline3.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline3.c)
type ApplicationCommandLine interface {
	gextras.Objector

	// CreateFileForArg creates a #GFile corresponding to a filename that was
	// given as part of the invocation of @cmdline.
	//
	// This differs from g_file_new_for_commandline_arg() in that it resolves
	// relative pathnames using the current working directory of the invoking
	// process rather than the local process.
	CreateFileForArg(arg string) File
	// Arguments gets the list of arguments that was passed on the command line.
	//
	// The strings in the array may contain non-UTF-8 data on UNIX (such as
	// filenames or arguments given in the system locale) but are always in
	// UTF-8 on Windows.
	//
	// If you wish to use the return value with Context, you must use
	// g_option_context_parse_strv().
	//
	// The return value is nil-terminated and should be freed using
	// g_strfreev().
	Arguments() []string
	// Cwd gets the working directory of the command line invocation. The string
	// may contain non-utf8 data.
	//
	// It is possible that the remote application did not send a working
	// directory, so this may be nil.
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	Cwd() string
	// Environ gets the contents of the 'environ' variable of the command line
	// invocation, as would be returned by g_get_environ(), ie as a
	// nil-terminated list of strings in the form 'NAME=VALUE'. The strings may
	// contain non-utf8 data.
	//
	// The remote application usually does not send an environment. Use
	// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
	// is possible that the environment is still not available (due to
	// invocation messages from other applications).
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	//
	// See g_application_command_line_getenv() if you are only interested in the
	// value of a single environment variable.
	Environ() []string
	// ExitStatus gets the exit status of @cmdline. See
	// g_application_command_line_set_exit_status() for more information.
	ExitStatus() int
	// IsRemote determines if @cmdline represents a remote invocation.
	IsRemote() bool
	// OptionsDict gets the options there were passed to
	// g_application_command_line().
	//
	// If you did not override local_command_line() then these are the same
	// options that were parsed according to the Entrys added to the application
	// with g_application_add_main_option_entries() and possibly modified from
	// your GApplication::handle-local-options handler.
	//
	// If no options were sent then an empty dictionary is returned so that you
	// don't need to check for nil.
	OptionsDict() *glib.VariantDict
	// PlatformData gets the platform data associated with the invocation of
	// @cmdline.
	//
	// This is a #GVariant dictionary containing information about the context
	// in which the invocation occurred. It typically contains information like
	// the current working directory and the startup notification ID.
	//
	// For local invocation, it will be nil.
	PlatformData() *glib.Variant
	// Stdin gets the stdin of the invoking process.
	//
	// The Stream can be used to read data passed to the standard input of the
	// invoking process. This doesn't work on all platforms. Presently, it is
	// only available on UNIX when using a D-Bus daemon capable of passing file
	// descriptors. If stdin is not available then nil will be returned. In the
	// future, support may be expanded to other platforms.
	//
	// You must only call this function once per commandline invocation.
	Stdin() InputStream
	// env gets the value of a particular environment variable of the command
	// line invocation, as would be returned by g_getenv(). The strings may
	// contain non-utf8 data.
	//
	// The remote application usually does not send an environment. Use
	// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
	// is possible that the environment is still not available (due to
	// invocation messages from other applications).
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	env(name string) string
	// SetExitStatus sets the exit status that will be used when the invoking
	// process exits.
	//
	// The return value of the #GApplication::command-line signal is passed to
	// this function when the handler returns. This is the usual way of setting
	// the exit status.
	//
	// In the event that you want the remote invocation to continue running and
	// want to decide on the exit status in the future, you can use this call.
	// For the case of a remote invocation, the remote process will typically
	// exit when the last reference is dropped on @cmdline. The exit status of
	// the remote process will be equal to the last value that was set with this
	// function.
	//
	// In the case that the commandline invocation is local, the situation is
	// slightly more complicated. If the commandline invocation results in the
	// mainloop running (ie: because the use-count of the application increased
	// to a non-zero value) then the application is considered to have been
	// 'successful' in a certain sense, and the exit status is always zero. If
	// the application use count is zero, though, the exit status of the local
	// CommandLine is used.
	SetExitStatus(exitStatus int)
}

// applicationCommandLine implements the ApplicationCommandLine interface.
type applicationCommandLine struct {
	gextras.Objector
}

var _ ApplicationCommandLine = (*applicationCommandLine)(nil)

// WrapApplicationCommandLine wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplicationCommandLine(obj *externglib.Object) ApplicationCommandLine {
	return ApplicationCommandLine{
		Objector: obj,
	}
}

func marshalApplicationCommandLine(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplicationCommandLine(obj), nil
}

// CreateFileForArg creates a #GFile corresponding to a filename that was
// given as part of the invocation of @cmdline.
//
// This differs from g_file_new_for_commandline_arg() in that it resolves
// relative pathnames using the current working directory of the invoking
// process rather than the local process.
func (c applicationCommandLine) CreateFileForArg(arg string) File {
	var arg0 *C.GApplicationCommandLine
	var arg1 *C.gchar

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(arg))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.GFile)
	var goret File

	cret = C.g_application_command_line_create_file_for_arg(arg0, arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(File)

	return goret
}

// Arguments gets the list of arguments that was passed on the command line.
//
// The strings in the array may contain non-UTF-8 data on UNIX (such as
// filenames or arguments given in the system locale) but are always in
// UTF-8 on Windows.
//
// If you wish to use the return value with Context, you must use
// g_option_context_parse_strv().
//
// The return value is nil-terminated and should be freed using
// g_strfreev().
func (c applicationCommandLine) Arguments() []string {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret **C.gchar
	var arg1 *C.int
	var goret []string

	cret = C.g_application_command_line_get_arguments(arg0, arg1)

	goret = make([]string, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		goret[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}

	return ret1, goret
}

// Cwd gets the working directory of the command line invocation. The string
// may contain non-utf8 data.
//
// It is possible that the remote application did not send a working
// directory, so this may be nil.
//
// The return value should not be modified or freed and is valid for as long
// as @cmdline exists.
func (c applicationCommandLine) Cwd() string {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret *C.gchar
	var goret string

	cret = C.g_application_command_line_get_cwd(arg0)

	goret = C.GoString(cret)

	return goret
}

// Environ gets the contents of the 'environ' variable of the command line
// invocation, as would be returned by g_get_environ(), ie as a
// nil-terminated list of strings in the form 'NAME=VALUE'. The strings may
// contain non-utf8 data.
//
// The remote application usually does not send an environment. Use
// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
// is possible that the environment is still not available (due to
// invocation messages from other applications).
//
// The return value should not be modified or freed and is valid for as long
// as @cmdline exists.
//
// See g_application_command_line_getenv() if you are only interested in the
// value of a single environment variable.
func (c applicationCommandLine) Environ() []string {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret **C.gchar
	var goret []string

	cret = C.g_application_command_line_get_environ(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret[i] = C.GoString(src)
		}
	}

	return goret
}

// ExitStatus gets the exit status of @cmdline. See
// g_application_command_line_set_exit_status() for more information.
func (c applicationCommandLine) ExitStatus() int {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret C.int
	var goret int

	cret = C.g_application_command_line_get_exit_status(arg0)

	goret = int(cret)

	return goret
}

// IsRemote determines if @cmdline represents a remote invocation.
func (c applicationCommandLine) IsRemote() bool {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_application_command_line_get_is_remote(arg0)

	if cret {
		goret = true
	}

	return goret
}

// OptionsDict gets the options there were passed to
// g_application_command_line().
//
// If you did not override local_command_line() then these are the same
// options that were parsed according to the Entrys added to the application
// with g_application_add_main_option_entries() and possibly modified from
// your GApplication::handle-local-options handler.
//
// If no options were sent then an empty dictionary is returned so that you
// don't need to check for nil.
func (c applicationCommandLine) OptionsDict() *glib.VariantDict {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	var cret *C.GVariantDict
	var goret *glib.VariantDict

	cret = C.g_application_command_line_get_options_dict(arg0)

	goret = glib.WrapVariantDict(unsafe.Pointer(cret))

	return goret
}

// PlatformData gets the platform data associated with the invocation of
// @cmdline.
//
// This is a #GVariant dictionary containing information about the context
// in which the invocation occurred. It typically contains information like
// the current working directory and the startup notification ID.
//
// For local invocation, it will be nil.
func (c applicationCommandLine) PlatformData() *glib.Variant {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	cret := new(C.GVariant)
	var goret *glib.Variant

	cret = C.g_application_command_line_get_platform_data(arg0)

	goret = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Stdin gets the stdin of the invoking process.
//
// The Stream can be used to read data passed to the standard input of the
// invoking process. This doesn't work on all platforms. Presently, it is
// only available on UNIX when using a D-Bus daemon capable of passing file
// descriptors. If stdin is not available then nil will be returned. In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
func (c applicationCommandLine) Stdin() InputStream {
	var arg0 *C.GApplicationCommandLine

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	cret := new(C.GInputStream)
	var goret InputStream

	cret = C.g_application_command_line_get_stdin(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(InputStream)

	return goret
}

// env gets the value of a particular environment variable of the command
// line invocation, as would be returned by g_getenv(). The strings may
// contain non-utf8 data.
//
// The remote application usually does not send an environment. Use
// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
// is possible that the environment is still not available (due to
// invocation messages from other applications).
//
// The return value should not be modified or freed and is valid for as long
// as @cmdline exists.
func (c applicationCommandLine) env(name string) string {
	var arg0 *C.GApplicationCommandLine
	var arg1 *C.gchar

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret string

	cret = C.g_application_command_line_getenv(arg0, arg1)

	goret = C.GoString(cret)

	return goret
}

// SetExitStatus sets the exit status that will be used when the invoking
// process exits.
//
// The return value of the #GApplication::command-line signal is passed to
// this function when the handler returns. This is the usual way of setting
// the exit status.
//
// In the event that you want the remote invocation to continue running and
// want to decide on the exit status in the future, you can use this call.
// For the case of a remote invocation, the remote process will typically
// exit when the last reference is dropped on @cmdline. The exit status of
// the remote process will be equal to the last value that was set with this
// function.
//
// In the case that the commandline invocation is local, the situation is
// slightly more complicated. If the commandline invocation results in the
// mainloop running (ie: because the use-count of the application increased
// to a non-zero value) then the application is considered to have been
// 'successful' in a certain sense, and the exit status is always zero. If
// the application use count is zero, though, the exit status of the local
// CommandLine is used.
func (c applicationCommandLine) SetExitStatus(exitStatus int) {
	var arg0 *C.GApplicationCommandLine
	var arg1 C.int

	arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	arg1 = C.int(exitStatus)

	C.g_application_command_line_set_exit_status(arg0, arg1)
}
