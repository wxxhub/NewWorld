package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

type AddConcernController struct {
	beego.Controller
}

// Get .
func (a *AddConcernController) Get() {

}

// Post .
func (a *AddConcernController) Post() {
	userID := a.GetString("user_id")
	goalUserId := a.GetString("goal_user_id")

	addStatus := models.GetInstance().AddConcern(userID, goalUserId)

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
