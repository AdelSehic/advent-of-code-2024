package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func Move(startPos *helpers.Coord, field *helpers.Field, direction byte) *helpers.Coord {
	var directionFunc func(c *helpers.Coord) *helpers.Coord
	switch direction {
	case helpers.LETTER_UP:
		directionFunc = (*helpers.Coord).Up
	case helpers.LETTER_DOWN:
		directionFunc = (*helpers.Coord).Down
	case helpers.LETTER_LEFT:
		directionFunc = (*helpers.Coord).Left
	case helpers.LETTER_RIGHT:
		directionFunc = (*helpers.Coord).Right
	}
	if moveHelper(startPos, field, directionFunc) {
		return directionFunc(startPos)
	}
	return startPos
}

func moveHelper(position *helpers.Coord, field *helpers.Field, moveFunc func(*helpers.Coord) *helpers.Coord) bool {
	next := moveFunc(position)
	letter := field.GetLetter(next)
	if letter == '#' {
		return false
	}
	if letter == '.' {
		return true
	}
	if moveHelper(next, field, moveFunc) {
		next = moveFunc(next)
		field.SetLetter(next, letter)
		return true
	}
	return false
}

func main() {
	input := &helpers.Field{}
	input.LoadDataWithPadding(os.Args[1], "#")
	input.PrintData()

	robotPlace := input.FindLetter(input.MakeAllCoords(), '@')
	robot := helpers.NewFieldIterator(robotPlace[0])

	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(scanner.Text()) != 0 {
	}

	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	directions := sb.String()

	fmt.Println(directions)
	for _, dir := range directions {
		input.SetLetter(robot.Position, '.')
		robot.Position = Move(robot.Position, input, byte(dir))
		input.SetLetter(robot.Position, byte(dir))
	}
	fmt.Println()
	input.PrintData()

	part1 := 0
	boxPlaces := input.FindLetter(input.MakeAllCoords(), 'O')
	for _, v := range boxPlaces {
		part1 += (v.Y-1) * 100 + v.X-1
	}
	fmt.Println("Part1: ", part1)
}
