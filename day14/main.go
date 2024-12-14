package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	// "time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

const (
	HEIGHT       = 103
	WIDTH        = 101
	FIELD_LETTER = ' '
	ROBOT_LETTER = '#'
)

func RobotMoveFunc(x, y int) func(*helpers.Coord) *helpers.Coord {
	return func(crd *helpers.Coord) *helpers.Coord {
		return &helpers.Coord{
			X: (crd.X + x + WIDTH) % WIDTH,
			Y: (crd.Y + y + HEIGHT) % HEIGHT,
		}
	}
}

func MakeRobots(infile string) []*helpers.FieldIterator {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	robots := make([]*helpers.FieldIterator, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var posX, posY, deltaX, deltaY int
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &posX, &posY, &deltaX, &deltaY)
		fmt.Println(posX, posY, deltaX, deltaY)
		robot := helpers.NewFieldIterator(&helpers.Coord{X: posX, Y: posY})
		robot.MoveFunc = RobotMoveFunc(deltaX, deltaY)
		robots = append(robots, robot)
	}
	return robots
}

func main() {
	robots := MakeRobots(os.Args[1])
	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			robot.Move()
		}
	}

	positions := make(map[helpers.Coord]int)
	coordinateCount := make(map[int]int)
	for _, robot := range robots {
		positions[*robot.Position]++
		coordinateCount[DetermineQuadrant(robot)]++
	}
	delete(coordinateCount, 0)

	fmt.Println("Part1: ", coordinateCount[1]*coordinateCount[2]*coordinateCount[3]*coordinateCount[4])

	for i := 0; true; i++ {
		emptyField := helpers.GenerateEmptyField(HEIGHT, WIDTH, FIELD_LETTER)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		pos := make(map[helpers.Coord]bool)
		for _, robot := range robots {
			robot.Move()
			pos[*robot.Position] = true
			emptyField.SetLetterUnpadded(robot.Position, ROBOT_LETTER)
		}
		if len(pos) == len(robots) {
			emptyField.PrintData()
			fmt.Println("Part2 solution : ", i+101)
			break
		}
		// time.Sleep(250 * time.Millisecond)
	}
}

func DetermineQuadrant(it *helpers.FieldIterator) int {
	midX := WIDTH / 2
	midY := HEIGHT / 2

	// Check if the point is on the axis (midline)
	if it.Position.X == midX || it.Position.Y == midY {
		return 0 // Indicate it's on the midline and doesn't belong to a quadrant
	}
	if it.Position.X < midX && it.Position.Y < midY {
		return 1 // Top-left
	} else if it.Position.X > midX && it.Position.Y < midY {
		return 2 // Top-right
	} else if it.Position.X < midX && it.Position.Y > midY {
		return 3 // Bottom-left
	} else {
		return 4 // Bottom-right
	}
}

func GetQuadrants(field *helpers.Field) (*helpers.Field, *helpers.Field, *helpers.Field, *helpers.Field) {
	left, right := field.SplitVertically()

	topl, bottoml := left.SplitHorizontally()
	topr, bottomr := right.SplitHorizontally()

	topl.Contract(0, 1, 0, 1)
	topr.Contract(0, 1, 0, 0)
	bottoml.Contract(0, 0, 0, 1)

	return topl, topr, bottoml, bottomr
}
