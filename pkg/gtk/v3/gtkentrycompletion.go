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
		{T: externglib.Type(C.gtk_entry_completion_get_type()), F: marshalEntryCompletion},
	})
}

// EntryCompletion is an auxiliary object to be used in conjunction with Entry
// to provide the completion functionality. It implements the CellLayout
// interface, to allow the user to add extra cells to the TreeView with
// completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, EntryCompletion checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see gtk_entry_completion_set_text_column()), but
// this can be overridden with a custom match function (see
// gtk_entry_completion_set_match_func()).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// EntryCompletion::match-selected signal and updating the entry in the signal
// handler. Note that you should return true from the signal handler to suppress
// the default behaviour.
//
// To add completion functionality to an entry, use gtk_entry_set_completion().
//
// In addition to regular completion matches, which will be inserted into the
// entry when they are selected, EntryCompletion also allows to display
// “actions” in the popup window. Their appearance is similar to menuitems, to
// differentiate them clearly from completion strings. When an action is
// selected, the EntryCompletion::action-activated signal is emitted.
//
// GtkEntryCompletion uses a TreeModelFilter model to represent the subset of
// the entire model that is currently matching. While the GtkEntryCompletion
// signals EntryCompletion::match-selected and EntryCompletion::cursor-on-match
// take the original model and an iter pointing to that model as arguments,
// other callbacks and signals (such as CellLayoutDataFuncs or
// CellArea::apply-attributes) will generally take the filter model as argument.
// As long as you are only calling gtk_tree_model_get(), this will make no
// difference to you. If for some reason, you need the original model, use
// gtk_tree_model_filter_get_model(). Don’t forget to use
// gtk_tree_model_filter_convert_iter_to_child_iter() to obtain a matching iter.
type EntryCompletion interface {
	gextras.Objector
	Buildable
	CellLayout

	// Complete requests a completion operation, or in other words a refiltering
	// of the current list with completions, using the current key. The
	// completion list view will be updated accordingly.
	Complete()
	// ComputePrefix computes the common prefix that is shared by all rows in
	// @completion that start with @key. If no row matches @key, nil will be
	// returned. Note that a text column must have been set for this function to
	// work, see gtk_entry_completion_set_text_column() for details.
	ComputePrefix(key string) string
	// DeleteAction deletes the action at @index_ from @completion’s action
	// list.
	//
	// Note that @index_ is a relative position and the position of an action
	// may have changed since it was inserted.
	DeleteAction(index_ int)
	// CompletionPrefix: get the original text entered by the user that
	// triggered the completion or nil if there’s no completion ongoing.
	CompletionPrefix() string
	// Entry gets the entry @completion has been attached to.
	Entry() Widget
	// InlineCompletion returns whether the common prefix of the possible
	// completions should be automatically inserted in the entry.
	InlineCompletion() bool
	// InlineSelection returns true if inline-selection mode is turned on.
	InlineSelection() bool
	// MinimumKeyLength returns the minimum key length as set for @completion.
	MinimumKeyLength() int
	// Model returns the model the EntryCompletion is using as data source.
	// Returns nil if the model is unset.
	Model() TreeModel
	// PopupCompletion returns whether the completions should be presented in a
	// popup window.
	PopupCompletion() bool
	// PopupSetWidth returns whether the completion popup window will be resized
	// to the width of the entry.
	PopupSetWidth() bool
	// PopupSingleMatch returns whether the completion popup window will appear
	// even if there is only a single match.
	PopupSingleMatch() bool
	// TextColumn returns the column in the model of @completion to get strings
	// from.
	TextColumn() int
	// InsertActionMarkup inserts an action in @completion’s action item list at
	// position @index_ with markup @markup.
	InsertActionMarkup(index_ int, markup string)
	// InsertActionText inserts an action in @completion’s action item list at
	// position @index_ with text @text. If you want the action item to have
	// markup, use gtk_entry_completion_insert_action_markup().
	//
	// Note that @index_ is a relative position in the list of actions and the
	// position of an action can change when deleting a different action.
	InsertActionText(index_ int, text string)
	// InsertPrefix requests a prefix insertion.
	InsertPrefix()
	// SetInlineCompletion sets whether the common prefix of the possible
	// completions should be automatically inserted in the entry.
	SetInlineCompletion(inlineCompletion bool)
	// SetInlineSelection sets whether it is possible to cycle through the
	// possible completions inside the entry.
	SetInlineSelection(inlineSelection bool)
	// SetMinimumKeyLength requires the length of the search key for @completion
	// to be at least @length. This is useful for long lists, where completing
	// using a small key takes a lot of time and will come up with meaningless
	// results anyway (ie, a too large dataset).
	SetMinimumKeyLength(length int)
	// SetModel sets the model for a EntryCompletion. If @completion already has
	// a model set, it will remove it before setting the new model. If model is
	// nil, then it will unset the model.
	SetModel(model TreeModel)
	// SetPopupCompletion sets whether the completions should be presented in a
	// popup window.
	SetPopupCompletion(popupCompletion bool)
	// SetPopupSetWidth sets whether the completion popup window will be resized
	// to be the same width as the entry.
	SetPopupSetWidth(popupSetWidth bool)
	// SetPopupSingleMatch sets whether the completion popup window will appear
	// even if there is only a single match. You may want to set this to false
	// if you are using [inline
	// completion][GtkEntryCompletion--inline-completion].
	SetPopupSingleMatch(popupSingleMatch bool)
	// SetTextColumn: convenience function for setting up the most used case of
	// this code: a completion list with just strings. This function will set up
	// @completion to have a list displaying all (and just) strings in the
	// completion list, and to get those strings from @column in the model of
	// @completion.
	//
	// This functions creates and adds a CellRendererText for the selected
	// column. If you need to set the text column, but don't want the cell
	// renderer, use g_object_set() to set the EntryCompletion:text-column
	// property directly.
	SetTextColumn(column int)
}

