package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	input, _ := fileutil.ScanStrings("input.txt")
	algorithm := input[0]
	litPixels := make(map[string]bool)
	toConsider := make(map[string]bool)

	y := 0
	for i := 2; i < len(input); i++ {
		s := strings.Split(input[i], "")
		for x, s := range s {
			if s == "#" {
				litPixels[key(x, y)] = true
			}
			toConsider[key(x, y)] = true
		}
		y += 1
	}
	for k := range litPixels {
		x, y := parse(k)
		for xx := x - 200; xx <= x+200; xx++ {
			for yy := y - 200; yy <= y+200; yy++ {
				toConsider[key(xx, yy)] = true
			}
		}
	}

	output := make(map[string]bool)
	for k := range toConsider {
		x, y := parse(k)
		if lookupOutput(algorithm, readInputNumberAtPixel(litPixels, x, y)) {
			output[k] = true
		}
	}

	// draw(output, -5, -5, 10, 8)
	// log.Println(output)
	// log.Println(len(output))

	toConsider2 := make(map[string]bool, 0)
	for k := range output {
		toConsider2[k] = true
	}

	for k := range output {
		x, y := parse(k)
		for xx := x - 2; xx <= x+2; xx++ {
			for yy := y - 2; yy <= y+2; yy++ {
				toConsider2[key(xx, yy)] = true
			}
		}
	}

	output2 := make(map[string]bool)
	for k := range toConsider2 {
		x, y := parse(k)
		if lookupOutput(algorithm, readInputNumberAtPixel(output, x, y)) {
			output2[k] = true
		}
	}
	// draw(output2, -5, -5, 10, 8)
	log.Println(len(output2))
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func parse(key string) (int, int) {
	bla := strings.Split(key, ",")
	x, _ := strconv.Atoi(bla[0])
	y, _ := strconv.Atoi(bla[1])
	return x, y
}

func lookupOutput(algorithm string, i int) bool {
	return string(algorithm[i]) == "#"
}

func readInputNumberAtPixel(litPixels map[string]bool, x, y int) int {
	numString := ""
	for yy := y - 1; yy <= y+1; yy++ {
		for xx := x - 1; xx <= x+1; xx++ {
			if _, ok := litPixels[key(xx, yy)]; ok {
				numString += "1"
			} else {
				numString += "0"
			}
		}
	}
	res, _ := strconv.ParseInt(numString, 2, 64)
	return int(res)
}

func draw(pixels map[string]bool, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			r := key(x, y)
			if pixels[r] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
