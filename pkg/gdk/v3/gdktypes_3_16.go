// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
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
	GTypeGLError = coreglib.Type(girepository.MustFind("Gdk", "GLError").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLError, F: marshalGLError},
	})
}

// GLError: error enumeration for GLContext.
type GLError C.gint

const (
	// GLErrorNotAvailable: openGL support is not available.
	GLErrorNotAvailable GLError = iota
	// GLErrorUnsupportedFormat: requested visual format is not supported.
	GLErrorUnsupportedFormat
	// GLErrorUnsupportedProfile: requested profile is not supported.
	GLErrorUnsupportedProfile
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLError.
func (g GLError) String() string {
	switch g {
	case GLErrorNotAvailable:
		return "NotAvailable"
	case GLErrorUnsupportedFormat:
		return "UnsupportedFormat"
	case GLErrorUnsupportedProfile:
		return "UnsupportedProfile"
	default:
		return fmt.Sprintf("GLError(%d)", g)
	}
}
