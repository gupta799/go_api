package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	log.Printf("Starting respondWithJson with code: %d and payload: %v", code, payload)
	
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v, error: %v", payload, err)
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	_, writeErr := w.Write(data)
	if writeErr != nil {
		log.Printf("Failed to write JSON response: %v", writeErr)
	}
	log.Println("Completed respondWithJson")
}
