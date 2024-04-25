package http

import (
	"github.com/xian1367/layout-go-zero/pkg/http/validator"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type Controller struct {
	DB       *gorm.DB
	Logger   logx.Logger
	Request  *http.Request
	ClientIP string
	Success
	Failure
	Error error
}

func ControllerFunc(fn func(*Controller)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := ControllerDefault(r)
		fn(&c)

		if c.Error == nil {
			httpx.WriteJson(w, http.StatusOK, c.Success)
		} else {
			httpx.WriteJson(w, http.StatusBadRequest, c.Failure)
		}
	}
}

func ControllerHandler[T Request](fn func(*Controller, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request T
		err := ControllerRequest(w, r, &request)
		if err != nil {
			return
		}

		c := ControllerDefault(r)
		fn(&c, request)

		if c.Error == nil {
			httpx.WriteJson(w, http.StatusOK, c.Success)
		} else {
			httpx.WriteJson(w, http.StatusBadRequest, c.Failure)
		}
	}
}

func ControllerDefault(r *http.Request) Controller {
	return Controller{
		Request:  r,
		ClientIP: httpx.GetRemoteAddr(r),
		DB:       orm.DB.WithContext(r.Context()),
		Logger:   logx.WithContext(r.Context()),
		Success: Success{
			Code: http.StatusOK,
		},
		Failure: Failure{
			Code: http.StatusBadRequest,
		},
	}
}

func ControllerRequest[T Request](w http.ResponseWriter, r *http.Request, request *T) (err error) {
	if err = httpx.Parse(r, request); err != nil {
		ValidationError(w, err)
		return
	}

	validator.SetTranslator(strings.Split(r.Header.Get("Accept-Language"), ",")...)
	err, errs := validator.Validate(r.Context(), request)
	if err != nil {
		ValidationErrors(w, errs)
		return
	}

	v := *request
	if err = v.ValidateFunc(); err != nil {
		ValidationError(w, err)
	}
	return
}
