package middleware

import (
	"log"
	"net/http"
	"time"

	"goji.io"
	"golang.org/x/net/context"
)

const isDevEnvironment = true

// HTTPLogger borrowed heavily from https://github.com/philpearl/tt_goji_middleware/blob/master/base/loggingmiddleware.go
func HTTPLogger(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		httpBegin := "[http-begin]"
		httpEnd := "[http-end]"
		if isDevEnvironment {
			httpBegin = "\x1b[34;1m[http-begin]\x1b[0m"
			httpEnd = "\x1b[34;1m[http-end]\x1b[0m"
		}
		log.Println(httpBegin,
			"method:", r.Method,
			"uri:", r.RequestURI,
		)

		lw := WrapWriter(w)
		h.ServeHTTPC(ctx, lw, r)
		end := time.Now()

		if lw.Status() == 0 {
			lw.WriteHeader(http.StatusOK)
		}
		log.Println(httpEnd,
			"status: ", lw.Status(),
			"elapsed: ", end.Sub(start),
		)
	}
	return goji.HandlerFunc(handler)
}
