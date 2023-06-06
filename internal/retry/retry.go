// Package retry provides helpers for retrying.
//
// This package defines flexible interfaces for retrying Go functions that may
// be flakey or eventually consistent. It abstracts the "backoff" (how long to
// wait between tries) and "retry" (execute the function again) mechanisms for
// maximum flexibility. Furthermore, everything is an interface, so you can
// define your own implementations.
//
// The package is modeled after Go's built-in HTTP package, making it easy to
// customize the built-in backoff with your own custom logic. Additionally,
// callers specify which errors are retryable by wrapping them. This is helpful
// with complex operations where only certain results should retry.
package retry

import (
	"context"
	"errors"
	"time"
)

// Func is a function passed to retry.
type Func func(ctx context.Context) (interface{}, error)

type Error struct {
	err error
}

// RetryableError marks an error as retryable.
func RetryableError(err error) error {
	if err == nil {
		return nil
	}
	return &Error{err}
}

// Unwrap implements error wrapping.
func (e *Error) Unwrap() error {
	return e.err
}

// Error returns the error string.
func (e *Error) Error() string {
	if e.err == nil {
		return "retryable: <nil>"
	}
	return "retryable: " + e.err.Error()
}

// Do wrap a function with a backoff to retry. The provided context is the same
// context passed to the Func.
func Do(ctx context.Context, b Backoff, f Func) (interface{}, error) {
	for {
		ife, err := f(ctx)
		if err == nil {
			return ife, nil
		}

		// Not retryable
		var rErr *Error
		if !errors.As(err, &rErr) {
			return nil, err
		}

		next, stop := b.Next()
		if stop {
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(next):
			continue
		}
	}
}
