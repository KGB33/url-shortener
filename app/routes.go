package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const apiVer = "/api/v1"

func (s *Server) initRoutes() {

	s.Router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	)
	s.Router.HandleFunc("/", s.handleIndex())
	s.Router.HandleFunc("/r/{shortUrl}", s.handleRedirect()).Methods("GET")

	s.Router.HandleFunc(apiVer+"/c", s.handleCreateUrl()).Methods("POST")
	s.Router.HandleFunc(apiVer+"/r", s.handleReadUrls()).Methods("GET")
	s.Router.HandleFunc(apiVer+"/u/{orgUrl}", s.handleUpdateUrl()).Methods("PUT")
	s.Router.HandleFunc(apiVer+"/d/{shortUrl}", s.handledeleteUrl()).Methods("DELETE")
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("static/index.html")
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Cannot load homepage: "+err.Error())
		}
		defer f.Close()
		io.Copy(w, f)
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
		if newUrl.Dest == "" {
			respondWithError(w, http.StatusBadRequest, "Missing the Url destination field")
			return
		}
		if newUrl.Short == "" {
			newUrl.generateShort(s)
		}
		err = newUrl.Create(s)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusCreated, newUrl)
	}
}

func (s *Server) handleReadUrls() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urls, err := scanUrls(s)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJson(w, http.StatusOK, urls)
	}
}

func (s *Server) handleRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["shortUrl"]
		var url Url
		err := url.Get(shortUrl, s)
		if err != nil {
			fmt.Fprintf(w, "No matching url for %s", shortUrl)
		} else {
			http.Redirect(w, r, url.Dest, 302)
		}
	}
}

func (s *Server) handleUpdateUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgUrlShort := mux.Vars(r)["orgUrl"]
		reqBody, _ := ioutil.ReadAll(r.Body)
		var newUrl, orgUrl Url
		orgUrl.Get(orgUrlShort, s)
		err := json.Unmarshal(reqBody, &newUrl)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		if newUrl.Dest == "" {
			respondWithError(w, http.StatusBadRequest, "Missing the Url destination field")
			return
		}
		if newUrl.Short == "" {
			newUrl.generateShort(s)
		}
		if orgUrl.Short == newUrl.Short {
			if err := newUrl.Update(s); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
		} else {
			if err := orgUrl.Delete(s); err != nil {
				orgUrl.Create(s)
				respondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				if err := newUrl.Create(s); err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
				}
			}

		}
		respondWithJson(w, http.StatusOK, newUrl)
	}
}

func (s *Server) handledeleteUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["shortUrl"]
		url := Url{shortUrl, ""}
		if err := url.Delete(s); err != nil {
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
