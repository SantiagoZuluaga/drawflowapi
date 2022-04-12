package app

import (
	"fmt"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	fmt.Println("Starting server on: http://localhost" + port)
	http.ListenAndServe(port, router)
}
