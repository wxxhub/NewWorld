package main

import (
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/astaxie/beego/toolbox"

	models "NewWorld/models"
	_ "NewWorld/routers"
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
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	// 设置log打印级别
	logs.SetLevel(logs.LevelWarning)
	beego.BConfig.WebConfig.Session.SessionOn = true
	models.HotManager.Init()
	updateHotTask := toolbox.NewTask("updateHot", "0/30 * * * * *", func() error { models.HotManager.Update(); return nil })
	toolbox.AddTask("updateHot", updateHotTask)
	toolbox.StartTask()
	// go beego.Run()
	beego.Run()
}
