package main

import (
	"fmt"
)

type IntReducer interface {
	Init() int
	Reduce(int, int) int
}

type IntSummer struct{}

func (IntSummer) Init() int {
	return 0
}

func (IntSummer) Reduce(a, b int) int {
	return a + b
}

type IntMult struct{}

func (IntMult) Init() int {
	return 1
}

func (IntMult) Reduce(a, b int) int {
	return a * b
}

func IntReduce(data []int, reducer IntReducer) int {
	total := reducer.Init()
	for _, el := range data {
		total = reducer.Reduce(el, total)
	}
	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(IntReduce(data, IntSummer{}))
	fmt.Println(IntReduce(data, IntMult{}))
}
