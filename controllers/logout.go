package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/common/utils"
	"fmt"
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
	//accessToken := this.Ctx.GetCookie("accessToken")
	refreshToken := this.Ctx.GetCookie("refreshToken")
	isValidRefreshToken, token := utils.CheckToken(refreshToken)
	fmt.Println("=============")
	fmt.Println(isValidRefreshToken)
	fmt.Println(*token)
	utils.RClient.Set(string(refreshToken), refreshToken, 0)
	fmt.Println(utils.RClient.Get(string(refreshToken)))
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登出成功",
	}
	this.ServeJSON()
	return
}
