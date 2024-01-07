// Code generated by girgen. DO NOT EDIT.

package gio

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// FILE_ATTRIBUTE_DOS_IS_MOUNTPOINT: key in the "dos" namespace for checking if
// the file is a NTFS mount point (a volume mount or a junction point). This
// attribute is TRUE if file is a reparse point of type
// IO_REPARSE_TAG_MOUNT_POINT
// (https://msdn.microsoft.com/en-us/library/dd541667.aspx). This attribute is
// only available for DOS file systems. Corresponding AttributeType is
// G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
const FILE_ATTRIBUTE_DOS_IS_MOUNTPOINT = "dos::is-mountpoint"

// FILE_ATTRIBUTE_DOS_REPARSE_POINT_TAG: key in the "dos" namespace for getting
// the file NTFS reparse tag. This value is 0 for files that are not reparse
// points. See the Reparse Tags
// (https://msdn.microsoft.com/en-us/library/dd541667.aspx) page for possible
// reparse tag values. Corresponding AttributeType is
// G_FILE_ATTRIBUTE_TYPE_UINT32.
const FILE_ATTRIBUTE_DOS_REPARSE_POINT_TAG = "dos::reparse-point-tag"
