package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	URLs = []Url{{"https://google.com/", "goo"}, {"https://github.com/", "gh"}}
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
	url, err := getUrlFromShort(shortUrl)
	if err != nil {
		fmt.Fprintf(w, "No matching url for %s", shortUrl)
	} else {
		http.Redirect(w, r, url.Dest, 302)
	}
}

// Returns a given URL with matching `Short` key
// or an zero value'd URL and an error if the url is not found
func getUrlFromShort(s string) (Url, error) {
	for _, url := range URLs {
		if url.Short == s {
			return url, nil
		}
	}
	return Url{}, errors.New("No URL found")
}

type Url struct {
	Dest  string `json:"Dest"`
	Short string `json:"Short"`
}

var URLs []Url
