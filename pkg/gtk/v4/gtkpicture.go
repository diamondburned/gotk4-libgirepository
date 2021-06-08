// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_picture_get_type()), F: marshalPicture},
	})
}

// Picture: the `GtkPicture` widget displays a `GdkPaintable`.
//
// !An example GtkPicture (picture.png)
//
// Many convenience functions are provided to make pictures simple to use. For
// example, if you want to load an image from a file, and then display it,
// there’s a convenience function to do this:
//
// “`c GtkWidget *widget = gtk_picture_new_for_filename ("myfile.png"); “`
//
// If the file isn’t loaded successfully, the picture will contain a “broken
// image” icon similar to that used in many web browsers. If you want to handle
// errors in loading the file yourself, for example by displaying an error
// message, then load the image with [ctor@Gdk.Texture.new_from_file], then
// create the `GtkPicture` with [ctor@Gtk.Picture.new_for_paintable].
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of `GResource` for details. In
// this case, [ctor@Gtk.Picture.new_for_resource] and
// [method@Gtk.Picture.set_resource] should be used.
//
// `GtkPicture` displays an image at its natural size. See [class@Gtk.Image] if
// you want to display a fixed-size image, such as an icon.
//
//
// Sizing the paintable
//
// You can influence how the paintable is displayed inside the `GtkPicture`. By
// turning off [property@Gtk.Picture:keep-aspect-ratio] you can allow the
// paintable to get stretched. [property@Gtk.Picture:can-shrink] can be unset to
// make sure that paintables are never made smaller than their ideal size - but
// be careful if you do not know the size of the paintable in use (like when
// displaying user-loaded images). This can easily cause the picture to grow
// larger than the screen. And [property@GtkWidget:halign] and
// [property@GtkWidget:valign] can be used to make sure the paintable doesn't
// fill all available space but is instead displayed at its original size.
//
//
// CSS nodes
//
// `GtkPicture` has a single CSS node with the name `picture`.
//
//
// Accessibility
//
// `GtkPicture` uses the `GTK_ACCESSIBLE_ROLE_IMG` role.
type Picture interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// AlternativeText gets the alternative textual description of the picture.
	//
	// The returned string will be nil if the picture cannot be described
	// textually.
	AlternativeText() string
	// CanShrink returns whether the `GtkPicture` respects its contents size.
	CanShrink() bool
	// File gets the `GFile` currently displayed if @self is displaying a file.
	//
	// If @self is not displaying a file, for example when
	// [method@Gtk.Picture.set_paintable] was used, then nil is returned.
	File() gio.File
	// KeepAspectRatio returns whether the `GtkPicture` preserves its contents
	// aspect ratio.
	KeepAspectRatio() bool
	// Paintable gets the `GdkPaintable` being displayed by the `GtkPicture`.
	Paintable() gdk.Paintable
	// SetAlternativeText sets an alternative textual description for the
	// picture contents.
	//
	// It is equivalent to the "alt" attribute for images on websites.
	//
	// This text will be made available to accessibility tools.
	//
	// If the picture cannot be described textually, set this property to nil.
	SetAlternativeText(alternativeText string)
	// SetCanShrink: if set to true, the @self can be made smaller than its
	// contents.
	//
	// The contents will then be scaled down when rendering.
	//
	// If you want to still force a minimum size manually, consider using
	// [method@Gtk.Widget.set_size_request].
	//
	// Also of note is that a similar function for growing does not exist
	// because the grow behavior can be controlled via
	// [method@Gtk.Widget.set_halign] and [method@Gtk.Widget.set_valign].
	SetCanShrink(canShrink bool)
	// SetFile makes @self load and display @file.
	//
	// See [ctor@Gtk.Picture.new_for_file] for details.
	SetFile(file gio.File)
	// SetFilename makes @self load and display the given @filename.
	//
	// This is a utility function that calls [method@Gtk.Picture.set_file].
	SetFilename(filename string)
	// SetKeepAspectRatio: if set to true, the @self will render its contents
	// according to their aspect ratio.
	//
	// That means that empty space may show up at the top/bottom or left/right
	// of @self.
	//
	// If set to false or if the contents provide no aspect ratio, the contents
	// will be stretched over the picture's whole area.
	SetKeepAspectRatio(keepAspectRatio bool)
	// SetPaintable makes @self display the given @paintable.
	//
	// If @paintable is nil, nothing will be displayed.
	//
	// See [ctor@Gtk.Picture.new_for_paintable] for details.
	SetPaintable(paintable gdk.Paintable)
	// SetPixbuf sets a `GtkPicture` to show a `GdkPixbuf`.
	//
	// See [ctor@Gtk.Picture.new_for_pixbuf] for details.
	//
	// This is a utility function that calls [method@Gtk.Picture.set_paintable].
	SetPixbuf(pixbuf gdkpixbuf.Pixbuf)
	// SetResource makes @self load and display the resource at the given
	// @resource_path.
	//
	// This is a utility function that calls [method@Gtk.Picture.set_file].
	SetResource(resourcePath string)
}

// picture implements the Picture interface.
type picture struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Picture = (*picture)(nil)

