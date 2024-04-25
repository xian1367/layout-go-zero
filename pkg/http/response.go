package http

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"math/rand"
	"net/http"
	"reflect"
)

type Success struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Failure struct {
	Code   uint32            `json:"code"`
	Msg    string            `json:"msg"`
	Errors map[string]string `json:"errors"`
}

func (c *Controller) Ok(data any, codes ...uint32) {
	if len(codes) > 0 {
		c.Failure.Code = codes[0]
	}
	if str, ok := data.(string); ok {
		c.Success.Msg = str
	} else {
		c.Success.Data = data
	}
}

func (c *Controller) Created(data any) {
	c.Success.Code = http.StatusCreated
	c.Ok(data)
}

func (c *Controller) Abort(msg string, codes ...uint32) {
	if len(codes) > 0 {
		c.Failure.Code = codes[0]
	}
	c.Failure.Msg = msg
	c.Error = errors.New(c.Failure.Msg)
}

func (c *Controller) Abort401() {
	c.Abort("登录失败", http.StatusUnauthorized)
}

func (c *Controller) Abort404() {
	c.Abort("数据不存在", http.StatusNotFound)
}

func ValidationError(w http.ResponseWriter, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, Failure{
		Code: http.StatusBadRequest,
		Msg:  err.Error(),
	})
}

// ValidationErrors 处理表单验证不通过的错误，返回的 JSON 示例：
//
//	{
//	    Errors: {
//	        "mobile": "手机号长度必须为 11 位的数字"
//	    },
//	    Message: "手机号长度必须为 11 位的数字"
//	}
func ValidationErrors(w http.ResponseWriter, errors map[string]string) {
	//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Failure{
	//	"message": "表单验证错误",
	//	"errors":  errors,
	//})
	keys := reflect.ValueOf(errors).MapKeys()
	key := keys[rand.Intn(len(keys))].Interface()
	httpx.WriteJson(w, http.StatusUnprocessableEntity, Failure{
		Code:   http.StatusUnprocessableEntity,
		Msg:    errors[key.(string)],
		Errors: errors,
	})
}
