package main

import (
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	boards, callsList := parse()
	for _, call := range callsList {
		for _, b := range boards {
			b.CallNumber(call)
			if b.Winner() {
				fmt.Println(b.SumUncalled() * call)
				return
			}
		}
	}
}

func partTwo() {
	boards, callsList := parse()
	for _, call := range callsList {
		for _, b := range boards {
			b.CallNumber(call)
			allWon := true
			for _, bb := range boards {
				allWon = allWon && bb.Winner()
			}
			if allWon {
				fmt.Println(b.SumUncalled() * call)
				return
			}
		}
	}
}
