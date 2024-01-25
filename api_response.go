package main

import (
	"encoding/json"
	"errors"
)

// APIResponse represents the structure of the JSON response from the API.
type APIResponse struct { // APIResponse represents the structure of the JSON response from the API.
	Success bool   `json:"success"` // Indicates whether the API request was successful.
	Message string `json:"message"` // Contains any error or informational message from the API.
	Data    struct { // Contains additional data returned by the API.
		ImageURL string `json:"image_url"` // The URL of the image returned by the API.
	} `json:"data"` // Contains additional data returned by the API.
} // End of APIResponse struct

// DecodeAPIResponse decodes the given JSON string into an APIResponse struct.
// It returns the decoded APIResponse and an error if decoding fails.
func DecodeAPIResponse(responseString string) (APIResponse, error) { // DecodeAPIResponse decodes the given JSON string into an APIResponse struct. It returns the decoded APIResponse and an error if decoding fails.
	var apiResponse APIResponse // APIResponse struct

	// Unmarshal the JSON string into the APIResponse struct
	err := json.Unmarshal([]byte(responseString), &apiResponse) // Unmarshal the JSON string into the APIResponse struct
	if err != nil { // Check if there was an error
		return APIResponse{}, errors.New("failed to decode API response: " + err.Error()) // Return the error
	} // End of if statement

	// Check if the API response indicates success
	if !apiResponse.Success { // Check if the API response indicates success
		return APIResponse{}, errors.New("API responded with an error: " + apiResponse.Message) // Return the error
	} // End of if statement

	return apiResponse, nil // Return the decoded APIResponse and no error
}
