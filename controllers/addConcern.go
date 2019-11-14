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

	var code int
	if addStatus == models.FAILED {
		code = 214
	} else if addStatus == models.HAVEEXIST {
		code = 211
	} else {
		code = 200
	}
	// a.Ctx.ResponseWriter.Header().Set("Content-type", "application/text")
	a.Ctx.ResponseWriter.WriteHeader(code)
	// a.Ctx.ResponseWriter.Write([]byte(""))
}
