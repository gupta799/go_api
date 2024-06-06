package baseClient

import (
	"encoding/json"
	"log"

	"github.com/parnurzeal/gorequest"
)

type BasePayload[T any] struct {
	Data T `json:"data"`
}

type BaseResponse[T1 any] struct {
	Data struct {
		Resource T1 `json:"resource"`
	} `json:"data"`
	Code int `json:"code"`
}

type BaseService  interface {
	Post(url string, payload interface{}, result interface{}) interface{}
}

type BaseClientStruct struct {
	Request *gorequest.SuperAgent
}

func (baseClient *BaseClientStruct) Post(url string, payload interface{}, result interface{}) interface{} {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, body, errs := baseClient.Request.Post(url).
		Send(string(jsonData)).
		End()

	if errs != nil {
		log.Printf("Errors: %v", errs)
		return errs[0]
	}

	log.Println("Response Body:", body)

	// Unmarshal the response body into the result interface{}
	if err := json.Unmarshal([]byte(body), result); err != nil {
		return err
	}

	return nil
}
