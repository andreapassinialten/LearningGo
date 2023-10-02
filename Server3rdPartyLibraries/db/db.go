package db

import (
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
	return &MoviesDB{DB: db}, nil
}
