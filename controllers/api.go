package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	model "NewWorld/models"
)

// APIController .
type APIController struct {
	beego.Controller
}

// Get .
func (a *APIController) Get() {
	a.TplName = "test/api_test.html"
}

// Post .
func (a *APIController) Post() {
	name := a.GetString("name")

	fmt.Printf("I'm get [%s]", name)
	switch name {
	case "test_data":
		a.returnTestData()
	default:
		a.returnFalse()
	}
}

// return TestData .
func (a *APIController) returnTestData() {
	testData := model.GetTestData()
	datas, _ := json.Marshal(testData)
	a.Data["json"] = string(datas)
	a.ServeJSON()
}

// return false .
func (a *APIController) returnFalse() {
	data, _ := json.Marshal("False")
	a.Data["json"] = string(data)
	a.ServeJSON()
}
