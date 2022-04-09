package program

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createProgram(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func findPrograms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func findOneProgram(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func updateProgram(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func deleteProgram(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func Routes(route chi.Router) {
	route.Post("/", createProgram)
	route.Get("/", findPrograms)
	route.Get("/{id}", findOneProgram)
	route.Put("/{id}", updateProgram)
	route.Delete("/{id}", deleteProgram)
}
