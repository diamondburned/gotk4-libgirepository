// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern AtkHyperlink* _gotk4_atk1_HyperlinkImplIface_get_hyperlink(AtkHyperlinkImpl*);
import "C"

// GType values.
var (
	GTypeHyperlinkImpl = coreglib.Type(C.atk_hyperlink_impl_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHyperlinkImpl, F: marshalHyperlinkImpl},
	})
}

// HyperlinkImplOverrider contains methods that are overridable.
type HyperlinkImplOverrider interface {
	// Hyperlink gets the hyperlink associated with this object.
	//
	// The function returns the following values:
	//
	//    - hyperlink: atkHyperlink object which points to this implementing
	//      AtkObject.
	//
	Hyperlink() *Hyperlink
}

// HyperlinkImpl allows AtkObjects to refer to their associated AtkHyperlink
// instance, if one exists. AtkHyperlinkImpl differs from AtkHyperlink in that
// AtkHyperlinkImpl is an interface, whereas AtkHyperlink is a object type. The
// AtkHyperlinkImpl interface allows a client to query an AtkObject for the
// availability of an associated AtkHyperlink instance, and obtain that
// instance. It is thus particularly useful in cases where embedded content or
// inline content within a text object is present, since the embedding text
// object implements AtkHypertext and the inline/embedded objects are exposed as
// children which implement AtkHyperlinkImpl, in addition to their being
// obtainable via AtkHypertext:getLink followed by AtkHyperlink:getObject.
//
// The AtkHyperlinkImpl interface should be supported by objects exposed within
// the hierarchy as children of an AtkHypertext container which correspond to
// "links" or embedded content within the text. HTML anchors are not, for
// instance, normally exposed this way, but embedded images and components which
// appear inline in the content of a text object are. The AtkHyperlinkIface
// interface allows a means of determining which children are hyperlinks in this
// sense of the word, and for obtaining their corresponding AtkHyperlink object,
// from which the embedding range, URI, etc. can be obtained.
//
// To some extent this interface exists because, for historical reasons,
// AtkHyperlink was defined as an object type, not an interface. Thus, in order
// to interact with AtkObjects via AtkHyperlink semantics, a new interface was
// required.
//
// HyperlinkImpl wraps an interface. This means the user can get the
// underlying type by calling Cast().
type HyperlinkImpl struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*HyperlinkImpl)(nil)
)

// HyperlinkImpler describes HyperlinkImpl's interface methods.
type HyperlinkImpler interface {
	coreglib.Objector

	// Hyperlink gets the hyperlink associated with this object.
	Hyperlink() *Hyperlink
}

var _ HyperlinkImpler = (*HyperlinkImpl)(nil)

func ifaceInitHyperlinkImpler(gifacePtr, data C.gpointer) {
	iface := (*C.AtkHyperlinkImplIface)(unsafe.Pointer(gifacePtr))
	iface.get_hyperlink = (*[0]byte)(C._gotk4_atk1_HyperlinkImplIface_get_hyperlink)
}

//export _gotk4_atk1_HyperlinkImplIface_get_hyperlink
func _gotk4_atk1_HyperlinkImplIface_get_hyperlink(arg0 *C.AtkHyperlinkImpl) (cret *C.AtkHyperlink) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(HyperlinkImplOverrider)

	hyperlink := iface.Hyperlink()

	cret = (*C.AtkHyperlink)(unsafe.Pointer(coreglib.InternObject(hyperlink).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(hyperlink).Native()))

	return cret
}

func wrapHyperlinkImpl(obj *coreglib.Object) *HyperlinkImpl {
	return &HyperlinkImpl{
		Object: obj,
	}
}

func marshalHyperlinkImpl(p uintptr) (interface{}, error) {
	return wrapHyperlinkImpl(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Hyperlink gets the hyperlink associated with this object.
//
// The function returns the following values:
//
//    - hyperlink: atkHyperlink object which points to this implementing
//      AtkObject.
//
func (impl *HyperlinkImpl) Hyperlink() *Hyperlink {
	var _arg0 *C.AtkHyperlinkImpl // out
	var _cret *C.AtkHyperlink     // in

	_arg0 = (*C.AtkHyperlinkImpl)(unsafe.Pointer(coreglib.InternObject(impl).Native()))

	_cret = C.atk_hyperlink_impl_get_hyperlink(_arg0)
	runtime.KeepAlive(impl)

	var _hyperlink *Hyperlink // out

	_hyperlink = wrapHyperlink(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _hyperlink
}

// HyperlinkImplIface: instance of this type is always passed by reference.
type HyperlinkImplIface struct {
	*hyperlinkImplIface
}

// hyperlinkImplIface is the struct that's finalized.
type hyperlinkImplIface struct {
	native *C.AtkHyperlinkImplIface
}
