package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Field struct {
	Lines []string
	Width int
	debug []string
}

func (in *Field) LoadDataWithPadding(infile string, paddingChar string) {
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	in.Lines = make([]string, 1)
	for scanner.Scan() {
		in.Lines = append(in.Lines, strings.Join([]string{paddingChar, scanner.Text(), paddingChar}, ""))
	}
	in.Width = len(in.Lines[1])
	padding := strings.Repeat(paddingChar, in.Width)
	in.Lines[0] = padding
	in.Lines = append(in.Lines, padding)

	in.debug = make([]string, len(in.Lines))
	for i := 0; i < len(in.Lines); i++ {
		in.debug[i] = strings.Repeat(paddingChar, in.Width)
	}
}

func GenerateEmptyField(height, widht int, character byte) *Field {
	f := &Field{}
	for i := 0; i < height; i++ {
		f.Lines = append(f.Lines, strings.Repeat(string(character), widht))
	}
	f.Width = widht
	return f
}

func (in *Field) PrintDebug() {
	fmt.Println("----------------------------------------------")
	for _, v := range in.debug {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------------------")
}

func (in *Field) WithinBounds(crd *Coord) bool {
	return crd.X > 0 && crd.X < in.Width-1 && crd.Y > 0 && crd.Y < len(in.Lines)-1
}

func (in *Field) SetDebug(crd *Coord, letter rune) {
	row := []rune(in.debug[crd.Y]) // Convert the string to a slice of runes
	row[crd.X] = letter            // Modify the specific position
	in.debug[crd.Y] = string(row)  // Convert it back to a string and store it
}

func (in *Field) PrintData() {
	for _, v := range in.Lines {
		fmt.Println(v)
	}
}

func (in *Field) MakeAllCoords() []*Coord {
	crds := make([]*Coord, 0, len(in.Lines)*in.Width)
	for i := 1; i < len(in.Lines)-1; i++ {
		for j := 1; j < in.Width-1; j++ {
			crds = append(crds, &Coord{i, j})
		}
	}
	return crds
}

func (in *Field) FindSequence(start *Coord, next func(*Coord) *Coord, sequence []byte) bool {
	current := start

	for _, letter := range sequence {
		current = next(current)
		if in.GetLetter(current) != letter {
			return false
		}
	}

	return true
}

func (in *Field) FindLetter(input []*Coord, letter byte) []*Coord {
	out := make([]*Coord, 0)
	for _, v := range input {
		if in.GetLetter(v) == letter {
			out = append(out, v)
		}
	}
	return out
}

func (in *Field) GetLetter(crd *Coord) byte {
	if len(in.Lines) > crd.Y && len(in.Lines[crd.Y]) > crd.X {
		return in.Lines[crd.Y][crd.X]
	}
	return '.'
}

func (in *Field) SetLetter(crd *Coord, letter byte) {
	if !in.WithinBounds(crd) {
		return
	}
	line := []byte(in.Lines[crd.Y])
	line[crd.X] = letter
	in.Lines[crd.Y] = string(line)
}

func (in *Field) SetLetterUnpadded(crd *Coord, letter byte) {
	line := []byte(in.Lines[crd.Y])
	line[crd.X] = letter
	in.Lines[crd.Y] = string(line)
}

func (in *Field) SplitVertically() (*Field, *Field) {
	left, right := &Field{}, &Field{}

	half := in.Width / 2
	for _, line := range in.Lines {
		l, r := line[:half+in.Width%2], line[half+in.Width%2:]
		left.Lines = append(left.Lines, l)
		right.Lines = append(right.Lines, r)
	}
	left.Width = half
	right.Width = half

	return left, right
}

func (in *Field) SplitHorizontally() (*Field, *Field) {
	top, bottom := &Field{}, &Field{}

	height := len(in.Lines)
	half := height / 2
	top.Lines, bottom.Lines = in.Lines[:half+height%2], in.Lines[half+height%2:]
	top.Width = in.Width
	bottom.Width = in.Width

	return top, bottom
}

func (f *Field) Copy() *Field {
	newField := &Field{
		Width: f.Width,
		debug: append([]string(nil), f.debug...),
	}

	newField.Lines = make([]string, len(f.Lines))
	copy(newField.Lines, f.Lines)

	return newField
}

func (f *Field) ValuePlaces(exclude ...byte) map[byte][]*Coord {
	excluded := make(map[byte]bool)
	for _, l := range exclude {
		excluded[l] = true
	}

	locations := make(map[byte][]*Coord)
	for _, coord := range f.MakeAllCoords() {
		letter := f.GetLetter(coord)
		if excluded[letter] {
			continue
		}
		locations[letter] = append(locations[letter], coord)
	}
	return locations
}

func (in *Field) Contract(top, bottom, left, right int) {
	in.Lines = in.Lines[top : len(in.Lines)-bottom]

	for i := 0; i < len(in.Lines); i++ {
		line := in.Lines[i]
		in.Lines[i] = line[left : len(in.Lines[i])-right]
	}
}
