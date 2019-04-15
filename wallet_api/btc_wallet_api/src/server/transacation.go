package server

import (
	"encoding/json"
	"fmt"
	"http_post"
)

func Transaction(from_address string,to_address string,Value float64,Privatekey_1 string,Privatekey_2 string){

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
	fmt.Println(method,"查看utxo未花费的输出结果如下：",body_data)


	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err == nil {
		//fmt.Println("==============json str 转map=======================")
		//fmt.Println(dat)

			Result := dat["result"].([]interface{})



			fmt.Println()
			//取出对应金额的utxo的未花费的输出
			for i,data :=range Result{

				utxo_data :=data.(map[string]interface{})
				utxo_txid :=utxo_data["txid"].(string)
				utxo_vout :=utxo_data["vout"].(float64)
				utxo_redeemScript :=utxo_data["redeemScript"].(string)
				utxo_scriptPubKey :=utxo_data["scriptPubKey"].(string)


				//创建交易
				if  data != nil && i==0 {
					//查看utxo信息是否获取到
					//fmt.Println(utxo_txid, utxo_vout, utxo_redeemScript, utxo_scriptPubKey)

				//	{"jsonrpc": "1.0", "id":"curltest", "method": "createrawtransaction", "params": ["[{"txid":"myid","vout":0}]", "{"address":0.01}"] }

				
					//id :="3"
					//method := "createrawtransaction"


					type Unspent  struct {
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




					//utxo未花费的交易输出
					//fmt.Println("++++++++++++++++++++++",PostData_Unspent,"++++++++++++++++")

					type unspen_toaddress_0 struct {
						To_address string
						Value float64
					}

					unspent_toaddress :=unspen_toaddress_0{
						to_address,
						Value,
					}



					//发送请求构建交易
					id_create := "3"
					method_create := "createrawtransaction"
					params_create :=[]interface{}{PostData_Unspent,unspent_toaddress}
					//请求返回的数据为[]byte格式
					body :=http_post.HttpPost(id_create,method_create,params_create)


					body_info :=string(body)

					fmt.Println(method_create,"创建交易返回结果如下：",body_info)




					//获取创建交易生成的hexstring
					type hexstring struct {
						Result string `json:"result"`
						Error string  `error:"result"`
						Id string     `id:"result"`
					}
					var bode_info hexstring
					json.Unmarshal(body,&bode_info)
					hexstring_0 :=bode_info.Result
					fmt.Println("获取创建交易生成的hexstring:",hexstring_0)


					//第一次签名：
					id_sign_1 := "4"
					method_sign_1 := "signrawtransaction"
					privatekey_1 :=Privatekey_1
					params_sign_1 :=[]interface{}{hexstring_0,PostData_Unspent,[]string{privatekey_1}}

					//请求返回的数据为[]byte格式
					body_sign_1 :=http_post.HttpPost(id_sign_1,method_sign_1,params_sign_1)
					bode_info_1 :=string(body_sign_1)
					fmt.Println(method_sign_1,"第一次签名结果如下",bode_info_1)


					//获取第一次签名的交易hash===hexstring_1
					var dat_sign_1 map[string]interface{}
					if err := json.Unmarshal([]byte(body_sign_1), &dat_sign_1); err == nil {
						//fmt.Println("==============json str 转map=======================")
						//fmt.Println(dat)
						Result_sign_1 := dat_sign_1["result"].(map[string]interface{})
						Result_sign_1_hexstring :=Result_sign_1["hex"].(string)


						fmt.Println("第一次签名映射：",Result_sign_1_hexstring)
						if Result_sign_1_hexstring != hexstring_0 &&Result_sign_1_hexstring!=""{

							//第二次签名：
							id_sign_2 := "5"
							method_sign_2 := "signrawtransaction"
							privatekey_2 :=Privatekey_2
							params_sign_2 :=[]interface{}{Result_sign_1_hexstring,PostData_Unspent,[]string{privatekey_2}}

							//请求返回的数据为[]byte格式
							body_sign_2 :=http_post.HttpPost(id_sign_2,method_sign_2,params_sign_2)

							bode_info_2 :=string(body_sign_2)
							fmt.Println(method_sign_1,"第二次签名结果如下",bode_info_2)

							//获取第一次签名的交易hash===hexstring_2
							var dat_sign_2 map[string]interface{}
							if err := json.Unmarshal([]byte(body_sign_2), &dat_sign_2); err == nil {
								Result_sign_2 := dat_sign_2["result"].(map[string]interface{})
								Result_sign_2_hexstring :=Result_sign_2["hex"].(string)
								Result_sign_2_complete :=Result_sign_2["complete"].(bool)
								fmt.Println("第二次签名映射：",Result_sign_2_complete,Result_sign_2_hexstring)

								if Result_sign_2_complete == true{
											fmt.Println("成功了")

									//提交交易：
									id_send := "6"
									method_send := "sendrawtransaction"
									params_send :=[]interface{}{Result_sign_2_hexstring}
									body_send :=http_post.HttpPost(id_send,method_send,params_send)
									body_send_info :=string(body_send)
									fmt.Println("提交交易返回数据为",body_send_info)

									//获取交易hash
									type send_tx_hash struct {
										Result string
										Error  string
										Id	string

									}
									var send_tx_hash_info hexstring
									json.Unmarshal(body_send,&send_tx_hash_info)
									transaction_hash :=string(send_tx_hash_info.Result)
									if  transaction_hash != ""{
										fmt.Println("交易发送成功")
										fmt.Println("交易hash为:",transaction_hash)
										var tuichu string
										fmt.Println("输入任意值退出")
										fmt.Scan(&tuichu)
									}




								}else {
									fmt.Println("第二次签名失败")
								}
							}



						}else {
							fmt.Println("第一次签名失败")
						}


						/*

						//取出对应金额的utxo的未花费的输出
						for i,data :=range Result{

							utxo_data :=data.(map[string]interface{})
							utxo_txid :=utxo_data["txid"].(string)
							utxo_vout :=utxo_data["vout"].(float64)
							utxo_redeemScript :=utxo_data["redeemScript"].(string)
							utxo_scriptPubKey :=utxo_data["scriptPubKey"].(string)
						}

						*/


					}






				}
			}

		}


	}