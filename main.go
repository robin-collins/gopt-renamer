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
	forceFlag := flag.Bool("force", false, "Force rename without prompt")
	silentFlag := flag.Bool("silent", false, "No output except exit codes") // New silent flag

	// Parse the flags
	flag.Parse()

	// If help flag is provided, display help information
	if *helpFlag {
		displayHelpInformation(*silentFlag)
		return
	}

	// Check if the image flag is provided
	if *imageFlag == "" {
		handleError("Error: Image file path is required. Use --image=\"path/to/image\" to specify the image file.", *silentFlag)
	}

	// Check if the image file exists
	if _, err := os.Stat(*imageFlag); os.IsNotExist(err) {
		handleError(fmt.Sprintf("Error: The file %s does not exist.", *imageFlag), *silentFlag)
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(*imageFlag)
	if err != nil {
		handleError(fmt.Sprintf("Error reading the image file: %s", err), *silentFlag)
	}

	// Base64 encode the image data
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Prompt to send to the API
	prompt := "Analyze the content and context of the image. Generate an informative and descriptive file name that reflects the key elements or subject matter depicted in the image, for photographs consider who, what, where, when and why. Attempt to keep it Main_Title-Main_Topic, for example Google_Chrome_Browser_reddit-cat_information_page, alternatively Who_What_Where_When_Why-Title. Reply only with the file name, excluding the file extension. Do not offer reasoning or justification for the chosen name."

	// Send the encoded image to the API
	response, err := SendImageToAPI(prompt, encodedImage)
	if err != nil {
		handleError("Error sending image to API", *silentFlag)
	}

	// Check if the response length is less than 10 characters
	if len(response) < 10 {
		if !*silentFlag {
			fmt.Println("API response is too short for a file name. Exiting.")
		}
		os.Exit(0)
	}

	// Rename the file
	renameFile(*imageFlag, response, *forceFlag, *silentFlag)
}

func displayHelpInformation(silentFlag bool) {
	if !silentFlag {
		fmt.Println("Usage of the program:")
		fmt.Println("  --image=\"path/to/image\" : Specify the image file path to be encoded and sent to the API")
		fmt.Println("  --help : Display this help information")
		fmt.Println("  --force : Force rename without prompt")
		fmt.Println("  --silent : No output except exit codes") // New help info for silent flag
	}
}

func handleError(errorMessage string, silentFlag bool) {
	if !silentFlag {
		log.Fatal(errorMessage)
	} else {
		os.Exit(1)
	}
}

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
