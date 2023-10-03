package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

const port string = ":8080"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var dep *Dependencies = NewDependencies()

	r.Post("/movies", dep.movieHandler.CreateMovieHandler)

	log.Printf("Web app is running in port %s", port)
	http.ListenAndServe(port, r)
}
