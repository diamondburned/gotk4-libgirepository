// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_main_context_get_type()), F: marshalMainContext},
		{T: externglib.Type(C.g_main_loop_get_type()), F: marshalMainLoop},
	})
}

// SourceFunc specifies the type of function passed to g_timeout_add(),
// g_timeout_add_full(), g_idle_add(), and g_idle_add_full().
//
// When calling g_source_set_callback(), you may need to cast a function of a
// different type to this type. Use G_SOURCE_FUNC() to avoid warnings about
// incompatible function types.
type SourceFunc func() (ok bool)

//export gotk4_SourceFunc
func gotk4_SourceFunc(arg0 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(SourceFunc)
	ok := fn()

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// GetCurrentTime: equivalent to the UNIX gettimeofday() function, but portable.
//
// You may find g_get_real_time() to be more convenient.
func GetCurrentTime(result *TimeVal) {
	var _arg1 *C.GTimeVal // out

	_arg1 = (*C.GTimeVal)(unsafe.Pointer(result.Native()))

	C.g_get_current_time(_arg1)
}

// GetMonotonicTime queries the system monotonic time.
//
// The monotonic clock will always increase and doesn't suffer discontinuities
// when the user (or NTP) changes the system time. It may or may not continue to
// tick during times where the machine is suspended.
//
// We try to use the clock that corresponds as closely as possible to the
// passage of time as measured by system calls such as poll() but it may not
// always be possible to do this.
func GetMonotonicTime() int64 {
	var _cret C.gint64 // in

	_cret = C.g_get_monotonic_time()

	var _gint64 int64 // out

	_gint64 = (int64)(_cret)

	return _gint64
}

// GetRealTime queries the system wall-clock time.
//
// This call is functionally equivalent to g_get_current_time() except that the
// return value is often more convenient than dealing with a Val.
//
// You should only use this call if you are actually interested in the real
// wall-clock time. g_get_monotonic_time() is probably more useful for measuring
// intervals.
func GetRealTime() int64 {
	var _cret C.gint64 // in

	_cret = C.g_get_real_time()

	var _gint64 int64 // out

	_gint64 = (int64)(_cret)

	return _gint64
}

// IdleRemoveByData removes the idle function with the given data.
func IdleRemoveByData(data interface{}) bool {
	var _arg1 C.gpointer // out

	_arg1 = C.gpointer(data)

	var _cret C.gboolean // in

	_cret = C.g_idle_remove_by_data(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MainDepth returns the depth of the stack of calls to
// g_main_context_dispatch() on any Context in the current thread. That is, when
// called from the toplevel, it gives 0. When called from within a callback from
// g_main_context_iteration() (or g_main_loop_run(), etc.) it returns 1. When
// called from within a callback to a recursive call to
// g_main_context_iteration(), it returns 2. And so forth.
//
// This function is useful in a situation like the following: Imagine an
// extremely simple "garbage collected" system.
//
//    gpointer
//    allocate_memory (gsize size)
//    {
//      FreeListBlock *block = g_new (FreeListBlock, 1);
//      block->mem = g_malloc (size);
//      block->depth = g_main_depth ();
//      free_list = g_list_prepend (free_list, block);
//      return block->mem;
//    }
//
//    void
//    free_allocated_memory (void)
//    {
//      GList *l;
//
//      int depth = g_main_depth ();
//      for (l = free_list; l; );
//        {
//          GList *next = l->next;
//          FreeListBlock *block = l->data;
//          if (block->depth > depth)
//            {
//              g_free (block->mem);
//              g_free (block);
//              free_list = g_list_delete_link (free_list, l);
//            }
//
//          l = next;
//        }
//      }
//
// There is a temptation to use g_main_depth() to solve problems with
// reentrancy. For instance, while waiting for data to be received from the
// network in response to a menu item, the menu item might be selected again. It
// might seem that one could make the menu item's callback return immediately
// and do nothing if g_main_depth() returns a value greater than 1. However,
// this should be avoided since the user then sees selecting the menu item do
// nothing. Furthermore, you'll find yourself adding these checks all over your
// code, since there are doubtless many, many things that the user could do.
// Instead, you can use the following techniques:
//
// 1. Use gtk_widget_set_sensitive() or modal dialogs to prevent the user from
// interacting with elements while the main loop is recursing.
//
// 2. Avoid main loop recursion in situations where you can't handle arbitrary
// callbacks. Instead, structure your code so that you simply return to the main
// loop and then get called again when there is more work to do.
func MainDepth() int {
	var _cret C.gint // in

	_cret = C.g_main_depth()

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MainContext: the `GMainContext` struct is an opaque data type representing a
// set of sources to be handled in a main loop.
type MainContext struct {
	native C.GMainContext
}

// WrapMainContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMainContext(ptr unsafe.Pointer) *MainContext {
	if ptr == nil {
		return nil
	}

	return (*MainContext)(ptr)
}

func marshalMainContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMainContext(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MainContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Acquire tries to become the owner of the specified context. If some other
// thread is the owner of the context, returns false immediately. Ownership is
// properly recursive: the owner can require ownership again and will release
// ownership when g_main_context_release() is called as many times as
// g_main_context_acquire().
//
// You must be the owner of a context before you can call
// g_main_context_prepare(), g_main_context_query(), g_main_context_check(),
// g_main_context_dispatch().
func (c *MainContext) Acquire() bool {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_main_context_acquire(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddPoll adds a file descriptor to the set of file descriptors polled for this
// context. This will very seldom be used directly. Instead a typical event
// source will use g_source_add_unix_fd() instead.
func (c *MainContext) AddPoll(fd *PollFD, priority int) {
	var _arg0 *C.GMainContext // out
	var _arg1 *C.GPollFD      // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GPollFD)(unsafe.Pointer(fd.Native()))
	_arg2 = C.gint(priority)

	C.g_main_context_add_poll(_arg0, _arg1, _arg2)
}

// Check passes the results of polling back to the main loop. You should be
// careful to pass @fds and its length @n_fds as received from
// g_main_context_query(), as this functions relies on assumptions on how @fds
// is filled.
//
// You must have successfully acquired the context with g_main_context_acquire()
// before you may call this function.
func (c *MainContext) Check(maxPriority int, fds []PollFD) bool {
	var _arg0 *C.GMainContext // out
	var _arg1 C.gint          // out
	var _arg2 *C.GPollFD
	var _arg3 C.gint

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(maxPriority)
	_arg3 = C.gint(len(fds))
	_arg2 = (*C.GPollFD)(unsafe.Pointer(&fds[0]))

	var _cret C.gboolean // in

	_cret = C.g_main_context_check(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Dispatch dispatches all pending sources.
//
// You must have successfully acquired the context with g_main_context_acquire()
// before you may call this function.
func (c *MainContext) Dispatch() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_dispatch(_arg0)
}

// IsOwner determines whether this thread holds the (recursive) ownership of
// this Context. This is useful to know before waiting on another thread that
// may be blocking to get ownership of @context.
func (c *MainContext) IsOwner() bool {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_main_context_is_owner(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Iteration runs a single iteration for the given main loop. This involves
// checking to see if any event sources are ready to be processed, then if no
// events sources are ready and @may_block is true, waiting for a source to
// become ready, then dispatching the highest priority events sources that are
// ready. Otherwise, if @may_block is false sources are not waited to become
// ready, only those highest priority events sources will be dispatched (if
// any), that are ready at this given moment without further waiting.
//
// Note that even when @may_block is true, it is still possible for
// g_main_context_iteration() to return false, since the wait may be interrupted
// for other reasons than an event source becoming ready.
func (c *MainContext) Iteration(mayBlock bool) bool {
	var _arg0 *C.GMainContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))
	if mayBlock {
		_arg1 = C.TRUE
	}

	var _cret C.gboolean // in

	_cret = C.g_main_context_iteration(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Pending checks if any sources have pending events for the given context.
func (c *MainContext) Pending() bool {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_main_context_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopThreadDefault pops @context off the thread-default context stack
// (verifying that it was on the top of the stack).
func (c *MainContext) PopThreadDefault() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_pop_thread_default(_arg0)
}

// Prepare prepares to poll sources within a main loop. The resulting
// information for polling is determined by calling g_main_context_query ().
//
// You must have successfully acquired the context with g_main_context_acquire()
// before you may call this function.
func (c *MainContext) Prepare() (int, bool) {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	var _arg1 C.gint     // in
	var _cret C.gboolean // in

	_cret = C.g_main_context_prepare(_arg0, &_arg1)

	var _priority int // out
	var _ok bool      // out

	_priority = (int)(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _priority, _ok
}

// PushThreadDefault acquires @context and sets it as the thread-default context
// for the current thread. This will cause certain asynchronous operations (such
// as most [gio][gio]-based I/O) which are started in this thread to run under
// @context and deliver their results to its main loop, rather than running
// under the global default context in the main thread. Note that calling this
// function changes the context returned by g_main_context_get_thread_default(),
// not the one returned by g_main_context_default(), so it does not affect the
// context used by functions like g_idle_add().
//
// Normally you would call this function shortly after creating a new thread,
// passing it a Context which will be run by a Loop in that thread, to set a new
// default context for all async operations in that thread. In this case you may
// not need to ever call g_main_context_pop_thread_default(), assuming you want
// the new Context to be the default for the whole lifecycle of the thread.
//
// If you don't have control over how the new thread was created (e.g. in the
// new thread isn't newly created, or if the thread life cycle is managed by a
// Pool), it is always suggested to wrap the logic that needs to use the new
// Context inside a g_main_context_push_thread_default() /
// g_main_context_pop_thread_default() pair, otherwise threads that are re-used
// will end up never explicitly releasing the Context reference they hold.
//
// In some cases you may want to schedule a single operation in a non-default
// context, or temporarily use a non-default context in the main thread. In that
// case, you can wrap the call to the asynchronous operation inside a
// g_main_context_push_thread_default() / g_main_context_pop_thread_default()
// pair, but it is up to you to ensure that no other asynchronous operations
// accidentally get started while the non-default context is active.
//
// Beware that libraries that predate this function may not correctly handle
// being used from a thread with a thread-default context. Eg, see
// g_file_supports_thread_contexts().
func (c *MainContext) PushThreadDefault() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_push_thread_default(_arg0)
}

// Release releases ownership of a context previously acquired by this thread
// with g_main_context_acquire(). If the context was acquired multiple times,
// the ownership will be released only when g_main_context_release() is called
// as many times as it was acquired.
func (c *MainContext) Release() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_release(_arg0)
}

// RemovePoll removes file descriptor from the set of file descriptors to be
// polled for a particular context.
func (c *MainContext) RemovePoll(fd *PollFD) {
	var _arg0 *C.GMainContext // out
	var _arg1 *C.GPollFD      // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GPollFD)(unsafe.Pointer(fd.Native()))

	C.g_main_context_remove_poll(_arg0, _arg1)
}

// Unref decreases the reference count on a Context object by one. If the result
// is zero, free the context and free all associated memory.
func (c *MainContext) Unref() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_unref(_arg0)
}

// Wakeup: if @context is currently blocking in g_main_context_iteration()
// waiting for a source to become ready, cause it to stop blocking and return.
// Otherwise, cause the next invocation of g_main_context_iteration() to return
// without blocking.
//
// This API is useful for low-level control over Context; for example,
// integrating it with main loop implementations such as Loop.
//
// Another related use for this function is when implementing a main loop with a
// termination condition, computed from multiple threads:
//
//      perform_work();
//
//      if (g_atomic_int_dec_and_test (&tasks_remaining))
//        g_main_context_wakeup (NULL);
func (c *MainContext) Wakeup() {
	var _arg0 *C.GMainContext // out

	_arg0 = (*C.GMainContext)(unsafe.Pointer(c.Native()))

	C.g_main_context_wakeup(_arg0)
}

// MainLoop: the `GMainLoop` struct is an opaque data type representing the main
// event loop of a GLib or GTK+ application.
type MainLoop struct {
	native C.GMainLoop
}

// WrapMainLoop wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMainLoop(ptr unsafe.Pointer) *MainLoop {
	if ptr == nil {
		return nil
	}

	return (*MainLoop)(ptr)
}

func marshalMainLoop(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMainLoop(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MainLoop) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// IsRunning checks to see if the main loop is currently being run via
// g_main_loop_run().
func (l *MainLoop) IsRunning() bool {
	var _arg0 *C.GMainLoop // out

	_arg0 = (*C.GMainLoop)(unsafe.Pointer(l.Native()))

	var _cret C.gboolean // in

	_cret = C.g_main_loop_is_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Quit stops a Loop from running. Any calls to g_main_loop_run() for the loop
// will return.
//
// Note that sources that have already been dispatched when g_main_loop_quit()
// is called will still be executed.
func (l *MainLoop) Quit() {
	var _arg0 *C.GMainLoop // out

	_arg0 = (*C.GMainLoop)(unsafe.Pointer(l.Native()))

	C.g_main_loop_quit(_arg0)
}

// Run runs a main loop until g_main_loop_quit() is called on the loop. If this
// is called for the thread of the loop's Context, it will process events from
// the loop, otherwise it will simply wait.
func (l *MainLoop) Run() {
	var _arg0 *C.GMainLoop // out

	_arg0 = (*C.GMainLoop)(unsafe.Pointer(l.Native()))

	C.g_main_loop_run(_arg0)
}

// Unref decreases the reference count on a Loop object by one. If the result is
// zero, free the loop and free all associated memory.
func (l *MainLoop) Unref() {
	var _arg0 *C.GMainLoop // out

	_arg0 = (*C.GMainLoop)(unsafe.Pointer(l.Native()))

	C.g_main_loop_unref(_arg0)
}
