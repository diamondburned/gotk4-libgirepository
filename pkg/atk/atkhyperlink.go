// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern AtkObject* _gotk4_atk1_HyperlinkClass_get_object(void*, gint);
// extern gboolean _gotk4_atk1_HyperlinkClass_is_selected_link(void*);
// extern gboolean _gotk4_atk1_HyperlinkClass_is_valid(void*);
// extern gchar* _gotk4_atk1_HyperlinkClass_get_uri(void*, gint);
// extern gint _gotk4_atk1_HyperlinkClass_get_end_index(void*);
// extern gint _gotk4_atk1_HyperlinkClass_get_n_anchors(void*);
// extern gint _gotk4_atk1_HyperlinkClass_get_start_index(void*);
// extern guint _gotk4_atk1_HyperlinkClass_link_state(void*);
// extern void _gotk4_atk1_HyperlinkClass_link_activated(void*);
// extern void _gotk4_atk1_Hyperlink_ConnectLinkActivated(gpointer, guintptr);
import "C"

// glib.Type values for atkhyperlink.go.
var (
	GTypeHyperlinkStateFlags = coreglib.Type(C.atk_hyperlink_state_flags_get_type())
	GTypeHyperlink           = coreglib.Type(C.atk_hyperlink_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeHyperlinkStateFlags, F: marshalHyperlinkStateFlags},
		{T: GTypeHyperlink, F: marshalHyperlink},
	})
}

// HyperlinkStateFlags describes the type of link.
type HyperlinkStateFlags C.guint

const (
	// HyperlinkIsInline: link is inline.
	HyperlinkIsInline HyperlinkStateFlags = 0b1
)

