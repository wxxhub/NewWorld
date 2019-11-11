package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

// AddConcernController .
type AddConcernController struct {
	beego.Controller
}

// Get .
func (a *AddConcernController) Get() {

}

// Post .
func (a *AddConcernController) Post() {
	userID := a.GetSession("user_id").(string)
	goalUserID := a.GetString("goal_user_id")
	concern, _ := a.GetBool("concern")

	var addStatus models.ProcessStatus
	if concern {
		addStatus = uniqueModel.AddConcern(userID, goalUserID)
	} else {
		addStatus = uniqueModel.CancelConcern(userID, goalUserID)
	}

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
