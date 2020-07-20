package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "bagulm123@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Get() {
	c.Data["Name"] = "Mahendra Hiraman Bagul"
	c.Data["MobileNo"] = "+918484947814"
	c.TplName = "aboutme.tpl"
}
