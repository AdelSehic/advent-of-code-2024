package main

func repeatElement[T any](element T, count int) []T {
	if count < 1 {
		return make([]T, 0)
	}
	result := make([]T, count) // Create a slice with length `count`
	for i := 0; i < count; i++ {
		result[i] = element
	}
	return result
}

func swap[T any](a, b *T){
	*a, *b = *b, *a
}
