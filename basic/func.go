package main

import (
    "fmt"
)

func reduce(fn func(int, int) int, data []int, begin ...int) int {
    total := 0
    if len(begin) > 0 {
        total = begin[0]
    }
    for i := 0; i < len(data); i++ {
        total = fn(data[i], total)
    }
    return total
}

func sum(a, b int) int {
  return a + b
}

func mult(a, b int) int {
  return a * b
}

func sub(a, b int) int {
  return a - b
}


func main() {
    data := [5]int{2, 3, 4, 5, 10}
    fmt.Println(data)
    answer := reduce(sum, data[:])
    fmt.Println(answer)
    answer = reduce(mult, data[:], 1)
    fmt.Println(answer)
    answer = reduce(sub, data[:])
    fmt.Println(answer)
}
