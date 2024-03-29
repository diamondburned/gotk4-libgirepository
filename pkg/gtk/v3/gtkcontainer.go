// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_Container_ConnectSetFocusChild(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_Container_ConnectRemove(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_Container_ConnectCheckResize(gpointer, guintptr);
// extern void _gotk4_gtk3_Container_ConnectAdd(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeContainer = coreglib.Type(girepository.MustFind("Gtk", "Container").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContainer, F: marshalContainer},
	})
}

// ContainerOverrides contains methods that are overridable.
type ContainerOverrides struct {
}

func defaultContainerOverrides(v *Container) ContainerOverrides {
	return ContainerOverrides{}
}

// Container: GTK+ user interface is constructed by nesting widgets inside
// widgets. Container widgets are the inner nodes in the resulting tree of
// widgets: they contain other widgets. So, for example, you might have a Window
// containing a Frame containing a Label. If you wanted an image instead of a
// textual label inside the frame, you might replace the Label widget with a
// Image widget.
//
// There are two major kinds of container widgets in GTK+. Both are subclasses
// of the abstract GtkContainer base class.
//
// The first type of container widget has a single child widget and derives from
// Bin. These containers are decorators, which add some kind of functionality to
// the child. For example, a Button makes its child into a clickable button; a
// Frame draws a frame around its child and a Window places its child widget
// inside a top-level window.
//
// The second type of container can have more than one child; its purpose is to
// manage layout. This means that these containers assign sizes and positions to
// their children. For example, a HBox arranges its children in a horizontal
// row, and a Grid arranges the widgets it contains in a two-dimensional grid.
//
// For implementations of Container the virtual method ContainerClass.forall()
// is always required, since it's used for drawing and other internal operations
// on the children. If the Container implementation expect to have non internal
// children it's needed to implement both ContainerClass.add() and
// ContainerClass.remove(). If the GtkContainer implementation has internal
// children, they should be added with gtk_widget_set_parent() on init() and
// removed with gtk_widget_unparent() in the WidgetClass.destroy()
// implementation. See more about implementing custom widgets at
// https://wiki.gnome.org/HowDoI/CustomWidgets
//
//
// Height for width geometry management
//
// GTK+ uses a height-for-width (and width-for-height) geometry management
// system. Height-for-width means that a widget can change how much vertical
// space it needs, depending on the amount of horizontal space that it is given
// (and similar for width-for-height).
//
// There are some things to keep in mind when implementing container widgets
// that make use of GTK+’s height for width geometry management system. First,
// it’s important to note that a container must prioritize one of its
// dimensions, that is to say that a widget or container can only have a
// SizeRequestMode that is GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH or
// GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT. However, every widget and container must
// be able to respond to the APIs for both dimensions, i.e. even if a widget has
// a request mode that is height-for-width, it is possible that its parent will
// request its sizes using the width-for-height APIs.
//
// To ensure that everything works properly, here are some guidelines to follow
// when implementing height-for-width (or width-for-height) containers.
//
// Each request mode involves 2 virtual methods. Height-for-width apis run
// through gtk_widget_get_preferred_width() and then through
// gtk_widget_get_preferred_height_for_width(). When handling requests in the
// opposite SizeRequestMode it is important that every widget request at least
// enough space to display all of its content at all times.
//
// When gtk_widget_get_preferred_height() is called on a container that is
// height-for-width, the container must return the height for its minimum width.
// This is easily achieved by simply calling the reverse apis implemented for
// itself as follows:
//
//    static void
//    foo_container_get_preferred_width_for_height (GtkWidget *widget,
//                                                  gint for_height,
//                                                  gint *min_width,
//                                                  gint *nat_width)
//    {
//       if (i_am_in_height_for_width_mode)
//         {
//           GTK_WIDGET_GET_CLASS (widget)->get_preferred_width (widget,
//                                                               min_width,
//                                                               nat_width);
//         }
//       else
//         {
//           ... execute the real width-for-height request here based on
//           the required width of the children collectively if the
//           container were to be allocated the said height ...
//         }
//    }
//
// Height for width requests are generally implemented in terms of a virtual
// allocation of widgets in the input orientation. Assuming an height-for-width
// request mode, a container would implement the
// get_preferred_height_for_width() virtual function by first calling
// gtk_widget_get_preferred_width() for each of its children.
//
// For each potential group of children that are lined up horizontally, the
// values returned by gtk_widget_get_preferred_width() should be collected in an
// array of RequestedSize structures. Any child spacing should be removed from
// the input for_width and then the collective size should be allocated using
// the gtk_distribute_natural_allocation() convenience function.
//
// The container will then move on to request the preferred height for each
// child by using gtk_widget_get_preferred_height_for_width() and using the
// sizes stored in the RequestedSize array.
//
// To allocate a height-for-width container, it’s again important to consider
// that a container must prioritize one dimension over the other. So if a
// container is a height-for-width container it must first allocate all widgets
// horizontally using a RequestedSize array and
// gtk_distribute_natural_allocation() and then add any extra space (if and
// where appropriate) for the widget to expand.
//
// After adding all the expand space, the container assumes it was allocated
// sufficient height to fit all of its content. At this time, the container must
// use the total horizontal sizes of each widget to request the height-for-width
// of each of its children and store the requests in a RequestedSize array for
// any widgets that stack vertically (for tabular containers this can be
// generalized into the heights and widths of rows and columns). The vertical
// space must then again be distributed using
// gtk_distribute_natural_allocation() while this time considering the allocated
// height of the widget minus any vertical spacing that the container adds. Then
// vertical expand space should be added where appropriate and available and the
// container should go on to actually allocating the child widgets.
//
// See [GtkWidget’s geometry management section][geometry-management] to learn
// more about implementing height-for-width geometry management for widgets.
//
//
// Child properties
//
// GtkContainer introduces child properties. These are object properties that
// are not specific to either the container or the contained widget, but rather
// to their relation. Typical examples of child properties are the position or
// pack-type of a widget which is contained in a Box.
//
// Use gtk_container_class_install_child_property() to install child properties
// for a container class and gtk_container_class_find_child_property() or
// gtk_container_class_list_child_properties() to get information about existing
// child properties.
//
// To set the value of a child property, use gtk_container_child_set_property(),
// gtk_container_child_set() or gtk_container_child_set_valist(). To obtain the
// value of a child property, use gtk_container_child_get_property(),
// gtk_container_child_get() or gtk_container_child_get_valist(). To emit
// notification about child property changes, use gtk_widget_child_notify().
//
//
// GtkContainer as GtkBuildable
//
// The GtkContainer implementation of the GtkBuildable interface supports a
// <packing> element for children, which can contain multiple <property>
// elements that specify child properties for the child.
//
// Since 2.16, child properties can also be marked as translatable using the
// same “translatable”, “comments” and “context” attributes that are used for
// regular properties.
//
// Since 3.16, containers can have a <focus-chain> element containing multiple
// <widget> elements, one for each child that should be added to the focus
// chain. The ”name” attribute gives the id of the widget.
//
// An example of these properties in UI definitions:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkEntry" id="entry1"/>
//        <packing>
//          <property name="pack-type">start</property>
//        </packing>
//      </child>
//      <child>
//        <object class="GtkEntry" id="entry2"/>
//      </child>
//      <focus-chain>
//        <widget name="entry1"/>
//        <widget name="entry2"/>
//      </focus-chain>
//    </object>.
type Container struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Container)(nil)
)

