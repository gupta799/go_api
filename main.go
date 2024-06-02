package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/parnurzeal/gorequest"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"PUT", "POST", "GET", "DELET"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	request := gorequest.New()
	myStruct := httpClientStruct{
		Request: request,
	}
	port := ":8080"

	mountedRouter := chi.NewRouter()
	mountedRouter.Get("/health", handlerReadiness)

	r.Mount("/v1", mountedRouter)
    r.Post("/about",myStruct.request_handler)

	srv := &http.Server{
		Handler: r,
		Addr:    port,
	}
	fmt.Printf("Listening on port %s\n", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Something went wrong with the app...")
	}

}
