package main

import (
	"fmt"
	"log"
	"net/http"

	"Calculator/calculator"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to calculator back end!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	http.Handle("/", router)

	router.HandleFunc("/", calculator.Handler).Name("Homepage").Methods("POST")
	router.HandleFunc("/", homepageHandler).Name("Homepage")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
