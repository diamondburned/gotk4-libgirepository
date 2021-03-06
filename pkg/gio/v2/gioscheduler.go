// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// IOSchedulerCancelAllJobs cancels all cancellable I/O jobs.
//
// A job is cancellable if a #GCancellable was passed into
// g_io_scheduler_push_job().
//
// Deprecated: You should never call this function, since you don't know how
// other libraries in your program might be making use of gioscheduler.
func IOSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()
}
