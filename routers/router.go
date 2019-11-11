package routers

import (
	"NewWorld/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logon", &controllers.LogonController{})
	beego.Router("/commit", &controllers.AddCommitCotroller{})
	beego.Router("/praise", &controllers.AddPraiseController{})
	beego.Router("/add_message", &controllers.AddMessageController{})
	beego.Router("/concern", &controllers.AddConcernController{})
	beego.Router("/self_message", &controllers.GetSelfMessageController{})
	beego.Router("/concern_message", &controllers.GetConcernMessageController{})
	beego.Router("/hot_message", &controllers.GetHotMessageController{})
	// beego.Router("/api", &controllers.APIController{})
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/test_message", &controllers.GetTestMessageController{})
}
