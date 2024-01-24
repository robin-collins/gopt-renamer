package main

import (
	"encoding/json"
	"errors"
)

// APIResponse represents the structure of the JSON response from the API.
type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ImageURL string `json:"image_url"`
	} `json:"data"`
}

// DecodeAPIResponse takes a JSON string response from the API and decodes it into an APIResponse struct.
func DecodeAPIResponse(responseString string) (APIResponse, error) {
	var apiResponse APIResponse

	// Unmarshal the JSON string into the APIResponse struct
	err := json.Unmarshal([]byte(responseString), &apiResponse)
	if err != nil {
		return APIResponse{}, errors.New("failed to decode API response: " + err.Error())
	}

	// Check if the API response indicates success
	if !apiResponse.Success {
		return APIResponse{}, errors.New("API responded with an error: " + apiResponse.Message)
	}

	return apiResponse, nil
}
