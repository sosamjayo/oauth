package handler

import (
	"net/http"
)

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Invalid Code", http.StatusBadRequest)
		return
	}

	// Get the state from the query
}
