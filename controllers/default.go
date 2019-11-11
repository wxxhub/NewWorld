package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

var uniqueModel = models.GetInstance()

// MainController .
type MainController struct {
	beego.Controller
}

// Get .
func (m *MainController) Get() {
	m.TplName = "index.html"
}
