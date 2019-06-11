package main

import (
	"server"
	"sync"
	"time"
)

func main(){

	server.DepDelete()


}

var m *sync.RWMutex

func read(i int) {
	println(i, "read start")
	m.RLock()
	var p = 0
	var pr = "read"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(60 * time.Second)
		p++
		println(i, pr)

	}
	m.RUnlock()
	println(i, "read end")
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	var p = 0
	var pr = "write"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(60 * time.Second)
		p++
		println(i, pr)

	}
	m.Unlock()
	println(i, "write end")
}


/*

	erc20_job_aion.Insert()
	erc20_job_aion.DepDelete()
	erc20_job_aion.WitDelete()

	erc20_job_amb.Insert()
	erc20_job_amb.DepDelete()
	erc20_job_amb.WitDelete()

	erc20_job_cl.Insert()
	erc20_job_cl.DepDelete()
	erc20_job_cl.WitDelete()

	erc20_job_dlt.Insert()
	erc20_job_dlt.DepDelete()
	erc20_job_dlt.WitDelete()

	erc20_job_drop.Insert()
	erc20_job_drop.DepDelete()
	erc20_job_drop.WitDelete()

	erc20_job_hive.Insert()
	erc20_job_hive.DepDelete()
	erc20_job_hive.WitDelete()

	erc20_job_ico.Insert()
	erc20_job_ico.DepDelete()
	erc20_job_ico.WitDelete()

	erc20_job_iic.Insert()
	erc20_job_iic.DepDelete()
	erc20_job_iic.WitDelete()

	erc20_job_ind.Insert()
	erc20_job_ind.DepDelete()
	erc20_job_ind.WitDelete()


	erc20_job_key.Insert()
	erc20_job_key.DepDelete()
	erc20_job_key.WitDelete()

	erc20_job_mvc.Insert()
	erc20_job_mvc.DepDelete()
	erc20_job_mvc.WitDelete()

	erc20_job_ode.Insert()
	erc20_job_ode.DepDelete()
	erc20_job_ode.WitDelete()

	erc20_job_slrm.Insert()
	erc20_job_slrm.DepDelete()
	erc20_job_slrm.WitDelete()

	erc20_job_snt.Insert()
	erc20_job_snt.DepDelete()
	erc20_job_snt.WitDelete()

	erc20_job_storj.Insert()
	erc20_job_storj.DepDelete()
	erc20_job_storj.WitDelete()


	erc20_job_ufr.Insert()
	erc20_job_ufr.DepDelete()
	erc20_job_ufr.WitDelete()

	erc20_job_xnc.Insert()
	erc20_job_xnc.DepDelete()
	erc20_job_xnc.WitDelete()

*/


