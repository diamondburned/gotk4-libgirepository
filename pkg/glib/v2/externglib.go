// Code generated by girgen. DO NOT EDIT.

package glib

import (
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib.h>
import "C"

// IdleAdd is an alias for pkg/core/glib.IdleAdd.
func IdleAdd(f interface{}) SourceHandle {
	return externglib.IdleAdd(f)
}

// IdleAddPriority is an alias for pkg/core/glib.IdleAddPriority.
func IdleAddPriority(p Priority, f interface{}) SourceHandle {
	return externglib.IdleAddPriority(p, f)
}

// TimeoutAdd is an alias for pkg/core/glib.TimeoutAdd.
func TimeoutAdd(ms uint, f interface{}) SourceHandle {
	return externglib.TimeoutAdd(ms, f)
}

// TimeoutAddPriority is an alias for pkg/core/glib.TimeoutAddPriority.
func TimeoutAddPriority(ms uint, p Priority, f interface{}) SourceHandle {
	return externglib.TimeoutAddPriority(ms, p, f)
}

// TimeoutSecondsAdd is an alias for pkg/core/glib.TimeoutSecondsAdd.
func TimeoutSecondsAdd(s uint, f interface{}) SourceHandle {
	return externglib.TimeoutSecondsAdd(s, f)
}

// TimeoutSecondsAddPriority is an alias for pkg/core/glib.TimeoutSecondsAddPriority.
func TimeoutSecondsAddPriority(s uint, p Priority, f interface{}) SourceHandle {
	return externglib.TimeoutSecondsAddPriority(s, p, f)
}

// TypeFromName is an alias for pkg/core/glib.TypeFromName.
func TypeFromName(typeName string) Type {
	return externglib.TypeFromName(typeName)
}

// NewValue is an alias for pkg/core/glib.NewValue.
func NewValue(v interface{}) *Value {
	return externglib.NewValue(v)
}

// SourceRemove is an alias for pkg/core/glib.SourceRemove.
func SourceRemove(src SourceHandle) bool {
	return externglib.SourceRemove(src)
}

// ObjectEq is an alias for pkg/core/glib.ObjectEq.
func ObjectEq(obj1 Objector, obj2 Objector) bool {
	return externglib.ObjectEq(obj1, obj2)
}

// Object is an alias for pkg/core/glib.Object.
type Object = externglib.Object

// Objector is an alias for pkg/core/glib.Objector.
type Objector = externglib.Objector

// Type is an alias for pkg/core/glib.Type.
type Type = externglib.Type

// Value is an alias for pkg/core/glib.Value.
type Value = externglib.Value

// Priority is an alias for pkg/core/glib.Priority.
type Priority = externglib.Priority

// SourceHandle is an alias for pkg/core/glib.SourceHandle.
type SourceHandle = externglib.SourceHandle

// SignalHandle is an alias for pkg/core/glib.SignalHandle.
type SignalHandle = externglib.SignalHandle
