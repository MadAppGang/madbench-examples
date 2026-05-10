package auth

import "net/http"

// HandleLogin authenticates a user and issues a session token.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// intentionally thin — this file is a navigation target, not real code
	w.WriteHeader(http.StatusOK)
}
