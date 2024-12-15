package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	testHeader := http.Header{}

	_, err := GetAPIKey(testHeader)

	if err == nil {
		t.Error("Error should not be nil, since the headers don't include the Authorization")
	}

	testHeader.Set("Authorization", "ApiKey99999")

	_, err = GetAPIKey(testHeader)

	if err == nil {
		t.Error("Error should not be nil, since the header is malformed")
	}

	_, err = GetAPIKey(testHeader)

	testHeader.Set("Authorization", "ApiKey99999 999")

	if err == nil {
		t.Error("Error should not be nil, since the header is malformed")
	}

	testHeader.Set("Authorization", "ApiKey 999")

	key, err := GetAPIKey(testHeader)

	if err != nil {
		t.Errorf("Error: %v in valid header, %v", err, testHeader)
	}

	if key != "999" {
		t.Errorf("Invalid key, key should be 999 but is %v", key)
	}
}
