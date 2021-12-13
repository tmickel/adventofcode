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
	input, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0

	grid := make([][]int, 0)
	for _, line := range input {
		levels := strings.Split(line, "")
		row := make([]int, 0)
		for _, level := range levels {
			z, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, z)
		}
		grid = append(grid, row)
	}

	height := len(grid)
	width := len(grid[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			adj := make([]int, 0)
			if y > 0 {
				adj = append(adj, grid[y-1][x])
			}
			if y < height-1 {
				adj = append(adj, grid[y+1][x])
			}
			if x > 0 {
				adj = append(adj, grid[y][x-1])
			}
			if x < width-1 {
				adj = append(adj, grid[y][x+1])
			}
			isLowest := true
			for _, a := range adj {
				if a <= grid[y][x] {
					isLowest = false
				}
			}
			if isLowest {
				result += 1 + grid[y][x]
			}
		}
	}

	fmt.Println(result)
}

func partTwo() {
	input, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := make([][]int, 0)

	for _, line := range input {
		levels := strings.Split(line, "")
		row := make([]int, 0)
		for _, level := range levels {
			z, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, z)
		}
		grid = append(grid, row)
	}

	height := len(grid)
	width := len(grid[0])

	visited := make(map[string]bool, 0) // "x,y" => bool

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 9 {
				visited[key(x, y)] = true
				continue
			}
			visited[key(x, y)] = false
		}
	}

	basins := make([]int, 0)

	for {
		nextBasinStart := "none"

		// check all positions of the grid for unvisited to start new basins
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !visited[key(x, y)] {
					nextBasinStart = key(x, y)
					break
				}
			}
		}

		if nextBasinStart == "none" {
			break
		}

		// do something bfs-like to grow the basin
		visited[nextBasinStart] = true
		queue := make([]string, 0)
		queue = append(queue, nextBasinStart)
		basinSize := 1
		for len(queue) > 0 {
			pt := queue[0]
			queue = queue[1:]

			ptParse := strings.Split(pt, ",")
			x, _ := strconv.Atoi(ptParse[0])
			y, _ := strconv.Atoi(ptParse[1])

			if y > 0 && grid[y-1][x] < 9 {
				p := key(x, y-1)
				if !visited[p] {
					queue = append(queue, p)
					visited[p] = true
					basinSize++
				}
			}
			if y < height-1 && grid[y+1][x] < 9 {
				p := key(x, y+1)
				if !visited[p] {
					queue = append(queue, p)
					visited[p] = true
					basinSize++
				}
			}
			if x > 0 && grid[y][x-1] < 9 {
				p := key(x-1, y)
				if !visited[p] {
					queue = append(queue, p)
					visited[p] = true
					basinSize++
				}
			}
			if x < width-1 && grid[y][x+1] < 9 {
				p := key(x+1, y)
				if !visited[p] {
					queue = append(queue, p)
					visited[p] = true
					basinSize++
				}
			}
		}
		basins = append(basins, basinSize)
	}

	sort.Ints(basins)
	fmt.Println(
		basins[len(basins)-1] *
			basins[len(basins)-2] *
			basins[len(basins)-3],
	)
}

func key(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
