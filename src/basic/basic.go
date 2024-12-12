package basic

import (
	"fmt"
	"net/http"
)

type BasicAuth struct {
}

// TODO
func (ba *BasicAuth) Configure(username, password string) *BasicAuth {
	return ba
}

// TODO
func (ba *BasicAuth) Build() func(http.Handler) http.Handler {
	return ba.Middleware
}

// TODO
func (ba *BasicAuth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HELLO FROM FORTRESS")
		next.ServeHTTP(w, r)
	})
}
