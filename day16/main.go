package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func PathTrace(start, end *helpers.Coord, maze *helpers.Field) []*PathTracer {
	solutions := make([]*PathTracer, 0)
	visited := make(map[helpers.Coord]bool)
	queue := []*PathTracer{NewTracer(start, start)}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for {
			visited[*current.Pos()] = true
			maze.SetLetter(current.Pos(), ' ')

			left, right := CheckSides(current.Iter, maze)
			if left {
				queue = append(queue, current.RotateOther())
			}
			if right {
				queue = append(queue, current.Rotate())
			}

			current.Move()
			if maze.GetLetter(current.Iter.Position) == '#' {
				break
			}
			if maze.GetLetter(current.Iter.Position) == 'E' {
				current.Solution()
				solutions = append(solutions, current)
				break
			}
		}
	}
	return solutions
}

func main() {
	timeStart := time.Now()
	maze := &helpers.Field{}
	maze.LoadData(os.Args[1])

	places := maze.ValuePlaces('.', '#')
	start := places['S'][0]
	end := places['E'][0]

	paths := PathTrace(start, end, maze)

	pathfield := maze.Copy()
	bestSol := math.MaxInt
	for _, path := range paths {
		if path.Solution() < bestSol {
			bestSol = path.Solution()
		}
	}

	for _, path := range paths {
		if path.Solution() != bestSol {
			continue
		}
		for _, coord := range path.Path {
			pathfield.SetLetter(coord, 'O')
		}
	}
	pathfield.PrintData()
	fmt.Println("BEST SOLUTION", bestSol)

	part2 := pathfield.FindLetter(pathfield.MakeAllCoords(), 'O')
	fmt.Println("Part2 : ", len(part2))

	fmt.Println(time.Since(timeStart))
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
