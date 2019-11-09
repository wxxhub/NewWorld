package controllers

import (
	models "NewWorld/models"
	"fmt"

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
	userID := a.GetSession("user_id").(string)
	messageID := a.GetString("message_id")
	havePraise, _ := a.GetBool("praise")

	fmt.Println(userID, messageID)
	var addStatus models.ProcessStatus
	if havePraise == true {
		addStatus = uniqueModel.AddPraise(messageID, userID)
	} else {
		addStatus = uniqueModel.CancelPraise(messageID, userID)
	}

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
