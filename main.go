package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	// Prompt to send to the API
	prompt := "Analyze the content and context of the screenshot and generate an informative and descriptive file name that reflects the key elements or subject matter depicted, Attempt to keep it Main_Title-Main_Topic, for example Google_Chrome_Browser_reddit-cat_information_page. Reply only with the file name, excluding the file extension. Do not offer reasoning or justification for the chosen name."

	// Send the encoded image to the API
	response, err := SendImageToAPI(prompt, encodedImage)
	if err != nil {
		log.Fatalf("Error sending the image to the API: %s", err)
	}

	// Ask the user if they want to rename the file
	dir, file := filepath.Split(*imageFlag)
	ext := filepath.Ext(file)
	// baseName := strings.TrimSuffix(file, ext)
	newFileName := fmt.Sprintf("%s%s", response, ext)
	newFilePath := filepath.Join(dir, newFileName)

	fmt.Printf("Do you want to rename the file from '%s' to '%s'? (yes/no): ", file, newFileName)
	var userResponse string
	_, err = fmt.Scan(&userResponse)
	if err != nil {
		log.Fatalf("Error reading user response: %s", err)
	}

	// If the user responds with 'yes', rename the file
	if userResponse == "yes" {
		err := os.Rename(*imageFlag, newFilePath)
		if err != nil {
			log.Fatalf("Error renaming the file: %s", err)
		}
		fmt.Printf("File renamed to '%s'\n", newFilePath)
	} else {
		fmt.Println("File rename operation cancelled.")
	}
}
