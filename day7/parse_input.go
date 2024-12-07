package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Lines []*Line
}

type Line struct {
	Target int
	Values []int
}

func LoadInput(infile string) *Input {
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	input := &Input{
		Lines: make([]*Line, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input.Lines = append(input.Lines, MakeLine(scanner.Text()))
	}

	return input
}

func MakeLine(scanned string) *Line {
	in := &Line{
		Values: make([]int, 0),
	}

	line := strings.Split(scanned, " ")
	line[0] = strings.TrimRight(line[0], ":")

	num, err := strconv.Atoi(line[0])
	if err != nil {
		panic(err)
	}
	in.Target = num
	for i := 1; i < len(line); i++ {
		num, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}
		in.Values = append(in.Values, num)
	}

	return in
}

func (in *Input) Print() {
	for _, ln := range in.Lines {
		fmt.Println(ln)
	}
}
