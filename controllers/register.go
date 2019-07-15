package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// Operations about Register
type RegisterController struct {
	beego.Controller
}

// @Title Register
// @Description Register
// @Param   username    formData    string  "ant"  true    "用户名"
// @Param   password    formData    string  "123"    true    "密码"
// @Param   email    formData    string  "demo@qq.com"    true    "邮箱"
// @Success 1000   参数错误
// @Success 1001   请求出错
// @Success 2000 用户名已存在
// @Success 2005 邮箱已存在
// @router / [post]
func (this *RegisterController) Register() {
	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	email := strings.TrimSpace(this.GetString("email"))
	// 判断参数是否为空
	if username == "" || password == "" || email == "" {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_PARAMETER_ILLEGAL,
			"msg":  consts.ERROR_DES_PARAMETER_ILLEGAL,
		}
		this.ServeJSON()
		return
	}

	// 判断用户名是否已经存在
	_, err := models.GetUserByUsername(username)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_USER_EXIST,
			"msg":  consts.ERROR_DES_USER_EXIST,
		}
		this.ServeJSON()
		return
	}

	// 判断邮箱是否已经存在
	_, err = models.GetUserByEmail(email)
	if err != orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_EMAIL_EXIST,
			"msg":  consts.ERROR_DES_EMAIL_EXIST,
		}
		this.ServeJSON()
		return
	}
	// 创建账户
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 加密密码
	encodePW := string(hash)                                                       // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	user := models.User{}
	user.Username = username
	user.Password = encodePW
	user.Email = email
	err = models.Register(user)
	if err == nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.SUCCECC,
			"msg":  "创建成功",
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
	}
}
