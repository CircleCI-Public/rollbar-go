package rollbar_test

import (
	"errors"
	"runtime"

	"github.com/CircleCI-Public/rollbar-go"
)

type CustomTraceError struct {
	error
	trace []runtime.Frame
}

func (e CustomTraceError) GetTrace() []runtime.Frame {
	return e.trace
}

func ExampleSetStackTracer() {
	rollbar.SetStackTracer(func(err error) ([]runtime.Frame, bool) {
		// preserve the default behavior for other types of errors
		if trace, ok := rollbar.DefaultStackTracer(err); ok {
			return trace, ok
		}

		var cerr CustomTraceError
		if errors.As(err, &cerr) {
			return cerr.GetTrace(), true
		}

		return nil, false
	})
}
