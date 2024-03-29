// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// STOCK_ADD: “Add” item and icon.
//
// Deprecated: Use named icon &quot;list-add&quot; or the label
// &quot;_Add&quot;.
const STOCK_ADD = "gtk-add"

// STOCK_APPLY: “Apply” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_Apply&quot;.
const STOCK_APPLY = "gtk-apply"

// STOCK_BOLD: “Bold” item and icon.
//
// Deprecated: Use named icon &quot;format-text-bold&quot;.
const STOCK_BOLD = "gtk-bold"

// STOCK_CANCEL: “Cancel” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_Cancel&quot;.
const STOCK_CANCEL = "gtk-cancel"

// STOCK_CDROM: “CD-Rom” item and icon.
//
// Deprecated: Use named icon &quot;media-optical&quot;.
const STOCK_CDROM = "gtk-cdrom"

// STOCK_CLEAR: “Clear” item and icon.
//
// Deprecated: Use named icon &quot;edit-clear&quot;.
const STOCK_CLEAR = "gtk-clear"

// STOCK_CLOSE: “Close” item and icon.
//
// Deprecated: Use named icon &quot;window-close&quot; or the label
// &quot;_Close&quot;.
const STOCK_CLOSE = "gtk-close"

// STOCK_CONVERT: “Convert” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_CONVERT = "gtk-convert"

// STOCK_COPY: “Copy” item and icon.
//
// Deprecated: Use the named icon &quot;edit-copy&quot; or the label
// &quot;_Copy&quot;.
const STOCK_COPY = "gtk-copy"

// STOCK_CUT: “Cut” item and icon.
//
// Deprecated: Use the named icon &quot;edit-cut&quot; or the label
// &quot;Cu_t&quot;.
const STOCK_CUT = "gtk-cut"

// STOCK_DELETE: “Delete” item and icon.
//
// Deprecated: Use the named icon &quot;edit-delete&quot; or the label
// &quot;_Delete&quot;.
const STOCK_DELETE = "gtk-delete"

// STOCK_DIALOG_ERROR: “Error” item and icon.
//
// Deprecated: Use named icon &quot;dialog-error&quot;.
const STOCK_DIALOG_ERROR = "gtk-dialog-error"

// STOCK_DIALOG_INFO: “Information” item and icon.
//
// Deprecated: Use named icon &quot;dialog-information&quot;.
const STOCK_DIALOG_INFO = "gtk-dialog-info"

// STOCK_DIALOG_QUESTION: “Question” item and icon.
//
// Deprecated: Use named icon &quot;dialog-question&quot;.
const STOCK_DIALOG_QUESTION = "gtk-dialog-question"

// STOCK_DIALOG_WARNING: “Warning” item and icon.
//
// Deprecated: Use named icon &quot;dialog-warning&quot;.
const STOCK_DIALOG_WARNING = "gtk-dialog-warning"

// STOCK_DND: “Drag-And-Drop” icon.
//
// Deprecated: since version 3.10.
const STOCK_DND = "gtk-dnd"

// STOCK_DND_MULTIPLE: “Drag-And-Drop multiple” icon.
//
// Deprecated: since version 3.10.
const STOCK_DND_MULTIPLE = "gtk-dnd-multiple"

// STOCK_EXECUTE: “Execute” item and icon.
//
// Deprecated: Use named icon &quot;system-run&quot;.
const STOCK_EXECUTE = "gtk-execute"

// STOCK_FIND: “Find” item and icon.
//
// Deprecated: Use named icon &quot;edit-find&quot;.
const STOCK_FIND = "gtk-find"

// STOCK_FIND_AND_REPLACE: “Find and Replace” item and icon.
//
// Deprecated: Use named icon &quot;edit-find-replace&quot;.
const STOCK_FIND_AND_REPLACE = "gtk-find-and-replace"

// STOCK_FLOPPY: “Floppy” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_FLOPPY = "gtk-floppy"

// STOCK_GOTO_BOTTOM: “Bottom” item and icon.
//
// Deprecated: Use named icon &quot;go-bottom&quot;.
const STOCK_GOTO_BOTTOM = "gtk-goto-bottom"

// STOCK_GOTO_FIRST: “First” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-first&quot;.
const STOCK_GOTO_FIRST = "gtk-goto-first"

// STOCK_GOTO_LAST: “Last” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-last&quot;.
const STOCK_GOTO_LAST = "gtk-goto-last"

// STOCK_GOTO_TOP: “Top” item and icon.
//
// Deprecated: Use named icon &quot;go-top&quot;.
const STOCK_GOTO_TOP = "gtk-goto-top"

