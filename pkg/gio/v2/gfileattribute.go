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
	GTypeFileAttributeInfoList = coreglib.Type(girepository.MustFind("Gio", "FileAttributeInfoList").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileAttributeInfoList, F: marshalFileAttributeInfoList},
	})
}

// FileAttributeInfo: information about a specific attribute.
//
// An instance of this type is always passed by reference.
type FileAttributeInfo struct {
	*fileAttributeInfo
}

// fileAttributeInfo is the struct that's finalized.
type fileAttributeInfo struct {
	native unsafe.Pointer
}

var GIRInfoFileAttributeInfo = girepository.MustFind("Gio", "FileAttributeInfo")

// FileAttributeInfoList acts as a lightweight registry for possible valid file
// attributes. The registry stores Key-Value pair formats as AttributeInfos.
//
// An instance of this type is always passed by reference.
type FileAttributeInfoList struct {
	*fileAttributeInfoList
}

// fileAttributeInfoList is the struct that's finalized.
type fileAttributeInfoList struct {
	native unsafe.Pointer
}

var GIRInfoFileAttributeInfoList = girepository.MustFind("Gio", "FileAttributeInfoList")

func marshalFileAttributeInfoList(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &FileAttributeInfoList{&fileAttributeInfoList{(unsafe.Pointer)(b)}}, nil
}

// Infos: array of AttributeInfos.
func (f *FileAttributeInfoList) Infos() *FileAttributeInfo {
	offset := GIRInfoFileAttributeInfoList.StructFieldOffset("infos")
	valptr := (**FileAttributeInfo)(unsafe.Add(f.native, offset))
	var _v *FileAttributeInfo // out
	_v = (*FileAttributeInfo)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// NInfos: number of values in the array.
func (f *FileAttributeInfoList) NInfos() int {
	offset := GIRInfoFileAttributeInfoList.StructFieldOffset("n_infos")
	valptr := (*int)(unsafe.Add(f.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// NInfos: number of values in the array.
func (f *FileAttributeInfoList) SetNInfos(nInfos int) {
	offset := GIRInfoFileAttributeInfoList.StructFieldOffset("n_infos")
	valptr := (*C.int)(unsafe.Add(f.native, offset))
	*valptr = C.int(nInfos)
}
