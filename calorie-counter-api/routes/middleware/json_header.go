package middleware

import (
	"goji.io"
	"golang.org/x/net/context"
	"net/http"
)

// JSONHeader sets the content-type of responses to application/json.
func JSONHeader(h goji.Handler) goji.Handler {
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTPC(ctx, w, r)
	}
	return goji.HandlerFunc(handler)

}
