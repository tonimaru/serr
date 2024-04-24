package serr

import (
	"errors"
	"fmt"
)

type wrapped struct {
	err error
	st  *stackTrace
}

var _ error = (*wrapped)(nil)
var _ StackTracer = (*wrapped)(nil)

func (w *wrapped) Error() string {
	return w.err.Error()
}

func (w *wrapped) StackTrace() *stackTrace {
	return w.st
}

func (w *wrapped) Unwrap() error {
	return w.err
}

func (w *wrapped) Is(err error) bool {
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
	var e *wrapped
	if errors.As(err, &e) {
		return err
	}
	return &wrapped{
		err: err,
		st:  callers(),
	}
}

func StackTrace(err error) (*stackTrace, bool) {
	var st StackTracer
	if errors.As(err, &st) {
		return st.StackTrace(), true
	}
	return nil, false
}
