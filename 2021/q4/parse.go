package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() ([]*BingoBoard, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	callsList := scanner.Text()

	boards := make([]*BingoBoard, 0)
	boardSpec := ""
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if boardSpec != "" {
				b := NewBingoBoard(boardSpec)
				boards = append(boards, b)
				boardSpec = ""
			}
			continue
		}
		boardSpec += txt + "\n"
	}
	b := NewBingoBoard(boardSpec)
	boards = append(boards, b)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calls := strings.Split(callsList, ",")
	callsInts := make([]int, 0)
	for _, call := range calls {
		callInt, err := strconv.Atoi(call)
		if err != nil {
			log.Fatal(err)
		}
		callsInts = append(callsInts, callInt)
	}

	return boards, callsInts
}
