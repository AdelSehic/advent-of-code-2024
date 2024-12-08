package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func FrequencyPlacesPart1(coordinates []*helpers.Coord) []*helpers.Coord {
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

func FrequencyPlacesPart2(coordinates []*helpers.Coord) []*helpers.FieldIterator {
	iters := make([]*helpers.FieldIterator, 0)
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			x1, y1 := coordinates[i].Distance(coordinates[j])
			iter1 := helpers.NewFieldIterator(coordinates[i])
			iter1.MoveFunc = helpers.NewMoveFunc(x1, y1)
			iters = append(iters, iter1)
			// fmt.Printf("New iterator at %+v, move func %d %d\r\n", iter1.Position, x1, y1)

			x2, y2 := coordinates[j].Distance(coordinates[i])
			iter2 := helpers.NewFieldIterator(coordinates[j])
			iter2.MoveFunc = helpers.NewMoveFunc(x2, y2)
			iters = append(iters, iter2)
			// fmt.Printf("New iterator at %+v, move func %d %d\r\n", iter2.Position, x2, y2)
		}
	}
	return iters
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
		frequencies := FrequencyPlacesPart1(coordinates)
		for _, freq := range frequencies {
			if input.WithinBounds(freq) {
				uniqueFreqs[*freq] = true
				input.SetLetter(freq, '#')
			}
		}
	}

	part2 := make(map[helpers.Coord]bool)
	for _, coordinates := range input.ValuePlaces('.') {
		frequencies := FrequencyPlacesPart2(coordinates)
		for _, iter := range frequencies {
			for input.WithinBounds(iter.Position) {
				part2[*iter.Position] = true
				if input.GetLetter(iter.Position) == '.' {
					input.SetLetter(iter.Position, '#')
				}
				iter.Move()
			}
		}
	}

	input.PrintData()
	// fmt.Println(len(part1))
	fmt.Println(len(part2))
}
