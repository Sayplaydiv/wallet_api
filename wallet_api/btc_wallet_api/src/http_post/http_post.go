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

const (
	url = "http://13.57.140.140:88"
)



func HttpPost(id string,method string,params []interface{}) []byte {
	PostData_0 := Data{
		Id:      id,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	}
	PostData_1, err := json.Marshal(PostData_0)
	if err != nil {
		fmt.Println("error", err)
	}
	if method == "createrawtransaction" {

		PostData_2 := string(PostData_1)
		if strings.Contains(PostData_2, "redeemScript") {
			str_0 := strings.Replace(PostData_2, "\"To_address\":", "", -1)
			str_1 := strings.Replace(str_0, ",\"Value\"", "", -1)
			str_2 := strings.Replace(str_1, ":[{", ":[[{", -1)
			str_3 := strings.Replace(str_2, "\"From_address\":", "", -1)
			str_4 := strings.Replace(str_3, ",\"Vout_value\"", "", -1)

			PostData := strings.Replace(str_4, "},{", "}],{", -1)

			fmt.Println("请求的数据",method,":", PostData)

			reqJson := PostData
			client := &http.Client{}
			req, err := http.NewRequest("POST", url, strings.NewReader(reqJson))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Basic YnRjcnBjOmRhY2IyMDE5MDMyMQ==")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			fmt.Println("returnJson:", resp.Status)
			body, _ := ioutil.ReadAll(resp.Body)
			return body

		}

	}else if method=="signrawtransaction" {
		PostData_2 := string(PostData_1)
		if strings.Contains(PostData_2, "redeemScript") {

			str_0 := strings.Replace(PostData_2, "{\"txid\"", "[{\"txid\"", -1)
			PostData := strings.Replace(str_0, "\"},[\"", "\"}],[\"", -1)

			fmt.Println("请求的数据",method,":", PostData)

			reqJson := PostData
			client := &http.Client{}
			req, err := http.NewRequest("POST", url, strings.NewReader(reqJson))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Basic YnRjcnBjOmRhY2IyMDE5MDMyMQ==")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			fmt.Println("returnJson:", resp.Status)
			body, _ := ioutil.ReadAll(resp.Body)
			return body

		}

	} else {
		PostData := string(PostData_1)

		fmt.Println("请求的数据",method,":", PostData)


		reqJson := PostData
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, strings.NewReader(reqJson))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Basic YnRjcnBjOmRhY2IyMDE5MDMyMQ==")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println("returnJson:", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	}
return []byte{}
}