package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiURL = "https://api.openai.com/v1/chat/completions" // API URL

// getResponse takes a string prompt and a base64 image string and returns the response from the API.
func SendImageToAPI(prompt, encodedImage string) (string, error) {
	// Retrieve the OpenAI API key from an environment variable
	openAIKey := os.Getenv("OPENAI_API_KEY") // OpenAI API key
	if openAIKey == "" { 				   // Check if the API key is empty
		return "", errors.New("OPENAI_API_KEY is not set in environment variables") // Return an error
	}

	// Prepare the JSON payload
	payload := map[string]interface{}{ // Create a map of string to interface
		"model": "gpt-4-vision-preview", // Set the model
		"messages": []map[string]interface{}{ // Set the messages
			{ // Create a map of string to interface
				"role": "user", // Set the role
				"content": []map[string]interface{}{ // Set the content
					{ // Create a map of string to interface
						"type": "text", // Set the type
						"text": prompt, // Set the text
					}, // End of map
					{ // Create a map of string to interface
						"type": "image_url", // Set the type
						"image_url": map[string]string{ // Set the image URL
							"url": fmt.Sprintf("data:image/jpeg;base64,%s", encodedImage), // Set the URL
						}, // End of map
					}, // End of map
				}, // End of map
			}, // End of map
		}, // End of map
		"max_tokens":  300, // Set the max tokens
		"temperature": 0.9, // Set the temperature
	} // End of map
	jsonPayload, err := json.Marshal(payload) // Marshal the JSON payload
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload)) // Create a new HTTP request
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json") // Set the content type
	req.Header.Set("Authorization", "Bearer "+openAIKey) // Set the authorization

	// Perform the HTTP request
	client := &http.Client{} // Create a new HTTP client
	resp, err := client.Do(req) // Perform the HTTP request
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement
	defer resp.Body.Close() // Close the response body

	// Check if the status code is successful
	if resp.StatusCode != http.StatusOK { // Check if the status code is not OK
		return "", errors.New("API request failed with status code: " + resp.Status) // Return an error
	} // End of if statement

	// Read the response body
	respBody, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement

	// Unmarshal the JSON response
	var result map[string]interface{} // Create a map of string to interface
	err = json.Unmarshal(respBody, &result) // Unmarshal the JSON response
	if err != nil { // Check if there was an error
		return "", err // Return the error
	} // End of if statement

	// Access the desired fields
	choices := result["choices"].([]interface{}) // Get the choices
	firstChoice := choices[0].(map[string]interface{}) // Get the first choice
	message := firstChoice["message"].(map[string]interface{}) // Get the message
	content := message["content"].(string) // Get the content

	return content, nil // Return the content
}
