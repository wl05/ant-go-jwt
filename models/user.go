package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `orm:"null" json:"salt"`
	Email    string `json:"email"`
}
type GetUsersStruct struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 自定义表名
func (m *User) TableName() string {
	return "user"
}

func Register(u User) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into user (username,password,email) values (?,?,?)", u.Username, u.Password, u.Email).Exec()
	return err
}

func GetUsers() (user []GetUsersStruct, err error) {
	var users []GetUsersStruct
	o := orm.NewOrm()
	_, _err := o.Raw("select id,username,email from user").QueryRows(&users)
	return users, _err
}

func GetUserById(id int) (user User, err error) {
	var _user User
	o := orm.NewOrm()
	act := o.Raw("select * from user where id = ?", id)
	_err := act.QueryRow(&_user)
	return _user, _err
}

func GetUserByUsername(username string) (user User, err error) {
	var _user User
	o := orm.NewOrm()
	o.Using("default")
	act := o.Raw("select * from user where username = ?", username)
	_err := act.QueryRow(&_user)
	return _user, _err
}

func GetUserByEmail(email string) (user User, err error) {
	var _user User
	o := orm.NewOrm()
	o.Using("default")
	act := o.Raw("select * from user where email = ?", email)
	_err := act.QueryRow(&_user)
	return _user, _err
}
func UpdateTagByUserId(id int, username string, password string) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("update user set username = ?,password = ? where id = ?", username, password, id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}
