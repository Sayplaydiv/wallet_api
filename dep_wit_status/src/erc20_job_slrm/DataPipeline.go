package erc20_job_slrm

import (
	"fmt"
	"log"
	"strconv"
)

func Insert() {

	//冲币状态入库
	dep_result := InitDB_dep()
	for i := 0; i < len(dep_result); i++ {

		dep_asset := string(dep_result[i]["asset"])
		dep_address := string(dep_result[i]["address"])
		dep_amount, err := strconv.ParseFloat(dep_result[i]["amount"], 64)
		dep_tx_hash := string(dep_result[i]["tx_hash"])
		dep_block_height, err := strconv.Atoi(dep_result[i]["block_height"])
		dep_confirm,err := strconv.Atoi(dep_result[i]["status"])
		dep_confirm_height,err := strconv.Atoi(dep_result[i]["confirm_height"])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("冲币信息：",i,dep_address)
		user_id := InitDB_audit("deposit",dep_address)

		if user_id !=0 {
			//asset string,address string,user_id int,amount float64,tx_hash string,block_height int,tx_time time.Time,confirm int
			InitDB_insert(dep_asset,dep_address,user_id,dep_amount,dep_tx_hash,dep_block_height,dep_confirm,dep_confirm_height)
		}

	}

	//提币状态入库
	wit_result := InitDB_wit()
	for i :=0;i<len(wit_result) ; i++ {
		wit_asset := string(wit_result[i]["asset"])
		wit_address := string(wit_result[i]["to_address"])
		wit_id := string(wit_result[i]["withdraw_id"])
		wit_amount, err := strconv.ParseFloat(wit_result[i]["amount"], 64)
		wit_tx_hash := string(wit_result[i]["txhash"])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(wit_asset,wit_address,wit_amount,wit_tx_hash)
		fmt.Println("提币信息：",wit_id)
		user_id := InitDB_audit("withdraw",wit_id)
		fmt.Println(user_id)
		if user_id == 1000000000000 {
			//asset string,address string,user_id int,amount float64,tx_hash string,block_height int,tx_time time.Time,confirm int
			InitDB_insert(wit_asset,wit_address,user_id,wit_amount,wit_tx_hash,0,2,0)
		}
	}
}

