package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: hexreplacer <file> <oldHexString> <newHexString>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	oldHex := os.Args[2]
	newHex := os.Args[3]

	oldBytes, err := hex.DecodeString(oldHex)
	if err != nil {
		fmt.Println("Error decoding old hex string:", err)
		os.Exit(1)
	}

	newBytes, err := hex.DecodeString(newHex)
	if err != nil {
		fmt.Println("Error decoding new hex string:", err)
		os.Exit(1)
	}

	// Read the original file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close() // Close the file after use

	// Read file content into a byte slice
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Check if the old hex string exists in the file
	if !strings.Contains(string(fileContent), string(oldBytes)) {
		fmt.Println("Error: Hex string not found in the file.")
		os.Exit(1)
	}

	// Create a backup file
	backupFilePath := filePath + ".bak"
	err = os.WriteFile(backupFilePath, fileContent, 0644)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		os.Exit(1)
	}

	// Replace hex string in memory
	newContent := strings.Replace(string(fileContent), string(oldBytes), string(newBytes), -1)

	// Write the modified content back to the original file
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Hex string replaced successfully! Backup created at:", backupFilePath)
}
