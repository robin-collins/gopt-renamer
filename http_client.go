package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const apiURL = "http://example.com/api/upload" // Replace with the actual API URL

// SendImageToAPI takes a base64 encoded image string, sends it to the API,
// and returns the API response as a string.
func SendImageToAPI(encodedImage string) (string, error) {
	// Prepare the JSON payload
	payload := map[string]string{"image": encodedImage}
	jsonPayload, err := json.Marshal(payload)a
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the status code is successful
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("API request failed with status code: " + resp.Status)
	}

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the response body to a string
	responseString := string(respBody)

	return responseString, nil
}
