package controllers

import (
	"encoding/json"
	"fmt"
	"myAppApi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type JsonReturn struct {
	Msg  string      `json:"message"`
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	var JsonReturn JsonReturn
	JsonReturn.Msg = "操作成功"
	JsonReturn.Code = 200
	JsonReturn.Data = users
	u.Data["json"] = JsonReturn
	u.ServeJSON()
	return
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	fmt.Println("uid----------------------------------", uid)
	if uid != "" {
		// users := models.GetAllUsers()
		// var JsonReturn JsonReturn
		// JsonReturn.Msg = "操作成功2"
		// JsonReturn.Code = 200
		// JsonReturn.Data = users
		type List struct {
			Name  string
			Age   int
			count int
		}

		type Data struct {
			page  string
			size  int
			count int
			list  []List
		}

		type Ret struct {
			Code int
			Msg  string
			Data interface{} `json:"data"` //	Data interface{} `json:"data"`//Data []Data
		}

		ret := new(Ret)
		ret.Code = 200
		ret.Msg = "success"
		list2 := []List{
			{Name: "a", Age: 1, count: 90},
			{Name: "b", Age: 2, count: 80},
			{Name: "c", Age: 3, count: 70},
		}
		data := Data{page: "why", size: 18, count: 31, list: list2}
		ret.Data = data
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
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
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
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
func (u *UserController) Login() {
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
func (u *UserController) Logout() {
	var JsonReturn JsonReturn
	JsonReturn.Msg = "操作成功"
	JsonReturn.Code = 200
	u.Data["json"] = JsonReturn
	u.ServeJSON()
}
