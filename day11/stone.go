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
