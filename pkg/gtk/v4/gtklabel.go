// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_Label_ConnectCopyClipboard(gpointer, guintptr);
// extern void _gotk4_gtk4_Label_ConnectActivateCurrentLink(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_Label_ConnectActivateLink(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeLabel = coreglib.Type(girepository.MustFind("Gtk", "Label").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLabel, F: marshalLabel},
	})
}

// Label: GtkLabel widget displays a small amount of text.
//
// As the name implies, most labels are used to label another widget such as a
// button.
//
// !An example GtkLabel (label.png)
//
// CSS nodes
//
//    label
//    ├── [selection]
//    ├── [link]
//    ┊
//    ╰── [link]
//
//
// GtkLabel has a single CSS node with the name label. A wide variety of style
// classes may be applied to labels, such as .title, .subtitle, .dim-label, etc.
// In the GtkShortcutsWindow, labels are used with the .keycap style class.
//
// If the label has a selection, it gets a subnode with name selection.
//
// If the label has links, there is one subnode per link. These subnodes carry
// the link or visited state depending on whether they have been visited. In
// this case, label node also gets a .link style class.
//
//
// GtkLabel as GtkBuildable
//
// The GtkLabel implementation of the GtkBuildable interface supports a custom
// <attributes> element, which supports any number of <attribute> elements. The
// <attribute> element has attributes named “name“, “value“, “start“ and “end“
// and allows you to specify pango.Attribute values for this label.
//
// An example of a UI definition fragment specifying Pango attributes:
//
//    <object class="GtkLabel">
//      <attributes>
//        <attribute name="weight" value="PANGO_WEIGHT_BOLD"/>
//        <attribute name="background" value="red" start="5" end="10"/>
//      </attributes>
//    </object>
//
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
//
// Accessibility
//
// GtkLabel uses the K_ACCESSIBLE_ROLE_LABEL role.
//
//
// Mnemonics
//
// Labels may contain “mnemonics”. Mnemonics are underlined characters in the
// label, used for keyboard navigation. Mnemonics are created by providing a
// string with an underscore before the mnemonic character, such as "_File", to
// the functions gtk.Label.NewWithMnemonic or gtk.Label.SetTextWithMnemonic().
//
// Mnemonics automatically activate any activatable widget the label is inside,
// such as a gtk.Button; if the label is not inside the mnemonic’s target
// widget, you have to tell the label about the target using
// gtk.Label.SetMnemonicWidget. Here’s a simple example where the label is
// inside a button:
//
//    // Pressing Alt+H will activate this button
//    GtkWidget *button = gtk_button_new ();
//    GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
//    gtk_button_set_child (GTK_BUTTON (button), label);
//
//
// There’s a convenience function to create buttons with a mnemonic label
// already inside:
//
//    // Pressing Alt+H will activate this button
//    GtkWidget *button = gtk_button_new_with_mnemonic ("_Hello");
//
//
// To create a mnemonic for a widget alongside the label, such as a gtk.Entry,
// you have to point the label at the entry with gtk.Label.SetMnemonicWidget():
//
//    // Pressing Alt+H will focus the entry
//    GtkWidget *entry = gtk_entry_new ();
//    GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
//    gtk_label_set_mnemonic_widget (GTK_LABEL (label), entry);
//
//
// Markup (styled text)
//
// To make it easy to format text in a label (changing colors, fonts, etc.),
// label text can be provided in a simple markup format:
//
// Here’s how to create a label with a small font:
//
//    GtkWidget *label = gtk_label_new (NULL);
//    gtk_label_set_markup (GTK_LABEL (label), "<small>Small text</small>");
//
//
// (See the Pango manual for complete documentation] of available tags,
// pango.ParseMarkup())
//
// The markup passed to gtk_label_set_markup() must be valid; for example,
// literal <, > and & characters must be escaped as &lt;, &gt;, and &amp;. If
// you pass text obtained from the user, file, or a network to
// gtk.Label.SetMarkup(), you’ll want to escape it with g_markup_escape_text()
// or g_markup_printf_escaped().
//
// Markup strings are just a convenient way to set the pango.AttrList on a
// label; gtk.Label.SetAttributes() may be a simpler way to set attributes in
// some cases. Be careful though; pango.AttrList tends to cause
// internationalization problems, unless you’re applying attributes to the
// entire string (i.e. unless you set the range of each attribute to [0,
// G_MAXINT)). The reason is that specifying the start_index and end_index for a
// pango.Attribute requires knowledge of the exact string being displayed, so
// translations will cause problems.
//
//
// Selectable labels
//
// Labels can be made selectable with gtk.Label.SetSelectable(). Selectable
// labels allow the user to copy the label contents to the clipboard. Only
// labels that contain useful-to-copy information — such as error messages —
// should be made selectable.
//
//
// Text layout
//
// A label can contain any number of paragraphs, but will have performance
// problems if it contains more than a small number. Paragraphs are separated by
// newlines or other paragraph separators understood by Pango.
//
// Labels can automatically wrap text if you call gtk.Label.SetWrap().
//
// gtk.Label.SetJustify() sets how the lines in a label align with one another.
// If you want to set how the label as a whole aligns in its available space,
// see the gtk.Widget:halign and gtk.Widget:valign properties.
//
// The gtk.Label:width-chars and gtk.Label:max-width-chars properties can be
// used to control the size allocation of ellipsized or wrapped labels. For
// ellipsizing labels, if either is specified (and less than the actual text
// size), it is used as the minimum width, and the actual text size is used as
// the natural width of the label. For wrapping labels, width-chars is used as
// the minimum width, if specified, and max-width-chars is used as the natural
// width. Even if max-width-chars specified, wrapping labels will be rewrapped
// to use all of the available width.
//
//
// Links
//
// GTK supports markup for clickable hyperlinks in addition to regular Pango
// markup. The markup for links is borrowed from HTML, using the <a> with
// “href“, “title“ and “class“ attributes. GTK renders links similar to the way
// they appear in web browsers, with colored, underlined text. The “title“
// attribute is displayed as a tooltip on the link. The “class“ attribute is
// used as style class on the CSS node for the link.
//
// An example looks like this:
//
//    const char *text =
//    "Go to the"
//    "<a href=\"http://www.gtk.org title=\"&lt;i&gt;Our&lt;/i&gt; website\">"
//    "GTK website</a> for more...";
//    GtkWidget *label = gtk_label_new (NULL);
//    gtk_label_set_markup (GTK_LABEL (label), text);
//
//
// It is possible to implement custom handling for links and their tooltips with
// the gtk.Label::activate-link signal and the gtk.Label.GetCurrentURI()
// function.
type Label struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Label)(nil)
)

