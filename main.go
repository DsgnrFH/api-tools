package main

import (
	"fmt"
	"log"
	"net/http"
	
	"dsgnrfh/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/id/uuid", handlers.UUIDHandler)
	mux.HandleFunc("/api/v1/id/uuidv7", handlers.UUIDv7Handler)

	mux.HandleFunc("/api/v1/network/subnet", handlers.SubnetHandler)

	mux.HandleFunc("/api/v1/crypto/sha256", handlers.SHA256Handler)

	fmt.Println("API Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
