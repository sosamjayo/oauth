package handler

import (
	"net/http"
)

func (h *Handler) AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: クエリパラメータからクライアントIDなどを取得する
}
