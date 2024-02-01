// http_client_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up environment for all tests
	os.Setenv("OPENAI_API_KEY", "dummy-key")
	defer os.Unsetenv("OPENAI_API_KEY")

	// Run the tests
	code := m.Run()

	// Any teardown or cleanup goes here

	os.Exit(code)
}
func TestSendImageToAPI(t *testing.T) {
	// Mock the HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response based on request
	}))
	defer server.Close()

	apiURL = server.URL // Use the mock server URL

	tests := []struct {
		name         string
		prompt       string
		encodedImage string
		confPath     string
		wantErr      bool
	}{
		// Define test cases here
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := SendImageToAPI(tc.prompt, tc.encodedImage, tc.confPath)
			if (err != nil) != tc.wantErr {
				t.Errorf("SendImageToAPI() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
