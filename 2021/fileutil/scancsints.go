package fileutil

import (
	"os"
	"strconv"
	"strings"
)

// ScanCSInts reads a single line of comma-separated ints into an int slice
func ScanCSInts(filename string) ([]int, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ns := strings.Split(strings.TrimSpace(string(file)), ",")
	result := make([]int, 0)
	for _, n := range ns {
		parsed, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		result = append(result, parsed)
	}

	return result, nil
}
