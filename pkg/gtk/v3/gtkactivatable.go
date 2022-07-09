// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ActivatableIface_sync_action_properties(void*, void*);
// extern void _gotk4_gtk3_ActivatableIface_update(void*, void*, void*);
import "C"

// GTypeActivatable returns the GType for the type Activatable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeActivatable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Activatable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalActivatable)
	return gtype
}

// ActivatableOverrider contains methods that are overridable.
type ActivatableOverrider interface {
	// SyncActionProperties: this is called to update the activatable
	// completely, this is called internally when the Activatable:related-action
	// property is set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	//
	// The function takes the following parameters:
	//
	//    - action (optional): related Action or NULL.
	//
	SyncActionProperties(action *Action)
	// The function takes the following parameters:
	//
	//    - action
	//    - propertyName
	//
	Update(action *Action, propertyName string)
}

// Activatable widgets can be connected to a Action and reflects the state of
// its action. A Activatable can also provide feedback through its action, as
// they are responsible for activating their related actions.
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
//    }.
//
// Activatable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Activatable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Activatable)(nil)
)

// Activatabler describes Activatable's interface methods.
type Activatabler interface {
	coreglib.Objector

	// DoSetRelatedAction: this is a utility function for Activatable
	// implementors.
	DoSetRelatedAction(action *Action)
	// RelatedAction gets the related Action for activatable.
	RelatedAction() *Action
	// UseActionAppearance gets whether this activatable should reset its layout
	// and appearance when setting the related action or when the action changes
	// appearance.
	UseActionAppearance() bool
	// SetRelatedAction sets the related action on the activatable object.
	SetRelatedAction(action *Action)
	// SetUseActionAppearance sets whether this activatable should reset its
	// layout and appearance when setting the related action or when the action
	// changes appearance > Activatable implementors need to handle the >
	// Activatable:use-action-appearance property and call >
	// gtk_activatable_sync_action_properties() to update activatable > if
	// needed.
	SetUseActionAppearance(useAppearance bool)
	// SyncActionProperties: this is called to update the activatable
	// completely, this is called internally when the Activatable:related-action
	// property is set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	SyncActionProperties(action *Action)
}

var _ Activatabler = (*Activatable)(nil)

func ifaceInitActivatabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "ActivatableIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("sync_action_properties"))) = unsafe.Pointer(C._gotk4_gtk3_ActivatableIface_sync_action_properties)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("update"))) = unsafe.Pointer(C._gotk4_gtk3_ActivatableIface_update)
}

//export _gotk4_gtk3_ActivatableIface_sync_action_properties
func _gotk4_gtk3_ActivatableIface_sync_action_properties(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActivatableOverrider)

	var _action *Action // out

	if arg1 != nil {
		_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))
	}

	iface.SyncActionProperties(_action)
}

//export _gotk4_gtk3_ActivatableIface_update
func _gotk4_gtk3_ActivatableIface_update(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActivatableOverrider)

	var _action *Action      // out
	var _propertyName string // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))
	_propertyName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	iface.Update(_action, _propertyName)
}

func wrapActivatable(obj *coreglib.Object) *Activatable {
	return &Activatable{
		Object: obj,
	}
}

func marshalActivatable(p uintptr) (interface{}, error) {
	return wrapActivatable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DoSetRelatedAction: this is a utility function for Activatable implementors.
//
// When implementing Activatable you must call this when handling changes of the
// Activatable:related-action, and you must also use this to break references in
// #GObject->dispose().
//
// This function adds a reference to the currently set related action for you,
// it also makes sure the Activatable->update() method is called when the
// related Action properties change and registers to the action’s proxy list.
//
// > Be careful to call this before setting the local > copy of the Action
// property, since this function uses > gtk_activatable_get_related_action() to
// retrieve the > previous action.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action to set.
//
func (activatable *Activatable) DoSetRelatedAction(action *Action) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_info := girepository.MustFind("Gtk", "Activatable")
	_info.InvokeIfaceMethod("do_set_related_action", _args[:], nil)

	runtime.KeepAlive(activatable)
	runtime.KeepAlive(action)
}

// RelatedAction gets the related Action for activatable.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - action: related Action if one is set.
//
func (activatable *Activatable) RelatedAction() *Action {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))

	_info := girepository.MustFind("Gtk", "Activatable")
	_gret := _info.InvokeIfaceMethod("get_related_action", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(activatable)

	var _action *Action // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// UseActionAppearance gets whether this activatable should reset its layout and
// appearance when setting the related action or when the action changes
// appearance.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: whether activatable uses its actions appearance.
//
func (activatable *Activatable) UseActionAppearance() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))

	_info := girepository.MustFind("Gtk", "Activatable")
	_gret := _info.InvokeIfaceMethod("get_use_action_appearance", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(activatable)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetRelatedAction sets the related action on the activatable object.
//
// > Activatable implementors need to handle the Activatable:related-action >
// property and call gtk_activatable_do_set_related_action() when it changes.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action to set.
//
func (activatable *Activatable) SetRelatedAction(action *Action) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_info := girepository.MustFind("Gtk", "Activatable")
	_info.InvokeIfaceMethod("set_related_action", _args[:], nil)

	runtime.KeepAlive(activatable)
	runtime.KeepAlive(action)
}

// SetUseActionAppearance sets whether this activatable should reset its layout
// and appearance when setting the related action or when the action changes
// appearance
//
// > Activatable implementors need to handle the >
// Activatable:use-action-appearance property and call >
// gtk_activatable_sync_action_properties() to update activatable > if needed.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - useAppearance: whether to use the actions appearance.
//
func (activatable *Activatable) SetUseActionAppearance(useAppearance bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))
	if useAppearance {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Activatable")
	_info.InvokeIfaceMethod("set_use_action_appearance", _args[:], nil)

	runtime.KeepAlive(activatable)
	runtime.KeepAlive(useAppearance)
}

// SyncActionProperties: this is called to update the activatable completely,
// this is called internally when the Activatable:related-action property is set
// or unset and by the implementing class when Activatable:use-action-appearance
// changes.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action (optional): related Action or NULL.
//
func (activatable *Activatable) SyncActionProperties(action *Action) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(activatable).Native()))
	if action != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	}

	_info := girepository.MustFind("Gtk", "Activatable")
	_info.InvokeIfaceMethod("sync_action_properties", _args[:], nil)

	runtime.KeepAlive(activatable)
	runtime.KeepAlive(action)
}
