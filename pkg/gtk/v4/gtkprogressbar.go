// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeProgressBar = coreglib.Type(girepository.MustFind("Gtk", "ProgressBar").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeProgressBar, F: marshalProgressBar},
	})
}

// ProgressBar: GtkProgressBar is typically used to display the progress of a
// long running operation.
//
// It provides a visual clue that processing is underway. GtkProgressBar can be
// used in two different modes: percentage mode and activity mode.
//
// !An example GtkProgressBar (progressbar.png)
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the GtkProgressBar in percentage mode and the user sees a growing bar
// indicating the percentage of the work that has been completed. In this mode,
// the application is required to call gtk.ProgressBar.SetFraction()
// periodically to update the progress bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the GtkProgressBar in activity mode, which shows activity by a
// block moving back and forth within the progress area. In this mode, the
// application is required to call gtk.ProgressBar.Pulse() periodically to
// update the progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// GtkProgressBar. Functions are provided to control the orientation of the bar,
// optional text can be displayed along with the bar, and the step size used in
// activity mode can be set.
//
// CSS nodes
//
//    progressbar[.osd]
//    ├── [text]
//    ╰── trough[.empty][.full]
//        ╰── progress[.pulse]
//
//
// GtkProgressBar has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
//
//
// Accessibility
//
// GtkProgressBar uses the K_ACCESSIBLE_ROLE_PROGRESS_BAR role.
type ProgressBar struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*ProgressBar)(nil)
	_ coreglib.Objector = (*ProgressBar)(nil)
)

func wrapProgressBar(obj *coreglib.Object) *ProgressBar {
	return &ProgressBar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	return wrapProgressBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
