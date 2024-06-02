package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type httpClientStruct struct {
	Request *gorequest.SuperAgent
}

type Payload struct {
	Question string `json:"question"`
}

type Response struct {
	Contextvars string `json:"contextvars"`
}

func (client *httpClientStruct) asyncRequest(resultChan chan<- *Response, payload Payload) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	_, body, errs := client.Request.Post("http://localhost:5000/about").
		Send(string(jsonData)).
		End()
	if errs != nil {
		log.Printf("Errors: %v", errs)
		resultChan <- nil
	} else {
		var response Response
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error unmarshalling response: %v", err)
			resultChan <- nil
			return
		} else {
			resultChan <- &response
		}
	}
}

func (client *httpClientStruct) request_handler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan *Response)
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go client.asyncRequest(resultChan, payload)

	// Start a goroutine to collect the result and write the response
	go func() {
		response := <-resultChan
		if response != nil {
			respondWithJson(w, 200, response)
		} else {
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}
		close(resultChan) // Close the channel after sending the result
	}()
}
