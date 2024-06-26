package skeleton

import (
	"net/http"
	"os"
)

func restrictMethods(h http.Handler, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method == method {
				h.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// GetOrDefault gets environment variable 'key'.
// If not found, returns the supplied 'fallback'.
func GetOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
