package fileutil

import (
	"bufio"
	"os"
	"strconv"
)

// ScanInts opens a newline-separated file and returns its contents parsed into an int slice
func ScanInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ints := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ints, nil
}
