package fortress

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HELLO FROM FORTRESS")
		next.ServeHTTP(w, r)
	})
}
