// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_glib2_SpawnChildSetupFunc(gpointer);
import "C"

// SpawnError: error codes returned by spawning processes.
type SpawnError C.gint

const (
	// SpawnErrorFork: fork failed due to lack of memory.
	SpawnErrorFork SpawnError = 0
	// SpawnErrorRead: read or select on pipes failed.
	SpawnErrorRead SpawnError = 1
	// SpawnErrorChdir: changing to working directory failed.
	SpawnErrorChdir SpawnError = 2
	// SpawnErrorAcces: execv() returned EACCES.
	SpawnErrorAcces SpawnError = 3
	// SpawnErrorPerm: execv() returned EPERM.
	SpawnErrorPerm SpawnError = 4
	// SpawnErrorTooBig: execv() returned E2BIG.
	SpawnErrorTooBig SpawnError = 5
	// SpawnError2Big: deprecated alias for G_SPAWN_ERROR_TOO_BIG (deprecated
	// since GLib 2.32).
	SpawnError2Big SpawnError = 5
	// SpawnErrorNoexec: execv() returned ENOEXEC.
	SpawnErrorNoexec SpawnError = 6
	// SpawnErrorNametoolong: execv() returned ENAMETOOLONG.
	SpawnErrorNametoolong SpawnError = 7
	// SpawnErrorNoent: execv() returned ENOENT.
	SpawnErrorNoent SpawnError = 8
	// SpawnErrorNOMEM: execv() returned ENOMEM.
	SpawnErrorNOMEM SpawnError = 9
	// SpawnErrorNotdir: execv() returned ENOTDIR.
	SpawnErrorNotdir SpawnError = 10
	// SpawnErrorLoop: execv() returned ELOOP.
	SpawnErrorLoop SpawnError = 11
	// SpawnErrorTxtbusy: execv() returned ETXTBUSY.
	SpawnErrorTxtbusy SpawnError = 12
	// SpawnErrorIO: execv() returned EIO.
	SpawnErrorIO SpawnError = 13
	// SpawnErrorNfile: execv() returned ENFILE.
	SpawnErrorNfile SpawnError = 14
	// SpawnErrorMfile: execv() returned EMFILE.
	SpawnErrorMfile SpawnError = 15
	// SpawnErrorInval: execv() returned EINVAL.
	SpawnErrorInval SpawnError = 16
	// SpawnErrorIsdir: execv() returned EISDIR.
	SpawnErrorIsdir SpawnError = 17
	// SpawnErrorLibbad: execv() returned ELIBBAD.
	SpawnErrorLibbad SpawnError = 18
	// SpawnErrorFailed: some other fatal failure, error->message should
	// explain.
	SpawnErrorFailed SpawnError = 19
)

// String returns the name in string for SpawnError.
func (s SpawnError) String() string {
	switch s {
	case SpawnErrorFork:
		return "Fork"
	case SpawnErrorRead:
		return "Read"
	case SpawnErrorChdir:
		return "Chdir"
	case SpawnErrorAcces:
		return "Acces"
	case SpawnErrorPerm:
		return "Perm"
	case SpawnErrorTooBig:
		return "TooBig"
	case SpawnErrorNoexec:
		return "Noexec"
	case SpawnErrorNametoolong:
		return "Nametoolong"
	case SpawnErrorNoent:
		return "Noent"
	case SpawnErrorNOMEM:
		return "NOMEM"
	case SpawnErrorNotdir:
		return "Notdir"
	case SpawnErrorLoop:
		return "Loop"
	case SpawnErrorTxtbusy:
		return "Txtbusy"
	case SpawnErrorIO:
		return "IO"
	case SpawnErrorNfile:
		return "Nfile"
	case SpawnErrorMfile:
		return "Mfile"
	case SpawnErrorInval:
		return "Inval"
	case SpawnErrorIsdir:
		return "Isdir"
	case SpawnErrorLibbad:
		return "Libbad"
	case SpawnErrorFailed:
		return "Failed"
	default:
		return fmt.Sprintf("SpawnError(%d)", s)
	}
}

// SpawnFlags flags passed to g_spawn_sync(), g_spawn_async() and
// g_spawn_async_with_pipes().
type SpawnFlags C.guint

