package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// input, err := fileutil.ScanStrings("example.txt")
	input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	result := 0

	for _, line := range input {
		stack := make([]string, 0)
		chars := strings.Split(line, "")

		for _, char := range chars {
			if char == "{" || char == "(" || char == "<" || char == "[" {
				stack = append(stack, char)
			} else {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if char == "}" {
					if popped != "{" {
						result += 1197
						break
					}
				}
				if char == "]" {
					if popped != "[" {
						result += 57
						break
					}
				}
				if char == ">" {
					if popped != "<" {
						result += 25137
						break
					}
				}
				if char == ")" {
					if popped != "(" {
						result += 3
						break
					}
				}

			}
		}

	}

	fmt.Println(result)
}

func partTwo() {
	input, err := fileutil.ScanStrings("input.txt")
	// input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}

	remainingAfterPartOne := make([]string, 0)

	for _, line := range input {
		stack := make([]string, 0)
		chars := strings.Split(line, "")

		remaining := true
		for _, char := range chars {
			if char == "{" || char == "(" || char == "<" || char == "[" {
				stack = append(stack, char)
			} else {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if char == "}" {
					if popped != "{" {
						remaining = false
						break
					}
				}
				if char == "]" {
					if popped != "[" {
						remaining = false
						break
					}
				}
				if char == ">" {
					if popped != "<" {
						remaining = false
						break
					}
				}
				if char == ")" {
					if popped != "(" {
						remaining = false
						break
					}
				}

			}
		}

		if remaining {
			remainingAfterPartOne = append(remainingAfterPartOne, line)
		}
	}

	allScores := make([]int, 0)

	for _, line := range remainingAfterPartOne {
		stack := make([]string, 0)
		chars := strings.Split(line, "")

		for _, char := range chars {
			if char == "{" || char == "(" || char == "<" || char == "[" {
				stack = append(stack, char)
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			r := stack[i]
			if r == "(" {
				score = score*5 + 1
			}
			if r == "[" {
				score = score*5 + 2
			}
			if r == "{" {
				score = score*5 + 3
			}
			if r == "<" {
				score = score*5 + 4
			}
		}
		allScores = append(allScores, score)
	}

	sort.Ints(allScores)
	fmt.Println(allScores[len(allScores)/2])
}
