package db

import (
	"Server3rdPartyLibraries/internal/movie"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MoviesDB struct {
	DB *gorm.DB
}

func CreateDB() (*MoviesDB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// panic("failed to connect database") // Cut the execution of the program
		return nil, err
	}

	Migrate(db)
	return &MoviesDB{DB: db}, nil
}

// migrations
func Migrate(db *gorm.DB) {
	log.Error().Msgf("%s", db.AutoMigrate(movie.Movie{}))
	log.Error().Msgf("%s", db.AutoMigrate(&movie.Actor{}))
}
