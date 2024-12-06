package main

import (
	"fmt"
	"github.com/AdelSehic/advent-of-code-2024/helpers"
	"os"
)

func main() {
	var field helpers.Field

	field.LoadDataWithPadding(os.Args[1], "|")
	field.PrintData()

	guardPos := field.FindLetter(field.MakeAllCoords(), '^')[0]
	guard := &Guard{
		MoveFunc: (*helpers.Coord).Up,
		Positing: guardPos,
		Facing:   FACING_UP,
	}

	positions := make(map[helpers.Coord]bool)
	positions[*guard.Positing] = true

	for guard.InFront(&field) != '|' {
		if guard.InFront(&field) == '#' {
			guard.Rotate()
		}
		field.SetLetter(guard.Positing, 'X')
		guard.Move()
		positions[*guard.Positing] = true
	}

	field.SetLetter(guard.Positing, 'X')
	field.PrintData()
	fmt.Printf("Guard has visited %d different positions\r\n", len(positions))
}
