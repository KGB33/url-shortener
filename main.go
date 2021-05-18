package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	URLs = []Url{{"google.com", "goo"}, {"Github.com", "gh"}}
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/r", redirect)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	log.Printf("Homepage index hit by %v\n", r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Redirect Page")
	if err := json.NewEncoder(w).Encode(URLs); err != nil {
		log.Fatal(err)
	}

}

type Url struct {
	Destination string `json:"Destination"`
	Shortened   string `json:"Shortened"`
}

var URLs []Url
