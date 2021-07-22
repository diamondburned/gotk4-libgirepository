// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_vulkan_context_get_type()), F: marshalVulkanContexter},
	})
}

// VulkanContext: GdkVulkanContext is an object representing the
// platform-specific Vulkan draw context.
//
// GdkVulkanContexts are created for a surface using
// gdk.Surface.CreateVulkanContext(), and the context will match the the
// characteristics of the surface.
//
// Support for GdkVulkanContext is platform-specific and context creation can
// fail, returning NULL context.
type VulkanContext struct {
	DrawContext

	gio.Initable
	*externglib.Object
}

// VulkanContexter describes VulkanContext's abstract methods.
type VulkanContexter interface {
	gextras.Objector

	privateVulkanContext()
}

var _ VulkanContexter = (*VulkanContext)(nil)

func wrapVulkanContext(obj *externglib.Object) *VulkanContext {
	return &VulkanContext{
		DrawContext: DrawContext{
			Object: obj,
		},
		Initable: gio.Initable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalVulkanContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVulkanContext(obj), nil
}

func (*VulkanContext) privateVulkanContext() {}