// entryCompletion implements the EntryCompletion class.
type entryCompletion struct {
	gextras.Objector
	Buildable
	CellLayout
}

var _ EntryCompletion = (*entryCompletion)(nil)

// WrapEntryCompletion wraps a GObject to the right type. It is
// primarily used internally.
func WrapEntryCompletion(obj *externglib.Object) EntryCompletion {
	return entryCompletion{
		Objector:   obj,
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
	}
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntryCompletion(obj), nil
}

// NewEntryCompletion constructs a class EntryCompletion.
func NewEntryCompletion() EntryCompletion {
	var _cret C.GtkEntryCompletion // in

	_cret = C.gtk_entry_completion_new()

	var _entryCompletion EntryCompletion // out

	_entryCompletion = WrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// NewEntryCompletionWithArea constructs a class EntryCompletion.
func NewEntryCompletionWithArea(area CellArea) EntryCompletion {
	var _arg1 *C.GtkCellArea       // out
	var _cret C.GtkEntryCompletion // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_entry_completion_new_with_area(_arg1)

	var _entryCompletion EntryCompletion // out

	_entryCompletion = WrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// Complete requests a completion operation, or in other words a refiltering
// of the current list with completions, using the current key. The
// completion list view will be updated accordingly.
func (c entryCompletion) Complete() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_complete(_arg0)
}

// ComputePrefix computes the common prefix that is shared by all rows in
// @completion that start with @key. If no row matches @key, nil will be
// returned. Note that a text column must have been set for this function to
// work, see gtk_entry_completion_set_text_column() for details.
func (c entryCompletion) ComputePrefix(key string) string {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.char               // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_entry_completion_compute_prefix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DeleteAction deletes the action at @index_ from @completion’s action
// list.
//
// Note that @index_ is a relative position and the position of an action
// may have changed since it was inserted.
func (c entryCompletion) DeleteAction(index_ int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(index_)

	C.gtk_entry_completion_delete_action(_arg0, _arg1)
}

// CompletionPrefix: get the original text entered by the user that
// triggered the completion or nil if there’s no completion ongoing.
func (c entryCompletion) CompletionPrefix() string {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_completion_prefix(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Entry gets the entry @completion has been attached to.
func (c entryCompletion) Entry() Widget {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_entry(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// InlineCompletion returns whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (c entryCompletion) InlineCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InlineSelection returns true if inline-selection mode is turned on.
func (c entryCompletion) InlineSelection() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MinimumKeyLength returns the minimum key length as set for @completion.
func (c entryCompletion) MinimumKeyLength() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gint                // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_minimum_key_length(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Model returns the model the EntryCompletion is using as data source.
// Returns nil if the model is unset.
func (c entryCompletion) Model() TreeModel {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

// PopupCompletion returns whether the completions should be presented in a
// popup window.
func (c entryCompletion) PopupCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSetWidth returns whether the completion popup window will be resized
// to the width of the entry.
func (c entryCompletion) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSingleMatch returns whether the completion popup window will appear
// even if there is only a single match.
func (c entryCompletion) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextColumn returns the column in the model of @completion to get strings
// from.
func (c entryCompletion) TextColumn() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gint                // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_text_column(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// InsertActionMarkup inserts an action in @completion’s action item list at
// position @index_ with markup @markup.
func (c entryCompletion) InsertActionMarkup(index_ int, markup string) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out
	var _arg2 *C.gchar              // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(index_)
	_arg2 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_completion_insert_action_markup(_arg0, _arg1, _arg2)
}

// InsertActionText inserts an action in @completion’s action item list at
// position @index_ with text @text. If you want the action item to have
// markup, use gtk_entry_completion_insert_action_markup().
//
// Note that @index_ is a relative position in the list of actions and the
// position of an action can change when deleting a different action.
func (c entryCompletion) InsertActionText(index_ int, text string) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out
	var _arg2 *C.gchar              // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(index_)
	_arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_completion_insert_action_text(_arg0, _arg1, _arg2)
}

// InsertPrefix requests a prefix insertion.
func (c entryCompletion) InsertPrefix() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_insert_prefix(_arg0)
}

// SetInlineCompletion sets whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (c entryCompletion) SetInlineCompletion(inlineCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_completion(_arg0, _arg1)
}

// SetInlineSelection sets whether it is possible to cycle through the
// possible completions inside the entry.
func (c entryCompletion) SetInlineSelection(inlineSelection bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineSelection {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_selection(_arg0, _arg1)
}

// SetMinimumKeyLength requires the length of the search key for @completion
// to be at least @length. This is useful for long lists, where completing
// using a small key takes a lot of time and will come up with meaningless
// results anyway (ie, a too large dataset).
func (c entryCompletion) SetMinimumKeyLength(length int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(length)

	C.gtk_entry_completion_set_minimum_key_length(_arg0, _arg1)
}

// SetModel sets the model for a EntryCompletion. If @completion already has
// a model set, it will remove it before setting the new model. If model is
// nil, then it will unset the model.
func (c entryCompletion) SetModel(model TreeModel) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_entry_completion_set_model(_arg0, _arg1)
}

// SetPopupCompletion sets whether the completions should be presented in a
// popup window.
func (c entryCompletion) SetPopupCompletion(popupCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_completion(_arg0, _arg1)
}

// SetPopupSetWidth sets whether the completion popup window will be resized
// to be the same width as the entry.
func (c entryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
}

// SetPopupSingleMatch sets whether the completion popup window will appear
// even if there is only a single match. You may want to set this to false
// if you are using [inline
// completion][GtkEntryCompletion--inline-completion].
func (c entryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
}

// SetTextColumn: convenience function for setting up the most used case of
// this code: a completion list with just strings. This function will set up
// @completion to have a list displaying all (and just) strings in the
// completion list, and to get those strings from @column in the model of
// @completion.
//
// This functions creates and adds a CellRendererText for the selected
// column. If you need to set the text column, but don't want the cell
// renderer, use g_object_set() to set the EntryCompletion:text-column
// property directly.
func (c entryCompletion) SetTextColumn(column int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(column)

	C.gtk_entry_completion_set_text_column(_arg0, _arg1)
}
