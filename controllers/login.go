package controllers

import (
	"ant-go-jwt/common/utils"
	"ant-go-jwt/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"ant-go-jwt/common/consts"
)

// Operations about Login
type LoginController struct {
	beego.Controller
}

// @Title 用户登录
// @Description 用户登录
// @Param   email    formData    string  "demo@qq.com"  true  "邮箱"
// @Param   password    formData    string  "123"    true "密码"
// @Success 200 登录成功
// @Success 1000   参数错误
// @Success 1001   请求出错
// @Success 2001 用户不存在
// @Success 2002 用户名或者密码错误
// @router /login [post]
func (this *LoginController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")
	// 参数是否为空
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}
	// 验证邮箱是否存在
	user, err := models.GetUserByEmail(email)
	if err == orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_NOT_EXIST,
			"msg":  consts.ERROR_DES_USER_NOT_EXIST,
		}
		this.ServeJSON()
		return
	}
	// 验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USERNAME_OR_PASSWORD_ERROR,
			"msg":  consts.ERROR_DES_USERNAME_OR_PASSWORD_ERROR,
		}
		this.ServeJSON()
		return
	}
	token, err := utils.CreateToken(email)
	this.Ctx.SetCookie("token", token, 24*60*60, "/") // 设置cookie
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登录成功",
		"data": token,
	}
	this.ServeJSON()
	return
}
