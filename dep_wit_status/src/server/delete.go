package server

import (
	"dbConnect"
	"fmt"
)

func DepDelete(){

	dep_status_result :=dbConnect.InitDB_wallet_deposit("deposit")
	//fmt.Println("dep_status_result",dep_status_result)

	for i:=0;i< len(dep_status_result); i++ {
		check_tx_hash :=dep_status_result[i]["tx_hash"]
		//fmt.Println("check_tx_hash",check_tx_hash)
		deposit_status :=dbConnect.InitDB_wallet_delete("select_depsoit",check_tx_hash)
		deposit_height :=dbConnect.InitDB_wallet_delete("select_deposit_height",check_tx_hash)
		//fmt.Println("DepDelete update height:",deposit_height)
		fmt.Println("DepDelete update height:",deposit_height)
		dbConnect.InitDB_wallet_update(deposit_height,check_tx_hash)
		//fmt.Println("deposit_status",deposit_status)
		if deposit_status ==2 {
			fmt.Println("状态值：",deposit_status)
			fmt.Println("checkhash：",check_tx_hash)
			dbConnect.InitDB_wallet_delete_1("delete",check_tx_hash)
		}


	}
}

func WitDelete()  {
	wit_status_result :=dbConnect.InitDB_wallet_deposit("witdraw")
	//fmt.Println("wit_status_result",wit_status_result)
	for i:=0;i< len(wit_status_result);i++  {
		check_tx_hash :=wit_status_result[i]["tx_hash"]
		//fmt.Println("check_tx_hash",check_tx_hash)
		withdraw_process :=dbConnect.InitDB_wallet_delete("select_witdraw_process",check_tx_hash)
		withdraw_status := dbConnect.InitDB_wallet_delete("select_witdraw_status",check_tx_hash)
		if withdraw_process==4{
			dbConnect.InitDB_wallet_delete_1("delete",check_tx_hash)
		}else if withdraw_status==4{
			dbConnect.InitDB_wallet_delete_1("delete",check_tx_hash)
		}else if withdraw_process<0 {
			dbConnect.InitDB_wallet_delete_1("delete",check_tx_hash)
		}
	}

}

