package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

func Runmode() bool {
	if beego.AppConfig.String("runmode") == "prod" {
		return true
	}
	return false
}

func Conf() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("mysql::db_user"),
		beego.AppConfig.String("mysql::db_pwd"),
		beego.AppConfig.String("mysql::db_host"),
		beego.AppConfig.String("mysql::db_port"),
		beego.AppConfig.String("mysql::db_name"),
	)
}
