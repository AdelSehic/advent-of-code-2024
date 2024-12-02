package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sliceIsSorted(input []int) bool {
	if len(input) < 2 {
		return true
	}

	compareFunc := func(a, b int) bool {
		return a > b
	}
	if input[0] < input[1] {
		compareFunc = func(a, b int) bool {
			return a < b
		}
	}

	return sort.SliceIsSorted(input, func(i, j int) bool {
		return compareFunc(input[i], input[j])
	})
}

func changesAreSafe(input []int) bool {
	if len(input) < 2 {
		return false
	}

	for i := 1; i < len(input); i++ {
		delta := abs(input[i-1] - input[i])
		if delta == 0 || delta > 3 {
			return false
		}
	}
	return true
}

func inputIsSafe(input []int) bool {
	return sliceIsSorted(input) && changesAreSafe(input)
}

func bruteForceFix(input []int) bool {
	for i := 0; i < len(input); i++ {
		newSlice := append([]int{}, input...)
		newSlice = append(newSlice[:i], newSlice[i+1:]...)
		if inputIsSafe(newSlice) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	var sol1, sol2 int

	in := bufio.NewScanner(file)
	for in.Scan() {
		inSlice := stringToIntSlice(in.Text())
		conclusion := "Unsafe"
		if inputIsSafe(inSlice) {
			conclusion = "Safe"
			sol1++
			sol2++
		} else if bruteForceFix(inSlice) {
			conclusion = "Safe after fix"
			sol2++
		}
		fmt.Println(inSlice, conclusion)
	}
	fmt.Printf("Solution 1: %d; Solution2: %d\r\n", sol1, sol2)

	return
}