func marshalHyperlinkStateFlags(p uintptr) (interface{}, error) {
	return HyperlinkStateFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for HyperlinkStateFlags.
func (h HyperlinkStateFlags) String() string {
	if h == 0 {
		return "HyperlinkStateFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(17)

	for h != 0 {
		next := h & (h - 1)
		bit := h - next

		switch bit {
		case HyperlinkIsInline:
			builder.WriteString("Inline|")
		default:
			builder.WriteString(fmt.Sprintf("HyperlinkStateFlags(0b%b)|", bit))
		}

		h = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if h contains other.
func (h HyperlinkStateFlags) Has(other HyperlinkStateFlags) bool {
	return (h & other) == other
}

// HyperlinkOverrider contains methods that are overridable.
type HyperlinkOverrider interface {
	// EndIndex gets the index with the hypertext document at which this link
	// ends.
	//
	// The function returns the following values:
	//
	//    - gint: index with the hypertext document at which this link ends.
	//
	EndIndex() int32
	// NAnchors gets the number of anchors associated with this hyperlink.
	//
	// The function returns the following values:
	//
	//    - gint: number of anchors associated with this hyperlink.
	//
	NAnchors() int32
	// GetObject returns the item associated with this hyperlinks nth anchor.
	// For instance, the returned Object will implement Text if link_ is a text
	// hyperlink, Image if link_ is an image hyperlink etc.
	//
	// Multiple anchors are primarily used by client-side image maps.
	//
	// The function takes the following parameters:
	//
	//    - i: (zero-index) integer specifying the desired anchor.
	//
	// The function returns the following values:
	//
	//    - object associated with this hyperlinks i-th anchor.
	//
	GetObject(i int32) *ObjectClass
	// StartIndex gets the index with the hypertext document at which this link
	// begins.
	//
	// The function returns the following values:
	//
	//    - gint: index with the hypertext document at which this link begins.
	//
	StartIndex() int32
	// URI: get a the URI associated with the anchor specified by i of link_.
	//
	// Multiple anchors are primarily used by client-side image maps.
	//
	// The function takes the following parameters:
	//
	//    - i: (zero-index) integer specifying the desired anchor.
	//
	// The function returns the following values:
	//
	//    - utf8: string specifying the URI.
	//
	URI(i int32) string
	// IsSelectedLink determines whether this AtkHyperlink is selected
	//
	// Deprecated: Please use ATK_STATE_FOCUSABLE for all links, and
	// ATK_STATE_FOCUSED for focused links.
	//
	// The function returns the following values:
	//
	//    - ok: true if the AtkHyperlink is selected, False otherwise.
	//
	IsSelectedLink() bool
	// IsValid: since the document that a link is associated with may have
	// changed this method returns TRUE if the link is still valid (with respect
	// to the document it references) and FALSE otherwise.
	//
	// The function returns the following values:
	//
	//    - ok: whether or not this link is still valid.
	//
	IsValid() bool
	LinkActivated()
	// The function returns the following values:
	//
	LinkState() uint32
}

// Hyperlink: ATK object which encapsulates a link or set of links (for instance
// in the case of client-side image maps) in a hypertext document. It may
// implement the AtkAction interface. AtkHyperlink may also be used to refer to
// inline embedded content, since it allows specification of a start and end
// offset within the host AtkHypertext object.
type Hyperlink struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Action
}

var (
	_ coreglib.Objector = (*Hyperlink)(nil)
)

func classInitHyperlinker(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.AtkHyperlinkClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.AtkHyperlinkClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ EndIndex() int32 }); ok {
		pclass.get_end_index = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_get_end_index)
	}

	if _, ok := goval.(interface{ NAnchors() int32 }); ok {
		pclass.get_n_anchors = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_get_n_anchors)
	}

	if _, ok := goval.(interface{ GetObject(i int32) *ObjectClass }); ok {
		pclass.get_object = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_get_object)
	}

	if _, ok := goval.(interface{ StartIndex() int32 }); ok {
		pclass.get_start_index = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_get_start_index)
	}

	if _, ok := goval.(interface{ URI(i int32) string }); ok {
		pclass.get_uri = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_get_uri)
	}

	if _, ok := goval.(interface{ IsSelectedLink() bool }); ok {
		pclass.is_selected_link = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_is_selected_link)
	}

	if _, ok := goval.(interface{ IsValid() bool }); ok {
		pclass.is_valid = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_is_valid)
	}

	if _, ok := goval.(interface{ LinkActivated() }); ok {
		pclass.link_activated = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_link_activated)
	}

	if _, ok := goval.(interface{ LinkState() uint32 }); ok {
		pclass.link_state = (*[0]byte)(C._gotk4_atk1_HyperlinkClass_link_state)
	}
}

//export _gotk4_atk1_HyperlinkClass_get_end_index
func _gotk4_atk1_HyperlinkClass_get_end_index(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ EndIndex() int32 })

	gint := iface.EndIndex()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_n_anchors
func _gotk4_atk1_HyperlinkClass_get_n_anchors(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NAnchors() int32 })

	gint := iface.NAnchors()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_object
func _gotk4_atk1_HyperlinkClass_get_object(arg0 *C.void, arg1 C.gint) (cret *C.AtkObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ GetObject(i int32) *ObjectClass })

	var _i int32 // out

	_i = int32(arg1)

	object := iface.GetObject(_i)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_start_index
func _gotk4_atk1_HyperlinkClass_get_start_index(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ StartIndex() int32 })

	gint := iface.StartIndex()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_uri
func _gotk4_atk1_HyperlinkClass_get_uri(arg0 *C.void, arg1 C.gint) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ URI(i int32) string })

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.URI(_i)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_atk1_HyperlinkClass_is_selected_link
func _gotk4_atk1_HyperlinkClass_is_selected_link(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ IsSelectedLink() bool })

	ok := iface.IsSelectedLink()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_HyperlinkClass_is_valid
func _gotk4_atk1_HyperlinkClass_is_valid(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ IsValid() bool })

	ok := iface.IsValid()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_HyperlinkClass_link_activated
