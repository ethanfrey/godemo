package main

import (
	"fmt"
)

func fill(data []int) {
	for i := range data {
		data[i] = 2*i + 1
	}
	// no effect
	data = data[2:]
}

func main() {
	a, b := [10]int{}, [10]int{}
	c, d := a[:], b[5:8]

	fill(c)
	fmt.Println(c)
	fmt.Println(a)

	fill(d)
	fmt.Println(d)
	fmt.Println(b)

	e := append(d, 17, 19, 1, 4)
	fmt.Println(e)

	foo := append(d, e...)
	fmt.Println(foo)
}
