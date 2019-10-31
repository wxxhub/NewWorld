package controllers

import (
	models "NewWorld/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// LoginController .
type LoginController struct {
	beego.Controller
}

var uniqueModel = models.GetInstance(1)

// Get .
func (l *LoginController) Get() {
	// l.Ctx.WriteString("Yes, you find it.")
	l.TplName = "test/login_test.html"
}

// Post .
func (l *LoginController) Post() {
	name := l.GetString("name")
	pwd := l.GetString("pwd")

	ok := uniqueModel.AuthenticateUser(name, pwd)

	var result string
	if ok {
		uniqueCode := getUniqueCode()
		l.SetSession("name", name)
		l.SetSession("unique_code", uniqueCode)
		l.Ctx.SetCookie("name", name)
		l.Ctx.SetCookie("unique_code", uniqueCode)
		result = "ok"
	} else {
		result = "false"
	}

	data, _ := json.Marshal(result)
	l.Data["json"] = string(data)
	l.ServeJSON()
}

// getUniqueCode .
func getUniqueCode() (code string) {
	return "123"
}