// Containerer describes types inherited from class Container.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Containerer interface {
	coreglib.Objector
	baseContainer() *Container
}

var _ Containerer = (*Container)(nil)

func init() {
	coreglib.RegisterClassInfo[*Container, *ContainerClass, ContainerOverrides](
		GTypeContainer,
		initContainerClass,
		wrapContainer,
		defaultContainerOverrides,
	)
}

func initContainerClass(gclass unsafe.Pointer, overrides ContainerOverrides, classInitFunc func(*ContainerClass)) {
	if classInitFunc != nil {
		class := (*ContainerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapContainer(obj *coreglib.Object) *Container {
	return &Container{
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
	}
}

func marshalContainer(p uintptr) (interface{}, error) {
	return wrapContainer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Container) baseContainer() *Container {
	return v
}

// BaseContainer returns the underlying base object.
func BaseContainer(obj Containerer) *Container {
	return obj.baseContainer()
}

func (v *Container) ConnectAdd(f func(object Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "add", false, unsafe.Pointer(C._gotk4_gtk3_Container_ConnectAdd), f)
}

func (v *Container) ConnectCheckResize(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "check-resize", false, unsafe.Pointer(C._gotk4_gtk3_Container_ConnectCheckResize), f)
}

func (v *Container) ConnectRemove(f func(object Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "remove", false, unsafe.Pointer(C._gotk4_gtk3_Container_ConnectRemove), f)
}

func (v *Container) ConnectSetFocusChild(f func(object Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "set-focus-child", false, unsafe.Pointer(C._gotk4_gtk3_Container_ConnectSetFocusChild), f)
}

// ContainerClass: base class for containers.
//
// An instance of this type is always passed by reference.
type ContainerClass struct {
	*containerClass
}

// containerClass is the struct that's finalized.
type containerClass struct {
	native unsafe.Pointer
}

var GIRInfoContainerClass = girepository.MustFind("Gtk", "ContainerClass")
