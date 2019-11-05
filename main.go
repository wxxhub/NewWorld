package main

import (
	_ "NewWorld/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"strings"
)

// FilterUser .
var FilterUser = func(ctx *context.Context) {

	ok := strings.Contains(ctx.Request.RequestURI, "/login")
	ok2 := strings.Contains(ctx.Request.RequestURI, "/test_message")
	_, nameOk := ctx.Input.Session("name").(string)

	if !ok && !nameOk && !ok2 {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Redirect(302, "/login")
	}

}

func main() {
	//注册过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
