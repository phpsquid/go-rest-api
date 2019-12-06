package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go + Docker")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Error test")
	log.Fatal("intentional error")
}

func main() {
	fmt.Println("Go Docker tutorial")

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/error", ErrorHandler)
	
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
