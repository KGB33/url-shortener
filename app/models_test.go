package app

import (
	"testing"

	"github.com/matryer/is"
)

func TestUrl_Get(t *testing.T) {
	clearDB()
	popDB()

	url := Url{}
	expected := Url{"gh", "https://github.com/"}
	err := url.Get("gh", s)
	if err != nil {
		t.Errorf("Error when getting URL: %s", err)
	}
	if url != expected {
		t.Errorf("Returned URL does not match. Got: %v, Expected: %v", url, expected)
	}
}

func TestUrl_Get_InvalidKey(t *testing.T) {
	clearDB()

	url := Url{}
	expected := Url{}
	err := url.Get("gh", s)
	if err.Error() != "redis: nil" {
		t.Errorf("Error when getting URL: %s", err)
	}
	if url != expected {
		t.Errorf("Returned URL does not match. Got: %v, Expected: %v", url, expected)
	}
}

func TestUrl_Create(t *testing.T) {
	clearDB()

	url := Url{"short", "dest"}
	urlBlank := Url{}

	err := url.Create(s)
	if err != nil {
		t.Errorf("Error when creating URL: %s", err)
	}

	if err := urlBlank.Get("short", s); err != nil {
		t.Errorf("Error when retreaving URL from db: %s", err)
	}
	if url != urlBlank {
		t.Errorf("Url Insterted or retreaved incorrectly.")
	}
}

func TestUrl_Create_DuplicateKey(t *testing.T) {
	clearDB()
	popDB()

	url := Url{"gh", "https://New.Github.com/"}

	err := url.Create(s)
	if err == nil {
		t.Errorf("No Error when creating a duplicate key.")
	}
}

func TestUrl_Update(t *testing.T) {
	clearDB()
	popDB()

	url := Url{"gh", "https://New.Github.com/"}

	if err := url.Update(s); err != nil {
		t.Errorf("Failed to update URL - %s", err)
	}
}

func TestUrl_Update_InvalidKey(t *testing.T) {
	clearDB()
	popDB()

	url := Url{"Not a Key", "this is all fake"}
	err := url.Update(s)
	if err == nil {
		t.Error("Expected an error when updating a non-existant URL")
	}
}

func TestUrl_Delete(t *testing.T) {
	clearDB()
	popDB()

	var u Url
	if err := u.Get("gh", s); err != nil {
		t.Errorf("Unable to get URL...")
	}

	if err := u.Delete(s); err != nil {
		t.Errorf("Error when deleting URL: %s", err)
	}

	if err := u.Get("gh", s); err == nil {
		t.Error("Url was not deleted")
	}

}

func TestUrl_Delete_InvalidKey(t *testing.T) {
	clearDB()
	popDB()

	u := Url{"Not A Key", "Fake Val too"}

	if err := u.Delete(s); err == nil {
		t.Errorf("Expected an error when deleting a non-existant key")
	}
}

func TestScanUrl_PopulatedDB(t *testing.T) {
	clearDB()
	popDB()
	is := is.New(t)

	u, err := scanUrls(s)
	is.NoErr(err)
	is.Equal(u, urls)
}

func TestScanUrl_EmptyDB(t *testing.T) {
	clearDB()

	expected := []Url{}
	is := is.New(t)
	u, err := scanUrls(s)
	is.NoErr(err)
	is.Equal(u, expected)
}
