// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expression_watch_get_type()), F: marshalExpressionWatch},
	})
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func()

//export gotk4_ExpressionNotify
func gotk4_ExpressionNotify(arg0 C.gpointer) {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ExpressionNotify)
	fn()
}

// ParamSpecExpression creates a new `GParamSpec` instance for a property
// holding a `GtkExpression`.
//
// See `g_param_spec_internal()` for details on the property strings.
func ParamSpecExpression(name string, nick string, blurb string, flags gobject.ParamFlags) gobject.ParamSpec {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 *C.char
	var arg4 C.GParamFlags

	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.char)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret gobject.ParamSpec

	cret = C.gtk_param_spec_expression(arg1, arg2, arg3, arg4)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gobject.ParamSpec)

	return goret
}

// ValueDupExpression retrieves the `GtkExpression` stored inside the given
// `value`, and acquires a reference to it.
func ValueDupExpression(value *externglib.Value) Expression {
	var arg1 *C.GValue

	arg1 = (*C.GValue)(value.GValue)

	cret := new(C.GtkExpression)
	var goret Expression

	cret = C.gtk_value_dup_expression(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Expression)

	return goret
}

// ValueGetExpression retrieves the `GtkExpression` stored inside the given
// `value`.
func ValueGetExpression(value *externglib.Value) Expression {
	var arg1 *C.GValue

	arg1 = (*C.GValue)(value.GValue)

	var cret *C.GtkExpression
	var goret Expression

	cret = C.gtk_value_get_expression(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Expression)

	return goret
}

// ValueSetExpression stores the given `GtkExpression` inside `value`.
//
// The `GValue` will acquire a reference to the `expression`.
func ValueSetExpression(value *externglib.Value, expression Expression) {
	var arg1 *C.GValue
	var arg2 *C.GtkExpression

	arg1 = (*C.GValue)(value.GValue)
	arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_value_set_expression(arg1, arg2)
}

// ValueTakeExpression stores the given `GtkExpression` inside `value`.
//
// This function transfers the ownership of the `expression` to the `GValue`.
func ValueTakeExpression(value *externglib.Value, expression Expression) {
	var arg1 *C.GValue
	var arg2 *C.GtkExpression

	arg1 = (*C.GValue)(value.GValue)
	arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_value_take_expression(arg1, arg2)
}

// ExpressionWatch: an opaque structure representing a watched `GtkExpression`.
//
// The contents of `GtkExpressionWatch` should only be accessed through the
// provided API.
type ExpressionWatch struct {
	native C.GtkExpressionWatch
}

// WrapExpressionWatch wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapExpressionWatch(ptr unsafe.Pointer) *ExpressionWatch {
	if ptr == nil {
		return nil
	}

	return (*ExpressionWatch)(ptr)
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapExpressionWatch(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *ExpressionWatch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Evaluate evaluates the watched expression and on success stores the result in
// `value`.
//
// This is equivalent to calling [method@Gtk.Expression.evaluate] with the
// expression and this pointer originally used to create `watch`.
func (w *ExpressionWatch) Evaluate(value *externglib.Value) bool {
	var arg0 *C.GtkExpressionWatch
	var arg1 *C.GValue

	arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))
	arg1 = (*C.GValue)(value.GValue)

	var cret C.gboolean
	var goret bool

	cret = C.gtk_expression_watch_evaluate(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Ref acquires a reference on the given `GtkExpressionWatch`.
func (w *ExpressionWatch) Ref() *ExpressionWatch {
	var arg0 *C.GtkExpressionWatch

	arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	cret := new(C.GtkExpressionWatch)
	var goret *ExpressionWatch

	cret = C.gtk_expression_watch_ref(arg0)

	goret = WrapExpressionWatch(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *ExpressionWatch) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Unref releases a reference on the given `GtkExpressionWatch`.
//
// If the reference was the last, the resources associated to `self` are freed.
func (w *ExpressionWatch) Unref() {
	var arg0 *C.GtkExpressionWatch

	arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	C.gtk_expression_watch_unref(arg0)
}

// Unwatch stops watching an expression.
//
// See [method@Gtk.Expression.watch] for how the watch was established.
func (w *ExpressionWatch) Unwatch() {
	var arg0 *C.GtkExpressionWatch

	arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	C.gtk_expression_watch_unwatch(arg0)
}
