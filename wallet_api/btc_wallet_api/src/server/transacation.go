package server

import (
	"encoding/json"
	"fmt"
	"http_post"
)

func Transaction(from_address string,to_address string,Value float64){

	var minconf int
	var maxconf int
	type minimumAmount_0 struct {
		MinimumAmount float64 `json:"minimumAmount"`
	}
	minimumAmount := minimumAmount_0{
		MinimumAmount: Value,
	}
	unspen_address := []string{from_address}
	minconf = 1
	maxconf = 9999999

	id := "2"
	method := "listunspent"
	params := []interface{}{minconf, maxconf, unspen_address, true, minimumAmount}

	//请求返回的数据为[]byte格式
	body := http_post.HttpPost(id, method, params)

	body_data := string(body)
	fmt.Println(body_data)


	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)

			Result := dat["result"].([]interface{})



			fmt.Println()
			//取出对应金额的utxo的未花费的输出
			for i,data :=range Result{
				fmt.Println(i,data)
				utxo_data :=data.(map[string]interface{})
				utxo_txid :=utxo_data["txid"].(string)
				utxo_vout :=utxo_data["vout"].(float64)
				utxo_redeemScript :=utxo_data["redeemScript"].(string)
				utxo_scriptPubKey :=utxo_data["scriptPubKey"].(string)


				//创建交易
				if  data != nil && i==0 {
					fmt.Println(utxo_txid, utxo_vout, utxo_redeemScript, utxo_scriptPubKey)

				//	{"jsonrpc": "1.0", "id":"curltest", "method": "createrawtransaction", "params": ["[{\"txid":"myid","vout":0}]", "{"address":0.01}"] }

				
					//id :="3"
					//method := "createrawtransaction"

					type Unspent struct {
						Txid string `json:"txid"`
						Vout float64   `json:"vout"`
						RedeemScript string  `json:"redeemScript"`
						ScriptPubKey string  `json:"scriptPubKey"`
					}
					PostData_Unspent := Unspent{
						Txid:utxo_txid,
						Vout:utxo_vout,
						RedeemScript:utxo_redeemScript,
						ScriptPubKey:utxo_scriptPubKey,
					}

					PostData_Unspent_0,err :=json.Marshal(PostData_Unspent)
					if err !=nil{
						fmt.Println("error",err)
					}

					//utxo未花费的交易输出
					PostData_Unspent_json_1 := string(PostData_Unspent_0)
					 PostData_Unspent_json  :=  []string{PostData_Unspent_json_1}
					fmt.Println(PostData_Unspent_json)

					type unspent_toaddress struct {

					}




					fmt.Println("======",unspent_toaddress_0,"======")



					/*

					params :=[]interface{}{password}
					//请求返回的数据为[]byte格式
					body :=http_post.HttpPost(id,method,params)

					*/

					//fmt.Println("------------------",PostData_Unspent_json,"---------------")

				}
			}

		}






	}