package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)





func main() {

	json_str := `{"sn":1,"ls":false,"bg":0,"ed":0,"ws":[{"bg":0,"cw":[{"sc":0,"w":"还"}]},{"bg":0,"cw":[{"sc":0,"w":"有点"}]},{"bg":0,"cw":[{"sc":0,"w":"眼熟"}]}]}`

	fmt.Println(json_str)
//var recResult string
var dat map[string]interface{}

	json.Unmarshal([]byte(json_str), &dat)

	fmt.Println(dat)

	fmt.Println(reflect.TypeOf(dat))
	ws:=dat["ws"].([]interface{})
	fmt.Println("wss==========",reflect.TypeOf(ws))
	for i,data:=range ws{
		wsMap := data.(map[string]interface{})
		fmt.Println(i,data)
		fmt.Println(wsMap["bg"])

	}






	fmt.Println(dat["ws"].([]interface{}))
}


