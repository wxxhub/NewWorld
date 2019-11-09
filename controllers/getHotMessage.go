package controllers

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

// GetHotMessageController .
type GetHotMessageController struct {
	beego.Controller
}

// Get .
func (g *GetHotMessageController) Get() {

}

// Post .
func (g *GetHotMessageController) Post() {
	messages := uniqueModel.GetHotMessage("")
	meshalData, _ := json.Marshal(messages)
	g.Data["json"] = string(meshalData)
	g.ServeJSON()
}

// RandGroup 生成随机数组.
func RandGroup(start, end, size int) ([]int, bool) {
	result := make([]int, size)
	if (end-start) > size || start >= end {
		return result, false
	}

	num := end - start

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		result[i] = start + rand.Intn(num)
	}

	return result, true
}
