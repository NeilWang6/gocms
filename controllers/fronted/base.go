package fronted

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocms/enums"
	"gocms/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	ControllerName string //当前控制名称
	ActionName     string //当前action名称
}

func (c *BaseController) Prepare() {
	//附值
	c.ControllerName, c.ActionName = c.GetControllerAndAction()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) SetTpl(template ...string) {
	var tplName string
	layout := "main/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.ControllerName[0 : len(c.ControllerName)-10])
		ActionName := strings.ToLower(c.ActionName)
		tplName = ctrlName + "/" + ActionName + ".html"
	}
	c.Data["cont"] = c.ControllerName
	c.Layout = layout
	c.TplName = tplName
}

// 第一个参数模板，第二个参数为layout 移动端模板
func (c *BaseController) SetMTpl(template ...string) {
	var tplName string
	layout := "main/m/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.ControllerName[0 : len(c.ControllerName)-10])
		ActionName := strings.ToLower(c.ActionName)
		tplName = ctrlName + "/" + ActionName + ".html"
	}
	fmt.Println(tplName)
	c.Data["cont"] = c.ControllerName
	c.Layout = layout
	c.TplName = tplName
}
func (c *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) PageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}
