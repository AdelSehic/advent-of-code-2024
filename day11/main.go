package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	// "os"
)

const (
	INPUT1 = "input1"
	INPUT2 = "input2"
	TEST1  = "test1"
	TEST2  = "test2"
)

type Cached struct {
	Number     uint64
	Iterations int
}

var cache = make(map[Cached]uint64)

func RecursiveEvolve(number uint64, exit, iteration int) uint64 {
	if iteration >= exit {
		return 1
	}

	cacheEntry := Cached{
		Number:     number,
		Iterations: exit - iteration,
	}

	if _, found := cache[cacheEntry]; found {
		return cache[cacheEntry]
	}

	if number == 0 {
		cache[cacheEntry] = RecursiveEvolve(1, exit, iteration+1)
		return cache[cacheEntry]
	}

	digits := []uint64{}
	modifiedNumber := number
	for modifiedNumber > 0 {
		digits = append([]uint64{modifiedNumber % 10}, digits...)
		modifiedNumber /= 10
	}

	if len(digits)%2 == 0 {
		mid := len(digits) / 2
		left := uint64(0)
		right := uint64(0)
		for i := 0; i < mid; i++ {
			left = left*10 + digits[i]
			right = right*10 + digits[mid+i]
		}
		cache[cacheEntry] = RecursiveEvolve(left, exit, iteration+1) + RecursiveEvolve(right, exit, iteration+1)
		return cache[cacheEntry]
	}
	cache[cacheEntry] = RecursiveEvolve(number*2024, exit, iteration+1)
	return cache[cacheEntry]
}

func main() {
	start := time.Now()
	sum := 0
	input, exit := MakeInput(os.Args[1])
	var wg sync.WaitGroup
	for _, v := range input {
		wg.Add(1)
		// go func() {
		sum += int(RecursiveEvolve(v, exit, 0))
		wg.Done()
		// }()
	}
	wg.Wait()
	fmt.Println(sum, time.Since(start))
}

func MakeInput(what string) ([]uint64, int) {
	switch what {
	case TEST1:
		return []uint64{0, 1, 10, 99, 999}, 1
	case TEST2:
		return []uint64{125, 17}, 6
	case INPUT1:
		return []uint64{8435, 234, 928434, 14, 0, 7, 92446, 8992692}, 25
	case INPUT2:
		return []uint64{8435, 234, 928434, 14, 0, 7, 92446, 8992692}, 75
	default:
		return []uint64{0}, 10
	}
}
