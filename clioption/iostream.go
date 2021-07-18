// https://github.com/kubernetes/cli-runtime/blob/master/pkg/genericclioptions/io_options.go
package clioption

import (
	"bytes"
	"io"
)

// IOStreams provides the standard names for iostreams.  This is useful for embedding and for unit testing.
type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

type DexOptions struct {
	RestartDex bool
}

// NewTestIOStreams returns a valid IOStreams and in, out, errout buffers for unit tests
func NewTestIOStreams() (IOStreams, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	return IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,
	}, in, out, errOut
}
