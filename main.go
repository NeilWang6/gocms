package main

import (
	_ "gocms/routers"
	_ "gocms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
