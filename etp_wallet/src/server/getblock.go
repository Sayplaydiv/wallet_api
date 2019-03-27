package server

import (
	"mysql"
)

 func GetBlock()  {



		id := "4"

		method  := "getblock"

		//userName := string_manage.Json02("wangyaxing")
		//passWord := string_manage.Json02("wang182323")
		//account     := string_manage.Json02("test233")
		//passWord    := string_manage.Json02("wang812323")
		//amount   := string_manage.Json02("1000000")
		//Params_0 :=[]string{account,passWord,}

		//Params :=strings.Join(Params_0,", ")
		//fmt.Print(Params)

		mysql.InitDB("getblock","etp",id,method)
}