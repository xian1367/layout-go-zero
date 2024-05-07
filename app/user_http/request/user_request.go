package request

import (
	"github.com/xian1367/layout-go-zero/pkg/http"
)

// UserIndexReq 列表
type UserIndexReq struct {
	//Name string `form:"name,optional"`
}

func (r UserIndexReq) ValidateFunc(c *http.Controller) error {

	return nil
}

// UserShowReq 详情
type UserShowReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserShowReq) ValidateFunc(c *http.Controller) error {
	return nil
}

// UserStoreReq 新增
type UserStoreReq struct{}

func (r UserStoreReq) ValidateFunc(c *http.Controller) error {
	return nil
}

// UserUpdateReq 修改
type UserUpdateReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserUpdateReq) ValidateFunc(c *http.Controller) error {
	return nil
}

// UserDestroyReq 删除
type UserDestroyReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserDestroyReq) ValidateFunc(c *http.Controller) error {
	return nil
}
