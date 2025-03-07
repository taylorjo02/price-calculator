package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

var ErrFileOpenFailed = errors.New("unable to open file")
var ErrFileReaderFailed = errors.New("unable to read file contents")
var ErrFileCreateFailed = errors.New("unable to create file")
var ErrJSONEncodeFailed = errors.New("unable to convert data to JSON")

type FileManager struct {
	InputFilePath  string `json:"inputFilePath"`
	OutputFilePath string `json:"outputFilePath"`
}


func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, ErrFileOpenFailed
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, ErrFileReaderFailed
	}

	file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return ErrFileCreateFailed
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return ErrJSONEncodeFailed
	}

	file.Close()

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}