package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

//args ...string
func InitDB_select_address(to_address string,tx_index string,tx_value string,hash string) {
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
		fmt.Println(inuse)



		fmt.Println("connnect success")

		
	}
}