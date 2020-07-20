package routers

import (
	"beego-mvc-example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/aboutme", &controllers.UserController{})
}
