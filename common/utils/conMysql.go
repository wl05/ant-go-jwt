package utils

import "github.com/astaxie/beego/orm"

// 连接mysql
func MysqlClient() {
	// 自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/ant-go-jwt?charset=utf8")
	//err := orm.RunSyncdb("default", true, false)
	// 遇到错误立即返回
	//if err != nil {
	//	fmt.Println(err)
	//}
}
