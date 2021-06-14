// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_activatable_get_type()), F: marshalActivatable},
	})
}

// ActivatableOverrider contains methods that are overridable. This
// interface is a subset of the interface Activatable.
type ActivatableOverrider interface {
	// SyncActionProperties: this is called to update the activatable
	// completely, this is called internally when the Activatable:related-action
	// property is set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	SyncActionProperties(action Action)

	Update(action Action, propertyName string)
}

// Activatable: activatable widgets can be connected to a Action and reflects
// the state of its action. A Activatable can also provide feedback through its
// action, as they are responsible for activating their related actions.
//
//
// Implementing GtkActivatable
//
// When extending a class that is already Activatable; it is only necessary to
// implement the Activatable->sync_action_properties() and Activatable->update()
// methods and chain up to the parent implementation, however when introducing a
// new Activatable class; the Activatable:related-action and
// Activatable:use-action-appearance properties need to be handled by the
// implementor. Handling these properties is mostly a matter of installing the
// action pointer and boolean flag on your instance, and calling
// gtk_activatable_do_set_related_action() and
// gtk_activatable_sync_action_properties() at the appropriate times.
//
// A class fragment implementing Activatable
//
//
//    enum {
//    ...
//
//    PROP_ACTIVATABLE_RELATED_ACTION,
//    PROP_ACTIVATABLE_USE_ACTION_APPEARANCE
//    }
//
//    struct _FooBarPrivate
//    {
//
//      ...
//
//      GtkAction      *action;
//      gboolean        use_action_appearance;
//    };
//
//    ...
//
//    static void foo_bar_activatable_interface_init         (GtkActivatableIface  *iface);
//    static void foo_bar_activatable_update                 (GtkActivatable       *activatable,
//    						           GtkAction            *action,
//    						           const gchar          *property_name);
//    static void foo_bar_activatable_sync_action_properties (GtkActivatable       *activatable,
//    						           GtkAction            *action);
//    ...
//
//
//    static void
//    foo_bar_class_init (FooBarClass *klass)
//    {
//
//      ...
//
//      g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_RELATED_ACTION, "related-action");
//      g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_USE_ACTION_APPEARANCE, "use-action-appearance");
//
//      ...
//    }
//
//
//    static void
//    foo_bar_activatable_interface_init (GtkActivatableIface  *iface)
//    {
//      iface->update = foo_bar_activatable_update;
//      iface->sync_action_properties = foo_bar_activatable_sync_action_properties;
//    }
//
//    ... Break the reference using gtk_activatable_do_set_related_action()...
//
//    static void
//    foo_bar_dispose (GObject *object)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      ...
//
//      if (priv->action)
//        {
//          gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), NULL);
//          priv->action = NULL;
//        }
//      G_OBJECT_CLASS (foo_bar_parent_class)->dispose (object);
//    }
//
//    ... Handle the “related-action” and “use-action-appearance” properties ...
//
//    static void
//    foo_bar_set_property (GObject         *object,
//                          guint            prop_id,
//                          const GValue    *value,
//                          GParamSpec      *pspec)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      switch (prop_id)
//        {
//
//          ...
//
//        case PROP_ACTIVATABLE_RELATED_ACTION:
//          foo_bar_set_related_action (bar, g_value_get_object (value));
//          break;
//        case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
//          foo_bar_set_use_action_appearance (bar, g_value_get_boolean (value));
//          break;
//        default:
//          G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
//          break;
//        }
//    }
//
//    static void
//    foo_bar_get_property (GObject         *object,
//                             guint            prop_id,
//                             GValue          *value,
//                             GParamSpec      *pspec)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      switch (prop_id)
//        {
//
//          ...
//
//        case PROP_ACTIVATABLE_RELATED_ACTION:
//          g_value_set_object (value, priv->action);
//          break;
//        case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
//          g_value_set_boolean (value, priv->use_action_appearance);
//          break;
//        default:
//          G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
//          break;
//        }
//    }
//
//
//    static void
//    foo_bar_set_use_action_appearance (FooBar   *bar,
//    				   gboolean  use_appearance)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      if (priv->use_action_appearance != use_appearance)
//        {
//          priv->use_action_appearance = use_appearance;
//
//          gtk_activatable_sync_action_properties (GTK_ACTIVATABLE (bar), priv->action);
//        }
//    }
//
//    ... call gtk_activatable_do_set_related_action() and then assign the action pointer,
//    no need to reference the action here since gtk_activatable_do_set_related_action() already
//    holds a reference here for you...
//    static void
//    foo_bar_set_related_action (FooBar    *bar,
//    			    GtkAction *action)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      if (priv->action == action)
//        return;
//
//      gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), action);
//
//      priv->action = action;
//    }
//
//    ... Selectively reset and update activatable depending on the use-action-appearance property ...
//    static void
//    gtk_button_activatable_sync_action_properties (GtkActivatable       *activatable,
//    		                                  GtkAction            *action)
//    {
//      GtkButtonPrivate *priv = GTK_BUTTON_GET_PRIVATE (activatable);
//
//      if (!action)
//        return;
//
//      if (gtk_action_is_visible (action))
//        gtk_widget_show (GTK_WIDGET (activatable));
//      else
//        gtk_widget_hide (GTK_WIDGET (activatable));
//
//      gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
//      ...
//
//      if (priv->use_action_appearance)
//        {
//          if (gtk_action_get_stock_id (action))
//    	foo_bar_set_stock (button, gtk_action_get_stock_id (action));
//          else if (gtk_action_get_label (action))
//    	foo_bar_set_label (button, gtk_action_get_label (action));
//
//          ...
//
//        }
//    }
//
//    static void
//    foo_bar_activatable_update (GtkActivatable       *activatable,
//    			       GtkAction            *action,
//    			       const gchar          *property_name)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (activatable);
//
//      if (strcmp (property_name, "visible") == 0)
//        {
//          if (gtk_action_is_visible (action))
//    	gtk_widget_show (GTK_WIDGET (activatable));
//          else
//    	gtk_widget_hide (GTK_WIDGET (activatable));
//        }
//      else if (strcmp (property_name, "sensitive") == 0)
//        gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
//      ...
//
//      if (!priv->use_action_appearance)
//        return;
//
//      if (strcmp (property_name, "stock-id") == 0)
//        foo_bar_set_stock (button, gtk_action_get_stock_id (action));
//      else if (strcmp (property_name, "label") == 0)
//        foo_bar_set_label (button, gtk_action_get_label (action));
//
//      ...
//    }
type Activatable interface {
	gextras.Objector
	ActivatableOverrider

	// DoSetRelatedAction: this is a utility function for Activatable
	// implementors.
	//
	// When implementing Activatable you must call this when handling changes of
	// the Activatable:related-action, and you must also use this to break
	// references in #GObject->dispose().
	//
	// This function adds a reference to the currently set related action for
	// you, it also makes sure the Activatable->update() method is called when
	// the related Action properties change and registers to the action’s proxy
	// list.
	//
	// > Be careful to call this before setting the local > copy of the Action
	// property, since this function uses > gtk_activatable_get_related_action()
	// to retrieve the > previous action.
	DoSetRelatedAction(action Action)
	// UseActionAppearance gets whether this activatable should reset its layout
	// and appearance when setting the related action or when the action changes
	// appearance.
	UseActionAppearance() bool
	// SetRelatedAction sets the related action on the @activatable object.
	//
	// > Activatable implementors need to handle the Activatable:related-action
	// > property and call gtk_activatable_do_set_related_action() when it
	// changes.
	SetRelatedAction(action Action)
	// SetUseActionAppearance sets whether this activatable should reset its
	// layout and appearance when setting the related action or when the action
	// changes appearance
	//
	// > Activatable implementors need to handle the >
	// Activatable:use-action-appearance property and call >
	// gtk_activatable_sync_action_properties() to update @activatable > if
	// needed.
	SetUseActionAppearance(useAppearance bool)
}

