// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// FileError values corresponding to @errno codes returned from file operations
// on UNIX. Unlike @errno codes, GFileError values are available on all systems,
// even Windows. The exact meaning of each code depends on what sort of file
// operation you were performing; the UNIX documentation gives more details. The
// following error code descriptions come from the GNU C Library manual, and are
// under the copyright of that manual.
//
// It's not very portable to make detailed assumptions about exactly which
// errors will be returned from a given operation. Some errors don't occur on
// some systems, etc., sometimes there are subtle differences in when a system
// will report a given error, etc.
type FileError int

const (
	// FileErrorExist: operation not permitted; only the owner of the file (or
	// other resource) or processes with special privileges can perform the
	// operation.
	FileErrorExist FileError = 0
	// FileErrorIsdir: file is a directory; you cannot open a directory for
	// writing, or create or remove hard links to it.
	FileErrorIsdir FileError = 1
	// FileErrorAcces: permission denied; the file permissions do not allow the
	// attempted operation.
	FileErrorAcces FileError = 2
	// FileErrorNametoolong: filename too long.
	FileErrorNametoolong FileError = 3
	// FileErrorNoent: no such file or directory. This is a "file doesn't exist"
	// error for ordinary files that are referenced in contexts where they are
	// expected to already exist.
	FileErrorNoent FileError = 4
	// FileErrorNotdir: a file that isn't a directory was specified when a
	// directory is required.
	FileErrorNotdir FileError = 5
	// FileErrorNxio: no such device or address. The system tried to use the
	// device represented by a file you specified, and it couldn't find the
	// device. This can mean that the device file was installed incorrectly, or
	// that the physical device is missing or not correctly attached to the
	// computer.
	FileErrorNxio FileError = 6
	// FileErrorNodev: the underlying file system of the specified file does not
	// support memory mapping.
	FileErrorNodev FileError = 7
	// FileErrorRofs: the directory containing the new link can't be modified
	// because it's on a read-only file system.
	FileErrorRofs FileError = 8
	// FileErrorTxtbsy: text file busy.
	FileErrorTxtbsy FileError = 9
	// FileErrorFault: you passed in a pointer to bad memory. (GLib won't
	// reliably return this, don't pass in pointers to bad memory.)
	FileErrorFault FileError = 10
	// FileErrorLoop: too many levels of symbolic links were encountered in
	// looking up a file name. This often indicates a cycle of symbolic links.
	FileErrorLoop FileError = 11
	// FileErrorNospc: no space left on device; write operation on a file failed
	// because the disk is full.
	FileErrorNospc FileError = 12
	// FileErrorNOMEM: no memory available. The system cannot allocate more
	// virtual memory because its capacity is full.
	FileErrorNOMEM FileError = 13
	// FileErrorMfile: the current process has too many files open and can't
	// open any more. Duplicate descriptors do count toward this limit.
	FileErrorMfile FileError = 14
	// FileErrorNfile: there are too many distinct file openings in the entire
	// system.
	FileErrorNfile FileError = 15
	// FileErrorBadf: bad file descriptor; for example, I/O on a descriptor that
	// has been closed or reading from a descriptor open only for writing (or
	// vice versa).
	FileErrorBadf FileError = 16
	// FileErrorInval: invalid argument. This is used to indicate various kinds
	// of problems with passing the wrong argument to a library function.
	FileErrorInval FileError = 17
	// FileErrorPipe: broken pipe; there is no process reading from the other
	// end of a pipe. Every library function that returns this error code also
	// generates a 'SIGPIPE' signal; this signal terminates the program if not
	// handled or blocked. Thus, your program will never actually see this code
	// unless it has handled or blocked 'SIGPIPE'.
	FileErrorPipe FileError = 18
	// FileErrorAgain: resource temporarily unavailable; the call might work if
	// you try again later.
	FileErrorAgain FileError = 19
	// FileErrorIntr: interrupted function call; an asynchronous signal occurred
	// and prevented completion of the call. When this happens, you should try
	// the call again.
	FileErrorIntr FileError = 20
	// FileErrorIO: input/output error; usually used for physical read or write
	// errors. i.e. the disk or other physical device hardware is returning
	// errors.
	FileErrorIO FileError = 21
	// FileErrorPerm: operation not permitted; only the owner of the file (or
	// other resource) or processes with special privileges can perform the
	// operation.
	FileErrorPerm FileError = 22
	// FileErrorNosys: function not implemented; this indicates that the system
	// is missing some functionality.
	FileErrorNosys FileError = 23
	// FileErrorFailed does not correspond to a UNIX error code; this is the
	// standard "failed for unspecified reason" error code present in all
	// #GError error code enumerations. Returned if no specific code applies.
	FileErrorFailed FileError = 24
)

