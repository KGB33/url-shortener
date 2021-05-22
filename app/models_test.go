package app

import "testing"

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

	urlBlank.Get("short", s)
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
	t.Errorf("not Implemented")
}

func TestUrl_Update_InvalidKey(t *testing.T) {
	t.Errorf("not Implemented")
}

func TestUrl_Delete(t *testing.T) {
	t.Errorf("not Implemented")
}

func TestUrl_Delete_InvalidKey(t *testing.T) {
	t.Errorf("not Implemented")
}

func TestScanUrl_PopulatedDB(t *testing.T) {
	t.Errorf("Not Implemented")
}

func TestScanUrl_EmptyDB(t *testing.T) {
	t.Errorf("Not Implemented")
}
