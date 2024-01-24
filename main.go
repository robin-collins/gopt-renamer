package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Define command line flags
	imageFlag := flag.String("image", "", "Path to the image file to be processed")
	helpFlag := flag.Bool("help", false, "Display help information")

	// Parse the flags
	flag.Parse()

	// If help flag is provided, display help information
	if *helpFlag {
		fmt.Println("Usage of the program:")
		fmt.Println("  --image=\"path/to/image\" : Specify the image file path to be encoded and sent to the API")
		fmt.Println("  --help : Display this help information")
		return
	}

	// Check if the image flag is provided
	if *imageFlag == "" {
		log.Fatal("Error: Image file path is required. Use --image=\"path/to/image\" to specify the image file.")
	}

	// Check if the image file exists
	if _, err := os.Stat(*imageFlag); os.IsNotExist(err) {
		log.Fatalf("Error: The file %s does not exist.", *imageFlag)
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(*imageFlag)
	if err != nil {
		log.Fatalf("Error reading the image file: %s", err)
	}

	// Base64 encode the image data
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Send the encoded image to the API
	response, err := SendImageToAPI(encodedImage)
	if err != nil {
		log.Fatalf("Error sending the image to the API: %s", err)
	}

	// Print the API response
	fmt.Println("API Response:", response)
}
