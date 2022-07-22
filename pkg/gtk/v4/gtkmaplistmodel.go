// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gpointer _gotk4_gtk4_MapListModelMapFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GType values.
var (
	GTypeMapListModel = coreglib.Type(C.gtk_map_list_model_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMapListModel, F: marshalMapListModel},
	})
}

// MapListModelMapFunc: user function that is called to map an item of the
// original model to an item expected by the map model.
//
// The returned items must conform to the item type of the model they are used
// with.
type MapListModelMapFunc func(item *coreglib.Object) (object *coreglib.Object)

//export _gotk4_gtk4_MapListModelMapFunc
func _gotk4_gtk4_MapListModelMapFunc(arg1 C.gpointer, arg2 C.gpointer) (cret C.gpointer) {
	var fn MapListModelMapFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(MapListModelMapFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.AssumeOwnership(unsafe.Pointer(arg1))

	object := fn(_item)

	cret = C.gpointer(unsafe.Pointer(object.Native()))
	C.g_object_ref(C.gpointer(object.Native()))

	return cret
}

// MapListModelOverrider contains methods that are overridable.
type MapListModelOverrider interface {
}

// MapListModel: GtkMapListModel maps the items in a list model to different
// items.
//
// GtkMapListModel uses a gtk.MapListModelMapFunc.
//
// Example: Create a list of GtkEventControllers
//
//    static gpointer
//    map_to_controllers (gpointer widget,
//                        gpointer data)
//    {
//      gpointer result = gtk_widget_observe_controllers (widget);
//      g_object_unref (widget);
//      return result;
//    }
//
//    widgets = gtk_widget_observe_children (widget);
//
//    controllers = gtk_map_list_model_new (G_TYPE_LIST_MODEL,
//                                          widgets,
//                                          map_to_controllers,
//                                          NULL, NULL);
//
//    model = gtk_flatten_list_model_new (GTK_TYPE_EVENT_CONTROLLER,
//                                        controllers);
//
//
// GtkMapListModel will attempt to discard the mapped objects as soon as they
// are no longer needed and recreate them if necessary.
type MapListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*MapListModel)(nil)
)

func initClassMapListModel(gclass unsafe.Pointer, goval any) {
}

func wrapMapListModel(obj *coreglib.Object) *MapListModel {
	return &MapListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalMapListModel(p uintptr) (interface{}, error) {
	return wrapMapListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMapListModel creates a new GtkMapListModel for the given arguments.
//
// The function takes the following parameters:
//
//    - model (optional) to map or NULL for none.
//    - mapFunc (optional): map function or NULL to not map items.
//
// The function returns the following values:
//
//    - mapListModel: new GtkMapListModel.
//
func NewMapListModel(model gio.ListModeller, mapFunc MapListModelMapFunc) *MapListModel {
	var _arg1 *C.GListModel            // out
	var _arg2 C.GtkMapListModelMapFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _cret *C.GtkMapListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	if mapFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk4_MapListModelMapFunc)
		_arg3 = C.gpointer(gbox.Assign(mapFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_map_list_model_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(model)
	runtime.KeepAlive(mapFunc)

	var _mapListModel *MapListModel // out

	_mapListModel = wrapMapListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mapListModel
}

// Model gets the model that is currently being mapped or NULL if none.
//
// The function returns the following values:
//
//    - listModel (optional): model that gets mapped.
//
func (self *MapListModel) Model() *gio.ListModel {
	var _arg0 *C.GtkMapListModel // out
	var _cret *C.GListModel      // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_map_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// HasMap checks if a map function is currently set on self.
//
// The function returns the following values:
//
//    - ok: TRUE if a map function is set.
//
func (self *MapListModel) HasMap() bool {
	var _arg0 *C.GtkMapListModel // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_map_list_model_has_map(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMapFunc sets the function used to map items.
//
// The function will be called whenever an item needs to be mapped and must
// return the item to use for the given input item.
//
// Note that GtkMapListModel may call this function multiple times on the same
// item, because it may delete items it doesn't need anymore.
//
// GTK makes no effort to ensure that map_func conforms to the item type of
// self. It assumes that the caller knows what they are doing and the map
// function returns items of the appropriate type.
//
// The function takes the following parameters:
//
//    - mapFunc (optional): map function or NULL to not map items.
//
func (self *MapListModel) SetMapFunc(mapFunc MapListModelMapFunc) {
	var _arg0 *C.GtkMapListModel       // out
	var _arg1 C.GtkMapListModelMapFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if mapFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_MapListModelMapFunc)
		_arg2 = C.gpointer(gbox.Assign(mapFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_map_list_model_set_map_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mapFunc)
}

// SetModel sets the model to be mapped.
//
// GTK makes no effort to ensure that model conforms to the item type expected
// by the map function. It assumes that the caller knows what they are doing and
// have set up an appropriate map function.
//
// The function takes the following parameters:
//
//    - model (optional) to be mapped.
//
func (self *MapListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkMapListModel // out
	var _arg1 *C.GListModel      // out

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_map_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
