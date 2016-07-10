package context

type RequestContext interface {
	RequestName() string
}

type MainContext interface {
	RequestContext
	MainController
}

type BasicContext struct {
	Name string
}

type RealContext struct {
	*BasicContext
	*MainControl
}

func (c *BasicContext) RequestName() string {
	return c.Name
}
