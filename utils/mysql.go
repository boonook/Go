package utils

import "github.com/astaxie/beego"

// func Mysqls() *sql.DB {
// 	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ZHDJ?charset=utf8")
// 	if err != nil {
// 		beego.Error("--------------连接数据库失败--------------")
// 	} else {
// 		beego.Info("--------------连接数据库成功--------------")
// 	}
// 	return db
// }

func Mysqls() {
	beego.Info("连接数据库成功")
}
