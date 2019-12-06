package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DB_USERNAME")
	fmt.Fprintln(w, "DB_USERNAME =", dbURL)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error test")
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
