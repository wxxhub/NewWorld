package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// GetConcernMessageController .
type GetConcernMessageController struct {
	beego.Controller
}

// Get .
func (g *GetConcernMessageController) Get() {

}

// Post .
func (g *GetConcernMessageController) Post() {
	userID := g.GetSession("user_id").(string)
	size, _ := g.GetUint64("size")

	concerns, _ := uniqueModel.GetConcern(userID)
	messages := uniqueModel.GetConcernMessage(concerns, size)

	marshal, _ := json.Marshal(messages)

	g.Data["json"] = string(marshal)
	g.ServeJSON()
}
