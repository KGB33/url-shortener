package app

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"
)

var s *Server
var urls []Url

func TestMain(m *testing.M) {
	s = NewServer("localhost:6379", "", 1)

	urls = []Url{
		{"burrito", "https://www.chipotle.com/"},
		{"gh", "https://github.com/"},
		{"goJson", "https://blog.golang.org/json"},
	}
	// Some tests assume the urls are sorted :(
	sort.Slice(urls, func(i, j int) bool { return urls[i].Short < urls[j].Short })

	clearDB()
	code := m.Run()
	os.Exit(code)
}

func clearDB() {
	if err := s.DB.FlushDB(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}

func popDB() {
	for _, u := range urls {
		u.Create(s)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr

}
