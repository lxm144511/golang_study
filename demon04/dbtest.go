package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "220.197.198.41"
	dbname = "my_db"
	use    = "root"
	port   = "3306"
	passwd = "123456"
)

// 连接信息拼接
var dbinfo = use + ":" + passwd + "@tcp" + "(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

func main() {
	fmt.Println(dbinfo)
	// 创建连接
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("open database failed", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("database connected failed", err)
		return
	}
	fmt.Println("database connect  success")
	// 查询数据
	res, err := db.Exec("use my_db  ")
	if err != nil {
		fmt.Println("sql exec failed", err)
		panic(err)
	}
	fmt.Println("into database success", &res)
	type User struct {
		Id   int
		Name string
		Age  int
		Sex  string
		Add  string
		Tel  string
	}
	rows, err := db.Query("select * from tb_user")
	if err != nil {
		fmt.Println("sql exec failed", err)
		panic(err)
	}
	fmt.Println(rows.Next())
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Sex, &u.Add, &u.Tel)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}

}