func wrapLabel(obj *coreglib.Object) *Label {
	return &Label{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalLabel(p uintptr) (interface{}, error) {
	return wrapLabel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivateCurrentLink gets emitted when the user activates a link in the
// label.
//
// The ::activate-current-link is a keybinding signal (class.SignalAction.html).
//
// Applications may also emit the signal with g_signal_emit_by_name() if they
// need to control activation of URIs programmatically.
//
// The default bindings for this signal are all forms of the Enter key.
func (v *Label) ConnectActivateCurrentLink(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate-current-link", false, unsafe.Pointer(C._gotk4_gtk4_Label_ConnectActivateCurrentLink), f)
}

// ConnectActivateLink gets emitted to activate a URI.
//
// Applications may connect to it to override the default behaviour, which is to
// call gtk_show_uri().
func (v *Label) ConnectActivateLink(f func(uri string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate-link", false, unsafe.Pointer(C._gotk4_gtk4_Label_ConnectActivateLink), f)
}

// ConnectCopyClipboard gets emitted to copy the slection to the clipboard.
//
// The ::copy-clipboard signal is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Ctrl-c.
func (v *Label) ConnectCopyClipboard(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "copy-clipboard", false, unsafe.Pointer(C._gotk4_gtk4_Label_ConnectCopyClipboard), f)
}
