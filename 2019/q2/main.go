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
	input, err := fileutil.ScanCSInts("input.txt")
	// input, err := fileutil.ScanCSInts("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	input[1] = 12
	input[2] = 2
	result := run(input)
	fmt.Println(result[0])
}

func partTwo() {
	input, err := fileutil.ScanCSInts("input.txt")
	// input, err := fileutil.ScanCSInts("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			mem := make([]int, len(input))
			for i, x := range input {
				mem[i] = x
			}
			mem[1] = noun
			mem[2] = verb
			result := run(mem)
			if result[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}

func run(memory []int) []int {
	ip := 0
	for {
		instruction := memory[ip]
		if instruction == 99 {
			// log.Println(memory)
			return memory
		}
		if instruction == 1 {
			left := memory[memory[ip+1]]
			right := memory[memory[ip+2]]
			dst := memory[ip+3]
			// log.Printf("writing %d+%d to %d", left, right, dst)
			memory[dst] = left + right
			ip += 4
		}
		if instruction == 2 {
			left := memory[memory[ip+1]]
			right := memory[memory[ip+2]]
			dst := memory[ip+3]
			// log.Printf("writing %d*%d to %d", left, right, dst)
			memory[dst] = left * right
			ip += 4
		}
	}
}
