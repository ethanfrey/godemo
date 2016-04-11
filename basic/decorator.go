package main

import (
    "fmt"
)

type Op interface {
    Mult(int) int
}

type Trans func(int) int

func square(in int) int {
    return in * in
}

func double_in(orig Trans) Trans {
    wrap := func (in int) int {
        return orig(in * 2)
    }
    return wrap
}

func double_out(orig Trans) Trans {
    wrap := func (in int) int {
        return 2 * orig(in)
    }
    return wrap
}


type Orig struct {
    Val int
}

func (o *Orig)Mult(in int) int {
    return o.Val * in
}

type Wrap struct {
    Val Op
}

func (w *Wrap)Mult(in int) int {
    return w.Val.Mult(in + 5)
}


func main() {
    // decorate some functions
    foo := double_in(square)
    bar := double_out(square)
    // returns  6*6=36 and 2*3*3=18
    fmt.Println(foo(3), bar(3))

    item := Orig{4}
    fmt.Println(item.Mult(5))

    // use struct to decorate methods
    wrap := Wrap{&item}
    fmt.Println(wrap.Mult(5))

    double := Wrap{&wrap}
    fmt.Println(double.Mult(5))


}
