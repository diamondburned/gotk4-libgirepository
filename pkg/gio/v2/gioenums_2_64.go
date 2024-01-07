// Code generated by girgen. DO NOT EDIT.

package gio

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
	GTypeMemoryMonitorWarningLevel = coreglib.Type(girepository.MustFind("Gio", "MemoryMonitorWarningLevel").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMemoryMonitorWarningLevel, F: marshalMemoryMonitorWarningLevel},
	})
}

// MemoryMonitorWarningLevel: memory availability warning levels.
//
// Note that because new values might be added, it is recommended that
// applications check MonitorWarningLevel as ranges, for example:
//
//    if (warning_level > G_MEMORY_MONITOR_WARNING_LEVEL_LOW)
//      drop_caches ();.
type MemoryMonitorWarningLevel C.gint

const (
	// MemoryMonitorWarningLevelLow: memory on the device is low, processes
	// should free up unneeded resources (for example, in-memory caches) so they
	// can be used elsewhere.
	MemoryMonitorWarningLevelLow MemoryMonitorWarningLevel = 50
	// MemoryMonitorWarningLevelMedium: same as
	// G_MEMORY_MONITOR_WARNING_LEVEL_LOW but the device has even less free
	// memory, so processes should try harder to free up unneeded resources. If
	// your process does not need to stay running, it is a good time for it to
	// quit.
	MemoryMonitorWarningLevelMedium MemoryMonitorWarningLevel = 100
	// MemoryMonitorWarningLevelCritical: system will soon start terminating
	// processes to reclaim memory, including background processes.
	MemoryMonitorWarningLevelCritical MemoryMonitorWarningLevel = 255
)

func marshalMemoryMonitorWarningLevel(p uintptr) (interface{}, error) {
	return MemoryMonitorWarningLevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MemoryMonitorWarningLevel.
func (m MemoryMonitorWarningLevel) String() string {
	switch m {
	case MemoryMonitorWarningLevelLow:
		return "Low"
	case MemoryMonitorWarningLevelMedium:
		return "Medium"
	case MemoryMonitorWarningLevelCritical:
		return "Critical"
	default:
		return fmt.Sprintf("MemoryMonitorWarningLevel(%d)", m)
	}
}