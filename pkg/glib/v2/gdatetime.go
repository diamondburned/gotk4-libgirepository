// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"time"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// TimeSpan: value representing an interval of time, in microseconds.
type TimeSpan = int64

// NewTimeZoneFromGo creates a new TimeZone instance from Go's Location.
// The location's accuracy is down to the second.
func NewTimeZoneFromGo(loc *time.Location) *TimeZone {
	switch loc {
	case time.UTC:
		return NewTimeZoneUTC()
	case time.Local:
		return NewTimeZoneLocal()
	}

	t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	t2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	return NewTimeZoneOffset(int32(t2.Sub(t1) / time.Second))
}

// NewDateTimeFromGo creates a new DateTime instance from Go's Time. The
// TimeZone of the DateTime will be implicitly converted from the Time.
func NewDateTimeFromGo(t time.Time) *DateTime {
	tz := NewTimeZoneFromGo(t.Location())

	Y, M, D := t.Date()
	h, m, s := t.Clock()

	// Second offset within a minute in nanoseconds.
	seconds := (time.Duration(s) * time.Second) + time.Duration(t.Nanosecond())

	return NewDateTime(tz, int(Y), int(M), int(D), int(h), int(m), seconds.Seconds())
}
