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
	input, err := fileutil.ScanCSInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	current := input
	for day := 0; day < 80; day++ {
		current = step(current)
	}
	fmt.Println(len(current))
}

func step(input []int) []int {
	result := make([]int, 0)
	for _, n := range input {
		next := n - 1
		if next < 0 {
			next = 6
			result = append(result, 8)
		}
		result = append(result, next)
	}
	return result
}

func partTwo() {
	input, err := fileutil.ScanCSInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	countPerDay := make(map[int]int, 0)
	for i := 0; i <= 8; i++ {
		countPerDay[i] = 0
	}
	for _, n := range input {
		countPerDay[n]++
	}

	for day := 0; day < 256; day++ {
		cleverStep(countPerDay)
	}
	fmt.Println(cleverTotal(countPerDay))
}

func cleverStep(countPerDay map[int]int) {
	zeros := countPerDay[0]
	countPerDay[0] = countPerDay[1]
	countPerDay[1] = countPerDay[2]
	countPerDay[2] = countPerDay[3]
	countPerDay[3] = countPerDay[4]
	countPerDay[4] = countPerDay[5]
	countPerDay[5] = countPerDay[6]
	countPerDay[6] = countPerDay[7] + zeros
	countPerDay[7] = countPerDay[8]
	countPerDay[8] = zeros
}

func cleverTotal(countPerDay map[int]int) int {
	sum := 0
	for _, v := range countPerDay {
		sum += v
	}
	return sum
}
