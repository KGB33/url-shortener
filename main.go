package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()
var rdb *redis.Client
var URLs []Url

// var ErrNil = errors.New("No Matching Records Found")

func main() {
	var err error
	rdb, err = NewDBClient("localhost:6379", "", 0)
	if err != nil {
		log.Fatal(err)
	}
	TEMP_insertUrls()
	handleRequests()
}

func TEMP_insertUrls() {
	URLs = []Url{{"https://google.com/", "goo"}, {"https://github.com/", "gh"}}
	for _, u := range URLs {
		if err := rdb.Set(ctx, u.Short, u.Dest, 0).Err(); err != nil {
			log.Fatal(err)
		}
	}
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

func NewDBClient(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil

}

type Url struct {
	Dest  string `json:"Dest"`
	Short string `json:"Short"`
}
