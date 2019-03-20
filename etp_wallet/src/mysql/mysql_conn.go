package mysql

import (
	"RpcApi"
	"connect_etp"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "wang812323"
	ip = "18.144.38.221"
	port = "3306"
	dbName = "eth_wallet"
)


var DB *sql.DB
//args ...string
func InitDB(method string,args ...string) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}

	//defer DB.Close()
	method_type := method


	switch method_type{
	case "getnewaddress":
		Insert_addr(DB,args[0],args[1],args[2])
	case "getheight":
		Update_height(DB,args[0],args[1])
	case "sendfrom":
		Insert_tran(DB,args[0],args[1],args[2])
	case "getblock":
		Select_tran(args[0],args[1],args[2])
		//height_status := Select_tran(args[0])
		//return height_status
	//	fmt.Println(height_status)
	}

	fmt.Println("connnect success")

}




//address string,asset string

func Insert_addr(DB *sql.DB,address string,asset string,inuse string) {

	stmt, err := DB.Prepare("INSERT address SET address=?,asset=?,inuse=?")
	res, err := stmt.Exec(address,asset,inuse)
	if err != nil {
		fmt.Println("数据库执行插入出错", err)
		return
	}
	rowsaffected, err := res.RowsAffected()
	if err != nil {

	}
	fmt.Println("共计", rowsaffected, "行受到影响")

}

func Update_height(DB *sql.DB,height string,asset string) {

	stmt, err := DB.Prepare("Update last_height SET height=? where asset=?")
	res, err := stmt.Exec(height,asset)
		if err != nil {
		fmt.Println("数据库执行更新出错", err)
		return
	}
	rowsaffected, err := res.RowsAffected()
	if err != nil {

	}
	fmt.Println("共计", rowsaffected, "行受到影响")

}


func Insert_tran(DB *sql.DB,address string,asset string,inuse string) {


	stmt, err := DB.Prepare("INSERT test SET address=?,asset=?,inuse=?")
	res, err := stmt.Exec(address,asset,inuse)
	if err != nil {
		fmt.Println("数据库执行插入出错", err)
		return
	}
	rowsaffected, err := res.RowsAffected()
	if err != nil {

	}
	fmt.Println("共计", rowsaffected, "行受到影响")

}

func Select_tran(Asset string,id string,method string) {

	var height string
	err := DB.QueryRow("select height from height_status where  Asset= ?", Asset).Scan(&height)
	if err != nil {
		log.Fatal(err)
	}

	//userName := string_manage.Json02("wangyaxing")
	//passWord := string_manage.Json02("wang182323")
	//account     := string_manage.Json02("test233")
	//passWord    := string_manage.Json02("wang812323")
	//amount   := string_manage.Json02("1000000")
	//Params_0 :=[]string{account,passWord,}

	//Params :=strings.Join(Params_0,", ")
	//fmt.Print(Params)

	Params := height

	Request := RpcApi.RpcApi(id, method, Params)

	body :=connect_etp.Connetc_etp(Request)


	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		result := dat["result"].(map[string]interface{})
		transaction :=result["transactions"].([]interface{})

		for i, data :=range  transaction{
			fmt.Println(i,data)
			fmt.Println(data)



		}

	}



}




