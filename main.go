package main

// 加下划线表明只调用init方法
import (
	"fmt"
	_ "myAppApi/routers" 
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		mysqlType :=beego.AppConfig.String("mysqlType")
		fmt.Println(beego.AppConfig.String("mysqlUserName"))
		db, err := sql.Open(mysqlType, "root:1234@tcp(127.0.0.1:3306)/ZHDJ?charset=utf8")
		if err != nil {
			beego.Error("连接数据库出错", err)
			return
		}else{
			beego.Info("连接数据库成功")
		}
		rows, err := db.Query("select id,userName from user")
		type UserInfo struct {
			id int
			userName string
		}
		var u UserInfo
		for rows.Next() {
			err = rows.Scan(&u.id, &u.userName)
			fmt.Println(u)
		}
		// 更新数据
		_, err = db.Exec("update user set userName='pd' where id=17")

		fmt.Println("testing...")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		// 关闭数据库连接
		defer db.Close()
		////获取beego的配置信息
		fmt.Println("beego.AppConfig",beego.AppConfig.String("mysqlType"))
	}
	beego.Run()
}
