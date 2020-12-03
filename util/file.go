package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ReadFile opens the input file if exists
func ReadFile(filePath string) *os.File {
	inputFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return inputFile
}

// CreateOutputFile creates an output file
func CreateOutputFile() *os.File {
	filePath := fmt.Sprintf("../output-%d.txt", time.Now().Unix())

	outputFile, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return outputFile
}
