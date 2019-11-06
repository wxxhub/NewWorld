package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

// AddPraiseController .
type AddPraiseController struct {
	beego.Controller
}

// Get .
func (a *AddPraiseController) Get() {

}

// Post .
func (a *AddPraiseController) Post() {
	userID := a.GetString("user_id")
	messageID := a.GetString("message_id")
	havePraise, _ := a.GetBool("praise")
	var addStatus models.AddStatus
	if havePraise == true {
		addStatus = uniqueModel.AddPraise(messageID, userID)
	} else {
		addStatus = uniqueModel.CancelPraise(messageID, userID)
	}

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
