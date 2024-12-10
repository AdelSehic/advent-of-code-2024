package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func Part1(crd *helpers.Coord, input *helpers.Field) int {

	height := input.GetLetter(crd)
	if height == '9' {
		input.SetLetter(crd, '.')
		return 1
	}

	rval := 0
	up, down, left, right := input.GetLetter(crd.Up()), input.GetLetter(crd.Down()), input.GetLetter(crd.Left()), input.GetLetter(crd.Right())
	if up == height+1 && up != '.' {
		rval += Part1(crd.Up(), input)
	}
	if down == height+1 && down != '.' {
		rval += Part1(crd.Down(), input)
	}
	if left == height+1 && left != '.' {
		rval += Part1(crd.Left(), input)
	}
	if right == height+1 && right != '.' {
		rval += Part1(crd.Right(), input)
	}

	return rval
}

func Part2(crd *helpers.Coord, input *helpers.Field) int {

	height := input.GetLetter(crd)
	if height == '9' {
		return 1
	}

	rval := 0
	up, down, left, right := input.GetLetter(crd.Up()), input.GetLetter(crd.Down()), input.GetLetter(crd.Left()), input.GetLetter(crd.Right())
	if up == height+1 && up != '.' {
		rval += Part2(crd.Up(), input)
	}
	if down == height+1 && down != '.' {
		rval += Part2(crd.Down(), input)
	}
	if left == height+1 && left != '.' {
		rval += Part2(crd.Left(), input)
	}
	if right == height+1 && right != '.' {
		rval += Part2(crd.Right(), input)
	}

	return rval
}

func main() {
	start := time.Now()
	input := &helpers.Field{}
	input.LoadDataWithPadding(os.Args[1], ".")

	part1, part2 := 0, 0
	starts := input.FindLetter(input.MakeAllCoords(), '0')
	for _, v := range starts {
		part1 += Part1(v, input.Copy())
		part2 += Part2(v, input.Copy())
	}
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
	fmt.Println(time.Since(start))
}
