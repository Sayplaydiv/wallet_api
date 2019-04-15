package main

import (
	"fmt"
	"server"
)

func main() {


	//获取时间datatime
	//server.GetNewAddress("wang812323")



	// 生成新地址 13。57。140。140
	//server.GetNewAddress("wang812323")


	var	from_address string
	var to_address string
	var value float64
	var privateke_1 string
	var privateke_2 string

	fmt.Println("id=2：listunspent查看utxo未花费的输出；id=3：createrawtransaction构建多重签名交易，id=4：signrawtransaction第一次签名，id=5:signrawtransaction第二次签名，id=6：sendrawtransaction提交多重签名交易")
	fmt.Println("请输入转帐地址：例如2N9ieB6WYf1FM4KAJAx9p3gs6ySmHfExGee")
	fmt.Scan(&from_address)
	fmt.Println("请输出收账地址：例如mvhRq2e5mjFnpz23QLoE5nwhknj7fT5wDt")
	fmt.Scan(&to_address)
	fmt.Println("请输入收账金额：例如0.001")
	fmt.Scan(&value)
	fmt.Println("请输入多签地址私钥1：cVHsAMmacJma77j4aYtPjn7bDoxeqsfgvL1BU6rxHjdX2YydVD6t")
	fmt.Scan(&privateke_1)
	fmt.Println("请输入多签地址私钥2：cTCsv3WFJCKukR4XUMvKxnXntiFaSLFToRM99HuLc6y9Gsx1qyuM")
	fmt.Scan(&privateke_2)


	/*

		// 生成新地址
		//server.GetNewAddress("wang812323")
		fmt.Println("id=2：listunspent查看utxo未花费的输出；id=3：createrawtransaction构建多重签名交易，id=4：signrawtransaction第一次签名，id=5:signrawtransaction第二次签名，id=6：sendrawtransaction提交多重签名交易")
		fmt.Println("请输入多重签名转帐地址：2NE4NX6desTqPsHzzCDqkQNvMtXjzgW7Yxy")
		fmt.Scan(&from_address)
		fmt.Println("请输出收账地址：mvhRq2e5mjFnpz23QLoE5nwhknj7fT5wDt")
		fmt.Scan(&to_address)
		fmt.Println("请输入收账金额：例如0.001")
		fmt.Scan(&value)
		fmt.Println("请输入多签地址私钥1：cW8ByfHyPcvRtGugaY4eQvosKsPmwKX9WbneFjZ96ixzk8LcsPER")
		fmt.Scan(&privateke_1)
		fmt.Println("请输入多签地址私钥2：cRE6o1faeyDaMzeLKBYWtbyrXAvJsgU2AnDGDSjSfTRY4neW2Pdm")
		fmt.Scan(&privateke_2)

	*/







	// 发送交易
	server.Transaction(from_address,to_address,value,privateke_1,privateke_2)


}