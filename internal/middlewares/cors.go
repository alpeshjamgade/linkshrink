package middlewares

import (
	"net/http"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Header().Set("Strict-Transport-Security", "max-age=3600;includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src'self'")

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1;mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}
