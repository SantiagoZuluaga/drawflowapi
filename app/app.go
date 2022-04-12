package app

import (
	"fmt"
	"net/http"

	"github.com/SantiagoZuluaga/drawflowapi/app/config"
	"github.com/SantiagoZuluaga/drawflowapi/app/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RunServer() {
	router := chi.NewRouter()
	router.Use(cors.Handler(config.CORS()))
	routes := routes.Routes()
	router.Mount("/api", routes)

	fmt.Println("Starting server on: http://localhost" + config.PORT)
	http.ListenAndServe(config.PORT, router)
}
