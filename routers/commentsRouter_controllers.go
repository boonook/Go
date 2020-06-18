package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:UserController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myAppApi/controllers:MenuController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:MenuController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/getAll`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["myAppApi/controllers:MenuController"] = append(beego.GlobalControllerRouter["myAppApi/controllers:MenuController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
