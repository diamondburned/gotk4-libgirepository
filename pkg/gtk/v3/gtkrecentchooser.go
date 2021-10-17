// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gint _gotk4_gtk3_RecentSortFunc(GtkRecentInfo*, GtkRecentInfo*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_chooser_error_get_type()), F: marshalRecentChooserError},
		{T: externglib.Type(C.gtk_recent_sort_type_get_type()), F: marshalRecentSortType},
		{T: externglib.Type(C.gtk_recent_chooser_get_type()), F: marshalRecentChooserer},
	})
}

// RecentChooserError: these identify the various errors that can occur while
// calling RecentChooser functions.
type RecentChooserError int

const (
	// RecentChooserErrorNotFound indicates that a file does not exist.
	RecentChooserErrorNotFound RecentChooserError = iota
	// RecentChooserErrorInvalidURI indicates a malformed URI.
	RecentChooserErrorInvalidURI
)

func marshalRecentChooserError(p uintptr) (interface{}, error) {
	return RecentChooserError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentChooserError.
func (r RecentChooserError) String() string {
	switch r {
	case RecentChooserErrorNotFound:
		return "NotFound"
	case RecentChooserErrorInvalidURI:
		return "InvalidURI"
	default:
		return fmt.Sprintf("RecentChooserError(%d)", r)
	}
}

// RecentSortType: used to specify the sorting method to be applyed to the
// recently used resource list.
type RecentSortType int

const (
	// RecentSortNone: do not sort the returned list of recently used resources.
	RecentSortNone RecentSortType = iota
	// RecentSortMru: sort the returned list with the most recently used items
	// first.
	RecentSortMru
	// RecentSortLru: sort the returned list with the least recently used items
	// first.
	RecentSortLru
	// RecentSortCustom: sort the returned list using a custom sorting function
	// passed using gtk_recent_chooser_set_sort_func().
	RecentSortCustom
)

func marshalRecentSortType(p uintptr) (interface{}, error) {
	return RecentSortType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentSortType.
func (r RecentSortType) String() string {
	switch r {
	case RecentSortNone:
		return "None"
	case RecentSortMru:
		return "Mru"
	case RecentSortLru:
		return "Lru"
	case RecentSortCustom:
		return "Custom"
	default:
		return fmt.Sprintf("RecentSortType(%d)", r)
	}
}

type RecentSortFunc func(a, b *RecentInfo) (gint int)

//export _gotk4_gtk3_RecentSortFunc
func _gotk4_gtk3_RecentSortFunc(arg0 *C.GtkRecentInfo, arg1 *C.GtkRecentInfo, arg2 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var a *RecentInfo // out
	var b *RecentInfo // out

	a = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(arg0)))
	C.gtk_recent_info_ref(arg0)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(a)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)
	b = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_recent_info_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(b)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)

	fn := v.(RecentSortFunc)
	gint := fn(a, b)

	cret = C.gint(gint)

	return cret
}

// RecentChooserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RecentChooserOverrider interface {
	// AddFilter adds filter to the list of RecentFilter objects held by
	// chooser.
	//
	// If no previous filter objects were defined, this function will call
	// gtk_recent_chooser_set_filter().
	AddFilter(filter *RecentFilter)
	// CurrentURI gets the URI currently selected by chooser.
	CurrentURI() string
	// Items gets the list of recently used resources in form of RecentInfo
	// objects.
	//
	// The return value of this function is affected by the “sort-type” and
	// “limit” properties of chooser.
	Items() []*RecentInfo
	ItemActivated()
	// ListFilters gets the RecentFilter objects held by chooser.
	ListFilters() []RecentFilter
	// RemoveFilter removes filter from the list of RecentFilter objects held by
	// chooser.
	RemoveFilter(filter *RecentFilter)
	// SelectAll selects all the items inside chooser, if the chooser supports
	// multiple selection.
	SelectAll()
	// SelectURI selects uri inside chooser.
	SelectURI(uri string) error
	SelectionChanged()
	// SetCurrentURI sets uri as the current URI for chooser.
	SetCurrentURI(uri string) error
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func. If the chooser has the sort type set to K_RECENT_SORT_CUSTOM
	// then the chooser will sort using this function.
	//
	// To the comparison function will be passed two RecentInfo structs and
	// sort_data; sort_func should return a positive integer if the first item
	// comes before the second, zero if the two items are equal and a negative
	// integer if the first item comes after the second.
	SetSortFunc(sortFunc RecentSortFunc)
	// UnselectAll unselects all the items inside chooser.
	UnselectAll()
	// UnselectURI unselects uri inside chooser.
	UnselectURI(uri string)
}

