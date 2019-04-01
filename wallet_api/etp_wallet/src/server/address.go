package server

import (
	"RpcApi"
	"connect_etp"
	"fmt"
	"mysql"

	//"mysql"
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

	body :=connect_etp.Connetc_etp(Request)


	Body_info := string(body)

	fmt.Println("===========", Body_info)

	//buf1 := json.Unmarshal(Body_info)
	//buf := strings.Split(Body_info," , ")
	buf := strings.Fields(Body_info)


	for i, data :=range buf{

		if i == 10 {
			NewAddress := strings.Trim(data," \", ")
			//fmt.Println(i,",",GetLastBlock)
			fmt.Println(NewAddress)

			mysql.InitDB("getnewaddress",NewAddress,"etp","1")

		}
	}






}