package main

import "net/http"

type EmptyStruct struct{} 

func handlerReadiness(w http.ResponseWriter, r *http.Request){
    respondWithJson(w,200,new(EmptyStruct) )
}