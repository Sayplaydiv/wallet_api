package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)


type DepwitStatus struct {
	Id int `json:"id"`
	Asset string `json:"asset"`
	User_id int `json:"user_id"`
	Amount float64 `json:"amount"`
	Tx_hash string `json:"tx_hash"`
	Block_height int `json:"block_height"`
	Tx_time time.Time `json:"tx_time"`
	Status int `json:"status"`
}


func GetAllDeposits() []*DepwitStatus {
	o := orm.NewOrm()
	o.Using("default")
	var deposits []*DepwitStatus
	q:= o.QueryTable("depwit_status")
	q.All(&deposits)
	return deposits
}



func AddDeposit(object DepwitStatus) (DepwitStatusId string) {
	object.Id = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.Id] = &object
	return object.Id
}




func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(DepwitStatus))
}