// STOCK_GO_BACK: “Back” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-previous&quot;.
const STOCK_GO_BACK = "gtk-go-back"

// STOCK_GO_DOWN: “Down” item and icon.
//
// Deprecated: Use named icon &quot;go-down&quot;.
const STOCK_GO_DOWN = "gtk-go-down"

// STOCK_GO_FORWARD: “Forward” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-next&quot;.
const STOCK_GO_FORWARD = "gtk-go-forward"

// STOCK_GO_UP: “Up” item and icon.
//
// Deprecated: Use named icon &quot;go-up&quot;.
const STOCK_GO_UP = "gtk-go-up"

// STOCK_HELP: “Help” item and icon.
//
// Deprecated: Use named icon &quot;help-browser&quot;.
const STOCK_HELP = "gtk-help"

// STOCK_HOME: “Home” item and icon.
//
// Deprecated: Use named icon &quot;go-home&quot;.
const STOCK_HOME = "gtk-home"

// STOCK_INDEX: “Index” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_INDEX = "gtk-index"

// STOCK_ITALIC: “Italic” item and icon.
//
// Deprecated: Use named icon &quot;format-text-italic&quot;.
const STOCK_ITALIC = "gtk-italic"

// STOCK_JUMP_TO: “Jump to” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-jump&quot;.
const STOCK_JUMP_TO = "gtk-jump-to"

// STOCK_JUSTIFY_CENTER: “Center” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-center&quot;.
const STOCK_JUSTIFY_CENTER = "gtk-justify-center"

// STOCK_JUSTIFY_FILL: “Fill” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-fill&quot;.
const STOCK_JUSTIFY_FILL = "gtk-justify-fill"

// STOCK_JUSTIFY_LEFT: “Left” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-left&quot;.
const STOCK_JUSTIFY_LEFT = "gtk-justify-left"

// STOCK_JUSTIFY_RIGHT: “Right” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-right&quot;.
const STOCK_JUSTIFY_RIGHT = "gtk-justify-right"

// STOCK_MISSING_IMAGE: “Missing image” icon.
//
// Deprecated: Use named icon &quot;image-missing&quot;.
const STOCK_MISSING_IMAGE = "gtk-missing-image"

// STOCK_NEW: “New” item and icon.
//
// Deprecated: Use named icon &quot;document-new&quot; or the label
// &quot;_New&quot;.
const STOCK_NEW = "gtk-new"

// STOCK_NO: “No” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_NO = "gtk-no"

// STOCK_OK: “OK” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_OK&quot;.
const STOCK_OK = "gtk-ok"

// STOCK_OPEN: “Open” item and icon.
//
// Deprecated: Use named icon &quot;document-open&quot; or the label
// &quot;_Open&quot;.
const STOCK_OPEN = "gtk-open"

// STOCK_PASTE: “Paste” item and icon.
//
// Deprecated: Use named icon &quot;edit-paste&quot; or the label
// &quot;_Paste&quot;.
const STOCK_PASTE = "gtk-paste"

// STOCK_PREFERENCES: “Preferences” item and icon.
//
// Deprecated: Use named icon &quot;preferences-system&quot; or the label
// &quot;_Preferences&quot;.
const STOCK_PREFERENCES = "gtk-preferences"

// STOCK_PRINT: “Print” item and icon.
//
// Deprecated: Use named icon &quot;document-print&quot; or the label
// &quot;_Print&quot;.
const STOCK_PRINT = "gtk-print"

// STOCK_PRINT_PREVIEW: “Print Preview” item and icon.
//
// Deprecated: Use label &quot;Pre_view&quot;.
const STOCK_PRINT_PREVIEW = "gtk-print-preview"

// STOCK_PROPERTIES: “Properties” item and icon.
//
// Deprecated: Use named icon &quot;document-properties&quot; or the label
// &quot;_Properties&quot;.
const STOCK_PROPERTIES = "gtk-properties"

// STOCK_QUIT: “Quit” item and icon.
//
// Deprecated: Use named icon &quot;application-exit&quot; or the label
// &quot;_Quit&quot;.
const STOCK_QUIT = "gtk-quit"

// STOCK_REDO: “Redo” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;edit-redo&quot; or the label
// &quot;_Redo&quot;.
const STOCK_REDO = "gtk-redo"

// STOCK_REFRESH: “Refresh” item and icon.
//
// Deprecated: Use named icon &quot;view-refresh&quot; or the label
// &quot;_Refresh&quot;.
const STOCK_REFRESH = "gtk-refresh"

