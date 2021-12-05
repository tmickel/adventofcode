package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

type Grid [][]int

func main() {
	partOne()
	partTwo()
}

func partOne() {
	width := 1000
	height := 1000

	grid := make([][]int, 0)
	for y := 0; y < height; y++ {
		row := make([]int, 0)
		for x := 0; x < width; x++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	in, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, lineSpec := range in {
		line, err := parseLine(lineSpec)
		if err != nil {
			log.Fatal(err)
		}
		markLine(line, grid, false)
	}
	fmt.Println(countOverlaps(grid))
}

func partTwo() {
	width := 1000
	height := 1000

	grid := make([][]int, 0)
	for y := 0; y < height; y++ {
		row := make([]int, 0)
		for x := 0; x < width; x++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	in, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, lineSpec := range in {
		line, err := parseLine(lineSpec)
		if err != nil {
			log.Fatal(err)
		}
		markLine(line, grid, true)
	}

	fmt.Println(countOverlaps(grid))
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func parseLine(lineSpec string) (*Line, error) {
	parts := strings.Fields(lineSpec)
	a := parts[0]
	aparts := strings.Split(a, ",")
	x1, err := strconv.Atoi(aparts[0])
	if err != nil {
		return nil, err
	}
	y1, err := strconv.Atoi(aparts[1])
	if err != nil {
		return nil, err
	}
	b := parts[2]
	bparts := strings.Split(b, ",")
	x2, err := strconv.Atoi(bparts[0])
	if err != nil {
		return nil, err
	}
	y2, err := strconv.Atoi(bparts[1])
	if err != nil {
		return nil, err
	}
	return &Line{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}, nil
}

func markLine(line *Line, grid Grid, markDiagonals bool) {
	dx := 0
	if line.x2 > line.x1 {
		dx = 1
	}
	if line.x2 < line.x1 {
		dx = -1
	}
	dy := 0
	if line.y2 > line.y1 {
		dy = 1
	}
	if line.y2 < line.y1 {
		dy = -1
	}

	if dx != 0 && dy != 0 && !markDiagonals {
		return
	}

	x := line.x1
	y := line.y1
	grid[y][x]++
	for x != line.x2 || y != line.y2 {
		x += dx
		y += dy
		grid[y][x]++
	}
}

func debugGrid(grid Grid) {
	for _, row := range grid {
		for _, col := range row {
			if col == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func countOverlaps(grid Grid) int {
	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col > 1 {
				count++
			}
		}
	}
	return count
}
