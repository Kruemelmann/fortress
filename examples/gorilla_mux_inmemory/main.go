package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/kruemelmann/fortress"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r.Use(fortress.AuthMiddleware)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
