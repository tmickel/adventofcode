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
	// partTwo()
}

func partOne() {
	p, i := getPaper()
	np := fold(p, i[0])
	fmt.Println(countDots(np))
}

// func partTwo() {
// 	p, i := getPaper()
// 	for _, ins := range i {
// 		p = fold(p, ins)
// 		//	printPaper(p)
// 	}
// 	printPaper(p)
// }

type Paper [][]bool

type Instruction struct {
	alongY bool
	pos    int
}

func getPaper() (Paper, []Instruction) {
	input, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	width := 0
	height := 0
	for _, line := range input {
		if line == "" {
			break
		}
		ps := strings.Split(line, ",")
		x, _ := strconv.Atoi(ps[0])
		y, _ := strconv.Atoi(ps[1])
		if x > width {
			width = x
		}
		if y > height {
			height = y
		}
	}

	pp := make(Paper, 0)
	for y := 0; y <= height; y++ {
		gl := make([]bool, 0)
		for x := 0; x <= width; x++ {
			gl = append(gl, false)
		}
		pp = append(pp, gl)
	}

	ins := make([]Instruction, 0)

	coordinatesMode := true
	for _, line := range input {
		if coordinatesMode {
			if line == "" {
				coordinatesMode = false
				continue
			}
			ps := strings.Split(line, ",")
			x, _ := strconv.Atoi(ps[0])
			y, _ := strconv.Atoi(ps[1])
			pp[y][x] = true
			continue
		}

		i := Instruction{}
		pts := strings.Split(line, "=")
		i.alongY = pts[0] == "fold along y"
		i.pos, _ = strconv.Atoi(pts[1])
		ins = append(ins, i)
	}

	return pp, ins
}

// func printPaper(p Paper) {
// 	for y := 0; y < len(p); y++ {
// 		for x := 0; x < len(p[0]); x++ {
// 			if p[y][x] {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 			if (x+1)%5 == 0 {
// 				fmt.Print("     ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func fold(p Paper, i Instruction) Paper {
	np := make(Paper, 0)
	if i.alongY {
		for y := 0; y < i.pos; y++ {
			gl := make([]bool, 0)
			for x := 0; x < len(p[0]); x++ {
				v := p[y][x]
				if y >= len(p)-2*(len(p)-i.pos) {
					v = v || p[len(p)-y-1][x]
				}
				gl = append(gl, v)
			}
			np = append(np, gl)
		}
	} else {
		for y := 0; y < len(p); y++ {
			gl := make([]bool, 0)
			for x := 0; x < i.pos; x++ {
				v := p[y][x]
				if x > len(p[0])-2*(len(p[0])-i.pos) {
					v = v || p[y][len(p[0])-x-1]
				}
				gl = append(gl, v)
			}
			np = append(np, gl)
		}
	}
	return np
}

func countDots(p Paper) int {
	ct := 0
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[0]); x++ {
			if p[y][x] {
				ct++
			}
		}
	}
	return ct
}
