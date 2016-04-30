package main

import (
	"errors"
	"fmt"
	"sort"
)

func findMatch(data []int, target int) ([]int, error) {
	if target <= 0 {
		// prevent infinite recursion depth...
		panic("Bad code...")
	}
	// this is the only case that should happen....
	cur := data[0]
	if cur == target {
		// we have a winner!
		fmt.Println("Stopped at", cur)
		return []int{cur}, nil
	} else if len(data) == 1 {
		// we hit the end, stop...
		return nil, errors.New("Bad path")
	} else {
		// try to use it if possible
		if cur < target {
			matches, err := findMatch(data[1:], target-cur)
			if err == nil {
				return append(matches, cur), nil
			}
		}
		// and we can recurse without our value in the stack
		matches, err := findMatch(data[1:], target)
		if err == nil {
			return matches, nil
		} else {
			return nil, err
		}
	}
}

func main() {
	// data := []int{3, 6, 8, 9, 12, 7, 22}
	// target := 30
	// data := []int{1, 2, 3, 4}
	// target := 3
	data := []int{18897109, 12828837, 9461105, 6371773, 5965343, 5946800, 5582170, 5564635, 5268860, 4552402, 4335391, 4296250, 4224851, 4192887, 3439809, 3279833, 3095313, 2812896, 2783243, 2710489, 2543482, 2356285, 2226009, 2149127, 2142508, 2134411}
	target := 100000000

	answer, err := findMatch(data, target)
	if err != nil {
		fmt.Println("No solution")
	} else {
		count := 0
		for _, el := range answer {
			count += el
		}
		if count != target {
			fmt.Println("Bad answer, some bug?")
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(answer)))
			fmt.Println("Found", len(answer), "cities that match:")
			fmt.Println(answer)
		}
	}
}
