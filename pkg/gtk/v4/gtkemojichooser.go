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
// extern void _gotk4_gtk4_EmojiChooser_ConnectEmojiPicked(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeEmojiChooser = coreglib.Type(girepository.MustFind("Gtk", "EmojiChooser").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEmojiChooser, F: marshalEmojiChooser},
	})
}

// EmojiChooser: GtkEmojiChooser is used by text widgets such as GtkEntry or
// GtkTextView to let users insert Emoji characters.
//
// !An example GtkEmojiChooser (emojichooser.png)
//
// GtkEmojiChooser emits the gtk.EmojiChooser::emoji-picked signal when an Emoji
// is selected.
//
// CSS nodes
//
//    popover
//    ├── box.emoji-searchbar
//    │   ╰── entry.search
//    ╰── box.emoji-toolbar
//        ├── button.image-button.emoji-section
//        ├── ...
//        ╰── button.image-button.emoji-section
//
//
// Every GtkEmojiChooser consists of a main node called popover. The contents of
// the popover are largely implementation defined and supposed to inherit
// general styles. The top searchbar used to search emoji and gets the
// .emoji-searchbar style class itself. The bottom toolbar used to switch
// between different emoji categories consists of buttons with the
// .emoji-section style class and gets the .emoji-toolbar style class itself.
type EmojiChooser struct {
	_ [0]func() // equal guard
	Popover
}

var (
	_ Widgetter         = (*EmojiChooser)(nil)
	_ coreglib.Objector = (*EmojiChooser)(nil)
)

func wrapEmojiChooser(obj *coreglib.Object) *EmojiChooser {
	return &EmojiChooser{
		Popover: Popover{
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
			Object: obj,
			NativeSurface: NativeSurface{
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
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalEmojiChooser(p uintptr) (interface{}, error) {
	return wrapEmojiChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEmojiPicked is emitted when the user selects an Emoji.
func (v *EmojiChooser) ConnectEmojiPicked(f func(text string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "emoji-picked", false, unsafe.Pointer(C._gotk4_gtk4_EmojiChooser_ConnectEmojiPicked), f)
}
