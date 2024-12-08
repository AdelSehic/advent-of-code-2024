package main

import (
	"os"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func main() {
	input := &helpers.Field{}
	input.LoadDataWithPadding(os.Args[1], "X")
}
