// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// CSSParserError errors that can occur while parsing CSS.
//
// These errors are unexpected and will cause parts of the given CSS to be
// ignored.
type CSSParserError C.gint

const (
	// CSSParserErrorFailed: unknown failure.
	CSSParserErrorFailed CSSParserError = iota
	// CSSParserErrorSyntax: given text does not form valid syntax.
	CSSParserErrorSyntax
	// CSSParserErrorImport: failed to import a resource.
	CSSParserErrorImport
	// CSSParserErrorName: given name has not been defined.
	CSSParserErrorName
	// CSSParserErrorUnknownValue: given value is not correct.
	CSSParserErrorUnknownValue
)

// String returns the name in string for CSSParserError.
func (c CSSParserError) String() string {
	switch c {
	case CSSParserErrorFailed:
		return "Failed"
	case CSSParserErrorSyntax:
		return "Syntax"
	case CSSParserErrorImport:
		return "Import"
	case CSSParserErrorName:
		return "Name"
	case CSSParserErrorUnknownValue:
		return "UnknownValue"
	default:
		return fmt.Sprintf("CSSParserError(%d)", c)
	}
}

// CSSParserWarning warnings that can occur while parsing CSS.
//
// Unlike CssParserErrors, warnings do not cause the parser to skip any input,
// but they indicate issues that should be fixed.
type CSSParserWarning C.gint

const (
	// CSSParserWarningDeprecated: given construct is deprecated and will be
	// removed in a future version.
	CSSParserWarningDeprecated CSSParserWarning = iota
	// CSSParserWarningSyntax: syntax construct was used that should be avoided.
	CSSParserWarningSyntax
	// CSSParserWarningUnimplemented: feature is not implemented.
	CSSParserWarningUnimplemented
)

// String returns the name in string for CSSParserWarning.
func (c CSSParserWarning) String() string {
	switch c {
	case CSSParserWarningDeprecated:
		return "Deprecated"
	case CSSParserWarningSyntax:
		return "Syntax"
	case CSSParserWarningUnimplemented:
		return "Unimplemented"
	default:
		return fmt.Sprintf("CSSParserWarning(%d)", c)
	}
}
