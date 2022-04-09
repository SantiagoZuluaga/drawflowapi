package config

import (
	"github.com/go-chi/cors"
)

func CORS() cors.Options {
	handler := cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
			"https://drawflowapp.herokuapp.com",
			"http://drawflowapp.herokuapp.com",
			"https://drawflowapi.herokuapp.com",
			"http://drawflowapi.herokuapp.com",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Access-Control-Allow-Origin",
			"Accept",
			"Authorization",
			"Content-Type",
		},
	}
	return handler
}
