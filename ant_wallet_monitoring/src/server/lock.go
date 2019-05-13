package server

import (
	"config"
	"encoding/json"
	"fmt"
)

//导入配置文件
var configMap = conf.InitConfig("src/config/conf.txt")
//获取配置value
var fromAddress=configMap["fromAddress"]
var passWord=configMap["passWord"]


func UnLock(){
	type unlock struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}



	PostParams :=[]interface{}{fromAddress,passWord,60}

	PostData:=unlock{
		Jsonrpc:"2.0",
		Method:"personal_unlockAccount",
		Params:PostParams,
		ID:3,
	}
	PostJsons, errs := json.Marshal(PostData)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("解锁请求数据为：",string(PostJsons))

	body :=HttpPost(string(PostJsons))

	type getunlock_result struct {
		Result string `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}

	//json转struct
	var GetUnlock getunlock_result
	json.Unmarshal(body,&GetUnlock)
	GetUnlock_status :=GetUnlock.Result

	fmt.Println("解锁返回数据为：",string(GetUnlock_status))
}

func Lock() {

	type lock struct {
		Jsonrpc string   `json:"jsonrpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		ID      int      `json:"id"`
	}
	PostParams :=[]string{fromAddress}

	PostData := lock{
			Jsonrpc:"2.0",
			Method:"personal_lockAccount",
			Params:PostParams,
		    ID:5,
	}


	PostJsons, errs := json.Marshal(PostData)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("加锁请求数据为",string(PostJsons))


	body :=HttpPost(string(PostJsons))

	type getlock_result struct {
		Result string `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}

	//json转struct
	var Getlock getlock_result
	json.Unmarshal(body,&Getlock)
	Getlock_status :=Getlock.Result

	fmt.Println("解锁返回数据为",string(Getlock_status))


}
