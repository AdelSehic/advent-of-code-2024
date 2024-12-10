package main

import (
	"fmt"
	"os"
)

func main() {
	in := MakeInput(os.Args[1])
	in.ExpandLines()
	part1 := in.Lines[0].Copy()
	part2 := in.Lines[0].Copy()

	part1.DefragmentBreak()
	fmt.Println("Part1: ", part1.Checksum())

	// fmt.Println(part2.Decode())
	part2.DefragmentWhole()
	part2.ApplyExpansions()
	// fmt.Println(part2.Decode())
	fmt.Println("Part2: ", part2.Checksum())
}
