package controllers

import (
	"myAppApi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type MenuController struct {
	beego.Controller
}

type JsonReturn2 struct {
	Msg  string      `json:"message"`
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

func (u *MenuController) GetAll() {
	var JsonReturn JsonReturn2
	users := models.GetAllUsers()
	JsonReturn.Msg = "操作成功"
	JsonReturn.Code = 200
	JsonReturn.Data = users
	u.Data["json"] = JsonReturn
	u.ServeJSON()
	return
}

func (u *MenuController) Get() {
	uid := u.GetString(":uid")

	if uid != "" {
		var JsonReturn JsonReturn2
		users := models.GetAllUsers()
		JsonReturn.Msg = "操作成功2"
		JsonReturn.Code = 200
		JsonReturn.Data = users
		u.Data["json"] = JsonReturn
		u.ServeJSON()
		return
	}
}
