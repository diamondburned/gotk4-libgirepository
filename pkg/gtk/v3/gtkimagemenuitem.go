// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkimagemenuitem.go.
var GTypeImageMenuItem = coreglib.Type(C.gtk_image_menu_item_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeImageMenuItem, F: marshalImageMenuItem},
	})
}

// ImageMenuItemOverrider contains methods that are overridable.
type ImageMenuItemOverrider interface {
}

// ImageMenuItem is a menu item which has an icon next to the text label.
//
// This is functionally equivalent to:
//
//      GtkWidget *box = gtk_box_new (GTK_ORIENTATION_HORIZONTAL, 6);
//      GtkWidget *icon = gtk_image_new_from_icon_name ("folder-music-symbolic", GTK_ICON_SIZE_MENU);
//      GtkWidget *label = gtk_accel_label_new ("Music");
//      GtkWidget *menu_item = gtk_menu_item_new ();
//      GtkAccelGroup *accel_group = gtk_accel_group_new ();
//
//      gtk_container_add (GTK_CONTAINER (box), icon);
//
//      gtk_label_set_use_underline (GTK_LABEL (label), TRUE);
//      gtk_label_set_xalign (GTK_LABEL (label), 0.0);
//
//      gtk_widget_add_accelerator (menu_item, "activate", accel_group,
//                                  GDK_KEY_m, GDK_CONTROL_MASK, GTK_ACCEL_VISIBLE);
//      gtk_accel_label_set_accel_widget (GTK_ACCEL_LABEL (label), menu_item);
//
//      gtk_box_pack_end (GTK_BOX (box), label, TRUE, TRUE, 0);
//
//      gtk_container_add (GTK_CONTAINER (menu_item), box);
//
//      gtk_widget_show_all (menu_item);.
type ImageMenuItem struct {
	_ [0]func() // equal guard
	MenuItem
}

var (
	_ Binner            = (*ImageMenuItem)(nil)
	_ coreglib.Objector = (*ImageMenuItem)(nil)
)

func classInitImageMenuItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapImageMenuItem(obj *coreglib.Object) *ImageMenuItem {
	return &ImageMenuItem{
		MenuItem: MenuItem{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalImageMenuItem(p uintptr) (interface{}, error) {
	return wrapImageMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewImageMenuItem creates a new ImageMenuItem with an empty label.
//
// Deprecated: Use gtk_menu_item_new() instead.
//
// The function returns the following values:
//
//    - imageMenuItem: new ImageMenuItem.
//
func NewImageMenuItem() *ImageMenuItem {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("new_ImageMenuItem", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _imageMenuItem *ImageMenuItem // out

	_imageMenuItem = wrapImageMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemFromStock creates a new ImageMenuItem containing the image
// and text from a stock item. Some stock ids have preprocessor macros like
// K_STOCK_OK and K_STOCK_APPLY.
//
// If you want this menu item to have changeable accelerators, then pass in NULL
// for accel_group. Next call gtk_menu_item_set_accel_path() with an appropriate
// path for the menu item, use gtk_stock_lookup() to look up the standard
// accelerator for the stock item, and if one is found, call
// gtk_accel_map_add_entry() to register it.
//
// Deprecated: Use gtk_menu_item_new_with_mnemonic() instead.
//
// The function takes the following parameters:
//
//    - stockId: name of the stock item.
//    - accelGroup (optional) to add the menu items accelerator to, or NULL.
//
// The function returns the following values:
//
//    - imageMenuItem: new ImageMenuItem.
//
func NewImageMenuItemFromStock(stockId string, accelGroup *AccelGroup) *ImageMenuItem {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg0))
	if accelGroup != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(accelGroup).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("new_ImageMenuItem_from_stock", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stockId)
	runtime.KeepAlive(accelGroup)

	var _imageMenuItem *ImageMenuItem // out

	_imageMenuItem = wrapImageMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemWithLabel creates a new ImageMenuItem containing a label.
//
// Deprecated: Use gtk_menu_item_new_with_label() instead.
//
// The function takes the following parameters:
//
//    - label: text of the menu item.
//
// The function returns the following values:
//
//    - imageMenuItem: new ImageMenuItem.
//
func NewImageMenuItemWithLabel(label string) *ImageMenuItem {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg0))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("new_ImageMenuItem_with_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _imageMenuItem *ImageMenuItem // out

	_imageMenuItem = wrapImageMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemWithMnemonic creates a new ImageMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
//
// Deprecated: Use gtk_menu_item_new_with_mnemonic() instead.
//
// The function takes the following parameters:
//
//    - label: text of the menu item, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - imageMenuItem: new ImageMenuItem.
//
func NewImageMenuItemWithMnemonic(label string) *ImageMenuItem {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg0))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("new_ImageMenuItem_with_mnemonic", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _imageMenuItem *ImageMenuItem // out

	_imageMenuItem = wrapImageMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// AlwaysShowImage returns whether the menu item will ignore the
// Settings:gtk-menu-images setting and always show the image, if available.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if the menu item will always show the image.
//
func (imageMenuItem *ImageMenuItem) AlwaysShowImage() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("get_always_show_image", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(imageMenuItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Image gets the widget that is currently set as the image of image_menu_item.
// See gtk_image_menu_item_set_image().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - widget set as image of image_menu_item.
//
func (imageMenuItem *ImageMenuItem) Image() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("get_image", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(imageMenuItem)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// UseStock checks whether the label set in the menuitem is used as a stock id
// to select the stock item for the item.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if the label set in the menuitem is used as a stock id to select
//      the stock item for the item.
//
func (imageMenuItem *ImageMenuItem) UseStock() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("get_use_stock", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(imageMenuItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAccelGroup specifies an accel_group to add the menu items accelerator to
// (this only applies to stock items so a stock item must already be set, make
// sure to call gtk_image_menu_item_set_use_stock() and
// gtk_menu_item_set_label() with a valid stock item first).
//
// If you want this menu item to have changeable accelerators then you shouldnt
// need this (see gtk_image_menu_item_new_from_stock()).
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - accelGroup: AccelGroup.
//
func (imageMenuItem *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(accelGroup).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("set_accel_group", _args[:], nil)

	runtime.KeepAlive(imageMenuItem)
	runtime.KeepAlive(accelGroup)
}

// SetAlwaysShowImage: if TRUE, the menu item will ignore the
// Settings:gtk-menu-images setting and always show the image, if available.
//
// Use this property if the menuitem would be useless or hard to use without the
// image.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - alwaysShow: TRUE if the menuitem should always show the image.
//
func (imageMenuItem *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("set_always_show_image", _args[:], nil)

	runtime.KeepAlive(imageMenuItem)
	runtime.KeepAlive(alwaysShow)
}

// SetImage sets the image of image_menu_item to the given widget. Note that it
// depends on the show-menu-images setting whether the image will be displayed
// or not.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - image (optional): widget to set as the image for the menu item.
//
func (imageMenuItem *ImageMenuItem) SetImage(image Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))
	if image != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("set_image", _args[:], nil)

	runtime.KeepAlive(imageMenuItem)
	runtime.KeepAlive(image)
}

// SetUseStock: if TRUE, the label set in the menuitem is used as a stock id to
// select the stock item for the item.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - useStock: TRUE if the menuitem should use a stock item.
//
func (imageMenuItem *ImageMenuItem) SetUseStock(useStock bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imageMenuItem).Native()))
	if useStock {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ImageMenuItem").InvokeMethod("set_use_stock", _args[:], nil)

	runtime.KeepAlive(imageMenuItem)
	runtime.KeepAlive(useStock)
}
