package helpers

import "strconv"

func IntPow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func IntLen(num int) int {
	return len(strconv.Itoa(num))
}

func Uint64Len(num uint64) int {
	return len(strconv.FormatUint(num, 10))
}

func ConcatInts(a, b int) int {
	return a*IntPow(10, IntLen(b)) + b
}
