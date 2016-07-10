package context

type RequestContext interface {
	RequestName() string
}

type BasicContext struct {
	Name string
}

func (c *BasicContext) RequestName() string {
	return c.Name
}

// NewContext should really take a request object, this is just and example
func NewContext(req string) RequestContext {
	return &BasicContext{req}
}
