package main

import (
	"encoding/base64"
	"os"
)

// ProcessImage takes a file path as input, checks if the file exists,
// reads the file, and returns the base64 encoded string of the file content.
func ProcessImage(filePath string) (string, error) {
	// Check if the image file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}

	// Read the image file
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Base64 encode the image data
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	return encodedImage, nil
}
