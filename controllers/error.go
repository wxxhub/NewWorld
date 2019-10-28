package controllers

import "github.com/astaxie/beego"

/*
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

// Error401 .
func (e *ErrorController) Error401() {
	e.Data["content"] = "未经授权，请求要求验证身份"
	e.TplName = "error/401.html"
}

// Error403 .
func (e *ErrorController) Error403() {
	e.Data["content"] = "服务器拒绝请求"
	e.TplName = "error/403.html"
}

// Error404 .
func (e *ErrorController) Error404() {
	e.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	e.TplName = "error/404.html"
}

// Error500 .
func (e *ErrorController) Error500() {
	e.Data["content"] = "server error"
	e.TplName = "error/500.html"
}

// Error503 .
func (e *ErrorController) Error503() {
	e.Data["content"] = "服务器目前无法使用（由于超载或停机维护）"
	e.TplName = "error/503.html"
}
