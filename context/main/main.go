package main

import (
	"fmt"
	"github.com/ethanfrey/godemo/context"
)

func main() {
	m := &context.MainControl{"My Controller"}

	for _, name := range []string{"Req 1", "Req 2", "FooBar"} {
		r := m.NewMainContext(name)
		fmt.Println("***")
		r.Log("Hello")
		fmt.Println(r.ControllerName())
		fmt.Println(r.RequestName())
	}
}
