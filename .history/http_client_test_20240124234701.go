package main

import (
    "errors"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestSendImageToAPI_MissingAPIKey(t *testing.T) {
    // Unset the environment variable for this test
    os.Unsetenv("OPENAI_API_KEY")

    _, err := SendImageToAPI("test prompt", "test image")
    expectedErr := errors.New("OPENAI_API_KEY is not set in environment variables")

    if err == nil || err.Error() != expectedErr.Error() {
        t.Errorf("Expected error message '%v', got '%v'", expectedErr, err)
    }
}

func TestSendImageToAPI_APIRequestFailure(t *testing.T) {
    // Set the environment variable for this test
    os.Setenv("OPENAI_API_KEY", "test key")

    // Create a test server that always responds with a 500 status code
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))
    defer testServer.Close()

    // Temporarily replace the API URL with the test server's URL
    oldAPIURL := apiURL
    apiURL = testServer.URL
    defer func() { apiURL = oldAPIURL }() // Restore original API URL after this test

    _, err := SendImageToAPI("test prompt", "test image")
    expectedErr := errors.New("API request failed with status code: 500 Internal Server Error")

    if err == nil || err.Error() != expectedErr.Error() {
        t.Errorf("Expected error message '%v', got '%v'", expectedErr, err)
    }
}

func TestSendImageToAPI_Success(t *testing.T) {
    // Set the environment variable for this test
    os.Setenv("OPENAI_API_KEY", "test key")

    // Create a test server that always responds with a successful response
    testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"choices": [{"message": {"content": "test content"}}]}`))
    }))
    defer testServer.Close()

    // Temporarily replace the API URL with the test server's URL
    oldAPIURL := apiURL
    apiURL = testServer.URL
    defer func() { apiURL = oldAPIURL }() // Restore original API URL after this test

    result, err := SendImageToAPI("test prompt", "test image")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if result != "test content" {
        t.Errorf("Expected 'test content', got '%v'", result)
    }
}