// RecentChooser is an interface that can be implemented by widgets displaying
// the list of recently used files. In GTK+, the main objects that implement
// this interface are RecentChooserWidget, RecentChooserDialog and
// RecentChooserMenu.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooser struct {
	*externglib.Object
}

// RecentChooserer describes RecentChooser's abstract methods.
type RecentChooserer interface {
	externglib.Objector

	// AddFilter adds filter to the list of RecentFilter objects held by
	// chooser.
	AddFilter(filter *RecentFilter)
	// CurrentItem gets the RecentInfo currently selected by chooser.
	CurrentItem() *RecentInfo
	// CurrentURI gets the URI currently selected by chooser.
	CurrentURI() string
	// Filter gets the RecentFilter object currently used by chooser to affect
	// the display of the recently used resources.
	Filter() *RecentFilter
	// Items gets the list of recently used resources in form of RecentInfo
	// objects.
	Items() []*RecentInfo
	// Limit gets the number of items returned by gtk_recent_chooser_get_items()
	// and gtk_recent_chooser_get_uris().
	Limit() int
	// LocalOnly gets whether only local resources should be shown in the
	// recently used resources selector.
	LocalOnly() bool
	// SelectMultiple gets whether chooser can select multiple items.
	SelectMultiple() bool
	// ShowIcons retrieves whether chooser should show an icon near the
	// resource.
	ShowIcons() bool
	// ShowNotFound retrieves whether chooser should show the recently used
	// resources that were not found.
	ShowNotFound() bool
	// ShowPrivate returns whether chooser should display recently used
	// resources registered as private.
	ShowPrivate() bool
	// ShowTips gets whether chooser should display tooltips containing the full
	// path of a recently user resource.
	ShowTips() bool
	// SortType gets the value set by gtk_recent_chooser_set_sort_type().
	SortType() RecentSortType
	// URIs gets the URI of the recently used resources.
	URIs() []string
	// ListFilters gets the RecentFilter objects held by chooser.
	ListFilters() []RecentFilter
	// RemoveFilter removes filter from the list of RecentFilter objects held by
	// chooser.
	RemoveFilter(filter *RecentFilter)
	// SelectAll selects all the items inside chooser, if the chooser supports
	// multiple selection.
	SelectAll()
	// SelectURI selects uri inside chooser.
	SelectURI(uri string) error
	// SetCurrentURI sets uri as the current URI for chooser.
	SetCurrentURI(uri string) error
	// SetFilter sets filter as the current RecentFilter object used by chooser
	// to affect the displayed recently used resources.
	SetFilter(filter *RecentFilter)
	// SetLimit sets the number of items that should be returned by
	// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
	SetLimit(limit int)
	// SetLocalOnly sets whether only local resources, that is resources using
	// the file:// URI scheme, should be shown in the recently used resources
	// selector.
	SetLocalOnly(localOnly bool)
	// SetSelectMultiple sets whether chooser can select multiple items.
	SetSelectMultiple(selectMultiple bool)
	// SetShowIcons sets whether chooser should show an icon near the resource
	// when displaying it.
	SetShowIcons(showIcons bool)
	// SetShowNotFound sets whether chooser should display the recently used
	// resources that it didn’t find.
	SetShowNotFound(showNotFound bool)
	// SetShowPrivate: whether to show recently used resources marked registered
	// as private.
	SetShowPrivate(showPrivate bool)
	// SetShowTips sets whether to show a tooltips containing the full path of
	// each recently used resource in a RecentChooser widget.
	SetShowTips(showTips bool)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortFunc RecentSortFunc)
	// SetSortType changes the sorting order of the recently used resources list
	// displayed by chooser.
	SetSortType(sortType RecentSortType)
	// UnselectAll unselects all the items inside chooser.
	UnselectAll()
	// UnselectURI unselects uri inside chooser.
	UnselectURI(uri string)
}

var _ RecentChooserer = (*RecentChooser)(nil)

func wrapRecentChooser(obj *externglib.Object) *RecentChooser {
	return &RecentChooser{
		Object: obj,
	}
}

