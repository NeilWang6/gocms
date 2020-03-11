package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cuua/gocms/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	"time"
)

var db *gorm.DB

func init() {
	ConnectDB()
}

// connectDB from database
func ConnectDB() {
	var err error
	dbType := beego.AppConfig.String("db_type")
	if dbType == "" {
		logs.Error("please specify database")
		return
	}
	db, err = gorm.Open(dbType, models.Conf())
	if err != nil {
		logs.Error("opens database failed: " + err.Error())
		return
	}
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.LogMode(!models.Runmode())
}

// 获取db
func DB() *gorm.DB {
	return db
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		logs.Error("Disconnect from database failed: " + err.Error())
	}
}
