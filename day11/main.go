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

type Input struct {
	Stones []*Stone
}

func MakeInput(what string) *Input {
	stones := make([]uint64, 0)
	switch what {
	case TEST1:
		stones = []uint64{0, 1, 10, 99, 999}
	case TEST2:
		stones = []uint64{125, 17}
	case INPUT:
		stones = []uint64{8435, 234, 928434, 14, 0, 7, 92446, 8992692}
	}
	input := &Input{
		Stones: make([]*Stone, 0, len(stones)),
	}
	for _, s := range stones {
		input.Stones = append(input.Stones, &Stone{s})
	}
	return input
}

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

func (in *Input) StoneCount() int {
	return len(in.Stones)
}

func main() {
	start := time.Now()
	input := MakeInput(os.Args[1])
	input.Blink(25)
	fmt.Println("Part1: ", input.StoneCount(), time.Since(start))
}
