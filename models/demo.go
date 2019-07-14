package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Demo map[string]*Demos
)

func init() {
	Demo = make(map[string]*Demos)
	u := Demos{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	Demo["user_11111"] = &u
}

type Demos struct {
	Id       string
	Demosname string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddDemos(u Demos) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Demo[u.Id] = &u
	return u.Id
}

func GetDemos(uid string) (u *Demos, err error) {
	if u, ok := Demo[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Demos not exists")
}

func GetAllDemoss() map[string]*Demos {
	return Demo
}

func UpdateDemos(uid string, uu *Demos) (a *Demos, err error) {
	if u, ok := Demo[uid]; ok {
		if uu.Demosname != "" {
			u.Demosname = uu.Demosname
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("Demos Not Exist")
}

func Login(username, password string) bool {
	for _, u := range Demo {
		if u.Demosname == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteDemos(uid string) {
	delete(Demo, uid)
}
