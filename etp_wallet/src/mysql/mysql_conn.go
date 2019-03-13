package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	_"github.com/go-sql-driver/mysql"
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
func InitDB(address string,asset string,inuse string)  {
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
//
	Insert(DB,address,asset,inuse)


	fmt.Println("connnect success")




}


//address string,asset string

func Insert(DB *sql.DB,address string,asset string,inuse string) {

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

func Update(DB *sql.DB,address string,asset string,inuse string) {

	stmt, err := DB.Prepare("Update test SET address=?,asset=?,inuse=?")
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


func Query(DB *sql.DB,address string,asset string,inuse string) {

	stmt, err := DB.Prepare("Update test SET address=?,asset=?,inuse=?")
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