package main

import (
	"erc20_job_aion"
	"erc20_job_amb"
	"erc20_job_cl"
	"erc20_job_dlt"
	"erc20_job_drop"
	"erc20_job_hive"
	"erc20_job_ico"
	"erc20_job_iic"
	"erc20_job_ind"
	"erc20_job_key"
	"erc20_job_mvc"
	"erc20_job_ode"
	"erc20_job_slrm"
	"erc20_job_snt"
	"erc20_job_storj"
	"erc20_job_ufr"
	"erc20_job_xnc"
	"fmt"
	"router"
	"runtime"
	"server"
	"time"
)
var quit chan int = make(chan int)

func apiStart()  {
	router.RouterInit()
	runtime.Gosched()
	quit <- 0
}
func selectStart(){
	server.Insert()
	server.DepDelete()
	//server.WitDelete()
	runtime.Gosched()
	quit <- 1
}

func aionStart()  {
	erc20_job_aion.Insert()
	erc20_job_aion.DepDelete()
	//erc20_job_aion.WitDelete()
	runtime.Gosched()
	quit <- 2

}

func ambStart()  {
	erc20_job_amb.Insert()
	erc20_job_amb.DepDelete()
	//erc20_job_amb.WitDelete()
	runtime.Gosched()
	quit <- 3
}

func clStart()  {
	erc20_job_cl.Insert()
	erc20_job_cl.DepDelete()
	//erc20_job_cl.WitDelete()
	runtime.Gosched()
	quit <- 4
}

func dltStart()  {
	erc20_job_dlt.Insert()
	erc20_job_dlt.DepDelete()
	//erc20_job_dlt.WitDelete()
	runtime.Gosched()
	quit <- 5
}

func dropStart()  {
	erc20_job_drop.Insert()
	erc20_job_drop.DepDelete()
	//erc20_job_drop.WitDelete()
	runtime.Gosched()
	quit <- 6
}

func hiveStart()  {
	erc20_job_hive.Insert()
	erc20_job_hive.DepDelete()
	//erc20_job_hive.WitDelete()
	runtime.Gosched()
	quit <- 7
}

func icoStart(){
	erc20_job_ico.Insert()
	erc20_job_ico.DepDelete()
	//erc20_job_ico.WitDelete()
	runtime.Gosched()
	quit <- 8
}

func iicStart(){
	erc20_job_iic.Insert()
	erc20_job_iic.DepDelete()
	//erc20_job_iic.WitDelete()
	runtime.Gosched()
	quit <- 9
}

func indStart()  {
	erc20_job_ind.Insert()
	erc20_job_ind.DepDelete()
	//erc20_job_ind.WitDelete()
	runtime.Gosched()
	quit <- 10
}

func keyStart()  {
	erc20_job_key.Insert()
	erc20_job_key.DepDelete()
	//erc20_job_key.WitDelete()
	runtime.Gosched()
	quit <- 11
}


func mvcStart(){
	erc20_job_mvc.Insert()
	erc20_job_mvc.DepDelete()
	//erc20_job_mvc.WitDelete()
	runtime.Gosched()
	quit <- 12
}

func odeStart()  {
	erc20_job_ode.Insert()
	erc20_job_ode.DepDelete()
	//erc20_job_ode.WitDelete()
	runtime.Gosched()
	quit <- 13

}

func slrmStart(){
	erc20_job_slrm.Insert()
	erc20_job_slrm.DepDelete()
	//erc20_job_slrm.WitDelete()
	runtime.Gosched()
	quit <- 14
}

func sntStart()  {
	erc20_job_snt.Insert()
	erc20_job_snt.DepDelete()
	//erc20_job_snt.WitDelete()
	runtime.Gosched()
	quit <- 15

}

func storjStart()  {
	erc20_job_storj.Insert()
	erc20_job_storj.DepDelete()
	//erc20_job_storj.WitDelete()
	runtime.Gosched()
	quit <- 16

}

func ufrStart()  {
	erc20_job_ufr.Insert()
	erc20_job_ufr.DepDelete()
	//erc20_job_ufr.WitDelete()
	runtime.Gosched()
	quit <- 17
}

func xncStart()  {
	erc20_job_xnc.Insert()
	erc20_job_xnc.DepDelete()
	//erc20_job_xnc.WitDelete()
	runtime.Gosched()
	quit <- 18

}



func main() {

	// 设置最大开n个原生线程
	runtime.GOMAXPROCS(30)
	fmt.Println("start ---")
	go apiStart()
	for{
		go selectStart()
		go aionStart()
		go ambStart()
		go clStart()
		go dltStart()
		go dropStart()
		go hiveStart()
		go icoStart()
		go iicStart()
		go indStart()
		go keyStart()
		go mvcStart()
		go odeStart()
		go slrmStart()
		go sntStart()
		go storjStart()
		go ufrStart()
		go xncStart()

		fmt.Println("start ===")
		for i := 1; i <18; i++ {
			sc := <-quit
			fmt.Println("线程管道:", sc)
			time.Sleep(1*time.Second)
		}

	}

	fmt.Println("end")



}


