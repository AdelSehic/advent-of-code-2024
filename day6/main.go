package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
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
	start := time.Now()
	var field helpers.Field

	field.LoadDataWithPadding(os.Args[1], "|")
	guard := NewGuard(field.FindLetter(field.MakeAllCoords(), '^')[0])

	positions := make(map[helpers.Coord]bool)
	positions[*guard.Position] = true
	path := make([]*Guard, 0)

	newField := field.Copy()
	for guard.InFront(newField) != '|' {
		if guard.InFront(newField) == '#' {
			guard.Rotate()
			continue
		}
		guard.Move()
		path = append(path, guard.Copy())
		positions[*guard.Position] = true
	}

	loops := 0
	alterations := make(map[string]bool, 0)
	var wg sync.WaitGroup
	for i := 1; i < len(path); i++ {
		pos := path[i].Position
		if alterations[fmt.Sprintf("%05d%05d", pos.Y, pos.X)] {
			continue
		}
		alterations[fmt.Sprintf("%05d%05d", pos.Y, pos.X)] = true

		wg.Add(1)
		go func() {
			g := path[i-1]
			bumped := make(map[string]bool)
			newField := field.Copy()
			newField.SetLetter(pos, '#')

			for g.InFront(newField) != '|' {
				if g.InFront(newField) == '#' {
					if CheckForLoop(g, bumped) {
						loops++
						break
					}
					g.Rotate()
					continue
				}
				g.Move()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Part 1: Guard has visited %d different positions\r\n", len(positions))
	fmt.Printf("Part 2: Detected %d possible loops\r\n", loops)
	fmt.Println(time.Since(start))
}
