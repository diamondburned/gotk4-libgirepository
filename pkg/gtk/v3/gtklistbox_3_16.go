// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a Model with
// gtk_list_box_bind_model() for each item that gets added to the model.
//
// Versions of GTK+ prior to 3.18 called gtk_widget_show_all() on the rows
// created by the GtkListBoxCreateWidgetFunc, but this forced all widgets inside
// the row to be shown, and is no longer the case. Applications should be
// updated to show the desired row widgets.
type ListBoxCreateWidgetFunc func(item *coreglib.Object) (widget Widgetter)
