package erc20_job_xnc

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)



//根据user_id 获取地址
func InitDB_audit(method string,get_userid string)(int){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_audit:=configMap["userName_audit"]
	password_audit:=configMap["password_audit"]
	ip_audit:=configMap["ip_audit"]
	port_audit:=configMap["port_audit"]
	dbName_audit:=configMap["dbName_audit"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_audit, ":", password_audit, "@tcp(",ip_audit, ":", port_audit, ")/", dbName_audit, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}
	//timestamp:=time.Unix(time.Now().Unix(), 0)

	switch method {
	case "deposit":
		var userId int
		err := DB.QueryRow("select userId from t_user_deposit_address where address= ?", get_userid).Scan(&userId)
		if err == sql.ErrNoRows{
			fmt.Println(get_userid)
			log.Fatal("user_id:deposit没找到对应的id，请排查")
			userId := 0
			return userId
		}else {
			return userId
		}
	case "withdraw":
		var userId int
		err := DB.QueryRow("select userId from t_withdraw_audit where orderId= ?", get_userid).Scan(&userId)
		if err == sql.ErrNoRows{
			fmt.Println(get_userid)
			log.Fatal("user_id:withdraw没找到对应的id，请排查")
			userId := 0
			return userId
		}else {
			return userId
		}

	}
	//fmt.Println("connnect success")
	defer DB.Close();
	return 1
}

//获取冲币信息
func InitDB_dep()  (map[int]map[string]string){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_wallet:=configMap["userName_wallet"]
	password_wallet:=configMap["password_wallet"]
	ip_wallet:=configMap["ip_wallet"]
	port_wallet:=configMap["port_wallet"]
	dbName_wallet:=configMap["dbName_wallet"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_wallet, ":", password_wallet, "@tcp(",ip_wallet, ":", port_wallet, ")/", dbName_wallet, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}
	//timestamp:=time.Unix(time.Now().Unix(), 0)


	//查询数据，取所有字段
	rows2, _ := DB.Query("select * from deposit_xnc where status=1");
	//返回所有列
	cols, _ := rows2.Columns();
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols));
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k];
	}

	i := 0;
	result := make(map[int]map[string]string);
	for rows2.Next(){
		//填充数据
		rows2.Scan(scans...);
		//每行数据
		row := make(map[string]string);
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k];
			//这里把[]byte数据转成string
			row[key] = string(v);
		}
		//放入结果集
		result[i] = row;
		i++;
	}
	defer DB.Close();
	return result

}



//获取提币信息
func InitDB_wit()  (map[int]map[string]string){
	var DB *sql.DB


	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_wallet:=configMap["userName_wallet"]
	password_wallet:=configMap["password_wallet"]
	ip_wallet:=configMap["ip_wallet"]
	port_wallet:=configMap["port_wallet"]
	dbName_wallet:=configMap["dbName_wallet"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_wallet, ":", password_wallet, "@tcp(",ip_wallet, ":", port_wallet, ")/", dbName_wallet, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}
	//timestamp:=time.Unix(time.Now().Unix(), 0)


	//查询数据，取所有字段
	rows2, _ := DB.Query("select * from withdraw_xnc where process<4 and process>=0 and status<4");
	//返回所有列
	cols, _ := rows2.Columns();
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols));
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k];
	}

	i := 0;
	result := make(map[int]map[string]string);
	for rows2.Next(){
		//填充数据
		rows2.Scan(scans...);
		//每行数据
		row := make(map[string]string);
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k];
			//这里把[]byte数据转成string
			row[key] = string(v);
		}
		//放入结果集
		result[i] = row;
		i++;
	}
	defer DB.Close();
	return result

}


//定时job 存储数据
func InitDB_insert(asset string,address string,user_id int,amount float64,tx_hash string,block_height int,confirm int,confirm_height int)(){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_audit:=configMap["userName_audit"]
	password_audit:=configMap["password_audit"]
	ip_audit:=configMap["ip_audit"]
	port_audit:=configMap["port_audit"]
	dbName_audit:=configMap["dbName_audit"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_audit, ":", password_audit, "@tcp(",ip_audit, ":", port_audit, ")/", dbName_audit, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}
	timestamp:=time.Unix(time.Now().Unix(), 0)

	var address_status string
	err := DB.QueryRow("select address from depwit_status where tx_hash= ?",tx_hash).Scan(&address_status)   //db为sql.DB
	if err == sql.ErrNoRows{

		stmt, err := DB.Prepare("INSERT depwit_status SET asset=?,address=?,user_id=?,amount=?,tx_hash=?,block_height=?,tx_time=?,confirm=?,confirm_height=?")
		res, err := stmt.Exec(asset,address,user_id,amount,tx_hash,block_height,timestamp,confirm,confirm_height)
		if err != nil {
			log.Fatal(err)
		}
		rowsaffected, err := res.RowsAffected()
		fmt.Println("本次共计插入数据：", rowsaffected, "行")

	}

	defer DB.Close();

}


