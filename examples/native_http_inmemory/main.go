package main

import (
	"io"
	"log"
	"net/http"

	"github.com/kruemelmann/fortress"
)

func main() {
	mux := http.NewServeMux()

	healthHandler := http.HandlerFunc(HealthCheckHandler)
	mux.Handle("/health", fortress.AuthMiddleware(healthHandler))

	log.Println("starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}
