package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events to code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTrace struct {}

func (t *nilTrace) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace
func Off() Tracer {
	return &nilTrace{}
}

func (t *tracer) Trace(a ...interface{}) {
	_, _ = fmt.Fprint(t.out, a...)
	_, _ = fmt.Fprintln(t.out)
}

func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}

