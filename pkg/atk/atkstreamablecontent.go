// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gint _gotk4_atk1_StreamableContentIface_get_n_mime_types(void*);
// extern void* _gotk4_atk1_StreamableContentIface_get_mime_type(void*, gint);
// extern void* _gotk4_atk1_StreamableContentIface_get_stream(void*, void*);
// extern void* _gotk4_atk1_StreamableContentIface_get_uri(void*, void*);
import "C"

// GTypeStreamableContent returns the GType for the type StreamableContent.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStreamableContent() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "StreamableContent").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStreamableContent)
	return gtype
}

// StreamableContentOverrider contains methods that are overridable.
type StreamableContentOverrider interface {
	// MIMEType gets the character string of the specified mime type. The first
	// mime type is at position 0, the second at position 1, and so on.
	//
	// The function takes the following parameters:
	//
	//    - i: gint representing the position of the mime type starting from 0.
	//
	// The function returns the following values:
	//
	//    - utf8: gchar* representing the specified mime type; the caller should
	//      not free the character string.
	//
	MIMEType(i int32) string
	// NMIMETypes gets the number of mime types supported by this object.
	//
	// The function returns the following values:
	//
	//    - gint which is the number of mime types supported by the object.
	//
	NMIMETypes() int32
	// Stream gets the content in the specified mime type.
	//
	// The function takes the following parameters:
	//
	//    - mimeType: gchar* representing the mime type.
	//
	// The function returns the following values:
	//
	//    - ioChannel which contains the content in the specified mime type.
	//
	Stream(mimeType string) *glib.IOChannel
	// URI: get a string representing a URI in IETF standard format (see
	// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may
	// be streamed in the specified mime-type, if one is available. If mime_type
	// is NULL, the URI for the default (and possibly only) mime-type is
	// returned.
	//
	// Note that it is possible for get_uri to return NULL but for get_stream to
	// work nonetheless, since not all GIOChannels connect to URIs.
	//
	// The function takes the following parameters:
	//
	//    - mimeType: gchar* representing the mime type, or NULL to request a URI
	//      for the default mime type.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional) returns a string representing a URI, or NULL if no
	//      corresponding URI can be constructed.
	//
	URI(mimeType string) string
}

// StreamableContent: interface whereby an object allows its backing content to
// be streamed to clients. Typical implementors would be images or icons, HTML
// content, or multimedia display/rendering widgets.
//
// Negotiation of content type is allowed. Clients may examine the backing data
// and transform, convert, or parse the content in order to present it in an
// alternate form to end-users.
//
// The AtkStreamableContent interface is particularly useful for saving,
// printing, or post-processing entire documents, or for persisting alternate
// views of a document. If document content itself is being serialized, stored,
// or converted, then use of the AtkStreamableContent interface can help address
// performance issues. Unlike most ATK interfaces, this interface is not
// strongly tied to the current user-agent view of the a particular document,
// but may in some cases give access to the underlying model data.
//
// StreamableContent wraps an interface. This means the user can get the
// underlying type by calling Cast().
type StreamableContent struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StreamableContent)(nil)
)

// StreamableContenter describes StreamableContent's interface methods.
type StreamableContenter interface {
	coreglib.Objector

	// MIMEType gets the character string of the specified mime type.
	MIMEType(i int32) string
	// NMIMETypes gets the number of mime types supported by this object.
	NMIMETypes() int32
	// Stream gets the content in the specified mime type.
	Stream(mimeType string) *glib.IOChannel
	// URI: get a string representing a URI in IETF standard format (see
	// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may
	// be streamed in the specified mime-type, if one is available.
	URI(mimeType string) string
}

var _ StreamableContenter = (*StreamableContent)(nil)

func ifaceInitStreamableContenter(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Atk", "StreamableContentIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("get_mime_type"))) = unsafe.Pointer(C._gotk4_atk1_StreamableContentIface_get_mime_type)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("get_n_mime_types"))) = unsafe.Pointer(C._gotk4_atk1_StreamableContentIface_get_n_mime_types)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("get_stream"))) = unsafe.Pointer(C._gotk4_atk1_StreamableContentIface_get_stream)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("get_uri"))) = unsafe.Pointer(C._gotk4_atk1_StreamableContentIface_get_uri)
}

//export _gotk4_atk1_StreamableContentIface_get_mime_type
func _gotk4_atk1_StreamableContentIface_get_mime_type(arg0 *C.void, arg1 C.gint) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.MIMEType(_i)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_n_mime_types
func _gotk4_atk1_StreamableContentIface_get_n_mime_types(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	gint := iface.NMIMETypes()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_stream
func _gotk4_atk1_StreamableContentIface_get_stream(arg0 *C.void, arg1 *C.void) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _mimeType string // out

	_mimeType = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ioChannel := iface.Stream(_mimeType)

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(ioChannel)))

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_uri
func _gotk4_atk1_StreamableContentIface_get_uri(arg0 *C.void, arg1 *C.void) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _mimeType string // out

	_mimeType = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.URI(_mimeType)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

func wrapStreamableContent(obj *coreglib.Object) *StreamableContent {
	return &StreamableContent{
		Object: obj,
	}
}

func marshalStreamableContent(p uintptr) (interface{}, error) {
	return wrapStreamableContent(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MIMEType gets the character string of the specified mime type. The first mime
// type is at position 0, the second at position 1, and so on.
//
// The function takes the following parameters:
//
//    - i: gint representing the position of the mime type starting from 0.
//
// The function returns the following values:
//
//    - utf8: gchar* representing the specified mime type; the caller should not
//      free the character string.
//
func (streamable *StreamableContent) MIMEType(i int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(streamable).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_info := girepository.MustFind("Atk", "StreamableContent")
	_gret := _info.InvokeIfaceMethod("get_mime_type", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(streamable)
	runtime.KeepAlive(i)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NMIMETypes gets the number of mime types supported by this object.
//
// The function returns the following values:
//
//    - gint which is the number of mime types supported by the object.
//
func (streamable *StreamableContent) NMIMETypes() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(streamable).Native()))

	_info := girepository.MustFind("Atk", "StreamableContent")
	_gret := _info.InvokeIfaceMethod("get_n_mime_types", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(streamable)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Stream gets the content in the specified mime type.
//
// The function takes the following parameters:
//
//    - mimeType: gchar* representing the mime type.
//
// The function returns the following values:
//
//    - ioChannel which contains the content in the specified mime type.
//
func (streamable *StreamableContent) Stream(mimeType string) *glib.IOChannel {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(streamable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Atk", "StreamableContent")
	_gret := _info.InvokeIfaceMethod("get_stream", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(streamable)
	runtime.KeepAlive(mimeType)

	var _ioChannel *glib.IOChannel // out

	_ioChannel = (*glib.IOChannel)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ioChannel)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _ioChannel
}

// URI: get a string representing a URI in IETF standard format (see
// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may be
// streamed in the specified mime-type, if one is available. If mime_type is
// NULL, the URI for the default (and possibly only) mime-type is returned.
//
// Note that it is possible for get_uri to return NULL but for get_stream to
// work nonetheless, since not all GIOChannels connect to URIs.
//
// The function takes the following parameters:
//
//    - mimeType: gchar* representing the mime type, or NULL to request a URI for
//      the default mime type.
//
// The function returns the following values:
//
//    - utf8 (optional) returns a string representing a URI, or NULL if no
//      corresponding URI can be constructed.
//
func (streamable *StreamableContent) URI(mimeType string) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(streamable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Atk", "StreamableContent")
	_gret := _info.InvokeIfaceMethod("get_uri", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(streamable)
	runtime.KeepAlive(mimeType)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
