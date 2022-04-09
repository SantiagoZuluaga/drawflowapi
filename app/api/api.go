package api

import (
	"github.com/SantiagoZuluaga/drawflowapi/app/api/program"
	"github.com/SantiagoZuluaga/drawflowapi/app/api/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/users", user.Routes)
	router.Route("/programs", program.Routes)

	return router
}
