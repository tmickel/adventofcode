package fileutil

import (
	"bufio"
	"os"
)

// ScanStrings reads a newline-separated file into a string array
func ScanStrings(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ss := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ss = append(ss, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}
