package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error)  {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)	

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Failed to open file")
	}

	// file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to return file")
	}

	// when executing a function with defer
	// Go will not execute the code right away but
	// only when the surrounding func or method finished
	// either because of an error or because it is done
	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to JSON")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager  {
	return FileManager {
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}