package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

//配置参数
const (
	userName = "root"
	password = "wang812323"
	ip = "18.144.38.221"
	port = "3306"
	dbName = "eth_wallet"
)



//数据库链接池
var DB *sql.DB
func InitDB(method string,args ...string){
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
		fmt.Println(DB,"opon database fail")
		return
	}
	timestamp:=time.Unix(time.Now().Unix(), 0)

	//fmt.Println(timestamp)

	switch method {
	case "getnewaddress":
		fmt.Println(args[0])
		getnewaddress(DB,"BTC",1,timestamp,args[0],)
	}

	fmt.Println("connnect success")
}

func getnewaddress(DB *sql.DB,asset string,inuse int,timestamp time.Time,address string)  {
	stmt, err := DB.Prepare("INSERT address SET address=?,asset=?,inuse=?,create_time=? ")
	res, err := stmt.Exec(address,asset,inuse,timestamp)
	if err != nil {
		fmt.Println("数据库执行插入出错", err)
		return
	}
	rowsaffected, err := res.RowsAffected()
	if err != nil {

	}
	fmt.Println("共计", rowsaffected, "行受到影响")

}

