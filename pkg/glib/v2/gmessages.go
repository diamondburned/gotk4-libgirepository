// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// LogWriterOutput: return values from WriterFuncs to indicate whether the given
// log entry was successfully handled by the writer, or whether there was an
// error in handling it (and hence a fallback writer should be used).
//
// If a WriterFunc ignores a log entry, it should return G_LOG_WRITER_HANDLED.
type LogWriterOutput int

const (
	// LogWriterOutputHandled: log writer has handled the log entry.
	LogWriterOutputHandled LogWriterOutput = 1
	// LogWriterOutputUnhandled: log writer could not handle the log entry.
	LogWriterOutputUnhandled LogWriterOutput = 0
)

// LogLevelFlags flags specifying the level of log messages.
//
// It is possible to change how GLib treats messages of the various levels using
// g_log_set_handler() and g_log_set_fatal_mask().
type LogLevelFlags int

const (
	// LogLevelFlagsFlagRecursion: internal flag
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	// LogLevelFlagsFlagFatal: internal flag
	LogLevelFlagsFlagFatal LogLevelFlags = 2
	// LogLevelFlagsLevelError: log level for errors, see g_error(). This level
	// is also used for messages produced by g_assert().
	LogLevelFlagsLevelError LogLevelFlags = 4
	// LogLevelFlagsLevelCritical: log level for critical warning messages, see
	// g_critical(). This level is also used for messages produced by
	// g_return_if_fail() and g_return_val_if_fail().
	LogLevelFlagsLevelCritical LogLevelFlags = 8
	// LogLevelFlagsLevelWarning: log level for warnings, see g_warning()
	LogLevelFlagsLevelWarning LogLevelFlags = 16
	// LogLevelFlagsLevelMessage: log level for messages, see g_message()
	LogLevelFlagsLevelMessage LogLevelFlags = 32
	// LogLevelFlagsLevelInfo: log level for informational messages, see
	// g_info()
	LogLevelFlagsLevelInfo LogLevelFlags = 64
	// LogLevelFlagsLevelDebug: log level for debug messages, see g_debug()
	LogLevelFlagsLevelDebug LogLevelFlags = 128
	// LogLevelFlagsLevelMask: a mask including all log levels
	LogLevelFlagsLevelMask LogLevelFlags = -4
)

// LogFunc specifies the prototype of log handler functions.
//
// The default log handler, g_log_default_handler(), automatically appends a
// new-line character to @message when printing it. It is advised that any
// custom log handler functions behave similarly, so that logging calls in user
// code do not need modifying to add a new-line character to the message if the
// log handler is changed.
//
// This is not used if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
type LogFunc func()

//export gotk4_LogFunc
func gotk4_LogFunc(arg0 *C.gchar, arg1 C.GLogLevelFlags, arg2 *C.gchar, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(LogFunc)
	fn()
}

// LogWriterFunc: writer function for log entries. A log entry is a collection
// of one or more Fields, using the standard [field names from journal
// specification](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html).
// See g_log_structured() for more information.
//
// Writer functions must ignore fields which they do not recognise, unless they
// can write arbitrary binary output, as field values may be arbitrary binary.
//
// @log_level is guaranteed to be included in @fields as the `PRIORITY` field,
// but is provided separately for convenience of deciding whether or where to
// output the log entry.
//
// Writer functions should return G_LOG_WRITER_HANDLED if they handled the log
// message successfully or if they deliberately ignored it. If there was an
// error handling the message (for example, if the writer function is meant to
// send messages to a remote logging server and there is a network error), it
// should return G_LOG_WRITER_UNHANDLED. This allows writer functions to be
// chained and fall back to simpler handlers in case of failure.
type LogWriterFunc func() (logWriterOutput LogWriterOutput)

//export gotk4_LogWriterFunc
func gotk4_LogWriterFunc(arg0 C.GLogLevelFlags, arg1 *C.GLogField, arg2 C.gsize, arg3 C.gpointer) C.GLogWriterOutput {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(LogWriterFunc)
	fn(logWriterOutput)

	cret = (C.GLogWriterOutput)(logWriterOutput)
}

