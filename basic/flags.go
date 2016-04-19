package main

import (
	"flag"
	"fmt"
)

func fib(cnt int) int {
	a, b := 0, 1
	for i := 1; i < cnt; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	val := flag.Int("step", 3, "Which fibonaci number do you want?")
	flag.Parse()

	res := fib(*val)
	fmt.Printf("fib(%d) = %d\n", *val, res)
}
