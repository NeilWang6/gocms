package controllers

import (
	"github.com/astaxie/beego/orm"
	"gocms/models/Student"
	"gocms/models/Teacher"
	"strings"
	"time"

	"gocms/enums"
	"gocms/models"
	"gocms/utils"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()

	c.SetTpl()
}
func (c *HomeController) Page404() {
	c.SetTpl()
}
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.SetTpl("home/error.html", "shared/layout_pullbox.html")
}
func (c *HomeController) Login() {
	c.Data["cont"] = c.ControllerName
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.SetTpl("home/login.html", "main/layout_page.html")

	//c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	//c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	//c.SetTpl("home/login.html", "shared/layout_base.html")
}
func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.JsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := models.BackendUserOneByUserName(username, userpwd)
	//fmt.Println(username,userpwd,user,err)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			c.JsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)
		//获取用户信息
		c.JsonResult(enums.JRCodeSucc, "登录成功", "")
	} else {
		c.JsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	}
}
func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.PageLogin()
}

func (c *HomeController) CkeditorUpload() {
	type ck struct {
		Uploaded int               `json:"uploaded"`
		FileName string            `json:"file_name"`
		Url      string            `json:"url"`
		Error    map[string]string `json:"error"`
	}
	var resp ck
	f, h, err := c.GetFile("upload")
	if err != nil {
		resp.Error = map[string]string{"message": err.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
		c.StopRun()
	}
	defer f.Close()
	// 保存位置在 static/upload, 没有文件夹要先创建
	file, err := utils.UploadImage(h, "news")
	if err != nil {
		resp.Error = map[string]string{"message": err.Error()}
	} else {
		resp.Uploaded = 1
		resp.Url = "/static/upload/" + file
		resp.FileName = file
	}
	c.Data["json"] = resp
	c.ServeJSON()
	c.StopRun()

}
