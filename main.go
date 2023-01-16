package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/razak17/go-poke-api/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/pokemon/{name}", routes.GetPokemon)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
