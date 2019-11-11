package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	models "NewWorld/models"
)

type GetTestMessageController struct {
	beego.Controller
}

// Get .
func (g *GetTestMessageController) Get() {
	g.TplName = "test/api_test.html"
}

// Post .
func (g *GetTestMessageController) Post() {
	testData := models.GetTestData()
	datas, _ := json.Marshal(testData)
	g.Data["json"] = string(datas)
	g.ServeJSON()
}
