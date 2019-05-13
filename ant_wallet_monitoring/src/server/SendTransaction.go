package server

import (
	"config"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var phone_list string = "18301908030,13651931118"

type sendTransaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []interface{} `json:"params"`
	ID int `json:"id"`
}

type params  struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

func SendTransaction(){

	//导入配置文件
	configMap := conf.InitConfig("src/config/conf.txt")


	getGasPrice,getGasPrice_16 := GasPrice()
	GasPrice_val :=int64(getGasPrice)

	getBlance :=GetBalance()

	//账户最低预留金额
	/*
	ethFee_0:=configMap["ethFee"]
	ethFee, err := strconv.ParseInt(ethFee_0, 10, 64)
	if err !=nil {
		fmt.Println(err.Error())
	}
	*/

	//转帐金额
	transactionValue_0 := getBlance-(GasPrice_val*21000)
	transactionValue_1 := strconv.FormatInt(transactionValue_0, 16)

	if transactionValue_0 >0 {

		transactionValue :=string(transactionValue_1)

		//获取配置value
		fromAddress:=configMap["fromAddress"]
		toAddress:=configMap["toAddress"]
		passWord:=configMap["passWord"]

		fmt.Println("passWord:",passWord)

		PostParams := params{
			From:fromAddress,
			To:toAddress,
			Gas:"0x5208",
			GasPrice:getGasPrice_16,
			Value:"0x"+transactionValue,
			Data:"",
		}

		PostData := sendTransaction{
			Jsonrpc:"2.0",
			Method:"eth_sendTransaction",
			Params: []interface{}{PostParams},
			ID:4,
		}
		PostJsons, errs := json.Marshal(PostData)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(PostJsons))

		//解锁
		UnLock()

		body := HttpPost(string(PostJsons))

		type sendtransaction_result struct {
			Result string `json:"result"`
			Error string  `error:"result"`
			Id string     `id:"result"`
		}

		//json转struct
		var GetTxHash sendtransaction_result
		json.Unmarshal(body,&GetTxHash)
		TxHash :=GetTxHash.Result
		fmt.Println("本次转帐金额为：",transactionValue_0)
		fmt.Println("交易hash为：",string(TxHash))
		//加锁
		Lock()

		//短信通知：
		type message  struct {
			Account  string `json:"account"`
			Password string `json:"password"`
			Msg      string `json:"msg"`
			Phone    string `json:"phone"`
		}

		transactionValue_3 :=strconv.FormatInt((transactionValue_0/1000000000000000000),10)

		transactionValue_4 := strconv.FormatInt((transactionValue_0%1000000000000000000),10)
		transactionValue_5 := transactionValue_3 + "."+transactionValue_4


		timestamp:=time.Unix(time.Now().Unix(), 0)
		phone_list:=configMap["phoneList"]

		message_json := "蚂蚁钱包提示："+"本次归集账户转帐金额为："+transactionValue_5+"个ETH,"+"交易hash为："+"https://ropsten.etherscan.io/tx/"+string(TxHash)+"；"+"时间为："+timestamp.String()

		Post_message :=message{
			Account:"N3562857",
			Password:"5e10{fd}@KH2N#",
			Msg:message_json,
			Phone:phone_list,
		}

		message_result, errs := json.Marshal(Post_message)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(message_result))

		type msg_result struct {
			Code     string `json:"code"`
			MsgID    string `json:"msgId"`
			Time     string `json:"time"`
			ErrorMsg string `json:"errorMsg"`
		}
		body_msg_info := HttpPost_1(string(message_result))

		var body_msg_result msg_result
		json.Unmarshal(body_msg_info,&body_msg_result)
		if body_msg_result.ErrorMsg ==""{
			fmt.Println("短信发送成功，时间为：",body_msg_result.Time)
		}




	}else {

		timestamp:=time.Unix(time.Now().Unix(), 0)

		//短信通知：
		type message  struct {
			Account  string `json:"account"`
			Password string `json:"password"`
			Msg      string `json:"msg"`
			Phone    string `json:"phone"`
		}


		message_json := "目前金额不足,今天充值金额不足以转帐;"+"时间："+timestamp.String()
		phone_list:=configMap["phoneList"]

		Post_message :=message{
			Account:"N3562857",
			Password:"5e10{fd}@KH2N#",
			Msg:message_json,
			Phone:phone_list,
		}

		message_result, errs := json.Marshal(Post_message)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(message_result))

		type msg_result struct {
			Code     string `json:"code"`
			MsgID    string `json:"msgId"`
			Time     string `json:"time"`
			ErrorMsg string `json:"errorMsg"`
		}
		body_msg_info := HttpPost_1(string(message_result))

		var body_msg_result msg_result
		json.Unmarshal(body_msg_info,&body_msg_result)
		if body_msg_result.ErrorMsg ==""{
			fmt.Println("短信发送成功，时间为：",body_msg_result.Time)
		}
		fmt.Printf("\n目前金额不足,今天充值金额不足以转帐")
	}
	}
	


