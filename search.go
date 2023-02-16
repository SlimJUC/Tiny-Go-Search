package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Prompt user for search directory and substring
	fmt.Print("Enter the search directory: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchDir := scanner.Text()

	fmt.Print("Enter the search string: ")
	scanner.Scan()
	searchString := scanner.Text()

	// Walk the directory tree and search for files
	totalFiles := 0
	matchingFiles := 0
	filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalFiles++
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if err := scanner.Err(); err != nil {
					return err
				}
				if strings.Contains(scanner.Text(), searchString) {
					matchingFiles++
					fmt.Printf("Found \"%s\" in file: %s\n", searchString, path)
					break
				}
			}
		}
		return nil
	})

	// Print search summary
	fmt.Printf("Searched %d files and found \"%s\" in %d files\n", totalFiles, searchString, matchingFiles)
}
