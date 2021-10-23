// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_state_type_get_type()), F: marshalStateType},
	})
}

type State = uint64

// StateType: possible types of states of an object.
type StateType C.gint

const (
	// StateInvalid indicates an invalid state - probably an error condition.
	StateInvalid StateType = iota
	// StateActive indicates a window is currently the active window, or an
	// object is the active subelement within a container or table.
	// ATK_STATE_ACTIVE should not be used for objects which have
	// ATK_STATE_FOCUSABLE or ATK_STATE_SELECTABLE: Those objects should use
	// ATK_STATE_FOCUSED and ATK_STATE_SELECTED respectively. ATK_STATE_ACTIVE
	// is a means to indicate that an object which is not focusable and not
	// selectable is the currently-active item within its parent container.
	StateActive
	// StateArmed indicates that the object is 'armed', i.e. will be activated
	// by if a pointer button-release event occurs within its bounds. Buttons
	// often enter this state when a pointer click occurs within their bounds,
	// as a precursor to activation. ATK_STATE_ARMED has been deprecated since
	// ATK-2.16 and should not be used in newly-written code.
	StateArmed
	// StateBusy indicates the current object is busy, i.e. onscreen
	// representation is in the process of changing, or the object is
	// temporarily unavailable for interaction due to activity already in
	// progress. This state may be used by implementors of Document to indicate
	// that content loading is underway. It also may indicate other 'pending'
	// conditions; clients may wish to interrogate this object when the
	// ATK_STATE_BUSY flag is removed.
	StateBusy
	// StateChecked indicates this object is currently checked, for instance a
	// checkbox is 'non-empty'.
	StateChecked
	// StateDefunct indicates that this object no longer has a valid backing
	// widget (for instance, if its peer object has been destroyed).
	StateDefunct
	// StateEditable indicates that this object can contain text, and that the
	// user can change the textual contents of this object by editing those
	// contents directly. For an object which is expected to be editable due to
	// its type, but which cannot be edited due to the application or platform
	// preventing the user from doing so, that object's StateSet should lack
	// ATK_STATE_EDITABLE and should contain ATK_STATE_READ_ONLY.
	StateEditable
	// StateEnabled indicates that this object is enabled, i.e. that it
	// currently reflects some application state. Objects that are "greyed out"
	// may lack this state, and may lack the STATE_SENSITIVE if direct user
	// interaction cannot cause them to acquire STATE_ENABLED. See also:
	// ATK_STATE_SENSITIVE.
	StateEnabled
	// StateExpandable indicates this object allows progressive disclosure of
	// its children.
	StateExpandable
	// StateExpanded indicates this object its expanded - see
	// ATK_STATE_EXPANDABLE above.
	StateExpanded
	// StateFocusable indicates this object can accept keyboard focus, which
	// means all events resulting from typing on the keyboard will normally be
	// passed to it when it has focus.
	StateFocusable
	// StateFocused indicates this object currently has the keyboard focus.
	StateFocused
	// StateHorizontal indicates the orientation of this object is horizontal;
	// used, for instance, by objects of ATK_ROLE_SCROLL_BAR. For objects where
	// vertical/horizontal orientation is especially meaningful.
	StateHorizontal
	// StateIconified indicates this object is minimized and is represented only
	// by an icon.
	StateIconified
	// StateModal indicates something must be done with this object before the
	// user can interact with an object in a different window.
	StateModal
	// StateMultiLine indicates this (text) object can contain multiple lines of
	// text.
	StateMultiLine
	// StateMultiselectable indicates this object allows more than one of its
	// children to be selected at the same time, or in the case of text objects,
	// that the object supports non-contiguous text selections.
	StateMultiselectable
	// StateOpaque indicates this object paints every pixel within its
	// rectangular region.
	StateOpaque
	// StatePressed indicates this object is currently pressed.
	StatePressed
	// StateResizable indicates the size of this object is not fixed.
	StateResizable
	// StateSelectable indicates this object is the child of an object that
	// allows its children to be selected and that this child is one of those
	// children that can be selected.
	StateSelectable
	// StateSelected indicates this object is the child of an object that allows
	// its children to be selected and that this child is one of those children
	// that has been selected.
	StateSelected
	// StateSensitive indicates this object is sensitive, e.g. to user
	// interaction. STATE_SENSITIVE usually accompanies STATE_ENABLED for
	// user-actionable controls, but may be found in the absence of
	// STATE_ENABLED if the current visible state of the control is
	// "disconnected" from the application state. In such cases, direct user
	// interaction can often result in the object gaining STATE_SENSITIVE, for
	// instance if a user makes an explicit selection using an object whose
	// current state is ambiguous or undefined. see STATE_ENABLED,
	// STATE_INDETERMINATE.
	StateSensitive
	// StateShowing indicates this object, the object's parent, the object's
	// parent's parent, and so on, are all 'shown' to the end-user, i.e. subject
	// to "exposure" if blocking or obscuring objects do not interpose between
	// this object and the top of the window stack.
	StateShowing
	// StateSingleLine indicates this (text) object can contain only a single
	// line of text.
	StateSingleLine
	// StateStale indicates that the information returned for this object may no
	// longer be synchronized with the application state. This is implied if the
	// object has STATE_TRANSIENT, and can also occur towards the end of the
	// object peer's lifecycle. It can also be used to indicate that the index
	// associated with this object has changed since the user accessed the
	// object (in lieu of "index-in-parent-changed" events).
	StateStale
	// StateTransient indicates this object is transient, i.e. a snapshot which
	// may not emit events when its state changes. Data from objects with
	// ATK_STATE_TRANSIENT should not be cached, since there may be no
	// notification given when the cached data becomes obsolete.
	StateTransient
	// StateVertical indicates the orientation of this object is vertical.
	StateVertical
	// StateVisible indicates this object is visible, e.g. has been explicitly
	// marked for exposure to the user. **note**: ATK_STATE_VISIBLE is no
	// guarantee that the object is actually unobscured on the screen, only that
	// it is 'potentially' visible, barring obstruction, being scrolled or
	// clipped out of the field of view, or having an ancestor container that
	// has not yet made visible. A widget is potentially onscreen if it has both
	// ATK_STATE_VISIBLE and ATK_STATE_SHOWING. The absence of ATK_STATE_VISIBLE
	// and ATK_STATE_SHOWING is semantically equivalent to saying that an object
	// is 'hidden'. See also ATK_STATE_TRUNCATED, which applies if an object
	// with ATK_STATE_VISIBLE and ATK_STATE_SHOWING set lies within a viewport
	// which means that its contents are clipped, e.g. a truncated spreadsheet
	// cell or an image within a scrolling viewport. Mostly useful for
	// screen-review and magnification algorithms.
	StateVisible
	// StateManagesDescendants indicates that "active-descendant-changed" event
	// is sent when children become 'active' (i.e. are selected or navigated to
	// onscreen). Used to prevent need to enumerate all children in very large
	// containers, like tables. The presence of STATE_MANAGES_DESCENDANTS is an
	// indication to the client. that the children should not, and need not, be
	// enumerated by the client. Objects implementing this state are expected to
	// provide relevant state notifications to listening clients, for instance
	// notifications of visibility changes and activation of their contained
	// child objects, without the client having previously requested references
	// to those children.
	StateManagesDescendants
	// StateIndeterminate indicates that the value, or some other quantifiable
	// property, of this AtkObject cannot be fully determined. In the case of a
	// large data set in which the total number of items in that set is unknown
	// (e.g. 1 of 999+), implementors should expose the currently-known set size
	// (999) along with this state. In the case of a check box, this state
	// should be used to indicate that the check box is a tri-state check box
	// which is currently neither checked nor unchecked.
	StateIndeterminate
	// StateTruncated indicates that an object is truncated, e.g. a text value
	// in a speradsheet cell.
	StateTruncated
	// StateRequired indicates that explicit user interaction with an object is
	// required by the user interface, e.g. a required field in a "web-form"
	// interface.
	StateRequired
	// StateInvalidEntry indicates that the object has encountered an error
	// condition due to failure of input validation. For instance, a form
	// control may acquire this state in response to invalid or malformed user
	// input.
	StateInvalidEntry
	// StateSupportsAutocompletion indicates that the object in question
	// implements some form of ¨typeahead¨ or pre-selection behavior whereby
	// entering the first character of one or more sub-elements causes those
	// elements to scroll into view or become selected. Subsequent character
	// input may narrow the selection further as long as one or more
	// sub-elements match the string. This state is normally only useful and
	// encountered on objects that implement Selection. In some cases the
	// typeahead behavior may result in full or partial ¨completion¨ of the data
	// in the input field, in which case these input events may trigger
	// text-changed events from the AtkText interface. This state supplants
	// ATK_ROLE_AUTOCOMPLETE.
	StateSupportsAutocompletion
	// StateSelectableText indicates that the object in question supports text
	// selection. It should only be exposed on objects which implement the Text
	// interface, in order to distinguish this state from ATK_STATE_SELECTABLE,
	// which infers that the object in question is a selectable child of an
	// object which implements Selection. While similar, text selection and
	// subelement selection are distinct operations.
	StateSelectableText
	// StateDefault indicates that the object is the "default" active component,
	// i.e. the object which is activated by an end-user press of the "Enter" or
	// "Return" key. Typically a "close" or "submit" button.
	StateDefault
	// StateAnimated indicates that the object changes its appearance
	// dynamically as an inherent part of its presentation. This state may come
	// and go if an object is only temporarily animated on the way to a 'final'
	// onscreen presentation. **note**: some applications, notably content
	// viewers, may not be able to detect all kinds of animated content.
	// Therefore the absence of this state should not be taken as definitive
	// evidence that the object's visual representation is static; this state is
	// advisory.
	StateAnimated
	// StateVisited indicates that the object (typically a hyperlink) has
	// already been 'activated', and/or its backing data has already been
	// downloaded, rendered, or otherwise "visited".
	StateVisited
	// StateCheckable indicates this object has the potential to be checked,
	// such as a checkbox or toggle-able table cell. Since: ATK-2.12.
	StateCheckable
	// StateHasPopup indicates that the object has a popup context menu or
	// sub-level menu which may or may not be showing. This means that
	// activation renders conditional content. Note that ordinary tooltips are
	// not considered popups in this context. Since: ATK-2.12.
	StateHasPopup
	// StateHasTooltip indicates this object has a tooltip. Since: ATK-2.16.
	StateHasTooltip
	// StateReadOnly indicates that a widget which is ENABLED and SENSITIVE has
	// a value which can be read, but not modified, by the user. Note that this
	// state should only be applied to widget types whose value is normally
	// directly user modifiable, such as check boxes, radio buttons, spin
	// buttons, text input fields, and combo boxes, as a means to convey that
	// the expected interaction with that widget is not possible. When the
	// expected interaction with a widget does not include modification by the
	// user, as is the case with labels and containers, ATK_STATE_READ_ONLY
	// should not be applied. See also ATK_STATE_EDITABLE. Since: ATK-2-16.
	StateReadOnly
	// StateLastDefined: not a valid state, used for finding end of enumeration.
	StateLastDefined
)

