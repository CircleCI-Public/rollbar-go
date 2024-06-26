package rollbar_test

import (
	"errors"

	"github.com/CircleCI-Public/rollbar-go"
)

type CustomWrappingError struct {
	error
	wrapped error
}

func (e CustomWrappingError) GetWrappedError() error {
	return e.wrapped
}

func ExampleSetUnwrapper() {
	rollbar.SetUnwrapper(func(err error) error {
		// preserve the default behavior for other types of errors
		if unwrapped := rollbar.DefaultUnwrapper(err); unwrapped != nil {
			return unwrapped
		}

		var ex CustomWrappingError
		if errors.As(err, &ex) {
			return ex.GetWrappedError()
		}

		return nil
	})
}
