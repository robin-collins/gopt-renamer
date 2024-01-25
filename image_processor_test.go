package main

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestProcessImage_NonExistentFile(t *testing.T) {
    _, err := ProcessImage("non_existent_file.jpg")
    if err == nil || !os.IsNotExist(err) {
        t.Errorf("Expected error for non-existent file, got %v", err)
    }
}

func TestProcessImage_Success(t *testing.T) {
    // Create a temporary file
    tmpfile, err := os.CreateTemp("", "example.*.jpg")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name()) // clean up

    text := []byte("Hello, World!")
    if _, err := tmpfile.Write(text); err != nil {
        t.Fatal(err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    // Process the temporary file
    result, err := ProcessImage(tmpfile.Name())
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    // Check the result
    expectedResult := base64.StdEncoding.EncodeToString(text)
    if result != expectedResult {
        t.Errorf("Expected %v, got %v", expectedResult, result)
    }
}