func _gotk4_atk1_HyperlinkClass_link_activated(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ LinkActivated() })

	iface.LinkActivated()
}

//export _gotk4_atk1_HyperlinkClass_link_state
func _gotk4_atk1_HyperlinkClass_link_state(arg0 *C.void) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ LinkState() uint32 })

	guint := iface.LinkState()

	cret = C.guint(guint)

	return cret
}

func wrapHyperlink(obj *coreglib.Object) *Hyperlink {
	return &Hyperlink{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalHyperlink(p uintptr) (interface{}, error) {
	return wrapHyperlink(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Hyperlink_ConnectLinkActivated
func _gotk4_atk1_Hyperlink_ConnectLinkActivated(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLinkActivated: signal link-activated is emitted when a link is
// activated.
func (link_ *Hyperlink) ConnectLinkActivated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(link_, "link-activated", false, unsafe.Pointer(C._gotk4_atk1_Hyperlink_ConnectLinkActivated), f)
}

// EndIndex gets the index with the hypertext document at which this link ends.
//
// The function returns the following values:
//
//    - gint: index with the hypertext document at which this link ends.
//
func (link_ *Hyperlink) EndIndex() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("get_end_index", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NAnchors gets the number of anchors associated with this hyperlink.
//
// The function returns the following values:
//
//    - gint: number of anchors associated with this hyperlink.
//
func (link_ *Hyperlink) NAnchors() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("get_n_anchors", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// GetObject returns the item associated with this hyperlinks nth anchor. For
// instance, the returned Object will implement Text if link_ is a text
// hyperlink, Image if link_ is an image hyperlink etc.
//
// Multiple anchors are primarily used by client-side image maps.
//
// The function takes the following parameters:
//
//    - i: (zero-index) integer specifying the desired anchor.
//
// The function returns the following values:
//
//    - object associated with this hyperlinks i-th anchor.
//
func (link_ *Hyperlink) GetObject(i int32) *ObjectClass {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))
	_arg1 = C.gint(i)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("get_object", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)
	runtime.KeepAlive(i)

	var _object *ObjectClass // out

	_object = wrapObject(coreglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// StartIndex gets the index with the hypertext document at which this link
// begins.
//
// The function returns the following values:
//
//    - gint: index with the hypertext document at which this link begins.
//
func (link_ *Hyperlink) StartIndex() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("get_start_index", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// URI: get a the URI associated with the anchor specified by i of link_.
//
// Multiple anchors are primarily used by client-side image maps.
//
// The function takes the following parameters:
//
//    - i: (zero-index) integer specifying the desired anchor.
//
// The function returns the following values:
//
//    - utf8: string specifying the URI.
//
func (link_ *Hyperlink) URI(i int32) string {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))
	_arg1 = C.gint(i)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("get_uri", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)
	runtime.KeepAlive(i)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// IsInline indicates whether the link currently displays some or all of its
// content inline. Ordinary HTML links will usually return FALSE, but an inline
// &lt;src&gt; HTML element will return TRUE.
//
// The function returns the following values:
//
//    - ok: whether or not this link displays its content inline.
//
func (link_ *Hyperlink) IsInline() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("is_inline", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelectedLink determines whether this AtkHyperlink is selected
//
// Deprecated: Please use ATK_STATE_FOCUSABLE for all links, and
// ATK_STATE_FOCUSED for focused links.
//
// The function returns the following values:
//
//    - ok: true if the AtkHyperlink is selected, False otherwise.
//
func (link_ *Hyperlink) IsSelectedLink() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("is_selected_link", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsValid: since the document that a link is associated with may have changed
// this method returns TRUE if the link is still valid (with respect to the
// document it references) and FALSE otherwise.
//
// The function returns the following values:
//
//    - ok: whether or not this link is still valid.
//
func (link_ *Hyperlink) IsValid() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(link_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Atk", "Hyperlink").InvokeMethod("is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(link_)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
