package main

import (
    "fmt"
    "strconv"
)

func main() {
    data := [5]int{4, 7, 2, 1, 7}
    output := make([]string, 0, 6)

    fmt.Println(len(data))
    fmt.Println(len(output))
    fmt.Println(data)
    fmt.Println(output)

    for i, v := range data {
        output = output[0:i+1]
        output[i] = strconv.Itoa(v)
    }

    fmt.Println(len(data))
    fmt.Println(len(output))
    fmt.Println(data)
    fmt.Println(output)
}
