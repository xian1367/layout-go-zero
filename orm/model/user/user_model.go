package user

import (
	"github.com/xian1367/layout-go-zero/orm/gen/gen"
	"github.com/xian1367/layout-go-zero/orm/model"
)

type User struct {
	model.BaseModel
	gen.User
}
