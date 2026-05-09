package routes

import (
	"net/http"

	"example.com/codenav/auth"
)

// Register wires the HTTP routes onto the given mux.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleLogin(w, r)
	})
}
