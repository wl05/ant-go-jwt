package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/common/utils"
	"ant-go-jwt/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"github.com/dgrijalva/jwt-go"

)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Register
// @Description Register
// @Param   username    formData    string  "ant"  true    "用户名"
// @Param   password    formData    string  "123"    true    "密码"
// @Param   email    formData    string  "demo@qq.com"    true    "邮箱"
// @Success 200  请求成功
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Register() {
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	user := models.User{}
	user.Username = username
	user.Password = password
	user.Email = email

	err := models.Register(user)
	if err == nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.SUCCECC,
			"msg":  "创建成功",
		}
		this.ServeJSON()
	}
}

// @Title 用户登录
// @Description 用户登录
// @Param   email    formData    string  "ant"  true  "邮箱"
// @Param   password    formData    string  "123"    true "密码"
// @Success 200 登录成功
// @Success 1101   参数错误
// @Success 1102   请求出错
// @Success 1104 用户名不存在
// @Success 1105 用户名或者密码错误
// @router / [post]
func (this *UserController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")
	fmt.Println("===========")
	fmt.Println(reflect.TypeOf(this))

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
	if user.Password != utils.Crypto(password) {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USERNAME_OR_PASSWORD_ERROR,
			"msg":  consts.ERROR_DES_USERNAME_OR_PASSWORD_ERROR,
		}
		this.ServeJSON()
		return
	}
	//ip := this.Ctx.Input.IP() + ":" + strconv.Itoa(this.Ctx.Input.Port())
	//token, err := models.CreateToken(user.Id, ip, this.Ctx.Request.UserAgent())
	//this.Ctx.SetCookie("token", token, 24*60*60, "/") // 设置cookie
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登录成功",
		"data": "",
	}
	this.ServeJSON()
	return
}
