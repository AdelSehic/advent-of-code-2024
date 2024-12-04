package main

import (
	"fmt"
	"os"
)

func main() {
	input := &Input{}
	input.LoadDataWithPadding(os.Args[1])
	fmt.Println(input.xmasCount())
}
