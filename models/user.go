package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int
	Username string
	Password string
	Salt     string `orm:"null"`
	Email    string
}

// 自定义表名
func (m *User) TableName() string {
	return "user"
}

func Register(u User) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into user (username,password,email) values (?,?,?)", u.Username, u.Password, u.Email).Exec()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
	return err
}

func GetUsers() (user []User, err error) {
	var users []User
	o := orm.NewOrm()
	_, _err := o.QueryTable("user").All(&users, "Id", "Username")
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
	act := o.Raw("SELECT * FROM user WHERE username = ?", username)
	_err := act.QueryRow(&_user)
	return _user, _err
}

func GetUserByEmail(email string) (user User, err error) {
	var _user User
	o := orm.NewOrm()
	o.Using("default")
	act := o.Raw("SELECT * FROM user WHERE email = ?", email)
	_err := act.QueryRow(&_user)
	return _user, _err
}
func UpdateTagByUserId(id int, username string, password string) (num int64, err error) {
	DB := orm.NewOrm()
	var act orm.RawSeter
	DB.Using("default")
	act = DB.Raw("UPDATE user SET username = ?,password = ? where id = ?", username, password, id)
	res, _err := act.Exec()
	_num, _ := res.RowsAffected()
	return _num, _err
}
