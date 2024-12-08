package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func FrequencyPlaces(coordinates []*helpers.Coord) []*helpers.FieldIterator {
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

	part1Output := input.Copy()
	part1 := make(map[helpers.Coord]bool)
	for _, coordinates := range input.ValuePlaces('.') {
		frequencies := FrequencyPlaces(coordinates)
		for _, iter := range frequencies {
			iter.Move()
			if input.WithinBounds(iter.Position) {
				part1[*iter.Position] = true
				part1Output.SetLetter(iter.Position, '#')
			}
		}
	}
	fmt.Println("Part1 field:")
	part1Output.PrintData()

	part2 := make(map[helpers.Coord]bool)
	for _, coordinates := range input.ValuePlaces('.') {
		frequencies := FrequencyPlaces(coordinates)
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
	fmt.Println()
	fmt.Println("Part2 field:")
	input.PrintData()

	fmt.Println()
	fmt.Println("Part1 solution : ", len(part1))
	fmt.Println("Part2 solution : " , len(part2))
}
