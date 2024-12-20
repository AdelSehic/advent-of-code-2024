package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var substrings = map[string]bool{}

var memo = map[string]bool{}

func RecursivePrints(str string) bool {

	if val, exists := memo[str]; exists {
		return val
	}

	width := len(str)
	if width == 0 {
		memo[str] = true
		return true
	}

	for i := 1; i <= width; i++ {
		substr := str[:i]
		remaining := str[i:]

		if substrings[substr] {
			if RecursivePrints(remaining) {
				memo[str] = true
				return true
			}
		}
	}

	memo[str] = false
	return false
}

func main() {
	start := time.Now()
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for _, v := range strings.Split(scanner.Text(), ", ") {
		substrings[v] = true
	}
	scanner.Scan()

	possible := 0
	for scanner.Scan() {
		if RecursivePrints(scanner.Text()) {
			possible++
		}
	}
	fmt.Println("Part1: ", possible, time.Since(start))
}
