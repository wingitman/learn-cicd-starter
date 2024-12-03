package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func TestAPIKey1(t *testing.T) {
  r := strings.NewReader("")
  re,_ := http.NewRequest(http.MethodGet, "https://google.co.uk", r) 
  re.Header.Add("Authorization","oo_spooky")

  _,err1 := GetAPIKey(re.Header) 
  if (err1 != nil) {
    t.Fatalf("Failed to get API Key due to error: %v",err1)
  }
}

func TestAPIKey2(t *testing.T) {
  t.Fatal("Rawr") 
}
