package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func (in *Input) ValidPermCount(base int) uint64 {
	sum := uint64(0)

	for _, ln := range in.Lines {
		in.WG.Add(1)
		go func() {
			defer in.WG.Done()
			if ln.ValidPermutation(base) {
				sum += uint64(ln.Target)
			}
		}()
	}
	in.WG.Wait()
	return sum
}

// check if a line contains a single valid permutation
func (line *Line) ValidPermutation(base int) bool {
	// fmt.Printf("Checking line %+v ...\r\n", *line)
	permutations := helpers.ByteSequence(len(line.Values)-1, base)
	for _, opcode := range permutations {
		if line.ApplyOperators(opcode) == line.Target {
			// fmt.Printf("Found a valid permutation %+v for %d: %+v\r\n", opcode, line.Target, line.Values)
			return true
		}
	}
	// fmt.Printf("No valid permutations for %+v\r\n", line.Values)
	return false
}

func (line *Line) ApplyOperators(opcode []int) uint64 {
	result := line.Values[0]

	// fmt.Println(opcode)
	for i := 0; i < len(line.Values)-1; i++ {
		if result > line.Target {
			return result
		}
		switch opcode[i] {
		case 0:
			result += line.Values[i+1]
		case 1:
			result *= line.Values[i+1]
		case 2:
			concat := helpers.ConcatInts(int(result), int(line.Values[i+1]))
			result = uint64(concat)
		}
	}
	return result
}

func main() {
	input := LoadInput(os.Args[1])

	start := time.Now()
	fmt.Println("Part1: ", input.ValidPermCount(2), time.Since(start))

	start = time.Now()
	fmt.Println("Part2: ", input.ValidPermCount(3), time.Since(start))
}
