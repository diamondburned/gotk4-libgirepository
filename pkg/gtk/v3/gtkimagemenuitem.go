// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_menu_item_get_type()), F: marshalImageMenuItem},
	})
}

// ImageMenuItem: a GtkImageMenuItem is a menu item which has an icon next to
// the text label.
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
//      gtk_widget_show_all (menu_item);
type ImageMenuItem interface {
	MenuItem
	Actionable
	Activatable
	Buildable

	// AlwaysShowImage returns whether the menu item will ignore the
	// Settings:gtk-menu-images setting and always show the image, if available.
	AlwaysShowImage() bool
	// Image gets the widget that is currently set as the image of
	// @image_menu_item. See gtk_image_menu_item_set_image().
	Image() Widget
	// UseStock checks whether the label set in the menuitem is used as a stock
	// id to select the stock item for the item.
	UseStock() bool
	// SetAccelGroup specifies an @accel_group to add the menu items accelerator
	// to (this only applies to stock items so a stock item must already be set,
	// make sure to call gtk_image_menu_item_set_use_stock() and
	// gtk_menu_item_set_label() with a valid stock item first).
	//
	// If you want this menu item to have changeable accelerators then you
	// shouldnt need this (see gtk_image_menu_item_new_from_stock()).
	SetAccelGroup(accelGroup AccelGroup)
	// SetAlwaysShowImage: if true, the menu item will ignore the
	// Settings:gtk-menu-images setting and always show the image, if available.
	//
	// Use this property if the menuitem would be useless or hard to use without
	// the image.
	SetAlwaysShowImage(alwaysShow bool)
	// SetImage sets the image of @image_menu_item to the given widget. Note
	// that it depends on the show-menu-images setting whether the image will be
	// displayed or not.
	SetImage(image Widget)
	// SetUseStock: if true, the label set in the menuitem is used as a stock id
	// to select the stock item for the item.
	SetUseStock(useStock bool)
}

// imageMenuItem implements the ImageMenuItem interface.
type imageMenuItem struct {
	MenuItem
	Actionable
	Activatable
	Buildable
}

var _ ImageMenuItem = (*imageMenuItem)(nil)

// WrapImageMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapImageMenuItem(obj *externglib.Object) ImageMenuItem {
	return ImageMenuItem{
		MenuItem:    WrapMenuItem(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalImageMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImageMenuItem(obj), nil
}

// NewImageMenuItem constructs a class ImageMenuItem.
func NewImageMenuItem() ImageMenuItem {
	var _cret C.GtkImageMenuItem

	cret = C.gtk_image_menu_item_new()

	var _imageMenuItem ImageMenuItem

	_imageMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ImageMenuItem)

	return _imageMenuItem
}

// NewImageMenuItemFromStock constructs a class ImageMenuItem.
func NewImageMenuItemFromStock(stockId string, accelGroup AccelGroup) ImageMenuItem {
	var _arg1 *C.gchar
	var _arg2 *C.GtkAccelGroup

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	var _cret C.GtkImageMenuItem

	cret = C.gtk_image_menu_item_new_from_stock(_arg1, _arg2)

	var _imageMenuItem ImageMenuItem

	_imageMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ImageMenuItem)

	return _imageMenuItem
}

// NewImageMenuItemWithLabel constructs a class ImageMenuItem.
func NewImageMenuItemWithLabel(label string) ImageMenuItem {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkImageMenuItem

	cret = C.gtk_image_menu_item_new_with_label(_arg1)

	var _imageMenuItem ImageMenuItem

	_imageMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ImageMenuItem)

	return _imageMenuItem
}

// NewImageMenuItemWithMnemonic constructs a class ImageMenuItem.
func NewImageMenuItemWithMnemonic(label string) ImageMenuItem {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkImageMenuItem

	cret = C.gtk_image_menu_item_new_with_mnemonic(_arg1)

	var _imageMenuItem ImageMenuItem

	_imageMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ImageMenuItem)

	return _imageMenuItem
}

// AlwaysShowImage returns whether the menu item will ignore the
// Settings:gtk-menu-images setting and always show the image, if available.
func (i imageMenuItem) AlwaysShowImage() bool {
	var _arg0 *C.GtkImageMenuItem

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.gtk_image_menu_item_get_always_show_image(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Image gets the widget that is currently set as the image of
// @image_menu_item. See gtk_image_menu_item_set_image().
func (i imageMenuItem) Image() Widget {
	var _arg0 *C.GtkImageMenuItem

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_image_menu_item_get_image(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// UseStock checks whether the label set in the menuitem is used as a stock
// id to select the stock item for the item.
func (i imageMenuItem) UseStock() bool {
	var _arg0 *C.GtkImageMenuItem

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.gtk_image_menu_item_get_use_stock(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetAccelGroup specifies an @accel_group to add the menu items accelerator
// to (this only applies to stock items so a stock item must already be set,
// make sure to call gtk_image_menu_item_set_use_stock() and
// gtk_menu_item_set_label() with a valid stock item first).
//
// If you want this menu item to have changeable accelerators then you
// shouldnt need this (see gtk_image_menu_item_new_from_stock()).
func (i imageMenuItem) SetAccelGroup(accelGroup AccelGroup) {
	var _arg0 *C.GtkImageMenuItem
	var _arg1 *C.GtkAccelGroup

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_image_menu_item_set_accel_group(_arg0, _arg1)
}

// SetAlwaysShowImage: if true, the menu item will ignore the
// Settings:gtk-menu-images setting and always show the image, if available.
//
// Use this property if the menuitem would be useless or hard to use without
// the image.
func (i imageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	var _arg0 *C.GtkImageMenuItem
	var _arg1 C.gboolean

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	if alwaysShow {
		_arg1 = C.gboolean(1)
	}

	C.gtk_image_menu_item_set_always_show_image(_arg0, _arg1)
}

// SetImage sets the image of @image_menu_item to the given widget. Note
// that it depends on the show-menu-images setting whether the image will be
// displayed or not.
func (i imageMenuItem) SetImage(image Widget) {
	var _arg0 *C.GtkImageMenuItem
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(image.Native()))

	C.gtk_image_menu_item_set_image(_arg0, _arg1)
}

// SetUseStock: if true, the label set in the menuitem is used as a stock id
// to select the stock item for the item.
func (i imageMenuItem) SetUseStock(useStock bool) {
	var _arg0 *C.GtkImageMenuItem
	var _arg1 C.gboolean

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	if useStock {
		_arg1 = C.gboolean(1)
	}

	C.gtk_image_menu_item_set_use_stock(_arg0, _arg1)
}