package main

import (
	"encoding/base64"
	"os"
)

// ProcessImage takes a file path as input, checks if the file exists,
// reads the file, and returns the base64 encoded string of the file content.
func ProcessImage(filePath string) (string, error) { // ProcessImage takes a file path as input, checks if the file exists, reads the file, and returns the base64 encoded string of the file content.
	// Check if the image file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) { // Check if the image file exists
		return "", err // Return the error
	} // End of if statement

	// Read the image file
	imageData, err := os.ReadFile(filePath) // Read the image file
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement

	// Base64 encode the image data
	encodedImage := base64.StdEncoding.EncodeToString(imageData) // Base64 encode the image data

	return encodedImage, nil // Return the encoded image and no error
} // End of ProcessImage function
