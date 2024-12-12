package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kruemelmann/fortress"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r.Use(fortress.AuthMiddleware)

	log.Println("starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}
