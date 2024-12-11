package main

import (
	"math"
)

type Stone struct {
	Value uint64
}

func (s *Stone) Evolve() []*Stone {
	rval := []*Stone{s}
	if s.Value == 0 {
		s.Value = 1
		return rval
	}
	valLen := int(math.Log10(float64(s.Value))) + 1
	if valLen%2 == 0 {
		divider := int(math.Pow(10, float64(valLen/2)))
		leftDigits := s.Value / uint64(divider)
		rightDigits := s.Value % uint64(divider)
		s.Value = leftDigits
		rval = append(rval, &Stone{rightDigits})
		return rval
	}
	s.Value *= 2024
	return rval
}

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

func (in *Input) StoneCount() int {
	return len(in.Stones)
}
