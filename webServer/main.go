package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	method := r.Method

	switch method {
	case http.MethodGet:
		w.WriteHeader(http.StatusCreated)

		values := r.URL.Query()

		name := values.Get("name")

		_, _ = io.WriteString(w, name+"\n")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "DEFAULT")
	}

}
