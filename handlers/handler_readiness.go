package handlers

import ("net/http"
   "github.com/gupta799/go_api/response"
)

type EmptyStruct struct{} 

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
   response.RespondWithJson(w,200,new(EmptyStruct))
}