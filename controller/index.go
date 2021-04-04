package controller

import (
	"strconv"
	"time"

	"github.com/anthonyzero/go-quick-api/core"
	"github.com/anthonyzero/go-quick-api/models/dao"
	"github.com/anthonyzero/go-quick-api/models/req"
	"github.com/anthonyzero/go-quick-api/pkg/logger"
	"github.com/anthonyzero/go-quick-api/public"
	"github.com/gin-gonic/gin"
)

//路由注册
func IndexRegister(router *gin.RouterGroup) {
	index := IndexController{}
	router.POST("/user_create", index.UserCreate)
	router.GET("/user_delete", index.UserDelete)
}

type IndexController struct {
}

// UserCreate godoc
// @Summary 创建用户
// @Description 创建用户
// @Tags 用户管理
// @ID /index/user_create
// @Accept  json
// @Produce  json
// @Param body body req.CreateUserRequest true "body"
// @Success 200 {object} core.Result "success"
// @Router /index/user_create [post]
func (index *IndexController) UserCreate(c *gin.Context) {
	params := &req.CreateUserRequest{}
	if err := c.ShouldBindJSON(params); err != nil {
		core.Error(c, public.ErrParameter)
		return
	}
	user := &dao.SmUser{
		UserName:   params.Username,
		Password:   params.Password,
		UserCode:   "testone",
		DepartID:   123,
		State:      1,
		CreateTime: time.Now(),
	}
	if err := user.CreateUser(); err != nil {
		logger.Errorf("save user occur err : %v", err)
		core.Error(c, public.InternalServerError)
		return
	}
	core.Success(c, nil)
}

// UserDelete godoc
// @Summary 删除用户
// @Description 根据用户ID删除
// @Tags 用户管理
// @ID /index/user_delete
// @Accept  json
// @Produce  json
// @Param userId query int true "用户ID"
// @Success 200 {object} core.Result "success"
// @Router /index/user_delete [get]
func (index *IndexController) UserDelete(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		core.Error(c, err)
		return
	}
	if userId <= 0 {
		core.Error(c, public.ErrParameter)
		return
	}
	user := dao.SmUser{UserID: userId}
	if err := user.DeleteUser(); err != nil {
		core.Error(c, public.InternalServerError)
		return
	}
	core.Success(c, nil)
}
