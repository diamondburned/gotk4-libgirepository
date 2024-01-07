// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeVFS = coreglib.Type(girepository.MustFind("Gio", "Vfs").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVFS, F: marshalVFS},
	})
}

// VFS_EXTENSION_POINT_NAME: extension point for #GVfs functionality. See
// [Extending GIO][extending-gio].
const VFS_EXTENSION_POINT_NAME = "gio-vfs"

// VFSOverrides contains methods that are overridable.
type VFSOverrides struct {
}

func defaultVFSOverrides(v *VFS) VFSOverrides {
	return VFSOverrides{}
}

// VFS: entry point for using GIO functionality.
type VFS struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*VFS)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VFS, *VFSClass, VFSOverrides](
		GTypeVFS,
		initVFSClass,
		wrapVFS,
		defaultVFSOverrides,
	)
}

func initVFSClass(gclass unsafe.Pointer, overrides VFSOverrides, classInitFunc func(*VFSClass)) {
	if classInitFunc != nil {
		class := (*VFSClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVFS(obj *coreglib.Object) *VFS {
	return &VFS{
		Object: obj,
	}
}

func marshalVFS(p uintptr) (interface{}, error) {
	return wrapVFS(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// VFSClass: instance of this type is always passed by reference.
type VFSClass struct {
	*vfsClass
}

// vfsClass is the struct that's finalized.
type vfsClass struct {
	native unsafe.Pointer
}

var GIRInfoVFSClass = girepository.MustFind("Gio", "VfsClass")