//数据查询
func InitDB_wallet_deposit(method string)(map[int]map[string]string){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_audit:=configMap["userName_audit"]
	password_audit:=configMap["password_audit"]
	ip_audit:=configMap["ip_audit"]
	port_audit:=configMap["port_audit"]
	dbName_audit:=configMap["dbName_audit"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_audit, ":", password_audit, "@tcp(",ip_audit, ":", port_audit, ")/", dbName_audit, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}

	//查询数据，取所有字段
	switch method {
	case "deposit":

		//查询数据，取所有字段
		rows2, _ := DB.Query("select tx_hash from depwit_status where confirm=1");
		//返回所有列
		cols, _ := rows2.Columns();
		//这里表示一行所有列的值，用[]byte表示
		vals := make([][]byte, len(cols));
		//这里表示一行填充数据
		scans := make([]interface{}, len(cols));
		//这里scans引用vals，把数据填充到[]byte里
		for k, _ := range vals {
			scans[k] = &vals[k];
		}

		i := 0;
		result := make(map[int]map[string]string);
		for rows2.Next(){
			//填充数据
			rows2.Scan(scans...);
			//每行数据
			row := make(map[string]string);
			//把vals中的数据复制到row中
			for k, v := range vals {
				key := cols[k];
				//这里把[]byte数据转成string
				row[key] = string(v);
			}
			//放入结果集
			result[i] = row;
			i++;
		}
		defer DB.Close();
		return result

	case "witdraw":
		//查询数据，取所有字段
		rows2, _ := DB.Query("select tx_hash from depwit_status where confirm=2");
		//返回所有列
		cols, _ := rows2.Columns();
		//这里表示一行所有列的值，用[]byte表示
		vals := make([][]byte, len(cols));
		//这里表示一行填充数据
		scans := make([]interface{}, len(cols));
		//这里scans引用vals，把数据填充到[]byte里
		for k, _ := range vals {
			scans[k] = &vals[k];
		}

		i := 0;
		result := make(map[int]map[string]string);
		for rows2.Next(){
			//填充数据
			rows2.Scan(scans...);
			//每行数据
			row := make(map[string]string);
			//把vals中的数据复制到row中
			for k, v := range vals {
				key := cols[k];
				//这里把[]byte数据转成string
				row[key] = string(v);
			}
			//放入结果集
			result[i] = row;
			i++;
		}
		defer DB.Close();
		return result
	}
	//fmt.Println("connnect success")
		defer DB.Close();

		return map[int]map[string]string{}
}












func InitDB_wallet_delete(method string,tx_hash string)(int){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_wallet:=configMap["userName_wallet"]
	password_wallet:=configMap["password_wallet"]
	ip_wallet:=configMap["ip_wallet"]
	port_wallet:=configMap["port_wallet"]
	dbName_wallet:=configMap["dbName_wallet"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_wallet, ":", password_wallet, "@tcp(",ip_wallet, ":", port_wallet, ")/", dbName_wallet, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}

	switch method {
	case "select_depsoit":
		var status int
		err := DB.QueryRow("select status from deposit_xnc where tx_hash= ?", tx_hash).Scan(&status)
		if err != nil {
			//fmt.Println(err.Error())
		} else if err == sql.ErrNoRows {
			//fmt.Println("no rows in result set")
		}

		return status
	case "select_deposit_height":
		var confirm_height int
		err := DB.QueryRow("select confirm_height from deposit_xnc where tx_hash= ?", tx_hash).Scan(&confirm_height)
		if err != nil {
			//fmt.Println(err.Error())
		} else if err == sql.ErrNoRows {
			//fmt.Println("no rows in result set")
		}
		return confirm_height
	case "select_witdraw_status":
		var status int
		err := DB.QueryRow("select status from withdraw_xnc where txhash= ?", tx_hash).Scan(&status)
		if err != nil {
			//fmt.Println(err.Error())
		} else if err == sql.ErrNoRows {
			//fmt.Println("no rows in result set")
		}

		return status
	}
	//fmt.Println("connnect success")
	defer DB.Close();
	return 999
}


func InitDB_wallet_delete_1(method string,tx_hash string){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_audit:=configMap["userName_audit"]
	password_audit:=configMap["password_audit"]
	ip_audit:=configMap["ip_audit"]
	port_audit:=configMap["port_audit"]
	dbName_audit:=configMap["dbName_audit"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_audit, ":", password_audit, "@tcp(",ip_audit, ":", port_audit, ")/", dbName_audit, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}

	switch method {
	case "delete":
		res,err:=DB.Exec("delete from depwit_status where tx_hash=?",tx_hash)

		if err != nil {
			log.Fatal("delect:",err)
		}
		rowsaffected, err := res.RowsAffected()
		fmt.Println("本次共计删除数据：", rowsaffected, "行")
	}

	//fmt.Println("connnect success")
	defer DB.Close();
}


func InitDB_wallet_update(select_deposit_height int,check_tx_hash string){
	var DB *sql.DB

	//导入配置文件
	configMap := conf.InitConfig("/opt/dep_wit_status/src/config/conf.ini")
	//获取配置value
	userName_audit:=configMap["userName_audit"]
	password_audit:=configMap["password_audit"]
	ip_audit:=configMap["ip_audit"]
	port_audit:=configMap["port_audit"]
	dbName_audit:=configMap["dbName_audit"]

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName_audit, ":", password_audit, "@tcp(",ip_audit, ":", port_audit, ")/", dbName_audit, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println(DB,"opon database fail")
	}

	stmt, err := DB.Prepare("Update depwit_status SET confirm_height=? where tx_hash=?")
	res, err := stmt.Exec(select_deposit_height,check_tx_hash)
	if err != nil {
		fmt.Println("数据库执行更新出错", err)
		return
	}
	rowsaffected, err := res.RowsAffected()
	if err != nil {

	}
	fmt.Println("共计", rowsaffected, "行更新")

	//fmt.Println("connnect success")
	defer DB.Close();
}
