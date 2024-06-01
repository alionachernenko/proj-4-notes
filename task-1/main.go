package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileNames := getFileNames()
	notes := getNotes(fileNames)

	var searchInput string

	fmt.Println("What do you want to find?")

	fmt.Scan(&searchInput)

	searchNotes(notes, searchInput)
}

func getFileNames() []string {
	files, err := os.ReadDir("./notes")

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

		if strings.HasSuffix(fileName, ".txt") {
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames
}

func getNotes(fileNames []string) []string {
	var notes []string

	for _, noteName := range fileNames {
		note, err := os.Open(noteName)

		if err != nil {
			fmt.Printf("Error reading note %v", note)
			continue
		}

		defer note.Close()

		scanner := bufio.NewScanner(note)

		for scanner.Scan() {
			line := scanner.Text()
			notes = append(notes, line)
		}
	}

	return notes
}

func searchNotes(notes []string, input string) {
	for _, note := range notes {
		if strings.Contains(strings.ToLower(note), input) {
			fmt.Printf("%v\n\n", note)
		}
	}
}
