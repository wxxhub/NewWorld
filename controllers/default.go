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
	m.TplName = "index.html"
}
