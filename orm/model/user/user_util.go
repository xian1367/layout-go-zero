package user

import (
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"net/http"
)

func Get(id interface{}) (user User) {
	orm.DB.Where("id = ?", id).First(&user)
	return
}

func All() (users []User) {
	orm.DB.Find(&users)
	return
}

func Paginate(r *http.Request) (users []User, paging server.Paging) {
	paging = server.Paginate(
		r,
		orm.DB.Model(User{}),
		&users,
	)
	return
}
