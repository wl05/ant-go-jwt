package controllers

import (
	"ant-go-jwt/common/consts"
	"ant-go-jwt/common/utils"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	accessToken := base.Ctx.GetCookie("accessToken")
	refreshToken := base.Ctx.GetCookie("refreshToken")

	isValidAccessToken, _, _ := utils.CheckToken(accessToken)
	isValidRefreshToken, _, refreshTokenClaims := utils.CheckToken(refreshToken)

	// 先判断refreshToken是否已经过期，过期直接返回让客户端重新登录
	if !isValidRefreshToken {
		base.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_LOGIN_ERROR,
			"msg":  consts.ERROR_DES_LOGIN_ERROR,
		}
		base.ServeJSON()
		return
	}
	// 再判断redis中的黑名单里面是否有isValidRefreshToken如果有说明客户端已经登出，解决jwt登出问题
	val, err := utils.RClient.Get(string(refreshToken)).Result()
	if err != redis.Nil && err != nil {
		base.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		base.ServeJSON()
		return
	}

	if val == "exited" {
		base.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_LOGIN_ERROR,
			"msg":  consts.ERROR_DES_LOGIN_ERROR,
		}
		base.ServeJSON()
		return
	}
	// refreshToken有效但是accessToken过期，可以重新生成新的accessToken让客户端重新发起请求,解决token刷新问题
	if !isValidAccessToken {
		t, _ := utils.CreateToken(refreshTokenClaims.Email, time.Now().Add(1*time.Minute))
		base.Ctx.SetCookie("accessToken", accessToken, "/")   // 设置cookie
		base.Ctx.SetCookie("refreshToken", refreshToken, "/") // 设置cookie
		type Data struct {
			AccessToken  string `json:"accessToken"`
			RefreshToken string `json:"refreshToken"`
		}
		data := Data{t, refreshToken}
		base.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REFRESH_TOKEN,
			"msg":  consts.ERROR_DES_REFRESH_TOKEN,
			"data": data,
		}
		base.ServeJSON()
		return
	}
}
