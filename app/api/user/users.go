package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create user")
	w.Write([]byte("Hello World!"))
}

func findUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Find users")
	w.Write([]byte("Hello World!"))
}

func findOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Find one user")
	w.Write([]byte("Hello World!"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user")
	w.Write([]byte("Hello World!"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user")
	w.Write([]byte("Hello World!"))
}

func validate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Validate token")
	w.Write([]byte("Hello World!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user")
	w.Write([]byte("Hello World!"))
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register user")
	w.Write([]byte("Hello World!"))
}

func Routes(route chi.Router) {
	route.Post("/", createUser)
	route.Get("/", findUsers)
	route.Get("/validate", validate)
	route.Post("/auth/login", login)
	route.Post("/auth/register", register)
	route.Get("/{id}", findOneUser)
	route.Put("/{id}", updateUser)
	route.Delete("/{id}", deleteUser)
}
