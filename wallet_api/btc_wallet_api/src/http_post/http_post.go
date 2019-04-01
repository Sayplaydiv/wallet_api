package http_post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Id string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string   `json:"method"`
	Params []interface{} `json:"params"`
}


func HttpPost(id string,method string,params []interface{}) []byte{
		PostData_0 := Data{
			Id:id,
			Jsonrpc:"2.0",
			Method:method,
			Params:params,
		}
		PostData_1,err :=json.Marshal(PostData_0)
		if err !=nil{
			fmt.Println("error",err)
		}
		PostData :=string(PostData_1)

		fmt.Println(PostData)

	url := "http://13.57.140.140:88"
	reqJson := PostData
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqJson))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Basic YnRjcnBjOmRhY2IyMDE5MDMyMQ==")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("returnJson:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	return body

}