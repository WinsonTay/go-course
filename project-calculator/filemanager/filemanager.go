package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	var lines []string
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading The Content File Failed")
		return nil, errors.New("could not load text")
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	// Create the file where the data will be written.
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close() // Close the file when the function returns.
	time.Sleep(3 * time.Second)
	// Encode the data as JSON and write it to the file.
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
