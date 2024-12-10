package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Exp struct {
	Number int
	Repeat int
}

func (e *Exp) ToString() string {
	str := strconv.Itoa(e.Number)
	if e.Number == -1 {
		str = "."
	}
	return strings.Repeat(str, e.Repeat)
}

func (line *Line) ToString() string {
	var sb strings.Builder
	for _, v := range line.Expansions {
		sb.WriteString(v.ToString())
	}
	return sb.String()
}

type Line struct {
	Data       string
	Expanded   string
	Expansions []Exp
	ExpSlice   []int
	DotCount   int
	MaxDigit   int
}

// Copy creates and returns a deep copy of the Line struct
func (line *Line) Copy() *Line {
	// Create a new Line instance
	copyLine := &Line{
		Data:     line.Data,
		Expanded: line.Expanded,
		DotCount: line.DotCount,
		MaxDigit: line.MaxDigit,
	}

	// Deep copy Expansions slice
	if line.Expansions != nil {
		copyLine.Expansions = make([]Exp, len(line.Expansions))
		copy(copyLine.Expansions, line.Expansions)
	}

	// Deep copy ExpSlice
	if line.ExpSlice != nil {
		copyLine.ExpSlice = make([]int, len(line.ExpSlice))
		copy(copyLine.ExpSlice, line.ExpSlice)
	}

	return copyLine
}

func (line *Line) Print() {
	fmt.Println(line.Data)
}

func (line *Line) Expand() {
	var sb strings.Builder
	digit := 0
	dots := false
	for i := 0; i < len(line.Data); i++ {
		num, _ := strconv.Atoi(string(line.Data[i]))
		if dots {
			sb.WriteString(strings.Repeat(".", num))
			line.ExpSlice = append(line.ExpSlice, repeatElement(-1, num)...)
			line.DotCount += num
			line.Expansions = append(line.Expansions, Exp{-1, num})
			dots = false
		} else {
			sb.WriteString(strings.Repeat(strconv.Itoa(digit), num))
			line.Expansions = append(line.Expansions, Exp{digit, num})
			line.ExpSlice = append(line.ExpSlice, repeatElement(digit, num)...)
			digit++
			dots = true
		}
	}
	line.MaxDigit = digit
	line.Expanded = sb.String()
}

func (line *Line) DefragmentBreak() {
	left := 0
	right := len(line.ExpSlice) - 1

	for left < right {
		if line.ExpSlice[left] != -1 {
			left++
		}
		if line.ExpSlice[right] == -1 {
			right--
		}
		for line.ExpSlice[right] != -1 && line.ExpSlice[left] == -1 {
			tmp := line.ExpSlice[left]
			line.ExpSlice[left] = line.ExpSlice[right]
			line.ExpSlice[right] = tmp
			left++
			right--
		}
	}
}

func (line *Line) DefragmentWhole() {
	for i := len(line.Expansions) - 1; i >= 0; i-- {
		if line.Expansions[i].Number == -1 {
			continue
		}
		place := line.FindPlace(line.Expansions[i].Repeat, i)
		if place == -1 {
			// fmt.Println("No suitable place to insert ", line.Expansions[i].ToString(), " found")
			continue
		}
		// fmt.Println(line.ToString())

		line.Expansions[place].Repeat -= line.Expansions[i].Repeat
		if line.Expansions[place].Repeat <= 0 {
			line.Expansions[place] = line.Expansions[i]
		} else {
			line.Expansions = slices.Insert(line.Expansions, place, line.Expansions[i])
			i++
		}
		line.Expansions[i].Number = -1

		// fmt.Println(line.ToString())
	}
}

func (line *Line) ApplyExpansions() {
	nums := make([]int, 0)
	for _, v := range line.Expansions {
		nums = append(nums, repeatElement(v.Number, v.Repeat)...)
	}
	line.ExpSlice = nums
}

func (line *Line) FindPlace(length, limit int) int {
	for i := 0; i < len(line.Expansions); i++ {
		if i >= limit {
			return -1
		}
		exp := line.Expansions[i]
		// fmt.Println(exp)
		if exp.Number == -1 && exp.Repeat >= length {
			// line.Expansions[i].Repeat -= length
			return i
		}
	}
	return -1
}

func (line *Line) Decode() string {
	var sb strings.Builder
	for _, v := range line.ExpSlice {
		if v == -1 {
			sb.WriteString(".")
		} else {
			sb.WriteString(strconv.Itoa(v))
		}
	}
	return sb.String()
}

func (line *Line) Checksum() uint64 {
	sum := uint64(0)
	for i, v := range line.ExpSlice {
		if v == -1 {
			continue
		}
		sum += uint64(i) * uint64(v)
	}
	return sum
}
