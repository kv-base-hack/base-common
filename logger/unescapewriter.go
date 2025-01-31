package logger

import (
	"bytes"
	"io"
)

// UnescapeWriter ...
type UnescapeWriter struct {
	w io.WriteCloser
}

var (
	eq = []byte{'\\', '"'}
	qq = []byte{'"'}
)

// Write ...
func (u *UnescapeWriter) Write(p []byte) (n int, err error) {
	nw := len(p)
	p = bytes.ReplaceAll(p, eq, qq)
	m, err := u.w.Write(p)
	if m == len(p) {
		return nw, err
	}
	return m, err
}

// CloseWrite ...
func (u UnescapeWriter) CloseWrite() error {
	return u.w.Close()
}
