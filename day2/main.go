package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToIntSlice(input string) []int {
	inSlice := strings.Split(input, " ")
	outSlice := make([]int, 0, len(inSlice))
	for _, v := range inSlice {
		num, _ := strconv.Atoi(v)
		outSlice = append(outSlice, num)
	}
	return outSlice
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sliceIsSortedSet(input *[]int, oneWrong *bool) bool {
	if len(*input) < 2 {
		return true
	}

	compareFunc := func(a, b int) bool {
		return a > b
	}
	if (*input)[0] < (*input)[1] {
		compareFunc = func(a, b int) bool {
			return a < b
		}
	}

	for i := 0; i < len((*input))-1; i++ {
		if !compareFunc((*input)[i], (*input)[i+1]) {
			if *oneWrong {
				return false
			}
			(*input) = append((*input)[:i], (*input)[i+1:]...)
			*oneWrong = true
			return sliceIsSortedSet(input, oneWrong)
		}
	}
	return true
}

func deltasAreSafe(input []int, oneWrong *bool) bool {
	if len(input) < 2 {
		return false
	}

	for i := 1; i < len(input); i++ {
		delta := abs(input[i-1] - input[i])
		if delta > 3 {
			if *oneWrong {
				return false
			}
			*oneWrong = true

			alt1 := make([]int, len(input))
			copy(alt1, input)
			alt1 = append(alt1[:i], alt1[i+1:]...)
			if deltasAreSafe(alt1, oneWrong) {
				return true
			}

			alt2 := make([]int, len(input))
			copy(alt2, input)
			alt2 = append(alt2[:i-1], alt2[i:]...)
			if deltasAreSafe(alt2, oneWrong) {
				return true
			}

			if i < len(input)-1 {
				return false
			}

			alt3 := make([]int, len(input))
			copy(alt3, input)
			alt2 = append(alt2[:i+1], alt2[i+2:]...)
			return deltasAreSafe(alt2, oneWrong)
		}
	}
	return true
}

func isInputSafe(input *[]int) bool {
	oneWrong := false
	sorted := sliceIsSortedSet(input, &oneWrong)
	return sorted && deltasAreSafe(*input, &oneWrong)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	safeNum := 0
	in := bufio.NewScanner(file)
	for in.Scan() {
		inSlice := stringToIntSlice(in.Text())
		conclusion := "Unsafe"
		if isInputSafe(&inSlice) {
			conclusion = "Safe"
			safeNum++
		}
		fmt.Println(inSlice, conclusion)
	}
	fmt.Println(safeNum)

	return
}
