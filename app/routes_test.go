package app

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/matryer/is"
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
	if body != "[]" {
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
	is := is.New(t)
	jsonStr := []byte(`{"ShortUrl":"short", "DestUrl":"dest"}`)
	req, _ := http.NewRequest("POST", apiVer+"/c", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	is.Equal(http.StatusCreated, response.Code)

	u := Url{}
	err := u.Get("short", s)
	is.NoErr(err)
	expectedUrl := Url{"short", "dest"}
	is.Equal(expectedUrl, u) // Url inserted into DB does not match.
}

// When malformed Json data is posted to the
// `/c/` endpoint, ensure that the server responds
// with a 400 bad request status code.
func TestEndpoint_CreateUrl_MalformedRequest(t *testing.T) {
	clearDB()
	is := is.New(t)
	badJson := bytes.NewBuffer([]byte(`{"Im Bad": "Json", "This isn't right": "at all"}`))
	req, _ := http.NewRequest("POST", apiVer+"/c", badJson)
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	is.Equal(http.StatusBadRequest, response.Code)

	expected := `{"Error":"Missing the Url destination field"}`
	body := response.Body.String()
	is.Equal(expected, body)
}

func TestEndpoint_CreateUrl_MissingShortKey(t *testing.T) {
	clearDB()

	is := is.New(t)

	data := bytes.NewBuffer([]byte(`{"DestUrl":"https://https://redis.io/commands/EXISTS"}`))
	req, _ := http.NewRequest("POST", apiVer+"/c", data)
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	is.Equal(http.StatusCreated, response.Code)
	// The shortURL is random, but it seems like when running tests it always uses the same seed
	expected := `{"ShortUrl":"ejLqV3Wkyd6","DestUrl":"https://https://redis.io/commands/EXISTS"}`
	is.Equal(expected, response.Body.String())
}

func TestEndpoint_CreateUrl_DuplicateShort(t *testing.T) {
	clearDB()
	popDB()
	is := is.New(t)

	js := bytes.NewBuffer([]byte(`{"ShortUrl":"goJson", "DestUrl":"https://blog.golang.org/json"}`))
	req, _ := http.NewRequest("POST", apiVer+"/c", js)
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	is.Equal(http.StatusInternalServerError, response.Code)

	expected := `{"Error":"Unable to insert the URL into the database. This is likely due to a duplicate ShortUrl."}`
	body := response.Body.String()
	is.Equal(expected, body)
}
