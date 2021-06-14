// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_properties_get_type()), F: marshalEditableProperties},
		{T: externglib.Type(C.gtk_debug_flags_get_type()), F: marshalDebugFlags},
	})
}

type EditableProperties int

const (
	EditablePropertiesPropText           EditableProperties = 0
	EditablePropertiesPropCursorPosition EditableProperties = 1
	EditablePropertiesPropSelectionBound EditableProperties = 2
	EditablePropertiesPropEditable       EditableProperties = 3
	EditablePropertiesPropWidthChars     EditableProperties = 4
	EditablePropertiesPropMaxWidthChars  EditableProperties = 5
	EditablePropertiesPropXalign         EditableProperties = 6
	EditablePropertiesPropEnableUndo     EditableProperties = 7
	EditablePropertiesNumProperties      EditableProperties = 8
)

func marshalEditableProperties(p uintptr) (interface{}, error) {
	return EditableProperties(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type DebugFlags int

const (
	DebugFlagsText           DebugFlags = 1
	DebugFlagsTree           DebugFlags = 2
	DebugFlagsKeybindings    DebugFlags = 4
	DebugFlagsModules        DebugFlags = 8
	DebugFlagsGeometry       DebugFlags = 16
	DebugFlagsIcontheme      DebugFlags = 32
	DebugFlagsPrinting       DebugFlags = 64
	DebugFlagsBuilder        DebugFlags = 128
	DebugFlagsSizeRequest    DebugFlags = 256
	DebugFlagsNoCSSCache     DebugFlags = 512
	DebugFlagsInteractive    DebugFlags = 1024
	DebugFlagsTouchscreen    DebugFlags = 2048
	DebugFlagsActions        DebugFlags = 4096
	DebugFlagsLayout         DebugFlags = 8192
	DebugFlagsSnapshot       DebugFlags = 16384
	DebugFlagsConstraints    DebugFlags = 32768
	DebugFlagsBuilderObjects DebugFlags = 65536
	DebugFlagsA11Y           DebugFlags = 131072
	DebugFlagsIconfallback   DebugFlags = 262144
)

func marshalDebugFlags(p uintptr) (interface{}, error) {
	return DebugFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
