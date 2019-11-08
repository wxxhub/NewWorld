package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"time"
)

// LoginController .
type LoginController struct {
	beego.Controller
}

// Get .
func (l *LoginController) Get() {
	l.TplName = "test/login_test.html"
}

// Post .
func (l *LoginController) Post() {
	userID := l.GetString("user_id")
	pwd := l.GetString("pwd")

	name, image, ok := uniqueModel.AuthenticateUser(userID, pwd)

	if ok {
		returnData := make(map[string]string, 2)
		returnData["name"] = name
		returnData["image"] = image
		narshalData, _ := json.Marshal(returnData)
		l.Data["json"] = string(narshalData)
		l.ServeJSON()

		uniqueCode := getUniqueCode()
		l.SetSession("userID", userID)
		l.SetSession("unique_code", uniqueCode)
		l.Ctx.SetCookie("userID", userID)
		l.Ctx.SetCookie("unique_code", uniqueCode)
		l.Ctx.Redirect(302, "/")
	} else {
		l.Ctx.ResponseWriter.WriteHeader(401)
	}
}

// getUniqueCode .
func getUniqueCode() (code string) {
	end := time.Now().Unix()
	head := time.Now().Nanosecond()
	data := strconv.FormatInt(int64(head), 32) + strconv.FormatInt(int64(end), 32)
	return data
}