func AssertWarning(logDomain string, file string, line int, prettyFunction string, expression string) {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 C.int
	var arg4 *C.char
	var arg5 *C.char

	arg1 = (*C.char)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(file))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.int(line)
	arg4 = (*C.char)(C.CString(prettyFunction))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.char)(C.CString(expression))
	defer C.free(unsafe.Pointer(arg5))

	C.g_assert_warning(arg1, arg2, arg3, arg4, arg5)
}

// LogDefaultHandler: the default log handler set up by GLib;
// g_log_set_default_handler() allows to install an alternate default log
// handler. This is used if no log handler has been set for the particular log
// domain and log level combination. It outputs the message to stderr or stdout
// and if the log level is fatal it calls G_BREAKPOINT(). It automatically
// prints a new-line character after the message, so one does not need to be
// manually included in @message.
//
// The behavior of this log handler can be influenced by a number of environment
// variables:
//
// - `G_MESSAGES_PREFIXED`: A :-separated list of log levels for which messages
// should be prefixed by the program name and PID of the application.
//
// - `G_MESSAGES_DEBUG`: A space-separated list of log domains for which debug
// and informational messages are printed. By default these messages are not
// printed.
//
// stderr is used for levels G_LOG_LEVEL_ERROR, G_LOG_LEVEL_CRITICAL,
// G_LOG_LEVEL_WARNING and G_LOG_LEVEL_MESSAGE. stdout is used for the rest,
// unless stderr was requested by g_log_writer_default_set_use_stderr().
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogDefaultHandler(logDomain string, logLevel LogLevelFlags, message string, unusedData interface{}) {
	var arg1 *C.gchar
	var arg2 C.GLogLevelFlags
	var arg3 *C.gchar
	var arg4 C.gpointer

	arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GLogLevelFlags)(logLevel)
	arg3 = (*C.gchar)(C.CString(message))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gpointer(unusedData)

	C.g_log_default_handler(arg1, arg2, arg3, arg4)
}

// LogRemoveHandler removes the log handler.
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogRemoveHandler(logDomain string, handlerID uint) {
	var arg1 *C.gchar
	var arg2 C.guint

	arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.guint(handlerID)

	C.g_log_remove_handler(arg1, arg2)
}

// LogSetAlwaysFatal sets the message levels which are always fatal, in any log
// domain. When a message with any of these levels is logged the program
// terminates. You can only set the levels defined by GLib to be fatal.
// G_LOG_LEVEL_ERROR is always fatal.
//
// You can also make some message levels fatal at runtime by setting the
// `G_DEBUG` environment variable (see Running GLib Applications
// (glib-running.html)).
//
// Libraries should not call this function, as it affects all messages logged by
// a process, including those from other libraries.
//
// Structured log messages (using g_log_structured() and
// g_log_structured_array()) are fatal only if the default log writer is used;
// otherwise it is up to the writer function to determine which log messages are
// fatal. See [Using Structured Logging][using-structured-logging].
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	var arg1 C.GLogLevelFlags

	arg1 = (C.GLogLevelFlags)(fatalMask)

	var cret C.GLogLevelFlags
	var goret LogLevelFlags

	cret = C.g_log_set_always_fatal(arg1)

	goret = LogLevelFlags(cret)

	return goret
}

