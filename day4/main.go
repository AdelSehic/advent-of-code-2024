package main

import (
	"fmt"
	"os"
)

func main() {
	input := &Input{}
	input.LoadDataWithPadding(os.Args[1])
	fmt.Println("Part1: ", input.xmasCount())
	fmt.Println("Part2: ", input.XmasCount())
}
