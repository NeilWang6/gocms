package db

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("mysql::db_user"),
		beego.AppConfig.String("mysql::db_pwd"),
		beego.AppConfig.String("mysql::db_name"),
		beego.AppConfig.String("mysql::db_host"),
	)
	//config := "root:root@/db_center?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", config)
	db.LogMode(LogState())
	db.DB().SetMaxIdleConns(10)
	//db.AutoMigrate()
	//db.DB().SetMaxOpenConns(100)
	if err != nil {
		fmt.Println("DB连接失败", err.Error())
	} else {
		DB = db
	}
}

func LogState() bool {
	if beego.AppConfig.String("runmode") == "prod" {
		return false
	}
	return true
}