// WrapPicture wraps a GObject to the right type. It is
// primarily used internally.
func WrapPicture(obj *externglib.Object) Picture {
	return Picture{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalPicture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPicture(obj), nil
}

// NewPicture constructs a class Picture.
func NewPicture() Picture {
	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new()

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// NewPictureForFile constructs a class Picture.
func NewPictureForFile(file gio.File) Picture {
	var arg1 *C.GFile

	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new_for_file(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// NewPictureForFilename constructs a class Picture.
func NewPictureForFilename(filename string) Picture {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new_for_filename(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// NewPictureForPaintable constructs a class Picture.
func NewPictureForPaintable(paintable gdk.Paintable) Picture {
	var arg1 *C.GdkPaintable

	arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new_for_paintable(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// NewPictureForPixbuf constructs a class Picture.
func NewPictureForPixbuf(pixbuf gdkpixbuf.Pixbuf) Picture {
	var arg1 *C.GdkPixbuf

	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new_for_pixbuf(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// NewPictureForResource constructs a class Picture.
func NewPictureForResource(resourcePath string) Picture {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkPicture
	var goret Picture

	cret = C.gtk_picture_new_for_resource(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Picture)

	return goret
}

// AlternativeText gets the alternative textual description of the picture.
//
// The returned string will be nil if the picture cannot be described
// textually.
func (s picture) AlternativeText() string {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret *C.char
	var goret string

	cret = C.gtk_picture_get_alternative_text(arg0)

	goret = C.GoString(cret)

	return goret
}

// CanShrink returns whether the `GtkPicture` respects its contents size.
func (s picture) CanShrink() bool {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_picture_get_can_shrink(arg0)

	if cret {
		goret = true
	}

	return goret
}

// File gets the `GFile` currently displayed if @self is displaying a file.
//
// If @self is not displaying a file, for example when
// [method@Gtk.Picture.set_paintable] was used, then nil is returned.
func (s picture) File() gio.File {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret *C.GFile
	var goret gio.File

	cret = C.gtk_picture_get_file(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gio.File)

	return goret
}

// KeepAspectRatio returns whether the `GtkPicture` preserves its contents
// aspect ratio.
func (s picture) KeepAspectRatio() bool {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_picture_get_keep_aspect_ratio(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Paintable gets the `GdkPaintable` being displayed by the `GtkPicture`.
func (s picture) Paintable() gdk.Paintable {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret *C.GdkPaintable
	var goret gdk.Paintable

	cret = C.gtk_picture_get_paintable(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Paintable)

	return goret
}

// SetAlternativeText sets an alternative textual description for the
// picture contents.
//
// It is equivalent to the "alt" attribute for images on websites.
//
// This text will be made available to accessibility tools.
//
// If the picture cannot be described textually, set this property to nil.
func (s picture) SetAlternativeText(alternativeText string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(alternativeText))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_alternative_text(arg0, arg1)
}

// SetCanShrink: if set to true, the @self can be made smaller than its
// contents.
//
// The contents will then be scaled down when rendering.
//
// If you want to still force a minimum size manually, consider using
// [method@Gtk.Widget.set_size_request].
//
// Also of note is that a similar function for growing does not exist
// because the grow behavior can be controlled via
// [method@Gtk.Widget.set_halign] and [method@Gtk.Widget.set_valign].
func (s picture) SetCanShrink(canShrink bool) {
	var arg0 *C.GtkPicture
	var arg1 C.gboolean

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if canShrink {
		arg1 = C.gboolean(1)
	}

	C.gtk_picture_set_can_shrink(arg0, arg1)
}

// SetFile makes @self load and display @file.
//
// See [ctor@Gtk.Picture.new_for_file] for details.
func (s picture) SetFile(file gio.File) {
	var arg0 *C.GtkPicture
	var arg1 *C.GFile

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_picture_set_file(arg0, arg1)
}

// SetFilename makes @self load and display the given @filename.
//
// This is a utility function that calls [method@Gtk.Picture.set_file].
func (s picture) SetFilename(filename string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_filename(arg0, arg1)
}

// SetKeepAspectRatio: if set to true, the @self will render its contents
// according to their aspect ratio.
//
// That means that empty space may show up at the top/bottom or left/right
// of @self.
//
// If set to false or if the contents provide no aspect ratio, the contents
// will be stretched over the picture's whole area.
func (s picture) SetKeepAspectRatio(keepAspectRatio bool) {
	var arg0 *C.GtkPicture
	var arg1 C.gboolean

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if keepAspectRatio {
		arg1 = C.gboolean(1)
	}

	C.gtk_picture_set_keep_aspect_ratio(arg0, arg1)
}

// SetPaintable makes @self display the given @paintable.
//
// If @paintable is nil, nothing will be displayed.
//
// See [ctor@Gtk.Picture.new_for_paintable] for details.
func (s picture) SetPaintable(paintable gdk.Paintable) {
	var arg0 *C.GtkPicture
	var arg1 *C.GdkPaintable

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	C.gtk_picture_set_paintable(arg0, arg1)
}

// SetPixbuf sets a `GtkPicture` to show a `GdkPixbuf`.
//
// See [ctor@Gtk.Picture.new_for_pixbuf] for details.
//
// This is a utility function that calls [method@Gtk.Picture.set_paintable].
func (s picture) SetPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkPicture
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_picture_set_pixbuf(arg0, arg1)
}

// SetResource makes @self load and display the resource at the given
// @resource_path.
//
// This is a utility function that calls [method@Gtk.Picture.set_file].
func (s picture) SetResource(resourcePath string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_resource(arg0, arg1)
}
