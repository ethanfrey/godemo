package main

import (
    "fmt"
)

func Sqrt(x float64, iter ...int) float64 {
    count := 3
    if len(iter) > 0 {
      count = iter[0]
    }
    z := x
    for ; count > 0; count-- {
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2.0, 10))
    fmt.Println(Sqrt(4.0, 6))
    fmt.Println(Sqrt(10.0))
}
