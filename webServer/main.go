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

func main() {

	fmt.Println("Starting a Web Server")

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/actors", ActorHandler)
	http.HandleFunc("/newReq", MyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // We dont need to specify another handler, since we attach it before
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
	// _, _ = fmt.Fprint(w, "hello world")

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
