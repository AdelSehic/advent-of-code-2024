package main

import (
	"bufio"
	"os"
)

type Input struct {
	Lines []*Line
}

func MakeInput(infile string) *Input {
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	in := &Input{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in.Lines = append(in.Lines, &Line{
			Data: scanner.Text(),
		})
	}
	return in
}

func (in *Input) PrintLines() {
	for _, line := range in.Lines {
		line.Print()
	}
}

func (in *Input) ExpandLines() {
	for _, line := range in.Lines {
		line.Expand()
	}
}

func (in *Input) DefragLines() {
	for _, line := range in.Lines {
		line.DefragmentBreak()
	}
}
