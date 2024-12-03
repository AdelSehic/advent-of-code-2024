package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var sb strings.Builder

	sb.WriteString("prepend this to fix edge case : do()")
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	input := sb.String()

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	donts := strings.Split(input, "don't()")
	sum := uint64(0)
	for _, v := range donts {
		doos := strings.Split(v, "do()")[1:]
		str := ""
		for _, a := range doos {
			str += a
		}
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			num1, err1 := strconv.ParseUint(match[1], 10, 64)
			num2, err2 := strconv.ParseUint(match[2], 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing numbers:", err1, err2)
				continue
			}
			sum += num1 * num2
		}
	}
	fmt.Println(sum)

	//
	// sum := uint64(0)
	// // Extract numbers from matches
	// results := make([][2]uint64, 0, len(matches))
	// for _, match := range matches {
	// 	if len(match) == 3 {
	// 		var num1, num2 uint64
	// 		fmt.Sscanf(match[1], "%d", &num1)
	// 		fmt.Sscanf(match[2], "%d", &num2)
	// 		results = append(results, [2]uint64{num1, num2})
	// 		sum += num1 * num2
	// 	}
	// }
	//
	// fmt.Println(results, sum)
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
