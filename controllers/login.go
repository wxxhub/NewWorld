package controllers

import (
	models "NewWorld/models"
	"strconv"

	"github.com/astaxie/beego"

	"time"
)

// LoginController .
type LoginController struct {
	beego.Controller
}

var uniqueModel = models.GetInstance()

// Get .
func (l *LoginController) Get() {
	l.TplName = "test/login_test.html"
}

// Post .
func (l *LoginController) Post() {
	name := l.GetString("name")
	pwd := l.GetString("pwd")

	ok := uniqueModel.AuthenticateUser(name, pwd)

	if ok {
		uniqueCode := getUniqueCode()
		println("uniqueCode:" + uniqueCode)
		l.SetSession("name", name)
		l.SetSession("unique_code", uniqueCode)
		l.Ctx.SetCookie("name", name)
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
