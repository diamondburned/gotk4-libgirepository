// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gio2_SocketAddressClass_to_native(void*, gpointer, gsize, GError**);
// extern gssize _gotk4_gio2_SocketAddressClass_get_native_size(void*);
import "C"

// glib.Type values for gsocketaddress.go.
var GTypeSocketAddress = coreglib.Type(C.g_socket_address_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSocketAddress, F: marshalSocketAddress},
	})
}

// SocketAddressOverrider contains methods that are overridable.
type SocketAddressOverrider interface {
	// NativeSize gets the size of address's native struct sockaddr. You can use
	// this to allocate memory to pass to g_socket_address_to_native().
	//
	// The function returns the following values:
	//
	//    - gssize: size of the native struct sockaddr that address represents.
	//
	NativeSize() int
	// ToNative converts a Address to a native struct sockaddr, which can be
	// passed to low-level functions like connect() or bind().
	//
	// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
	// returned. If the address type is not known on the system then a
	// G_IO_ERROR_NOT_SUPPORTED error is returned.
	//
	// The function takes the following parameters:
	//
	//    - dest (optional): pointer to a memory location that will contain the
	//      native struct sockaddr.
	//    - destlen: size of dest. Must be at least as large as
	//      g_socket_address_get_native_size().
	//
	ToNative(dest unsafe.Pointer, destlen uint) error
}

// SocketAddress is the equivalent of struct sockaddr in the BSD sockets API.
// This is an abstract class; use SocketAddress for internet sockets, or
// SocketAddress for UNIX domain sockets.
type SocketAddress struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SocketConnectable
}

var (
	_ coreglib.Objector = (*SocketAddress)(nil)
)

// SocketAddresser describes types inherited from class SocketAddress.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type SocketAddresser interface {
	coreglib.Objector
	baseSocketAddress() *SocketAddress
}

var _ SocketAddresser = (*SocketAddress)(nil)

func classInitSocketAddresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GSocketAddressClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GSocketAddressClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ NativeSize() int }); ok {
		pclass.get_native_size = (*[0]byte)(C._gotk4_gio2_SocketAddressClass_get_native_size)
	}

	if _, ok := goval.(interface {
		ToNative(dest unsafe.Pointer, destlen uint) error
	}); ok {
		pclass.to_native = (*[0]byte)(C._gotk4_gio2_SocketAddressClass_to_native)
	}
}

//export _gotk4_gio2_SocketAddressClass_get_native_size
func _gotk4_gio2_SocketAddressClass_get_native_size(arg0 *C.void) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NativeSize() int })

	gssize := iface.NativeSize()

	cret = C.gssize(gssize)

	return cret
}

//export _gotk4_gio2_SocketAddressClass_to_native
func _gotk4_gio2_SocketAddressClass_to_native(arg0 *C.void, arg1 C.gpointer, arg2 C.gsize, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ToNative(dest unsafe.Pointer, destlen uint) error
	})

	var _dest unsafe.Pointer // out
	var _destlen uint        // out

	_dest = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_destlen = uint(arg2)

	_goerr := iface.ToNative(_dest, _destlen)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapSocketAddress(obj *coreglib.Object) *SocketAddress {
	return &SocketAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalSocketAddress(p uintptr) (interface{}, error) {
	return wrapSocketAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (address *SocketAddress) baseSocketAddress() *SocketAddress {
	return address
}

// BaseSocketAddress returns the underlying base object.
func BaseSocketAddress(obj SocketAddresser) *SocketAddress {
	return obj.baseSocketAddress()
}

// NewSocketAddressFromNative creates a Address subclass corresponding to the
// native struct sockaddr native.
//
// The function takes the following parameters:
//
//    - native: pointer to a struct sockaddr.
//    - len: size of the memory location pointed to by native.
//
// The function returns the following values:
//
//    - socketAddress: new Address if native could successfully be converted,
//      otherwise NULL.
//
func NewSocketAddressFromNative(native unsafe.Pointer, len uint) *SocketAddress {
	var _args [2]girepository.Argument
	var _arg0 C.gpointer // out
	var _arg1 C.gsize    // out
	var _cret *C.void    // in

	_arg0 = (C.gpointer)(unsafe.Pointer(native))
	_arg1 = C.gsize(len)

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gsize)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "SocketAddress").InvokeMethod("new_SocketAddress_from_native", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(native)
	runtime.KeepAlive(len)

	var _socketAddress *SocketAddress // out

	_socketAddress = wrapSocketAddress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketAddress
}

// NativeSize gets the size of address's native struct sockaddr. You can use
// this to allocate memory to pass to g_socket_address_to_native().
//
// The function returns the following values:
//
//    - gssize: size of the native struct sockaddr that address represents.
//
func (address *SocketAddress) NativeSize() int {
	var _args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gssize // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "SocketAddress").InvokeMethod("get_native_size", _args[:], nil)
	_cret = *(*C.gssize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

// ToNative converts a Address to a native struct sockaddr, which can be passed
// to low-level functions like connect() or bind().
//
// If not enough space is available, a G_IO_ERROR_NO_SPACE error is returned. If
// the address type is not known on the system then a G_IO_ERROR_NOT_SUPPORTED
// error is returned.
//
// The function takes the following parameters:
//
//    - dest (optional): pointer to a memory location that will contain the
//      native struct sockaddr.
//    - destlen: size of dest. Must be at least as large as
//      g_socket_address_get_native_size().
//
func (address *SocketAddress) ToNative(dest unsafe.Pointer, destlen uint) error {
	var _args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gpointer // out
	var _arg2 C.gsize    // out
	var _cerr *C.void    // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dest))
	_arg2 = C.gsize(destlen)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gsize)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gio", "SocketAddress").InvokeMethod("to_native", _args[:], nil)

	runtime.KeepAlive(address)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(destlen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
