package internal

import "io"

type CloseableReader struct {
	io.Reader
	closeable io.ReadCloser
}

func (cr CloseableReader) Close() error {
	return cr.closeable.Close()
}
