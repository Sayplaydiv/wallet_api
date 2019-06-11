package main

import (
	_ "DW_status/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

var userName = beego.AppConfig.String("mysqluser")
var userPass = beego.AppConfig.String("mysqlpass")
var	mysqlurls = beego.AppConfig.String("mysqlurls")
var	mysqldb  = beego.AppConfig.String("mysqldb")
var	mysqlprot =beego.AppConfig.String("mysqlprot")



func main() {
	if beego.BConfig.RunMode == beego.BConfig.RunMode {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}



func init() {


	//sql配置
	dataSource := strings.Join([]string{userName, ":", userPass, "@tcp(",mysqlurls, ":", mysqlprot, ")/", mysqldb, "?charset=utf8"}, "")
	//fmt.Println(dataSource)



	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
}