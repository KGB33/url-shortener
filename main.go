package main

import (
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
	router.HandleFunc("/r/{shortUrl}", redirect)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	log.Printf("Homepage index hit by %v\n", r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["shortUrl"]

	var url Url
	for _, u := range URLs {
		if u.Short == shortUrl {
			url = u
			break
		}
	}
	fmt.Fprintf(w, "Redirected from %s to %s", url.Short, url.Dest)
}

type Url struct {
	Dest  string `json:"Dest"`
	Short string `json:"Short"`
}

var URLs []Url
