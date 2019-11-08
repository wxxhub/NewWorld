package controllers

import (
	models "NewWorld/models"
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// GetSelfMessageController .
type GetSelfMessageController struct {
	beego.Controller
}

// Get .
func (g *GetSelfMessageController) Get() {

}

// Post .
func (g *GetSelfMessageController) Post() {
	userID := g.GetString("user_id")
	start, okStart := g.GetInt("start")
	end, okEnd := g.GetInt("end")
	messages := make([]models.Message, 0)
	if okStart == nil && okEnd == nil {
		list, ok := uniqueModel.GetMessages(userID, start, end)
		if ok == nil {
			for _, messageID := range list {
				message, findOk := uniqueModel.GetMessage(messageID)
				if findOk == true {
					message.MessageID = messageID
					messages = append(messages, message)
				}
			}
		}

		meshalData, _ := json.Marshal(messages)
		g.Data["json"] = string(meshalData)
		g.ServeJSON()
	} else {
		logs.Warn("GetSelfMessageController get sessoion or post failed!")
		g.Ctx.ResponseWriter.WriteHeader(0)
	}
}
