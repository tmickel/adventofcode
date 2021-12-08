package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// how many times 2, 4, 3, and 7 length
	input, err := fileutil.ScanStrings("input.txt")
	// input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, x := range input {
		o := strings.Split(x, " | ")
		outs := strings.Fields(o[1])
		for _, i := range outs {
			if len(i) == 2 || len(i) == 4 || len(i) == 3 || len(i) == 7 {
				result++
			}
		}
	}
	fmt.Println(result)
}

type ByYikes []string

func (a ByYikes) Len() int      { return len(a) }
func (a ByYikes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByYikes) Less(i, j int) bool {
	if len(a[i]) == 6 && len(a[j]) == 5 {
		return true
	}
	if len(a[i]) == 5 && len(a[j]) == 6 {
		return false
	}
	return len(a[i]) < len(a[j])
}

func partTwo() {
	input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, x := range input {
		o := strings.Split(x, " | ")
		patterns := strings.Fields(o[0])
		sort.Sort(ByYikes(patterns))
		matches := make(map[string]int, 0)
		byNumber := make(map[int]string, 0)
		for _, pattern := range patterns {
			pattern = sortString(pattern)
			if len(pattern) == 2 {
				// 2-length: 1
				matches[pattern] = 1
				byNumber[1] = pattern
			}
			if len(pattern) == 3 {
				// 3-length: 7
				matches[pattern] = 7
				byNumber[7] = pattern
			}
			if len(pattern) == 4 {
				// 4-length: 4
				matches[pattern] = 4
				byNumber[4] = pattern
			}
			if len(pattern) == 7 {
				// 7-length: 8
				matches[pattern] = 8
				byNumber[8] = pattern
			}
			if len(pattern) == 6 {
				// 6-length, but does not contain both from 2-length (1): 6
				oneParts := strings.Split(byNumber[1], "")
				fourParts := strings.Split(byNumber[4], "")
				if !strings.Contains(pattern, oneParts[0]) || !strings.Contains(pattern, oneParts[1]) {
					matches[pattern] = 6
					byNumber[6] = pattern
				} else if !strings.Contains(pattern, fourParts[0]) ||
					!strings.Contains(pattern, fourParts[1]) ||
					!strings.Contains(pattern, fourParts[2]) ||
					!strings.Contains(pattern, fourParts[3]) {
					// 6-length, but does not contain all from 4-length (4): 0
					matches[pattern] = 0
					byNumber[0] = pattern
				} else {
					// remaining 6-length (contains all from 4-length (4) and 2-length (1)): 9
					matches[pattern] = 9
					byNumber[9] = pattern
				}
			}
			if len(pattern) == 5 {
				// 5-length, and contains both from 2-length (1): 3
				oneParts := strings.Split(byNumber[1], "")
				sixParts := strings.Split(byNumber[6], "")

				if len(sixParts) == 0 {
					log.Panic("no six known")
				}

				if strings.Contains(pattern, oneParts[0]) && strings.Contains(pattern, oneParts[1]) {
					matches[pattern] = 3
					byNumber[3] = pattern
				} else if subset(strings.Split(pattern, ""), sixParts) {
					// 5-length, subset of 6: 5
					matches[pattern] = 5
					byNumber[5] = pattern
				} else {
					matches[pattern] = 2
					byNumber[2] = pattern
				}
			}
		}
		outs := strings.Fields(o[1])
		display := ""
		for _, i := range outs {
			display += fmt.Sprint(matches[sortString(i)])
		}
		out, err := strconv.Atoi(display)
		if err != nil {
			log.Fatal(err)
		}
		result += out
	}
	fmt.Println(result)
}

// https://stackoverflow.com/questions/18879109/subset-check-with-slices-in-go
func subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
