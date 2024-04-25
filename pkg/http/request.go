package http

type Request interface {
	ValidateFunc() error
}

type Req struct {
}

func (r Req) ValidateFunc() error {
	return nil
}
