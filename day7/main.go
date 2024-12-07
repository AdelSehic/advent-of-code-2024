package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func (in *Input) ValidPermCount() uint64 {
	sum := uint64(0)
	for _, ln := range in.Lines {
		if ln.ValidPermutation() {
			sum += uint64(ln.Target)
		}
	}
	return sum
}

// check if a line contains a single valid permutation
func (line *Line) ValidPermutation() bool {
	// fmt.Printf("Checking line %+v ...\r\n", *line)
	permutations := helpers.ByteSequence(len(line.Values) - 1)
	for _, opcode := range permutations {
		if line.ApplyOperators(opcode) == line.Target {
			return true
		}
	}
	return false
}

func (line *Line) ApplyOperators(opcode []int) int {
	result := line.Values[0]

	for i := 0; i < len(line.Values)-1; i++ {
		switch opcode[i] {
		case 0:
			result += line.Values[i+1]
		case 1:
			result *= line.Values[i+1]
		}
	}
	// fmt.Printf("Opcode %+v result : %d\r\n", opcode, result)

	return result
}

func main() {
	input := LoadInput(os.Args[1])

	// input.Print()

	fmt.Println("Part1: ", input.ValidPermCount())
}
