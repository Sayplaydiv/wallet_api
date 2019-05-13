package server

import (
	"config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HttpPost(PostData string) []byte{

	//导入配置文件
	configMap := conf.InitConfig("src/config/conf.txt")
	//获取配置里host属性的value
	url:=configMap["url"]

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(PostData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//查看通信是否正常
	//fmt.Println("returnJson:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}


func HttpPost_1(PostData string) ([]byte){

	//获取配置里host属性的value
	url:="https://smssh1.253.com/msg/send/json"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(PostData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//查看通信是否正常
	//fmt.Println("returnJson:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	return body

}
