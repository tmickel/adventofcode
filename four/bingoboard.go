package main

import (
	"log"
	"strconv"
	"strings"
)

type BingoBoard struct {
	numbers [][]int
	marked  [][]bool
}

func NewBingoBoard(boardSpec string) *BingoBoard {
	numbers := make([][]int, 0)
	marked := make([][]bool, 0)

	rows := strings.Split(strings.TrimSpace(boardSpec), "\n")
	if len(rows) != 5 {
		log.Fatal("wrong row count for board spec")
	}
	for _, row := range rows {
		cols := strings.Fields(row)
		rowResult := make([]int, 0)
		markedResult := make([]bool, 0)
		for _, col := range cols {
			n, err := strconv.Atoi(col)
			if err != nil {
				log.Fatal(err)
			}
			rowResult = append(rowResult, n)
			markedResult = append(markedResult, false)
		}
		numbers = append(numbers, rowResult)
		marked = append(marked, markedResult)
	}

	return &BingoBoard{
		numbers: numbers,
		marked:  marked,
	}
}

func (b *BingoBoard) CallNumber(n int) {
	for i, row := range b.numbers {
		for j, col := range row {
			if col == n {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *BingoBoard) SumUncalled() int {
	total := 0
	for i, row := range b.marked {
		for j, col := range row {
			if !col {
				total += b.numbers[i][j]
			}
		}
	}
	return total
}

func (b *BingoBoard) Winner() bool {
	for _, row := range b.marked {
		rowWon := true
		for _, col := range row {
			rowWon = rowWon && col
		}
		if rowWon {
			return true
		}
	}
	for colI := 0; colI < len(b.marked[0]); colI++ {
		colWon := true
		for _, row := range b.marked {
			colWon = colWon && row[colI]
		}
		if colWon {
			return true
		}
	}

	return false
}

func (b *BingoBoard) DebugPrint() {
	log.Println(b.numbers)
	log.Println(b.marked)
}
