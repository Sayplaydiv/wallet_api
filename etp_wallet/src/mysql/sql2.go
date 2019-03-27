package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

//args ...string
func InitDB_select_address(to_address string,tx_value string,hash string,tx_index string) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
	fmt.Println(to_address)

	rows, err := DB.Query("select * from address where  address=?", to_address)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var address string
		var asset string
		var inuse int
		var create_time string
		rows.Columns()
		err = rows.Scan(&id, &address, &asset, &inuse, &create_time)
		fmt.Println(id,address,asset,inuse,create_time)
		if inuse == 1 {
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			fmt.Println(to_address,tx_index,tx_value,hash)
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			//Insert_all(DB,to_address,tx_index,tx_value,hash)

				stmt, err := DB.Prepare("INSERT deposit SET  to_address =?,tx_value=?,tx_hash=?,tx_index=?")

				res, err := stmt.Exec(to_address,tx_value,hash,tx_index)
				if err != nil {
					fmt.Println("数据库执行插入出错", err)
					return
				}
				rowsaffected, err := res.RowsAffected()
				if err != nil {

				}
				fmt.Println("共计", rowsaffected, "行受到影响")




		}





	}
}