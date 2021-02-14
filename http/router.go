package http

import (
	"github.com/gorilla/mux"

	"github.com/sumelms/microkit/http/handler"
	"github.com/sumelms/microkit/http/middleware"
)

// NewRouter creates a new mux.Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Default middlewares
	r.Use(middleware.JSONEncodeMiddleware)
	// Default handlers
	r.Handle("/health", handler.HealthCheckerHandler())
	return r
}
