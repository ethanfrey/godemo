package main

import (
	"fmt"
	"github.com/ethanfrey/godemo/context"
)

func main() {
	m := &context.MainControl{"My Controller"}
	c := &context.BasicContext{"My Request"}
	r := &context.RealContext{c, m}
	r.Log("Hello")
	fmt.Println(r.ControllerName())
	fmt.Println(r.RequestName())
}
