package mysql

import (
	"database/sql"
	"fmt"
	"strings"
)

//args ...string
func InitDB_insert_tx(to_address string,tx_index string,tx_value string,hash string) {
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


		stmt, err := DB.Prepare("INSERT deposit SET to=?,value=?,hash=?,tx_index=?")
		res, err := stmt.Exec(to_address,tx_index,tx_value,hash)
		if err != nil {
			fmt.Println("数据库执行插入出错", err)
			return
		}
		rowsaffected, err := res.RowsAffected()
		if err != nil {

		}
		fmt.Println("共计", rowsaffected, "行受到影响")





	fmt.Println("connnect success")

}
