package server

import (
	"encoding/json"
	"fmt"
	"http_post"
)

func Transaction(from_address string,to_address string,Value float64)  {


	var minconf int
	var maxconf int
	type minimumAmount_0 struct {
		MinimumAmount float64 `json:"minimumAmount"`
	}
	minimumAmount := minimumAmount_0{
		MinimumAmount:Value,
	}
	unspen_address:=[]string{}
	minconf=1
	maxconf=9999999

	id :="2"
	method := "listunspent"
	params :=[]interface{}{minconf,maxconf,unspen_address,true,minimumAmount}


	//请求返回的数据为[]byte格式
	body := http_post.HttpPost(id,method,params)


	body_data := string(body)
	fmt.Println(body_data)


	type listunspent_result struct {
		Result []interface{} `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}

	//json转struct

	var bode_info_0 listunspent_result
	json.Unmarshal(body,&bode_info_0)
	 bode_info_1 :=bode_info_0.Result
	//test:=bode_info_1["txid"]

	var	bode_info_2 = bode_info_1.(map[string]interface{})

	fmt.Println(bode_info_1)



}
