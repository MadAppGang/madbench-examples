package main

import (
	"net/http"

	"example.com/codenav/auth"
	"example.com/codenav/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)
	// Direct call so a literal `HandleLogin(` reference lives in this file too.
	mux.HandleFunc("/login-direct", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleLogin(w, r)
	})
	_ = http.ListenAndServe(":8080", mux)
}
