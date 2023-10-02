package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var movies map[string]Movie = make(map[string]Movie)

func main() {

	// Add using unique key and a value
	movies["1"] = Movie{
		ImdbID:   "1",
		Title:    "Titanic",
		Genre:    "Drama",
		Director: "James Cameron",
	}

	movies["2"] = Movie{
		ImdbID:   "2",
		Title:    "Terminator",
		Genre:    "Sci-fi",
		Director: "James Cameron",
	}

	fmt.Println("Starting a Web Server")

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/actors", ActorHandler)
	http.HandleFunc("/newReq", MyHandler)
	http.HandleFunc("/movies", MovieHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // We dont need to specify another handler, since we attach it before

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

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	var method string = r.Method

	switch method {
	case http.MethodGet:
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

	case http.MethodPost:
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

	case http.MethodDelete:
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

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// Tell the client that there is an error
		movieError := MovieError{
			Error:      errors.New("Something went wrong"),
			HappenedAt: time.Now(),
		}
		fmt.Println(movieError)
		_ = json.NewEncoder(w).Encode(movieError)
	}
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	var method string = r.Method

	switch method {
	case http.MethodGet:
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, "Calling a Get \n")
		var values = r.URL.Query()
		var name string = values.Get("name")
		_, _ = io.WriteString(w, name+"\n")
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		_, _ = io.WriteString(w, "Calling a POST \n")
		var values = r.URL.Query()
		var surname string = values.Get("surname")
		_, _ = io.WriteString(w, surname+"\n")
	case http.MethodPut:
		w.WriteHeader(http.StatusAccepted)
		_, _ = io.WriteString(w, "Calling a PUT \n")
	}
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case http.MethodGet:
		_, _ = io.WriteString(w, "Hello from Italy")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "DEFAULT")
	}

}

func ActorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	method := r.Method
	switch method {
	case http.MethodPost:
		var body io.ReadCloser = r.Body
		var actorRequest ActorRequest
		_ = json.NewDecoder(body).Decode(&actorRequest)

		// I have all the info in actorRequest
		actorResponse := actorRequest.toResponse()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(actorResponse)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// Tell the client that there is an error
		actorError := ActorError{
			Error:      errors.New("Something went wrong"),
			HappenedAt: time.Now(),
		}
		_ = json.NewEncoder(w).Encode(actorError)
	}
}

type ActorError struct {
	HappenedAt time.Time `json:"happened_at"`
	Error      error     `json:"error"`
}

type ActorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func (a ActorRequest) toResponse() ActorResponse {
	return ActorResponse{
		Name:     fmt.Sprintf("Name : %s, Last name: %s, Age: %d", a.FirstName, a.LastName, a.Age),
		CreateAt: time.Now(),
	}
}

type ActorResponse struct {
	Name     string    `json:"full_name"`
	CreateAt time.Time `json:"create_at"`
}
