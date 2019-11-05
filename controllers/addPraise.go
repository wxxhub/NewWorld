package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

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

	addStatus := models.GetInstance().AddPraise(messageID, userID)

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