// LogSetFatalMask sets the log levels which are fatal in the given domain.
// G_LOG_LEVEL_ERROR is always fatal.
//
// This has no effect on structured log messages (using g_log_structured() or
// g_log_structured_array()). To change the fatal behaviour for specific log
// messages, programs must install a custom log writer function using
// g_log_set_writer_func(). See [Using Structured
// Logging][using-structured-logging].
//
// This function is mostly intended to be used with G_LOG_LEVEL_CRITICAL. You
// should typically not set G_LOG_LEVEL_WARNING, G_LOG_LEVEL_MESSAGE,
// G_LOG_LEVEL_INFO or G_LOG_LEVEL_DEBUG as fatal except inside of test
// programs.
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	var arg1 *C.gchar
	var arg2 C.GLogLevelFlags

	arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GLogLevelFlags)(fatalMask)

	var cret C.GLogLevelFlags
	var goret LogLevelFlags

	cret = C.g_log_set_fatal_mask(arg1, arg2)

	goret = LogLevelFlags(cret)

	return goret
}

// LogSetHandlerFull: like g_log_set_handler(), but takes a destroy notify for
// the @user_data.
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogSetHandlerFull() uint {
	var cret C.guint
	var goret uint

	cret = C.g_log_set_handler_full(arg1, arg2, arg3, arg4, arg5)

	goret = uint(cret)

	return goret
}

// LogSetWriterFunc: set a writer function which will be called to format and
// write out each log message. Each program should set a writer function, or the
// default writer (g_log_writer_default()) will be used.
//
// Libraries **must not** call this function — only programs are allowed to
// install a writer function, as there must be a single, central point where log
// messages are formatted and outputted.
//
// There can only be one writer function. It is an error to set more than one.
func LogSetWriterFunc() {
	C.g_log_set_writer_func(arg1, arg2, arg3)
}

// LogStructuredArray: log a message with structured data. The message will be
// passed through to the log writer set by the application using
// g_log_set_writer_func(). If the message is fatal (i.e. its log level is
// G_LOG_LEVEL_ERROR), the program will be aborted at the end of this function.
//
// See g_log_structured() for more documentation.
//
// This assumes that @log_level is already present in @fields (typically as the
// `PRIORITY` field).
func LogStructuredArray() {
	C.g_log_structured_array(arg1, arg2, arg3)
}

// LogVariant: log a message with structured data, accepting the data within a
// #GVariant. This version is especially useful for use in other languages, via
// introspection.
//
// The only mandatory item in the @fields dictionary is the "MESSAGE" which must
// contain the text shown to the user.
//
// The values in the @fields dictionary are likely to be of type String
// (VARIANT_TYPE_STRING). Array of bytes (VARIANT_TYPE_BYTESTRING) is also
// supported. In this case the message is handled as binary and will be
// forwarded to the log writer as such. The size of the array should not be
// higher than G_MAXSSIZE. Otherwise it will be truncated to this size. For
// other types g_variant_print() will be used to convert the value into a
// string.
//
// For more details on its usage and about the parameters, see
// g_log_structured().
func LogVariant(logDomain string, logLevel LogLevelFlags, fields *Variant) {
	var arg1 *C.gchar
	var arg2 C.GLogLevelFlags
	var arg3 *C.GVariant

	arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GLogLevelFlags)(logLevel)
	arg3 = (*C.GVariant)(unsafe.Pointer(fields.Native()))

	C.g_log_variant(arg1, arg2, arg3)
}

// LogWriterDefaultSetUseStderr: configure whether the built-in log functions
// (g_log_default_handler() for the old-style API, and both
// g_log_writer_default() and g_log_writer_standard_streams() for the structured
// API) will output all log messages to `stderr`.
//
// By default, log messages of levels G_LOG_LEVEL_INFO and G_LOG_LEVEL_DEBUG are
// sent to `stdout`, and other log messages are sent to `stderr`. This is
// problematic for applications that intend to reserve `stdout` for structured
// output such as JSON or XML.
//
// This function sets global state. It is not thread-aware, and should be called
// at the very start of a program, before creating any other threads or creating
// objects that could create worker threads of their own.
func LogWriterDefaultSetUseStderr(useStderr bool) {
	var arg1 C.gboolean

	if useStderr {
		arg1 = C.gboolean(1)
	}

	C.g_log_writer_default_set_use_stderr(arg1)
}

