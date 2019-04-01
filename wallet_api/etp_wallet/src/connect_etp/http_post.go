package connect_etp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
	"strings"
)


func Connetc_etp(Request string) []byte {
	url := "http://18.144.38.221:8820/rpc/v3"
	reqJson := Request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqJson))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("returnJson:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)

		return body

}

