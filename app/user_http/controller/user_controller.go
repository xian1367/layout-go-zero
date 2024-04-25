package controller

import (
	"github.com/xian1367/layout-go-zero/app/user_http/request"
	"github.com/xian1367/layout-go-zero/orm/model/user"
	"github.com/xian1367/layout-go-zero/pkg/http"
)

type UserController struct{}

// Index
// @Summary 列表
// @Description 列表
// @Tags user_http.user
// @Accept */*
// @Produce json
// @Param payload query http.Query true "payload"
// @Success 200 {object} http.Success{data=http.Paging{list=[]user.User}}
// @Failure 500 {object} http.Failure
// @Router /user [get]
func (ctrl *UserController) Index(c *http.Controller, req request.UserIndexReq) {
	_, pager := user.Paginate(c.Request)
	c.Ok(pager)
}

// Show
// @Summary 详情
// @Description 详情
// @Tags user_http.user
// @Accept */*
// @Produce json
// @Param payload body request.UserShowReq true "payload"
// @Success 200 {object} http.Success{data=user.User}
// @Failure 500 {object} http.Failure
// @Router /user/{id} [get]
func (ctrl *UserController) Show(c *http.Controller, req request.UserShowReq) {
	userModel := user.Get(req.ID)
	if userModel.ID == 0 {
		c.Abort404()
		return
	}
	c.Ok(userModel)
}

// Store
// @Summary 新增
// @Description 新增
// @Tags user_http.user
// @Accept json
// @Produce json
// @Param payload body request.UserStoreReq true "payload"
// @Success 201 {object} http.Success{data=user.User}
// @Failure 500 {object} http.Failure
// @Router /user [post]
func (ctrl *UserController) Store(c *http.Controller, req request.UserStoreReq) {
	userModel := user.User{}
	c.DB.Create(&userModel)
	if userModel.ID == 0 {
		c.Abort("创建失败，请稍后尝试~")
	}
	c.Created(userModel)
}

// Update
// @Summary 更新
// @Description 更新
// @Tags user_http.user
// @Accept json
// @Produce json
// @Param payload body request.UserUpdateReq true "payload"
// @Success 200 {object} http.Success{data=user.User}
// @Failure 500 {object} http.Failure
// @Router /user/{id} [put]
func (ctrl *UserController) Update(c *http.Controller, req request.UserUpdateReq) {
	userModel := user.Get(req.ID)
	if userModel.ID == 0 {
		c.Abort404()
		return
	}

	rowsAffected := c.DB.Save(&userModel).RowsAffected
	if rowsAffected == 0 {
		c.Abort("更新失败，请稍后尝试~")
		return
	}
	c.Ok(userModel)
}

// Destroy
// @Summary 删除
// @Description 删除
// @Tags user_http.user
// @Accept json
// @Produce json
// @Param payload body request.UserDestroyReq true "payload"
// @Success 200
// @Failure 500 {object} http.Failure
// @Router /user/{id} [delete]
func (ctrl *UserController) Destroy(c *http.Controller, req request.UserDestroyReq) {
	userModel := user.Get(req.ID)
	if userModel.ID == 0 {
		c.Abort404()
		return
	}

	rowsAffected := c.DB.Delete(&userModel).RowsAffected
	if rowsAffected == 0 {
		c.Abort("删除失败，请稍后尝试~")
		return
	}
}
