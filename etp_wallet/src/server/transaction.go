package server

import (
	"RpcApi"
	"connect_etp"
	"fmt"
	"mysql"
	"string_manage"
	"strings"
)



func Transction()  {

	id := "1"

	method  := "sendfrom"

	userName := string_manage.Json02("wangyaxing")
	passWord := string_manage.Json02("wang182323")
	from     := string_manage.Json02("tPak5DS4aR577dAEtCUDMzvZe2RKzQticA")
	to       := string_manage.Json02("t82Z9X6aazyBJuq76rCtRpU7speMqbstzy")
	amount   := string_manage.Json02("1000000")
	Params_0 :=[]string{userName,passWord,from,to,amount}

	Params :=strings.Join(Params_0,", ")



//	fmt.Print(Params)
	Request := RpcApi.RpcApi(id,method,Params)

	Body_msg :=connect_etp.Connetc_etp(Request)


	for i,data :=range Body_msg{


		if i == 12 {
			Txhash := strings.Trim(data," \", ")

			fmt.Println(Txhash)

				mysql.InitDB(Txhash,"etp","1")

	}
	}

	//console log :transaction hash


		//RpcApi.RpcApi(id,method,)




}




