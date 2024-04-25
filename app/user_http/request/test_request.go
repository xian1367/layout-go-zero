package request

type TestGetReq struct {
	Test string `form:"test,optional"`
}

func (r TestGetReq) ValidateFunc() error {
	return nil
}
