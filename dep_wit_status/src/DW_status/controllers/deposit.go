package controllers

import (
	"DW_status/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type DepositController struct {
	beego.Controller
}
// @Title 获得所有
// @Description 返回所有的数据
// @Success 200 {object}
// @router / [get]
func (u *DepositController) GetAll() {
	deposits := models.GetAllDeposits()
	u.Data["json"] = map[string]interface{}{"deposits": deposits}
	u.ServeJSON()
}



func (u *DepositController) Post() {
	var s models.DepwitStatus
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddDeposit(&s)
	u.Data["json"] = uid
	u.ServeJSON()
}


