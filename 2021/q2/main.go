package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	in, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hpos := 0
	depth := 0
	for _, instruction := range in {
		parts := strings.Split(instruction, " ")
		cmd := parts[0]
		x, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("invalid amount: %v", err)
		}

		if cmd == "forward" {
			hpos += x
			continue
		}
		if cmd == "up" {
			depth -= x
			continue
		}
		if cmd == "down" {
			depth += x
			continue
		}
		log.Fatalf("invalid command: %s", cmd)
	}

	fmt.Println(hpos * depth)
}

func partTwo() {
	in, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hpos := 0
	depth := 0
	aim := 0
	for _, instruction := range in {
		parts := strings.Split(instruction, " ")
		cmd := parts[0]
		x, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("invalid amount: %v", err)
		}

		if cmd == "forward" {
			hpos += x
			depth += aim * x
			continue
		}
		if cmd == "up" {
			aim -= x
			continue
		}
		if cmd == "down" {
			aim += x
			continue
		}
		log.Fatalf("invalid command: %s", cmd)
	}

	fmt.Println(hpos * depth)
}
