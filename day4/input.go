package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (in *Input) SetDebug(crd *coord, letter rune) {
	row := []rune(in.debug[crd.y]) // Convert the string to a slice of runes
	row[crd.x] = letter            // Modify the specific position
	in.debug[crd.y] = string(row)  // Convert it back to a string and store it
}

func (in *Input) PrintData() {
	for _, v := range in.lines {
		fmt.Println(v)
	}
}

func (in *Input) makeAllCoords() []*coord {
	crds := make([]*coord, 0, len(in.lines)*in.width)
	for i := 1; i < len(in.lines)-1; i++ {
		for j := 1; j < in.width-1; j++ {
			crds = append(crds, &coord{i, j})
		}
	}
	return crds
}

func (in *Input) findMAS(start *coord, next func(*coord) *coord) bool {
	target := []byte{'M', 'A', 'S'}
	current := start

	for _, letter := range target {
		current = next(current)
		if in.getLetter(current) != letter {
			return false
		}
		in.SetDebug(current, rune(letter))
	}

	return true
}

func (in *Input) findLetter(input []*coord, letter byte) []*coord {
	out := make([]*coord, 0)
	for _, v := range input {
		if in.getLetter(v) == letter {
			out = append(out, v)
			in.SetDebug(v, rune(letter))
		}
	}
	return out
}

func (in *Input) xmasCount() int {
	directions := []func(*coord) *coord{
		(*coord).right,
		(*coord).left,
		(*coord).up,
		(*coord).down,
		(*coord).topLeft,
		(*coord).topRight,
		(*coord).bottomLeft,
		(*coord).bottomRight,
	}

	xs := in.findLetter(in.makeAllCoords(), 'X')
	sum := 0
	for _, x := range xs {
		for _, move := range directions {
			if in.findMAS(x, move) {
				sum++
			}
		}
		in.PrintDebug()
	}

	return sum
}

func (in *Input) getLetter(crd *coord) byte {
	if len(in.lines) > crd.y && len(in.lines[crd.y]) > crd.x {
		return in.lines[crd.y][crd.x]
	}
	return '.'
}
