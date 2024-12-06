package main

import (
	"fmt"
	"github.com/AdelSehic/advent-of-code-2024/helpers"
	"os"
)

func CheckForLoop(g *Guard, bump map[string]bool) bool {
	hash := fmt.Sprintf("%05d%05d%s", g.NextField().Y, g.NextField().X, g.Facing)
	if found := bump[hash]; found {
		return true
	}
	bump[hash] = true
	return false
}

func main() {
	var field helpers.Field

	field.LoadDataWithPadding(os.Args[1], "|")
	guard := NewGuard(field.FindLetter(field.MakeAllCoords(), '^')[0])

	positions := make(map[helpers.Coord]bool)
	positions[*guard.Position] = true
	path := make([]*helpers.Coord, 0)

	newField := field.Copy()
	for guard.InFront(newField) != '|' {
		if guard.InFront(newField) == '#' {
			guard.Rotate()
			continue
		}
		guard.Move()
		path = append(path, guard.Position)
		positions[*guard.Position] = true
	}

	loops := 0
	alterations := make(map[string]bool, 0)
	for i := 1; i < len(path); i++ {
		pos := path[i]
		if alterations[fmt.Sprintf("%05d%05d", pos.Y, pos.X)] {
			continue
		}
		alterations[fmt.Sprintf("%05d%05d", pos.Y, pos.X)] = true

		guard.Reset()
		bumped := make(map[string]bool)
		newField := field.Copy()
		newField.SetLetter(pos, '#')

		for guard.InFront(newField) != '|' {
			if guard.InFront(newField) == '#' {
				if CheckForLoop(guard, bumped) {
					loops++
					break
				}
				guard.Rotate()
				continue
			}
			guard.Move()
		}
	}

	fmt.Printf("Part 1: Guard has visited %d different positions\r\n", len(positions))
	fmt.Printf("Part 2: Detected %d possible loops\r\n", loops)
}
