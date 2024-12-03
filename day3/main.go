package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var sb strings.Builder

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	input := sb.String()

	firstDont := strings.Index(input, "don't()")
	lastDo := strings.LastIndex(input, "do()")
	inputBegin := input[:firstDont]
	inputEnd := input[lastDo:]

	left := "do()"
	right := "don't()"
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	matches := rx.FindAllStringSubmatch(input, -1)

	extractMul := func(text string) [][2]uint64 {
		mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		mulMatches := mulRegex.FindAllStringSubmatch(text, -1)

		results := make([][2]uint64, 0, len(mulMatches))
		for _, match := range mulMatches {
			if len(match) == 3 {
				var num1, num2 uint64
				fmt.Sscanf(match[1], "%d", &num1)
				fmt.Sscanf(match[2], "%d", &num2)
				results = append(results, [2]uint64{num1, num2})
			}
		}
		return results
	}

	// Collect results from all regions
	results := extractMul(inputBegin) // From inputBegin
	for _, match := range matches {   // From all matched substrings
		results = append(results, extractMul(match[1])...)
	}
	results = append(results, extractMul(inputEnd)...) // From inputEnd

	// Sum the results
	sum := uint64(0)
	for _, pair := range results {
		sum += pair[0] * pair[1]
	}

	// Output the results
	fmt.Println("Extracted mul pairs:", results)
	fmt.Println("Sum of all multiplications:", sum)
}

// func main() {
// 	file, err := os.Open(os.Args[1])
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	scanner := bufio.NewScanner(file)
//
// 	var sb strings.Builder
//
// 	for scanner.Scan() {
// 		sb.WriteString(scanner.Text())
// 	}
// 	input := sb.String()
//
// 	doAndDont := strings.Split(input, "don't()")
// 	dos := make([]string, 0)
//
// 	dos = append(dos, doAndDont[0])
//
// 	for i := 1; i < len(doAndDont); i++ {
// 		doSplit := strings.Split(doAndDont[i], "do()")
// 		if len(doSplit) == 2 {
// 			dos = append(dos, doSplit[1])
// 		}
// 	}
//
// 	dos = append(dos, doAndDont[len(doAndDont)-1])
// 	allDos := strings.Join(dos, "")
//
// 	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
// 	matches := re.FindAllStringSubmatch(allDos, -1)
//
// 	sum := uint64(0)
// 	// Extract numbers from matches
// 	results := make([][2]uint64, 0, len(matches))
// 	for _, match := range matches {
// 		if len(match) == 3 {
// 			var num1, num2 uint64
// 			fmt.Sscanf(match[1], "%d", &num1)
// 			fmt.Sscanf(match[2], "%d", &num2)
// 			results = append(results, [2]uint64{num1, num2})
// 			sum += num1 * num2
// 		}
// 	}
//
// 	fmt.Println(results, sum)
// }
