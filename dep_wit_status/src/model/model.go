package model

import (
	"dbConnect"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	DepwitPost struct {
		Method string `json:"method"`
		User_id int `json:"user_id"`
	}

	// response entity
	DepwitStatus struct {
		Address string `json:"address"`
		Asset string `json:"asset"`
		User_id int `json:"user_id"`
		Amount float64 `json:"amount"`
		Tx_hash string `json:"tx_hash"`
		Block_height int `json:"block_height"`
		Tx_time string `json:"tx_time"`
		Confirm_height int `json:"confirm_height"`
	}
)
var timestamp =time.Unix(time.Now().Unix(), 0)
// createAPI add a new api
func PostDepwit(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	//获取post body数据
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	test := string(buf[0:n])

	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(test), &mapResult)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}

	//json转化为struct
	var requst DepwitPost
	json.Unmarshal([]byte(test), &requst)



	if requst.Method == "deposit" {

		response :=Respon(requst.User_id,1)



		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"deposit": response,
		})

	} else if requst.Method == "withdraw"{

		//提币中状态返回
		response :=Respon(requst.User_id,2)

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"withdraw":response,
		})

	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": "method error，pls {\"method\":\"withdraw\",\"user_id\":123}\"",
		})
	}
}



//返回多组数据处理
func Respon(user_id int,confirm int) ([]interface{}) {

	var n []interface{}

	result :=dbConnect.InitDB_wallet(user_id,confirm)
	for i := 0; i < len(result); i++ {
		if confirm == 1 {
			dep_asset := string(result[i]["asset"])
			dep_address := string(result[i]["address"])
			dep_amount, err := strconv.ParseFloat(result[i]["amount"], 64)
			dep_tx_hash := string(result[i]["tx_hash"])
			dep_block_height, err := strconv.Atoi(result[i]["block_height"])
			dep_confirm_height, err := strconv.Atoi(result[i]["confirm_height"])
			tx_time :=result[i]["tx_time"]
			if err != nil {
				log.Fatal(err)
			}

			//冲币中状态返回
			repose_0 := DepwitStatus{
				Address:      dep_address,
				Asset:        dep_asset,
				User_id:      user_id,
				Amount:       dep_amount,
				Tx_hash:      dep_tx_hash,
				Block_height: dep_block_height,
				Tx_time:      tx_time,
				Confirm_height:dep_confirm_height,
			}
			n =append(n,repose_0)

		}else if confirm == 2{
			wit_asset := string(result[i]["asset"])
			wit_address := string(result[i]["address"])
			wit_amount, err := strconv.ParseFloat(result[i]["amount"], 64)
			wit_tx_hash := string(result[i]["tx_hash"])
			tx_time :=result[i]["tx_time"]
			if err != nil {
				log.Fatal(err)
			}
			//冲币中状态返回
			repose_0 := DepwitStatus{
				Address:      wit_address,
				Asset:        wit_asset,
				User_id:      user_id,
				Amount:       wit_amount,
				Tx_hash:      wit_tx_hash,
				Tx_time:      tx_time,
			}
			n =append(n,repose_0)
		}
	}
	return n
}


