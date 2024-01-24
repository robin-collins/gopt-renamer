package main

import (
    "testing"
)

func TestDecodeAPIResponse_Success(t *testing.T) {
    responseString := `{"success": true, "message": "OK", "data": {"image_url": "http://example.com/image.jpg"}}`
    expected := APIResponse{
        Success: true,
        Message: "OK",
        Data: struct {
            ImageURL string `json:"image_url"`
        }{
            ImageURL: "http://example.com/image.jpg",
        },
    }

    result, err := DecodeAPIResponse(responseString)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestDecodeAPIResponse_Failure(t *testing.T) {
    responseString := `{"success": false, "message": "Something went wrong", "data": {}}`
    expected := APIResponse{}

    result, err := DecodeAPIResponse(responseString)
    if err == nil || err.Error() != "API responded with an error: Something went wrong" {
        t.Errorf("Expected error message 'API responded with an error: Something went wrong', got %v", err)
    }

    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestDecodeAPIResponse_InvalidJSON(t *testing.T) {
    responseString := `{"success": false, "message": "Something went wrong", "data": {}`
    expected := APIResponse{}

    result, err := DecodeAPIResponse(responseString)
    if err == nil || err.Error() != "failed to decode API response: unexpected end of JSON input" {
        t.Errorf("Expected error message 'failed to decode API response: unexpected end of JSON input', got %v", err)
    }

    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}