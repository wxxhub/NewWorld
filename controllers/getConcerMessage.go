package controllers

import (
	"encoding/json"

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

	// 获取关注的人
	concerns, _ := uniqueModel.GetConcern(userID)
	// fmt.Println("concerns:", concerns)

	// 根据关注的人获取消息
	messages := uniqueModel.GetConcernMessage(concerns, size)
	// fmt.Println("messages:", messages)
	marshal, _ := json.Marshal(messages)
	// fmt.Println("json:", string(marshal))
	g.Data["json"] = string(marshal)
	g.ServeJSON()
}
