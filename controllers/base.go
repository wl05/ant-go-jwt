package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"ant-go-jwt/common/utils"
	"ant-go-jwt/common/consts"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	token := base.Ctx.GetCookie("token")
	fmt.Println(token)
	b, _ := utils.CheckToken(token)
	if !b {
		base.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_LOGIN_ERROR,
			"msg":  consts.ERROR_DES_LOGIN_ERROR,
		}
		base.ServeJSON()
		return
	}
}
