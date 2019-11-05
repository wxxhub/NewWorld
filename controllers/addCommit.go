package controllers

import (
	models "NewWorld/models"

	"github.com/astaxie/beego"
)

type AddCommitCotroller struct {
	beego.Controller
}

// Get .
func (a *AddCommitCotroller) Get() {

}

// Post .
func (a *AddCommitCotroller) Post() {
	userID := a.GetString("user_id")
	messageID := a.GetString("message_id")
	commit := a.GetString("commit")

	addStatus := models.GetInstance().AddCommit(messageID, userID, commit)

	a.Ctx.ResponseWriter.WriteHeader(int(addStatus))
}
