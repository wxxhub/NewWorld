package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

// AddMessageController .
type AddMessageController struct {
	beego.Controller
}

// Get .
func (a *AddMessageController) Get() {

}

// Post .
func (a *AddMessageController) Post() {
	userID := a.GetString("user_id")
	text := a.GetString("text")
	image := a.GetString("image")

	addStatus := models.GetInstance().AddMessage(userID, text, image)

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
