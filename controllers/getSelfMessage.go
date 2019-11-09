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
	userID := g.GetSession("user_id").(string)
	start, okStart := g.GetUint64("start")
	end, okEnd := g.GetUint64("end")
	messages := make([]models.Message, 0)
	if okStart == nil && okEnd == nil {
		list, ok := uniqueModel.GetMessages(userID, start, end)
		if ok == nil {
			for _, messageID := range list {
				message, findOk := uniqueModel.GetMessage(messageID)
				if findOk == true {
					message.MessageID = messageID
					message.HavePraise = uniqueModel.HavePraise(messageID, userID)
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
