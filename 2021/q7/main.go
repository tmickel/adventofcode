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
	crabs, err := fileutil.ScanCSInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, crab := range crabs {
		if crab > max {
			max = crab
		}
	}

	leastFuel := 100000000
	for pos := 0; pos <= max; pos++ {
		fuel := 0
		for _, crab := range crabs {
			fuel += abs(pos - crab)
		}
		if fuel <= leastFuel {
			leastFuel = fuel
		}
	}
	fmt.Println(leastFuel)
}

func partTwo() {
	crabs, err := fileutil.ScanCSInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, crab := range crabs {
		if crab > max {
			max = crab
		}
	}

	leastFuel := 100000000
	for pos := 0; pos <= max; pos++ {
		fuel := 0
		for _, crab := range crabs {
			nSteps := abs(pos - crab)
			// triangular numbers
			fuel += nSteps * (nSteps + 1) / 2
		}
		if fuel <= leastFuel {
			leastFuel = fuel
		}
	}
	fmt.Println(leastFuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
