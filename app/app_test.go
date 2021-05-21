package app

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var s Server

func TestMain(m *testing.M) {
	s = NewServer("localhost:6379", "", 1)
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
	urls := []Url{
		{"goJson", "https://blog.golang.org/json"},
		{"gh", "https://github.com/"},
		{"burrito", "https://www.chipotle.com/"},
	}
	for _, u := range urls {
		postUrl(u, &s)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected: %d -- Got %d\n", expected, actual)
	}
}
