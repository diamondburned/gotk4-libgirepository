// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeInputPurpose = coreglib.Type(girepository.MustFind("Gtk", "InputPurpose").RegisteredGType())
	GTypeLevelBarMode = coreglib.Type(girepository.MustFind("Gtk", "LevelBarMode").RegisteredGType())
	GTypeInputHints   = coreglib.Type(girepository.MustFind("Gtk", "InputHints").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInputPurpose, F: marshalInputPurpose},
		coreglib.TypeMarshaler{T: GTypeLevelBarMode, F: marshalLevelBarMode},
		coreglib.TypeMarshaler{T: GTypeInputHints, F: marshalInputHints},
	})
}

// InputPurpose describes primary purpose of the input widget. This information
// is useful for on-screen keyboards and similar input methods to decide which
// keys should be presented to the user.
//
// Note that the purpose is not meant to impose a totally strict rule about
// allowed characters, and does not replace input validation. It is fine for an
// on-screen keyboard to let the user override the character set restriction
// that is expressed by the purpose. The application is expected to validate the
// entry contents, even if it specified a purpose.
//
// The difference between GTK_INPUT_PURPOSE_DIGITS and GTK_INPUT_PURPOSE_NUMBER
// is that the former accepts only digits while the latter also some punctuation
// (like commas or points, plus, minus) and “e” or “E” as in 3.14E+000.
//
// This enumeration may be extended in the future; input methods should
// interpret unknown values as “free form”.
type InputPurpose C.gint

const (
	// InputPurposeFreeForm: allow any character.
	InputPurposeFreeForm InputPurpose = iota
	// InputPurposeAlpha: allow only alphabetic characters.
	InputPurposeAlpha
	// InputPurposeDigits: allow only digits.
	InputPurposeDigits
	// InputPurposeNumber: edited field expects numbers.
	InputPurposeNumber
	// InputPurposePhone: edited field expects phone number.
	InputPurposePhone
	// InputPurposeURL: edited field expects URL.
	InputPurposeURL
	// InputPurposeEmail: edited field expects email address.
	InputPurposeEmail
	// InputPurposeName: edited field expects the name of a person.
	InputPurposeName
	// InputPurposePassword: like GTK_INPUT_PURPOSE_FREE_FORM, but characters
	// are hidden.
	InputPurposePassword
	// InputPurposePIN: like GTK_INPUT_PURPOSE_DIGITS, but characters are
	// hidden.
	InputPurposePIN
	// InputPurposeTerminal: allow any character, in addition to control codes.
	InputPurposeTerminal
)

func marshalInputPurpose(p uintptr) (interface{}, error) {
	return InputPurpose(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputPurpose.
func (i InputPurpose) String() string {
	switch i {
	case InputPurposeFreeForm:
		return "FreeForm"
	case InputPurposeAlpha:
		return "Alpha"
	case InputPurposeDigits:
		return "Digits"
	case InputPurposeNumber:
		return "Number"
	case InputPurposePhone:
		return "Phone"
	case InputPurposeURL:
		return "URL"
	case InputPurposeEmail:
		return "Email"
	case InputPurposeName:
		return "Name"
	case InputPurposePassword:
		return "Password"
	case InputPurposePIN:
		return "PIN"
	case InputPurposeTerminal:
		return "Terminal"
	default:
		return fmt.Sprintf("InputPurpose(%d)", i)
	}
}

// LevelBarMode describes how LevelBar contents should be rendered. Note that
// this enumeration could be extended with additional modes in the future.
type LevelBarMode C.gint

const (
	// LevelBarModeContinuous: bar has a continuous mode.
	LevelBarModeContinuous LevelBarMode = iota
	// LevelBarModeDiscrete: bar has a discrete mode.
	LevelBarModeDiscrete
)

func marshalLevelBarMode(p uintptr) (interface{}, error) {
	return LevelBarMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for LevelBarMode.
func (l LevelBarMode) String() string {
	switch l {
	case LevelBarModeContinuous:
		return "Continuous"
	case LevelBarModeDiscrete:
		return "Discrete"
	default:
		return fmt.Sprintf("LevelBarMode(%d)", l)
	}
}

// InputHints describes hints that might be taken into account by input methods
// or applications. Note that input methods may already tailor their behaviour
// according to the InputPurpose of the entry.
//
// Some common sense is expected when using these flags - mixing
// GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.
//
// This enumeration may be extended in the future; input methods should ignore
// unknown values.
type InputHints C.guint

const (
	// InputHintNone: no special behaviour suggested.
	InputHintNone InputHints = 0b0
	// InputHintSpellcheck: suggest checking for typos.
	InputHintSpellcheck InputHints = 0b1
	// InputHintNoSpellcheck: suggest not checking for typos.
	InputHintNoSpellcheck InputHints = 0b10
	// InputHintWordCompletion: suggest word completion.
	InputHintWordCompletion InputHints = 0b100
	// InputHintLowercase: suggest to convert all text to lowercase.
	InputHintLowercase InputHints = 0b1000
	// InputHintUppercaseChars: suggest to capitalize all text.
	InputHintUppercaseChars InputHints = 0b10000
	// InputHintUppercaseWords: suggest to capitalize the first character of
	// each word.
	InputHintUppercaseWords InputHints = 0b100000
	// InputHintUppercaseSentences: suggest to capitalize the first word of each
	// sentence.
	InputHintUppercaseSentences InputHints = 0b1000000
	// InputHintInhibitOSK: suggest to not show an onscreen keyboard (e.g for a
	// calculator that already has all the keys).
	InputHintInhibitOSK InputHints = 0b10000000
	// InputHintVerticalWriting: text is vertical. Since 3.18.
	InputHintVerticalWriting InputHints = 0b100000000
	// InputHintEmoji: suggest offering Emoji support. Since 3.22.20.
	InputHintEmoji InputHints = 0b1000000000
	// InputHintNoEmoji: suggest not offering Emoji support. Since 3.22.20.
	InputHintNoEmoji InputHints = 0b10000000000
)

func marshalInputHints(p uintptr) (interface{}, error) {
	return InputHints(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for InputHints.
func (i InputHints) String() string {
	if i == 0 {
		return "InputHints(0)"
	}

	var builder strings.Builder
	builder.Grow(251)

	for i != 0 {
		next := i & (i - 1)
		bit := i - next

		switch bit {
		case InputHintNone:
			builder.WriteString("None|")
		case InputHintSpellcheck:
			builder.WriteString("Spellcheck|")
		case InputHintNoSpellcheck:
			builder.WriteString("NoSpellcheck|")
		case InputHintWordCompletion:
			builder.WriteString("WordCompletion|")
		case InputHintLowercase:
			builder.WriteString("Lowercase|")
		case InputHintUppercaseChars:
			builder.WriteString("UppercaseChars|")
		case InputHintUppercaseWords:
			builder.WriteString("UppercaseWords|")
		case InputHintUppercaseSentences:
			builder.WriteString("UppercaseSentences|")
		case InputHintInhibitOSK:
			builder.WriteString("InhibitOSK|")
		case InputHintVerticalWriting:
			builder.WriteString("VerticalWriting|")
		case InputHintEmoji:
			builder.WriteString("Emoji|")
		case InputHintNoEmoji:
			builder.WriteString("NoEmoji|")
		default:
			builder.WriteString(fmt.Sprintf("InputHints(0b%b)|", bit))
		}

		i = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if i contains other.
func (i InputHints) Has(other InputHints) bool {
	return (i & other) == other
}
