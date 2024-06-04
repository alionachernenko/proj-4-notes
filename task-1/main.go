package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filesDirectory := flag.String("dir", "./notes", "files directory")

	fileNames := getFileNames(filesDirectory, ".txt")
	notes := getFiles(filesDirectory, fileNames)

	var searchInput string

	fmt.Println("What do you want to find?")

	fmt.Scan(&searchInput)

	searchFiles(notes, searchInput)
}

func getFileNames(dir *string, ext string) []string {

	files, err := os.ReadDir(*dir)

	if err != nil {
		fmt.Printf("Error reading directory :(")
		return nil
	}

	var fileNames []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if strings.HasSuffix(fileName, ext) {
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames
}

func getFiles(dir *string, fileNames []string) []string {
	var files []string

	for _, fileName := range fileNames {
		file, err := os.Open(fmt.Sprintf("%v/%v", dir, fileName))

		if err != nil {
			fmt.Printf("Error reading file %v: %v", fileName, err)
			continue
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			files = append(files, line)
		}
	}

	return files
}

func searchFiles(files []string, input string) {
	for _, file := range files {
		if strings.Contains(strings.ToLower(file), strings.ToLower(input)) {
			fmt.Printf("%v\n\n", file)
		}
	}
}
