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

// FlowBoxCreateWidgetFunc: called for flow boxes that are bound to a Model with
// gtk_flow_box_bind_model() for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func(item *coreglib.Object) (widget Widgetter)