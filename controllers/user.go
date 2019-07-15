package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/models"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @获取用户列表
// @Description 获取用户列表的
// @Success 200 请求成功
// @Success 1102   请求出错
// @router / [get]
func (this *UserController) GetUsers() {
	users, err := models.GetUsers()
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"data": users,
	}
	this.ServeJSON()
	return
}
