// image_processor_test.go
package main

import (
	"os"
	"testing"
)

func TestProcessImage(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_image.*.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write some data to the file
	tempFile.WriteString("test data")
	tempFile.Close()

	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Valid File",
			filePath: tempFile.Name(),
			wantErr:  false,
		},
		{
			name:     "Invalid File",
			filePath: "nonexistent.jpg",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ProcessImage(tc.filePath)
			if (err != nil) != tc.wantErr {
				t.Errorf("ProcessImage() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
