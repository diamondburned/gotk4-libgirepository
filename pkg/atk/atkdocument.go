// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_document_get_type()), F: marshalDocumenter},
	})
}

// DocumentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DocumentOverrider interface {
	// CurrentPageNumber retrieves the current page number inside document.
	CurrentPageNumber() int
	// Document gets a gpointer that points to an instance of the DOM. It is up
	// to the caller to check atk_document_get_type to determine how to cast
	// this pointer.
	//
	// Deprecated: Since 2.12. document is already a representation of the
	// document. Use it directly, or one of its children, as an instance of the
	// DOM.
	Document() cgo.Handle
	// DocumentAttributeValue retrieves the value of the given attribute_name
	// inside document.
	DocumentAttributeValue(attributeName string) string
	// DocumentLocale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES
	// locale of the content of this document instance. Individual text
	// substrings or images within this document may have a different locale,
	// see atk_text_get_attributes and atk_image_get_image_locale.
	//
	// Deprecated: Please use atk_object_get_object_locale() instead.
	DocumentLocale() string
	// DocumentType gets a string indicating the document type.
	//
	// Deprecated: Since 2.12. Please use atk_document_get_attributes() to ask
	// for the document type if it applies.
	DocumentType() string
	// PageCount retrieves the total number of pages inside document.
	PageCount() int
	// SetDocumentAttribute sets the value for the given attribute_name inside
	// document.
	SetDocumentAttribute(attributeName, attributeValue string) bool
}

// Document interface should be supported by any object whose content is a
// representation or view of a document. The AtkDocument interface should appear
// on the toplevel container for the document content; however AtkDocument
// instances may be nested (i.e. an AtkDocument may be a descendant of another
// AtkDocument) in those cases where one document contains "embedded content"
// which can reasonably be considered a document in its own right.
type Document struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Document)(nil)
)

// Documenter describes Document's interface methods.
type Documenter interface {
	externglib.Objector

	// AttributeValue retrieves the value of the given attribute_name inside
	// document.
	AttributeValue(attributeName string) string
	// CurrentPageNumber retrieves the current page number inside document.
	CurrentPageNumber() int
	// Document gets a gpointer that points to an instance of the DOM.
	Document() cgo.Handle
	// DocumentType gets a string indicating the document type.
	DocumentType() string
	// Locale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale
	// of the content of this document instance.
	Locale() string
	// PageCount retrieves the total number of pages inside document.
	PageCount() int
	// SetAttributeValue sets the value for the given attribute_name inside
	// document.
	SetAttributeValue(attributeName, attributeValue string) bool
}

var _ Documenter = (*Document)(nil)

func wrapDocument(obj *externglib.Object) *Document {
	return &Document{
		Object: obj,
	}
}

func marshalDocumenter(p uintptr) (interface{}, error) {
	return wrapDocument(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AttributeValue retrieves the value of the given attribute_name inside
// document.
//
// The function takes the following parameters:
//
//    - attributeName: character string representing the name of the attribute
//    whose value is being queried.
//
func (document *Document) AttributeValue(attributeName string) string {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attributeName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_document_get_attribute_value(_arg0, _arg1)
	runtime.KeepAlive(document)
	runtime.KeepAlive(attributeName)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CurrentPageNumber retrieves the current page number inside document.
func (document *Document) CurrentPageNumber() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))

	_cret = C.atk_document_get_current_page_number(_arg0)
	runtime.KeepAlive(document)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Document gets a gpointer that points to an instance of the DOM. It is up to
// the caller to check atk_document_get_type to determine how to cast this
// pointer.
//
// Deprecated: Since 2.12. document is already a representation of the document.
// Use it directly, or one of its children, as an instance of the DOM.
func (document *Document) Document() cgo.Handle {
	var _arg0 *C.AtkDocument // out
	var _cret C.gpointer     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))

	_cret = C.atk_document_get_document(_arg0)
	runtime.KeepAlive(document)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// DocumentType gets a string indicating the document type.
//
// Deprecated: Since 2.12. Please use atk_document_get_attributes() to ask for
// the document type if it applies.
func (document *Document) DocumentType() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))

	_cret = C.atk_document_get_document_type(_arg0)
	runtime.KeepAlive(document)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Locale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale of
// the content of this document instance. Individual text substrings or images
// within this document may have a different locale, see atk_text_get_attributes
// and atk_image_get_image_locale.
//
// Deprecated: Please use atk_object_get_object_locale() instead.
func (document *Document) Locale() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))

	_cret = C.atk_document_get_locale(_arg0)
	runtime.KeepAlive(document)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageCount retrieves the total number of pages inside document.
func (document *Document) PageCount() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))

	_cret = C.atk_document_get_page_count(_arg0)
	runtime.KeepAlive(document)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetAttributeValue sets the value for the given attribute_name inside
// document.
//
// The function takes the following parameters:
//
//    - attributeName: character string representing the name of the attribute
//    whose value is being set.
//    - attributeValue: string value to be associated with attribute_name.
//
func (document *Document) SetAttributeValue(attributeName, attributeValue string) bool {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(document.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attributeName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(attributeValue)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.atk_document_set_attribute_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(document)
	runtime.KeepAlive(attributeName)
	runtime.KeepAlive(attributeValue)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectLoadComplete: 'load-complete' signal is emitted when a pending load of
// a static document has completed. This signal is to be expected by ATK clients
// if and when AtkDocument implementors expose ATK_STATE_BUSY. If the state of
// an AtkObject which implements AtkDocument does not include ATK_STATE_BUSY, it
// should be safe for clients to assume that the AtkDocument's static contents
// are fully loaded into the container. (Dynamic document contents should be
// exposed via other signals.).
func (document *Document) ConnectLoadComplete(f func()) externglib.SignalHandle {
	return document.Connect("load-complete", f)
}

// ConnectLoadStopped: 'load-stopped' signal is emitted when a pending load of
// document contents is cancelled, paused, or otherwise interrupted by the user
// or application logic. It should not however be emitted while waiting for a
// resource (for instance while blocking on a file or network read) unless a
// user-significant timeout has occurred.
func (document *Document) ConnectLoadStopped(f func()) externglib.SignalHandle {
	return document.Connect("load-stopped", f)
}

// ConnectPageChanged: 'page-changed' signal is emitted when the current page of
// a document changes, e.g. pressing page up/down in a document viewer.
func (document *Document) ConnectPageChanged(f func(pageNumber int)) externglib.SignalHandle {
	return document.Connect("page-changed", f)
}

// ConnectReload: 'reload' signal is emitted when the contents of a document is
// refreshed from its source. Once 'reload' has been emitted, a matching
// 'load-complete' or 'load-stopped' signal should follow, which clients may
// await before interrogating ATK for the latest document content.
func (document *Document) ConnectReload(f func()) externglib.SignalHandle {
	return document.Connect("reload", f)
}
