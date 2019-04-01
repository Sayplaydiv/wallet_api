package main

import "server"

func main()  {

	//"{\"id\":125,\"jsonrpc\":\"2.0\",\"method\":\"sendfrom\",\"params\":[\"jiang2\",\"jiang2\",\"MRaZS6Yh3JeAo3v4NWEypZsGuZyoqgGiMa\",\"MAGGFkPSCAW9ZhUzCbw61LGMfg5Ks8ASx6\",\"1000000\"]}"

	//RpcApi.RpcApi("1","sendfrom","something")

	//交易
	//server.Transction()


	//生成地址
	//server.Address()

	//创建信道
	//complete := make(chan int)


	//获取最新块高,更新数据库高度

	  server.GetHeight()


	//扫区块获取交易
	defer server.GetBlock()

	//数据库清除
	//mysql.InitDB("delete","75")

}