package server

import (
	"config"
	"encoding/json"
	"fmt"
	"strconv"
)

//{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xc94770007dda54cF92009BFF0dE90c06F603a09f", "latest"],"id":1}
func GetBalance() (int64) {
	type getbalance struct {
		Jsonrpc string   `json:"jsonrpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		ID      int      `json:"id"`
	}
	//导入配置文件
	configMap := conf.InitConfig("src/config/conf.txt")
	//获取配置value
	fromAddress:=configMap["fromAddress"]


	PostParams :=[]string{fromAddress,"latest"}

	getBalance := getbalance{
		Jsonrpc:"2.0",
		Method:"eth_getBalance",
		Params:PostParams,
		ID:1,
	}
	PostData, errs := json.Marshal(getBalance)
	if errs !=nil {
			fmt.Println(errs.Error())
	}
	body :=HttpPost(string(PostData))



	type getbalance_result struct {
		Result string `json:"result"`
		Error string  `error:"result"`
		Id string     `id:"result"`
	}

	//json转struct
	var GetBlance getbalance_result
	json.Unmarshal(body,&GetBlance)
	balance :=GetBlance.Result

	fmt.Println(string(balance))

	balance_val_0 := balance[2:]

	n, err := strconv.ParseUint(balance_val_0, 16, 64)
	if err != nil {
		panic(err)
	}

	balance_return :=int64(n)
	//转化为eth
	balance_val_1 := strconv.FormatInt((int64(n)/1000000000000000000),10)
	balance_val_2 := strconv.FormatInt((int64(n)%1000000000000000000),10)
	balance_val := balance_val_1 + "."+balance_val_2
	fmt.Print("归集账户目前金额为：",balance_val,"个ETH")
	return balance_return

}