// STOCK_REMOVE: “Remove” item and icon.
//
// Deprecated: Use named icon &quot;list-remove&quot; or the label
// &quot;_Remove&quot;.
const STOCK_REMOVE = "gtk-remove"

// STOCK_REVERT_TO_SAVED: “Revert” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;document-revert&quot; or the label
// &quot;_Revert&quot;.
const STOCK_REVERT_TO_SAVED = "gtk-revert-to-saved"

// STOCK_SAVE: “Save” item and icon.
//
// Deprecated: Use named icon &quot;document-save&quot; or the label
// &quot;_Save&quot;.
const STOCK_SAVE = "gtk-save"

// STOCK_SAVE_AS: “Save As” item and icon.
//
// Deprecated: Use named icon &quot;document-save-as&quot; or the label
// &quot;Save _As&quot;.
const STOCK_SAVE_AS = "gtk-save-as"

// STOCK_SELECT_COLOR: “Color” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_SELECT_COLOR = "gtk-select-color"

// STOCK_SELECT_FONT: “Font” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_SELECT_FONT = "gtk-select-font"

// STOCK_SORT_ASCENDING: “Ascending” item and icon.
//
// Deprecated: Use named icon &quot;view-sort-ascending&quot;.
const STOCK_SORT_ASCENDING = "gtk-sort-ascending"

// STOCK_SORT_DESCENDING: “Descending” item and icon.
//
// Deprecated: Use named icon &quot;view-sort-descending&quot;.
const STOCK_SORT_DESCENDING = "gtk-sort-descending"

// STOCK_SPELL_CHECK: “Spell Check” item and icon.
//
// Deprecated: Use named icon &quot;tools-check-spelling&quot;.
const STOCK_SPELL_CHECK = "gtk-spell-check"

// STOCK_STOP: “Stop” item and icon.
//
// Deprecated: Use named icon &quot;process-stop&quot; or the label
// &quot;_Stop&quot;.
const STOCK_STOP = "gtk-stop"

// STOCK_STRIKETHROUGH: “Strikethrough” item and icon.
//
// Deprecated: Use named icon &quot;format-text-strikethrough&quot; or the label
// &quot;_Strikethrough&quot;.
const STOCK_STRIKETHROUGH = "gtk-strikethrough"

// STOCK_UNDELETE: “Undelete” item and icon. The icon has an RTL variant.
//
// Deprecated: since version 3.10.
const STOCK_UNDELETE = "gtk-undelete"

// STOCK_UNDERLINE: “Underline” item and icon.
//
// Deprecated: Use named icon &quot;format-text-underline&quot; or the label
// &quot;_Underline&quot;.
const STOCK_UNDERLINE = "gtk-underline"

// STOCK_UNDO: “Undo” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;edit-undo&quot; or the label
// &quot;_Undo&quot;.
const STOCK_UNDO = "gtk-undo"

// STOCK_YES: “Yes” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_YES = "gtk-yes"

// STOCK_ZOOM_100: “Zoom 100%” item and icon.
//
// Deprecated: Use named icon &quot;zoom-original&quot; or the label
// &quot;_Normal Size&quot;.
const STOCK_ZOOM_100 = "gtk-zoom-100"

// STOCK_ZOOM_FIT: “Zoom to Fit” item and icon.
//
// Deprecated: Use named icon &quot;zoom-fit-best&quot; or the label &quot;Best
// _Fit&quot;.
const STOCK_ZOOM_FIT = "gtk-zoom-fit"

// STOCK_ZOOM_IN: “Zoom In” item and icon.
//
// Deprecated: Use named icon &quot;zoom-in&quot; or the label &quot;Zoom
// _In&quot;.
const STOCK_ZOOM_IN = "gtk-zoom-in"

// STOCK_ZOOM_OUT: “Zoom Out” item and icon.
//
// Deprecated: Use named icon &quot;zoom-out&quot; or the label &quot;Zoom
// _Out&quot;.
const STOCK_ZOOM_OUT = "gtk-zoom-out"

type Stock = string

// TranslateFunc: function used to translate messages in e.g. IconFactory and
// ActionGroup.
//
// Deprecated: since version 3.10.
type TranslateFunc func(path string) (utf8 string)

// StockItem: deprecated: since version 3.10.
//
// An instance of this type is always passed by reference.
type StockItem struct {
	*stockItem
}

// stockItem is the struct that's finalized.
type stockItem struct {
	native unsafe.Pointer
}

var GIRInfoStockItem = girepository.MustFind("Gtk", "StockItem")
