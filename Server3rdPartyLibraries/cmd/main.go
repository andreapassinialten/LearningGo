package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

const port string = ":8080"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := chi.NewRouter()

	log.Printf("Web app is running in port %s", port)
	http.ListenAndServe(port, r)
}
