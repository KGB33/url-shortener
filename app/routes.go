package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/", s.handleIndex())

	s.Router.HandleFunc("/c", s.handleCreateUrl()).Methods("POST")
	s.Router.HandleFunc("/r/{shortUrl}", s.handleRedirect()).Methods("GET")
	s.Router.HandleFunc("/u/{orgUrl}", s.handleUpdateUrl()).Methods("PUT")
	s.Router.HandleFunc("/d/{shortUrl}", s.handledeleteUrl()).Methods("DELETE")
}

// Main Page - a list of all shortened URLS
func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urls, err := scanUrls(s)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, urls)
	}
}

func (s *Server) handleCreateUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var newUrl Url
		err := json.Unmarshal(reqBody, &newUrl)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		if newUrl.Short == "" || newUrl.Dest == "" {
			respondWithError(w, http.StatusBadRequest, "Missing some Url fields")
			return
		}
		if err := postUrl(newUrl, s); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusCreated, newUrl)
	}
}

func (s *Server) handleRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["shortUrl"]
		url, err := getUrl(shortUrl, s)
		if err != nil {
			fmt.Fprintf(w, "No matching url for %s", shortUrl)
		} else {
			http.Redirect(w, r, url.Dest, 302)
		}
	}
}

func (s *Server) handleUpdateUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgUrl := mux.Vars(r)["orgUrl"]
		shortUrl := r.URL.Query().Get("shortUrl")
		destUrl := r.URL.Query().Get("destUrl")
		newUrl := Url{shortUrl, destUrl}
		if _, err := getUrl(orgUrl, s); err == nil {
			if err := deleteUrl(orgUrl, s); err != nil {
				log.Fatal(err)
			}
		}
		if err := postUrl(newUrl, s); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Server) handledeleteUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["shortUrl"]
		if err := deleteUrl(shortUrl, s); err != nil {
			log.Fatal(err)
		}
	}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Packs an error message into a json Object.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"Error": message})
}
