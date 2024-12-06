package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

type Input struct {
	lines []string
	width int
	debug []string
}

func (in *Input) LoadDataWithPadding(infile string) {
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	in.lines = make([]string, 1)
	for scanner.Scan() {
		in.lines = append(in.lines, strings.Join([]string{".", scanner.Text(), "."}, ""))
	}
	in.width = len(in.lines[1])
	padding := strings.Repeat(".", in.width)
	in.lines[0] = padding
	in.lines = append(in.lines, padding)

	in.debug = make([]string, len(in.lines))
	for i := 0; i < len(in.lines); i++ {
		in.debug[i] = strings.Repeat(".", in.width)
	}
}

func (in *Input) PrintDebug() {
	fmt.Println("----------------------------------------------")
	for _, v := range in.debug {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------------------")
}

func (in *Input) SetDebug(crd *helpers.Coord, letter rune) {
	row := []rune(in.debug[crd.Y]) // Convert the string to a slice of runes
	row[crd.X] = letter            // Modify the specific position
	in.debug[crd.Y] = string(row)  // Convert it back to a string and store it
}

func (in *Input) PrintData() {
	for _, v := range in.lines {
		fmt.Println(v)
	}
}

func (in *Input) makeAllCoords() []*helpers.Coord {
	crds := make([]*helpers.Coord, 0, len(in.lines)*in.width)
	for i := 1; i < len(in.lines)-1; i++ {
		for j := 1; j < in.width-1; j++ {
			crds = append(crds, &helpers.Coord{i, j})
		}
	}
	return crds
}

func (in *Input) findSequence(start *helpers.Coord, next func(*helpers.Coord) *helpers.Coord, sequence []byte) bool {
	current := start

	for _, letter := range sequence {
		current = next(current)
		if in.getLetter(current) != letter {
			return false
		}
	}

	return true
}

func (in *Input) findMAS(start *helpers.Coord, next func(*helpers.Coord) *helpers.Coord) bool {
	target := []byte{'M', 'A', 'S'}
	return in.findSequence(start, next, target)
}

func (in *Input) findAS(start *helpers.Coord, next func(*helpers.Coord) *helpers.Coord) bool {
	target := []byte{'A', 'S'}
	if in.findSequence(start, next, target) {
		in.SetDebug(start, 'M')
		n := next(start)
		in.SetDebug(n, 'A')
		n = next(n)
		in.SetDebug(n, 'S')
		return true
	}
	return false
}

func (in *Input) findLetter(input []*helpers.Coord, letter byte) []*helpers.Coord {
	out := make([]*helpers.Coord, 0)
	for _, v := range input {
		if in.getLetter(v) == letter {
			out = append(out, v)
		}
	}
	return out
}

func (in *Input) xmasCount() int {
	directions := []func(*helpers.Coord) *helpers.Coord{
		(*helpers.Coord).Right,
		(*helpers.Coord).Left,
		(*helpers.Coord).Up,
		(*helpers.Coord).Down,
		(*helpers.Coord).TopLeft,
		(*helpers.Coord).TopRight,
		(*helpers.Coord).BottomLeft,
		(*helpers.Coord).BottomRight,
	}

	xs := in.findLetter(in.makeAllCoords(), 'X')
	sum := 0
	for _, x := range xs {
		for _, move := range directions {
			if in.findMAS(x, move) {
				sum++
			}
		}
		// in.PrintDebug()
	}

	return sum
}

func (in *Input) XmasCount() int {
	directions := []func(*helpers.Coord) *helpers.Coord{
		(*helpers.Coord).TopLeft,
		(*helpers.Coord).TopRight,
		(*helpers.Coord).BottomLeft,
		(*helpers.Coord).BottomRight,
	}
	sum := 0

	crosses := make(map[helpers.Coord]int)
	ms := in.findLetter(in.makeAllCoords(), 'M')
	for _, m := range ms {
		for _, move := range directions {
			if in.findAS(m, move) {
				A := move(m)
				crosses[helpers.Coord{A.Y, A.X}]++
			}
		}
	}

	for k, v := range crosses {
		if v >= 2 {
			sum++
			in.SetDebug(&k, '*')
		}
	}

	// in.PrintDebug()
	return sum
}

func (in *Input) getLetter(crd *helpers.Coord) byte {
	if len(in.lines) > crd.Y && len(in.lines[crd.Y]) > crd.X {
		return in.lines[crd.Y][crd.X]
	}
	return '.'
}
