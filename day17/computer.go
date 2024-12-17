package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	RegA    int64
	RegB    int64
	RegC    int64
	Program []int64
}

// LoadProgram reads register values and program instructions from a file
func (pc *Computer) LoadProgram(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register A: %d", &pc.RegA)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register B: %d", &pc.RegB)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register C: %d", &pc.RegC)

	scanner.Scan()
	scanner.Scan()

	opcodes := strings.Split(strings.TrimPrefix(scanner.Text(), "Program: "), ",")
	for _, v := range opcodes {
		op, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		pc.Program = append(pc.Program, op)
	}

	return nil
}
