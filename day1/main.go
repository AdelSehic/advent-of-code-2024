package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sortSlice(slc []int) {
	sort.Slice(slc, func(i, j int) bool {
		return slc[i] < slc[j]
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInputs(file *os.File, left, right *[]int) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, r int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		if err != nil {
			return err
		}
		*left = append(*left, l)
		*right = append(*right, r)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	return nil
}

func firstSolution(left, right []int) int {
	sortSlice(left)
	sortSlice(right)

	sum := 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}
	return sum
}

func secondSolution(left, right []int) int {
	repetitions := make(map[int]int)

	for _, v := range right {
		repetitions[v]++
	}

	score := 0
	for _, v := range left {
		score += v * repetitions[v]
	}

	return score
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	left, right := make([]int, 0), make([]int, 0)

	if err := getInputs(file, &left, &right); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	leftc := append([]int(nil), left...)
	rightc := append([]int(nil), right...)
	fmt.Println(firstSolution(leftc, rightc))

	fmt.Println(secondSolution(left, right))
}
