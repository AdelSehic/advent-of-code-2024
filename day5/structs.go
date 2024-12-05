package main

type Pair struct {
	A int
	B int
}

type Page struct {
	Sequnce []int
	Members map[int]bool
}

type Rule struct {
	Value  int
	Before map[int]bool // values which have to come before
	After  map[int]bool // values which have to be after
}

type Input struct {
	Rules map[int]*Rule
	Pages []*Page
}
