package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	Outputer
}

func NewPerson(name string, o Outputer) Person {
	if o == nil {
		o = NormalOutput{}
	}
	return Person{name, o}
}

func (p *Person) Display() {
	p.Output(p.name)
}

type Outputer interface {
	Output(name string)
}

type NormalOutput struct{}

func (o NormalOutput) Output(name string) {
	fmt.Println(name)
}

type BoldOutput struct{}

func (o BoldOutput) Output(name string) {
	fmt.Println(strings.ToUpper(name))
}

func main() {
	fmt.Println("Example of exposing some method to override (as embedding doesn't do this)")
	p1 := NewPerson("John Hughes", nil)
	p2 := NewPerson("Someone else", BoldOutput{})
	p1.Display()
	p2.Display()
}
