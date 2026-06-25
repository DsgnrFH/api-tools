package handlers

import (
	"encoding/json"
	"net/http"
	"dsgnrfh/pkg/id"
)

func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	uuid := id.GenerateUUIDv4()

	if r.URL.Query().Get("format") == "text" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(uuid))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"uuid": uuid})
}

func UUIDv7Handler(w http.ResponseWriter, r *http.Request) {
	uuidStr := id.GenerateUUIDv7()

	if r.URL.Query().Get("format") == "text" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(uuidStr))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"uuid": uuidStr,
		"type": "uuidv7",
	})
}
