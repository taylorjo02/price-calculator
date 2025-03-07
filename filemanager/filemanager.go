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

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
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

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)

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