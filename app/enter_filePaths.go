package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (a *App) enterFilePaths() []string {

	var (
		filePaths []string
		iterPaths []string
	)

	filePaths = []string{}
	// Prompt user for input
	fmt.Println("Enter file paths (use * for wildcard) or type 'done' to finish:")

	// Read input from standard input in a loop until 'done' is entered
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {

			input := scanner.Text()

			if input == "done" {
				break
			}
			iterPaths = expandPaths(input)
			for _, path := range iterPaths {
				// Check if the file exists and has the .xpm extension
				if isValidFile(path) {
					filePaths = append(filePaths, path)
				} else {
					fmt.Printf("Error: %s is not a valid .xpm file or does not exist\n", path)
				}
			}
			if len(filePaths) > 0 {
				fmt.Println("Selected files:")
				for _, filePath := range filePaths {
					fmt.Printf("* %s\n", filePath)
				}
			}
			fmt.Println("\nEnter more file paths or type 'done' to finish:")
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			break
		}
	}
	return filePaths
}

// expandPaths expands file paths with * wildcard
func expandPaths(input string) []string {
	var expandedPaths []string

	// Split input by space to get individual paths
	paths := strings.Fields(input)

	for _, path := range paths {
		// Check if * is present in the path
		if strings.Contains(path, "*") {
			// Expand wildcard using filepath.Glob
			matches, err := filepath.Glob(path)
			if err != nil {
				fmt.Println("Error expanding wildcard:", err)
				continue
			}
			for _, matchedFile := range matches {
				if isValidFile(matchedFile) {
					expandedPaths = append(expandedPaths, matchedFile)
				}
			}
			// Append expanded paths to the result
		} else if isValidFile(path) {
			// If no wildcard, simply add the path to the result
			expandedPaths = append(expandedPaths, path)
		}
	}

	return expandedPaths
}

// isValidFile checks if the file exists and has the .xpm extension
func isValidFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // File doesn't exist
	}

	// Check if it's a regular file and has the .xpm extension
	return !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".xpm")
}
