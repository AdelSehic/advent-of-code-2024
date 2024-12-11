package main

import (
	"fmt"
	"os"
	"time"
)

const (
	INPUT      = "input"
	TEST1      = "test1"
	TEST2      = "test2"
)

func (in *Input) Blink(times int) {
	for i := 0; i < times; i++ {
		evolution := make([]*Stone, 0)
		for _, stone := range in.Stones {
			evolution = append(evolution, stone.Evolve()...)
		}
		in.Stones = evolution
		fmt.Println(i + 1)
	}
}

func main() {
	start := time.Now()
	input := MakeInput(os.Args[1])
	input.Blink(25)
	fmt.Println("Part1: ", input.StoneCount(), time.Since(start))
}
