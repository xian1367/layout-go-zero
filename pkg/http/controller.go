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
	R        *http.Request
	W        http.ResponseWriter
	UserID   string
	ClientIP string
	Success  Success
	Failure  Failure
	Error    error
}

func ControllerFunc(fn func(*Controller)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := ControllerDefault(w, r)
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
		c := ControllerDefault(w, r)

		var request T
		err := ControllerRequest(&c, &request)
		if err != nil {
			return
		}

		fn(&c, request)

		if c.Error == nil {
			httpx.WriteJson(w, http.StatusOK, c.Success)
		} else {
			httpx.WriteJson(w, http.StatusBadRequest, c.Failure)
		}
	}
}

func ControllerDefault(w http.ResponseWriter, r *http.Request) Controller {
	return Controller{
		W:        w,
		R:        r,
		UserID:   r.Context().Value("user_id").(string),
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

func ControllerRequest[T Request](c *Controller, request *T) (err error) {
	if err = httpx.Parse(c.R, request); err != nil {
		ValidationError(c.W, err)
		return
	}

	validator.SetTranslator(strings.Split(c.R.Header.Get("Accept-Language"), ",")...)
	err, errs := validator.Validate(c.R.Context(), request)
	if err != nil {
		ValidationErrors(c.W, errs)
		return
	}

	v := *request
	if err = v.ValidateFunc(c); err != nil {
		ValidationError(c.W, err)
	}
	return
}
