package app

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// When `/` is requested, it should return a list of
// all the entries in the DB, if the DB is
// empty, then an empty list should be returned.
func TestEmptyDB_Index(t *testing.T) {
	clearDB()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	fmt.Println(response.Header())
	body := response.Body.String()
	body = strings.TrimSpace(body)
	if body != "null" {
		t.Errorf("Expected null, got: %s\n", body)
	}
}

func TestCreateUrl(t *testing.T) {
	jsonStr := []byte(`{"ShortUrl":"short", "DestUrl":"dest"}`)
	req, _ := http.NewRequest("POST", "/c", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	u, err := getUrl("short", &s)
	if err != nil {
		t.Errorf("Failed to retreave new URL from database: %s\n", err)
	}
	expectedUrl := Url{"short", "dest"}
	if u != expectedUrl {
		t.Errorf("Url inserted into DB does not match expected\n\tExpected: %v\n\tGot: %v", expectedUrl, u)
	}

}
