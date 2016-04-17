package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) String() string {
	fmt.Println("foo")
	fmt.Println(s[0])
	a := sort.IntSlice(s)
	a.Sort()

	fmt.Println(s[0])
	return fmt.Sprint([]int(a))
}

func main() {
	d := Sequence{5, 1, 7, 3, 9}
	fmt.Println(d)
}
