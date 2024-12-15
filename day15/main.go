package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
	"github.com/eiannone/keyboard"
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

	if moveHelper(startPos, field, direction, directionFunc) {
		return directionFunc(startPos)
	}
	return startPos
}

func dfWallSearch(position *helpers.Coord, field *helpers.Field, direction byte, moveFunc func(*helpers.Coord) *helpers.Coord) bool {
	if position == nil {
		return false
	}
	next := moveFunc(position)
	letter := field.GetLetter(next)
	if letter == '#' {
		return true
	}
	if letter == '.' {
		return false
	}
	hash := string(letter) + string(direction)
	var halfBox *helpers.Coord
	halfBox = nil
	switch hash {
	case "[^", "[v":
		halfBox = next.Right()
	case "]^", "]v":
		halfBox = next.Left()
	}
	return dfWallSearch(next, field, direction, moveFunc) || dfWallSearch(halfBox, field, direction, moveFunc)
}

func moveHelper(position *helpers.Coord, field *helpers.Field, direction byte, moveFunc func(*helpers.Coord) *helpers.Coord) bool {
	next := moveFunc(position)
	letter := field.GetLetter(next)
	if letter == '#' {
		return false
	}
	if letter == '.' {
		return true
	}

	if dfWallSearch(position, field, direction, moveFunc) {
		return false
	}

	hash := string(letter) + string(direction)
	var halfBox *helpers.Coord
	halfBox = nil
	otherLetter := byte('.')
	switch hash {
	case "[^", "[v":
		halfBox = next.Right()
		otherLetter = ']'
	case "]^", "]v":
		halfBox = next.Left()
		otherLetter = '['
	}

	if dfWallSearch(position, field, direction, moveFunc) {
		return false
	}

	if halfBox != nil {
		if moveHelper(next, field, direction, moveFunc) && moveHelper(halfBox, field, direction, moveFunc) {
			field.SetLetter(next, '.')
			field.SetLetter(halfBox, '.')
			field.SetLetter(moveFunc(next), letter)
			field.SetLetter(moveFunc(halfBox), otherLetter)
			return true
		}
	} else {
		if moveHelper(next, field, direction, moveFunc) {
			field.SetLetter(moveFunc(next), letter)
			return true
		}
	}
	return false
}

func main() {
	input := &helpers.Field{}
	input.LoadDataWithPadding(os.Args[1], "#")
	part2 := CreatePart2(input)
	input.Contract(1, 1, 1, 1)
	part2.Contract(1, 1, 2, 2)

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

	for _, dir := range directions {
		input.SetLetter(robot.Position, '.')
		robot.Position = Move(robot.Position, input, byte(dir))
		input.SetLetter(robot.Position, byte(dir))
	}

	part1 := 0
	boxPlaces := input.FindLetter(input.MakeAllCoords(), 'O')
	for _, v := range boxPlaces {
		part1 += (v.Y)*100 + v.X
	}
	fmt.Println("Part1: ", part1)

	robot2place := part2.FindLetter(part2.MakeAllCoords(), '@')
	robot2 := helpers.NewFieldIterator(robot2place[0])
	for _, dir := range directions {
		part2.SetLetter(robot2.Position, '.')
		robot2.Position = Move(robot2.Position, part2, byte(dir))
		part2.SetLetter(robot2.Position, byte(dir))
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	part2result := 0
	boxPlacesPart2 := part2.FindLetter(part2.MakeAllCoords(), '[')
	for _, v := range boxPlacesPart2 {
		part2result += (v.Y)*100 + v.X
	}
	fmt.Println("Part2: ", part2result)
}

func CreatePart2(input *helpers.Field) *helpers.Field {
	field := &helpers.Field{
		Width: 2 * input.Width,
		Lines: make([]string, len(input.Lines)),
	}
	field.Width = 2 * input.Width
	i := 0
	for _, line := range input.Lines {
		var sb strings.Builder
		for _, char := range line {
			switch char {
			case '#':
				sb.WriteString("##")
			case '@':
				sb.WriteString("@.")
			case 'O':
				sb.WriteString("[]")
			default:
				sb.WriteString("..")
			}
		}
		field.Lines[i] = sb.String()
		i++
	}
	return field
}
