package filemanager

import (
	"bufio"
	"errors"
	"os"
)

var ErrFileOpenFailed = errors.New("unable to open file")
var ErrFileReaderFailed = errors.New("unable to read file contents")

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