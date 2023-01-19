package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// this code is generated
// TODO: review
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <template_name> <new_name>")
		os.Exit(1)
	}

	templateName := os.Args[1]
	newName := strings.ReplaceAll(os.Args[2], "-", "_")

	// Convert the new name to CamelCase for the ##name## placeholder
	newNameCamel := strings.ReplaceAll(
		strings.Title(
			strings.ReplaceAll(newName, "_", " "),
		),
		" ",
		"",
	)

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	// Define the template file names
	templateFile := fmt.Sprintf("%s.tmpl", templateName)
	templateTestFile := fmt.Sprintf("%s_test.tmpl", templateName)

	// Define the target file names
	targetFile := fmt.Sprintf("leetcode/%s.go", newName)
	targetTestFile := fmt.Sprintf("leetcode/%s_test.go", newName)

	// Read the template files
	templateBytes, err := ioutil.ReadFile(filepath.Join("templates", templateFile))
	if err != nil {
		fmt.Println("Error reading template file:", err)
		os.Exit(1)
	}

	templateTestBytes, err := ioutil.ReadFile(filepath.Join("templates", templateTestFile))
	if err != nil {
		fmt.Println("Error reading template test file:", err)
		os.Exit(1)
	}

	// Replace the ##name## placeholder with the new name
	newBytes := bytes.ReplaceAll(templateBytes, []byte("##name##"), []byte(newNameCamel))
	newTestBytes := bytes.ReplaceAll(templateTestBytes, []byte("##name##"), []byte(newNameCamel))

	// Write the new files
	err = ioutil.WriteFile(filepath.Join(currentDir, targetFile), newBytes, 0644)
	if err != nil {
		fmt.Println("Error writing new file:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(filepath.Join(currentDir, targetTestFile), newTestBytes, 0644)
	if err != nil {
		fmt.Println("Error writing new test file:", err)
		os.Exit(1)
	}
}
