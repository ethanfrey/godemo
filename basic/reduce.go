package main

import (
	"fmt"
)

type IntReducer func(int, int) int

func Sum(x, total int) int {
	return x + total
}

func Mult(x, total int) int {
	return x * total
}

func IntReduce(data []int, reducer IntReducer, total int) int {
	for _, el := range data {
		total = reducer(el, total)
	}
	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(IntReduce(data, Sum, 0))
	fmt.Println(IntReduce(data, Mult, 1))
}
