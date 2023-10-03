package repository

import (
	"Server3rdPartyLibraries/db"
	"Server3rdPartyLibraries/internal/movie"
	"fmt"
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
	fmt.Println(m)
	log.Printf("Movie is saved")
	return movie.Movie{}, nil
}
