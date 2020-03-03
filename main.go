package main

import (
	_ "github.com/cuua/gocms/routers"
	_ "github.com/cuua/gocms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
