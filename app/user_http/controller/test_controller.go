package controller

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/xian1367/layout-go-zero/pkg/http"
)

type TestController struct{}

func (ctrl *TestController) Test(c *http.Controller) {
	spew.Dump(c.Success)
	return
}