// FileSetContentsFlags flags to pass to g_file_set_contents_full() to affect
// its safety and performance.
type FileSetContentsFlags int

const (
	// FileSetContentsFlagsNone: no guarantees about file consistency or
	// durability. The most dangerous setting, which is slightly faster than
	// other settings.
	FileSetContentsFlagsNone FileSetContentsFlags = 0
	// FileSetContentsFlagsConsistent: guarantee file consistency: after a
	// crash, either the old version of the file or the new version of the file
	// will be available, but not a mixture. On Unix systems this equates to an
	// `fsync()` on the file and use of an atomic `rename()` of the new version
	// of the file over the old.
	FileSetContentsFlagsConsistent FileSetContentsFlags = 1
	// FileSetContentsFlagsDurable: guarantee file durability: after a crash,
	// the new version of the file will be available. On Unix systems this
	// equates to an `fsync()` on the file (if G_FILE_SET_CONTENTS_CONSISTENT is
	// unset), or the effects of G_FILE_SET_CONTENTS_CONSISTENT plus an
	// `fsync()` on the directory containing the file after calling `rename()`.
	FileSetContentsFlagsDurable FileSetContentsFlags = 2
	// FileSetContentsFlagsOnlyExisting: only apply consistency and durability
	// guarantees if the file already exists. This may speed up file operations
	// if the file doesn’t currently exist, but may result in a corrupted
	// version of the new file if the system crashes while writing it.
	FileSetContentsFlagsOnlyExisting FileSetContentsFlags = 4
)

// FileTest: a test to perform on a file using g_file_test().
type FileTest int

const (
	// FileTestIsRegular: true if the file is a regular file (not a directory).
	// Note that this test will also return true if the tested file is a symlink
	// to a regular file.
	FileTestIsRegular FileTest = 1
	// FileTestIsSymlink: true if the file is a symlink.
	FileTestIsSymlink FileTest = 2
	// FileTestIsDir: true if the file is a directory.
	FileTestIsDir FileTest = 4
	// FileTestIsExecutable: true if the file is executable.
	FileTestIsExecutable FileTest = 8
	// FileTestExists: true if the file exists. It may or may not be a regular
	// file.
	FileTestExists FileTest = 16
)

