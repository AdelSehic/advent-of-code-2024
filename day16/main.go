package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

var directNeighbors = []func(*helpers.Coord) *helpers.Coord{
	(*helpers.Coord).Up,
	(*helpers.Coord).Right,
	(*helpers.Coord).Left,
	(*helpers.Coord).Down,
}

func DepthFirstFill(startpos, endpos *helpers.Coord, maze *helpers.Field) (int, int) {
	visited := make(map[helpers.Coord]bool)
	stack := []*helpers.Coord{startpos}
	steps := 2
	turns := 0
	prevDir := &helpers.Coord{X: 0, Y: 0}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		visited[*current] = true

		if *current == *endpos {
			return steps, turns
		}

		end, moveFunc := findDirectionFunc(current, maze, visited)
		if end {
			stack = stack[:len(stack)-1]
			continue
		}

		next := moveFunc(current)

		direction := &helpers.Coord{
			X: current.X - next.X,
			Y: current.Y - next.Y,
		}

		if *prevDir != *direction && !oppositeDirections(prevDir, direction) {
			turns++
		}
		steps++
		prevDir = direction

		maze.SetLetter(next, '?')
		drawMaze(maze)
		time.Sleep(100 * time.Millisecond)

		stack = append(stack, next)
	}
	return steps, turns
}

func oppositeDirections(crd1, crd2 *helpers.Coord) bool {
	// fmt.Println("Opposite ", crd1, crd2)
	isit := (crd1.X+crd2.X) == 0 && (crd1.Y+crd2.Y) == 0
	// fmt.Println(isit)
	return isit
}

// func moveToNextCrossroad(position *helpers.Coord, maze *helpers.Field, visited map[helpers.Coord]bool) *helpers.Coord {
// 	for !detectCrossroad(position, maze, visited) {
// 		visited[*position] = true
// 		end, moveFunc := findDirectionFunc(position, maze, visited)
// 		if end {
// 			return nil
// 		}
// 		position = moveFunc(position)
// 		maze.SetLetter(position, '?')
// 		maze.PrintData()
// 		time.Sleep(500 * time.Millisecond)
// 	}
// 	return position
// }

// func detectCrossroad(position *helpers.Coord, maze *helpers.Field, visited map[helpers.Coord]bool) bool {
// 	pathCount := 0
// 	for _, fn := range directNeighbors {
// 		coord := fn(position)
// 		if visited[*coord] {
// 			continue
// 		}
// 		if maze.GetLetter(coord) == '.' {
// 			pathCount++
// 		}
// 	}
// 	return pathCount > 1
// }

func findDirectionFunc(position *helpers.Coord, maze *helpers.Field, visited map[helpers.Coord]bool) (bool, func(*helpers.Coord) *helpers.Coord) {
	for _, fn := range directNeighbors {
		coord := fn(position)
		letter := maze.GetLetter(coord)
		if !visited[*coord] && (letter == '.' || letter == 'E') {
			return false, fn
		}
	}
	return true, nil
}

func main() {

	if err := termbox.Init(); err != nil {
		panic(err)
	}

	maze := &helpers.Field{}
	maze.LoadData(os.Args[1])

	maze.PrintData()
	start := maze.FindLetter(maze.MakeAllCoords(), 'S')[0]
	end := maze.FindLetter(maze.MakeAllCoords(), 'E')[0]

	steps, turn := DepthFirstFill(start, end, maze)
	termbox.Close()

	fmt.Printf("Steps: %d, turns: %d, part1 result = %d\r\n", steps, turn, steps+turn*1000)
}

func drawMaze(maze *helpers.Field) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y, line := range maze.Lines {
		for x, char := range line {
			switch char {
			case '#':
				termbox.SetCell(x, y, '█', termbox.ColorDarkGray, termbox.ColorDefault)
			case '.':
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			case 'S', 'E':
				termbox.SetCell(x, y, char, termbox.ColorBlue, termbox.ColorDefault)
			default:
				termbox.SetCell(x, y, char, termbox.ColorGreen, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}
