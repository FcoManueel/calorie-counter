package middleware

import (
	"log"
	"net/http"
	"time"

	"bytes"
	"goji.io"
	"golang.org/x/net/context"
)

const isDevEnvironment = true

// HTTPLogger borrowed heavily from https://github.com/philpearl/tt_goji_middleware/blob/master/base/loggingmiddleware.go
func HTTPLogger(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		printStart("http-begin", r)

		lw := WrapWriter(w)
		h.ServeHTTPC(ctx, lw, r)
		end := time.Now()

		if lw.Status() == 0 {
			lw.WriteHeader(http.StatusOK)
		}
		printEnd("http-end", lw, end.Sub(start))
	}
	return goji.HandlerFunc(handler)
}

func printStart(reqID string, r *http.Request) {
	var buf bytes.Buffer

	if reqID != "" {
		cW(&buf, bBlack, "[%s] ", reqID)
	}
	buf.WriteString("Started ")
	cW(&buf, bMagenta, "%s ", r.Method)
	cW(&buf, nBlue, "%q ", r.URL.String())
	buf.WriteString("from ")
	buf.WriteString(r.RemoteAddr)

	log.Print(buf.String())
}

func printEnd(reqID string, w WriterProxy, dt time.Duration) {
	var buf bytes.Buffer

	if reqID != "" {
		cW(&buf, bBlack, "[%s] ", reqID)
	}
	buf.WriteString("Returning ")
	status := w.Status()
	if status < 200 {
		cW(&buf, bBlue, "%03d", status)
	} else if status < 300 {
		cW(&buf, bGreen, "%03d", status)
	} else if status < 400 {
		cW(&buf, bCyan, "%03d", status)
	} else if status < 500 {
		cW(&buf, bYellow, "%03d", status)
	} else {
		cW(&buf, bRed, "%03d", status)
	}
	buf.WriteString(" in ")
	if dt < 500*time.Millisecond {
		cW(&buf, nGreen, "%s", dt)
	} else if dt < 5*time.Second {
		cW(&buf, nYellow, "%s", dt)
	} else {
		cW(&buf, nRed, "%s", dt)
	}

	log.Print(buf.String())
}
