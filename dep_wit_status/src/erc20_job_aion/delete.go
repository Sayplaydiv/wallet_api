package erc20_job_aion

import (
	"fmt"
)

func DepDelete(){

	dep_status_result :=InitDB_wallet_deposit("deposit")
	//fmt.Println("dep_status_result",dep_status_result)

	for i:=0;i< len(dep_status_result); i++ {
		check_tx_hash :=dep_status_result[i]["tx_hash"]
		//fmt.Println("check_tx_hash",check_tx_hash)
		deposit_status :=InitDB_wallet_delete("select_depsoit",check_tx_hash)
		deposit_height :=InitDB_wallet_delete("select_deposit_height",check_tx_hash)
		//fmt.Println("DepDelete update height:",deposit_height)
		fmt.Println("DepDelete update height:",deposit_height)
		InitDB_wallet_update(deposit_height,check_tx_hash)
		//fmt.Println("deposit_status",deposit_status)
		if deposit_status ==2 {
			InitDB_wallet_delete_1("delete",check_tx_hash)
		}
	}
}

func WitDelete()  {
	wit_status_result :=InitDB_wallet_deposit("witdraw")
	//fmt.Println("wit_status_result",wit_status_result)
	for i:=0;i< len(wit_status_result);i++  {
		check_tx_hash :=wit_status_result[i]["tx_hash"]
		//fmt.Println("check_tx_hash",check_tx_hash)
		withdraw_process :=InitDB_wallet_delete("select_witdraw_process",check_tx_hash)
		withdraw_status := InitDB_wallet_delete("select_witdraw_status",check_tx_hash)
		if withdraw_process==4{
			InitDB_wallet_delete_1("delete",check_tx_hash)
		}else if withdraw_status==4{
			InitDB_wallet_delete_1("delete",check_tx_hash)
		}else if withdraw_process<0 {
			InitDB_wallet_delete_1("delete",check_tx_hash)
		}
	}

}

