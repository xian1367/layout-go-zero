package user

import (
	http2 "github.com/xian1367/layout-go-zero/pkg/http"
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

func Paginate(r *http.Request) (users []User, paging http2.Paging) {
	paging = http2.Paginate(
		r,
		orm.DB.Model(User{}),
		&users,
	)
	return
}
