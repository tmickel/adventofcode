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
	// input, err := fileutil.ScanStrings("example.txt")
	input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}

	grid := make([][]int, 0)
	for _, row := range input {
		energies := make([]int, 0)
		oct := strings.Split(row, "")
		for _, o := range oct {
			energy, _ := strconv.Atoi(o)
			energies = append(energies, energy)
		}
		grid = append(grid, energies)
	}

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += step(grid)
	}

	fmt.Println(flashes)
}

func step(grid [][]int) int {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x]++
		}
	}

	flashedInRound := make(map[string]bool, 0)

	flashes := 0
	for {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				height := len(grid)
				width := len(grid[y])
				e := grid[y][x]
				if e > 9 && !flashedInRound[fmt.Sprintf("%d,%d", x, y)] {
					flashes++
					flashedInRound[fmt.Sprintf("%d,%d", x, y)] = true
					if y > 0 && !flashedInRound[fmt.Sprintf("%d,%d", x, y-1)] {
						grid[y-1][x]++
					}
					if y < height-1 && !flashedInRound[fmt.Sprintf("%d,%d", x, y+1)] {
						grid[y+1][x]++
					}
					if x < width-1 && !flashedInRound[fmt.Sprintf("%d,%d", x+1, y)] {
						grid[y][x+1]++
					}
					if x > 0 && !flashedInRound[fmt.Sprintf("%d,%d", x-1, y)] {
						grid[y][x-1]++
					}
					if x > 0 && y > 0 && !flashedInRound[fmt.Sprintf("%d,%d", x-1, y-1)] {
						grid[y-1][x-1]++
					}
					if y < height-1 && x < width-1 && !flashedInRound[fmt.Sprintf("%d,%d", x+1, y+1)] {
						grid[y+1][x+1]++
					}

					if y < height-1 && x > 0 && !flashedInRound[fmt.Sprintf("%d,%d", x-1, y+1)] {
						grid[y+1][x-1]++
					}

					if y > 0 && x < width-1 && !flashedInRound[fmt.Sprintf("%d,%d", x+1, y-1)] {
						grid[y-1][x+1]++
					}

					grid[y][x] = 0
				}
			}
		}
		over := true
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] > 9 && !flashedInRound[fmt.Sprintf("%d,%d", x, y)] {
					over = false
				}
			}
		}
		if over {
			break
		}
	}

	return flashes
}

func partTwo() {
	// input, err := fileutil.ScanStrings("example.txt")
	input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}

	grid := make([][]int, 0)
	for _, row := range input {
		energies := make([]int, 0)
		oct := strings.Split(row, "")
		for _, o := range oct {
			energy, _ := strconv.Atoi(o)
			energies = append(energies, energy)
		}
		grid = append(grid, energies)
	}

	i := 0
	for {
		i++
		step(grid)
		allF := true
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] != 0 {
					allF = false
				}
			}
		}
		if allF {
			break
		}
	}

	fmt.Println(i)
}
