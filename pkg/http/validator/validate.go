package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
)

func Validate(rtx context.Context, request interface{}) (err error, errors map[string]string) {
	err = Validator.StructCtx(rtx, request)
	if err == nil {
		return
	}
	errors = err.(validator.ValidationErrors).Translate(Trans)
	return
}