func marshalRecentChooserer(p uintptr) (interface{}, error) {
	return wrapRecentChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddFilter adds filter to the list of RecentFilter objects held by chooser.
//
// If no previous filter objects were defined, this function will call
// gtk_recent_chooser_set_filter().
//
// The function takes the following parameters:
//
//    - filter: RecentFilter.
//
func (chooser *RecentChooser) AddFilter(filter *RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_chooser_add_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// CurrentItem gets the RecentInfo currently selected by chooser.
func (chooser *RecentChooser) CurrentItem() *RecentInfo {
	var _arg0 *C.GtkRecentChooser // out
	var _cret *C.GtkRecentInfo    // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_current_item(_arg0)
	runtime.KeepAlive(chooser)

	var _recentInfo *RecentInfo // out

	_recentInfo = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_recentInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)

	return _recentInfo
}

// CurrentURI gets the URI currently selected by chooser.
func (chooser *RecentChooser) CurrentURI() string {
	var _arg0 *C.GtkRecentChooser // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_current_uri(_arg0)
	runtime.KeepAlive(chooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Filter gets the RecentFilter object currently used by chooser to affect the
// display of the recently used resources.
func (chooser *RecentChooser) Filter() *RecentFilter {
	var _arg0 *C.GtkRecentChooser // out
	var _cret *C.GtkRecentFilter  // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_filter(_arg0)
	runtime.KeepAlive(chooser)

	var _recentFilter *RecentFilter // out

	_recentFilter = wrapRecentFilter(externglib.Take(unsafe.Pointer(_cret)))

	return _recentFilter
}

// Items gets the list of recently used resources in form of RecentInfo objects.
//
// The return value of this function is affected by the “sort-type” and “limit”
// properties of chooser.
func (chooser *RecentChooser) Items() []*RecentInfo {
	var _arg0 *C.GtkRecentChooser // out
	var _cret *C.GList            // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_items(_arg0)
	runtime.KeepAlive(chooser)

	var _list []*RecentInfo // out

	_list = make([]*RecentInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkRecentInfo)(v)
		var dst *RecentInfo // out
		dst = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// Limit gets the number of items returned by gtk_recent_chooser_get_items() and
// gtk_recent_chooser_get_uris().
func (chooser *RecentChooser) Limit() int {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_limit(_arg0)
	runtime.KeepAlive(chooser)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// LocalOnly gets whether only local resources should be shown in the recently
// used resources selector. See gtk_recent_chooser_set_local_only().
func (chooser *RecentChooser) LocalOnly() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_local_only(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectMultiple gets whether chooser can select multiple items.
func (chooser *RecentChooser) SelectMultiple() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_select_multiple(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowIcons retrieves whether chooser should show an icon near the resource.
func (chooser *RecentChooser) ShowIcons() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_show_icons(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowNotFound retrieves whether chooser should show the recently used
// resources that were not found.
func (chooser *RecentChooser) ShowNotFound() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_show_not_found(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowPrivate returns whether chooser should display recently used resources
// registered as private.
func (chooser *RecentChooser) ShowPrivate() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_show_private(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowTips gets whether chooser should display tooltips containing the full
// path of a recently user resource.
func (chooser *RecentChooser) ShowTips() bool {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_show_tips(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SortType gets the value set by gtk_recent_chooser_set_sort_type().
func (chooser *RecentChooser) SortType() RecentSortType {
	var _arg0 *C.GtkRecentChooser // out
	var _cret C.GtkRecentSortType // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_sort_type(_arg0)
	runtime.KeepAlive(chooser)

	var _recentSortType RecentSortType // out

	_recentSortType = RecentSortType(_cret)

	return _recentSortType
}

// URIs gets the URI of the recently used resources.
//
// The return value of this function is affected by the “sort-type” and “limit”
// properties of chooser.
//
// Since the returned array is NULL terminated, length may be NULL.
func (chooser *RecentChooser) URIs() []string {
	var _arg0 *C.GtkRecentChooser // out
	var _cret **C.gchar           // in
	var _arg1 C.gsize             // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_get_uris(_arg0, &_arg1)
	runtime.KeepAlive(chooser)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice(_cret, _arg1)
		_utf8s = make([]string, _arg1)
		for i := 0; i < int(_arg1); i++ {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ListFilters gets the RecentFilter objects held by chooser.
func (chooser *RecentChooser) ListFilters() []RecentFilter {
	var _arg0 *C.GtkRecentChooser // out
	var _cret *C.GSList           // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_recent_chooser_list_filters(_arg0)
	runtime.KeepAlive(chooser)

	var _sList []RecentFilter // out

	_sList = make([]RecentFilter, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkRecentFilter)(v)
		var dst RecentFilter // out
		dst = *wrapRecentFilter(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveFilter removes filter from the list of RecentFilter objects held by
// chooser.
//
// The function takes the following parameters:
//
//    - filter: RecentFilter.
//
func (chooser *RecentChooser) RemoveFilter(filter *RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_chooser_remove_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// SelectAll selects all the items inside chooser, if the chooser supports
// multiple selection.
func (chooser *RecentChooser) SelectAll() {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	C.gtk_recent_chooser_select_all(_arg0)
	runtime.KeepAlive(chooser)
}

// SelectURI selects uri inside chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) SelectURI(uri string) error {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_chooser_select_uri(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCurrentURI sets uri as the current URI for chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) SetCurrentURI(uri string) error {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_chooser_set_current_uri(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetFilter sets filter as the current RecentFilter object used by chooser to
// affect the displayed recently used resources.
//
// The function takes the following parameters:
//
//    - filter: RecentFilter.
//
func (chooser *RecentChooser) SetFilter(filter *RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if filter != nil {
		_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	}

	C.gtk_recent_chooser_set_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// SetLimit sets the number of items that should be returned by
// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
//
// The function takes the following parameters:
//
//    - limit: positive integer, or -1 for all items.
//
func (chooser *RecentChooser) SetLimit(limit int) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gint              // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = C.gint(limit)

	C.gtk_recent_chooser_set_limit(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(limit)
}

// SetLocalOnly sets whether only local resources, that is resources using the
// file:// URI scheme, should be shown in the recently used resources selector.
// If local_only is TRUE (the default) then the shown resources are guaranteed
// to be accessible through the operating system native file system.
//
// The function takes the following parameters:
//
//    - localOnly: TRUE if only local files can be shown.
//
func (chooser *RecentChooser) SetLocalOnly(localOnly bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if localOnly {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_local_only(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(localOnly)
}

// SetSelectMultiple sets whether chooser can select multiple items.
//
// The function takes the following parameters:
//
//    - selectMultiple: TRUE if chooser can select more than one item.
//
func (chooser *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if selectMultiple {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_select_multiple(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(selectMultiple)
}

// SetShowIcons sets whether chooser should show an icon near the resource when
// displaying it.
//
// The function takes the following parameters:
//
//    - showIcons: whether to show an icon near the resource.
//
func (chooser *RecentChooser) SetShowIcons(showIcons bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if showIcons {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_show_icons(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showIcons)
}

// SetShowNotFound sets whether chooser should display the recently used
// resources that it didn’t find. This only applies to local resources.
//
// The function takes the following parameters:
//
//    - showNotFound: whether to show the local items we didn’t find.
//
func (chooser *RecentChooser) SetShowNotFound(showNotFound bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if showNotFound {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_show_not_found(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showNotFound)
}

// SetShowPrivate: whether to show recently used resources marked registered as
// private.
//
// The function takes the following parameters:
//
//    - showPrivate: TRUE to show private items, FALSE otherwise.
//
func (chooser *RecentChooser) SetShowPrivate(showPrivate bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if showPrivate {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_show_private(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showPrivate)
}

// SetShowTips sets whether to show a tooltips containing the full path of each
// recently used resource in a RecentChooser widget.
//
// The function takes the following parameters:
//
//    - showTips: TRUE if tooltips should be shown.
//
func (chooser *RecentChooser) SetShowTips(showTips bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	if showTips {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_set_show_tips(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showTips)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the chooser has the sort type set to K_RECENT_SORT_CUSTOM then the chooser
// will sort using this function.
//
// To the comparison function will be passed two RecentInfo structs and
// sort_data; sort_func should return a positive integer if the first item comes
// before the second, zero if the two items are equal and a negative integer if
// the first item comes after the second.
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (chooser *RecentChooser) SetSortFunc(sortFunc RecentSortFunc) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.GtkRecentSortFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_RecentSortFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_recent_chooser_set_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(sortFunc)
}

// SetSortType changes the sorting order of the recently used resources list
// displayed by chooser.
//
// The function takes the following parameters:
//
//    - sortType: sort order that the chooser should use.
//
func (chooser *RecentChooser) SetSortType(sortType RecentSortType) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.GtkRecentSortType // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = C.GtkRecentSortType(sortType)

	C.gtk_recent_chooser_set_sort_type(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(sortType)
}

// UnselectAll unselects all the items inside chooser.
func (chooser *RecentChooser) UnselectAll() {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))

	C.gtk_recent_chooser_unselect_all(_arg0)
	runtime.KeepAlive(chooser)
}

// UnselectURI unselects uri inside chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) UnselectURI(uri string) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_chooser_unselect_uri(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)
}

// ConnectItemActivated: this signal is emitted when the user "activates" a
// recent item in the recent chooser. This can happen by double-clicking on an
// item in the recently used resources list, or by pressing Enter.
func (chooser *RecentChooser) ConnectItemActivated(f func()) externglib.SignalHandle {
	return chooser.Connect("item-activated", f)
}

// ConnectSelectionChanged: this signal is emitted when there is a change in the
// set of selected recently used resources. This can happen when a user modifies
// the selection with the mouse or the keyboard, or when explicitly calling
// functions to change the selection.
func (chooser *RecentChooser) ConnectSelectionChanged(f func()) externglib.SignalHandle {
	return chooser.Connect("selection-changed", f)
}
