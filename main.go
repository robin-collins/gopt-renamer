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
		if !*silentFlag {
			fmt.Println("Usage of the program:")
			fmt.Println("  --image=\"path/to/image\" : Specify the image file path to be encoded and sent to the API")
			fmt.Println("  --help : Display this help information")
			fmt.Println("  --force : Force rename without prompt")
			fmt.Println("  --silent : No output except exit codes") // New help info for silent flag
		}
		return
	}

	// Check if the image flag is provided
	if *imageFlag == "" {
		if !*silentFlag {
			log.Fatal("Error: Image file path is required. Use --image=\"path/to/image\" to specify the image file.")
		} else {
			os.Exit(1)
		}
	}

	// Check if the image file exists
	if _, err := os.Stat(*imageFlag); os.IsNotExist(err) {
		if !*silentFlag {
			log.Fatalf("Error: The file %s does not exist.", *imageFlag)
		} else {
			os.Exit(1)
		}
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(*imageFlag)
	if err != nil {
		if !*silentFlag {
			log.Fatalf("Error reading the image file: %s", err)
		} else {
			os.Exit(1)
		}
	}

	// Base64 encode the image data
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Prompt to send to the API
	prompt := "Analyze the content and context of the image. Generate an informative and descriptive file name that reflects the key elements or subject matter depicted in the image, for photographs consider who, what, where, when and why. Attempt to keep it Main_Title-Main_Topic, for example Google_Chrome_Browser_reddit-cat_information_page, alternatively Who_What_Where_When_Why-Title. Reply only with the file name, excluding the file extension. Do not offer reasoning or justification for the chosen name."

	// Send the encoded image to the API
	response, err := SendImageToAPI(prompt, encodedImage)
	if err != nil {
		// Existing error handling code
	} else {
		// Check if the response length is less than 10 characters
		if len(response) < 10 {
			if !*silentFlag {
				fmt.Println("API response is too short for a file name. Exiting.")
			}
			os.Exit(0)
		}
	}

	// Ask the user if they want to rename the file
	dir, file := filepath.Split(*imageFlag)
	ext := filepath.Ext(file)
	newFileName := fmt.Sprintf("%s%s", response, ext)
	newFilePath := filepath.Join(dir, newFileName)

	if *forceFlag || *silentFlag {
		// If force flag or silent flag is true, rename the file without prompt
		err := os.Rename(*imageFlag, newFilePath)
		if err != nil {
			if !*silentFlag {
				log.Fatalf("Error renaming the file: %s", err)
			} else {
				os.Exit(1)
			}
		}
		if !*silentFlag {
			fmt.Printf("File renamed to '%s'\n", newFilePath)
		}
	} else {
		// Otherwise, prompt the user
		fmt.Printf("Do you want to rename the file from '%s' to '%s'? (yes/no): ", file, newFileName)
		var userResponse string
		_, err = fmt.Scan(&userResponse)
		if err != nil {
			if !*silentFlag {
				log.Fatalf("Error reading user response: %s", err)
			} else {
				os.Exit(1)
			}
		}

		if userResponse == "yes" {
			err := os.Rename(*imageFlag, newFilePath)
			if err != nil {
				if !*silentFlag {
					log.Fatalf("Error renaming the file: %s", err)
				} else {
					os.Exit(1)
				}
			}
			if !*silentFlag {
				fmt.Printf("File renamed to '%s'\n", newFilePath)
			}
		} else {
			if !*silentFlag {
				fmt.Println("File rename operation cancelled.")
			}
		}
	}

	// If silent flag is set, suppress all output and exit with code 0 upon success
	if *silentFlag {
		os.Exit(0)
	}
}
