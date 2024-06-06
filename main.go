package main

import (
	"fmt"
	"net/http"
	"github.com/gupta799/go_api/baseClient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/parnurzeal/gorequest"
	"github.com/gupta799/go_api/handlers"
	"github.com/gupta799/go_api/middlewares"
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
    baseStruct:= &baseClient.BaseClientStruct{
        Request: request,
    }
	
	myStruct := handlers.HttpClientStruct{
		Request: request,
        BaseRequest: baseStruct,
	}
	port := ":8080"

	mountedRouter := chi.NewRouter()
	mountedRouter.Get("/health", handlers.HandlerReadiness)

	r.Mount("/v1", mountedRouter)
	r.Get("/getToken",handlers.AuthHandler)
	   r.Group(func(r chi.Router) {
        r.Use(middlewares.JWTAuth) // Apply the JWTAuth middleware to this group
		r.Post("/about",myStruct.RagQauery_handler)

    })

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
