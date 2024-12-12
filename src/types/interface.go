package types

import "net/http"

type Fortress interface {
	Middleware(next http.Handler) http.Handler

	Configure(conf any) Fortress
	Build() func(http.Handler) http.Handler
}
