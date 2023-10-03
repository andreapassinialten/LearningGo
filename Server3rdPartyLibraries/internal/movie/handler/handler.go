package handler

import (
	"Server3rdPartyLibraries/internal/movie"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	*movie.Service
}

func NewHandler(s *movie.Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close() // defer will put the fun call just before the return

	var _movie movie.Movie
	err := json.NewDecoder(body).Decode(&_movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Msgf("cannot decode request, %s", err)
		_ = json.NewEncoder(w).Encode(struct {
			Err string `json:"err"`
		}{Err: "Cannot decode request"})
		return
	}

	movieCreated, err := h.Create(_movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Msgf("cannot save movie, %s", err)
		_ = json.NewEncoder(w).Encode(struct {
			Err string `json:"err"`
		}{Err: "Cannot save request"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(movieCreated)
}
