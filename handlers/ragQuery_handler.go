package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gupta799/go_api/baseClient"
	"github.com/parnurzeal/gorequest"
	  "github.com/gupta799/go_api/response"
)


type HttpClientStruct struct {
	Request     *gorequest.SuperAgent
	BaseRequest baseClient.BaseService
}


type Payload struct {
	Question string `json:"question"`
}

type Response struct {
	Contextvars string `json:"contextvars"`
}

func (client *HttpClientStruct) asyncRequest(resultChan chan<- *baseClient.BaseResponse[Response], payload Payload) {
	var ragRespose baseClient.BaseResponse[Response]
	errs := client.BaseRequest.Post("http://localhost:5000/about", payload, &ragRespose)

	if errs != nil {
		log.Printf("Errors: %v", errs)
		resultChan <- nil
	} else {

		resultChan <- &ragRespose

	}
}

func (client *HttpClientStruct) RagQauery_handler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan *baseClient.BaseResponse[Response])
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go client.asyncRequest(resultChan, payload)

	// Start a goroutine to collect the result and write the response

	apiResponse := <-resultChan
	if apiResponse != nil {
		response.RespondWithJson(w, 200, apiResponse)
	} else {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}
	close(resultChan) // Close the channel after sending the result

}