// activatable implements the Activatable interface.
type activatable struct {
	gextras.Objector
}

var _ Activatable = (*activatable)(nil)

// WrapActivatable wraps a GObject to a type that implements interface
// Activatable. It is primarily used internally.
func WrapActivatable(obj *externglib.Object) Activatable {
	return Activatable{
		Objector: obj,
	}
}

func marshalActivatable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActivatable(obj), nil
}

// DoSetRelatedAction: this is a utility function for Activatable
// implementors.
//
// When implementing Activatable you must call this when handling changes of
// the Activatable:related-action, and you must also use this to break
// references in #GObject->dispose().
//
// This function adds a reference to the currently set related action for
// you, it also makes sure the Activatable->update() method is called when
// the related Action properties change and registers to the action’s proxy
// list.
//
// > Be careful to call this before setting the local > copy of the Action
// property, since this function uses > gtk_activatable_get_related_action()
// to retrieve the > previous action.
func (a activatable) DoSetRelatedAction(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_do_set_related_action(_arg0, _arg1)
}

// UseActionAppearance gets whether this activatable should reset its layout
// and appearance when setting the related action or when the action changes
// appearance.
func (a activatable) UseActionAppearance() bool {
	var _arg0 *C.GtkActivatable // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_activatable_get_use_action_appearance(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRelatedAction sets the related action on the @activatable object.
//
// > Activatable implementors need to handle the Activatable:related-action
// > property and call gtk_activatable_do_set_related_action() when it
// changes.
func (a activatable) SetRelatedAction(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_set_related_action(_arg0, _arg1)
}

// SetUseActionAppearance sets whether this activatable should reset its
// layout and appearance when setting the related action or when the action
// changes appearance
//
// > Activatable implementors need to handle the >
// Activatable:use-action-appearance property and call >
// gtk_activatable_sync_action_properties() to update @activatable > if
// needed.
func (a activatable) SetUseActionAppearance(useAppearance bool) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	if useAppearance {
		_arg1 = C.TRUE
	}

	C.gtk_activatable_set_use_action_appearance(_arg0, _arg1)
}

// SyncActionProperties: this is called to update the activatable
// completely, this is called internally when the Activatable:related-action
// property is set or unset and by the implementing class when
// Activatable:use-action-appearance changes.
func (a activatable) SyncActionProperties(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_sync_action_properties(_arg0, _arg1)
}
