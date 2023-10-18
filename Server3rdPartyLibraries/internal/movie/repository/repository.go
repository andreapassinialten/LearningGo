package repository

import (
	"Server3rdPartyLibraries/db"
	"Server3rdPartyLibraries/internal/movie"
	"github.com/rs/zerolog/log"
)

type RepositorySQLITE struct {
	db *db.MoviesDB
}

type RepositoryMONGODB struct {
	db *db.MoviesDB
}

func NewRepo(db *db.MoviesDB) RepositorySQLITE {
	return RepositorySQLITE{
		db: db,
	}
}

func (r RepositorySQLITE) CreateMovie(m movie.Movie) (movie.Movie, error) {
	err := r.db.DB.Create(&m).Error

	if err != nil {
		return movie.Movie{}, err
	}

	log.Info().Msgf("movie %s saved OK to the DB", m.Title)
	return m, nil
}
