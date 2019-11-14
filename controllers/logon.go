package controllers

import (
	models "NewWorld/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// LogonController .
type LogonController struct {
	beego.Controller
}

// Get .
func (l *LogonController) Get() {
	logs.Info("LogonController Get()")
	l.TplName = "logon.html"
}

// Post .
func (l *LogonController) Post() {
	logs.Info("user logon!")
	name := l.GetString("name")
	ID := l.GetString("user_id")
	pwd := l.GetString("pwd")
	fmt.Println(name, ID, pwd)
	image := l.GetString("image")

	addStatus := models.GetInstance().AddUser(ID, name, pwd, image)

	fmt.Println(addStatus)
	if addStatus == models.HAVEEXIST {
		l.Ctx.ResponseWriter.WriteHeader(204)
	} else {
		l.Ctx.ResponseWriter.WriteHeader(200)
	}
	l.ServeJSON()
}
