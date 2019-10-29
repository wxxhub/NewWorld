package routers

import (
	"NewWorld/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/api", &controllers.APIController{})
	beego.ErrorController(&controllers.ErrorController{})
}
