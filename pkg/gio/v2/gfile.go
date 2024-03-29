// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFile = coreglib.Type(girepository.MustFind("Gio", "File").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFile, F: marshalFile},
	})
}

// FileOverrider contains methods that are overridable.
type FileOverrider interface {
}

// File is a high level abstraction for manipulating files on a virtual file
// system. #GFiles are lightweight, immutable objects that do no I/O upon
// creation. It is necessary to understand that #GFile objects do not represent
// files, merely an identifier for a file. All file content I/O is implemented
// as streaming operations (see Stream and Stream).
//
// To construct a #GFile, you can use:
//
// - g_file_new_for_path() if you have a path.
//
// - g_file_new_for_uri() if you have a URI.
//
// - g_file_new_for_commandline_arg() for a command line argument.
//
// - g_file_new_tmp() to create a temporary file from a template.
//
// - g_file_parse_name() from a UTF-8 string gotten from
// g_file_get_parse_name().
//
// - g_file_new_build_filename() to create a file from path elements.
//
// One way to think of a #GFile is as an abstraction of a pathname. For normal
// files the system pathname is what is stored internally, but as #GFiles are
// extensible it could also be something else that corresponds to a pathname in
// a userspace implementation of a filesystem.
//
// #GFiles make up hierarchies of directories and files that correspond to the
// files on a filesystem. You can move through the file system with #GFile using
// g_file_get_parent() to get an identifier for the parent directory,
// g_file_get_child() to get a child within a directory,
// g_file_resolve_relative_path() to resolve a relative path between two
// #GFiles. There can be multiple hierarchies, so you may not end up at the same
// root if you repeatedly call g_file_get_parent() on two different files.
//
// All #GFiles have a basename (get with g_file_get_basename()). These names are
// byte strings that are used to identify the file on the filesystem (relative
// to its parent directory) and there is no guarantees that they have any
// particular charset encoding or even make any sense at all. If you want to use
// filenames in a user interface you should use the display name that you can
// get by requesting the G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME attribute with
// g_file_query_info(). This is guaranteed to be in UTF-8 and can be used in a
// user interface. But always store the real basename or the #GFile to use to
// actually access the file, because there is no way to go from a display name
// to the actual name.
//
// Using #GFile as an identifier has the same weaknesses as using a path in that
// there may be multiple aliases for the same file. For instance, hard or soft
// links may cause two different #GFiles to refer to the same file. Other
// possible causes for aliases are: case insensitive filesystems, short and long
// names on FAT/NTFS, or bind mounts in Linux. If you want to check if two
// #GFiles point to the same file you can query for the G_FILE_ATTRIBUTE_ID_FILE
// attribute. Note that #GFile does some trivial canonicalization of pathnames
// passed in, so that trivial differences in the path string used at creation
// (duplicated slashes, slash at end of path, "." or ".." path segments, etc)
// does not create different #GFiles.
//
// Many #GFile operations have both synchronous and asynchronous versions to
// suit your application. Asynchronous versions of synchronous functions simply
// have _async() appended to their function names. The asynchronous I/O
// functions call a ReadyCallback which is then used to finalize the operation,
// producing a GAsyncResult which is then passed to the function's matching
// _finish() operation.
//
// It is highly recommended to use asynchronous calls when running within a
// shared main loop, such as in the main thread of an application. This avoids
// I/O operations blocking other sources on the main loop from being dispatched.
// Synchronous I/O operations should be performed from worker threads. See the
// [introduction to asynchronous programming section][async-programming] for
// more.
//
// Some #GFile operations almost always take a noticeable amount of time, and so
// do not have synchronous analogs. Notable cases include:
//
// - g_file_mount_mountable() to mount a mountable file.
//
// - g_file_unmount_mountable_with_operation() to unmount a mountable file.
//
// - g_file_eject_mountable_with_operation() to eject a mountable file.
//
//
// Entity Tags
//
// One notable feature of #GFiles are entity tags, or "etags" for short. Entity
// tags are somewhat like a more abstract version of the traditional mtime, and
// can be used to quickly determine if the file has been modified from the
// version on the file system. See the HTTP 1.1 specification
// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html) for HTTP Etag
// headers, which are a very similar concept.
//
// File wraps an interface. This means the user can get the
// underlying type by calling Cast().
type File struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*File)(nil)
)

// Filer describes File's interface methods.
type Filer interface {
	coreglib.Objector

	baseFile() *File
}

var _ Filer = (*File)(nil)

func ifaceInitFiler(gifacePtr, data C.gpointer) {
}

func wrapFile(obj *coreglib.Object) *File {
	return &File{
		Object: obj,
	}
}

func marshalFile(p uintptr) (interface{}, error) {
	return wrapFile(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *File) baseFile() *File {
	return v
}

// BaseFile returns the underlying base object.
func BaseFile(obj Filer) *File {
	return obj.baseFile()
}

// FileIface: interface for writing VFS file handles.
//
// An instance of this type is always passed by reference.
type FileIface struct {
	*fileIface
}

// fileIface is the struct that's finalized.
type fileIface struct {
	native unsafe.Pointer
}

var GIRInfoFileIface = girepository.MustFind("Gio", "FileIface")

// SupportsThreadContexts: boolean that indicates whether the #GFile
// implementation supports thread-default contexts. Since 2.22.
func (f *FileIface) SupportsThreadContexts() bool {
	offset := GIRInfoFileIface.StructFieldOffset("supports_thread_contexts")
	valptr := (*bool)(unsafe.Add(f.native, offset))
	var _v bool // out
	if *valptr != 0 {
		_v = true
	}
	return _v
}
