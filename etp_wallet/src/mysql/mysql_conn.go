package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
func InitDB(method string,args ...string)  {
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

	switch method_type {
	case "getnewaddress":
		Insert_addr(DB,args[0],args[1],args[2])
	case "getheight":
		Update_height(DB,args[0],args[1])
	case "sendfrom":
		Insert_tran(DB,args[0],args[1],args[2])
	case "getblock":
		Select_tran(DB,args[0])
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


func Select_tran(DB *sql.DB,asset string) {

	stmt, err := DB.Prepare("select height from height_status where asset=?")
	res, err := stmt.Exec(asset)
	if err != nil {
		fmt.Println("数据库执行查询出错", err)
		return
	}


	fmt.Println("共计",res )

}






