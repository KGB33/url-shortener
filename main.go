package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	var err error
	rdb, err = NewDBClient("localhost:6379", "", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.ShutdownSave(ctx)
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)

	router.HandleFunc("/c", createUrl).Methods("POST")
	router.HandleFunc("/r/{shortUrl}", redirect).Methods("GET")
	router.HandleFunc("/u/{orgUrl}", updateUrl).Methods("PUT")
	router.HandleFunc("/d/{shortUrl}", deleteUrlRoute).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Main Page - also a list of all shortened URLS
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	urls, err := scanUrls()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Homepage index hit by %v\n", r)
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUrl Url
	err := json.Unmarshal(reqBody, &newUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := postUrl(newUrl); err != nil {
		log.Fatal(err)
	}

}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["shortUrl"]
	url, err := getUrl(shortUrl)
	if err != nil {
		fmt.Fprintf(w, "No matching url for %s", shortUrl)
	} else {
		http.Redirect(w, r, url.Dest, 302)
	}
}

func updateUrl(w http.ResponseWriter, r *http.Request) {
	orgUrl := mux.Vars(r)["orgUrl"]
	shortUrl := r.URL.Query().Get("shortUrl")
	destUrl := r.URL.Query().Get("destUrl")
	newUrl := Url{shortUrl, destUrl}
	if err := putUrl(orgUrl, newUrl); err != nil {
		log.Fatal(err)
	}
}

func deleteUrlRoute(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["shortUrl"]
	if err := deleteUrl(shortUrl); err != nil {
		log.Fatal(err)
	}
}

// Returns a given URL with matching `Short` key
// or an zero value'd URL and an error if the url is not found
func getUrl(s string) (Url, error) {
	destUrl, err := rdb.Get(ctx, s).Result()
	if err != nil {
		return Url{}, err
	}
	return Url{s, destUrl}, nil
}

func putUrl(org string, u Url) error {
	if org != u.Short {
		if err := deleteUrl(org); err != nil {
			return err
		}
	}
	return postUrl(u)
}

func postUrl(u Url) error {
	return rdb.Set(ctx, u.Short, u.Dest, 0).Err()
}

func deleteUrl(s string) error {
	return rdb.Del(ctx, s).Err()
}

func scanUrls() ([]Url, error) {
	var urls []Url
	iter := rdb.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		nextUrl, err := getUrl(iter.Val())
		if err != nil {
			return nil, err
		}
		urls = append(urls, nextUrl)
	}
	return urls, nil

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
