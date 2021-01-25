package trace

import "io"

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) Tracer {
	return nil
}
