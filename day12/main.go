package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

var Directions = []string{
	"topleft",
	"up",
	"topright",
	"left",
	"right",
	"bottomleft",
	"down",
	"bottomright",
}

type Plot struct {
	Input  *helpers.Field
	Places []*helpers.Coord
	Fences uint64
	Sides  int
}

func (p *Plot) PrintPlot() {
	for _, k := range p.Places {
		fmt.Printf("{%d, %d}", k.X, k.Y)
	}
	fmt.Printf(" - Fence size: %d\r\n", p.Fences)
}

func (p *Plot) CalculateFenceSize() {
	for _, v := range p.Places {
		p.Fences += CheckBorder(p.Input, v)
	}
}

func (p *Plot) GetPrice() int {
	return len(p.Places) * int(p.Fences)
}

func CheckBorder(input *helpers.Field, place *helpers.Coord) uint64 {
	letter := input.GetLetter(place)
	sum := uint64(0)
	if input.GetLetter(place.Up()) != letter {
		sum++
	}
	if input.GetLetter(place.Down()) != letter {
		sum++
	}
	if input.GetLetter(place.Left()) != letter {
		sum++
	}
	if input.GetLetter(place.Right()) != letter {
		sum++
	}
	return sum
}

func FindPlot(letter byte, input *helpers.Field, place *helpers.Coord) []*helpers.Coord {
	rval := make([]*helpers.Coord, 0)
	if input.GetLetter(place) != letter {
		return rval
	}
	input.SetLetter(place, '*')
	rval = append(rval, place)
	rval = append(rval, FindPlot(letter, input, place.Up())...)
	rval = append(rval, FindPlot(letter, input, place.Down())...)
	rval = append(rval, FindPlot(letter, input, place.Left())...)
	rval = append(rval, FindPlot(letter, input, place.Right())...)
	return rval
}

func main() {
	var input helpers.Field
	input.LoadDataWithPadding(os.Args[1], "*")
	letterPlaces := input.ValuePlaces('*')

	plots := make(map[byte][]*Plot)
	visited := make(map[helpers.Coord]bool)
	for letter, places := range letterPlaces {
		for _, coord := range places {
			if visited[*coord] {
				continue
			}

			p := &Plot{
				Input:  &input,
				Places: FindPlot(letter, input.Copy(), coord),
			}
			p.CalculateFenceSize()
			plots[letter] = append(plots[letter], p)

			for _, c := range p.Places {
				visited[*c] = true
			}
		}
	}

	part1 := uint64(0)
	for _, v := range plots {
		for _, p := range v {
			part1 += uint64(p.GetPrice())
		}
	}
	fmt.Println("Part1 :", part1)

	original := input.Copy()

	for letter, v := range plots {
		for _, p := range v {
			for _, k := range p.Places {
				num := CornerState(letter, k, original)
				input.SetLetter(k, byte(num+48))
				p.Sides += num
			}
			fmt.Println(string(letter), p.Sides)
		}
	}

	input.PrintData()
	original.PrintData()
}

func CornerState(letter byte, coord *helpers.Coord, input *helpers.Field) int {
	neighCount := 0
	neighbors := coord.Neighbors()
	letters := make(map[string]byte)
	for i := 0; i < len(neighbors); i++ {
		letters[Directions[i]] = input.GetLetter(neighbors[i])
	}

	for _, c := range letters {
		if c != letter {
			neighCount++
		}
	}

	switch neighCount {
	case 8:
		return 4
	case 7:
		return 2
	}

	if (letters["up"] == letter && letters["down"] == letter) ||
		(letters["left"] == letter && letters["right"] == letter) {
		return 0
	}


	return 0
}
