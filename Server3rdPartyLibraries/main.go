package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"net/http"
	"time"
)

var movies map[string]Movie = make(map[string]Movie)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.With(SayHelloMiddleware).Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Movies GET: ")
		var body io.ReadCloser = r.Body
		var movie Movie
		_ = json.NewDecoder(body).Decode(&movie)

		val, ok := movies[movie.ImdbID]

		var movieResponse = val

		if ok {
			w.WriteHeader(http.StatusAccepted)
			_ = json.NewEncoder(w).Encode(movieResponse)
			fmt.Println(movieResponse)
		} else {
			w.WriteHeader(http.StatusNoContent)
			fmt.Println("No movie found")
			movieError := MovieError{
				Error:        errors.New("No movie found"),
				HappenedAt:   time.Now(),
				ErrorMessage: "No movie found",
			}
			fmt.Println(movieError)
			_ = json.NewEncoder(w).Encode(movieError)
		}
	})
	r.Post("/movies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Movies POST: ")
		var body io.ReadCloser = r.Body
		var movie Movie
		_ = json.NewDecoder(body).Decode(&movie)

		val, ok := movies[movie.ImdbID]
		if !ok {
			movies[movie.ImdbID] = Movie{
				ImdbID:   movie.ImdbID,
				Title:    movie.Title,
				Genre:    movie.Genre,
				Director: movie.Director,
				Cast:     movie.Cast,
			}
			w.WriteHeader(http.StatusAccepted)
			_, _ = io.WriteString(w, "Movie added")
			fmt.Println("Movie added")
			fmt.Println(val)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Cannot ADD movie")
			movieError := MovieError{
				Error:        errors.New("Cannot ADD movie"),
				HappenedAt:   time.Now(),
				ErrorMessage: "Cannot ADD movie",
			}
			fmt.Println(movieError)
			_ = json.NewEncoder(w).Encode(movieError)
		}
	})
	r.Delete("/movies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Movies DELETE: ")
		var body io.ReadCloser = r.Body
		var movie Movie
		_ = json.NewDecoder(body).Decode(&movie)

		delete(movies, movie.ImdbID)

		_, ok := movies[movie.ImdbID]

		if !ok {
			w.WriteHeader(http.StatusAccepted)
			_, _ = io.WriteString(w, "Movie deleted")
			fmt.Println("Movie deleted")
			fmt.Println(movie.ImdbID)
		} else {
			w.WriteHeader(http.StatusNoContent)
			fmt.Println("No movie to delete")
			movieError := MovieError{
				Error:        errors.New("No movie to delete"),
				HappenedAt:   time.Now(),
				ErrorMessage: "No movie to delete",
			}
			fmt.Println(movieError)
			_ = json.NewEncoder(w).Encode(movieError)
		}
	})
	http.ListenAndServe("localhost:8080", r)

}

func SayHelloMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello there")

		next.ServeHTTP(writer, request)
	})
}

type Movie struct {
	ImdbID   string  `json:"imdb_id"`
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Director string  `json:"director"`
	Cast     []Actor `json:"cast"`
}

type Actor struct {
	Name string `json:"name"`
}

type MovieError struct {
	HappenedAt   time.Time `json:"happened_at"`
	Error        error     `json:"error"`
	ErrorMessage string
}
