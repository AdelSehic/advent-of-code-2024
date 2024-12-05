package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *Input) MakePair(str string) {
	nums := strings.Split(str, "|")
	numA, _ := strconv.Atoi(nums[0])
	numB, _ := strconv.Atoi(nums[1])

	if _, found := in.Rules[numA]; !found {
		in.Rules[numA] = &Rule{
			Value:  numA,
			Before: make(map[int]bool),
			After:  make(map[int]bool),
		}
	}
	if _, found := in.Rules[numB]; !found {
		in.Rules[numB] = &Rule{
			Value:  numB,
			Before: make(map[int]bool),
			After:  make(map[int]bool),
		}
	}

	in.Rules[numA].After[numB] = true
	in.Rules[numB].Before[numA] = true
}

func (in *Input) MakePage(str string) {
	page := &Page{
		Sequnce: make([]int, 0),
		Members: make(map[int]bool),
	}

	vals := strings.Split(str, ",")
	for _, v := range vals {
		num, _ := strconv.Atoi(v)
		page.Members[num] = true
		page.Sequnce = append(page.Sequnce, num)
	}

	in.Pages = append(in.Pages, page)
}

func (page *Page) CheckRules(rule map[int]*Rule) bool {
	for i := range page.Sequnce {
		left, right := splitSlice(page.Sequnce, i)
		for _, v := range left {
			if rule[page.Sequnce[i]].After[v] {
				fmt.Printf("Rule violation for %+v ... %d placed after %d\r\n", page.Sequnce, page.Sequnce[i], v)
				return false
			}
		}
		for _, v := range right {
			if rule[page.Sequnce[i]].Before[v] {
				fmt.Printf("Rule violation for %+v ... %d placed before %d\r\n", page.Sequnce, page.Sequnce[i], v)
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	in := &Input{}
	in.Rules = make(map[int]*Rule)

	scanner := bufio.NewScanner(file)
	for scanner.Scan(); scanner.Text() != ""; scanner.Scan() {
		in.MakePair(scanner.Text())
	}

	for scanner.Scan() {
		in.MakePage(scanner.Text())
	}

	properPages := make([]int, 0)
	wrongPages := make([]int, 0)
	for i, page := range in.Pages {
		if page.CheckRules(in.Rules) {
			fmt.Printf("Satisfactory order for %+v\r\n", page.Sequnce)
			properPages = append(properPages, i)
		} else {
			wrongPages = append(wrongPages, i)
		}
	}

	sum := 0
	for _, v := range properPages {
		page := in.Pages[v].Sequnce
		sum += page[len(page)/2]
	}
	fmt.Println("Sum of the middle of correct pages:", sum)
}
