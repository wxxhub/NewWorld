package main

import (
	_ "NewWorld/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"fmt"
	"strings"
)

// FilterUser .
var FilterUser = func(ctx *context.Context) {

	ok := strings.Contains(ctx.Request.RequestURI, "/login")

	if !ok {
		name, nameOk := ctx.Input.Session("name").(string)
		code, codeOk := ctx.Input.Session("unique_code").(string)

		println(name)
		println(nameOk)
		println(code)
		println(codeOk)
		if !nameOk && !codeOk {
			ctx.Redirect(302, "/login")
		} else if name != ctx.GetCookie("name") && code != ctx.GetCookie("unique_code") {
			ctx.Redirect(302, "/login")
		}

		fmt.Println("pass!")
	}

	fmt.Println("login...!")
}

func main() {
	//注册过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