func marshalStateType(p uintptr) (interface{}, error) {
	return StateType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for StateType.
func (s StateType) String() string {
	switch s {
	case StateInvalid:
		return "Invalid"
	case StateActive:
		return "Active"
	case StateArmed:
		return "Armed"
	case StateBusy:
		return "Busy"
	case StateChecked:
		return "Checked"
	case StateDefunct:
		return "Defunct"
	case StateEditable:
		return "Editable"
	case StateEnabled:
		return "Enabled"
	case StateExpandable:
		return "Expandable"
	case StateExpanded:
		return "Expanded"
	case StateFocusable:
		return "Focusable"
	case StateFocused:
		return "Focused"
	case StateHorizontal:
		return "Horizontal"
	case StateIconified:
		return "Iconified"
	case StateModal:
		return "Modal"
	case StateMultiLine:
		return "MultiLine"
	case StateMultiselectable:
		return "Multiselectable"
	case StateOpaque:
		return "Opaque"
	case StatePressed:
		return "Pressed"
	case StateResizable:
		return "Resizable"
	case StateSelectable:
		return "Selectable"
	case StateSelected:
		return "Selected"
	case StateSensitive:
		return "Sensitive"
	case StateShowing:
		return "Showing"
	case StateSingleLine:
		return "SingleLine"
	case StateStale:
		return "Stale"
	case StateTransient:
		return "Transient"
	case StateVertical:
		return "Vertical"
	case StateVisible:
		return "Visible"
	case StateManagesDescendants:
		return "ManagesDescendants"
	case StateIndeterminate:
		return "Indeterminate"
	case StateTruncated:
		return "Truncated"
	case StateRequired:
		return "Required"
	case StateInvalidEntry:
		return "InvalidEntry"
	case StateSupportsAutocompletion:
		return "SupportsAutocompletion"
	case StateSelectableText:
		return "SelectableText"
	case StateDefault:
		return "Default"
	case StateAnimated:
		return "Animated"
	case StateVisited:
		return "Visited"
	case StateCheckable:
		return "Checkable"
	case StateHasPopup:
		return "HasPopup"
	case StateHasTooltip:
		return "HasTooltip"
	case StateReadOnly:
		return "ReadOnly"
	case StateLastDefined:
		return "LastDefined"
	default:
		return fmt.Sprintf("StateType(%d)", s)
	}
}

// StateTypeForName gets the StateType corresponding to the description string
// name.
//
// The function takes the following parameters:
//
//    - name: character string state name.
//
func StateTypeForName(name string) StateType {
	var _arg1 *C.gchar       // out
	var _cret C.AtkStateType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_state_type_for_name(_arg1)
	runtime.KeepAlive(name)

	var _stateType StateType // out

	_stateType = StateType(_cret)

	return _stateType
}

// StateTypeGetName gets the description string describing the StateType type.
//
// The function takes the following parameters:
//
//    - typ whose name is required.
//
func StateTypeGetName(typ StateType) string {
	var _arg1 C.AtkStateType // out
	var _cret *C.gchar       // in

	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_type_get_name(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StateTypeRegister: register a new object state.
//
// The function takes the following parameters:
//
//    - name: character string describing the new state.
//
func StateTypeRegister(name string) StateType {
	var _arg1 *C.gchar       // out
	var _cret C.AtkStateType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_state_type_register(_arg1)
	runtime.KeepAlive(name)

	var _stateType StateType // out

	_stateType = StateType(_cret)

	return _stateType
}
