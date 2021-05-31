package app

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/matryer/is"
)

// When `/` is requested, it should return the index.html file.
// ./static/index.html
// BUT because go test runs the code from a different CWD than main.go
// This test should always return an internal server error
func TestIndex(t *testing.T) {
	clearDB()
	is := is.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	is.Equal(http.StatusInternalServerError, response.Code)
	body := response.Body.String()
	exp := `{"Error":"Cannot load homepage: open static/index.html: no such file or directory"}`
	is.Equal(exp, body)
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
