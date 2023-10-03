package main

import (
	"Server3rdPartyLibraries/db"
	"Server3rdPartyLibraries/internal/movie"
	"Server3rdPartyLibraries/internal/movie/handler"
	"Server3rdPartyLibraries/internal/movie/repository"
)

type Dependencies struct {
	movieHandler *handler.Handler
}

func NewDependencies() *Dependencies {
	movieDB, err := db.CreateDB()

	if err != nil {
		panic(err)
	}

	movieStorage := repository.NewRepo(movieDB)
	movieService := movie.NewService(movieStorage)

	movieHandler := handler.NewHandler(movieService)

	return &Dependencies{movieHandler: movieHandler}
}
