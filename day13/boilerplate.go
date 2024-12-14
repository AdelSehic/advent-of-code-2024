package main

import (
	"bufio"
	"fmt"
	"os"
)

type Button struct {
	XOffset int
	YOffset int
}

type Prize struct {
	ButtonA Button
	ButtonB Button
	X       int
	Y       int
}

type Input struct {
	Prizes []*Prize
}

func (in *Input) GetPrizes(infile string, offset int) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for {
		p := &Prize{}

		if !scanner.Scan() {
			break
		}
		lineA := scanner.Text()
		fmt.Sscanf(lineA, "Button A: X+%d, Y+%d", &p.ButtonA.XOffset, &p.ButtonA.YOffset)

		if !scanner.Scan() {
			break
		}
		lineB := scanner.Text()
		fmt.Sscanf(lineB, "Button B: X+%d, Y+%d", &p.ButtonB.XOffset, &p.ButtonB.YOffset)

		if !scanner.Scan() {
			break
		}
		linePrize := scanner.Text()
		fmt.Sscanf(linePrize, "Prize: X=%d, Y=%d", &p.X, &p.Y)
		p.X += offset
		p.Y += offset

		scanner.Scan()

		in.Prizes = append(in.Prizes, p)
	}
}

func (in *Input) PrintInputs() {
	for _, v := range in.Prizes {
		v.PrintPrize()
	}
}

func (pr *Prize) PrintPrize() {
	fmt.Printf("A: X+%d, Y+%d\r\nB: X+%d, Y+%d\r\nPrize: X=%d Y=%d\r\n",
		pr.ButtonA.XOffset, pr.ButtonA.YOffset, pr.ButtonB.XOffset, pr.ButtonB.YOffset, pr.X, pr.Y)
}
