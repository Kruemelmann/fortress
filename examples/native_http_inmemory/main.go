package main

import (
	"io"
	"log"
	"net/http"

	"github.com/kruemelmann/fortress"
	"github.com/kruemelmann/fortress/src/basic"
)

func main() {
	//setup fortress
	fortress := fortress.New(fortress.BasicAuth).Configure(basic.Config{Username: "username", Password: "password"}).Build()

	srv := http.NewServeMux()
	healthHandler := http.HandlerFunc(HealthCheckHandler)
	srv.Handle("/health", healthHandler)
	srv.Handle("/protected", fortress(http.HandlerFunc(ProtectedHandler)))

	log.Println("starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", srv))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `this route is protected by fortress`)
}
