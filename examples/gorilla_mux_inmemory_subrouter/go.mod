module github.com/kruemelmann/fortress/examples/gorilla_mux_inmemory_subrouter

go 1.22.2

replace github.com/kruemelmann/fortress => ../../../fortress

require (
	github.com/gorilla/mux v1.8.1
	github.com/kruemelmann/fortress v0.0.0-00010101000000-000000000000
)