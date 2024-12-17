package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()
	pc := &Computer{}

	if err := pc.LoadProgram(os.Args[1]); err != nil {
		panic(err)
	}

	output := make([]int64, 0)
	for i := 0; i < len(pc.Program); {
		op, arg := pc.Program[i], pc.Program[i+1]

		combo := arg
		switch combo {
		case 4:
			combo = pc.RegA
		case 5:
			combo = pc.RegB
		case 6:
			combo = pc.RegC
		}

		switch op {
		case 0:
			pc.RegA = pc.RegA / power(2, combo)
		case 1:
			pc.RegB = pc.RegB ^ arg
		case 2:
			pc.RegB = combo & 7
		case 3:
			if pc.RegA != 0 {
				i = int(arg)
				continue
			}
		case 4:
			pc.RegB = pc.RegB ^ pc.RegC
		case 5:
			output = append(output, combo%8)
		case 6:
			pc.RegB = pc.RegA / power(2, combo)
		case 7:
			pc.RegC = pc.RegA / power(2, combo)
		default:
			fmt.Println("Wtf is that opcode", op)
		}
		i += 2
	}

	for _, v := range output {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\b \r\n")
	fmt.Println(time.Since(start))
}

func power(base, exp int64) int64 {
    return int64(math.Pow(float64(base), float64(exp)))
}
