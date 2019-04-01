package server

import (
	"encoding/json"
	"http_post"
	"mysql"
)
//Method string,Params string
func GetNewAddress(password string){
	id :="1"
	method := "getnewaddress"
	params :=[]interface{}{password}

	//请求返回的数据为[]byte格式
	body :=http_post.HttpPost(id,method,params)

	/*查看返回原始数据
	body_data := string(body)
	fmt.Println(body_data)
	*/

	type address_result struct {
		Result string `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}

	//json转struct
	var bode_info address_result
	json.Unmarshal(body,&bode_info)
	NewAddress :=bode_info.Result

	//存入数据库
	if bode_info.Error == "" {
		mysql.InitDB(method,NewAddress)
	}


}
