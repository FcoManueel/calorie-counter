package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSHeader lets browsers play nicely with the API
func CORSHeader(h http.Handler) http.Handler {
	// Added "" as an AllowedMethod in case the browser fails to send 'Access-Control-Request-Method'
	// Added "Accept","Authorization","Content-Type" to Allowed Headers since cors implements strict checking

	middleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "POST", "DELETE", "GET", ""},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "x-debug", "accept-language"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Total-Count"},
		MaxAge:           1800,
		Debug:            false,
	})
	return middleware.Handler(h)
}
