package routers

import (
	"github.com/astaxie/beego"
	"gocms/controllers"
)

func init() {
	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/ckeditorupload", &controllers.HomeController{}, "Post:CkeditorUpload")
}
