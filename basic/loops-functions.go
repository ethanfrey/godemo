package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64, iter ...int) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    count := 3
    if len(iter) > 0 {
      count = iter[0]
    }
    z := x
    for ; count > 0; count-- {
        z = z - (z * z - x) / (2 * z)
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(-10.0))
    fmt.Println(Sqrt(2.0, 10))
    fmt.Println(Sqrt(4.0, 6))
    fmt.Println(Sqrt(10.0))
}
