package serr

import (
	"errors"
	"fmt"
)

type serr struct {
	err error
	st  *stackTrace
}

var _ error = (*serr)(nil)
var _ StackTracer = (*serr)(nil)

func (w *serr) Error() string {
	return w.err.Error()
}

func (w *serr) StackTrace() *stackTrace {
	return w.st
}

func (w *serr) Unwrap() error {
	return w.err
}

func (w *serr) Is(err error) bool {
	return errors.Is(w.err, err)
}

func New(text string) error {
	return Wrap(errors.New(text))
}

func Errorf(format string, a ...any) error {
	return Wrap(fmt.Errorf(format, a...))
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	var e *serr
	if errors.As(err, &e) {
		return err
	}
	return &serr{
		err: err,
		st:  callers(),
	}
}
