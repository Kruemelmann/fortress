package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kruemelmann/fortress"
	"github.com/kruemelmann/fortress/src/basic"
)

func main() {
	//setup fortress
	fortress := fortress.New(fortress.BasicAuth).Configure(basic.Config{Username: "username", Password: "password"}).Build()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	//make protected subRouter
	sr := r.PathPrefix("/protected").Subrouter()
	sr.HandleFunc("/", ProtectedHandler)
	sr.Use(fortress)

	log.Println("starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

// ProtectedHandler is inside the Subrouter (see main)
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `this route is protected with fortress`)
}
