// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	_ "runtime/cgo"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

func init() {
	girepository.Require("Gsk", "4.0", girepository.LoadFlagLazy)
}
