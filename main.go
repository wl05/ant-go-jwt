package main

import (
	_ "ant-go-jwt/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/ant-go-jwt?charset=utf8")
	// 自动创建表 参数二为是否开启创建表   参数三是否更新表
	//err := orm.RunSyncdb("default", true, false)
	// 遇到错误立即返回
	//if err != nil {
	//	fmt.Println(err)
	//}
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
