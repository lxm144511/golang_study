package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 创建连接数据库的句柄
	db, err := sql.Open("mysql", "root:123456@tcp(220.197.198.41:3306)/my_db?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("open database  failed ", err)
		return
	}
	db.SetMaxOpenConns(10)
	fmt.Println("connect success")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS my_db;")
	if err != nil {
		fmt.Println("failed to create databases", err.Error())
		return
	}
	_, err = db.Exec("USE my_db;")
	if err != nil {
		fmt.Println("select database failed")
		return
	}
	// 创建表
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tb_user(id int(10) primary key,name varchar(20),age int(10),sex varchar(5),addr varchar(64),tel varchar(11));")
	if err != nil {
		fmt.Println("create table  failed:", err.Error())
		return
	}
	// 插入数据
	result, err := db.Exec("INSERT INTO `tb_user`(`id`,`name`,`age`,`sex`,`addr`,`tel`) values(1025,'ALICE',20,'WOMAN','SZ','123')")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return
	}
	if RowAffect, err := result.RowsAffected(); nil == err {
		fmt.Println("RowAffect", RowAffect)
	}
	fmt.Println("insert new record", id)
}
