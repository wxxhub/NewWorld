package routers

import (
	"NewWorld/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})
}
