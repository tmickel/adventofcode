package main

import (
	"fmt"
	"log"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	vals, err := fileutil.ScanInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(IncreasesCount(vals))
}

func IncreasesCount(numbers []int) int {
	previous := -1
	increases := 0
	for i, n := range numbers {
		if i > 0 && n > previous {
			increases++
		}
		previous = n
	}
	return increases
}

func partTwo() {
	vals, err := fileutil.ScanInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(IncreasesCount(WindowSums(vals, 3)))
}

func WindowSums(numbers []int, windowSize int) []int {
	windows := make([]int, 0)
	for i := 0; i < len(numbers)+1-windowSize; i++ {
		sum := 0
		for j := 0; j < windowSize; j++ {
			sum += numbers[i+j]
		}
		windows = append(windows, sum)
	}
	return windows
}
