package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	BOILERPLATE_DIR = "boilerplate"
	TEMPALTE_DIR    = "templates"
	CONFIG_DIR      = "configs"
)

func main() {

	// home directory logic
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	templateLocation := filepath.Join(homeDir, BOILERPLATE_DIR, TEMPALTE_DIR)

	fileName := "hello.txt"
	filePath := filepath.Join(templateLocation, fileName)

	data := []byte("Hello, World!")

	os.MkdirAll(templateLocation, os.ModePerm)

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File wrote successfully at : ", filePath)
}
