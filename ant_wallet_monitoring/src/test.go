package main

import (
	"server"
	"time"
)

func main(){


	//test使用
	for i:=1;i<100000000000000;i++ {
		timer:=time.NewTimer(time.Minute * 5)
		<-timer.C
		server.SendTransaction()
	}

	/*

	//生产使用
	for i:=1;i<100000000000000;i++ {
		timer:=time.NewTimer(time.Hour * 24)
		<-timer.C
		server.SendTransaction()
	}

	 */

}