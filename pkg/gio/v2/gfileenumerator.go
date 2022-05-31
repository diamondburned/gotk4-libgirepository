// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GFileInfo* _gotk4_gio2_FileEnumeratorClass_next_file(void*, void*, GError**);
// extern GList* _gotk4_gio2_FileEnumeratorClass_next_files_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_FileEnumeratorClass_close_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_FileEnumeratorClass_close_fn(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// glib.Type values for gfileenumerator.go.
var GTypeFileEnumerator = coreglib.Type(C.g_file_enumerator_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFileEnumerator, F: marshalFileEnumerator},
	})
}

// FileEnumeratorOverrider contains methods that are overridable.
type FileEnumeratorOverrider interface {
	// CloseFinish finishes closing a file enumerator, started from
	// g_file_enumerator_close_async().
	//
	// If the file enumerator was already closed when
	// g_file_enumerator_close_async() was called, then this function will
	// report G_IO_ERROR_CLOSED in error, and return FALSE. If the file
	// enumerator had pending operation when the close operation was started,
	// then this function will report G_IO_ERROR_PENDING, and return FALSE. If
	// cancellable was not NULL, then the operation may have been cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and FALSE will
	// be returned.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	CloseFinish(result AsyncResulter) error
	// The function takes the following parameters:
	//
	CloseFn(ctx context.Context) error
	// NextFile returns information for the next file in the enumerated object.
	// Will block until the information is available. The Info returned from
	// this function will contain attributes that match the attribute string
	// that was passed when the Enumerator was created.
	//
	// See the documentation of Enumerator for information about the order of
	// returned files.
	//
	// On error, returns NULL and sets error to the error. If the enumerator is
	// at the end, NULL will be returned and error will be unset.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//
	// The function returns the following values:
	//
	//    - fileInfo (optional) or NULL on error or end of enumerator. Free the
	//      returned object with g_object_unref() when no longer needed.
	//
	NextFile(ctx context.Context) (*FileInfo, error)
	// NextFilesFinish finishes the asynchronous operation started with
	// g_file_enumerator_next_files_async().
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - list of Infos. You must free the list with g_list_free() and unref
	//      the infos with g_object_unref() when you're done with them.
	//
	NextFilesFinish(result AsyncResulter) ([]*FileInfo, error)
}

// FileEnumerator allows you to operate on a set of #GFiles, returning a Info
// structure for each file enumerated (e.g. g_file_enumerate_children() will
// return a Enumerator for each of the children within a directory).
//
// To get the next file's information from a Enumerator, use
// g_file_enumerator_next_file() or its asynchronous version,
// g_file_enumerator_next_files_async(). Note that the asynchronous version will
// return a list of Infos, whereas the synchronous will only return the next
// file in the enumerator.
//
// The ordering of returned files is unspecified for non-Unix platforms; for
// more information, see g_dir_read_name(). On Unix, when operating on local
// files, returned files will be sorted by inode number. Effectively you can
// assume that the ordering of returned files will be stable between successive
// calls (and applications) assuming the directory is unchanged.
//
// If your application needs a specific ordering, such as by name or
// modification time, you will have to implement that in your application code.
//
// To close a Enumerator, use g_file_enumerator_close(), or its asynchronous
// version, g_file_enumerator_close_async(). Once a Enumerator is closed, no
// further actions may be performed on it, and it should be freed with
// g_object_unref().
type FileEnumerator struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileEnumerator)(nil)
)

func classInitFileEnumeratorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GFileEnumeratorClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GFileEnumeratorClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		CloseFinish(result AsyncResulter) error
	}); ok {
		pclass.close_finish = (*[0]byte)(C._gotk4_gio2_FileEnumeratorClass_close_finish)
	}

	if _, ok := goval.(interface {
		CloseFn(ctx context.Context) error
	}); ok {
		pclass.close_fn = (*[0]byte)(C._gotk4_gio2_FileEnumeratorClass_close_fn)
	}

	if _, ok := goval.(interface {
		NextFile(ctx context.Context) (*FileInfo, error)
	}); ok {
		pclass.next_file = (*[0]byte)(C._gotk4_gio2_FileEnumeratorClass_next_file)
	}

	if _, ok := goval.(interface {
		NextFilesFinish(result AsyncResulter) ([]*FileInfo, error)
	}); ok {
		pclass.next_files_finish = (*[0]byte)(C._gotk4_gio2_FileEnumeratorClass_next_files_finish)
	}
}

