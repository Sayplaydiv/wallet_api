package main

import (
	"server"
)

func main() {


	//获取时间datatime
	//server.GetNewAddress("wang812323")



	// 生成新地址 13。57。140。140
	//server.GetNewAddress("wang812323")

	/*
		var	from_address string
		var to_address string
		var value float64
		var privateke_1 string
		var privateke_2 string


		fmt.Println("id=2：listunspent查看utxo未花费的输出；id=3：createrawtransaction构建多重签名交易，id=4：signrawtransaction第一次签名，id=5:signrawtransaction第二次签名，id=6：sendrawtransaction提交多重签名交易")
		fmt.Println("请输入转帐地址：例如2NEQo5P3AFNfUf8cBVvnuSHBJiYF9mEPR3P")
		fmt.Scan(&from_address)
		fmt.Println("请输出收账地址：例如mvhRq2e5mjFnpz23QLoE5nwhknj7fT5wDt")
		fmt.Scan(&to_address)
		fmt.Println("请输入收账金额：例如0.001")
		fmt.Scan(&value)
		fmt.Println("请输入多签地址私钥1：cVbNQMDy3K9ncpAregcbiRMtvz7hfsqiWAv4xSVex9apN35otBHn")
		fmt.Scan(&privateke_1)
		fmt.Println("请输入多签地址私钥2：cSRHPY7QGxeicLqUVKddjRH3oxUJKBHKVrBFuVEDFpJhn48xRnty")
		fmt.Scan(&privateke_2)

		*/







	// 发送交易
	server.Transaction("2NEQo5P3AFNfUf8cBVvnuSHBJiYF9mEPR3P","mvhRq2e5mjFnpz23QLoE5nwhknj7fT5wDt",0.1,"cVbNQMDy3K9ncpAregcbiRMtvz7hfsqiWAv4xSVex9apN35otBHn","cSRHPY7QGxeicLqUVKddjRH3oxUJKBHKVrBFuVEDFpJhn48xRnty")


}