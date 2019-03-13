package server

import (
	"RpcApi"
	"connect_etp"
	"fmt"
	"mysql"
	"string_manage"
	"strings"
)

func Address()  {
	id := "2"

	method  := "getnewaddress"

	//userName := string_manage.Json02("wangyaxing")
	//passWord := string_manage.Json02("wang182323")
	account     := string_manage.Json02("test233")
	passWord    := string_manage.Json02("wang812323")
	//amount   := string_manage.Json02("1000000")
	Params_0 :=[]string{account,passWord,}

	Params :=strings.Join(Params_0,", ")
	//fmt.Print(Params)
	Request := RpcApi.RpcApi(id,method,Params)

	Body_msg :=connect_etp.Connetc_etp(Request)

	fmt.Println(Body_msg)


	for i,data :=range Body_msg{

		if i == 10 {
			NewAddress := strings.Trim(data," \", ")

			fmt.Println(NewAddress)

			mysql.InitDB(NewAddress,"etp","1")

		}
	}



}