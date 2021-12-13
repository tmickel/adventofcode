package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	xs, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Gamma(xs) * Epsilon(xs))
}

func partTwo() {
	xs, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Oxygen(xs) * CO2Scrubber(xs))
}

func Gamma(ns []string) int {
	result := 0
	bitCount := len(ns[0])
	for bitIndex := bitCount - 1; bitIndex >= 0; bitIndex-- {
		result |= MostCommonBit(BitsFromPosition(ns, bitIndex), 1) << (bitCount - bitIndex - 1)
	}
	return result
}

func Epsilon(ns []string) int {
	result := 0
	bitCount := len(ns[0])
	for bitIndex := bitCount - 1; bitIndex >= 0; bitIndex-- {
		result |= LeastCommonBit(BitsFromPosition(ns, bitIndex), 1) << (bitCount - bitIndex - 1)
	}
	return result

}

func Oxygen(ns []string) int {
	remaining := make([]string, 0)
	remaining = append(remaining, ns...)

	consideredBit := 0

	for len(remaining) > 1 {
		goalBit := MostCommonBit(BitsFromPosition(remaining, consideredBit), 1)
		newFiltered := make([]string, 0)
		for _, x := range remaining {
			bit := BitAtPosition(x, consideredBit)
			if bit == goalBit {
				newFiltered = append(newFiltered, x)
			}
		}

		remaining = newFiltered
		consideredBit++
	}
	result := remaining[0]
	r, err := strconv.ParseInt(result, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(r)
}

func CO2Scrubber(ns []string) int {
	remaining := make([]string, 0)
	remaining = append(remaining, ns...)

	consideredBit := 0

	for len(remaining) > 1 {
		goalBit := LeastCommonBit(BitsFromPosition(remaining, consideredBit), 0)
		newFiltered := make([]string, 0)
		for _, x := range remaining {
			bit := BitAtPosition(x, consideredBit)
			if bit == goalBit {
				newFiltered = append(newFiltered, x)
			}
		}

		remaining = newFiltered
		consideredBit++
	}
	result := remaining[0]
	r, err := strconv.ParseInt(result, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(r)
}

func MostCommonBit(bits []int, tiebreaker int) int {
	sum := 0
	for _, bit := range bits {
		sum += bit
	}
	if sum*2 == len(bits) {
		return tiebreaker
	}
	if sum*2 > len(bits) {
		return 1
	}
	return 0
}

func LeastCommonBit(bits []int, tiebreaker int) int {
	sum := 0
	for _, bit := range bits {
		sum += bit
	}
	if sum*2 == len(bits) {
		return tiebreaker
	}
	if sum*2 < len(bits) {
		return 1
	}
	return 0
}

// pos is from the left
func BitsFromPosition(ns []string, pos int) []int {
	result := make([]int, 0)
	for _, n := range ns {
		bit := BitAtPosition(n, pos)
		result = append(result, bit)
	}
	return result
}

// pos is from the left
func BitAtPosition(n string, pos int) int {
	bit, err := strconv.Atoi(string(n[pos]))
	if err != nil {
		log.Fatal(err)
	}
	return bit
}
