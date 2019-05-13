package server

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GasPrice()(uint64,string){
	type unlock struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}



	PostParams :=[]interface{}{}

	PostData:=unlock{
		Jsonrpc:"2.0",
		Method:"eth_gasPrice",
		Params:PostParams,
		ID:2,
	}
	PostJsons, errs := json.Marshal(PostData)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("获取gas请求数据为：",string(PostJsons))

	 body:=HttpPost(string(PostJsons))

	type gasprice_result struct {
		Result string `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}
	//json转struct
	var GetGasPrice gasprice_result
	json.Unmarshal(body,&GetGasPrice)
	GasPrice :=GetGasPrice.Result

	GasPrice_val_0 := GasPrice[2:]

	n, err := strconv.ParseUint(GasPrice_val_0, 16, 64)
	if err != nil {
		panic(err)
	}
	GasPrice_val :=int64(n)
		//fmt.Println(reflect.TypeOf(n))

	fmt.Println("本次交易gasPrice为：",GasPrice_val)

	return n,string(GasPrice)

}

