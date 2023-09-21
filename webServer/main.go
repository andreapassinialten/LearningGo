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

	log.Fatal(http.ListenAndServe(":8080", nil)) // We dont need to specify another handler, since we attach it before
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "hello world")
	_, _ = io.WriteString(w, "Hello from Italy")
}
