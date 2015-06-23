package middleware

import (
	"net/http"
)

func SetApiVersion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-API-Version", "1")
		next.ServeHTTP(w, r)
	})
}
