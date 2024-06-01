package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to respond with Json response %v", payload)
		w.WriteHeader(500)
	}
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
