package RpcApi

import (
	"fmt"
	"string_manage"
	"strings"
)

//jsonrpc 2。0
func RpcApi(Id string,Method string,Params string) string{
	//id
	Id_0 := []string{"id","1"}
	Id_1 := strings.Join(Id_0,"\":")
	Id_2 := "\""+ Id_1

	//jsonrpc version
	JsonRpc_0 := []string{"jsonrpc","2.0"}
	JsonRpc_1 := strings.Join(JsonRpc_0,"\":\"")
	JsonRpc_2 := "\""+ JsonRpc_1 + "\""

	//etp method
	Method_0 :=[]string{"method",Method}
	Method_1 :=strings.Join(Method_0,"\":\"")
	Method_2 := "\""+ Method_1 + "\""

	//etp method Params
	Params_0 := "["+ Params + "]"
	Params_1 := string_manage.Json02("params")
	Params_2 :=[]string{Params_1,Params_0}

	Params_3 :=strings.Join(Params_2,":")


	Jsonrpc_0 :=[]string{Id_2,JsonRpc_2,Method_2,Params_3}
	Json_3 := strings.Join(Jsonrpc_0,", ")
	Json := "{" + Json_3 + "}"

	 //console log

	//Json :="{\"id\":\"1\",\"jsonrpc\":\"2.0\",\"method\":\"sendfrom\",params:[\"wangyaxing\",\"wang182323\",\"tPak5DS4aR577dAEtCUDMzvZe2RKzQticA\",\"MAGGFkPSCAW9ZhUzCbw61LGMfg5Ks8ASx6\",\"1000000\"]}"

	//Json1 :="{\"id\":1, \"jsonrpc\":\"2.0\", \"method\":\"sendfrom\", \"params\":[\"wangyaxing\", \"wang182323\", \"tPak5DS4aR577dAEtCUDMzvZe2RKzQticA\", \"MAGGFkPSCAW9ZhUzCbw61LGMfg5Ks8ASx6\", \"10000\"]}"
	fmt.Println("请求json数据为：",Json)

	//fmt.Println(Json1)

	var Request_0 string = Json

	//request json to etp node


			return Request_0




}





