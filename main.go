package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/greet", basicHandler)

	fmt.Println("Hello World")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Err occured %s\n", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD\n"))
}
