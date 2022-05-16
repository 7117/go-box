package data_source

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {

	mysqlConf := LoadMysqlConf()

	logoMode := mysqlConf.LogoMode

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		mysqlConf.UserName,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DataBase,
	)

	Db, err = gorm.Open("mysql", dataSource)

	if err != nil {
		fmt.Println(err)
	}

	Db.LogMode(logoMode)

	Db.DB().SetMaxOpenConns(100) // 最大连接数
	Db.DB().SetMaxIdleConns(50)  // 最大空闲数

	//Db.AutoMigrate(&first.User{})
}
