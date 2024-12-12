package main

import (
	"fmt"
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

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

	plots := make(map[byte][][]*helpers.Coord)
	for letter, places := range letterPlaces {
		for _, coord := range places {
			if input.GetLetter(coord) == '*' {
				continue
			}
			plots[letter] = append(plots[letter], FindPlot(letter, &input, coord))
		}
	}

	for k, v := range plots {
		fmt.Printf("----------- %s -----------\r\n", string(k))
		for _, i := range v {
			for _, k := range i {
				fmt.Printf("{%d, %d}", k.X, k.Y)
			}
			fmt.Println()
		}
		fmt.Printf("-------------------------\r\n")
	}

	// part1 := uint64(0)
	// for k, v := range fences {
	// 	part1 += v * uint64(len(letterPlaces[k]))
	// 	fmt.Printf("%s - %d\r\n", string(k), v)
	// }
	// fmt.Println("Part1 :", part1)
}
