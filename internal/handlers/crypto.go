package handlers

import (
	"encoding/json"
	"net/http"
	"dsgnrfh/pkg/crypto"
)

func SHA256Handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing 'text' parameter"})
		return
	}

	hash := crypto.HashString(text)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"input":  text,
		"sha256": hash,
	})
}
