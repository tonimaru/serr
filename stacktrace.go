package serr

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type StackTracer interface {
	StackTrace() *stackTrace
}

type stackTrace []uintptr

func callers() *stackTrace {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	st := (stackTrace)(pcs[:n])
	return &st
}

func (s *stackTrace) format(st io.Writer, verb rune) {
	frames := runtime.CallersFrames(*s)
	var frame runtime.Frame
	more := true
	for {
		if !more {
			break
		}
		frame, more = frames.Next()
		if strings.Contains(frame.File, "/go/src/runtime/") {
			continue
		}
		fmt.Fprintf(st, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
	}
}

func (s *stackTrace) Format(st fmt.State, verb rune) {
	s.format(st, verb)
}

func (s *stackTrace) String() string {
	var sb strings.Builder
	s.format(&sb, 'v')
	return sb.String()
}
