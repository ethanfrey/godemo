package main

import (
	"fmt"
	"github.com/ethanfrey/godemo/context"
)

func main() {
	m := context.NewController("My Controller")

	for _, name := range []string{"Req 1", "Req 2", "FooBar"} {
		r := m.NewContext(name)
		fmt.Println("***")
		r.Log("Hello")
		fmt.Println(r.ControllerName())
		fmt.Println(r.RequestName())
	}
}
