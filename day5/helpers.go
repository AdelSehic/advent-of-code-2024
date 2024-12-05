package main

import "fmt"

func (in *Input) PrintRules() {
	for _, v := range in.Rules {
		v.printRule()
	}
}

func (r *Rule) printRule() {
	fmt.Println("Value: ", r.Value)
	fmt.Printf("Before: ")
	for v := range r.Before {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
	fmt.Printf("After: ")
	for v := range r.After {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
	fmt.Println("")
}

func (in *Input) PrintPages() {
	for _, v := range in.Pages {
		fmt.Println(v.Sequnce)
	}
}

func splitSlice(slice []int, index int) ([]int, []int) {
	if index < 0 || index > len(slice) {
		panic("index out of bound")
	}
	return slice[:index], slice[index+1:]
}
