// Code generated by girgen. DO NOT EDIT.

package gtk

import ()

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// ListBoxForEachFunc: function used by gtk_list_box_selected_foreach(). It will
// be called on every selected child of the box.
type ListBoxForEachFunc func(box *ListBox, row *ListBoxRow)