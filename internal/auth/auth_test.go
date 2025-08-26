package auth_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestSuccess(t *testing.T) {
	keyOriginal := "12345"
	headers := http.Header{}
	headers["Authorization"] = []string{"ApiKey " + keyOriginal}

	key, err := auth.GetAPIKey(headers)
	if err != nil || key != keyOriginal {
		t.FailNow()
	}
}

func TestNoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := auth.GetAPIKey(headers)
	if !errors.Is(err, auth.ErrNoAuthHeaderIncluded) {
		t.FailNow()
	}
}

func TestWrongFormat(t *testing.T) {
	keyOriginal := "12345"
	headers := http.Header{}
	headers["Authorization"] = []string{"ApiKey" + keyOriginal}

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.FailNow()
	}
}

func TestWrongName(t *testing.T) {
	keyOriginal := "12345"
	headers := http.Header{}
	headers["Authorization"] = []string{"ApiKew" + keyOriginal}

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.FailNow()
	}
}
