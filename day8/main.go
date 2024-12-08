package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func FrequencyPlaces(coordinates []*helpers.Coord) []*helpers.Coord {
	frequencies := make([]*helpers.Coord, 0)
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			freq1, freq2 := FrequencyPlace(coordinates[i], coordinates[j])
			frequencies = append(frequencies, freq1)
			frequencies = append(frequencies, freq2)
		}
	}
	return frequencies
}

func FrequencyPlace(first, second *helpers.Coord) (*helpers.Coord, *helpers.Coord) {
	x, y := first.Distance(second)
	freq1 := &helpers.Coord{
		Y: first.Y + y,
		X: first.X + x,
	}
	x2, y2 := second.Distance(first)
	freq2 := &helpers.Coord{
		Y: second.Y + y2,
		X: second.X + x2,
	}
	return freq1, freq2
}

func main() {
	input := &helpers.Field{}
	input.LoadDataWithPadding(os.Args[1], "+")

	uniqueFreqs := make(map[helpers.Coord]bool)
	for _, coordinates := range input.ValuePlaces('.') {
		frequencies := FrequencyPlaces(coordinates)
		for _, freq := range frequencies {
			if input.WithinBounds(freq) {
				uniqueFreqs[*freq] = true
				input.SetLetter(freq, '#')
			}
		}
	}

	fmt.Println(len(uniqueFreqs))
	input.PrintData()
}