// Basename gets the name of the file without any leading directory components.
// It returns a pointer into the given file name string.
func Basename(fileName string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_basename(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// BuildFilenamev behaves exactly like g_build_filename(), but takes the path
// elements as a string array, instead of varargs. This function is mainly meant
// for language bindings.
func BuildFilenamev(args []string) string {
	var _arg1 **C.gchar

	_arg1 = (**C.gchar)(C.malloc(C.ulong((len(args) + 1)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))

	{
		out := unsafe.Slice(_arg1, len(args))
		for i := range args {
			out[i] = (*C.gchar)(C.CString(args[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	var _cret *C.gchar // in

	_cret = C.g_build_filenamev(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// BuildPathv behaves exactly like g_build_path(), but takes the path elements
// as a string array, instead of varargs. This function is mainly meant for
// language bindings.
func BuildPathv(separator string, args []string) string {
	var _arg1 *C.gchar // out
	var _arg2 **C.gchar

	_arg1 = (*C.gchar)(C.CString(separator))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.gchar)(C.malloc(C.ulong((len(args) + 1)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))

	{
		out := unsafe.Slice(_arg2, len(args))
		for i := range args {
			out[i] = (*C.gchar)(C.CString(args[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	var _cret *C.gchar // in

	_cret = C.g_build_pathv(_arg1, _arg2)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// CanonicalizeFilename gets the canonical file name from @filename. All triple
// slashes are turned into single slashes, and all `..` and `.`s resolved
// against @relative_to.
//
// Symlinks are not followed, and the returned path is guaranteed to be
// absolute.
//
// If @filename is an absolute path, @relative_to is ignored. Otherwise,
// @relative_to will be prepended to @filename to make it absolute. @relative_to
// must be an absolute path, or nil. If @relative_to is nil, it'll fallback to
// g_get_current_dir().
//
// This function never fails, and will canonicalize file paths even if they
// don't exist.
//
// No file system I/O is done.
func CanonicalizeFilename(filename string, relativeTo string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(relativeTo))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret *C.gchar // in

	_cret = C.g_canonicalize_filename(_arg1, _arg2)

	var _ret string // out

	_ret = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _ret
}

// FileTest returns true if any of the tests in the bitfield @test are true. For
// example, `(G_FILE_TEST_EXISTS | G_FILE_TEST_IS_DIR)` will return true if the
// file exists; the check whether it's a directory doesn't matter since the
// existence test is true. With the current set of available tests, there's no
// point passing in more than one test at a time.
//
// Apart from G_FILE_TEST_IS_SYMLINK all tests follow symbolic links, so for a
// symbolic link to a regular file g_file_test() will return true for both
// G_FILE_TEST_IS_SYMLINK and G_FILE_TEST_IS_REGULAR.
//
// Note, that for a dangling symbolic link g_file_test() will return true for
// G_FILE_TEST_IS_SYMLINK and false for all other flags.
//
// You should never use g_file_test() to test whether it is safe to perform an
// operation, because there is always the possibility of the condition changing
// before you actually perform the operation. For example, you might think you
// could use G_FILE_TEST_IS_SYMLINK to know whether it is safe to write to a
// file without being tricked into writing into a different location. It doesn't
// work!
//
//    // DON'T DO THIS
//    if (!g_file_test (filename, G_FILE_TEST_IS_SYMLINK))
//      {
//        fd = g_open (filename, O_WRONLY);
//        // write to fd
//      }
//
// Another thing to note is that G_FILE_TEST_EXISTS and
// G_FILE_TEST_IS_EXECUTABLE are implemented using the access() system call.
// This usually doesn't matter, but if your program is setuid or setgid it means
// that these tests will give you the answer for the real user ID and group ID,
// rather than the effective user ID and group ID.
//
// On Windows, there are no symlinks, so testing for G_FILE_TEST_IS_SYMLINK will
// always return false. Testing for G_FILE_TEST_IS_EXECUTABLE will just check
// that the file exists and its name indicates that it is executable, checking
// for well-known extensions and those listed in the `PATHEXT` environment
// variable.
func FileTest(filename string, test FileTest) bool {
	var _arg1 *C.gchar    // out
	var _arg2 C.GFileTest // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GFileTest)(test)

	var _cret C.gboolean // in

	_cret = C.g_file_test(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GetCurrentDir gets the current directory.
//
// The returned string should be freed when no longer needed. The encoding of
// the returned string is system defined. On Windows, it is always UTF-8.
//
// Since GLib 2.40, this function will return the value of the "PWD" environment
// variable if it is set and it happens to be the same as the current directory.
// This can make a difference in the case that the current directory is the
// target of a symbolic link.
func GetCurrentDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_current_dir()

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// MkdirWithParents: create a directory if it doesn't already exist. Create
// intermediate parent directories as needed, too.
func MkdirWithParents(pathname string, mode int) int {
	var _arg1 *C.gchar // out
	var _arg2 C.gint   // out

	_arg1 = (*C.gchar)(C.CString(pathname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(mode)

	var _cret C.gint // in

	_cret = C.g_mkdir_with_parents(_arg1, _arg2)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// PathGetBasename gets the last component of the filename.
//
// If @file_name ends with a directory separator it gets the component before
// the last slash. If @file_name consists only of directory separators (and on
// Windows, possibly a drive letter), a single separator is returned. If
// @file_name is empty, it gets ".".
func PathGetBasename(fileName string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_path_get_basename(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// PathGetDirname gets the directory components of a file name. For example, the
// directory component of `/usr/bin/test` is `/usr/bin`. The directory component
// of `/` is `/`.
//
// If the file name has no directory components "." is returned. The returned
// string should be freed when no longer needed.
func PathGetDirname(fileName string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_path_get_dirname(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// PathIsAbsolute returns true if the given @file_name is an absolute file name.
// Note that this is a somewhat vague concept on Windows.
//
// On POSIX systems, an absolute file name is well-defined. It always starts
// from the single root directory. For example "/usr/local".
//
// On Windows, the concepts of current drive and drive-specific current
// directory introduce vagueness. This function interprets as an absolute file
// name one that either begins with a directory separator such as "\Users\tml"
// or begins with the root on a drive, for example "C:\Windows". The first case
// also includes UNC paths such as "\\\\myserver\docs\foo". In all cases, either
// slashes or backslashes are accepted.
//
// Note that a file name relative to the current drive root does not truly
// specify a file uniquely over time and across processes, as the current drive
// is a per-process value and can be changed.
//
// File names relative the current directory on some specific drive, such as
// "D:foo/bar", are not interpreted as absolute by this function, but they
// obviously are not relative to the normal current directory as returned by
// getcwd() or g_get_current_dir() either. Such paths should be avoided, or need
// to be handled using Windows-specific code.
func PathIsAbsolute(fileName string) bool {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_path_is_absolute(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PathSkipRoot returns a pointer into @file_name after the root component, i.e.
// after the "/" in UNIX or "C:\" under Windows. If @file_name is not an
// absolute path it returns nil.
func PathSkipRoot(fileName string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_path_skip_root(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}
