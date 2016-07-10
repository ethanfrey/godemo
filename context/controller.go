package context

import "fmt"

type Logger interface {
	Log(string)
}

type MainController interface {
	Logger
	ContollerName() string
}

type MainControl struct {
	Prefix string
}

func (m *MainControl) Log(msg string) {
	fmt.Println(m.Prefix, ":", msg)
}

func (m *MainControl) ControllerName() string {
	return "MainControl"
}
