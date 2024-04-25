package factory

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/xian1367/layout-go-zero/orm/model/user"
)

func MakeUsers(count int) []user.User {
	var objs []user.User

	for i := 0; i < count; i++ {
		userModel := user.User{}
		userModel.Name = gofakeit.Name()
		objs = append(objs, userModel)
	}

	return objs
}
