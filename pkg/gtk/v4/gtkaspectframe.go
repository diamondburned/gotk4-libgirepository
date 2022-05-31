// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkaspectframe.go.
var GTypeAspectFrame = coreglib.Type(C.gtk_aspect_frame_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAspectFrame, F: marshalAspectFrame},
	})
}

// AspectFrame: GtkAspectFrame preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*AspectFrame)(nil)
)

func wrapAspectFrame(obj *coreglib.Object) *AspectFrame {
	return &AspectFrame{
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
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	return wrapAspectFrame(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAspectFrame: create a new GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment of the child within the parent. Ranges from
//      0.0 (left aligned) to 1.0 (right aligned).
//    - yalign: vertical alignment of the child within the parent. Ranges from
//      0.0 (top aligned) to 1.0 (bottom aligned).
//    - ratio: desired aspect ratio.
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken from
//      the requistion of the child.
//
// The function returns the following values:
//
//    - aspectFrame: new GtkAspectFrame.
//
func NewAspectFrame(xalign, yalign, ratio float32, obeyChild bool) *AspectFrame {
	var _args [4]girepository.Argument
	var _arg0 C.float    // out
	var _arg1 C.float    // out
	var _arg2 C.float    // out
	var _arg3 C.gboolean // out
	var _cret *C.void    // in

	_arg0 = C.float(xalign)
	_arg1 = C.float(yalign)
	_arg2 = C.float(ratio)
	if obeyChild {
		_arg3 = C.TRUE
	}

	*(*C.float)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.float)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.float)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.gboolean)(unsafe.Pointer(&_args[3])) = _arg3

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("new_AspectFrame", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(ratio)
	runtime.KeepAlive(obeyChild)

	var _aspectFrame *AspectFrame // out

	_aspectFrame = wrapAspectFrame(coreglib.Take(unsafe.Pointer(_cret)))

	return _aspectFrame
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self@.
//
func (self *AspectFrame) Child() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// ObeyChild returns whether the child's size request should override the set
// aspect ratio of the GtkAspectFrame.
//
// The function returns the following values:
//
//    - ok: whether to obey the child's size request.
//
func (self *AspectFrame) ObeyChild() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_obey_child", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ratio returns the desired aspect ratio of the child.
//
// The function returns the following values:
//
//    - gfloat: desired aspect ratio.
//
func (self *AspectFrame) Ratio() float32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.float // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_ratio", _args[:], nil)
	_cret = *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// XAlign returns the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function returns the following values:
//
//    - gfloat: horizontal alignment.
//
func (self *AspectFrame) XAlign() float32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.float // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_xalign", _args[:], nil)
	_cret = *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign returns the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function returns the following values:
//
//    - gfloat: vertical alignment.
//
func (self *AspectFrame) YAlign() float32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.float // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_yalign", _args[:], nil)
	_cret = *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *AspectFrame) SetChild(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetObeyChild sets whether the aspect ratio of the child's size request should
// override the set aspect ratio of the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken from
//      the requistion of the child.
//
func (self *AspectFrame) SetObeyChild(obeyChild bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if obeyChild {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_obey_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(obeyChild)
}

// SetRatio sets the desired aspect ratio of the child.
//
// The function takes the following parameters:
//
//    - ratio: aspect ratio of the child.
//
func (self *AspectFrame) SetRatio(ratio float32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.float // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.float(ratio)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.float)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_ratio", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(ratio)
}

// SetXAlign sets the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment, from 0.0 (left aligned) to 1.0 (right
//      aligned).
//
func (self *AspectFrame) SetXAlign(xalign float32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.float // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.float(xalign)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.float)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_xalign", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - yalign: horizontal alignment, from 0.0 (top aligned) to 1.0 (bottom
//      aligned).
//
func (self *AspectFrame) SetYAlign(yalign float32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.float // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.float(yalign)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.float)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_yalign", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}
