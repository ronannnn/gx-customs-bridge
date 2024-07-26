package apis

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (hs *HttpServer) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(hs.infraMiddleware.Lang)
	r.Use(middleware.Recoverer)
	apiV1Router := chi.NewRouter()

	// biz routes
	apiV1Router.Group(func(r chi.Router) {
		r.Use(hs.infraMiddleware.AuthHandlers()...)
		r.Use(hs.infraMiddleware.ReqRecorder)
	})

	// mount /api/v1 to root router
	r.Mount("/api/v1", apiV1Router)
	return r
}
