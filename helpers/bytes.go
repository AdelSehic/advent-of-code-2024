package helpers

func ByteSequence(number int) [][]int {
	permutations := make([][]int, 1<<number)
	for i := 0; i < (1 << number); i++ {
		bits := make([]int, number)
		for j := 0; j < number; j++ {
			bits[j] = (i >> j) & 1
		}
		permutations[i] = bits
	}
	return permutations
}
