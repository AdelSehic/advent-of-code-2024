package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
	"github.com/nsf/termbox-go"
)

const (
	HEIGHT       = 7
	WIDHT        = 11
	FIELD_LETTER = '.'
	ROBOT_LETTER = '^'
)

func RobotMoveFunc(x, y int) func(*helpers.Coord) *helpers.Coord {
	return func(crd *helpers.Coord) *helpers.Coord {
		return &helpers.Coord{
			X: (crd.X + x + WIDHT) % WIDHT,
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

	// err := termbox.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termbox.Close()

	field := helpers.GenerateEmptyField(HEIGHT, WIDHT, FIELD_LETTER)

	robot := helpers.NewFieldIterator(&helpers.Coord{X: 5, Y: 4})
	robot.MoveFunc = RobotMoveFunc(1, -1)

	robots := MakeRobots(os.Args[1])
	for i := 0; i < 100; i++ {
		for _, robot := range robots {
			// field.SetLetterUnpadded(robot.Position, ROBOT_LETTER)
			robot.Move()
		}
	}

	positions := make(map[helpers.Coord]int)
	for _, robot := range robots {
		positions[*robot.Position]++
		field.SetLetterUnpadded(robot.Position, byte(positions[*robot.Position]+48))
	}

	topl, topr, bottoml, bottomr := GetQuadrants(field)

	// for i := 0; i < 100; i++ {
	// termbox.Clear(termbox.ColorBlue, termbox.ColorDefault)
	// drawGrid()
	// drawRobot(robot.Position)
	// termbox.Flush()
	// time.Sleep(200 * time.Millisecond)
	// robot.Move()
	// 	field.SetLetterUnpadded(robot.Position, ROBOT_LETTER)
	// 	field.PrintData()
	// 	time.Sleep(time.Millisecond * 500)
	// 	field.SetLetterUnpadded(robot.Position, FIELD_LETTER)
	// 	robot.Move()
	// 	fmt.Println(robot.Position)
	// }

	// field.PrintData()
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

func drawRobot(coord *helpers.Coord) {
	termbox.SetCell(coord.X, coord.Y, '@', termbox.ColorGreen, termbox.ColorDefault)
}

func drawGrid() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDHT; x++ {
			termbox.SetCell(x, y, '.', termbox.ColorWhite, termbox.ColorDefault)
		}
	}
}
