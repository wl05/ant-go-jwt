package controllers

import (
	"ant-go-jwt/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Demoss
type DemoController struct {
	beego.Controller
}

// @Title CreateDemos
// @Description create users
// @Param	body		body 	models.Demos	true		"body for user content"
// @Success 200 {int} models.Demos.Id
// @Failure 403 body is empty
// @router / [post]
func (u *DemoController) Post() {
	var user models.Demos
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddDemos(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Demoss
// @Success 200 {object} models.Demos
// @router / [get]
func (u *DemoController) GetAll() {
	users := models.GetAllDemoss()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Demos
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *DemoController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetDemos(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Demos	true		"body for user content"
// @Success 200 {object} models.Demos
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *DemoController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.Demos
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateDemos(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *DemoController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteDemos(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *DemoController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *DemoController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
