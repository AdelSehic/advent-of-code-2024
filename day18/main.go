package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

var (
	X_SIZE = 7
	Y_SIZE = 7
)

func main() {
	start := time.Now()

	infile := os.Args[1]
	if infile == "input.txt" {
		X_SIZE = 71
		Y_SIZE = 71
	}

	field := helpers.GenerateEmptyFieldPadded(X_SIZE, Y_SIZE, '.', '#')
	field.SetLetterUnpadded(&helpers.Coord{X: X_SIZE, Y: Y_SIZE}, 'E')
	field.PrintData()

	corruptions := LoadCoords(infile)
	itermax := 12
	if os.Args[1] == "input.txt" {
		itermax = 1024
	}

	for i := 0; i < itermax; i++ {
		field.SetLetter(corruptions[i], '#')
	}

	startPos := &helpers.Coord{X: 1, Y: 1}
	endPos := &helpers.Coord{X: X_SIZE, Y: Y_SIZE}

	solved := field.Copy()
	part1 := PathTraceDijkstra(startPos, endPos, field)
	for _, c := range part1.Path {
		solved.SetLetter(c, 'O')
	}
	solved.PrintData()
	fmt.Println("Part1 : ", part1.Steps)

	for i := itermax; i < len(corruptions); i++ {
		field.SetLetter(corruptions[i], '#')
		if PathTraceDijkstra(startPos, endPos, field) == nil {
			fmt.Printf("Part2 : %d,%d\r\n", corruptions[i].X-1, corruptions[i].Y-1)
			break
		}
	}

	fmt.Println(time.Since(start))
}

func LoadCoords(infile string) []*helpers.Coord {
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rval := make([]*helpers.Coord, 0)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		X, _ := strconv.Atoi(values[0])
		Y, _ := strconv.Atoi(values[1])
		rval = append(rval, &helpers.Coord{X: X + 1, Y: Y + 1})
	}
	return rval
}
