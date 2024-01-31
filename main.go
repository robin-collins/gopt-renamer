// Package main implements a command-line tool for renaming image files.
// The tool encodes the specified image file in Base64, sends it to an external API for analysis,
// and renames the file based on the API's response.
// Author: github-copilot and gpt-4-vision-preview
// Maintainer: Robin Collins (https://github.com/robin-collins)
// Created: 2024-01-24
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// main is the entry point of the program. It parses command line flags,
// validates inputs, processes the image file, and finally renames it based
// on the response from an external API.
func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to get OpenAI API Key: %v", err)
	}
	// Define command line flags
	imageFlag := flag.String("image", "", "Path to the image file to be processed")
	helpFlag := flag.Bool("help", false, "Display help information")
	forceFlag := flag.Bool("force", false, "Force rename without prompt")
	silentFlag := flag.Bool("silent", false, "No output except exit codes") // New silent flag

	// Parse the flags
	flag.Parse()

	// Display help information if the help flag is provided
	if *helpFlag {
		displayHelpInformation(*silentFlag)
		return
	}

	// Validate the image file path
	if *imageFlag == "" {
		handleError("Error: Image file path is required. Use --image=\"path/to/image\" to specify the image file.", *silentFlag)
	}

	// Check for the existence of the image file
	if _, err := os.Stat(*imageFlag); os.IsNotExist(err) {
		handleError(fmt.Sprintf("Error: The file %s does not exist.", *imageFlag), *silentFlag)
	}

	// Process the image file
	imageData, err := os.ReadFile(*imageFlag)
	if err != nil {
		handleError(fmt.Sprintf("Error reading the image file: %s", err), *silentFlag)
	}

	// Encode the image file in Base64
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Define the prompt for the API
	prompt := "Analyze the content and context of the image. Generate an informative and descriptive file name that reflects the key elements or subject matter depicted in the image, for photographs consider who, what, where, when and why. Attempt to keep it Main_Title-Main_Topic, for example Google_Chrome_Browser_reddit-cat_information_page, alternatively Who_What_Where_When_Why-Title. Reply only with the file name, excluding the file extension. Do not offer reasoning or justification for the chosen name."

	if config.OpenAI_API_Key == "" {
        log.Fatal("Error: OpenAI API Key not set")
    }

	// Send the encoded image to the API and obtain a response
	response, err := SendImageToAPI(prompt, encodedImage)
	if err != nil {
		handleError("Error sending image to API", *silentFlag)
	}

	// Validate the API's response
	if len(response) < 10 {
		if !*silentFlag {
			fmt.Println("API response is too short for a file name. Exiting.")
		}
		os.Exit(0)
	}

	// Rename the file based on the API's response
	renameFile(*imageFlag, response, *forceFlag, *silentFlag)
}
// displayHelpInformation prints the help information for the program.
// If the silentFlag is set, it suppresses the output.
func displayHelpInformation(silentFlag bool) {
	if !silentFlag {
		fmt.Println("Usage of the program:")
		fmt.Println("  --image=\"path/to/image\" : Specify the image file path to be encoded and sent to the API")
		fmt.Println("  --help : Display this help information")
		fmt.Println("  --force : Force rename without prompt")
		fmt.Println("  --silent : No output except exit codes") // New help info for silent flag
	}
}
// handleError logs or prints the specified error message.
// If the silentFlag is set, it exits with status code 1.
func handleError(errorMessage string, silentFlag bool) {
	if !silentFlag {
		log.Fatal(errorMessage)
	} else {
		os.Exit(1)
	}
}
// renameFile renames the specified image file based on the response.
// It considers forceFlag and silentFlag for forced renaming and suppressing output, respectively.
func renameFile(imageFlag, response string, forceFlag, silentFlag bool) {
	dir, file := filepath.Split(imageFlag)
	ext := filepath.Ext(file)
	newFileName := fmt.Sprintf("%s%s", response, ext)
	newFilePath := filepath.Join(dir, newFileName)

	if forceFlag || silentFlag {
		// If force flag or silent flag is true, rename the file without prompt
		err := os.Rename(imageFlag, newFilePath)
		if err != nil {
			handleError(fmt.Sprintf("Error renaming the file: %s", err), silentFlag)
		}
		if !silentFlag {
			fmt.Printf("File renamed to '%s'\n", newFilePath)
		}
	} else {
		// Otherwise, prompt the user
		fmt.Printf("Do you want to rename the file from '%s' to '%s'? (yes/no): ", file, newFileName)
		var userResponse string
		_, err := fmt.Scan(&userResponse)
		if err != nil {
			handleError(fmt.Sprintf("Error reading user response: %s", err), silentFlag)
		}

		if userResponse == "yes" {
			err := os.Rename(imageFlag, newFilePath)
			if err != nil {
				handleError(fmt.Sprintf("Error renaming the file: %s", err), silentFlag)
			}
			if !silentFlag {
				fmt.Printf("File renamed to '%s'\n", newFilePath)
			}
		} else {
			if !silentFlag {
				fmt.Println("File rename operation cancelled.")
			}
		}
	}

	// If silent flag is set, suppress all output and exit with code 0 upon success
	if silentFlag {
		os.Exit(0)
	}
}