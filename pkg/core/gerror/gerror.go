package gerror

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gmodule.h>
// #include <glib.h>
import "C"

import (
	"reflect"
	"sync"
	"unsafe"

	"golang.org/x/sync/singleflight"
)

// ErrorCoder is an interface that returns a GError code. Errors may optionally
// implement this interface to override the default error code.
type ErrorCoder interface {
	ErrorCode() int
}

var (
	quarkMap    sync.Map // reflect.Type -> uint32
	quarkFlight singleflight.Group
)

// getQuark returns the quark associated with the given error. It registers the
// type of the given error using reflect.
func getQuark(err error) C.GQuark {
	// This is actually very cheap.
	typ := reflect.TypeOf(err)

	quark, ok := quarkMap.Load(typ)
	if ok {
		return C.GQuark(quark.(uint32))
	}

	var typeName string

	// cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
	// BenchmarkReflectType-8   	 9097218	       123.1 ns/op

	pkgPath := typ.PkgPath()
	if pkgPath == "" {
		typeName = typ.Name()
	} else {
		typeName = pkgPath + "." + typ.Name()
	}

	quark, _, _ = quarkFlight.Do(pkgPath, func() (interface{}, error) {
		// Allocate the quark string inside the singlefligth group, so the
		// pointer is ensured to be the same for this type. This string must not
		// be freed, since g_quark_from_static_string will use it directly.
		quarkString := (*C.gchar)(C.CString(typeName))

		// Immediately convert this to uint32, since I don't trust interfaces
		// holding C types.
		quark := uint32(C.g_quark_from_static_string(quarkString))
		quarkMap.Store(typ, quark)

		return quark, nil
	})

	return C.GQuark(quark.(uint32))
}

// New creates a new *C.GError from the given error. The caller is responsible
// for freeing the error with g_error_free().
func New(err error) unsafe.Pointer {
	if err == nil {
		return nil
	}

	var code int
	if coder, ok := err.(ErrorCoder); ok {
		code = coder.ErrorCode()
	}

	errString := (*C.gchar)(C.CString(err.Error()))
	defer C.free(unsafe.Pointer(errString))

	return unsafe.Pointer(C.g_error_new_literal(getQuark(err), C.gint(code), errString))
}

// glibError is converted from a GError to implement Go's error interface.
type glibError struct {
	quark uint32
	code  int
	err   string
}

// Quark returns the internal quark for the error. Callers that want this quark
// must manually type assert using their own interface.
func (err glibError) Quark() uint32 {
	return err.quark
}

func (err glibError) ErrorCode() int {
	return err.code
}

func (err glibError) Error() string {
	return err.err
}

// Take returns a new Go error from a *GError and frees the *GError. If the
// *GError is nil, then a nil error is returned.
func Take(gerror unsafe.Pointer) error {
	if gerror == nil {
		return nil
	}

	v := (*C.GError)(gerror)
	defer C.g_error_free(v)

	return glibError{
		quark: uint32(v.domain),
		code:  int(v.code),
		err:   C.GoString(v.message),
	}
}
