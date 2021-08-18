// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_builder_list_item_factory_get_type()), F: marshalBuilderListItemFactorier},
	})
}

// BuilderListItemFactory: GtkBuilderListItemFactory is a GtkListItemFactory
// that creates widgets by instantiating GtkBuilder UI templates.
//
// The templates must be extending GtkListItem, and typically use GtkExpressions
// to obtain data from the items in the model.
//
// Example:
//
//      <interface>
//        <template class="GtkListItem">
//          <property name="child">
//            <object class="GtkLabel">
//              <property name="xalign">0</property>
//              <binding name="label">
//                <lookup name="name" type="SettingsKey">
//                  <lookup name="item">GtkListItem</lookup>
//                </lookup>
//              </binding>
//            </object>
//          </property>
//        </template>
//      </interface>
type BuilderListItemFactory struct {
	ListItemFactory
}

func wrapBuilderListItemFactory(obj *externglib.Object) *BuilderListItemFactory {
	return &BuilderListItemFactory{
		ListItemFactory: ListItemFactory{
			Object: obj,
		},
	}
}

func marshalBuilderListItemFactorier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBuilderListItemFactory(obj), nil
}

// NewBuilderListItemFactoryFromBytes creates a new GtkBuilderListItemFactory
// that instantiates widgets using bytes as the data to pass to GtkBuilder.
func NewBuilderListItemFactoryFromBytes(scope BuilderScoper, bytes []byte) *BuilderListItemFactory {
	var _arg1 *C.GtkBuilderScope    // out
	var _arg2 *C.GBytes             // out
	var _cret *C.GtkListItemFactory // in

	if scope != nil {
		_arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))
	}
	_arg2 = C.g_bytes_new_with_free_func(
		C.gconstpointer(unsafe.Pointer(&bytes[0])),
		C.gsize(len(bytes)),
		C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
		C.gpointer(gbox.Assign(bytes)),
	)
	defer C.g_bytes_unref(_arg2)

	_cret = C.gtk_builder_list_item_factory_new_from_bytes(_arg1, _arg2)
	runtime.KeepAlive(scope)
	runtime.KeepAlive(bytes)

	var _builderListItemFactory *BuilderListItemFactory // out

	_builderListItemFactory = wrapBuilderListItemFactory(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderListItemFactory
}

// NewBuilderListItemFactoryFromResource creates a new GtkBuilderListItemFactory
// that instantiates widgets using data read from the given resource_path to
// pass to GtkBuilder.
func NewBuilderListItemFactoryFromResource(scope BuilderScoper, resourcePath string) *BuilderListItemFactory {
	var _arg1 *C.GtkBuilderScope    // out
	var _arg2 *C.char               // out
	var _cret *C.GtkListItemFactory // in

	if scope != nil {
		_arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_builder_list_item_factory_new_from_resource(_arg1, _arg2)
	runtime.KeepAlive(scope)
	runtime.KeepAlive(resourcePath)

	var _builderListItemFactory *BuilderListItemFactory // out

	_builderListItemFactory = wrapBuilderListItemFactory(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderListItemFactory
}

// Bytes gets the data used as the GtkBuilder UI template for constructing
// listitems.
func (self *BuilderListItemFactory) Bytes() []byte {
	var _arg0 *C.GtkBuilderListItemFactory // out
	var _cret *C.GBytes                    // in

	_arg0 = (*C.GtkBuilderListItemFactory)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_builder_list_item_factory_get_bytes(_arg0)
	runtime.KeepAlive(self)

	var _bytes []byte // out

	_bytes = *(*[]byte)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_bytes_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(&_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// Resource: if the data references a resource, gets the path of that resource.
func (self *BuilderListItemFactory) Resource() string {
	var _arg0 *C.GtkBuilderListItemFactory // out
	var _cret *C.char                      // in

	_arg0 = (*C.GtkBuilderListItemFactory)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_builder_list_item_factory_get_resource(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Scope gets the scope used when constructing listitems.
func (self *BuilderListItemFactory) Scope() BuilderScoper {
	var _arg0 *C.GtkBuilderListItemFactory // out
	var _cret *C.GtkBuilderScope           // in

	_arg0 = (*C.GtkBuilderListItemFactory)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_builder_list_item_factory_get_scope(_arg0)
	runtime.KeepAlive(self)

	var _builderScope BuilderScoper // out

	if _cret != nil {
		_builderScope = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(BuilderScoper)
	}

	return _builderScope
}