//export _gotk4_gio2_FileEnumeratorClass_close_finish
func _gotk4_gio2_FileEnumeratorClass_close_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.CloseFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileEnumeratorClass_close_fn
func _gotk4_gio2_FileEnumeratorClass_close_fn(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFn(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.CloseFn(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileEnumeratorClass_next_file
func _gotk4_gio2_FileEnumeratorClass_next_file(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret *C.GFileInfo) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		NextFile(ctx context.Context) (*FileInfo, error)
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	fileInfo, _goerr := iface.NextFile(_cancellable)

	if fileInfo != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fileInfo).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(fileInfo).Native()))
	}
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileEnumeratorClass_next_files_finish
func _gotk4_gio2_FileEnumeratorClass_next_files_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret *C.GList) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		NextFilesFinish(result AsyncResulter) ([]*FileInfo, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	list, _goerr := iface.NextFilesFinish(_result)

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.void // out
		dst = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapFileEnumerator(obj *coreglib.Object) *FileEnumerator {
	return &FileEnumerator{
		Object: obj,
	}
}

func marshalFileEnumerator(p uintptr) (interface{}, error) {
	return wrapFileEnumerator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Close releases all resources used by this enumerator, making the enumerator
// return G_IO_ERROR_CLOSED on all calls.
//
// This will be automatically called when the last reference is dropped, but you
// might want to call this function to make sure resources are released as early
// as possible.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
func (enumerator *FileEnumerator) Close(ctx context.Context) error {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("close", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CloseAsync: asynchronously closes the file enumerator.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned in
// g_file_enumerator_close_finish().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (enumerator *FileEnumerator) CloseAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument
	var _arg0 *C.void    // out
	var _arg2 *C.void    // out
	var _arg1 C.int      // out
	var _arg3 C.gpointer // out
	var _arg4 C.gpointer

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(ioPriority)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.gpointer)(unsafe.Pointer(&_args[3])) = _arg3

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("close_async", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// CloseFinish finishes closing a file enumerator, started from
// g_file_enumerator_close_async().
//
// If the file enumerator was already closed when
// g_file_enumerator_close_async() was called, then this function will report
// G_IO_ERROR_CLOSED in error, and return FALSE. If the file enumerator had
// pending operation when the close operation was started, then this function
// will report G_IO_ERROR_PENDING, and return FALSE. If cancellable was not
// NULL, then the operation may have been cancelled by triggering the
// cancellable object from another thread. If the operation was cancelled, the
// error G_IO_ERROR_CANCELLED will be set, and FALSE will be returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (enumerator *FileEnumerator) CloseFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("close_finish", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Child: return a new #GFile which refers to the file named by info in the
// source directory of enumerator. This function is primarily intended to be
// used inside loops with g_file_enumerator_next_file().
//
// This is a convenience method that's equivalent to:
//
//    gchar *name = g_file_info_get_name (info);
//    GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
//                                     name);.
//
// The function takes the following parameters:
//
//    - info gotten from g_file_enumerator_next_file() or the async equivalents.
//
// The function returns the following values:
//
//    - file for the Info passed it.
//
func (enumerator *FileEnumerator) Child(info *FileInfo) *File {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(info).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(info)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// Container: get the #GFile container which is being enumerated.
//
// The function returns the following values:
//
//    - file which is being enumerated.
//
func (enumerator *FileEnumerator) Container() *File {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("get_container", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)

	var _file *File // out

	_file = wrapFile(coreglib.Take(unsafe.Pointer(_cret)))

	return _file
}

// HasPending checks if the file enumerator has pending operations.
//
// The function returns the following values:
//
//    - ok: TRUE if the enumerator has pending operations.
//
func (enumerator *FileEnumerator) HasPending() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("has_pending", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if the file enumerator has been closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the enumerator is closed.
//
func (enumerator *FileEnumerator) IsClosed() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("is_closed", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Iterate: this is a version of g_file_enumerator_next_file() that's easier to
// use correctly from C programs. With g_file_enumerator_next_file(), the
// gboolean return value signifies "end of iteration or error", which requires
// allocation of a temporary #GError.
//
// In contrast, with this function, a FALSE return from
// g_file_enumerator_iterate() *always* means "error". End of iteration is
// signaled by out_info or out_child being NULL.
//
// Another crucial difference is that the references for out_info and out_child
// are owned by direnum (they are cached as hidden properties). You must not
// unref them in your own code. This makes memory management significantly
// easier for C code in combination with loops.
//
// Finally, this function optionally allows retrieving a #GFile as well.
//
// You must specify at least one of out_info or out_child.
//
// The code pattern for correctly using g_file_enumerator_iterate() from C is:
//
//    direnum = g_file_enumerate_children (file, ...);
//    while (TRUE)
//      {
//        GFileInfo *info;
//        if (!g_file_enumerator_iterate (direnum, &info, NULL, cancellable, error))
//          goto out;
//        if (!info)
//          break;
//        ... do stuff with "info"; do not unref it! ...
//      }
//
//    out:
//      g_object_unref (direnum); // Note: frees the last info.
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//
// The function returns the following values:
//
//    - outInfo (optional): output location for the next Info, or NULL.
//    - outChild (optional): output location for the next #GFile, or NULL.
//
func (direnum *FileEnumerator) Iterate(ctx context.Context) (*FileInfo, *File, error) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _out0 *C.void // in
	var _out1 *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(direnum).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("iterate", _args[:], _outs[:])

	runtime.KeepAlive(direnum)
	runtime.KeepAlive(ctx)

	var _outInfo *FileInfo // out
	var _outChild *File    // out
	var _goerr error       // out
	_out1 = *(**C.void)(unsafe.Pointer(&_outs[1]))
	_out1 = *(**C.void)(unsafe.Pointer(&_outs[1]))

	if _out0 != nil {
		_outInfo = wrapFileInfo(coreglib.Take(unsafe.Pointer(_out0)))
	}
	if _out1 != nil {
		_outChild = wrapFile(coreglib.Take(unsafe.Pointer(_out1)))
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outInfo, _outChild, _goerr
}

// NextFile returns information for the next file in the enumerated object. Will
// block until the information is available. The Info returned from this
// function will contain attributes that match the attribute string that was
// passed when the Enumerator was created.
//
// See the documentation of Enumerator for information about the order of
// returned files.
//
// On error, returns NULL and sets error to the error. If the enumerator is at
// the end, NULL will be returned and error will be unset.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - fileInfo (optional) or NULL on error or end of enumerator. Free the
//      returned object with g_object_unref() when no longer needed.
//
func (enumerator *FileEnumerator) NextFile(ctx context.Context) (*FileInfo, error) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("next_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	if _cret != nil {
		_fileInfo = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}

// NextFilesAsync: request information for a number of files from the enumerator
// asynchronously. When all i/o for the operation is finished the callback will
// be called with the requested information.
//
// See the documentation of Enumerator for information about the order of
// returned files.
//
// The callback can be called with less than num_files files in case of error or
// at the end of the enumerator. In case of a partial error the callback will be
// called with any succeeding items and no error, and on the next request the
// error will be reported. If a request is cancelled the callback will be called
// with G_IO_ERROR_CANCELLED.
//
// During an async request no other sync and async calls are allowed, and will
// result in G_IO_ERROR_PENDING errors.
//
// Any outstanding i/o request with higher priority (lower numerical value) will
// be executed before an outstanding request with lower priority. Default
// priority is G_PRIORITY_DEFAULT.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - numFiles: number of file info objects to request.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (enumerator *FileEnumerator) NextFilesAsync(ctx context.Context, numFiles, ioPriority int32, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument
	var _arg0 *C.void    // out
	var _arg3 *C.void    // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out
	var _arg4 C.gpointer // out
	var _arg5 C.gpointer

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(numFiles)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.int)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.gpointer)(unsafe.Pointer(&_args[4])) = _arg4

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("next_files_async", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(numFiles)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// NextFilesFinish finishes the asynchronous operation started with
// g_file_enumerator_next_files_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - list of Infos. You must free the list with g_list_free() and unref the
//      infos with g_object_unref() when you're done with them.
//
func (enumerator *FileEnumerator) NextFilesFinish(result AsyncResulter) ([]*FileInfo, error) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("next_files_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(result)

	var _list []*FileInfo // out
	var _goerr error      // out

	_list = make([]*FileInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *FileInfo // out
		dst = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// SetPending sets the file enumerator as having pending operations.
//
// The function takes the following parameters:
//
//    - pending: boolean value.
//
func (enumerator *FileEnumerator) SetPending(pending bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	if pending {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "FileEnumerator").InvokeMethod("set_pending", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(pending)
}
