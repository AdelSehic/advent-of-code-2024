package helpers

func ByteSequence(number, base int) [][]int {
	totalPerms := IntPow(base, number)
	permutations := make([][]int, totalPerms)

	for i := 0; i < totalPerms; i++ {
		bits := make([]int, number)
		num := i
		for j := 0; j < number; j++ {
			bits[j] = num % base
			num /= base
		}
		permutations[i] = bits
	}
	return permutations
}
