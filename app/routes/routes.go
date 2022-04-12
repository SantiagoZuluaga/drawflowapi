package routes

import (
	"github.com/SantiagoZuluaga/drawflowapi/app/routes/auth"
	"github.com/SantiagoZuluaga/drawflowapi/app/routes/programs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/auth", auth.Routes)
	router.Route("/programs", programs.Routes)

	return router
}
