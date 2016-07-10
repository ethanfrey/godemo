package context

import "fmt"

type Logger interface {
	Log(string)
}

type MainController interface {
	Logger
	ControllerName() string
	NewContext(string) MainContext
}

type mainController struct {
	Prefix string
}

type MainContext interface {
	RequestContext
	MainController
}

type mainContext struct {
	RequestContext
	MainController
}

func (m *mainController) Log(msg string) {
	fmt.Println(m.Prefix, ":", msg)
}

func (m *mainController) ControllerName() string {
	return "MainControl"
}

// NewContext is called for each request that is handler by
// this controller
func (m *mainController) NewContext(req string) MainContext {
	// TODO: we can do a bit more customization that just appending
	// the two types if we wish here
	return &mainContext{
		NewContext(req),
		m,
	}
}

// NewController is called in the setup of the app
func NewController(prefix string) MainController {
	return &mainController{prefix}
}
