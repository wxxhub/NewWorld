package controllers

import (
	"github.com/astaxie/beego"
)

// LoginController .
type LoginController struct {
	beego.Controller
}

// Post .
func (l *LoginController) Post() {
	name := l.GetString("name")
	pwd := l.GetString("pwd")

	l.Ctx.WriteString("name:" + name + ", pwd:" + pwd)
}
