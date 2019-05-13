package main

import (
	"server"
	"time"
)

func main(){
	
	//生产使用
	for i:=1;i<100000000000000;i++ {
		timer:=time.NewTimer(time.Hour * 24)
		<-timer.C
		server.SendTransaction()
	}



}