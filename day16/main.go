package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func PathSearch(start, end *helpers.Coord, maze *helpers.Field) {
	visited := make(map[helpers.Coord]bool)
	queue := []*helpers.FieldIterator{helpers.NewFieldIterator(start)}
	queue[0].OriginalPos = &helpers.Coord{X: 0, Y: 1}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for {
			if visited[*current.Position] {
				break
			}
			visited[*current.Position] = true
			maze.SetLetter(current.Position, current.GetDirectionLetter())

			left, right := CheckSides(current, maze)
			if left {
				newit := current.NewCopy()
				newit.RotateOther()
				newit.Move()
				newit.OriginalPos.X++
				newit.OriginalPos.Y++
				queue = append(queue, newit)
			}
			if right {
				newit := current.NewCopy()
				newit.Rotate()
				newit.Move()
				newit.OriginalPos.X++
				newit.OriginalPos.Y++
				queue = append(queue, newit)
			}

			current.Move()
			current.OriginalPos.X++
			if maze.GetLetter(current.Position) == '#' {
				break
			}
			if maze.GetLetter(current.Position) == 'E' {
				sol := current.OriginalPos.X + (current.OriginalPos.Y*1000)
				fmt.Printf("ENDING\r\n-------------------------\r\nSteps: %d, Turns: %d, Solution: %d\r\n", current.OriginalPos.X, current.OriginalPos.Y, sol)
				return
			}
		}
	}
}

func main() {
	timeStart := time.Now()
	maze := &helpers.Field{}
	maze.LoadData(os.Args[1])

	places := maze.ValuePlaces('.', '#')
	start := places['S'][0]
	end := places['E'][0]

	PathSearch(start, end, maze)
	time.Since(timeStart)
}

func CheckSides(it *helpers.FieldIterator, maze *helpers.Field) (bool, bool) {
	if maze.GetLetter(it.Position) == '#' {
		return false, false
	}
	newit := it.NewCopy()
	newit.Rotate()
	right := newit.InFront(maze) == '.'

	newit.Rotate()
	newit.Rotate()
	left := newit.InFront(maze) == '.'
	return left, right
}
