package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/common/utils"
)

// Operations about Logout
type LogoutController struct {
	BaseController
}

// @Title 用户登出
// @Description 用户登出
// @Success 200 登出成功
// @router /logout [post]
func (this *LogoutController) Logout() {
	refreshToken := this.Ctx.GetCookie("refreshToken")
	utils.RClient.Set(string(refreshToken), "exited", 0) // 这还要设置一下过期时间

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登出成功",
	}
	this.ServeJSON()
	return
}
