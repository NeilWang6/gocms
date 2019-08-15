package db

import (
	"adsgoscr.xinxin.com/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.String("dsp.mysqluser"),
		c.String("dsp.mysqlpass"),
		c.String("dsp.mysqlurls"),
		c.String("dsp.mysqldb"),
	)
	//config := "root:root@/db_center?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", config)
	db.LogMode(LogState())
	db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
	if err != nil {
		fmt.Println("DB连接失败", err.Error())
	} else {
		DB = db
	}
}

func LogState() bool {
	c, _ := conf.NewFileConf("conf/app.conf")
	if c.String("env.runmode") == "prod" {
		return false
	}
	return true
}