const (
	// SpawnDefault: no flags, default behaviour.
	SpawnDefault SpawnFlags = 0b0
	// SpawnLeaveDescriptorsOpen parent's open file descriptors will be
	// inherited by the child; otherwise all descriptors except stdin, stdout
	// and stderr will be closed before calling exec() in the child.
	SpawnLeaveDescriptorsOpen SpawnFlags = 0b1
	// SpawnDoNotReapChild: child will not be automatically reaped; you must use
	// g_child_watch_add() yourself (or call waitpid() or handle SIGCHLD
	// yourself), or the child will become a zombie.
	SpawnDoNotReapChild SpawnFlags = 0b10
	// SpawnSearchPath: argv[0] need not be an absolute path, it will be looked
	// for in the user's PATH.
	SpawnSearchPath SpawnFlags = 0b100
	// SpawnStdoutToDevNull child's standard output will be discarded, instead
	// of going to the same location as the parent's standard output.
	SpawnStdoutToDevNull SpawnFlags = 0b1000
	// SpawnStderrToDevNull child's standard error will be discarded.
	SpawnStderrToDevNull SpawnFlags = 0b10000
	// SpawnChildInheritsStdin: child will inherit the parent's standard input
	// (by default, the child's standard input is attached to /dev/null).
	SpawnChildInheritsStdin SpawnFlags = 0b100000
	// SpawnFileAndArgvZero: first element of argv is the file to execute, while
	// the remaining elements are the actual argument vector to pass to the
	// file. Normally g_spawn_async_with_pipes() uses argv[0] as the file to
	// execute, and passes all of argv to the child.
	SpawnFileAndArgvZero SpawnFlags = 0b1000000
	// SpawnSearchPathFromEnvp: if argv[0] is not an absolute path, it will be
	// looked for in the PATH from the passed child environment. Since: 2.34.
	SpawnSearchPathFromEnvp SpawnFlags = 0b10000000
	// SpawnCloexecPipes: create all pipes with the O_CLOEXEC flag set. Since:
	// 2.40.
	SpawnCloexecPipes SpawnFlags = 0b100000000
)

