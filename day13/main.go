package main

import (
	"fmt"
	"os"
	"time"
)

func (p *Prize) Cramer() (int, int) {
	det := p.ButtonA.XOffset*p.ButtonB.YOffset - p.ButtonA.YOffset*p.ButtonB.XOffset

	detA := p.X*p.ButtonB.YOffset - p.Y*p.ButtonB.XOffset
	detB := p.ButtonA.XOffset*p.Y - p.ButtonA.YOffset*p.X

	if detA%det != 0 || detB%det != 0 {
		return 0, 0
	}
	return detA / det, detB / det
}

func main() {
	start := time.Now()
	in := &Input{}
	in.GetPrizes(os.Args[1], 0)

	costs := 0
	for _, pr := range in.Prizes {
		solA, solB := pr.Cramer()
		cost := solA*3 + solB
		costs += cost
	}
	fmt.Println("Part1 : ", costs, time.Since(start))

	start = time.Now()
	in = &Input{}
	in.GetPrizes(os.Args[1], 10000000000000)

	costs = 0
	for _, pr := range in.Prizes {
		solA, solB := pr.Cramer()
		cost := solA*3 + solB
		costs += cost
	}
	fmt.Println("Part2 : ", costs, time.Since(start))
}
