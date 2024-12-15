package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

const (
	HEIGHT       = 103
	WIDTH        = 101
	CELL_SIZE    = 20
	FIELD_LETTER = ' '
	ROBOT_LETTER = '#'
	BACKGROUND_R = 255
	BACKGROUND_G = 255
	BACKGROUND_B = 255
	ROBOT_R      = 0
	ROBOT_G      = 128
	ROBOT_B      = 0
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

	// cmd := exec.Command("clear")
	// cmd.Stdout = os.Stdout
	for i := 0; true; i++ {
		pos := make(map[helpers.Coord]bool)
		// emptyField := helpers.GenerateEmptyField(HEIGHT, WIDTH, FIELD_LETTER)
		for _, robot := range robots {
			robot.Move()
			pos[*robot.Position] = true
			// emptyField.SetLetterUnpadded(robot.Position, ROBOT_LETTER)
		}
		if len(pos) == len(robots) {
			saveGridAsImage(robots, "output.jpg")
			// emptyField.PrintData()
			fmt.Println("Part2: ", i+101)
			break
		}
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

func saveGridAsImage(robots []*helpers.FieldIterator, filename string) error {
	img := image.NewRGBA(image.Rect(0, 0, WIDTH*CELL_SIZE, HEIGHT*CELL_SIZE))
	for y := 0; y < HEIGHT*CELL_SIZE; y++ {
		for x := 0; x < WIDTH*CELL_SIZE; x++ {
			img.Set(x, y, color.RGBA{BACKGROUND_R, BACKGROUND_G, BACKGROUND_B, 255})
		}
	}

	for _, robot := range robots {
		x := robot.Position.X * CELL_SIZE
		y := robot.Position.Y * CELL_SIZE
		for dy := 0; dy < CELL_SIZE; dy++ {
			for dx := 0; dx < CELL_SIZE; dx++ {
				img.Set(x+dx, y+dy, color.RGBA{ROBOT_R, ROBOT_G, ROBOT_B, 255})
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}
