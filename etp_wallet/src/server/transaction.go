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
	to       := string_manage.Json02("tMrnmDJFDUPMNanCuu6ESkTPQSMWPoAGrg")
	amount   := string_manage.Json02("1000000")
	Params_0 :=[]string{userName,passWord,from,to,amount}

	Params :=strings.Join(Params_0,", ")



//	fmt.Print(Params)
	Request := RpcApi.RpcApi(id,method,Params)

	body :=connect_etp.Connetc_etp(Request)


	Body_info := string(body)

	fmt.Println("===========", Body_info)

	//buf1 := json.Unmarshal(Body_info)
	//buf := strings.Split(Body_info," , ")
	buf := strings.Fields(Body_info)

	for i,data :=range buf{


		if i == 12 {
			Txhash := strings.Trim(data," \", ")

			fmt.Println(Txhash)

				mysql.InitDB("sendfrom",Txhash,"etp","1")

	}
	}



	//console log :transaction hash
	//RpcApi.RpcApi(id,method,)

}




