package request

// UserIndexReq 列表
type UserIndexReq struct {
	//Name string `form:"name,optional"`
}

func (r UserIndexReq) ValidateFunc() error {
	return nil
}

// UserShowReq 详情
type UserShowReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserShowReq) ValidateFunc() error {
	return nil
}

// UserStoreReq 新增
type UserStoreReq struct{}

func (r UserStoreReq) ValidateFunc() error {
	return nil
}

// UserUpdateReq 修改
type UserUpdateReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserUpdateReq) ValidateFunc() error {
	return nil
}

// UserDestroyReq 删除
type UserDestroyReq struct {
	ID int32 `path:"id" validate:"required"`
}

func (r UserDestroyReq) ValidateFunc() error {
	return nil
}
