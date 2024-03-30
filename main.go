package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main()  {
	fmt.Println("Hello World")

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Err occured %s\n", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD\n"))
}
