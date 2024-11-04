package server

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// Logger is a middleware that logs the HTTP requests.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()

		// Create a response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w}

		// Process the request
		next.ServeHTTP(rw, r)

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Log the details
		log.Debug().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", rw.statusCode).
			Dur("latency", latency).
			Msg("request handled")
	})
}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
