package main

import (
	// 抽象包
	// 相当于php的扩展
	"database/sql"
	"fmt"
	// 实现层  有针对mysql的  sqlserver的
	// 类似于php的pdo  mysli

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8")

	stmt, _ := db.Prepare("insert into access (id,title) values (?,?)")
	res, _ := stmt.Exec("3", "aaa")
	id, _ := res.LastInsertId()

	fmt.Print(id)

}
