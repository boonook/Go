package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myAppApi/models"
)

// Operations about Users
type MenuController struct {
	beego.Controller
}

func (u *MenuController) GetAll() {
	fmt.Println("89089089089089089")
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()

}
