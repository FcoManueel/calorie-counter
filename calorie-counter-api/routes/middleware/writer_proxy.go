package middleware

// Copied from the previous goji version: https://github.com/zenazn/goji/blob/master/web/mutil/writer_proxy.go
import (
	"io"
	"net/http"
)

// WriterProxy is a proxy around an http.ResponseWriter that allows you to hook
// into various parts of the response process.
type WriterProxy interface {
	http.ResponseWriter
	// Status returns the HTTP status of the request, or 0 if one has not
	// yet been sent.
	Status() int
	// BytesWritten returns the total number of bytes sent to the client.
	BytesWritten() int
	// Tee causes the response body to be written to the given io.Writer in
	// addition to proxying the writes through. Only one io.Writer can be
	// tee'd to at once: setting a second one will overwrite the first.
	// Writes will be sent to the proxy before being written to this
	// io.Writer. It is illegal for the tee'd writer to be modified
	// concurrently with writes.
	Tee(io.Writer)
	// Unwrap returns the original proxied target.
	Unwrap() http.ResponseWriter
}

// WrapWriter wraps an http.ResponseWriter, returning a proxy that allows you to
// hook into various parts of the response process.
func WrapWriter(w http.ResponseWriter) WriterProxy {
	return &basicWriter{ResponseWriter: w}
}

// basicWriter wraps a http.ResponseWriter that implements the minimal
// http.ResponseWriter interface.
type basicWriter struct {
	http.ResponseWriter
	wroteHeader bool
	code        int
	bytes       int
	tee         io.Writer
}

func (b *basicWriter) WriteHeader(code int) {
	if !b.wroteHeader {
		b.code = code
		b.wroteHeader = true
		b.ResponseWriter.WriteHeader(code)
	}
}
func (b *basicWriter) Write(buf []byte) (int, error) {
	b.WriteHeader(http.StatusOK)
	n, err := b.ResponseWriter.Write(buf)
	if b.tee != nil {
		_, err2 := b.tee.Write(buf[:n])
		// Prefer errors generated by the proxied writer.
		if err == nil {
			err = err2
		}
	}
	b.bytes += n
	return n, err
}
func (b *basicWriter) maybeWriteHeader() {
	if !b.wroteHeader {
		b.WriteHeader(http.StatusOK)
	}
}
func (b *basicWriter) Status() int {
	return b.code
}
func (b *basicWriter) BytesWritten() int {
	return b.bytes
}
func (b *basicWriter) Tee(w io.Writer) {
	b.tee = w
}
func (b *basicWriter) Unwrap() http.ResponseWriter {
	return b.ResponseWriter
}