// String returns the names in string for SpawnFlags.
func (s SpawnFlags) String() string {
	if s == 0 {
		return "SpawnFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(203)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SpawnDefault:
			builder.WriteString("Default|")
		case SpawnLeaveDescriptorsOpen:
			builder.WriteString("LeaveDescriptorsOpen|")
		case SpawnDoNotReapChild:
			builder.WriteString("DoNotReapChild|")
		case SpawnSearchPath:
			builder.WriteString("SearchPath|")
		case SpawnStdoutToDevNull:
			builder.WriteString("StdoutToDevNull|")
		case SpawnStderrToDevNull:
			builder.WriteString("StderrToDevNull|")
		case SpawnChildInheritsStdin:
			builder.WriteString("ChildInheritsStdin|")
		case SpawnFileAndArgvZero:
			builder.WriteString("FileAndArgvZero|")
		case SpawnSearchPathFromEnvp:
			builder.WriteString("SearchPathFromEnvp|")
		case SpawnCloexecPipes:
			builder.WriteString("CloexecPipes|")
		default:
			builder.WriteString(fmt.Sprintf("SpawnFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpawnFlags) Has(other SpawnFlags) bool {
	return (s & other) == other
}

// SpawnChildSetupFunc specifies the type of the setup function passed to
// g_spawn_async(), g_spawn_sync() and g_spawn_async_with_pipes(), which can, in
// very limited ways, be used to affect the child's execution.
//
// On POSIX platforms, the function is called in the child after GLib has
// performed all the setup it plans to perform, but before calling exec().
// Actions taken in this function will only affect the child, not the parent.
//
// On Windows, the function is called in the parent. Its usefulness on Windows
// is thus questionable. In many cases executing the child setup function in the
// parent can have ill effects, and you should be very careful when porting
// software to Windows that uses child setup functions.
//
// However, even on POSIX, you are extremely limited in what you can safely do
// from a ChildSetupFunc, because any mutexes that were held by other threads in
// the parent process at the time of the fork() will still be locked in the
// child process, and they will never be unlocked (since the threads that held
// them don't exist in the child). POSIX allows only async-signal-safe functions
// (see signal(7)) to be called in the child between fork() and exec(), which
// drastically limits the usefulness of child setup functions.
//
// In particular, it is not safe to call any function which may call malloc(),
// which includes POSIX functions such as setenv(). If you need to set up the
// child environment differently from the parent, you should use
// g_get_environ(), g_environ_setenv(), and g_environ_unsetenv(), and then pass
// the complete environment list to the g_spawn... function.
type SpawnChildSetupFunc func()

//export _gotk4_glib2_SpawnChildSetupFunc
func _gotk4_glib2_SpawnChildSetupFunc(arg1 C.gpointer) {
	var fn SpawnChildSetupFunc
	{
		v := gbox.Get(uintptr(arg1))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SpawnChildSetupFunc)
	}

	fn()
}

// SpawnCheckExitStatus: set error if exit_status indicates the child exited
// abnormally (e.g. with a nonzero exit code, or via a fatal signal).
//
// The g_spawn_sync() and g_child_watch_add() family of APIs return an exit
// status for subprocesses encoded in a platform-specific way. On Unix, this is
// guaranteed to be in the same format waitpid() returns, and on Windows it is
// guaranteed to be the result of GetExitCodeProcess().
//
// Prior to the introduction of this function in GLib 2.34, interpreting
// exit_status required use of platform-specific APIs, which is problematic for
// software using GLib as a cross-platform layer.
//
// Additionally, many programs simply want to determine whether or not the child
// exited successfully, and either propagate a #GError or print a message to
// standard error. In that common case, this function can be used. Note that the
// error message in error will contain human-readable information about the exit
// status.
//
// The domain and code of error have special semantics in the case where the
// process has an "exit code", as opposed to being killed by a signal. On Unix,
// this happens if WIFEXITED() would be true of exit_status. On Windows, it is
// always the case.
//
// The special semantics are that the actual exit code will be the code set in
// error, and the domain will be G_SPAWN_EXIT_ERROR. This allows you to
// differentiate between different exit codes.
//
// If the process was terminated by some means other than an exit status, the
// domain will be G_SPAWN_ERROR, and the code will be G_SPAWN_ERROR_FAILED.
//
// This function just offers convenience; you can of course also check the
// available platform via a macro such as G_OS_UNIX, and use WIFEXITED() and
// WEXITSTATUS() on exit_status directly. Do not attempt to scan or parse the
// error message string; it may be translated and/or change in future versions
// of GLib.
//
// The function takes the following parameters:
//
//    - exitStatus: exit code as returned from g_spawn_sync().
//
func SpawnCheckExitStatus(exitStatus int32) error {
	var _arg1 C.gint    // out
	var _cerr *C.GError // in

	_arg1 = C.gint(exitStatus)

	C.g_spawn_check_exit_status(_arg1, &_cerr)
	runtime.KeepAlive(exitStatus)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SpawnCommandLineAsync: simple version of g_spawn_async() that parses a
// command line with g_shell_parse_argv() and passes it to g_spawn_async(). Runs
// a command line in the background. Unlike g_spawn_async(), the
// G_SPAWN_SEARCH_PATH flag is enabled, other flags are not. Note that
// G_SPAWN_SEARCH_PATH can have security implications, so consider using
// g_spawn_async() directly if appropriate. Possible errors are those from
// g_shell_parse_argv() and g_spawn_async().
//
// The same concerns on Windows apply as for g_spawn_command_line_sync().
//
// The function takes the following parameters:
//
//    - commandLine: command line.
//
func SpawnCommandLineAsync(commandLine string) error {
	var _arg1 *C.gchar  // out
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(commandLine)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_async(_arg1, &_cerr)
	runtime.KeepAlive(commandLine)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SpawnCommandLineSync: simple version of g_spawn_sync() with little-used
// parameters removed, taking a command line instead of an argument vector. See
// g_spawn_sync() for full details. command_line will be parsed by
// g_shell_parse_argv(). Unlike g_spawn_sync(), the G_SPAWN_SEARCH_PATH flag is
// enabled. Note that G_SPAWN_SEARCH_PATH can have security implications, so
// consider using g_spawn_sync() directly if appropriate. Possible errors are
// those from g_spawn_sync() and those from g_shell_parse_argv().
//
// If exit_status is non-NULL, the platform-specific exit status of the child is
// stored there; see the documentation of g_spawn_check_exit_status() for how to
// use and interpret this.
//
// On Windows, please note the implications of g_shell_parse_argv() parsing
// command_line. Parsing is done according to Unix shell rules, not Windows
// command interpreter rules. Space is a separator, and backslashes are special.
// Thus you cannot simply pass a command_line containing canonical Windows
// paths, like "c:\\program files\\app\\app.exe", as the backslashes will be
// eaten, and the space will act as a separator. You need to enclose such paths
// with single quotes, like "'c:\\program files\\app\\app.exe'
// 'e:\\folder\\argument.txt'".
//
// The function takes the following parameters:
//
//    - commandLine: command line.
//
// The function returns the following values:
//
//    - standardOutput (optional): return location for child output.
//    - standardError (optional): return location for child errors.
//    - exitStatus (optional): return location for child exit status, as returned
//      by waitpid().
//
func SpawnCommandLineSync(commandLine string) (standardOutput []byte, standardError []byte, exitStatus int32, goerr error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // in
	var _arg3 *C.gchar  // in
	var _arg4 C.gint    // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(commandLine)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_sync(_arg1, &_arg2, &_arg3, &_arg4, &_cerr)
	runtime.KeepAlive(commandLine)

	var _standardOutput []byte // out
	var _standardError []byte  // out
	var _exitStatus int32      // out
	var _goerr error           // out

	if _arg2 != nil {
		defer C.free(unsafe.Pointer(_arg2))
		{
			var i int
			var z C.gchar
			for p := _arg2; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_arg2, i)
			_standardOutput = make([]byte, i)
			for i := range src {
				_standardOutput[i] = byte(src[i])
			}
		}
	}
	if _arg3 != nil {
		defer C.free(unsafe.Pointer(_arg3))
		{
			var i int
			var z C.gchar
			for p := _arg3; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_arg3, i)
			_standardError = make([]byte, i)
			for i := range src {
				_standardError[i] = byte(src[i])
			}
		}
	}
	_exitStatus = int32(_arg4)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _standardOutput, _standardError, _exitStatus, _goerr
}

// SpawnSync executes a child synchronously (waits for the child to exit before
// returning). All output from the child is stored in standard_output and
// standard_error, if those parameters are non-NULL. Note that you must set the
// G_SPAWN_STDOUT_TO_DEV_NULL and G_SPAWN_STDERR_TO_DEV_NULL flags when passing
// NULL for standard_output and standard_error.
//
// If exit_status is non-NULL, the platform-specific exit status of the child is
// stored there; see the documentation of g_spawn_check_exit_status() for how to
// use and interpret this. Note that it is invalid to pass
// G_SPAWN_DO_NOT_REAP_CHILD in flags, and on POSIX platforms, the same
// restrictions as for g_child_watch_source_new() apply.
//
// If an error occurs, no data is returned in standard_output, standard_error,
// or exit_status.
//
// This function calls g_spawn_async_with_pipes() internally; see that function
// for full details on the other parameters and details on how these functions
// work on Windows.
//
// The function takes the following parameters:
//
//    - workingDirectory (optional) child's current working directory, or NULL to
//      inherit parent's.
//    - argv: child's argument vector.
//    - envp (optional): child's environment, or NULL to inherit parent's.
//    - flags from Flags.
//    - childSetup (optional): function to run in the child just before exec().
//
// The function returns the following values:
//
//    - standardOutput (optional): return location for child output, or NULL.
//    - standardError (optional): return location for child error messages, or
//      NULL.
//    - exitStatus (optional): return location for child exit status, as returned
//      by waitpid(), or NULL.
//
func SpawnSync(workingDirectory string, argv, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc) (standardOutput []byte, standardError []byte, exitStatus int32, goerr error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 *C.gchar  // in
	var _arg8 *C.gchar  // in
	var _arg9 C.gint    // in
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.calloc(C.size_t((len(argv) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.calloc(C.size_t((len(envp) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}

	C.g_spawn_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_arg7, &_arg8, &_arg9, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)

	var _standardOutput []byte // out
	var _standardError []byte  // out
	var _exitStatus int32      // out
	var _goerr error           // out

	if _arg7 != nil {
		defer C.free(unsafe.Pointer(_arg7))
		{
			var i int
			var z C.gchar
			for p := _arg7; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_arg7, i)
			_standardOutput = make([]byte, i)
			for i := range src {
				_standardOutput[i] = byte(src[i])
			}
		}
	}
	if _arg8 != nil {
		defer C.free(unsafe.Pointer(_arg8))
		{
			var i int
			var z C.gchar
			for p := _arg8; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_arg8, i)
			_standardError = make([]byte, i)
			for i := range src {
				_standardError[i] = byte(src[i])
			}
		}
	}
	_exitStatus = int32(_arg9)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _standardOutput, _standardError, _exitStatus, _goerr
}
