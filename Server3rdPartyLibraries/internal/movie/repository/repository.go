package repository

import (
	"Server3rdPartyLibraries/db"
	"Server3rdPartyLibraries/internal/movie"
)

type Repository struct {
	db *db.MoviesDB
}

func (r Repository) CreateMovie(m movie.Movie) (movie.Movie, error) {
	return movie.Movie{}, nil
}
