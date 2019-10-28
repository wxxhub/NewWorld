package controllers

import (
	"github.com/astaxie/beego"
)

// MainController .
type MainController struct {
	beego.Controller
}

// Get .
func (m *MainController) Get() {
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "astaxie@gmail.com"
	m.TplName = "index.html"
}