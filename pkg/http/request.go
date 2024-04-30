package http

type Request interface {
	ValidateFunc(*Controller) error
}

type Req struct {
}

func (r Req) ValidateFunc(c *Controller) error {
	return nil
}
