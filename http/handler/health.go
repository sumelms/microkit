package handler

import "net/http"

// HealthCheckerHandler service health
func HealthCheckerHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK"))
		rw.WriteHeader(200)
		return
	})
}
