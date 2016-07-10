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

type MainContext struct {
	RequestContext
	*MainControl
}

func (m *MainControl) Log(msg string) {
	fmt.Println(m.Prefix, ":", msg)
}

func (m *MainControl) ControllerName() string {
	return "MainControl"
}

func (m *MainControl) NewMainContext(req string) *MainContext {
	return &MainContext{
		NewContext(req),
		m,
	}
}
