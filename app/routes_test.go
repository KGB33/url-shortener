package app

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

// When `/` is requested, it should return a list of
// all the entries in the DB, if the DB is
// empty, then an empty list should be returned.
func TestIndex_EmptyDB(t *testing.T) {
	clearDB()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()
	body = strings.TrimSpace(body)
	if body != "null" {
		t.Errorf("Expected null, got: %s\n", body)
	}
}

func TestIndex_PopulatedDB(t *testing.T) {
	clearDB()
	popDB()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()
	expected := []string{
		`{"ShortUrl":"goJson","DestUrl":"https://blog.golang.org/json"}`,
		`{"ShortUrl":"burrito","DestUrl":"https://www.chipotle.com/"}`,
		`{"ShortUrl":"gh","DestUrl":"https://github.com/"}`,
	}
	body = strings.TrimSpace(body)
	for _, s := range expected {
		if !strings.Contains(body, s) {
			t.Errorf("Expected %s in response, got: %s\n", s, body)
		}
	}
}

func TestEndpoint_CreateUrl(t *testing.T) {
	clearDB()
	jsonStr := []byte(`{"ShortUrl":"short", "DestUrl":"dest"}`)
	req, _ := http.NewRequest("POST", "/c", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	u := Url{}
	err := u.Get("short", s)
	if err != nil {
		t.Errorf("Failed to retreave new URL from database: %s\n", err)
	}
	expectedUrl := Url{"short", "dest"}
	if u != expectedUrl {
		t.Errorf("Url inserted into DB does not match expected\n\tExpected: %v\n\tGot: %v", expectedUrl, u)
	}

}

// When malformed Json data is posted to the
// `/c/` endpoint, ensure that the server responds
// with a 400 bad request status code.
func TestEndpoint_CreateUrl_MalformedRequest(t *testing.T) {
	clearDB()
	badJson := bytes.NewBuffer([]byte(`{"Im Bad": "Json", "This isn't right": "at all"}`))
	req, _ := http.NewRequest("POST", "/c", badJson)
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	expected := `{"Error":"Missing some Url fields"}`
	if body := response.Body.String(); body != expected {
		t.Errorf("Server response does not match expected.\n\tGot: %s\n\tExpected: %s", body, expected)
	}
}

func TestEndpoint_CreateUrl_DuplicateShort(t *testing.T) {
	clearDB()
	popDB()

	js := bytes.NewBuffer([]byte(`{"ShortUrl":"goJson", "DestUrl":"https://blog.golang.org/json"}`))
	req, _ := http.NewRequest("POST", "/c", js)
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	expected := `{"Error":"Unable to insert the URL into the database. This is likely due to a duplicate ShortUrl."}`
	if body := response.Body.String(); body != expected {
		t.Errorf("Response does not match expected.\n\tGot: %s\n\tExpected: %s", body, expected)
	}
}