// LogWriterDefaultWouldDrop: check whether g_log_writer_default() and
// g_log_default_handler() would ignore a message with the given domain and
// level.
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or `all`) is listed in the space-separated
// `G_MESSAGES_DEBUG` environment variable.
//
// This can be used when implementing log writers with the same filtering
// behaviour as the default, but a different destination or output format:
//
//      if (!g_log_writer_default_would_drop (G_LOG_LEVEL_DEBUG, G_LOG_DOMAIN))
//        {
//          gchar *result = expensive_computation (my_object);
//
//          g_debug ("my_object result: s", result);
//          g_free (result);
//        }
func LogWriterDefaultWouldDrop(logLevel LogLevelFlags, logDomain string) bool {
	var arg1 C.GLogLevelFlags
	var arg2 *C.char

	arg1 = (C.GLogLevelFlags)(logLevel)
	arg2 = (*C.char)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean
	var goret bool

	cret = C.g_log_writer_default_would_drop(arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// LogWriterIsJournald: check whether the given @output_fd file descriptor is a
// connection to the systemd journal, or something else (like a log file or
// `stdout` or `stderr`).
//
// Invalid file descriptors are accepted and return false, which allows for the
// following construct without needing any additional error handling:
//
//    is_journald = g_log_writer_is_journald (fileno (stderr));
func LogWriterIsJournald(outputFd int) bool {
	var arg1 C.gint

	arg1 = C.gint(outputFd)

	var cret C.gboolean
	var goret bool

	cret = C.g_log_writer_is_journald(arg1)

	if cret {
		goret = true
	}

	return goret
}

// LogWriterSupportsColor: check whether the given @output_fd file descriptor
// supports ANSI color escape sequences. If so, they can safely be used when
// formatting log messages.
func LogWriterSupportsColor(outputFd int) bool {
	var arg1 C.gint

	arg1 = C.gint(outputFd)

	var cret C.gboolean
	var goret bool

	cret = C.g_log_writer_supports_color(arg1)

	if cret {
		goret = true
	}

	return goret
}

// ReturnIfFailWarning: internal function used to print messages from the public
// g_return_if_fail() and g_return_val_if_fail() macros.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 *C.char

	arg1 = (*C.char)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(prettyFunction))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.char)(C.CString(expression))
	defer C.free(unsafe.Pointer(arg3))

	C.g_return_if_fail_warning(arg1, arg2, arg3)
}

// WarnMessage: internal function used to print messages from the public
// g_warn_if_reached() and g_warn_if_fail() macros.
func WarnMessage(domain string, file string, line int, fn string, warnexpr string) {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 C.int
	var arg4 *C.char
	var arg5 *C.char

	arg1 = (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(file))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.int(line)
	arg4 = (*C.char)(C.CString(fn))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.char)(C.CString(warnexpr))
	defer C.free(unsafe.Pointer(arg5))

	C.g_warn_message(arg1, arg2, arg3, arg4, arg5)
}

// LogField: structure representing a single field in a structured log entry.
// See g_log_structured() for details.
//
// Log fields may contain arbitrary values, including binary with embedded nul
// bytes. If the field contains a string, the string must be UTF-8 encoded and
// have a trailing nul byte. Otherwise, @length must be set to a non-negative
// value.
type LogField struct {
	native C.GLogField
}

// WrapLogField wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLogField(ptr unsafe.Pointer) *LogField {
	if ptr == nil {
		return nil
	}

	return (*LogField)(ptr)
}

func marshalLogField(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapLogField(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *LogField) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// Key gets the field inside the struct.
func (l *LogField) Key() string {
	var v string
	v = C.GoString(l.native.key)
	return v
}

// Value gets the field inside the struct.
func (l *LogField) Value() interface{} {
	var v interface{}
	v = interface{}(l.native.value)
	return v
}

// Length gets the field inside the struct.
func (l *LogField) Length() int {
	var v int
	v = int(l.native.length)
	return v
}
