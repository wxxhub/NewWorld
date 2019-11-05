package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

// LogonController .
type LogonController struct {
	beego.Controller
}

// Get .
func (l *LogonController) Get() {

}

// Post .
func (l *LogonController) Post() {
	name := l.GetString("name")
	ID := l.GetString("user_id")
	pwd := l.GetString("pwd")

	addStatus := models.GetInstance().AddUser(ID, name, pwd)

	l.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
