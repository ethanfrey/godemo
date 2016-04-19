package main

import (
	"fmt"
)

// const (
// 	a int = 3
// 	b     = 5
// 	c     = iota
// 	d
// )

type P struct {
	K string
	V float64
}

const (
	_         = iota
	a float64 = 5
	b
	c = (iota * iota)
	d
)

var (
	data []P = []P{P{"a", a}, P{"b", b}, P{"c", c}, P{"d", d}}
)

func showConst() {
	for _, x := range data {
		fmt.Println(x.K, x.V)
	}
}

func trace(s string) string {
	fmt.Println("Enter:", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving:", s)
}

func setup() []string {
	defer un(trace("setup"))
	const a, b, c = 1, 4, 6
	val := [10]string{a: "A", b: "B", c: "C", 2: "Silly"}
	// val := []string{a: "A", b: "B", c: "C", 2: "Silly"}
	fmt.Println("len:", len(val))
	return val[:]
}

func nosetup() []int {
	defer un(trace("nosetup"))
	return make([]int, 9)
}

func setupDemo() {
	data := nosetup()
	fmt.Println("len:", len(data))
	fmt.Println("cap:", cap(data))
	fmt.Println(data)
	for i, v := range data {
		fmt.Println(i, ":", v)
	}
}

func main() {
	// setupDemo()
	showConst()
}
